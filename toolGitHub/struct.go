package toolGitHub

import (
	"github.com/google/go-github/v31/github"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type TypeGitHub struct {
	User *github.User
	Client *github.Client
	Valid bool

	Debug  bool
	State *ux.State
}


func New(runtime *toolRuntime.TypeRuntime) *TypeGitHub {
	runtime = runtime.EnsureNotNil()

	gh := &TypeGitHub {
		User:   nil,
		Client: nil,
		Valid:  true,

		Debug:  runtime.Debug,
		State:  ux.NewState(runtime.CmdName, runtime.Debug),
	}
	gh.State.SetPackage("")
	gh.State.SetFunctionCaller()
	return gh
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
