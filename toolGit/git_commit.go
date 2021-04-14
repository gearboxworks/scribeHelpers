package toolGit

import (
	"github.com/gearboxworks/scribeHelpers/ux"
)


// Usage:
//		{{- $cmd := $git.Commit }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) LastCommitId() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		g.State = g.Exec(gitCommandRevParse, "--verify", "HEAD")
		if g.State.IsError() {
			break
		}

		g.State.SetResponse(g.State.Output)
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.LastCommitMessage }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) LastCommitMessage() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		g.State = g.Exec(gitCommandLog, "-1", "--pretty=%B")
		if g.State.IsError() {
			break
		}

		g.State.SetResponse(g.State.Output)
	}

	return g.State
}
