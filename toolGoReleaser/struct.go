package toolGoReleaser

import (
	"github.com/newclarity/scribeHelpers/toolExec"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)

type GoReleaserGetter interface {
}


type TypeGoReleaser struct {
	exec    *toolExec.TypeExecCommand
	path    *toolPath.TypeOsPath

	runtime *toolRuntime.TypeRuntime
	State   *ux.State
}
func (gr *TypeGoReleaser) IsNil() *ux.State {
	return ux.IfNilReturnError(gr)
}


func New(runtime *toolRuntime.TypeRuntime) *TypeGoReleaser {
	var te TypeGoReleaser

	for range onlyOnce {
		runtime = runtime.EnsureNotNil()

		te.exec		= toolExec.New(runtime)
		te.path		= toolPath.New(runtime)

		te.runtime	= runtime
		te.State	= ux.NewState(runtime.CmdName, runtime.Debug)
		te.State.SetPackage("")
		te.State.SetFunctionCaller()

		te.State = te.exec.SetCmd(DefaultCmd)
		if te.State.IsNotOk() {
			break
		}
	}

	return &te
}

func (gr *TypeGoReleaser) IsRunnable() bool {
	return gr.exec.IsRunnable()
}

func (gr *TypeGoReleaser) ShowProgress() {
	gr.exec.ShowProgress()
}

func (gr *TypeGoReleaser) SilenceProgress() {
	gr.exec.SilenceProgress()
}
