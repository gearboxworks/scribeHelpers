package toolExample

import "github.com/newclarity/scribeHelpers/ux"


type ToolExample TypeExample
func (e *ToolExample) IsNil() *ux.State {
return ux.IfNilReturnError(e)
}

func (e *ToolExample) Reflect() *TypeExample {
return (*TypeExample)(e)
}

func (e *TypeExample) Reflect() *ToolExample {
return (*ToolExample)(e)
}


type State ux.State

func (s *State) Reflect() *ux.State {
return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
return (*State)(p)
}

func ReflectToolExample(p *TypeExample) *ToolExample {
return (*ToolExample)(p)
}
