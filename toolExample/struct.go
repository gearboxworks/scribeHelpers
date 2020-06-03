package toolExample

import (
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)

type ExampleGetter interface {
}


type TypeExample struct {
	name  string
	path  *toolPath.TypeOsPath

	State *ux.State
}


func New(runtime *toolRuntime.TypeRuntime) *TypeExample {
	runtime = runtime.EnsureNotNil()

	te := TypeExample{
		name:  "",
		path:  toolPath.New(runtime),

		State: ux.NewState(runtime.CmdName, runtime.Debug),
	}
	te.State.SetPackage("")
	te.State.SetFunctionCaller()
	return &te
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}
func ReflectHelperExample(p *TypeExample) *HelperExample {
	return (*HelperExample)(p)
}

func (c *TypeExample) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}

