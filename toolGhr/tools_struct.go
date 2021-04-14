package toolGhr

import "github.com/gearboxworks/scribeHelpers/ux"


type ToolGhr TypeGhr
func (c *ToolGhr) IsNil() *ux.State {
	return ux.IfNilReturnError(c)
}

func (c *ToolGhr) Reflect() *TypeGhr {
	return (*TypeGhr)(c)
}

func (c *TypeGhr) Reflect() *ToolGhr {
	return (*ToolGhr)(c)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolGhr(p *TypeGhr) *ToolGhr {
	return (*ToolGhr)(p)
}
