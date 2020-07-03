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
func (gh *TypeGitHub) IsNil() *ux.State {
	return ux.IfNilReturnError(gh)
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


func (gh *TypeGitHub) EnsureNotNil() *TypeGitHub {
	if gh == nil {
		return New(nil)
	}
	return gh
}
