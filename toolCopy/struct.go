package toolCopy

import (
	"github.com/gearboxworks/scribeHelpers/toolPath"
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
)


// @TODO - Look at several other copy options that provide "cloud" based copies.
// @TODO - https://rclone.org/
// @TODO - https://pkg.go.dev/github.com/Redundancy/go-sync?tab=doc
// @TODO - https://github.com/Redundancy/go-sync
// @TODO - https://pkg.go.dev/bitbucket.org/kardianos/rsync?tab=doc
// @TODO - https://github.com/zloylos/grsync


type OsCopyGetter interface {
}

type TypeOsCopy struct {
	Method *TypeCopyMethods				`json:"method" mapstructure:"method"`

	Paths  TypeOsCopyPaths

	Valid  bool							`json:"valid" mapstructure:"valid"`
	Debug  bool							`json:"debug" mapstructure:"debug"`
	State  *ux.State					`json:"state" mapstructure:"state"`
}
func (c *TypeOsCopy) IsNil() *ux.State {
	return ux.IfNilReturnError(c)
}

type TypeOsCopyPaths struct {
	Source       *toolPath.TypeOsPath	`json:"source" mapstructure:"source"`
	Destination  *toolPath.TypeOsPath	`json:"destination" mapstructure:"destination"`

	Exclude PathArray					`json:"exclude" mapstructure:"exclude"`
	Include PathArray					`json:"include" mapstructure:"include"`
}


func New(runtime *toolRuntime.TypeRuntime) *TypeOsCopy {
	runtime = runtime.EnsureNotNil()

	c := &TypeOsCopy{
		Paths: TypeOsCopyPaths {
			Source: toolPath.New(runtime),
			Destination: toolPath.New(runtime),

			Exclude: PathArray{},
			Include: PathArray{},
		},

		Method: NewCopyMethod(),

		Valid:  false,
		Debug:  runtime.Debug,
		State:  ux.NewState(runtime.CmdName, runtime.Debug),
	}
	c.State.SetPackage("")
	c.State.SetFunctionCaller()
	return c
}


func (c *TypeOsCopy) EnsureNotNil() *TypeOsCopy {
	if c == nil {
		return New(nil)
	}
	return c
}
