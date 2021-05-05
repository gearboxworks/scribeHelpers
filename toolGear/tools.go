package toolGear

import (
	"fmt"
	"github.com/gearboxworks/scribeHelpers/toolGear/gearConfig"
	"github.com/gearboxworks/scribeHelpers/ux"
	"github.com/mitchellh/mapstructure"
	"os"
)


// Usage:
//		{{ $copy := CopyFiles }}
func ToolGearbox(gcm interface{}, remote string) *Gears {
	var gears Gears

	for range onlyOnce {
		//if remote != "" {
		//	_ = os.Setenv("DOCKER_HOST", remote)
			ux.PrintfWhite("# Remote Docker is %s\n", os.Getenv("DOCKER_HOST"))
		//}

		gears = NewGears(nil)
		if gears.State.IsNotOk() {
			break
		}

		//gears.State = gears.SetProvider("docker")
		//if gears.State.IsError() {
		//	break
		//}

		gears.State = gears.SetProviderUrl(remote)
		if gears.State.IsNotOk() {
			break
		}

		gears.State = gears.Get()
		if gears.State.IsNotOk() {
			break
		}

		if gcm != nil {
			gc := NewGearConfig(gears.Runtime)
			err := mapstructure.Decode(gcm, &gc)
			if err != nil {
				gears.State.SetError(err)
				break
			}

			gears.State = gears.AddGears(gc)

			//gears.FindImage("mountebank", "2.4.0")
			//gears.Selected.Logs().GetOutput()
		}
	}

	if gears.State.IsNotOk() {
		gears.State.PrintResponse()
	}

	return &gears
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

		for k := range gear.GearConfig.Versions {
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

//goland:noinspection GoUnusedExportedFunction
func ToolNewGear() *Gear {
	ret := NewGear(nil, nil)

	for range onlyOnce {
		//ret.State.SetOk()
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
