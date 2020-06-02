package helperService

import (
	"github.com/newclarity/scribeHelpers/helperPath"
	"github.com/newclarity/scribeHelpers/ux"
)

type ExampleGetter interface {
}


type TypeExample struct {
	name  string
	path  *helperPath.TypeOsPath

	State *ux.State
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}
func ReflectHelperExample(p *TypeExample) *HelperExample {
	return (*HelperExample)(p)
}

func (c *TypeExample) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}

