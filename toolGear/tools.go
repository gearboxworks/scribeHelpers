package toolGear

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/newclarity/scribeHelpers/toolGear/gearConfig"
	"github.com/newclarity/scribeHelpers/ux"
)


// Usage:
//		{{ $copy := CopyFiles }}
func ToolNewGear() *Gear {
	ret := NewGear(nil, nil)

	for range onlyOnce {
		//ret.State.SetOk()
	}

	return (*Gear)(ret)
}


func ToolListVersions(gc interface{}) string {
	var ret string

	for range onlyOnce {
		gear := NewGear(nil, nil)
		err := mapstructure.Decode(gc, &gear.GearConfig)
		if err != nil {
			gear.State.SetError(err)
			break
		}

		for k, _ := range gear.GearConfig.Versions {
			ret += k + " "
		}
	}

	return ret
}


func ToolListVersionInfo(gc interface{}) string {
	var ret string

	for range onlyOnce {
		gear := NewGear(nil, nil)
		err := mapstructure.Decode(gc, &gear.GearConfig)
		if err != nil {
			gear.State.SetError(err)
			break
		}

		for k, v := range gear.GearConfig.Versions {
			if v.Latest {
				ret += ux.SprintfWhite("\t*")
			} else {
				ret += "\t "
			}

			ret += ux.SprintfBlue("%s", k)
			ret += " - "
			ret += ux.SprintfCyan("%s/%s:%s\n",
				gear.GearConfig.Meta.Organization,
				gear.GearConfig.Meta.Name,
				k,
				)
		}
	}

	return ret
}


func ToolShowGear(gc interface{}) string {
	var ret string

	for range onlyOnce {
		gear := NewGear(nil, nil)
		err := mapstructure.Decode(gc, &gear.GearConfig)
		if err != nil {
			gear.State.SetError(err)
			break
		}

		//gear.State = gear.ParseGearConfig(data)
		ret = gear.PrintGearConfig()
		fmt.Print(ret)
	}

	return ret
}


func ReflectGearConfig(ref interface{}) gearConfig.GearConfig {
	return ref.(gearConfig.GearConfig)
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


func (gear *TypeGear) ParseGearConfig(cs interface{}) string {
	var ret string
	for range onlyOnce {
		err := mapstructure.Decode(cs, &gear.gearConfig)
		if err == nil {
			gear.State.SetOk()
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


func (gear *Gear) PrintGearConfig() string {
	return fmt.Sprintf("%v", gear.GearConfig)
}
