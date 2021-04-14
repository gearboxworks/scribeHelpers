package toolGitHub

import "github.com/gearboxworks/scribeHelpers/ux"


type ToolGitHub TypeGitHub
func (gh *ToolGitHub) IsNil() *ux.State {
	return ux.IfNilReturnError(gh)
}

func (gh *ToolGitHub) Reflect() *TypeGitHub {
	return (*TypeGitHub)(gh)
}

func (gh *TypeGitHub) Reflect() *ToolGitHub {
	return (*ToolGitHub)(gh)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolGitHub(p *TypeGitHub) *ToolGitHub {
	return (*ToolGitHub)(p)
}
