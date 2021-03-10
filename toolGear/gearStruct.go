package toolGear

import (
	"github.com/newclarity/scribeHelpers/toolGear/gearConfig"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)
//"github.com/docker/docker/integration-cli/cli"
// DOCKER_HOST=tcp://macpro:2375


type Gear struct {
	Repo         *GitHubRepo
	GearConfig   *gearConfig.GearConfig

	Image        *Image
	Container    *Container

	Docker		 *Docker

	Runtime      *toolRuntime.TypeRuntime
	State        *ux.State
}
type Gears struct {
	Language    Language
	Array		map[string]*Gear
	Selected    *Gear

	Docker      *Docker

	Runtime     *toolRuntime.TypeRuntime
	State       *ux.State
}
type Language struct {
	AppName string
	ImageName string
	ContainerName string
}


func NewGear(runtime *toolRuntime.TypeRuntime) *Gear {
	var gear Gear

	for range onlyOnce {
		runtime = runtime.EnsureNotNil()

		gear = Gear {
			Repo:       NewRepo(runtime),
			GearConfig: gearConfig.New(runtime),
			Image:      NewImage(runtime),
			Container:  NewContainer(runtime),

			Docker:     nil,

			Runtime:    runtime,
			State:      ux.NewState(runtime.CmdName, runtime.Debug),
		}
		gear.State.SetPackage("")
		gear.State.SetFunctionCaller()

		//if gear.Repo.State.IsNotOk() {
		//	gear.State = gear.Repo.State
		//	break
		//}
		//
		//if gear.Docker.State.IsNotOk() {
		//	gear.State.SetError("can not connect to Docker service provider - maybe you haven't set DOCKER_HOST, or Docker not running on this host")
		//	//gear.State = gear.Docker.State
		//	break
		//}
		//
		//if gear.GearConfig.State.IsNotOk() {
		//	gear.State = gear.GearConfig.State
		//	break
		//}
	}

	return &gear
}


func NewGears(runtime *toolRuntime.TypeRuntime) Gears {
	var gears Gears

	for range onlyOnce {
		runtime = runtime.EnsureNotNil()

		l := Language {
			AppName:       "Gearbox",
			ImageName:     "Gear Image",
			ContainerName: "Gear",
		}

		gears = Gears {
			Language:   l,
			Array:      make(map[string]*Gear),
			Selected:   nil,

			Docker:     NewDocker(runtime),

			Runtime:    runtime,
			State:      ux.NewState(runtime.CmdName, runtime.Debug),
		}
		gears.State.SetPackage("")
		gears.State.SetFunctionCaller()

		gears.State = gears.Get()
		if gears.State.IsNotOk() {
			break
		}
	}

	return gears
}

func (gears *Gears) IsValid() *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.State.EnsureNotNil()

		if gears.Docker.Client == nil {
			gears.State.SetError("docker client is nil")
			break
		}
	}

	return gears.State
}


func (gears *Gears) IsNil() *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.State.EnsureNotNil()
	}

	return gears.State
}


func (gears *Gears) SetLanguage(appName string, imageName string, containerName string) *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.Language = Language {
			AppName:       appName,
			ImageName:     imageName,
			ContainerName: containerName,
		}
	}

	return gears.State
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


func (gear *Gear) IsRunning() bool {
	var ok bool
	if state := ux.IfNilReturnError(gear); state.IsError() {
		return ok
	}

	for range onlyOnce {
		switch gear.Container.Summary.State {
			case ux.StateUnknown:
				//
			case ux.StateRunning:
				ok = true
			case ux.StatePaused:
				//
			case ux.StateCreated:
				//
			case ux.StateRestarting:
				//
			case ux.StateRemoving:
				//
			case ux.StateExited:
				//
			case ux.StateDead:
				//
		}
	}

	return ok
}
func (gear *Gear) IsNotRunning() bool {
	return !gear.IsRunning()
}

func (gear *Gear) IsCreated() bool {
	var ok bool
	if state := ux.IfNilReturnError(gear); state.IsError() {
		return ok
	}

	for range onlyOnce {
		switch gear.Container.Summary.State {
			case ux.StateUnknown:
				//
			case ux.StateRunning:
				ok = true
			case ux.StatePaused:
				ok = true
			case ux.StateCreated:
				ok = true
			case ux.StateRestarting:
				ok = true
			case ux.StateRemoving:
				//
			case ux.StateExited:
				ok = true
			case ux.StateDead:
				ok = true
		}
	}

	return ok
}
func (gear *Gear) IsNotCreated() bool {
	return !gear.IsCreated()
}

func (gear *Gear) IsStopped() bool {
	var ok bool
	if state := ux.IfNilReturnError(gear); state.IsError() {
		return ok
	}

	for range onlyOnce {
		switch gear.Container.Summary.State {
			case ux.StateUnknown:
				//
			case ux.StateRunning:
				//
			case ux.StatePaused:
				//
			case ux.StateCreated:
				//
			case ux.StateRestarting:
				//
			case ux.StateRemoving:
				//
			case ux.StateExited:
				ok = true
			case ux.StateDead:
				//
		}
	}

	return ok
}
func (gear *Gear) IsNotStopped() bool {
	return !gear.IsStopped()
}

//func (gear *Gear) IsNotRunning() bool {
//	var ok bool
//	if state := ux.IfNilReturnError(gear); state.IsError() {
//		return ok
//	}
//
//	for range onlyOnce {
//		switch gear.Container.Summary.State {
//			case ux.StateUnknown:
//				//
//			case ux.StateRunning:
//				//
//			case ux.StatePaused:
//				//
//			case ux.StateCreated:
//				//
//			case ux.StateRestarting:
//				//
//			case ux.StateRemoving:
//				//
//			case ux.StateExited:
//				//
//			case ux.StateDead:
//				//
//		}
//	}
//
//	return ok
//}
//func (gear *Gear) IsNotRunning() bool {
//	return !gear.IsRunning()
//}


func (gear *Gear) Status() *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	//for range onlyOnce {
	//	//var found bool
	//
	//	//found, gear.State = gear.FindContainer(gear.GearConfig.Meta.Name)
	//	//found, gear.State = gear.FindImage()
	//	gear.State = gear.Image.Status()
	//
	//	//gear.Image.Name = name
	//	//gear.Image.Version = version
	//	//gear.Image.Status()
	//	//
	//	//gear.Container.Name = name
	//	//gear.Container.Version = version
	//	//gear.Container.Status()
	//
	//	gear.State = gear.Container.Status()
	//	//gear.State = gear.Docker.Status()
	//	if gear.State.IsError() {
	//		break
	//	}
	//
	//	if gear.Image.GearConfig != nil {
	//		gear.GearConfig = gear.Image.GearConfig
	//	}
	//	if gear.Container.GearConfig != nil {
	//		gear.GearConfig = gear.Container.GearConfig
	//	}
	//
	//	if gear.Image.ID == "" {
	//		gear.Image.ID = strings.TrimPrefix(gear.Container.Details.Image, "sha256:")
	//		gear.Image.Name = gear.Container.Name
	//		gear.Image.Version = gear.Container.Version
	//	}
	//
	//	state2 := gear.Image.Status()
	//	if state2.IsError() {
	//		break
	//	}
	//
	//	//state = runState
	//
	//	//state = gear.Docker.Image.State()
	//	//if state.IsError() {
	//	//	break
	//	//}
	//}

	gear.State = gear.Container.Status()

	return gear.State
}


func (gear *Gear) Start() *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}
	return gear.Container.Start()
}

func (gear *Gear) Stop() *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}
	return gear.Container.Stop()
}

func (gear *Gear) Remove() *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}
	return gear.Container.Remove()
}

func (gear *Gear) GetCommand(cmd []string) []string {
	return gear.GearConfig.GetCommand(cmd)
}

//func (gear *Gear) Create() *ux.State {
//	if state := gear.IsNil(); state.IsError() {
//		return state
//	}
//	return gear.Container.Create()
//}


func (gears *Gears) Get() *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.GetImages("")
		if gears.State.IsNotOk() {
			break
		}

		gears.State = gears.GetContainers("")
		if gears.State.IsNotOk() {
			break
		}
	}

	return gears.State
}


func (gears *Gears) Status() *ux.State {
	state := ux.EnsureStateNotNil(nil)

	for _, v := range gears.Array {
		state = v.Status()
		if state.IsNotOk() {
			break
		}
	}

	return state
}


//func (gear *Gears) FindContainer(gearName string, gearVersion string) (bool, *ux.State) {
//	var found bool
//	if state := gear.IsNil(); state.IsError() {
//		return false, state
//	}
//
//	for range onlyOnce {
//		found, gear.State = gear.FindContainer(gearName, gearVersion)
//		if !found {
//			break
//		}
//		if gear.State.IsError() {
//			break
//		}
//
//		gear.State = gear.Status()
//		if gear.State.IsError() {
//			break
//		}
//
//		gear.Selected.GearConfig = gear.Selected.Container.GearConfig
//	}
//
//	return found, gear.State
//}


//func (gear *Gears) FindImage(gearName string, gearVersion string) (bool, *ux.State) {
//	var found bool
//	if state := gear.IsNil(); state.IsError() {
//		return false, state
//	}
//
//	for range onlyOnce {
//		found, gear.State = gear.FindImage(gearName, gearVersion)
//		if !found {
//			//state.ClearError()
//			break
//		}
//		if gear.State.IsError() {
//			break
//		}
//
//		//if gear.GearConfig == nil {
//		//	gear.GearConfig = gear.Docker.Image.GearConfig
//		//}
//
//		//@TODO - TO CHECK
//		//state = gear.Status()
//		//if state.IsError() {
//		//	break
//		//}
//	}
//
//	return found, gear.State
//}


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
			case gear.IsErrConnectionFailed(err):
			case gear.IsErrNotFound(err):
			case gear.IsErrPluginPermissionDenied(err):
			case gear.IsErrUnauthorized(err):
			default:
		}
	}

	return ok, gear.State
}


func (gear *Gear) ListLinks(version string) *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}
	return gear.GearConfig.ListLinks(gear.Container.Version)
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


//func (gear *Gear) AddVolume(local string, remote string) bool {
//	return gear.AddVolume(local, remote)
//}


//func (gear *Gear) AddMount(local string, remote string) bool {
//	return gear.AddMount(local, remote)
//}


//func (gear *Gear) ContainerCreate(gearName string, gearVersion string) *ux.State {
//	return gear.Docker.ContainerCreate(gearName, gearVersion)
//}


//func (gear *Gear) GetContainers(gearName string) (Gears, *ux.State) {
//	return gear.GetContainers(gearName)
//}
