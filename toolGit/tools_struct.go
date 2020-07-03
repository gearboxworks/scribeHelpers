package toolGit

import "github.com/newclarity/scribeHelpers/ux"

type ToolGit TypeGit
func (c *ToolGit) IsNil() *ux.State {
	return ux.IfNilReturnError(c)
}

func (c *ToolGit) Reflect() *TypeGit {
	return (*TypeGit)(c)
}

func (c *TypeGit) Reflect() *ToolGit {
	return (*ToolGit)(c)
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
