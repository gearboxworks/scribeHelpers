package toolExec

import (
	"bytes"
	"github.com/gearboxworks/scribeHelpers/toolPath"
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/toolTypes"
	"github.com/gearboxworks/scribeHelpers/ux"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)


type TypeExecCommandGetter interface {
}

type TypeExecCommand struct {
	cmd        *toolPath.TypeOsPath
	args       []string

	workingDir *toolPath.TypeOsPath

	show       bool
	stdout     []byte
	stderr     []byte
	exit       int

	Runtime    *toolRuntime.TypeRuntime
	State      *ux.State
}
func (e *TypeExecCommand) IsNil() *ux.State {
	return ux.IfNilReturnError(e)
}


func New(runtime *toolRuntime.TypeRuntime) *TypeExecCommand {
	runtime = runtime.EnsureNotNil()

	ret := &TypeExecCommand {
		cmd:        toolPath.New(runtime),
		args:       nil,

		workingDir: nil,	// nil - means not defined.

		show:       false,
		stdout:     []byte{},
		stderr:     []byte{},
		exit:       0,

		Runtime:    runtime,
		State:      ux.NewState(runtime.CmdName, runtime.Debug),
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


func (e *TypeExecCommand) IsRunnable() bool {
	var ok bool
	if state := e.IsNil(); state.IsError() {
		return ok
	}

	for range onlyOnce {
		_, err := exec.LookPath(e.cmd.GetPath())
		if err != nil {
			e.State.SetError("Executable not found.")
			break
		}

		e.State = e.cmd.StatPath()
		if e.State.IsNotOk() {
			break
		}

		ok = true
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
		if !e.IsRunnable() {
			e.State.PrintResponse()
			break
		}

		for range onlyOnce {
			if e.workingDir == nil {
				e.State.SetOk()
				break
			}
			e.State = e.workingDir.StatPath()	// Re-stat as things may have changed.
			if e.workingDir.NotExists() {
				e.State.SetError("Working directory doesn't exist.")
				break
			}
			if e.State.IsError() {
				break
			}
			e.workingDir.Chdir()
			if e.State.IsError() {
				break
			}
		}
		if e.State.IsError() {
			break
		}

		//c := exec.Command((*cmds)[0], (*cmds)[1:]...)
		c := exec.Command(e.cmd.GetPath(), e.args...)

		var err error
		if e.show {
			var stdoutBuf, stderrBuf bytes.Buffer
			c.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
			c.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

			err = c.Run()
			//err = c.Run()
			//if err != nil {
			//	e.State.SetError(err)
			//}

			e.State.SetError(err)
			e.stdout = stdoutBuf.Bytes()
			e.stderr = stderrBuf.Bytes()
			err = nil	// @TODO - The 'err' value within here needs to be not visible outside.

		} else {
			e.stdout, err = c.CombinedOutput()
			e.State.SetError(err)
		}

		e.State.SetOutput(e.stdout)
		e.State.SetResponse(&e.stderr)

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
func (e *TypeExecCommand) SetCmd(path ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}

	for range onlyOnce {
		if e.cmd == nil {
			e.cmd = toolPath.New(e.Runtime)
		}

		e.cmd.SetPath(path...)
		if e.cmd.IsRelative() {
			p, err := exec.LookPath(e.cmd.GetPath())
			if err != nil {
				e.State.SetError("Executable not found.")
				break
			}
			e.cmd.SetPath(p)
		}

		if e.IsRunnable() {
			// Will set e.State, if error.
			break
		}
	}

	return e.State
}


func (e *TypeExecCommand) SetWorkingPath(path ...string) *ux.State {
	if state := e.IsNil(); state.IsError() {
		return nil
	}

	for range onlyOnce {
		if e.workingDir == nil {
			e.workingDir = toolPath.New(e.Runtime)
		}

		//@TODO - Sometimes a directory may not be present until actually running the command.
		//@TODO - So maybe don't actually stat the dir here, but do it within TypeExecCommand.Run()
		e.workingDir.SetPath(path...)
		e.State = e.workingDir.StatPath()
		if e.State.IsError() {
			e.workingDir = nil	// nil - means not set.
			break
		}
		if !e.workingDir.IsADir() {
			e.State.SetError("Working path is not a directory.")
			e.workingDir = nil	// nil - means not set.
			break
		}
	}

	return e.State
}
func (e *TypeExecCommand) GetWorkingPath() string {
	var ret string
	if state := e.IsNil(); state.IsError() {
		return ""
	}

	for range onlyOnce {
		if e.workingDir == nil {
			break
		}

		ret = e.workingDir.GetPath()
	}

	return ret
}
func (e *TypeExecCommand) GetWorkingPathAbs() string {
	var ret string
	if state := e.IsNil(); state.IsError() {
		return ""
	}

	for range onlyOnce {
		if e.workingDir == nil {
			break
		}

		ret = e.workingDir.GetPathAbs()
	}

	return ret
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
func (e *TypeExecCommand) GetStdoutArray(sep string) []string {
	return strings.Split(string(e.stdout), sep)
}


func (e *TypeExecCommand) GetStderr() []byte {
	return e.stderr
}
func (e *TypeExecCommand) GetStderrString() string {
	return string(e.stderr)
}
func (e *TypeExecCommand) GetStderrArray(sep string) []string {
	return strings.Split(string(e.stderr), sep)
}


func (e *TypeExecCommand) GetExitCode() int {
	return e.exit
}
