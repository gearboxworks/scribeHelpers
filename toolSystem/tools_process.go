package toolSystem

import (
	"github.com/newclarity/scribeHelpers/toolTypes"
	"github.com/newclarity/scribeHelpers/ux"
)


type ToolProcesses TypeProcesses
func (p *ToolProcesses) Reflect() *TypeProcesses {
	return (*TypeProcesses)(p)
}
func (p *TypeProcesses) Reflect() *ToolProcesses {
	return (*ToolProcesses)(p)
}
func (p *ToolProcesses) IsNil() *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}
	p.State = p.State.EnsureNotNil()
	return p.State
}


type ToolProcess TypeProcess
func (p *ToolProcess) Reflect() *TypeProcess {
	return (*TypeProcess)(p)
}
func (p *TypeProcess) Reflect() *ToolProcess {
	return (*ToolProcess)(p)
}
func (p *ToolProcess) IsNil() *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}
	p.State = p.State.EnsureNotNil()
	return p.State
}


// Usage:
//		{{ $procs := FindByName }}
func (p *ToolProcesses) FindByName(name interface{}) *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}

	for range onlyOnce {
		n := toolTypes.ReflectString(name)
		if n == nil {
			p.State.SetError("process name undefined")
			break
		}
		if *n == "" {
			p.State.SetError("process name undefined")
			break
		}

		p.State = p.Reflect().FindByName(*n)
	}

	return p.State
}


// Usage:
//		{{ $procs := FindByPid }}
func (p *ToolProcesses) FindByPid(pid interface{}) *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}

	for range onlyOnce {
		n := toolTypes.ReflectInt32(pid)
		if n == nil {
			p.State.SetError("PID undefined")
			break
		}
		if *n == 0 {
			p.State.SetError("PID undefined")
			break
		}

		p.State = p.Reflect().FindByPid(*n)
	}

	return p.State
}


// Usage:
//		{{ $procs := FindByName }}
func (p *ToolProcesses) Print() string {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return ""
	}
	var ret string

	for range onlyOnce {
		for _, proc := range p.procs {
			ret += ux.SprintfWhite("%d %d %s\n",
				proc.GetPid(),
				proc.GetPpid(),
				proc.GetExePath(),
				)
		}
	}

	return ret
}


// Usage:
func ToolFindProcByName(name interface{}) *ToolProcesses {
	p := NewProcesses(nil)
	p.State = p.Reflect().FindByName(name)
	return p.Reflect()
}


// Usage:
func ToolFindProcByPid(pid interface{}) *ToolProcesses {
	p := NewProcesses(nil)
	p.State = p.Reflect().FindByPid(pid)
	return p.Reflect()
}
