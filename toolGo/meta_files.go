package toolGo

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/fatih/astrewrite"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
)


func (g *TypeGo) ShowTree() string {
	if state := g.IsNil(); state.IsError() {
		return state.SprintError()
	}
	var ret string

	for range onlyOnce {
		g.Go.ShowTree()
	}

	return ret
}


type GoFile struct {
	Path    *toolPath.TypeOsPath
	Ast     *ast.File
	fset    *token.FileSet

	meta    *GoMeta
	runtime *toolRuntime.TypeRuntime
	State   *ux.State
}
func (gf *GoFile) IsNil() *ux.State {
	return ux.IfNilReturnError(gf)
}


type GoFiles struct {
	files []*GoFile
	Found *GoFile
}

type Mode uint

const (
	PackageClauseOnly Mode             = 1 << iota // stop parsing after package clause
	ImportsOnly                                    // stop parsing after import declarations
	ParseComments                                  // parse comments and add them to AST
	Trace                                          // print a trace of parsed productions
	DeclarationErrors                              // report declaration errors
	SpuriousErrors                                 // same as AllErrors, for backward-compatibility
	AllErrors         = SpuriousErrors             // report all errors (not just the first 10 on different lines)
)


func NewGoFile(rt *toolRuntime.TypeRuntime, fset *token.FileSet, path *toolPath.TypeOsPath) *GoFile {
	rt = rt.EnsureNotNil()

	g := GoFile{
		Path: path,
		Ast:  &ast.File{},
		fset: fset,

		meta:    &GoMeta{
			binaryName:    Name{},
			binaryVersion: Version{},
			sourceRepo:    Repo{},
			binaryRepo:    Repo{},
			Valid:         false,
			//State:         ux.NewState(rt.CmdName, rt.Debug),
		},
		runtime: rt,
		State:   ux.NewState(rt.CmdName, rt.Debug),
	}
	g.State.SetPackage("")
	g.State.SetFunctionCaller()
	//fset.AddFile(path.GetPath(), -1)
	return &g
}


func (gf *GoFile) Parse(mode Mode) *ux.State {
	for range onlyOnce {
		var err error

		if gf.runtime.Debug {
			ux.PrintflnBlue("Scan file => '%s' abs('%s')", gf.Path.GetPath(), gf.Path.GetPathAbs())
		}

		gf.Ast, err = parser.ParseFile(gf.fset, gf.Path.GetPath(), nil, parser.Mode(mode))
		if err != nil {
			gf.State.SetError(err)
			break
		}

		if gf.Ast.Name == nil {
			err = errors.New("not a goLang file")
			break
		}

		gf.meta = gf.GetMeta()
	}

	return gf.State
}


func (gf *GoFile) String() string {
	var ret string

	for range onlyOnce {
		ret += fmt.Sprintf("\n##########\n")
		ret += fmt.Sprintf("# Filename: %v\n\n", gf.Path.GetPathAbs())
		ret += fmt.Sprintf("Name: %v\n", gf.Ast.Name)
		ret += fmt.Sprintf("Package Pos: %v\n", gf.Ast.Package)
		ret += fmt.Sprintf("Comments: %v\n", gf.Ast.Comments)
		ret += fmt.Sprintf("Decls: %v\n", gf.Ast.Decls)
		ret += fmt.Sprintf("Doc: %v\n", gf.Ast.Doc.Text())
		ret += fmt.Sprintf("Scope: %v\n", gf.Ast.Scope.String())
		ret += fmt.Sprintf("Unresolved: %v\n", gf.Ast.Unresolved)
		ret += fmt.Sprintf("Imports: %v\n", gf.Ast.Imports)
	}

	return ret
}


func (gf *GoFile) GetMeta() *GoMeta {
	for range onlyOnce {
		var ok bool
		for _, object := range gf.Ast.Decls {
			if ok {
				break
			}

			//ux.PrintflnBlue("%s => %v", name, object)
			switch decl := object.(type) {
				case *ast.FuncDecl:
					//ux.PrintflnBlue("Func")
				case *ast.GenDecl:
					for _, spec := range decl.Specs {
						ok = gf.getGenDecl(spec)
					}
				default:
					ux.PrintflnBlue("Unknown declaration @\n", decl.Pos())
			}
		}
		if ok {
			gf.meta.Valid = true
			break
		}
	}

	return gf.meta
}


func (gf *GoFile) UpdateMeta(name string, value string) *ux.State {
	for range onlyOnce {
		rewriteFunc := func(node ast.Node) (ast.Node, bool) {
			spec, ok := node.(*ast.ValueSpec)
			if !ok {
				return node, true
			}

			for _, id := range spec.Names {
				n, v := getValueSpec(id)
				if n != name {
					continue
				}

				//id.Name = "Hey" + BinaryVersion
				expr1 := id.Obj.Decl.(*ast.ValueSpec).Values[0]
				expr := spec.Values[0]
				oldValue := getExpr(expr1)
				ux.PrintflnBlue("Modify src file '%s'", gf.Path.GetPathAbs())
				ux.PrintflnBlue("\tChanging '%s' FROM '%s'/'%s' TO '%s'", name, v, oldValue, value)
				setExpr(expr, value)
			}

			// change struct type name to "Bar"
			//spec.Name.Name = "Bar"
			return spec, true
		}

		rewritten := astrewrite.Walk(gf.Ast, rewriteFunc)

		var buf bytes.Buffer
		err := printer.Fprint(&buf, gf.fset, rewritten)
		if err != nil {
			gf.State.SetError(err)
			break
		}

		ux.PrintflnBlue("New file will be written as:")
		ux.PrintfBlue(buf.String())
		//if gf.runtime.Debug {
		//	fmt.Println(buf.String())
		//}

		gf.Path.SetContents(&buf)
		gf.Path.SetOverwriteable()
		gf.State = gf.Path.WriteFile()
		if gf.State.IsNotOk() {
			break
		}
	}

	return gf.State
}


func (gf *GoFile) getGenDecl(spec ast.Spec) bool {
	var ok bool
	switch spec := spec.(type) {
		case *ast.ImportSpec:
			//ux.PrintflnBlue("Import", spec.Path.Value)
		case *ast.TypeSpec:
			//ux.PrintflnBlue("Type", spec.Name.String())
		case *ast.ValueSpec:
			for _, id := range spec.Names {
				name, value := getValueSpec(id)
				if name == "" {
					continue
				}
				err := gf.meta.setValue(name, value)
				if err == nil {
					ok = true
				}
			}
		default:
			ux.PrintflnBlue("Unknown token type")	// : %s", decl.Tok)
	}
	return ok
}


func getValueSpec(id *ast.Ident) (string, string) {
	var name string
	var value string
	switch id.Name {
		case BinaryName:
			fallthrough
		case BinaryVersion:
			fallthrough
		case SourceRepo:
			fallthrough
		case BinaryRepo:
			name = id.Name
	}

	expr := id.Obj.Decl.(*ast.ValueSpec).Values[0]
	value = getExpr(expr)
	//fmt.Printf("%s => %s", name, value)
	//fmt.Printf("\n")

	//ux.PrintflnBlue("Var %s: %v", id.Name, id.Obj.Decl.(*ast.ValueSpec).Values[0].(*ast.BasicLit).Value)
	return name, value
}


func getExpr(expr ast.Expr) string {
	var ret string
	switch expr.(type) {
		case *ast.BasicLit:
			ret = expr.(*ast.BasicLit).Value
			ret = strings.TrimPrefix(ret, "\"")
			ret = strings.TrimSuffix(ret, "\"")
		case *ast.BinaryExpr:
			ret = getExpr(expr.(*ast.BinaryExpr).X) + getExpr(expr.(*ast.BinaryExpr).Y)
		case *ast.Ident:
			_, ret = getValueSpec(expr.(*ast.Ident))
			//id := expr.(*ast.BinaryExpr).Y.(*ast.Object).Decl
			//switch id.(type) {
			//	case *ast.ValueSpec:
			//		_, ret = getValueSpec(id.(*ast.ValueSpec))
			//}
	}
	return ret
}


func setExpr(expr ast.Expr, value string) {
	switch expr.(type) {
		case *ast.BasicLit:
			value = "\"" + value + "\""
			expr.(*ast.BasicLit).Value = value
		case *ast.BinaryExpr:
			setExpr(expr.(*ast.BinaryExpr).X, value)
		case *ast.Ident:
			value = "\"" + value + "\""
			expr.(*ast.Ident).Name = value
	}
}
