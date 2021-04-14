package toolNetwork

import "github.com/gearboxworks/scribeHelpers/ux"


type ToolNetwork TypeNetwork
func (r *ToolNetwork) IsNil() *ux.State {
	return ux.IfNilReturnError(r)
}

func (r *ToolNetwork) Reflect() *TypeNetwork {
	return (*TypeNetwork)(r)
}

func (r *TypeNetwork) Reflect() *ToolNetwork {
	return (*ToolNetwork)(r)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolNetwork(p *TypeNetwork) *ToolNetwork {
	return (*ToolNetwork)(p)
}
