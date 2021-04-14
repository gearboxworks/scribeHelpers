package toolGo

import "github.com/gearboxworks/scribeHelpers/ux"


type ToolGo TypeGo
func (g *ToolGo) IsNil() *ux.State {
	return ux.IfNilReturnError(g)
}

func (g *ToolGo) Reflect() *TypeGo {
	return (*TypeGo)(g)
}

func (g *TypeGo) Reflect() *ToolGo {
	return (*ToolGo)(g)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolGo(p *TypeGo) *ToolGo {
	return (*ToolGo)(p)
}
