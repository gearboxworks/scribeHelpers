package toolExec

import (
	"github.com/newclarity/scribeHelpers/ux"
)


type ToolExecCommand TypeExecCommand
func (g *ToolExecCommand) Reflect() *TypeExecCommand {
	return (*TypeExecCommand)(g)
}
func (e *TypeExecCommand) Reflect() *ToolExecCommand {
	return (*ToolExecCommand)(e)
}

func (c *ToolExecCommand) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}


// Usage:
//		{{ $output := ExecCommand "ps %s" "-eaf" ... }}
func ToolExecCmd(cmd ...interface{}) *ux.State {
	ret := New(false)

	for range OnlyOnce {
		ec := ReflectExecCommand(cmd...)
		if ec == nil {
			break
		}
		ec.ShowProgress()
		ret.State = ec.Run()
	}

	return ret.State
}
// Alias of ExecCommand
func ToolExec(cmd ...interface{}) *ux.State {
	return ToolExecCmd(cmd...)
}


// Usage:
//		{{ $cmd := ExecCommand "ps %s" "-eaf" ... }}
//		{{ $cmd.PrintError }}
func (e *TypeExecCommand) PrintError() string {
	return e.State.SprintError()
}


// Usage:
//		{{ $cmd.ExitOnError }}
func (e *TypeExecCommand) ExitOnError() string {
	e.State.ExitOnError()
	return ""
}


// Usage:
//		{{ $cmd.ExitOnWarning }}
func (e *TypeExecCommand) ExitOnWarning() string {
	e.State.ExitOnWarning()
	return ""
}


// Usage:
//		{{ OsExit 1 }}
func ToolOsExit(e ...interface{}) string {
	for range OnlyOnce {
		value := ux.ReflectInt(e)
		ux.Exit(*value)
	}
	return ""
}
