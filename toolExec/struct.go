package toolExec

import (
	"bytes"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/toolTypes"
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
	cmd     *toolPath.TypeOsPath
	args    []string

	show    bool
	stdout  []byte
	stderr  []byte
	exit    int

	Runtime *toolRuntime.TypeRuntime
	State   *ux.State
}


func New(runtime *toolRuntime.TypeRuntime) *TypeExecCommand {
	runtime = runtime.EnsureNotNil()

	ret := &TypeExecCommand {
		cmd:    toolPath.New(runtime),
		args:   nil,

		show:   false,
		stdout: []byte{},
		stderr: []byte{},
		exit:   0,

		Runtime:  runtime,
		State:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
	ret.State.SetPackage("")
	ret.State.SetFunctionCaller()
	return ret
}


func ReflectExecCommand(ref ...interface{}) *TypeExecCommand {
	ec := New(nil)

	for range onlyOnce {
		s := *toolTypes.ReflectStrings(ref)
		if len(s) == 0 {
			break
		}
		if len(s) >= 1 {
			ec.SetCmd(s[0])
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

	for range onlyOnce {
		_, err := exec.LookPath(e.cmd.GetPath())
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

	for range onlyOnce {
		e.State = e.SetCmd(cmd)
		if e.State.IsNotOk() {
			e.State.PrintResponse()
			break
		}

		e.State = e.SetArgs(args...)
		if e.State.IsNotOk() {
			e.State.PrintResponse()
			break
		}

		if e.Runtime.Debug {
			ux.PrintflnBlue("# Executing: %s %s", e.cmd, strings.Join(e.args, " "))
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

	for range onlyOnce {
		if e.IsRunnable() {
			e.State.PrintResponse()
			break
		}

		//c := exec.Command((*cmds)[0], (*cmds)[1:]...)
		c := exec.Command(e.cmd.GetPath(), e.args...)

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


func (e *TypeExecCommand) GetCmd() *toolPath.TypeOsPath {
	return e.cmd
}
func (e *TypeExecCommand) GetCmdPath() string {
	return e.cmd.GetPath()
}
//func (e *TypeExecCommand) GetExe() string {
//	return e.exe
//}
//func (e *TypeExecCommand) GetPath() string {
//	return e.exe
//}
func (e *TypeExecCommand) SetCmd(path ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}

	for range onlyOnce {
		if e.cmd == nil {
			e.cmd = toolPath.New(e.Runtime)
		}

		if e.cmd.SetPath(path...) {
			e.State.SetError("cannot set cmd path")
			break
		}

		if e.cmd.IsRelative() {
			p, err := exec.LookPath(e.cmd.GetPath())
			if err != nil {
				e.State.SetError("Executable not found.")
				break
			}
			if e.cmd.SetPath(p) {
				e.State.SetError("cannot set cmd path")
				break
			}
		}

		if e.IsRunnable() {
			// Will set e.State, if error.
			break
		}
	}

	return e.State
}
//func (e *TypeExecCommand) SetExe(path ...string) *ux.State {
//	return e.SetCmd(path...)
//}
//func (e *TypeExecCommand) SetPath(path ...string) *ux.State {
//	return e.SetCmd(path...)
//}


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
