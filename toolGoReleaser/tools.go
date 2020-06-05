package toolGoReleaser

import "github.com/newclarity/scribeHelpers/ux"


type ToolGoReleaser TypeGoReleaser
func (g *ToolGoReleaser) Reflect() *TypeGoReleaser {
	return (*TypeGoReleaser)(g)
}
func (gr *TypeGoReleaser) Reflect() *ToolGoReleaser {
	return (*ToolGoReleaser)(gr)
}

func (c *ToolGoReleaser) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}
