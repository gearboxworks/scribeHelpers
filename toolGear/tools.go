package toolGear

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/newclarity/scribeHelpers/ux"
)


// Usage:
//		{{ $copy := CopyFiles }}
func ToolNewGear() *Gear {
	ret := NewGear(nil)

	for range onlyOnce {
		//ret.State.SetOk()
	}

	return (*Gear)(ret)
}


//func (c *TypeGears) List() string {
//	var ret string
//	for range onlyOnce {
//		c.State = c.Reflect().List("")
//		if c.State.IsError() {
//			break
//		}
//	}
//	return ret
//}


func (c *TypeGear) ParseGearConfig(cs interface{}) string {
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


func (c *Gear) PrintGearConfig() string {
	return fmt.Sprintf("%v", c.GearConfig)
}
