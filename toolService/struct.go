package toolService

import (
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)

type ServiceGetter interface {
}


type TypeService struct {
	name  string
	path  *toolPath.TypeOsPath

	Debug bool
	State *ux.State
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}
func ReflectToolService(p *TypeService) *ToolService {
	return (*ToolService)(p)
}


func New(runtime *toolRuntime.TypeRuntime) *TypeService {
	runtime = runtime.EnsureNotNil()

	s := TypeService{
		name: "",
		path:   toolPath.New(runtime.Debug),

		Debug:  runtime.Debug,
		State:  ux.NewState(runtime.CmdName, runtime.Debug),
	}
	s.State.SetPackage("")
	s.State.SetFunctionCaller()
	return &s
}


func (s *TypeService) IsNil() *ux.State {
	if state := ux.IfNilReturnError(s); state.IsError() {
		return state
	}
	s.State = s.State.EnsureNotNil()
	return s.State
}


func (s *TypeService) EnsureNotNil() *TypeService {
	if s == nil {
		return New(true)
	}
	return s
}
