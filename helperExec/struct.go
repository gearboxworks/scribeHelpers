package helperExec

import (
	"bytes"
	"github.com/newclarity/scribeHelpers/helperPath"
	"github.com/newclarity/scribeHelpers/helperTypes"
	"github.com/newclarity/scribeHelpers/ux"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)


type TypeExecCommandGetter interface {
}

type TypeExecCommand struct {
	exe    string
	args   []string

	show   bool
	stdout []byte
	stderr []byte
	exit   int

	debug  bool

	State  *ux.State
}


func NewExecCommand(debugMode bool) *TypeExecCommand {
	ret := &TypeExecCommand {
		exe:    "",
		args:   nil,

		show:   false,
		stdout: []byte{},
		stderr: []byte{},
		exit:   0,

		debug:  debugMode,

		State:   ux.NewState(debugMode),
	}
	ret.State.SetPackage("")
	ret.State.SetFunctionCaller()

	return ret
}


func ReflectExecCommand(ref ...interface{}) *TypeExecCommand {
	ec := NewExecCommand(false)

	for range OnlyOnce {
		s := *helperTypes.ReflectStrings(ref)
		if len(s) == 0 {
			break
		}
		if len(s) >= 1 {
			ec.SetPath(s[0])
		}
		if len(s) >= 2 {
			ec.SetArgs(s[1:]...)
		}
	}

	return ec
}


func (e *TypeExecCommand) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}


func (e *TypeExecCommand) IsRunnable() bool {
	if state := e.IsNil(); state.IsError() {
		return false
	}
	var ok bool

	for range OnlyOnce {
		if e.State == nil {
			e.State = ux.NewState(false)
		}

		_, err := exec.LookPath(e.exe)
		if err != nil {
			e.State.SetError("Executable not found.")
			break
		}
	}

	return ok
}


func (e *TypeExecCommand) Exec(cmd string, args ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		if e.State == nil {
			e.State = ux.NewState(false)
		}

		e.State = e.SetPath(cmd)
		if e.State.IsNotOk() {
			e.State.PrintResponse()
			break
		}

		if e.IsRunnable() {
			e.State.PrintResponse()
			break
		}

		e.State = e.SetArgs(args...)
		if e.State.IsNotOk() {
			e.State.PrintResponse()
			break
		}

		if e.debug {
			ux.PrintflnBlue("# Executing: %s %s", e.exe, strings.Join(e.args, " "))
		}
		e.State = e.Run()
		if e.State.IsNotOk() {
			e.State.PrintResponse()
			break
		}
	}

	return e.State
}


func (e *TypeExecCommand) Run() *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}

	for range OnlyOnce {
		if e.State == nil {
			e.State = ux.NewState(false)
		}

		//c := exec.Command((*cmds)[0], (*cmds)[1:]...)
		c := exec.Command(e.exe, e.args...)

		var err error
		if e.show {
			var stdoutBuf, stderrBuf bytes.Buffer
			c.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
			c.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

			err := c.Run()
			if err != nil {
				e.State.SetError(err)
			}

			e.State.SetError(err)
			e.stdout = stdoutBuf.Bytes()
			e.stderr = stderrBuf.Bytes()

		} else {
			e.stdout, err = c.CombinedOutput()
			e.State.SetError(err)
		}

		if e.State.IsError() {
			if exitError, ok := err.(*exec.ExitError); ok {
				waitStatus := exitError.Sys().(syscall.WaitStatus)
				e.exit = waitStatus.ExitStatus()
				e.State.SetExitCode(e.exit)
			}
			break
		}

		waitStatus := c.ProcessState.Sys().(syscall.WaitStatus)
		e.exit = waitStatus.ExitStatus()
		e.State.SetExitCode(e.exit)
	}

	return e.State
}


func (e *TypeExecCommand) ShowProgress() {
	if state := e.IsNil(); state.IsError() {
		return
	}
	e.show = true
}
func (e *TypeExecCommand) SilenceProgress() {
	if state := e.IsNil(); state.IsError() {
		return
	}
	e.show = false
}


func (e *TypeExecCommand) GetExe() string {
	return e.exe
}
func (e *TypeExecCommand) GetPath() string {
	return e.exe
}
func (e *TypeExecCommand) SetPath(path ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}

	ep := helperPath.HelperNewPath()
	ep.SetPath(path...)
	e.exe = ep.GetPath()
	return e.State
}


func (e *TypeExecCommand) GetArgs() []string {
	return e.args
}
func (e *TypeExecCommand) SetArgs(args ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}
	e.args = []string{}
	e.State = e.AddArgs(args...)
	return e.State
}
func (e *TypeExecCommand) AddArgs(args ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}
	e.args = append(e.args, args...)
	return e.State
}
func (e *TypeExecCommand) AppendArgs(args ...string) *ux.State {
	e.State = e.AddArgs(args...)
	return e.State
}


func (e *TypeExecCommand) GetStdout() []byte {
	return e.stdout
}
func (e *TypeExecCommand) GetStdoutString() string {
	return string(e.stdout)
}


func (e *TypeExecCommand) GetStderr() []byte {
	return e.stderr
}
func (e *TypeExecCommand) GetStderrString() string {
	return string(e.stderr)
}


func (e *TypeExecCommand) GetExitCode() int {
	return e.exit
}
