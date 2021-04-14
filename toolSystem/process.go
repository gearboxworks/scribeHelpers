package toolSystem

import (
	"github.com/gearboxworks/scribeHelpers/toolPath"
	"github.com/gearboxworks/scribeHelpers/ux"
	"github.com/shirou/gopsutil/process"
)


func (p *TypeProcess) GetName() string {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return ""
	}

	return p.name
}


func (p *TypeProcess) GetPid() int32 {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return 0
	}

	return p.pid
}


func (p *TypeProcess) GetPpid() int32 {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return 0
	}

	return p.ppid
}


func (p *TypeProcess) GetExePath() string {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return ""
	}

	return p.pathExe.GetPath()
}


func (p *TypeProcess) GetCwd() string {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return ""
	}

	return p.pathCwd.GetPath()
}


func (p *TypeProcess) GetOpenFiles() *TypeOpenFiles {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return &TypeOpenFiles{State: state}
	}

	for range onlyOnce {
		of, err := p.proc.OpenFiles()
		if err != nil {
			p.State.SetError(err)
			p.State.SetOk() // Effectively ignore error.
		}

		for _, f := range of {
			path := toolPath.New(p.runtime)
			path.SetPath(f.Path)
			p.openFiles.Files = append(p.openFiles.Files, path)
		}
	}

	p.openFiles.State = p.State
	return p.openFiles
}


//func (p *TypeProcess) Find() *ux.State {
//	if state := ux.IfNilReturnError(p); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		p.State.SetOk()
//
//		if p.name != "" {
//			p.State = p.FindByName(p.name)
//			break
//		}
//
//		if p.pid != 0 {
//			p.State = p.FindByPid(p.pid)
//			break
//		}
//
//		p.State.SetError("No PID or name specified.")
//	}
//
//	return p.State
//}
//
//
//func (p *TypeProcess) FindByPid(pid int32) *ux.State {
//	if state := ux.IfNilReturnError(p); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		procs, err := process.Processes()
//		if err != nil {
//			p.State.SetError(err)
//			break
//		}
//
//		var proc *process.Process
//		for _, proc = range procs {
//			if proc.Pid == pid {
//				p.populateStruct(proc)
//				break
//			}
//		}
//	}
//
//	return p.State
//}
//
//
//func (p *TypeProcess) FindByName(name string) *ux.State {
//	if state := ux.IfNilReturnError(p); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		procs, err := process.Processes()
//		if err != nil {
//			p.State.SetError(err)
//			break
//		}
//
//		var proc *process.Process
//		for _, proc = range procs {
//			var err error
//			var n string
//			//ux.PrintflnBlue("# %s\n", proc.String())
//
//			n, err = proc.Name()
//			if err != nil {
//				p.State.SetError(err)
//				break
//			}
//			if n == name {
//				break
//			}
//
//			n, err = proc.Cmdline()
//			if err != nil {
//				p.State.SetError(err)
//				break
//			}
//			if n == name {
//				break
//			}
//		}
//
//		if proc != nil {
//			p.populateStruct(proc)
//		}
//
//		//infoStat, _ := host.Info()
//		//ux.PrintflnBlue("Total processes: %d\n", infoStat.Procs)
//		//
//		//miscStat, _ := load.Misc()
//		//ux.PrintflnBlue("Running processes: %d\n", miscStat.ProcsRunning)
//	}
//
//	return p.State
//}


func (p *TypeProcess) populateStruct(proc *process.Process) {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return
	}

	for range onlyOnce {
		p.proc = proc
		p.pid = proc.Pid

		var err error
		p.ppid, err = proc.Ppid()
		if err != nil {
			p.State.SetError(err)
			break
		}

		p.name, err = proc.Name()
		if err != nil {
			p.State.SetError(err)
			break
		}

		var path string
		path, err = proc.Cwd()
		if err == nil {
			p.pathCwd = toolPath.ToolNewPath(path)
		}

		p.openFiles = p.GetOpenFiles()

		//var of []process.OpenFilesStat
		//of, err = p.proc.OpenFiles()
		//for _, h := range of {
		//	ux.PrintflnBlue("String: %s => %s, %d", h.String(), h.Filename, h.Fd)
		//	//if strings.HasSuffix(h.Filename, fmt.Sprintf("%c%s", filepath.Separator, p.name)) {
		//	//	p.pathExe = toolPath.ToolNewPath(h.Filename)
		//	//}
		//}

		path, err = proc.Exe()
		// Errors will be ignored.
		if path != "" {
			p.pathExe = toolPath.ToolNewPath(path)
		}


		//a, _ := p.proc.Name()
		//b, _ := p.proc.Exe()
		//c, _ := p.proc.Status()
		//d, _ := p.proc.Cwd()
		//e, _ := p.proc.Cmdline()
		//f, _ := p.proc.CmdlineSlice()
		//
		//ux.PrintflnBlue("Name: %s", a)
		//ux.PrintflnBlue("Exe: %s", b)
		//ux.PrintflnBlue("Status: %s", c)
		//ux.PrintflnBlue("Cwd: %s", d)
		//ux.PrintflnBlue("Cmdline: %s", e)
		//ux.PrintflnBlue("CmdlineSlice: %s", strings.Join(f, " | "))
		//ux.PrintflnBlue("")
	}

	return
}
