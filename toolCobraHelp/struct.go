package toolCobraHelp

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
)

type CobraGetter interface {
}

type TypeCommands struct {
	Commands Cmds

	runtime  *toolRuntime.TypeRuntime
	State    *ux.State
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


type State ux.State
func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}
func ReflectToolCobra(e *TypeCommands) *ToolCobra {
	return (*ToolCobra)(e)
}

func (tc *TypeCommands) IsNil() *ux.State {
	if state := ux.IfNilReturnError(tc); state.IsError() {
		return state
	}
	tc.State = tc.State.EnsureNotNil()
	return tc.State
}

