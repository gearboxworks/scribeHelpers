package helperSystem

import (
	"github.com/newclarity/scribeHelpers/helperRuntime"
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
func ReflectHelperSystem(p *TypeSystem) *HelperSystem {
	return (*HelperSystem)(p)
}


func New(runtime *helperRuntime.TypeRuntime) *TypeSystem {
	runtime = runtime.EnsureNotNil()

	s := TypeSystem {
		Procs: NewProcesses(runtime.Debug),
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
