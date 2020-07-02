package toolCrawler

import (
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)

type ExampleGetter interface {
}


type TypeExample struct {
	name    string
	path    *toolPath.TypeOsPath

	runtime *toolRuntime.TypeRuntime
	State   *ux.State
}


func New(runtime *toolRuntime.TypeRuntime) *TypeExample {
	runtime = runtime.EnsureNotNil()

	te := TypeExample{
		name:    "",
		path:    toolPath.New(runtime),

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
func ReflectToolExample(e *TypeExample) *ToolExample {
	return (*ToolExample)(e)
}

func (e *TypeExample) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}

