package toolGit

import (
	"github.com/gearboxworks/scribeHelpers/toolPath"
	"github.com/gearboxworks/scribeHelpers/ux"
	"strconv"
	"strings"
)


// Usage:
//		{{ $cmd := $git.Chdir .Some.Directory }}
//		{{ if $git.IsOk }}Changed to directory {{ $git.Dir }}{{ end }}
func (g *TypeGit) Chdir() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	return toolPath.ToolChdir(g.Base.GetPath()).State
}


// Usage:
//		{{ $git.SetDryRun }}
func (g *TypeGit) SetDryRun() bool {
	g.GitOptions = append(g.GitOptions, "-n")
	return true
}


// Usage:
//		{{- $cmd := $git.Open }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) IsExisting() bool {
	var ok bool
	if state := g.IsNil(); state.IsError() {
		return false
	}
	g.State.SetFunction()

	for range onlyOnce {
		if g.IsNotAvailable() {
			break
		}

		ok = true
	}

	return ok
}
func (g *TypeGit) IsNotExisting() bool {
	return !g.IsExisting()
}


// Usage:
//		{{ $cmd := $git.Exec "tag" "-l" }}
//		{{ if $git.IsOk }}OK{{ end }}
// func (me *ToolGit) Exec(cmd interface{}, args ...interface{}) *ux.State {
func (g *TypeGit) Exec(cmd string, args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

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

		if err == nil {
			g.State.SetOk()
			break
		}

		checkExit := err.Error()
		//fmt.Printf(":%s:\n", checkExit)
		if strings.HasPrefix(checkExit, "exit status ") {
			checkExit = strings.TrimPrefix(checkExit, "exit status ")
			exitCode, err := strconv.Atoi(checkExit)
			if err != nil {
				g.State.SetExitCode(1) // Fake an exit code.
				break
			}
			g.State.SetExitCode(exitCode)
		}
	}

	return g.State
}
