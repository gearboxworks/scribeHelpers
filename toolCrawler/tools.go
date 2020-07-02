package toolCrawler

import "github.com/newclarity/scribeHelpers/ux"


type ToolExample TypeExample
func (e *ToolExample) Reflect() *TypeExample {
	return (*TypeExample)(e)
}
func (e *TypeExample) Reflect() *ToolExample {
	return (*ToolExample)(e)
}

func (e *ToolExample) IsNil() *ux.State {
	if state := ux.IfNilReturnError(e); state.IsError() {
		return state
	}
	e.State = e.State.EnsureNotNil()
	return e.State
}
