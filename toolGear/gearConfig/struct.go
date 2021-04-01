package gearConfig

import (
	"encoding/json"
	"fmt"
	"github.com/docker/go-connections/nat"
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


func (gc *GearConfig) IsValid() bool {// *ux.State {
	var ok bool
	if state := ux.IfNilReturnError(gc); state.IsError() {
		return ok
	}

	for range onlyOnce {
		//if gc == nil {
		//	gc.State.SetError("gear config is nil")
		//	break
		//}

		gc.State = gc.State.EnsureNotNil()
		if gc.Meta.Name == "" {
			gc.State.SetError("gear config has no name")
			break
		}

		if len(gc.Versions) == 0 {
			gc.State.SetError("gear config has no versions")
			break
		}

		ok = true
	}

	return ok
}
func (gc *GearConfig) IsNotValid() bool {
	return !gc.IsValid()
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

func (gc *GearConfig) GetFixedPorts() nat.PortMap {
	var ports nat.PortMap
	if state := gc.IsNil(); state.IsError() {
		return nil
	}

	for range onlyOnce {
		if len(gc.Build.FixedPorts) > 0 {
			ports = make(nat.PortMap)
			for k, v := range gc.Build.FixedPorts {
				fmt.Printf("%s => %v\n", k, v)
				var bind []nat.PortBinding
				bind = append(bind, nat.PortBinding {
					HostIP: "0.0.0.0",
					HostPort: v,
				})
				ports[(nat.Port)(v + "/tcp")] = bind
			}
		} else {
			ports = nil
		}
	}

	return ports
}

func (gc *GearConfig) GetBuildRun() string {
	if gc == nil {
		return ""
	}
	return gc.Build.Run
}

func (gc *GearConfig) GetBuildArgs() string {
	if gc == nil {
		return ""
	}
	return gc.Build.Args.String()
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

func (gc *GearConfig) GetVersion(version string) *GearVersion {
	return gc.Versions.GetVersion(version)
}

func (gc *GearConfig) GetVersions() *GearVersions {
	return gc.Versions.GetVersions()
}

func (gc *GearConfig) IsBaseRef(version string) bool {
	var ok bool
	if state := gc.IsNil(); state.IsError() {
		return ok
	}

	for range onlyOnce {
		vers := gc.Versions.GetVersion(version)
		ok = vers.IsBaseRef()
	}

	return ok
}
