package toolGear

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/newclarity/scribeHelpers/ux"
)

type ToolDockerGear DockerGear
func (c *ToolDockerGear) Reflect() *DockerGear {
	return (*DockerGear)(c)
}
func (gear *DockerGear) Reflect() *ToolDockerGear {
	return (*ToolDockerGear)(gear)
}

func (c *ToolDockerGear) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}


// Usage:
//		{{ $copy := CopyFiles }}
func ToolNewGear() *ToolDockerGear {
	ret := New(nil)

	for range onlyOnce {
		//ret.State.SetOk()
	}

	return (*ToolDockerGear)(ret)
}


func (c *ToolDockerGear) List() string {
	var ret string
	for range onlyOnce {
		c.State = c.Reflect().List("")
		if c.State.IsError() {
			break
		}
	}
	return ret
}


func (c *ToolDockerGear) ParseGearConfig(cs interface{}) string {
	var ret string
	for range onlyOnce {
		err := mapstructure.Decode(cs, &c.gearConfig)
		if err == nil {
			c.State.SetOk()
			break
		}

		//csString := toolTypes.ReflectString(cs)
		//if csString != nil {
		//	c.State = c.Reflect().ParseGearConfig(*csString)
		//	if c.State.IsOk() {
		//		//ret = c.State.SprintError()
		//		ret = "cs == toolTypes.ReflectString"
		//		break
		//	}
		//}
		//
		//csStruct := gearConfig.ReflectGearConfig(cs)
		//if csStruct != nil {
		//	c.gearConfig = csStruct
		//	ret = "cs == gearConfig.ReflectGearConfig"
		//	break
		//}
		//
		//csStructs := gearConfig.ReflectGearConfigs(cs)
		//if csStructs != nil {
		//	err := mapstructure.Decode(csStructs, &c.gearConfig)
		//	c.State.SetError(err)
		//	if c.State.IsError() {
		//		break
		//	}
		//	break
		//}

		ret = ux.SprintfRed("Invalid Gear config\n")
	}
	return ret
}


func (c *ToolDockerGear) PrintGearConfig() string {
	return fmt.Sprintf("%v", c.gearConfig)
}
