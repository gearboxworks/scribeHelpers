package toolCopy

import (
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


// @TODO - Look at several other copy options that provide "cloud" based copies.
// @TODO - https://rclone.org/
// @TODO - https://pkg.go.dev/github.com/Redundancy/go-sync?tab=doc
// @TODO - https://github.com/Redundancy/go-sync
// @TODO - https://pkg.go.dev/bitbucket.org/kardianos/rsync?tab=doc
// @TODO - https://github.com/zloylos/grsync


type OsCopyGetter interface {
}

type TypeOsPath toolPath.TypeOsPath

type TypeOsCopy struct {
	Source       *toolPath.TypeOsPath	`json:"source" mapstructure:"source"`
	Destination  *toolPath.TypeOsPath	`json:"destination" mapstructure:"destination"`

	Exclude PathArray					`json:"exclude" mapstructure:"exclude"`
	Include PathArray					`json:"include" mapstructure:"include"`

	Method *TypeCopyMethods				`json:"method" mapstructure:"method"`

	Valid  bool							`json:"valid" mapstructure:"valid"`
	Debug  bool							`json:"debug" mapstructure:"debug"`
	State  *ux.State					`json:"state" mapstructure:"state"`
}


type State ux.State
func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}
//func ReflectState(p *ux.State) *ux.State {
//	return (*State)(p)
//}
func ReflectToolOsCopy(p *TypeOsCopy) *ToolOsCopy {
	return (*ToolOsCopy)(p)
}


func New(runtime *toolRuntime.TypeRuntime) *TypeOsCopy {
	runtime = runtime.EnsureNotNil()

	c := &TypeOsCopy{
		Source:       toolPath.New(runtime),
		Destination:  toolPath.New(runtime),

		Exclude: PathArray{},
		Include: PathArray{},

		Method: NewCopyMethod(),

		Valid:  false,
		Debug:  runtime.Debug,
		State:  ux.NewState(runtime.CmdName, runtime.Debug),
	}
	c.State.SetPackage("")
	c.State.SetFunctionCaller()
	return c
}


func (c *TypeOsCopy) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}


func (c *TypeOsCopy) EnsureNotNil() *TypeOsCopy {
	if c == nil {
		return New(nil)
	}
	return c
}
