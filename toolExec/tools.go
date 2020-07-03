package toolExec

import (
	"github.com/newclarity/scribeHelpers/ux"
)


// Usage:
//		{{ $output := ExecCommand "ps %s" "-eaf" ... }}
func ToolExecCmd(cmd ...interface{}) *ux.State {
	ret := New(nil)

	for range onlyOnce {
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
	for range onlyOnce {
		value := ux.ReflectInt(e)
		ux.Exit(*value)
	}
	return ""
}
