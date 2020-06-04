package toolGit

import (
	"github.com/newclarity/scribeHelpers/ux"
)


// Usage:
//		{{- $cmd := $git.GitClone }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
// func (me *ToolGit) GitClone(args ...interface{}) *ux.State {
func (g *ToolGit) GitClone(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandClone, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitInit }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitInit(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandInit, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitAdd }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitAdd(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandAdd, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitMv }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitMv(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandMv, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitReset }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitReset(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandReset, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitRm }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitRm(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandRm, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitBisect }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitBisect(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandBisect, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitGrep }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitGrep(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandGrep, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitLog }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitLog(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandLog, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitShow }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitShow(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandShow, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitStatus }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitStatus(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandStatus, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitBranch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitBranch(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandBranch, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitCheckout }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitCheckout(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandCheckout, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitCommit }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitCommit(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandCommit, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitDiff }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitDiff(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandDiff, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitMerge }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitMerge(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandMerge, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitRebase }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitRebase(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandRebase, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitTag }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitTag(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandTag, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitFetch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitFetch(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandFetch, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitPull }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitPull(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandPull, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitPush }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *ToolGit) GitPush(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandPush, args...))
	return g.State
}
