package toolGit

import (
	"github.com/newclarity/scribeHelpers/ux"
	"gopkg.in/src-d/go-git.v4"
	"net/url"
	"strings"
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

		if g.Url.String() == "" {
			g.State.SetError("Git repo URL is empty")
			break
		}

		g.Base.StatPath()
		if g.Base.Exists() {
			g.State.SetError("cannot clone as path %s already exists", g.Base.GetPath())
			g.State.SetExitCode(1) // Fake an exit code.
			break
		}


		if g.State.IsVerboseMode() {
			ux.PrintfWhite("Cloning %s into %s\n", g.Url, g.Base.GetPath())
		}
		g.skipDirCheck = true
		g.State = g.Exec(gitCommandClone, g.Url.String(), g.Base.GetPath())
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
		g.State = g.SetUrl(c.Remotes["origin"].URLs[0])
		if g.State.IsNotOk() {
			break
		}

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

		if path == nil {
			g.State.SetError("path is empty")
			break
		}

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
			g.State.SetOk() // We may want to clone after we set the path.
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
//		{{- $cmd := $git.GetPath }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GetPath() string {
	if state := g.IsNil(); state.IsError() {
		return ""
	}
	return g.Base.GetPath()
}


// Usage:
//		{{- $cmd := $git.FetchUrl }}
//		{{- if $cmd.IsOk }}{{ $cmd.data }}{{- end }}
//func (g *TypeGit) FetchUrl() (string, *ux.State) {
func (g *TypeGit) FetchUrl() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		g.State = g.Exec("config", "--get", "remote.origin.url")
		if g.State.IsError() {
			break
		}

		g.SetUrl(g.State.Output)
		//g.Url = g.State.Output
		//g.State.SetResponse(&g.State.Output)
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.GetUrl }}
//		{{- if $cmd.IsOk }}{{ $cmd.data }}{{- end }}
//func (g *TypeGit) GetUrl() (string, *ux.State) {
func (g *TypeGit) GetUrl() string {
	if state := g.IsNil(); state.IsError() {
		return ""
	}
	return g.Url.String()
}


// Usage:
//		{{- $cmd := $git.SetUrl }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) SetUrl(u ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	var err error
	g.Url, err = url.Parse(addPrefix(u...))
	if err != nil {
		g.State.SetError(err)
	}
	return g.State
}


func addPrefix(urlString ...string) string {
	var ret string
	for range onlyOnce {
		ret = strings.Join(urlString, "/")

		if strings.HasPrefix(ret, "http") {
			// We have a full URL - no change.
			break
		}

		if strings.HasPrefix(ret, "github.com") {
			// We have a github.com specific string.
			ret = "https://" + ret
			break
		}

		ua := strings.Split(ret, "/")
		if len(ua) == 0 {
			// Dunno, leave as is.
			break
		}

		if strings.Contains(ua[0], ".") {
			// We have a host defined in the first segment.
			ret = "https://" + ret
			break
		}

		// We probably just have a "owner/repo_name" style URL.
		ret = "https://github.com/" + ret
	}

	return ret
}
