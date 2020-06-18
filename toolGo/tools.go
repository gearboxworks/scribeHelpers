package toolGo

import "github.com/newclarity/scribeHelpers/ux"


type ToolGo TypeGo
func (e *ToolGo) Reflect() *TypeGo {
	return (*TypeGo)(e)
}
func (g *TypeGo) Reflect() *ToolGo {
	return (*ToolGo)(g)
}

func (e *ToolGo) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}
