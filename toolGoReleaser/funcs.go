package toolGoReleaser

import (
	"github.com/newclarity/scribeHelpers/toolExec"
	"github.com/newclarity/scribeHelpers/ux"
)


func (gr *TypeGoReleaser) SetBasePath(path ...string) *ux.State {
	if state := gr.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if len(path) == 0 {
			//p.State.SetError("no path specified")
			break
		}

		if !gr.path.SetPath(path...) {
			gr.State.SetError("no path specified")
			break
		}

		gr.State = gr.path.StatPath()
		if gr.State.IsNotOk() {
			break
		}
	}

	return gr.State
}


func (gr *TypeGoReleaser) Release() *ux.State {
	if state := gr.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {

		grFile := NewArgFile(gr.Debug)
		gr.State = grFile.SetPath(DefaultFile)
		if grFile.NotExists() {
			gr.State = grFile.State
			break
		}

		ux.PrintflnBlue("Found goreleaser file: %s", DefaultFile)
		gr.State = exe.Exec("goreleaser", "--rm-dist")
		if gr.State.IsNotOk() {
			ux.PrintflnWarning("Error with goreleaser.")
			break
		}
	}

	return gr.State
}


func (gr *TypeGoReleaser) Build() *ux.State {
	if state := gr.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		grFile := NewArgFile(gr.runtime)
		gr.State = grFile.SetPath(DefaultFile)
		if grFile.NotExists() {
			gr.State = grFile.State
			break
		}

		exe := toolExec.New(gr.runtime)
		exe.ShowProgress()

		ux.PrintflnBlue("Found goreleaser file: %s", DefaultFile)
		gr.State = exe.Exec("goreleaser", "--snapshot", "--skip-publish", "--rm-dist")
		if gr.State.IsNotOk() {
			//ux.PrintflnWarning("goreleaser failed to build.")
			break
		}
	}

	return gr.State
}
