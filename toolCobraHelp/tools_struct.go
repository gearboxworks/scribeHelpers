package toolCobraHelp

import "github.com/newclarity/scribeHelpers/ux"


type ToolCobra TypeCommands
func (c *ToolCobra) IsNil() *ux.State {
	return ux.IfNilReturnError(c)
}

func (c *ToolCobra) Reflect() *TypeCommands {
	return (*TypeCommands)(c)
}

func (tc *TypeCommands) Reflect() *ToolCobra {
	return (*ToolCobra)(tc)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolCobra(e *TypeCommands) *ToolCobra {
	return (*ToolCobra)(e)
}
