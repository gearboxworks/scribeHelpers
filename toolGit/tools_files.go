package toolGit

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
)


// Usage:
//		{{- $cmd := $git.ChangedFiles }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) ChangedFiles() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range onlyOnce {
		g.State = g.Exec(gitCommandStatus, "--porcelain")
		if g.State.IsError() {
			break
		}

		var fps Filepaths
		fps = make(Filepaths, len(g.State.OutputArray))
		for i, fp := range g.State.OutputArray {
			s := strings.Fields(fp)
			if len(s) > 1 {
				fps[i] = s[1]
			}
		}

		g.State.Response = fps
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.AddFiles }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) AddFiles(files ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range onlyOnce {
		if len(files) == 0 {
			files = []string{"."}
		}

		files = append([]string{"--porcelain"}, files...)
		g.State = g.Exec(gitCommandAdd, files...)
		if g.State.IsNotOk() {
			break
		}

		var fps Filepaths
		fps = make(Filepaths, len(g.State.OutputArray))
		for i, fp := range g.State.OutputArray {
			s := strings.Fields(fp)
			if len(s) > 1 {
				fps[i] = s[1]
			}
		}

		g.State.Response = fps
	}

	return g.State
}


// Usage:
func (g *ToolGit) Push(comment string, args ...interface{}) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range onlyOnce {
		c := fmt.Sprintf(comment, args...)
		if c == "" {
			g.State.SetError("Missing comment to Git commit.")
			break
		}
		var count int

		ux.PrintflnBlue("Adding files to repo...")
		g.State = g.AddFiles(".")
		if g.State.IsNotOk() {
			break
		}
		count = len((g.State.Response).([]string))
		ux.PrintflnBlue("Added %d files to repo.", )


		g.State = g.ChangedFiles()
		if g.State.IsNotOk() {
			break
		}

		g.State = g.Exec(gitCommandAdd, "--porcelain")
		if g.State.IsNotOk() {
			break
		}

		var fps Filepaths
		fps = make(Filepaths, len(g.State.OutputArray))
		for i, fp := range g.State.OutputArray {
			s := strings.Fields(fp)
			if len(s) > 1 {
				fps[i] = s[1]
			}
		}

		g.State.Response = fps
	}

	return g.State
}
