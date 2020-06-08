package toolGhr

import "github.com/newclarity/scribeHelpers/ux"


type ToolGhr TypeGhr
func (ghr *ToolGhr) Reflect() *TypeGhr {
	return (*TypeGhr)(ghr)
}
func (ghr *TypeGhr) Reflect() *ToolGhr {
	return (*ToolGhr)(ghr)
}

func (ghr *ToolGhr) IsNil() *ux.State {
	if state := ux.IfNilReturnError(ghr); state.IsError() {
		return state
	}
	ghr.State = ghr.State.EnsureNotNil()
	return ghr.State
}
