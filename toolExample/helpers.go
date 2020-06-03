package toolExample

import "github.com/newclarity/scribeHelpers/ux"


type HelperExample TypeExample
func (g *HelperExample) Reflect() *TypeExample {
	return (*TypeExample)(g)
}
func (g *TypeExample) Reflect() *HelperExample {
	return (*HelperExample)(g)
}

func (c *HelperExample) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}
