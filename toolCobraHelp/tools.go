package toolCobraHelp

import "github.com/newclarity/scribeHelpers/ux"


type ToolCobra TypeCommands
func (e *ToolCobra) Reflect() *TypeCommands {
	return (*TypeCommands)(e)
}
func (tc *TypeCommands) Reflect() *ToolCobra {
	return (*ToolCobra)(tc)
}

func (e *ToolCobra) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}
