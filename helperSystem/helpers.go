package helperSystem

import "github.com/newclarity/scribeHelpers/ux"

type HelperSystem TypeSystem
func (p *HelperSystem) Reflect() *TypeSystem {
	return (*TypeSystem)(p)
}
func (p *TypeSystem) Reflect() *HelperSystem {
	return (*HelperSystem)(p)
}

func (p *HelperSystem) IsNil() *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}
	p.State = p.State.EnsureNotNil()
	return p.State
}


// Usage:
//		{{ $sys := NewSystem }}
func HelperNewSystem() *HelperSystem {
	ret := NewSystem(false)

	for range OnlyOnce {
	}

	return ReflectHelperSystem(ret)
}
