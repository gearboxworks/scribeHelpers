package toolGo

import (
	"github.com/newclarity/scribeHelpers/ux"
	"go/ast"
	"go/token"
	"sort"
)


func (gf *GoFiles) ShowTree() {
	for range onlyOnce {
		for _, f := range gf.files {
			ux.PrintflnBlue("\n##########")
			ux.PrintflnBlue("# Filename: %v\n", f.Path.GetPathAbs())

			locals, globals := make(map[string]int), make(map[string]int)

			v := newVisitor(f.Ast)
			ast.Walk(v, f.Ast)
			for k, v := range v.locals {
				locals[k] += v
			}
			for k, v := range v.globals {
				globals[k] += v
			}

			ux.PrintflnBlue("most common local variable names")
			printTopFive(locals)
			ux.PrintflnBlue("most common global variable names")
			printTopFive(globals)
		}
	}
}


func printTopFive(counts map[string]int) {
	type pair struct {
		s string
		n int
	}
	pairs := make([]pair, 0, len(counts))
	for s, n := range counts {
		pairs = append(pairs, pair{s, n})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].n > pairs[j].n })

	for i := 0; i < len(pairs) && i < 5; i++ {
		ux.PrintflnYellow("%6d %s", pairs[i].n, pairs[i].s)
	}
}


type visitor struct {
	pkgDecl map[*ast.GenDecl]bool
	locals  map[string]int
	globals map[string]int
}


func newVisitor(f *ast.File) visitor {
	decls := make(map[*ast.GenDecl]bool)
	for _, decl := range f.Decls {
		if v, ok := decl.(*ast.GenDecl); ok {
			decls[v] = true
		}
	}

	return visitor{
		decls,
		make(map[string]int),
		make(map[string]int),
	}
}


func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch d := n.(type) {
	case *ast.AssignStmt:
		if d.Tok != token.DEFINE {
			return v
		}
		for _, name := range d.Lhs {
			v.local(name)
		}
	case *ast.RangeStmt:
		v.local(d.Key)
		v.local(d.Value)
	case *ast.FuncDecl:
		if d.Recv != nil {
			v.localList(d.Recv.List)
		}
		v.localList(d.Type.Params.List)
		if d.Type.Results != nil {
			v.localList(d.Type.Results.List)
		}
	case *ast.GenDecl:
		if d.Tok != token.VAR {
			return v
		}
		for _, spec := range d.Specs {
			if value, ok := spec.(*ast.ValueSpec); ok {
				for _, name := range value.Names {
					if name.Name == "_" {
						continue
					}
					if v.pkgDecl[d] {
						v.globals[name.Name]++
					} else {
						v.locals[name.Name]++
					}
				}
			}
		}
	}

	return v
}


func (v visitor) local(n ast.Node) {
	ident, ok := n.(*ast.Ident)
	if !ok {
		return
	}
	if ident.Name == "_" || ident.Name == "" {
		return
	}
	if ident.Obj != nil && ident.Obj.Pos() == ident.Pos() {
		v.locals[ident.Name]++
	}
}


func (v visitor) localList(fs []*ast.Field) {
	for _, f := range fs {
		for _, name := range f.Names {
			v.local(name)
		}
	}
}
