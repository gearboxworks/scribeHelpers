package toolExec

import (
	"github.com/newclarity/scribeHelpers/toolTypes"
	"github.com/newclarity/scribeHelpers/ux"
	"io/ioutil"
	"log"
	"os"
)


func HelperExecBash(cmd ...interface{}) *ux.State {
	ret := New(false)

	for range OnlyOnce {
		a := toolTypes.ReflectStrings(cmd...)

		ret.exe = "bash"
		ret.args = []string{"-c"}
		ret.args = append(ret.args, *a...)
		ret.ShowProgress()
		ret.State = ret.Run()
	}

	return ret.State
}


func HelperNewBash(cmd ...interface{}) *HelperExecCommand {
	ret := New(false)

	for range OnlyOnce {
		ret.exe = "bash"
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
func (e *HelperExecCommand) AppendCommands(cmd ...interface{}) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}
	e.State.SetFunction("")

	for range OnlyOnce {
		a := toolTypes.ReflectStrings(cmd...)
		e.args = append(e.args, *a...)
	}

	return e.State
}
func (e *HelperExecCommand) Append(cmd ...interface{}) *ux.State {
	return e.AppendCommands(cmd...)
}


func (e *HelperExecCommand) Run() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}
	e.State.SetFunction("")

	for range OnlyOnce {
		file, err := ioutil.TempFile("tmp", "scribe-shell")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(file.Name())

		e.Reflect().exe = "bash"
		e.Reflect().args = []string{"-c"}
		e.Reflect().ShowProgress()
		e.State = e.Reflect().Run()
	}

	return e.State
}
