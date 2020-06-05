package toolGit

import (
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/tsuyoshiwada/go-gitcmd"
)

type ToolGit TypeGit
func (g *ToolGit) Reflect() *TypeGit {
	return (*TypeGit)(g)
}
func (g *TypeGit) Reflect() *ToolGit {
	return (*ToolGit)(g)
}

func (g *ToolGit) IsNil() *ux.State {
	if state := ux.IfNilReturnError(g); state.IsError() {
		return state
	}
	g.State = g.State.EnsureNotNil()
	return g.State
}


// Usage:
//		{{ $git := NewGit }}
func ToolNewGit(path ...interface{}) *ToolGit {
	ret := New(nil)

	for range onlyOnce {
		p := toolPath.ReflectAbsPath(path...)
		if p == nil {
			break
		}
		if ret.Base.SetPath(*p) {
			state := ret.Base.StatPath()
			ret.State = state
			if ret.Base.Exists() {

			}
			if ret.State.IsError() {
				break
			}

			// Can now set it after.
			//ret.State.SetError("%s destination empty", *p)
			//break
		}

		//ret.Cmd = toolExec.NewExecCommand(false)
		ret.client = gitcmd.New(ret.GitConfig)

		if ret.IsNotAvailable() {
			break
		}
	}

	return ReflectToolGit(ret)
}


// Usage:
//		{{ $cmd := $git.Chdir .Some.Directory }}
//		{{ if $git.IsOk }}Changed to directory {{ $git.Dir }}{{ end }}
func (g *TypeGit) Chdir() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	return toolPath.ToolChdir(g.Base.GetPath()).State
}


// Usage:
//		{{ $git.SetDryRun }}
func (g *TypeGit) SetDryRun() bool {
	g.GitOptions = append(g.GitOptions, "-n")
	return true
}


// Usage:
//		{{ $cmd := $git.Exec "tag" "-l" }}
//		{{ if $git.IsOk }}OK{{ end }}
// func (me *ToolGit) Exec(cmd interface{}, args ...interface{}) *ux.State {
func (g *TypeGit) Exec(cmd string, args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range onlyOnce {
		if g.IsNotAvailable() {
			break
		}

		//c := toolTypes.ReflectString(cmd)
		//if c == nil {
		//	break
		//}
		//g.Cmd.exe = *c
		//
		//a := toolTypes.ReflectStrings(args...)
		//if a == nil {
		//	break
		//}
		//
		//g.Cmd.args = []string{}
		//g.Cmd.args = append(g.Cmd.args, g.GitOptions...)
		//g.Cmd.args = append(g.Cmd.args, *a...)

		//g.Cmd.SetPath(cmd)
		//g.Cmd.AddArgs(g.GitOptions...)
		//g.Cmd.SetArgs(args...)
		a := g.GitOptions
		a = append(a, args...)

		for range onlyOnce {
			if g.skipDirCheck {
				break
			}
			if g.Base.IsCwd() {
				break
			}
			path := g.Base.GetPath()
			cd := toolPath.ToolChdir(path)
			if cd.State.IsError() {
				ux.PrintfError("Cannot change directory to '%s'", path)
				break
			}
		}

		//out, err := g.client.Exec(g.Cmd.GetExe(), g.Cmd.GetArgs()...)
		out, err := g.client.Exec(cmd, a...)
		g.State.SetOutput(out)
		g.State.OutputTrim()
		g.State.SetError(err)

		if g.State.IsError() {
			g.State.SetExitCode(1) // Fake an exit code.
			break
		}

		g.State.SetOk("")
	}

	return g.State
}


//// Usage:
////		{{- $cmd := $git.IsExec }}
////		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
//func (g *TypeGit) IsAvailable() *ux.State {
//	for range onlyOnce {
//		if g.Reflect().IsNotAvailable() {
//			break
//		}
//	}
//
//	//foo := &State{}
//	//foo = (*State)(g.State)
//	//foo = (*State)(g.Reflect().State)
//	//foo = ReflectState(g.State)
//
//	return g.State
//}
//
//
//// Usage:
////		{{ if $ret.IsError }}{{ $cmd.PrintError }}{{ end }}
//func (g *TypeGit) SetError(error ...interface{}) {
//	g.State.SetError(error...)
//}
//
//
//// Usage:
////		{{ if $ret.IsError }}{{ $cmd.PrintError }}{{ end }}
//func (g *TypeGit) IsError() bool {
//	return g.State.IsError()
//}
//
//
//// Usage:
////		{{ if $ret.IsOk }}OK{{ end }}
//func (g *TypeGit) IsOk() bool {
//	return g.State.IsOk()
//}
//
//
//// Usage:
////		{{ if $ret.IsOk }}OK{{ end }}
//func (g *TypeGit) PrintError() string {
//	return g.Cmd.PrintError()
//}
//
//
//// Usage:
////		{{ if $ret.IsOk }}OK{{ end }}
//func (g *TypeGit) ExitOnError() string {
//	return g.State.ExitOnError()
//}
