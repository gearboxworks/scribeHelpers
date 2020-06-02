package helperSystem

import (
	"github.com/newclarity/scribeHelpers/helperTypes"
	"github.com/newclarity/scribe/ux"
)


type HelperProcesses TypeProcesses
func (p *HelperProcesses) Reflect() *TypeProcesses {
	return (*TypeProcesses)(p)
}
func (p *TypeProcesses) Reflect() *HelperProcesses {
	return (*HelperProcesses)(p)
}
func (p *HelperProcesses) IsNil() *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}
	p.State = p.State.EnsureNotNil()
	return p.State
}


type HelperProcess TypeProcess
func (p *HelperProcess) Reflect() *TypeProcess {
	return (*TypeProcess)(p)
}
func (p *TypeProcess) Reflect() *HelperProcess {
	return (*HelperProcess)(p)
}
func (p *HelperProcess) IsNil() *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}
	p.State = p.State.EnsureNotNil()
	return p.State
}


// Usage:
//		{{ $procs := FindByName }}
func (p *HelperProcesses) FindByName(name interface{}) *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}

	for range OnlyOnce {
		n := helperTypes.ReflectString(name)
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
func (p *HelperProcesses) FindByPid(pid interface{}) *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}

	for range OnlyOnce {
		n := helperTypes.ReflectInt32(pid)
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
func (p *HelperProcesses) Print() string {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return ""
	}
	var ret string

	for range OnlyOnce {
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
func HelperFindProcByName(name interface{}) *HelperProcesses {
	p := NewProcesses(false)
	p.State = p.Reflect().FindByName(name)
	return p.Reflect()
}


// Usage:
func HelperFindProcByPid(pid interface{}) *HelperProcesses {
	p := NewProcesses(false)
	p.State = p.Reflect().FindByPid(pid)
	return p.Reflect()
}
