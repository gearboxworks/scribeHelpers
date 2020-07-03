package toolSystem


// Usage:
//		{{ $sys := NewSystem }}
func ToolNewSystem() *ToolSystem {
	ret := New(nil)

	for range onlyOnce {
	}

	return ReflectToolSystem(ret)
}
