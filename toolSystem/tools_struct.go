package toolSystem

import "github.com/newclarity/scribeHelpers/ux"


type ToolSystem TypeSystem
func (s *ToolSystem) IsNil() *ux.State {
	return ux.IfNilReturnError(s)
}

func (s *ToolSystem) Reflect() *TypeSystem {
	return (*TypeSystem)(s)
}

func (s *TypeSystem) Reflect() *ToolSystem {
	return (*ToolSystem)(s)
}


type ToolProcesses TypeProcesses
func (p *ToolProcesses) IsNil() *ux.State {
	return ux.IfNilReturnError(p)
}

func (p *ToolProcesses) Reflect() *TypeProcesses {
	return (*TypeProcesses)(p)
}

func (p *TypeProcesses) Reflect() *ToolProcesses {
	return (*ToolProcesses)(p)
}


type ToolProcess TypeProcess
func (p *ToolProcess) IsNil() *ux.State {
	return ux.IfNilReturnError(p)
}

func (p *ToolProcess) Reflect() *TypeProcess {
	return (*TypeProcess)(p)
}

func (p *TypeProcess) Reflect() *ToolProcess {
	return (*ToolProcess)(p)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolSystem(p *TypeSystem) *ToolSystem {
	return (*ToolSystem)(p)
}
