package toolGit

import "github.com/newclarity/scribeHelpers/ux"

type ToolGit TypeGit
func (g *ToolGit) IsNil() *ux.State {
	return ux.IfNilReturnError(g)
}

func (g *ToolGit) Reflect() *TypeGit {
	return (*TypeGit)(g)
}

func (g *TypeGit) Reflect() *ToolGit {
	return (*ToolGit)(g)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolGit(p *TypeGit) *ToolGit {
	return (*ToolGit)(p)
}
