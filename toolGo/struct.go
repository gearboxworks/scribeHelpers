package toolGo

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"go/token"
)


/*
@TODO - Consider using parser.ParseDir instead.
*/


type GoGetter interface {
}


type TypeGo struct {
	Go      GoFiles

	goFiles *toolPath.TypeOsPaths
	fset    *token.FileSet

	runtime *toolRuntime.TypeRuntime
	State   *ux.State
}
func (g *TypeGo) IsNil() *ux.State {
	return ux.IfNilReturnError(g)
}


func New(rt *toolRuntime.TypeRuntime) *TypeGo {
	rt = rt.EnsureNotNil()

	g := &TypeGo{
		Go: GoFiles{},

		goFiles: toolPath.NewPaths(rt),
		fset:    token.NewFileSet(),

		runtime: rt,
		State:   ux.NewState(rt.CmdName, rt.Debug),
	}
	g.State.SetPackage("")
	g.State.SetFunctionCaller()
	return g
}

func (g *TypeGo) SetPath(path ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	return g.goFiles.SetBasePath(path...)
}

func (g *TypeGo) GetPath() string {
	if state := g.IsNil(); state.IsError() {
		return ""
	}
	return g.goFiles.Base.GetPath()
}

func (g *TypeGo) SetRecursive() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	return g.goFiles.SetRecursive()
}

func (g *TypeGo) SetNonRecursive() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	return g.goFiles.SetNonRecursive()
}

func (g *TypeGo) Find(path ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		g.State = g.goFiles.FindByExt("go", path...)
		if g.State.IsNotOk() {
			break
		}

		g.Go = GoFiles{}
		for _, f := range g.goFiles.Paths {
			gf := NewGoFile(g.runtime, g.fset, f)
			if gf.State.IsNotOk() {
				g.State = gf.State
			}
			g.Go.files = append(g.Go.files, gf)
		}
	}

	return g.State
}

func (g *TypeGo) Parse(mode ...Mode) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if len(mode) == 0 {
			mode = append(mode, Mode(0))
		}

		for _, f := range g.Go.files {
			//g.Go[i].Parse(f.Path.GetPath(), mode[0])
			g.State = f.Parse(mode[0])
			if f.meta.Valid {
				g.Go.Found = f
				break
			}
		}
	}

	return g.State
}

func (g *TypeGo) Count() int {
	if state := g.IsNil(); state.IsError() {
		return 0
	}
	return len(g.Go.files)
}

func (g *TypeGo) String() string {
	if state := g.IsNil(); state.IsError() {
		return state.SprintError()
	}
	var ret string

	for range onlyOnce {
		for _, f := range g.Go.files {
			ret += fmt.Sprintf("%v", f.String())
		}
	}

	return ret
}

func (g *TypeGo) GetPackageName(path ...string) string {
	if state := g.IsNil(); state.IsError() {
		return state.SprintError()
	}
	var ret string

	for range onlyOnce {
		if len(path) == 0 {
			break
		}

		if g.State.IsNotOk() {
			break
		}

		g.State = g.SetNonRecursive()
		if g.State.IsNotOk() {
			break
		}

		g.State = g.Find(path...)
		if g.State.IsNotOk() {
			break
		}

		g.State = g.Parse()
		if g.State.IsNotOk() {
			break
		}

		if len(g.Go.files) == 0 {
			break
		}

		ret = g.Go.files[0].Ast.Name.Name
	}

	return ret
}

func (gf *GoFiles) GetMeta() *GoMeta {
	var ret *GoMeta

	for range onlyOnce {
		for _, f := range gf.files {
			//ux.PrintflnBlue("# %s", f.Path.GetPath())
			//ux.PrintflnCyan("%v", f.Ast.Name)
			ret = f.GetMeta()
			if ret != nil {
				break
			}
		}
	}

	return ret
}
