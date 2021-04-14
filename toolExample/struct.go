package toolExample

import (
	"github.com/gearboxworks/scribeHelpers/toolPath"
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
)

type ExampleGetter interface {
}


type TypeExample struct {
	name    string
	path    *toolPath.TypeOsPath

	runtime *toolRuntime.TypeRuntime
	State   *ux.State
}
func (e *TypeExample) IsNil() *ux.State {
	return ux.IfNilReturnError(e)
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
