package toolSystem

import (
	"github.com/gearboxworks/scribeHelpers/ux"
	"github.com/shirou/gopsutil/process"
	"strings"
)


func (p *TypeProcesses) FindByName(name string) *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}

	for range onlyOnce {
		procs, err := process.Processes()
		if err != nil {
			p.State.SetError(err)
			break
		}

		for _, proc := range procs {
			if !matchProcName(name, proc) {
				continue
			}

			newProc := NewProcess(p.runtime)
			newProc.populateStruct(proc)
			p.procs = append(p.procs, newProc)
		}
	}

	return p.State
}


func (p *TypeProcesses) FindByPid(pid int32) *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}

	for range onlyOnce {
		procs, err := process.Processes()
		if err != nil {
			p.State.SetError(err)
			break
		}

		for _, proc := range procs {
			if proc.Pid != pid {
				continue
			}

			newProc := NewProcess(p.runtime)
			newProc.populateStruct(proc)
			p.procs = append(p.procs, newProc)
		}
	}

	return p.State
}


func matchProcName(name string, proc *process.Process) bool {
	var ok bool

	for range onlyOnce {
		n, err := proc.Name()
		if err != nil {
			break
		}
		if n == name {
			ok = true
			break
		}
		if strings.HasSuffix(n, name) {
			ok = true
			break
		}

		n, err = proc.Cmdline()
		if err != nil {
			break
		}
		if n == name {
			ok = true
			break
		}
		if strings.HasSuffix(n, name) {
			ok = true
			break
		}
	}

	return ok
}
