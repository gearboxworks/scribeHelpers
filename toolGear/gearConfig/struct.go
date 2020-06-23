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


func (gc *GearConfig) IsNil() *ux.State {
	if state := ux.IfNilReturnError(gc); state.IsError() {
		return state
	}
	gc.State = gc.State.EnsureNotNil()
	return gc.State
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
