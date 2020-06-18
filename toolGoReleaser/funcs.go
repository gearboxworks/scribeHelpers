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


func (gr *TypeGoReleaser) Release(path ...string) *ux.State {
	if state := gr.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		ux.PrintflnBlue("Generating release with GoReleaser...")

		e := toolExec.NewMultiExec(gr.runtime)
		if e.State.IsNotOk() {
			gr.State = e.State
			break
		}

		gr.State = e.Set("goreleaser", "--rm-dist")
		if gr.State.IsNotOk() {
			break
		}

		gr.State = e.SetDontAppendFile()
		if gr.State.IsNotOk() {
			break
		}

		gr.State = e.SetChdir()
		if gr.State.IsNotOk() {
			break
		}

		gr.State = e.ShowProgress()
		if gr.State.IsNotOk() {
			break
		}

		gr.State = e.FindRegex(DefaultFile, path...)
		if gr.State.IsNotOk() {
			break
		}

		p := e.GetPaths()
		ux.PrintflnBlue("Releasing with GoReleaser in %d paths...", len(p))

		gr.State = e.Run()
		if gr.State.IsNotOk() {
			break
		}

		gr.State.SetOk("go module update OK")
	}

	return gr.State
}


func (gr *TypeGoReleaser) Build(recurse bool, path ...string) *ux.State {
	if state := gr.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		ux.PrintflnBlue("Building with GoReleaser...")

		e := toolExec.NewMultiExec(gr.runtime)
		if e.State.IsNotOk() {
			gr.State = e.State
			break
		}

		if recurse {
			e.Paths.SetRecursive()
		} else {
			e.Paths.SetNonRecursive()
		}

		gr.State = e.Set("goreleaser", "--snapshot", "--skip-publish", "--rm-dist")
		if gr.State.IsNotOk() {
			break
		}

		gr.State = e.SetDontAppendFile()
		if gr.State.IsNotOk() {
			break
		}

		gr.State = e.SetChdir()
		if gr.State.IsNotOk() {
			break
		}

		gr.State = e.ShowProgress()
		if gr.State.IsNotOk() {
			break
		}

		gr.State = e.FindRegex(DefaultFile, path...)
		if gr.State.IsNotOk() {
			break
		}

		p := e.GetPaths()
		if len(p) == 0 {
			ux.PrintflnYellow("No '%s' files found. Aborting.", DefaultFile)
			break
		}

		ux.PrintflnBlue("Building with GoReleaser in %d paths...", len(p))
		for _, pp := range p {
			ux.PrintflnBlue("%s => %s", pp.GetDirname(), pp.GetFilename())
		}

		gr.State = e.Run()
		if gr.State.IsNotOk() {
			break
		}

		gr.State.SetOk("GoReleaser build OK")
	}

	return gr.State
}
