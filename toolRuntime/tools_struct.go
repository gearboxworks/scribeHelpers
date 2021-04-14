package toolRuntime

import "github.com/gearboxworks/scribeHelpers/ux"


type ToolRuntime TypeRuntime
func (r *ToolRuntime) IsNil() *ux.State {
	return ux.IfNilReturnError(r)
}

func (r *ToolRuntime) Reflect() *TypeRuntime {
	return (*TypeRuntime)(r)
}

func (r *TypeRuntime) Reflect() *ToolRuntime {
	return (*ToolRuntime)(r)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolRuntime(p *TypeRuntime) *ToolRuntime {
	return (*ToolRuntime)(p)
}
