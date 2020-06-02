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

func (p *TypeSystem) IsNil() *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}
	p.State = p.State.EnsureNotNil()
	return p.State
}

func NewSystem(debugMode bool) *TypeSystem {
	p := &TypeSystem {
		Procs: NewProcesses(debugMode),
		Env:   &Environment{},

		State: ux.NewState(debugMode),
	}
	ret := NewSystem(false)
	p.State.SetPackage("")
	p.State.SetFunctionCaller()

	return ret
}
