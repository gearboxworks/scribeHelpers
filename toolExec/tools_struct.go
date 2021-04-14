package toolExec

import "github.com/gearboxworks/scribeHelpers/ux"


type ToolExecCommand TypeExecCommand
func (e *ToolExecCommand) IsNil() *ux.State {
	return ux.IfNilReturnError(e)
}

func (e *ToolExecCommand) Reflect() *TypeExecCommand {
	return (*TypeExecCommand)(e)
}

func (e *TypeExecCommand) Reflect() *ToolExecCommand {
	return (*ToolExecCommand)(e)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolExecCommand(p *TypeExecCommand) *ToolExecCommand {
	return (*ToolExecCommand)(p)
}
