package helperCopy

import (
	"github.com/newclarity/scribeHelpers/helperPath"
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

type TypeOsPath helperPath.TypeOsPath

type TypeOsCopy struct {
	Source       *helperPath.TypeOsPath	`json:"source" mapstructure:"source"`
	Destination  *helperPath.TypeOsPath	`json:"destination" mapstructure:"destination"`

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
func ReflectHelperOsCopy(p *TypeOsCopy) *HelperOsCopy {
	return (*HelperOsCopy)(p)
}


func New(debugFlag bool) *TypeOsCopy {
	c := &TypeOsCopy{
		Source:       helperPath.New(debugFlag),
		Destination:  helperPath.New(debugFlag),

		Exclude: PathArray{},
		Include: PathArray{},

		Method: NewCopyMethod(),

		Valid:  false,
		Debug:  debugFlag,
		State:  ux.NewState(debugFlag),
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
		return New(true)
	}
	return c
}
