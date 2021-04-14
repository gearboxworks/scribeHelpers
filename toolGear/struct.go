package toolGear

import (
	"github.com/gearboxworks/scribeHelpers/toolGear/gearConfig"
	"github.com/gearboxworks/scribeHelpers/toolGear/gearSsh"
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
)


type TypeGear struct {
	Image     *Image
	Container *Container
	gearConfig *gearConfig.GearConfig

	Ssh       *gearSsh.Ssh

	Runtime   *toolRuntime.TypeRuntime
	State     *ux.State
}
type GearConfigs []*gearConfig.GearConfig
//type TypeDockerGears []TypeDockerGear

//func (gear *TypeDockerGear) IsNil() *ux.State {
//	return ux.IfNilReturnError(gear)
//}


type TypeMatchImage struct {
	Organization string
	Name         string
	Version      string
}
type TypeMatchContainer TypeMatchImage
