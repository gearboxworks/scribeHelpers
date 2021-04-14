package toolSystem

import (
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
)

type SystemGetter interface {
}


type TypeSystem struct {
	Procs *TypeProcesses
	Env   *Environment

	State *ux.State
}
func (s *TypeSystem) IsNil() *ux.State {
	return ux.IfNilReturnError(s)
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
