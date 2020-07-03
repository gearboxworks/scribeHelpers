package toolDocker

import "github.com/newclarity/scribeHelpers/ux"


type ToolDocker TypeDocker
func (d *ToolDocker) IsNil() *ux.State {
	return ux.IfNilReturnError(d)
}

func (d *ToolDocker) Reflect() *TypeDocker {
	return (*TypeDocker)(d)
}

func (d *TypeDocker) Reflect() *ToolDocker {
	return (*ToolDocker)(d)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolDocker(p *TypeDocker) *ToolDocker {
	return (*ToolDocker)(p)
}
