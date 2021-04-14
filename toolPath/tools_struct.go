package toolPath

import "github.com/gearboxworks/scribeHelpers/ux"


type ToolOsPath TypeOsPath
func (p *ToolOsPath) IsNil() *ux.State {
	return ux.IfNilReturnError(p)
}

func (p *ToolOsPath) Reflect() *TypeOsPath {
	return (*TypeOsPath)(p)
}

func (p *TypeOsPath) Reflect() *ToolOsPath {
	return (*ToolOsPath)(p)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolOsPath(p *TypeOsPath) *ToolOsPath {
	return (*ToolOsPath)(p)
}
