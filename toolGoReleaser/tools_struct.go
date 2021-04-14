package toolGoReleaser

import "github.com/gearboxworks/scribeHelpers/ux"


type ToolGoReleaser TypeGoReleaser
func (c *ToolGoReleaser) IsNil() *ux.State {
	return ux.IfNilReturnError(c)
}

func (c *ToolGoReleaser) Reflect() *TypeGoReleaser {
	return (*TypeGoReleaser)(c)
}

func (c *TypeGoReleaser) Reflect() *ToolGoReleaser {
	return (*ToolGoReleaser)(c)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolGoReleaser(p *TypeGoReleaser) *ToolGoReleaser {
	return (*ToolGoReleaser)(p)
}
