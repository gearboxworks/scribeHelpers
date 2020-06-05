package toolGit

import (
	"github.com/newclarity/scribeHelpers/ux"
)

type Commit struct {
	Hash string
}


func _NewCommit(hash string) *Commit {
	return &Commit{
		Hash: hash,
	}
}


// Usage:
//		{{- $cmd := $git.Commit }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) Commit(format interface{}, a ...interface{}) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range onlyOnce {
		g.State = g.Exec("rev-parse", "--verify", "HEAD")
		if g.State.IsError() {
			break
		}

		g.State.Response = _NewCommit(g.State.Output)
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.LastCommitMessage }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) LastCommitMessage(format interface{}, a ...interface{}) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range onlyOnce {
		g.State = g.Exec("git", "log", "-1", "--pretty=%B")
		if g.State.IsError() {
			break
		}

		g.State.Response = _NewCommit(g.State.Output)
	}

	return g.State
}
