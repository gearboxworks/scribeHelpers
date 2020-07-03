package toolGear

import "github.com/newclarity/scribeHelpers/ux"


type ToolDockerGear TypeDockerGear
func (c *ToolDockerGear) IsNil() *ux.State {
	return ux.IfNilReturnError(c)
}

func (c *ToolDockerGear) Reflect() *TypeDockerGear {
	return (*TypeDockerGear)(c)
}

func (c *TypeDockerGear) Reflect() *ToolDockerGear {
	return (*ToolDockerGear)(c)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolDockerGear(p *TypeDockerGear) *ToolDockerGear {
	return (*ToolDockerGear)(p)
}
