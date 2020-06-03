package helperSystem

import (
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


func New(debugMode bool) *TypeSystem {
	ret := &TypeSystem {
		Procs: NewProcesses(debugMode),
		Env:   &Environment{},

		State: ux.NewState(debugMode),
	}
	ret.State.SetPackage("")
	ret.State.SetFunctionCaller()

	return ret
}


func (s *TypeSystem) IsNil() *ux.State {
	if state := ux.IfNilReturnError(s); state.IsError() {
		return state
	}
	s.State = s.State.EnsureNotNil()
	return s.State
}
