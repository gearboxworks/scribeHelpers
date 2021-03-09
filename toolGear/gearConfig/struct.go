package gearConfig

import (
	"encoding/json"
	"fmt"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)

const onlyOnce = "1"

const DefaultCommandName = "default"


type GearConfig struct {
	Meta       GearMeta       `json:"meta"`
	Build      GearBuild      `json:"build"`
	Run        GearRun        `json:"run"`
	Project    GearProject    `json:"project"`
	Extensions GearExtensions `json:"extensions"`
	Versions   GearVersions   `json:"versions"`

	Schema     string         `json:"schema"`

	Runtime    *toolRuntime.TypeRuntime
	State      *ux.State
}
func (gc *GearConfig) IsNil() *ux.State {
	return ux.IfNilReturnError(gc)
}


type GearConfigs map[string]GearConfig


func New(runtime *toolRuntime.TypeRuntime) *GearConfig {
	runtime = runtime.EnsureNotNil()

	gc := GearConfig{
		Meta:       GearMeta{},
		Build:      GearBuild{},
		Run:        GearRun{},
		Project:    GearProject{},
		Extensions: GearExtensions{},
		Versions:   nil,
		Schema:     "",
		Runtime:    runtime,
		State:      ux.NewState(runtime.CmdName, runtime.Debug),
	}
	gc.State.SetPackage("")
	gc.State.SetFunctionCaller()
	return &gc
}


func (gc *GearConfig) EnsureNotNil() *GearConfig {
	if gc == nil {
		return New(nil)
	}
	return gc
}


func (gc *GearConfig) IsValid() *ux.State {
	if state := ux.IfNilReturnError(gc); state.IsError() {
		return state
	}

	for range onlyOnce {
		gc.State = gc.State.EnsureNotNil()

		//if gc == nil {
		//	gc.State.SetError("gear config is nil")
		//	break
		//}
	}

	return gc.State
}


func (gc *GearConfig) ParseJson(cs string) *ux.State {
	if gc == nil {
		gc = gc.EnsureNotNil()
		ux.PrintfRed("GearConfig is nil!\n")
		return gc.State
	}

	for range onlyOnce {
		gc.State = gc.State.EnsureNotNil()

		if cs == "" {
			gc.State.SetError("gear config is empty")
			break
		}

		js := []byte(cs)
		if js == nil {
			gc.State.SetError("gear config json is nil")
			break
		}

		err := json.Unmarshal(js, &gc)
		if err != nil {
			gc.State.SetError("gearbox.json schema unknown: %s", err)
			break
		}
	}

	return gc.State
}


func (gc *GearConfig) IsMatchedGear(gearName string, gearVersion string, tagVersions []string) bool {
	var ok bool

	for range onlyOnce {
		if gc.Meta.Organization != defaultOrganization {
			break
		}

		if gc.Meta.Name != gearName {
			break
		}

		if !gc.Versions.HasVersion(gearVersion) {
			break
		}

		nameCheck := fmt.Sprintf("%s/%s:%s", defaultOrganization, gearName, gearVersion)
		for _, s := range tagVersions {
			if s == nameCheck {
				ok = true
				break
			}
		}
	}

	return ok
}


func (gc *GearConfig) String() string {
	var ret string
	if state := ux.IfNilReturnError(gc); state.IsError() {
		return ux.SprintfRed("GearConfig is empty.\n")
	}

	for range onlyOnce {
		if gc.Schema == "" {
			ret = ux.SprintfRed("GearConfig is empty.\n")
			break
		}

		ret += ux.SprintfBlue("Schema: %s\n", gc.Schema)
		ret += ux.SprintfBlue("%v", gc.Meta.String())
		ret += ux.SprintfBlue("%v", gc.Build.String())
		ret += ux.SprintfBlue("%v", gc.Run.String())
		ret += ux.SprintfBlue("%v", gc.Project.String())
		ret += ux.SprintfBlue("%v", gc.Extensions.String())
		ret += ux.SprintfBlue("%v", gc.Versions.String())
	}

	return ret
}


func (gc *GearConfig) GetClass() string {
	if gc == nil {
		return ""
	}
	return gc.Meta.Class
}

func (gc *GearConfig) GetName() string {
	if gc == nil {
		return ""
	}
	return gc.Meta.Name
}

func (gc *GearConfig) GetPorts() *GearPorts {
	if gc == nil {
		return &GearPorts{}
	}
	return &gc.Build.Ports
}

func (gc *GearConfig) GetCommand(cmd []string) []string {
	var retCmd []string

	for range onlyOnce {
		var cmdExec string
		switch {
			case len(cmd) == 0:
				cmdExec = DefaultCommandName

			case cmd[0] == "":
				cmdExec = DefaultCommandName

			case cmd[0] == gc.Meta.Name:
				cmdExec = DefaultCommandName

			case cmd[0] != "":
				cmdExec = cmd[0]

			default:
				//cmdExec = cmd[0]
				cmdExec = DefaultCommandName
		}

		c := gc.MatchCommand(cmdExec)
		if c == nil {
			retCmd = []string{}
			break
		}

		retCmd = append([]string{*c}, cmd[1:]...)
	}

	return retCmd
}

func (gc *GearConfig) MatchCommand(cmd string) *string {
	var c *string

	for range onlyOnce {
		if c2, ok := gc.Run.Commands[cmd]; ok {
			c = &c2
			break
		}
	}

	return c
}
