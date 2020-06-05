package toolGit

import (
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/ux"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)


// Usage:
//		{{- $cmd := $git.Clone }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
// func (me *ToolGit) Clone(url string, dir ...interface{}) *TypeExecCommand {
func (g *ToolGit) Clone() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range onlyOnce {
		if g.Reflect().IsNotAvailable() {
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
//func (g *ToolGit) Clone() *ux.State {
//	for range onlyOnce {
//		if g.Reflect().IsNotOk() {
//			break
//		}
//
//		if url == "" {
//			g.State.SetError("URL is nil")
//			break
//		}
//		g.SetUrl(url)
//
//
//		d := toolPath.ReflectAbsPath(dir...)
//		if d == nil {
//			g.State.SetError("dir is nil")
//			break
//		}
//
//		if !g.Base.SetPath(*d) {
//			g.State.SetError("error setting path to %s", g.Base.GetPath())
//			break
//		}
//
//		g.Base.StatPath()
//		if g.Base.Exists() {
//			g.State.SetError("cannot clone as path %s already exists", g.Base.GetPath())
//			g.Cmd.Exit = 1
//			break
//		}
//
//
//		ux.PrintfWhite("Cloning %s into %s\n", g.Url, g.Base.GetPath())
//		g.skipDirCheck = true
//		g.State = g.Exec(gitCommandClone, g.Url, g.Base.GetPath())
//		g.skipDirCheck = false
//	}
//
//	if g.State.IsError() {
//		g.State.SetError("Clone() - %s", g.State.Error)
//	}
//	return g.State
//}


// Usage:
//		{{- $cmd := $git.Open }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) IsExisting() bool {
	var ok bool
	if state := g.IsNil(); state.IsError() {
		return false
	}
	g.State.SetFunction("")

	for range onlyOnce {
		if g.Reflect().IsNotAvailable() {
			break
		}

		ok = true
	}

	return ok
}
func (g *ToolGit) IsNotExisting() bool {
	return !g.IsExisting()
}

// Usage:
//		{{- $cmd := $git.Open }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) Open() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range onlyOnce {
		if g.Reflect().IsNotAvailable() {
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
		g.State.Response = true
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.SetPath }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) SetPath(path ...interface{}) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range onlyOnce {
		if g.Reflect().IsNotAvailable() {
			break
		}

		p := toolPath.ReflectAbsPath(path...)
		if p == nil {
			g.State.SetError("path repo is nil")
			break
		}
		if *p == "" {
			g.State.SetError("path repo is nil")
			break
		}


		if !g.Base.SetPath(*p) {
			g.State.SetError("path repo '%s' cannot be set", *p)
			break
		}

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
			g.State.SetError("path repo '%s' exists and is a file", *p)
			break
		}
		g.State = g.Chdir()
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.GetUrl }}
//		{{- if $cmd.IsOk }}{{ $cmd.Data }}{{- end }}
func (g *ToolGit) GetUrl() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range onlyOnce {
		g.State = g.Exec("config", "--get", "remote.origin.url")
		if g.State.IsError() {
			break
		}

		g.Url = g.State.Output
		g.State.Response = g.State.Output
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.SetUrl }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) SetUrl(u Url) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.Url = u
	return g.State
}


// Usage:
//		{{- $cmd := $git.Clone }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
// func (me *ToolGit) Clone(url interface{}, dir ...interface{}) *TypeExecCommand {
func (g *ToolGit) Remove() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range onlyOnce {
		g.State = g.Base.StatPath()
		if g.State.IsError() {
			break
		}

		// @TODO - TO BE IMPLEMENTED

		//g.Base.Exists
		//ps := toolSystem.ResolveAbsPath(*d)
		//if ps.IsFile {
		//	break
		//}
		//if ps.IsDir {
		//	if ps.Exists {
		//		g.State.SetError("Repository exists for directory '%s'", ps.Filename)
		//		g.Cmd.Exit = 1
		//		break
		//	}
		//}
		//
		//g.SetUrl(*u)
		//g.Base = ps
		//ux.PrintfWhite("Cloning %s into %s\n", g.Url, g.Base.Filename)
		//
		//g.skipDirCheck = true
		//g.Cmd = (*toolTypes.TypeExecCommand)(g.Exec(gitCommandClone, g.Url, g.Base.Filename))
		//g.skipDirCheck = false
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.Lock }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) Lock() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range onlyOnce {
		g.State = g.GetTagObject(LockTag)
		if g.State.IsError() {
			break
		}

		var to *object.Tag
		to = g.State.Response.(*object.Tag)

		g.State.Response = to.ID()
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.GetStatus }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GetStatus() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range onlyOnce {
		var wt *git.Worktree
		var err error
		wt, err = g.repository.Worktree()
		g.State.SetError(err)
		if g.State.IsError() {
			break
		}

		var sts git.Status
		sts, err = wt.Status()
		g.State.SetError(err)
		if g.State.IsError() {
			break
		}

		g.State.Response = sts
	}

	return g.State
}
