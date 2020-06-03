package toolService

import "github.com/newclarity/scribeHelpers/ux"


type HelperService TypeService
func (g *HelperService) Reflect() *TypeService {
	return (*TypeService)(g)
}
func (s *TypeService) Reflect() *HelperService {
	return (*HelperService)(s)
}

func (c *HelperService) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}
