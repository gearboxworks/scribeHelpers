package toolCobraHelp

import (
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
	"github.com/spf13/cobra"
)


type CobraGetter interface {
}

type TypeCommands struct {
	Commands Cmds

	runtime  *toolRuntime.TypeRuntime
	State    *ux.State
}
func (tc *TypeCommands) IsNil() *ux.State {
	return ux.IfNilReturnError(tc)
}

type Cmds map[string][]*cobra.Command


func New(runtime *toolRuntime.TypeRuntime) *TypeCommands {
	runtime = runtime.EnsureNotNil()

	te := TypeCommands {
		Commands:    make(Cmds),

		runtime: runtime,
		State:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
	te.State.SetPackage("")
	te.State.SetFunctionCaller()
	return &te
}
