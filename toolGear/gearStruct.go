package toolGear

import (
	"github.com/newclarity/scribeHelpers/toolGear/gearConfig"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
)
//"github.com/docker/docker/integration-cli/cli"
// DOCKER_HOST=tcp://macpro:2375


type Gear struct {
	Repo         *GitHubRepo
	Docker       *TypeDockerGear
	GearConfig   *gearConfig.GearConfig

	Runtime      *toolRuntime.TypeRuntime
	State        *ux.State
}


func NewGear(runtime *toolRuntime.TypeRuntime) *Gear {
	var gear Gear

	for range onlyOnce {
		runtime = runtime.EnsureNotNil()

		gear = Gear{
			Repo:       NewRepo(runtime),
			Docker:     New(runtime),
			GearConfig: gearConfig.New(runtime),
			Runtime:    runtime,
			State:      ux.NewState(runtime.CmdName, runtime.Debug),
		}
		gear.State.SetPackage("")
		gear.State.SetFunctionCaller()

		if gear.Repo.State.IsNotOk() {
			gear.State = gear.Repo.State
			break
		}

		if gear.Docker.State.IsNotOk() {
			gear.State.SetError("can not connect to Docker service provider")
			//gear.State = gear.Docker.State
			break
		}

		if gear.GearConfig.State.IsNotOk() {
			gear.State = gear.GearConfig.State
			break
		}
	}

	return &gear
}


func NewGearConfig(runtime *toolRuntime.TypeRuntime) *gearConfig.GearConfig {
	return gearConfig.New(runtime)
}


func (gear *Gear) IsNil() *ux.State {
	if state := ux.IfNilReturnError(gear); state.IsError() {
		return state
	}

	for range onlyOnce {
		gear.State = gear.State.EnsureNotNil()

		gear.State = gear.Docker.IsNil()
		if gear.State.IsNotOk() {
			break
		}

		gear.State = gear.Repo.IsNil()
		if gear.State.IsNotOk() {
			break
		}
	}

	return gear.State
}


func (gear *Gear) IsValid() *ux.State {
	if state := ux.IfNilReturnError(gear); state.IsError() {
		return state
	}

	for range onlyOnce {
		gear.State = gear.State.EnsureNotNil()
	}

	return gear.State
}


func (gear *Gear) Status() *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gear.State = gear.Docker.Container.Status()
		if gear.State.IsError() {
			break
		}

		if gear.Docker.Image.GearConfig != nil {
			gear.GearConfig = gear.Docker.Image.GearConfig
		}
		if gear.Docker.Container.GearConfig != nil {
			gear.GearConfig = gear.Docker.Container.GearConfig
		}

		if gear.Docker.Image.ID == "" {
			gear.Docker.Image.ID = strings.TrimPrefix(gear.Docker.Container.Details.Image, "sha256:")
			gear.Docker.Image.Name = gear.Docker.Container.Name
			gear.Docker.Image.Version = gear.Docker.Container.Version
		}

		state2 := gear.Docker.Image.Status()
		if state2.IsError() {
			break
		}

		//state = runState

		//state = gear.Docker.Image.State()
		//if state.IsError() {
		//	break
		//}
	}

	return gear.State
}


func (gear *Gear) FindContainer(gearName string, gearVersion string) (bool, *ux.State) {
	var found bool
	if state := gear.IsNil(); state.IsError() {
		return false, state
	}

	for range onlyOnce {
		found, gear.State = gear.Docker.FindContainer(gearName, gearVersion)
		if !found {
			break
		}
		if gear.State.IsError() {
			break
		}

		gear.State = gear.Status()
		if gear.State.IsError() {
			break
		}

		gear.GearConfig = gear.Docker.Container.GearConfig
	}

	return found, gear.State
}


func (gear *Gear) FindImage(gearName string, gearVersion string) (bool, *ux.State) {
	var found bool
	if state := gear.IsNil(); state.IsError() {
		return false, state
	}

	for range onlyOnce {
		found, gear.State = gear.Docker.FindImage(gearName, gearVersion)
		if !found {
			//state.ClearError()
			break
		}
		if gear.State.IsError() {
			break
		}

		//if gear.GearConfig == nil {
		//	gear.GearConfig = gear.Docker.Image.GearConfig
		//}

		//@TODO - TO CHECK
		//state = gear.Status()
		//if state.IsError() {
		//	break
		//}
	}

	return found, gear.State
}


func (gear *Gear) DecodeError(err error) (bool, *ux.State) {
	var ok bool
	if state := gear.IsNil(); state.IsError() {
		return false, state
	}

	for range onlyOnce {
		switch {
			case err != nil:
				ok = true

			//case gear.Docker.IsErrContainerNotFound(err):
			case gear.Docker.IsErrConnectionFailed(err):
			case gear.Docker.IsErrNotFound(err):
			case gear.Docker.IsErrPluginPermissionDenied(err):
			case gear.Docker.IsErrUnauthorized(err):
			default:
		}
	}

	return ok, gear.State
}


func (gear *Gear) ListLinks(version string) *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gear.State = gear.GearConfig.ListLinks(version)
	}

	return gear.State
}


func (gear *Gear) CreateLinks(version string) *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gear.State = gear.GearConfig.CreateLinks(version)
	}

	return gear.State
}


func (gear *Gear) RemoveLinks(version string) *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gear.State = gear.GearConfig.RemoveLinks(version)
	}

	return gear.State
}


func (gear *Gear) AddVolume(local string, remote string) bool {
	return gear.Docker.AddVolume(local, remote)
}


func (gear *Gear) AddMount(local string, remote string) bool {
	return gear.Docker.AddMount(local, remote)
}


func (gear *Gear) ContainerCreate(gearName string, gearVersion string) *ux.State {
	return gear.Docker.ContainerCreate(gearName, gearVersion)
}
