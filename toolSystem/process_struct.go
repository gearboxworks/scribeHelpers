package toolSystem

import (
	"github.com/gearboxworks/scribeHelpers/toolPath"
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
	"github.com/shirou/gopsutil/process"
)

type ProcessGetter interface {
}


type TypeProcesses struct {
	procs   []*TypeProcess

	runtime *toolRuntime.TypeRuntime
	State   *ux.State
}

type TypeProcess struct {
	name      string
	ppid      int32
	pid       int32

	proc      *process.Process
	pathExe   *toolPath.TypeOsPath
	pathCwd   *toolPath.TypeOsPath
	openFiles *TypeOpenFiles

	runtime   *toolRuntime.TypeRuntime
	State     *ux.State
}

type TypeOpenFiles struct {
	Files   []*toolPath.TypeOsPath

	runtime *toolRuntime.TypeRuntime
	State   *ux.State
}


func ReflectToolProcess(p *TypeProcess) *ToolProcess {
	return (*ToolProcess)(p)
}

func (p *TypeProcess) IsNil() *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}
	p.State = p.State.EnsureNotNil()
	return p.State
}


func NewProcesses(runtime *toolRuntime.TypeRuntime) *TypeProcesses {
	runtime = runtime.EnsureNotNil()

	ret := &TypeProcesses {
		procs:   nil,

		runtime: runtime,
		State:   ux.NewState(runtime.CmdName, runtime.Debug),
	}

	return ret
}


func NewProcess(runtime *toolRuntime.TypeRuntime) *TypeProcess {
	runtime = runtime.EnsureNotNil()

	p := &TypeProcess{
		name:      "",
		ppid:      0,
		pid:       0,

		proc:      nil,
		pathExe:   nil,	// toolPath.ToolNewPath(),
		pathCwd:   nil,	// toolPath.ToolNewPath(),

		openFiles: NewOpenFiles(runtime),

		runtime:   runtime,
		State:     ux.NewState(runtime.CmdName, runtime.Debug),
	}
	p.State.SetPackage("")
	p.State.SetFunctionCaller()

	return p
}


func NewOpenFiles(runtime *toolRuntime.TypeRuntime) *TypeOpenFiles {
	runtime = runtime.EnsureNotNil()

	ret := &TypeOpenFiles {
		Files: nil,

		runtime: runtime,
		State: ux.NewState(runtime.CmdName, runtime.Debug),
	}

	return ret
}
