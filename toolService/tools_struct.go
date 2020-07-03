package toolService

import "github.com/newclarity/scribeHelpers/ux"


type ToolService TypeService
func (s *ToolService) IsNil() *ux.State {
	return ux.IfNilReturnError(s)
}

func (s *ToolService) Reflect() *TypeService {
	return (*TypeService)(s)
}

func (s *TypeService) Reflect() *ToolService {
	return (*ToolService)(s)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolService(p *TypeService) *ToolService {
	return (*ToolService)(p)
}
