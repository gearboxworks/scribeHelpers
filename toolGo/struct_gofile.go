package toolGo

import (
	"errors"
	"fmt"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"go/ast"
	"go/parser"
	"go/token"
)


func (g *TypeGo) ShowTree() string {
	if state := g.IsNil(); state.IsError() {
		return state.SprintError()
	}
	var ret string

	for range onlyOnce {
		g.Files.ShowTree()
	}

	return ret
}


type GoFile struct {
	Path    *toolPath.TypeOsPath
	Ast     *ast.File
	fset    *token.FileSet

	meta    *GoMeta
	runtime *toolRuntime.TypeRuntime
	state   *ux.State
}
type GoFiles []*GoFile

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
		state:   ux.NewState(rt.CmdName, rt.Debug),
	}
	g.state.SetPackage("")
	g.state.SetFunctionCaller()
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
			gf.state.SetError(err)
			break
		}

		if gf.Ast.Name == nil {
			err = errors.New("not a goLang file")
		}
	}

	return gf.state
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
