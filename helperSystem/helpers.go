package helperSystem

import "github.com/newclarity/scribeHelpers/ux"

type HelperSystem TypeSystem
func (s *HelperSystem) Reflect() *TypeSystem {
	return (*TypeSystem)(s)
}
func (s *TypeSystem) Reflect() *HelperSystem {
	return (*HelperSystem)(s)
}

func (s *HelperSystem) IsNil() *ux.State {
	if state := ux.IfNilReturnError(s); state.IsError() {
		return state
	}
	s.State = s.State.EnsureNotNil()
	return s.State
}


// Usage:
//		{{ $sys := NewSystem }}
func HelperNewSystem() *HelperSystem {
	ret := New(false)

	for range OnlyOnce {
	}

	return ReflectHelperSystem(ret)
}
