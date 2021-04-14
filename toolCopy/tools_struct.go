package toolCopy

import "github.com/gearboxworks/scribeHelpers/ux"


type ToolOsCopy TypeOsCopy
func (c *ToolOsCopy) IsNil() *ux.State {
	return ux.IfNilReturnError(c)
}

func (c *ToolOsCopy) Reflect() *TypeOsCopy {
	return (*TypeOsCopy)(c)
}

func (c *TypeOsCopy) Reflect() *ToolOsCopy {
	return (*ToolOsCopy)(c)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolOsCopy(p *TypeOsCopy) *ToolOsCopy {
	return (*ToolOsCopy)(p)
}
