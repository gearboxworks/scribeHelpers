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
	Files   GoFiles

	goFiles *toolPath.TypeOsPaths
	fset    *token.FileSet

	runtime *toolRuntime.TypeRuntime
	State   *ux.State
}


type State ux.State
func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}
func ReflectToolGo(e *TypeGo) *ToolGo {
	return (*ToolGo)(e)
}


func New(rt *toolRuntime.TypeRuntime) *TypeGo {
	rt = rt.EnsureNotNil()

	g := &TypeGo{
		Files:   GoFiles{},

		goFiles: toolPath.NewPaths(rt),
		fset:    token.NewFileSet(),

		runtime: rt,
		State:   ux.NewState(rt.CmdName, rt.Debug),
	}
	g.State.SetPackage("")
	g.State.SetFunctionCaller()
	return g
}


func (g *TypeGo) IsNil() *ux.State {
	if state := ux.IfNilReturnError(g); state.IsError() {
		return state
	}
	g.State = g.State.EnsureNotNil()
	return g.State
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

		g.Files = GoFiles{}
		for _, f := range g.goFiles.Paths {
			gf := NewGoFile(g.runtime, g.fset, f)
			if gf.state.IsNotOk() {
				g.State = gf.state
			}
			g.Files = append(g.Files, gf)
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

		for i, _ := range g.Files {
			//g.Files[i].Parse(f.Path.GetPath(), mode[0])
			g.Files[i].Parse(mode[0])
		}
	}

	return g.State
}

func (g *TypeGo) Count() int {
	if state := g.IsNil(); state.IsError() {
		return 0
	}
	return len(g.Files)
}

func (g *TypeGo) String() string {
	if state := g.IsNil(); state.IsError() {
		return state.SprintError()
	}
	var ret string

	for range onlyOnce {
		for _, f := range g.Files {
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

		if len(g.Files) == 0 {
			break
		}

		ret = g.Files[0].Ast.Name.Name
	}

	return ret
}
