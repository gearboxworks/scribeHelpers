package toolExample

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
