package toolGit

import (
	"github.com/gearboxworks/scribeHelpers/ux"
)


// Usage:
//		{{- $cmd := $git.GitClone }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
// func (me *ToolGit) GitClone(args ...interface{}) *ux.State {
func (g *TypeGit) GitClone(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandClone, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitClean }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
// func (me *ToolGit) GitClone(args ...interface{}) *ux.State {
func (g *TypeGit) GitClean(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandClean, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitInit }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitInit(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandInit, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitAdd }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitAdd(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandAdd, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitMv }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitMv(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandMv, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitReset }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitReset(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandReset, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitRm }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitRm(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandRm, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitBisect }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitBisect(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandBisect, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitGrep }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitGrep(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandGrep, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitLog }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitLog(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandLog, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitShow }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitShow(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandShow, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitStatus }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitStatus(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandStatus, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitBranch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitBranch(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandBranch, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitCheckout }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitCheckout(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandCheckout, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitCommit }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitCommit(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandCommit, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitDiff }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitDiff(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandDiff, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitMerge }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitMerge(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandMerge, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitRebase }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitRebase(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandRebase, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitTag }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitTag(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandTag, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitFetch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitFetch(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandFetch, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitPull }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitPull(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandPull, args...)
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitPush }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GitPush(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()
	g.State = g.Exec(gitCommandPush, args...)
	return g.State
}
