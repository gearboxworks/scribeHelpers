package gearConfig

import (
	"encoding/json"
	"fmt"
	"github.com/newclarity/scribeHelpers/helperRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)

const OnlyOnce = "1"

const DefaultCommandName = "default"


type GearConfig struct {
	Meta       GearMeta       `json:"meta"`
	Build      GearBuild      `json:"build"`
	Run        GearRun        `json:"run"`
	Project    GearProject    `json:"project"`
	Extensions GearExtensions `json:"extensions"`
	Versions   GearVersions   `json:"versions"`

	Schema     string         `json:"schema"`

	Runtime    *helperRuntime.TypeRuntime
	State      *ux.State
}
type GearConfigs map[string]GearConfig


func New(runtime *helperRuntime.TypeRuntime) *GearConfig {
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

	for range OnlyOnce {
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
		fmt.Printf("HEY")
		return gc.State
	}
	for range OnlyOnce {
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

	for range OnlyOnce {
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
