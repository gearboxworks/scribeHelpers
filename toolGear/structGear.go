package toolGear

import (
	"github.com/docker/go-connections/nat"
	"github.com/jedib0t/go-pretty/table"
	"github.com/newclarity/scribeHelpers/toolGear/gearConfig"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"strconv"
)


type Gear struct {
	Repo         *GitHubRepo
	GearConfig   *gearConfig.GearConfig

	Image        *Image
	Container    *Container

	Docker		 *Docker

	Runtime      *toolRuntime.TypeRuntime
	State        *ux.State
}


func NewGear(runtime *toolRuntime.TypeRuntime, docker *Docker) *Gear {
	var gear Gear

	for range onlyOnce {
		runtime = runtime.EnsureNotNil()

		gear = Gear {
			Repo:       NewRepo(runtime),
			GearConfig: gearConfig.New(runtime),
			Image:      NewImage(runtime),
			Container:  NewContainer(runtime),

			Docker:     docker,

			Runtime:    runtime,
			State:      ux.NewState(runtime.CmdName, runtime.Debug),
		}
		gear.State.SetPackage("")
		gear.State.SetFunctionCaller()

		gear.Image.Docker = gear.Docker
		gear.Container.Docker = gear.Docker

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

func (gear *Gear) Refresh() *ux.State {
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

	for range onlyOnce {
		gear.State = gear.Image.Refresh()
		if gear.State.IsError() {
			break
		}

		//if gear.Container.IsNotValid() {
		//	gear.Container.Name = gear.Image.Name
		//	gear.Container.Version = gear.Image.Version
		//}

		gear.State = gear.Container.Refresh()
		if gear.State.IsError() {
			break
		}
	}
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

func (gear *Gear) Logs() *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}
	return gear.Container.Logs()
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

func (gear *Gear) GetVolumeMounts() []string {
	if state := gear.IsNil(); state.IsError() {
		return []string{}
	}
	return gear.Container.GetVolumeMounts()
}

func (gear *Gear) GetFixedPorts() nat.PortMap {
	if state := gear.IsNil(); state.IsError() {
		return nil
	}
	return gear.GearConfig.GetFixedPorts()
}

func (gear *Gear) AddVolume(local string, remote string) bool {
	if gear.Container.VolumeMounts == nil {
		gear.Container.VolumeMounts = make(VolumeMounts)
	}
	return gear.Container.VolumeMounts.Add(local, remote)
}


// ******************************************************************************** //

func (gear *Gear) ParseGearConfig(cs string) *ux.State {
	if state := ux.IfNilReturnError(gear); state.IsError() {
		return state
	}

	for range onlyOnce {
		gear.State = gear.GearConfig.ParseJson(cs)
		if gear.State.IsNotOk() {
			break
		}
	}

	return gear.State
}


// ******************************************************************************** //

func (gear *Gear) ListContainerPorts() *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//var err error
		if gear.IsNotRunning() {
			break
		}

		ux.PrintfCyan("Open ports for Container: %s-%s\n", gear.Container.Name, gear.Container.Version)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Container",
			"Port Name",
			"Host Port",
			"Container Port",
		})

		ports, _ := gear.GetPorts()
		for _, v := range ports {
			if v.PrivatePort == 22 {
				t.AppendRow([]interface{} {
					ux.SprintfYellow("%s-%s\n", gear.Container.Name, gear.Container.Version),
					ux.SprintfYellow("ssh"),
					ux.SprintfYellow("%s:%d", v.IP, v.PublicPort),
					ux.SprintfYellow("%d", v.PrivatePort),
				})
				continue
			}

			t.AppendRow([]interface{} {
				ux.SprintfGreen("%s-%s\n", gear.Container.Name, gear.Container.Version),
				ux.SprintfGreen(v.Name),
				ux.SprintfGreen("%s:%d", v.IP, v.PublicPort),
				ux.SprintfGreen("%d", v.PrivatePort),
			})
		}

		count := t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		t.Render()
		ux.PrintflnGreen("Ports found: %d", count)
		ux.PrintflnBlue("")

		gear.State.SetOk("")
	}

	return gear.State
}

func (gear *Gear) GetPorts() (Ports, *ux.State) {
	ports := make(Ports)
	if state := gear.IsNil(); state.IsError() {
		return ports, state
	}

	for range onlyOnce {
		ports = gear.Container.GetPorts()

		//gcp := gear.gearConfig.Build.Ports
		for _, p := range ports {
			if p.PrivatePort == 22 {
				p.Name = "ssh"
				continue
			}
			if p.PrivatePort == 9970 {
				p.Name = "gearbox"
				continue
			}

			for k, v := range gear.GearConfig.Build.Ports {
				if k == "" {
					continue
				}
				i, _ := strconv.Atoi(v)
				if uint16(i) == p.PrivatePort {
					p.Name = k
					break
				}
			}
		}
	}

	return ports, gear.State
}
