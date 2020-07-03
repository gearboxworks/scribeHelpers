package toolSelfUpdate

import "github.com/newclarity/scribeHelpers/ux"


type ToolSelfUpdate TypeSelfUpdate
func (su *ToolSelfUpdate) IsNil() *ux.State {
	return ux.IfNilReturnError(su)
}

func (su *ToolSelfUpdate) Reflect() *TypeSelfUpdate {
	return (*TypeSelfUpdate)(su)
}

func (su *TypeSelfUpdate) Reflect() *ToolSelfUpdate {
	return (*ToolSelfUpdate)(su)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolSelfUpdate(p *TypeSelfUpdate) *ToolSelfUpdate {
	return (*ToolSelfUpdate)(p)
}
