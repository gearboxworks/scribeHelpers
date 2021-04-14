package toolExec

import (
	"github.com/gearboxworks/scribeHelpers/toolPath"
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
)


type TypeMultiExecCommand struct {
	Exec    *TypeExecCommand
	Paths   *toolPath.TypeOsPaths

	chDir   bool
	noAddFile bool

	Runtime *toolRuntime.TypeRuntime
	State   *ux.State
}


func NewMultiExec(runtime *toolRuntime.TypeRuntime) *TypeMultiExecCommand {
	runtime = runtime.EnsureNotNil()

	ret := &TypeMultiExecCommand {
		Exec:    New(runtime),
		Paths:   toolPath.NewPaths(runtime),

		Runtime: runtime,
		State:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
	ret.State.SetPackage("")
	ret.State.SetFunctionCaller()
	return ret
}


func (e *TypeMultiExecCommand) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}


func (e *TypeMultiExecCommand) ShowProgress() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	e.Exec.ShowProgress()

	return e.State
}


func (e *TypeMultiExecCommand) SilenceProgress() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	e.Exec.SilenceProgress()

	return e.State
}


func (e *TypeMultiExecCommand) IsRunnable() bool {
	if state := e.IsNil(); state.IsError() {
		return false
	}
	return e.Exec.IsRunnable()
}


func (e *TypeMultiExecCommand) Set(cmd string, path ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		e.State = e.Exec.SetCmd(cmd)
		if e.State.IsNotOk() {
			break
		}

		e.State = e.Exec.SetArgs(path...)
		if e.State.IsNotOk() {
			break
		}
	}

	return e.State
}


func (e *TypeMultiExecCommand) SetCmd(path ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		e.State = e.Exec.SetCmd(path...)
		if e.State.IsNotOk() {
			break
		}
	}

	return e.State
}


func (e *TypeMultiExecCommand) SetArgs(args ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		e.State = e.Exec.SetArgs(args...)
		if e.State.IsNotOk() {
			break
		}
	}

	return e.State
}


func (e *TypeMultiExecCommand) SetBasePath(path ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		e.State = e.Paths.SetBasePath(path...)
		if e.State.IsNotOk() {
			break
		}
	}

	return e.State
}


func (e *TypeMultiExecCommand) Find(path ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		e.Paths.Find(path...)
		if e.State.IsNotOk() {
			break
		}
	}

	return e.State
}


func (e *TypeMultiExecCommand) FindRegex(re string, path ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		e.Paths.FindRegex(re, path...)
		if e.State.IsNotOk() {
			break
		}
	}

	return e.State
}


func (e *TypeMultiExecCommand) SetDontAppendFile() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}
	e.noAddFile = true
	return e.State
}


func (e *TypeMultiExecCommand) SetChdir() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}
	e.chDir = true
	return e.State
}


func (e *TypeMultiExecCommand) GetPaths() []*toolPath.TypeOsPath {
	if state := e.IsNil(); state.IsError() {
		return []*toolPath.TypeOsPath{}
	}
	return e.Paths.Paths
}


func (e *TypeMultiExecCommand) Run() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if e.Paths.GetLength() == 0 {
			e.State.SetWarning("no files to operate on")
			break
		}

		var saveOutput []string
		saveArgs := e.Exec.GetArgs()
		for _, p := range e.Paths.Paths {
			if e.chDir {
				e.State = p.Chdir()
				if e.State.IsNotOk() {
					break
				}
			}

			var args []string
			if !e.noAddFile {
				args = appendArgs(p.GetPath(), saveArgs...)
			} else {
				args = saveArgs
			}

			e.State = e.Exec.SetArgs(args...)
			if e.State.IsNotOk() {
				break
			}

			e.State = e.Exec.Run()
			if e.State.IsNotOk() {
				break
			}

			//o2 := e.Exec.State.GetOutputArray()
			o2 := e.State.GetOutputArray()
			saveOutput = append(saveOutput, o2...)
			//fmt.Printf("%v\n", len(saveOutput))
		}

		e.State.SetOutput(saveOutput)
	}

	return e.State
}


func appendArgs(file string, args ...string) []string {
	for range onlyOnce {
		// If no args, append file to args.
		if len(args) == 0 {
			args = []string{ file }
			break
		}

		// Search for {} in args, if found, replace with file.
		var foundReplace bool
		for i, a := range args {
			if a == "{}" {
				foundReplace = true
				args[i] = file
				break
			}
		}
		if foundReplace {
			break
		}

		// Else just append file to args.
		args = append(args, file)
	}

	return args
}
