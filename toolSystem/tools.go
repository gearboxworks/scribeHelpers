package toolSystem

import "github.com/newclarity/scribeHelpers/ux"

type ToolSystem TypeSystem
func (s *ToolSystem) Reflect() *TypeSystem {
	return (*TypeSystem)(s)
}
func (s *TypeSystem) Reflect() *ToolSystem {
	return (*ToolSystem)(s)
}

func (s *ToolSystem) IsNil() *ux.State {
	if state := ux.IfNilReturnError(s); state.IsError() {
		return state
	}
	s.State = s.State.EnsureNotNil()
	return s.State
}


// Usage:
//		{{ $sys := NewSystem }}
func ToolNewSystem() *ToolSystem {
	ret := New(false)

	for range OnlyOnce {
	}

	return ReflectToolSystem(ret)
}
