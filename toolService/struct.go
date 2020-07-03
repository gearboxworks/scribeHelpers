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
func (s *TypeService) IsNil() *ux.State {
	return ux.IfNilReturnError(s)
}


func New(runtime *toolRuntime.TypeRuntime) *TypeService {
	runtime = runtime.EnsureNotNil()

	s := TypeService{
		name: "",
		path:   toolPath.New(runtime),

		Debug:  runtime.Debug,
		State:  ux.NewState(runtime.CmdName, runtime.Debug),
	}
	s.State.SetPackage("")
	s.State.SetFunctionCaller()
	return &s
}

func (s *TypeService) EnsureNotNil() *TypeService {
	if s == nil {
		return New(nil)
	}
	return s
}
