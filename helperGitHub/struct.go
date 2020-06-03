package helperGitHub

import (
	"github.com/google/go-github/v31/github"
	"github.com/newclarity/scribeHelpers/ux"
)


type TypeGitHub struct {
	User *github.User
	Client *github.Client
	Valid bool

	Debug  bool
	State *ux.State
}


func New(debugMode bool) *TypeGitHub {
	ret := &TypeGitHub {
		User:   nil,
		Client: nil,
		Valid:  true,

		Debug:  debugMode,
		State:  ux.NewState(debugMode),
	}
	ret.State.SetPackage("")
	ret.State.SetFunctionCaller()

	return ret
}


func (gh *TypeGitHub) IsNil() *ux.State {
	if state := ux.IfNilReturnError(gh); state.IsError() {
		return state
	}
	gh.State = gh.State.EnsureNotNil()
	return gh.State
}


func (gh *TypeGitHub) EnsureNotNil() *TypeGitHub {
	if gh == nil {
		return New(true)
	}
	return gh
}
