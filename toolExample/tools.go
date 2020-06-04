package toolExample

import "github.com/newclarity/scribeHelpers/ux"


type ToolExample TypeExample
func (g *ToolExample) Reflect() *TypeExample {
	return (*TypeExample)(g)
}
func (g *TypeExample) Reflect() *ToolExample {
	return (*ToolExample)(g)
}

func (c *ToolExample) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}
