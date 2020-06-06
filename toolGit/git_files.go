package toolGit

import (
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
)


// Usage:
//		{{- $cmd := $git.ChangedFiles }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) ChangedFiles() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		g.State = g.Exec(gitCommandStatus, "--porcelain")
		if g.State.IsError() {
			break
		}

		var fps []string
		fps = make([]string, len(g.State.OutputArray))
		for i, fp := range g.State.OutputArray {
			s := strings.Fields(fp)
			if len(s) > 1 {
				fps[i] = s[1]
			}
		}

		g.State.SetResponse(&fps)
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.AddFiles }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) AddFiles(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		if len(args) == 0 {
			args = []string{"."}
		}

		//args = append([]string{"--porcelain"}, args...)
		g.State = g.Exec(gitCommandAdd, args...)
		if g.State.IsNotOk() {
			break
		}

		var fps []string
		fps = make([]string, len(g.State.OutputArray))
		for i, fp := range g.State.OutputArray {
			s := strings.Fields(fp)
			if len(s) > 1 {
				fps[i] = s[1]
			}
		}

		g.State.SetResponse(&fps)
	}

	return g.State
}
