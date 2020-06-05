package toolSystem

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)

type SystemGetter interface {
}


type TypeSystem struct {
	Procs *TypeProcesses
	Env   *Environment

	State *ux.State
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}
func ReflectToolSystem(p *TypeSystem) *ToolSystem {
	return (*ToolSystem)(p)
}


func New(runtime *toolRuntime.TypeRuntime) *TypeSystem {
	runtime = runtime.EnsureNotNil()

	s := TypeSystem {
		Procs: NewProcesses(runtime),
		Env:   &Environment{},

		State: ux.NewState(runtime.CmdName, runtime.Debug),
	}
	s.State.SetPackage("")
	s.State.SetFunctionCaller()

	return &s
}


func (s *TypeSystem) IsNil() *ux.State {
	if state := ux.IfNilReturnError(s); state.IsError() {
		return state
	}
	s.State = s.State.EnsureNotNil()
	return s.State
}
