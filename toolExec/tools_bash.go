package toolExec

import (
	"github.com/newclarity/scribeHelpers/toolTypes"
	"github.com/newclarity/scribeHelpers/ux"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)


func ToolExecBash(cmd ...interface{}) *ux.State {
	ret := New(nil)

	for range onlyOnce {
		a := toolTypes.ReflectStrings(cmd...)

		bash, err := exec.LookPath("bash")
		if err != nil {
			ret.State.SetError("Executable not found.")
			break
		}
		ret.cmd.SetPath(bash)

		ret.args = []string{"-c"}
		ret.args = append(ret.args, *a...)
		ret.ShowProgress()
		ret.State = ret.Run()
	}

	return ret.State
}


func ToolNewBash(cmd ...interface{}) *ToolExecCommand {
	ret := New(nil)

	for range onlyOnce {
		bash, err := exec.LookPath("bash")
		if err != nil {
			ret.State.SetError("Executable not found.")
			break
		}
		ret.cmd.SetPath(bash)

		ret.args = []string{"-c"}

		a := toolTypes.ReflectStrings(cmd...)
		ret.args = append(ret.args, *a...)
	}

	return ret.Reflect()
}


// Intent: TBD
//
// Template examples:
//  {{ . }}
//  fmt.Println("Hello")
func (e *ToolExecCommand) AppendCommands(cmd ...interface{}) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}
	e.State.SetFunction("")

	for range onlyOnce {
		a := toolTypes.ReflectStrings(cmd...)
		e.args = append(e.args, *a...)
	}

	return e.State
}
func (e *ToolExecCommand) Append(cmd ...interface{}) *ux.State {
	return e.AppendCommands(cmd...)
}


func (e *ToolExecCommand) Run() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}
	e.State.SetFunction("")

	for range onlyOnce {
		file, err := ioutil.TempFile("tmp", "scribe-shell")
		if err != nil {
			log.Fatal(err)
		}
		//noinspection ALL
		defer os.Remove(file.Name())

		bash, err := exec.LookPath("bash")
		if err != nil {
			e.State.SetError("Executable not found.")
			break
		}
		e.cmd.SetPath(bash)

		e.Reflect().args = []string{"-c"}
		e.Reflect().ShowProgress()
		e.State = e.Reflect().Run()
	}

	return e.State
}
