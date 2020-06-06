package toolGit

import (
	"github.com/newclarity/scribeHelpers/ux"
	"gopkg.in/src-d/go-git.v4"
)


// Usage:
//		{{- $cmd := $git.Clone }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
// func (me *ToolGit) Clone(url string, dir ...interface{}) *TypeExecCommand {
func (g *TypeGit) Clone() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		if g.IsNotAvailable() {
			break
		}

		if g.Url == "" {
			g.State.SetError("Git repo URL is empty")
			break
		}

		g.Base.StatPath()
		if g.Base.Exists() {
			g.State.SetError("cannot clone as path %s already exists", g.Base.GetPath())
			g.State.SetExitCode(1) // Fake an exit code.
			break
		}


		ux.PrintfWhite("Cloning %s into %s\n", g.Url, g.Base.GetPath())
		g.skipDirCheck = true
		g.State = g.Exec(gitCommandClone, g.Url, g.Base.GetPath())
		g.skipDirCheck = false
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.Open }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) Open() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		if g.IsNotAvailable() {
			break
		}

		g.State = g.Exec("rev-parse", "--is-inside-work-tree")
		if !g.State.OutputEquals("true") {
			if g.State.IsError() {
				g.State.SetError("current directory does not contain a valid .Git repository: %s", g.State.GetError())
				break
			}

			g.State.SetError("current directory does not contain a valid Git repository")
			break
		}

		var err error
		g.repository, err = git.PlainOpen(g.Base.GetPath())
		if err != nil {
			g.State.SetError(err)
			break
		}

		c, _ := g.repository.Config()
		g.Url = c.Remotes["origin"].URLs[0]

		g.State.SetOk("Opened directory %s.\nRemote origin is set to %s\n", g.Base.GetPath(), g.Url)
		ok := true
		g.State.SetResponse(&ok)
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.SetPath }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) SetPath(path ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		if g.IsNotAvailable() {
			break
		}

		//p := toolPath.ReflectAbsPath(path...)
		//if p == nil {
		//	g.State.SetError("path repo is nil")
		//	break
		//}
		//if *p == "" {
		//	g.State.SetError("path repo is nil")
		//	break
		//}

		g.Base.SetPath(path...)
		//if ! g.Base.SetPath(path...) {
		//	g.State.SetError("path repo '%s' cannot be set", g.Base.GetPath())
		//	break
		//}

		g.State = g.Base.StatPath()
		//if state.IsError() {
		//	g.State = state
		//	break
		//}

		if g.Base.NotExists() {
			g.State.Clear()
			break
		}
		if g.Base.IsAFile() {
			g.State.SetError("path repo '%s' exists and is a file", g.Base.GetPath())
			break
		}
		g.State = g.Chdir()
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.GetUrl }}
//		{{- if $cmd.IsOk }}{{ $cmd.data }}{{- end }}
func (g *TypeGit) GetUrl() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		g.State = g.Exec("config", "--get", "remote.origin.url")
		if g.State.IsError() {
			break
		}

		g.Url = g.State.Output
		g.State.SetResponse(&g.State.Output)
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.SetUrl }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) SetUrl(u Url) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.Url = u
	return g.State
}
