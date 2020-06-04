package toolService

import "github.com/newclarity/scribeHelpers/ux"


type ToolService TypeService
func (g *ToolService) Reflect() *TypeService {
	return (*TypeService)(g)
}
func (s *TypeService) Reflect() *ToolService {
	return (*ToolService)(s)
}

func (c *ToolService) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}
