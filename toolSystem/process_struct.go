package toolSystem

import (
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/shirou/gopsutil/process"
)

type ProcessGetter interface {
}


type TypeProcesses struct {
	procs []*TypeProcess

	State *ux.State
	Debug bool
}

type TypeProcess struct {
	name      string
	ppid      int32
	pid       int32

	proc      *process.Process
	pathExe   *toolPath.TypeOsPath
	pathCwd   *toolPath.TypeOsPath
	openFiles *TypeOpenFiles

	State     *ux.State
	Debug     bool
}

type TypeOpenFiles struct {
	Files []*toolPath.TypeOsPath
	State *ux.State
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


func NewProcesses(debugMode bool) *TypeProcesses {
	ret := &TypeProcesses {
		procs: nil,

		Debug: debugMode,
		State: ux.NewState(debugMode),
	}

	return ret
}


func NewProcess(debugMode bool) *TypeProcess {
	p := &TypeProcess{
		name:      "",
		ppid:      0,
		pid:       0,

		proc:      nil,
		pathExe:   nil,	// toolPath.ToolNewPath(),
		pathCwd:   nil,	// toolPath.ToolNewPath(),

		openFiles: NewOpenFiles(debugMode),

		Debug:     debugMode,
		State:     ux.NewState(debugMode),
	}
	p.State.SetPackage("")
	p.State.SetFunctionCaller()

	return p
}


func NewOpenFiles(debugMode bool) *TypeOpenFiles {
	ret := &TypeOpenFiles {
		Files: nil,
		State: ux.NewState(debugMode),
	}

	return ret
}
