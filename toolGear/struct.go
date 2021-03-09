package toolGear

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
	"github.com/newclarity/scribeHelpers/toolGear/gearConfig"
	"github.com/newclarity/scribeHelpers/toolGear/gearSsh"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
)


type TypeGear struct {
	Image     *Image
	Container *Container
	gearConfig *gearConfig.GearConfig

	//Client    *client.Client
	Ssh       *gearSsh.Ssh

	//_Parent      *TypeDockerGear
	Runtime   *toolRuntime.TypeRuntime
	State     *ux.State
}
type GearConfigs []*gearConfig.GearConfig
//type TypeDockerGears []TypeDockerGear

//func (gear *TypeDockerGear) IsNil() *ux.State {
//	return ux.IfNilReturnError(gear)
//}


type TypeMatchImage struct {
	Organization string
	Name         string
	Version      string
}
type TypeMatchContainer TypeMatchImage


//func NewGear(runtime *toolRuntime.TypeRuntime) *Gear {
//	var gear Gear
//	runtime = runtime.EnsureNotNil()
//
//	for range onlyOnce {
//		gear = Gear {
//			Repo:       nil,
//			GearConfig: gearConfig.New(runtime),
//			Image:      NewImage(runtime),
//			Container:  NewContainer(runtime),
//			Ssh:        nil,
//
//			Docker:     nil,
//			Runtime:    runtime,
//			State:      ux.NewState(runtime.CmdName, runtime.Debug),
//		}
//
//		////foo := os.Getenv("DOCKER_HOST")
//		////fmt.Printf("DOCKER_HOST:%s\n", foo)
//		//
//		//var err error
//		//gear.Client, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
//		////cli.DockerClient, err = client.NewEnvClient()
//		//if err != nil {
//		//	gear.State.SetError("Docker client error: %s", err)
//		//	break
//		//}
//		//
//		//ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
//		////noinspection GoDeferInLoop
//		//defer cancel()
//		//
//		//_, err = gear.Client.Ping(ctx)
//		//if err != nil {
//		//	gear.State.SetError("Docker client error: %s", err)
//		//	break
//		//}
//
//		//gear.Image.Docker = &gear
//		//gear.Container.Docker = &gear
//	}
//
//	return &gear
//}


//func (gear *Gear) IsValid() *ux.State {
//	if state := ux.IfNilReturnError(gear); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		gear.State = gear.State.EnsureNotNil()
//
//		//if gear.Client == nil {
//		//	gear.State.SetError("docker client is nil")
//		//	break
//		//}
//	}
//
//	return gear.State
//}


//func (gear *Gear) Status(name string, version string) *ux.State {
//	if state := ux.IfNilReturnError(gear); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		gear.Image.Name = name
//		gear.Image.Version = version
//		gear.Image.Status()
//
//		gear.Container.Name = name
//		gear.Container.Version = version
//		gear.Container.Status()
//	}
//
//	return gear.State
//}


func (gear *Gear) SetSshStatusLine(s bool) {
	gear.Ssh.StatusLine.Enable = s
}


func (gear *Gear) SetSshShell(s bool) {
	gear.Ssh.Shell = s
}


func (gear *Gear) AddVolume(local string, remote string) bool {
	if gear.Container.VolumeMounts == nil {
		gear.Container.VolumeMounts = make(VolumeMounts)
	}
	return gear.Container.VolumeMounts.Add(local, remote)
}


//func (gear *Gear) ContainerCreate(gearName string, gearVersion string) *ux.State {
//	if state := ux.IfNilReturnError(gear); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		_, gear.State = gear.Docker.ContainerCreate(gearName, gearVersion)
//	}
//
//	return
//}


//func (gear *Gears) List(name string) *ux.State {
//	if state := ux.IfNilReturnError(gear); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		gear.State = gear.ImageList(name)
//		if gear.State.IsError() {
//			break
//		}
//
//		gear.State = gear.ContainerList(name)
//		if gear.State.IsError() {
//			break
//		}
//
//		gear.State = gear.NetworkList(DefaultNetwork)
//	}
//
//	return gear.State
//}


func (gears *Gears) ListContainers(name string) (*ux.State) {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.GetContainers("")
		if gears.State.IsNotOk() {
			break
		}

		ux.PrintfCyan("Gearbox containers: ")
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Name",
			"Class",
			"State",
			"Image",
			"Ports",
			"IP Address",
			"Mounts",
			"Size",
		})

		for _, gear := range gears.Array {
			if gear.Container == nil {
				continue
			}
			//if gear.Container.Summary == nil {
			//	continue
			//}
			if gear.Container.ID == "" {
				continue
			}
			name := gear.Container.Name

			sshPort := ""
			var ports string
			for _, p := range gear.Container.GetPorts() {
				//if p.PrivatePort == 22 {
				//	sshPort = fmt.Sprintf("%d", p.PublicPort)
				//	continue
				//}
				//ports += fmt.Sprintf("%s://%s:%d => %d\n", p.Type, p.IP, p.PublicPort, p.PrivatePort)
				if p.IP == "0.0.0.0" {
					ports += fmt.Sprintf("%d => %d\n", p.PublicPort, p.PrivatePort)
				} else {
					ports += fmt.Sprintf("%s://%s:%d => %d\n", p.Type, p.IP, p.PublicPort, p.PrivatePort)
				}
			}
			if sshPort == "0" {
				sshPort = "none"
			}

			var mounts string
			for _, m := range gear.Container.GetMounts() {
				// ms += fmt.Sprintf("%s(%s) host:%s => container:%s (RW:%v)\n", m.Name, m.Type, m.Source, m.Destination, m.RW)
				mounts += fmt.Sprintf("host:%s\n\t=> container:%s (RW:%v)\n", m.Source, m.Destination, m.RW)
			}

			var ipAddress string
			for k, n := range gear.Container.GetNetworks() {
				ipAddress += fmt.Sprintf("(%s) %s\n", k, n.IPAddress)
			}

			state := gear.Container.GetState()
			if state == ux.StateRunning {
				state = ux.SprintfGreen(state)
			} else {
				state = ux.SprintfYellow(state)
			}

			t.AppendRow([]interface{}{
				ux.SprintfWhite(name),
				ux.SprintfWhite(gear.GearConfig.GetClass()),
				state,
				ux.SprintfWhite(gear.Image.GetName()),
				ux.SprintfWhite(ports),
				ux.SprintfWhite(ipAddress),
				ux.SprintfWhite(mounts),
				ux.SprintfWhite(humanize.Bytes(gear.Container.GetSize())),
			})
		}

		gears.State.ClearError()
		count := t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		ux.PrintflnGreen("%d found", count)
		t.Render()
		gears.State.SetResponse(count)

		ux.PrintflnBlue("")
	}

	return gears.State
}


func (gears *Gears) ListImages(f string) (*ux.State) {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.GetImages("")
		if gears.State.IsNotOk() {
			break
		}

		ux.PrintfCyan("Downloaded Gearbox images: ")
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Class", "Image", "Ports", "Size"})

		for _, gear := range gears.Array {
			// foo := fmt.Sprintf("%s/%s", gc.Organization, gc.Name)
			t.AppendRow([]interface{}{
				ux.SprintfWhite(gear.GearConfig.GetClass()),
				//ux.SprintfWhite(gc.Meta.State),
				ux.SprintfWhite(gear.Image.GetName()),
				ux.SprintfWhite("%s", gear.GearConfig.GetPorts()),
				ux.SprintfWhite(humanize.Bytes(gear.Image.GetSize())),
			})
		}

		gears.State.ClearError()
		count := t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		ux.PrintflnGreen("%d found", count)
		t.Render()
		gears.State.SetResponse(count)

		ux.PrintflnBlue("")
	}

	return gears.State
}


func (gear *Gears) Ls(name string) *ux.State {
	if state := ux.IfNilReturnError(gear); state.IsError() {
		return state
	}

	for range onlyOnce {
		gear.State = gear.GetContainers(name)
		if gear.State.IsError() {
			break
		}

		gear.State = gear.NetworkList(DefaultNetwork)
	}

	return gear.State
}


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
