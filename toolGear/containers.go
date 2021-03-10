package toolGear

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
	"github.com/newclarity/scribeHelpers/toolGear/gearConfig"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"strings"
)


// Run a container
// This first example shows how to run a container using the Docker API.
// On the command line, you would use the docker run command, but this is just as easy to do from your own apps too.
// This is the equivalent of typing docker run alpine echo hello world at the command prompt:
func (gears *Gears) ContainerCreate(gearName string, gearVersion string) *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//if c.runtime.Debug {
		//	fmt.Printf("DEBUG: ContainerCreate(%s, %s)\n", gearName, gearVersion)
		//}

		if gearName == "" {
			gears.State.SetError("empty gearname")
			break
		}

		if gearVersion == "" {
			gearVersion = "latest"
		}

		var ok bool
		ok, gears.State = gears.FindContainer(gearName, gearVersion)
		if gears.State.IsError() {
			break
		}
		if !ok {
			// Find Gear image since we don't have a container.
			for range onlyOnce {
				ok, gears.State = gears.FindImage(gearName, gearVersion)
				if gears.State.IsError() {
					ok = false
					break
				}
				if ok {
					break
				}

				ux.PrintflnNormal("Downloading Gear '%s:%s'.", gearName, gearVersion)

				// Pull Gear image.
				gears.Selected.Image.ID = gearName
				gears.Selected.Image.Name = gearName
				gears.Selected.Image.Version = gearVersion
				gears.State = gears.Selected.Image.Pull()
				if gears.State.IsError() {
					gears.State.SetError("no such gear '%s'", gearName)
					break
				}

				// Confirm it's there.
				ok, gears.State = gears.FindImage(gearName, gearVersion)
				if gears.State.IsError() {
					ok = false
					break
				}
			}
			if !ok {
				gears.State.SetError("Cannot install Gear image '%s:%s' - %s.", gearName, gearVersion, gears.State.GetError())
				break
			}
			//c.State.Clear()
		}

		//c.Selected.Container.ID = c.Selected.Image.ID
		//c.Selected.Container.Name = c.Selected.Image.Name
		//c.Selected.Container.Version = c.Selected.Image.Version

		// c.Image.Details.Container = "gearboxworks/golang:1.14"
		// tag := fmt.Sprintf("", c.Image.Name, c.Image.Version)
		tag := fmt.Sprintf("gearboxworks/%s:%s", gearName, gearVersion)
		gn := fmt.Sprintf("%s-%s", gearName, gearVersion)

		var binds []string
		for k, v := range gears.Selected.Container.VolumeMounts {
			binds = append(binds, fmt.Sprintf("%s:%s", k, v))
		}

		config := container.Config {
			// Hostname:        "",
			// Domainname:      "",
			User:            "root",
			// AttachStdin:     false,
			AttachStdout:    true,
			AttachStderr:    true,
			ExposedPorts:    nil,
			Tty:             false,
			OpenStdin:       false,
			StdinOnce:       false,
			Env:             nil,
			Cmd:             []string{"/init"},
			// Healthcheck:     nil,
			// ArgsEscaped:     false,
			Image:           tag,
			// Volumes:         nil,
			// WorkingDir:      "",
			// Entrypoint:      nil,
			// NetworkDisabled: false,
			// MacAddress:      "",
			// OnBuild:         nil,
			// Labels:          nil,
			// StopSignal:      "",
			// StopTimeout:     nil,
			// Shell:           nil,
		}

		netConfig := network.NetworkingConfig {}

		// DockerMount
		// ms := mount.Mount {
		// 	Type:          "bind",
		// 	Source:        "/Users/mick/Documents/GitHub/containers/docker-golang",
		// 	Target:        "/foo",
		// 	ReadOnly:      false,
		// 	Consistency:   "",
		// 	BindOptions:   nil,
		// 	VolumeOptions: nil,
		// 	TmpfsOptions:  nil,
		// }

		var ports nat.PortMap
		if len(gears.Selected.GearConfig.Build.FixedPorts) > 0 {
			ports = make(nat.PortMap)
			for k, v := range gears.Selected.GearConfig.Build.FixedPorts {
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

		hostConfig := container.HostConfig {
			Binds:           binds,
			ContainerIDFile: "",
			LogConfig:       container.LogConfig{
				Type:   "",
				Config: nil,
			},
			NetworkMode:     DefaultNetwork,
			PortBindings:    ports,						// @TODO
			RestartPolicy:   container.RestartPolicy {
				Name:              "",
				MaximumRetryCount: 0,
			},
			AutoRemove:      false,
			VolumeDriver:    "",
			VolumesFrom:     nil,
			CapAdd:          nil,
			CapDrop:         nil,
			//Capabilities:    nil,
			//CgroupnsMode:    "",
			DNS:             []string{},
			DNSOptions:      []string{},
			DNSSearch:       []string{},
			ExtraHosts:      nil,
			GroupAdd:        nil,
			IpcMode:         "",
			Cgroup:          "",
			Links:           nil,
			OomScoreAdj:     0,
			PidMode:         "",
			Privileged:      true,
			PublishAllPorts: true,
			ReadonlyRootfs:  false,
			SecurityOpt:     nil,
			StorageOpt:      nil,
			Tmpfs:           nil,
			UTSMode:         "",
			UsernsMode:      "",
			ShmSize:         0,
			Sysctls:         nil,
			Runtime:         "runc",
			ConsoleSize:     [2]uint{},
			Isolation:       "",
			Resources:       container.Resources{},
			Mounts:          []mount.Mount{},
			//MaskedPaths:     nil,
			//ReadonlyPaths:   nil,
			Init:            nil,
		}

		//ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		////noinspection GoDeferInLoop
		//defer cancel()
		//
		//var resp container.ContainerCreateCreatedBody
		//var err error
		//resp, err = c.Docker.Client.ContainerCreate(ctx, &config, &hostConfig, &netConfig, gn)
		//if err != nil {
		//	c.State.SetError("error creating gear: %s", err)
		//	break
		//}
		//if resp.ID == "" {
		//	break
		//}

		var resp container.ContainerCreateCreatedBody
		resp, gears.State = gears.Docker.ContainerCreate(&config, &hostConfig, &netConfig, gn)
		if gears.State.IsNotOk() {
			break
		}

		gears.Selected.Container.ID = resp.ID
		gears.Selected.Container.Name = gearName
		gears.Selected.Container.Version = gearVersion
		gears.Selected.Container.Docker = gears.Docker
		if gearVersion == "latest" {
			gears.Selected.Container.IsLatest = true
		}

		// var response Response
		gears.State = gears.Selected.Status()
		if gears.State.IsError() {
			break
		}

		if gears.State.IsCreated() {
			break
		}

		//if c.State.IsRunning() {
		//	break
		//}
		//
		//if c.State.IsPaused() {
		//	break
		//}
		//
		//if c.State.IsRestarting() {
		//	break
		//}
	}

	return gears.State
}

//func (c *Gear) Create() *ux.State {
//	if state := c.IsNil(); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		//if c.runtime.Debug {
//		//	fmt.Printf("DEBUG: ContainerCreate(%s, %s)\n", gearName, gearVersion)
//		//}
//
//		if c.Container.Name == "" {
//			c.State.SetError("empty gearname")
//			break
//		}
//
//		if c.Container.Version == "" {
//			c.Container.Version = "latest"
//		}
//
//		var ok bool
//		ok, c.State = c.Container.FindContainer(gearName, c.Version)
//		if c.State.IsError() {
//			break
//		}
//		if !ok {
//			// Find Gear image since we don't have a container.
//			for range onlyOnce {
//				ok, c.State = c.FindImage(gearName, c.Version)
//				if c.State.IsError() {
//					ok = false
//					break
//				}
//				if ok {
//					break
//				}
//
//				ux.PrintflnNormal("Downloading Gear '%s:%s'.", gearName, c.Version)
//
//				// Pull Gear image.
//				c.Selected.Image.ID = gearName
//				c.Selected.Image.Name = gearName
//				c.Selected.Image.Version = gearVersion
//				c.State = c.Selected.Image.Pull()
//				if c.State.IsError() {
//					c.State.SetError("no such gear '%s'", gearName)
//					break
//				}
//
//				// Confirm it's there.
//				ok, c.State = c.FindImage(gearName, gearVersion)
//				if c.State.IsError() {
//					ok = false
//					break
//				}
//			}
//			if !ok {
//				c.State.SetError("Cannot install Gear image '%s:%s' - %s.", gearName, gearVersion, c.State.GetError())
//				break
//			}
//			//c.State.Clear()
//		}
//
//		c.Selected.Container.ID = c.Selected.Image.ID
//		c.Selected.Container.Name = c.Selected.Image.Name
//		c.Selected.Container.Version = c.Selected.Image.Version
//
//		// c.Image.Details.Container = "gearboxworks/golang:1.14"
//		// tag := fmt.Sprintf("", c.Image.Name, c.Image.Version)
//		tag := fmt.Sprintf("gearboxworks/%s:%s", c.Selected.Image.Name, c.Selected.Image.Version)
//		gn := fmt.Sprintf("%s-%s", c.Selected.Image.Name, c.Selected.Image.Version)
//
//		var binds []string
//		for k, v := range c.Selected.Container.VolumeMounts {
//			binds = append(binds, fmt.Sprintf("%s:%s", k, v))
//		}
//
//		config := container.Config {
//			// Hostname:        "",
//			// Domainname:      "",
//			User:            "root",
//			// AttachStdin:     false,
//			AttachStdout:    true,
//			AttachStderr:    true,
//			ExposedPorts:    nil,
//			Tty:             false,
//			OpenStdin:       false,
//			StdinOnce:       false,
//			Env:             nil,
//			Cmd:             []string{"/init"},
//			// Healthcheck:     nil,
//			// ArgsEscaped:     false,
//			Image:           tag,
//			// Volumes:         nil,
//			// WorkingDir:      "",
//			// Entrypoint:      nil,
//			// NetworkDisabled: false,
//			// MacAddress:      "",
//			// OnBuild:         nil,
//			// Labels:          nil,
//			// StopSignal:      "",
//			// StopTimeout:     nil,
//			// Shell:           nil,
//		}
//
//		netConfig := network.NetworkingConfig {}
//
//		// DockerMount
//		// ms := mount.Mount {
//		// 	Type:          "bind",
//		// 	Source:        "/Users/mick/Documents/GitHub/containers/docker-golang",
//		// 	Target:        "/foo",
//		// 	ReadOnly:      false,
//		// 	Consistency:   "",
//		// 	BindOptions:   nil,
//		// 	VolumeOptions: nil,
//		// 	TmpfsOptions:  nil,
//		// }
//
//		hostConfig := container.HostConfig {
//			Binds:           binds,
//			ContainerIDFile: "",
//			LogConfig:       container.LogConfig{
//				Type:   "",
//				Config: nil,
//			},
//			NetworkMode:     DefaultNetwork,
//			PortBindings:    nil,						// @TODO
//			RestartPolicy:   container.RestartPolicy {
//				Name:              "",
//				MaximumRetryCount: 0,
//			},
//			AutoRemove:      false,
//			VolumeDriver:    "",
//			VolumesFrom:     nil,
//			CapAdd:          nil,
//			CapDrop:         nil,
//			//Capabilities:    nil,
//			//CgroupnsMode:    "",
//			DNS:             []string{},
//			DNSOptions:      []string{},
//			DNSSearch:       []string{},
//			ExtraHosts:      nil,
//			GroupAdd:        nil,
//			IpcMode:         "",
//			Cgroup:          "",
//			Links:           nil,
//			OomScoreAdj:     0,
//			PidMode:         "",
//			Privileged:      true,
//			PublishAllPorts: true,
//			ReadonlyRootfs:  false,
//			SecurityOpt:     nil,
//			StorageOpt:      nil,
//			Tmpfs:           nil,
//			UTSMode:         "",
//			UsernsMode:      "",
//			ShmSize:         0,
//			Sysctls:         nil,
//			Runtime:         "runc",
//			ConsoleSize:     [2]uint{},
//			Isolation:       "",
//			Resources:       container.Resources{},
//			Mounts:          []mount.Mount{},
//			//MaskedPaths:     nil,
//			//ReadonlyPaths:   nil,
//			Init:            nil,
//		}
//
//		//ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
//		////noinspection GoDeferInLoop
//		//defer cancel()
//		//
//		//var resp container.ContainerCreateCreatedBody
//		//var err error
//		//resp, err = c.Docker.Client.ContainerCreate(ctx, &config, &hostConfig, &netConfig, gn)
//		//if err != nil {
//		//	c.State.SetError("error creating gear: %s", err)
//		//	break
//		//}
//		//if resp.ID == "" {
//		//	break
//		//}
//
//		var resp container.ContainerCreateCreatedBody
//		resp, c.State = c.Docker.ContainerCreate(&config, &hostConfig, &netConfig, gn)
//
//		c.Selected.Container.ID = resp.ID
//		//c.Container.Name = c.Image.Name
//		//c.Container.Version = c.Image.Version
//
//		// var response Response
//		c.State = c.Status()
//		if c.State.IsError() {
//			break
//		}
//
//		if c.State.IsCreated() {
//			break
//		}
//
//		//if c.State.IsRunning() {
//		//	break
//		//}
//		//
//		//if c.State.IsPaused() {
//		//	break
//		//}
//		//
//		//if c.State.IsRestarting() {
//		//	break
//		//}
//	}
//
//	return c.State
//}


func (gears *Gears) GetContainers(name string) (*ux.State) {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.Docker.ContainerList(true)
		if gears.State.IsNotOk() {
			break
		}

		var c types.Container
		for _, c = range gears.Docker.Containers {
			if _, ok := c.Labels["gearbox.json"]; !ok {
				continue
			}

			gear := NewGear(gears.Runtime)
			gear.Docker = gears.Docker
			gears.State = gear.GearConfig.ParseJson(c.Labels["gearbox.json"])
			if gears.State.IsError() {
				continue
			}

			if gear.GearConfig.Meta.Organization != DefaultOrganization {
				continue
			}

			if name != "" {
				if gear.GearConfig.Meta.Name != name {
					continue
				}
			}

			gear.Container.ID = c.ID
			gear.Container.Name = gear.GearConfig.GetName()
			gear.Container.Version = c.Labels["gearbox.version"]
			gear.Container.Summary = c
			gear.Container.Docker = gear.Docker
			gear.Container.GearConfig = gear.GearConfig
			gear.Container.Details, _ = gear.Docker.ContainerInspect(c.ID)
			//gear.State.RunState = c.State
			gear.State.RunState = c.State
			if c.Labels["container.latest"] == "true" {
				gear.Container.IsLatest = true
			}

			//gear.Image.ID = gear.Container.Summary.ImageID
			if _, ok := gears.Array[c.ImageID]; ok {
				gears.Array[c.ImageID].Container = gear.Container
			} else {
				gears.Array[c.ID] = gear
			}
		}

	}

	return gears.State
}

func (gears *Gears) ContainerListFiles(f string) (int, *ux.State) {
	var count int
	if state := gears.IsNil(); state.IsError() {
		return 0, state
	}

	for range onlyOnce {
		gears.State = gears.Docker.ContainerList(true)

		ux.PrintfCyan("Installed %s %s: ", gears.Language.AppName, gears.Language.ContainerName)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Name",
			"Class",
			"State",
			"Image",
			"Ports",
			"SSH port",
			"IP Address",
			"Mounts",
			"Size",
		})


		//gc := toolGear.NewGearConfig(gear.Runtime)
		gc := gearConfig.New(gears.Runtime)
		for _, c := range gears.Docker.Containers {
			gears.State = gc.ParseJson(c.Labels["gearbox.json"])
			if gears.State.IsError() {
				continue
			}

			if gc.Meta.Organization != DefaultOrganization {
				continue
			}

			if f != "" {
				if gc.Meta.Name != f {
					continue
				}
			}

			name := strings.TrimPrefix(c.Names[0], "/")

			sshPort := ""
			var ports string
			for _, p := range c.Ports {
				if p.PrivatePort == 22 {
					sshPort = fmt.Sprintf("%d", p.PublicPort)
					continue
				}
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
			for _, m := range c.Mounts {
				// ms += fmt.Sprintf("%s(%s) host:%s => container:%s (RW:%v)\n", m.Name, m.Type, m.Source, m.Destination, m.RW)
				mounts += fmt.Sprintf("host:%s\n\t=> %s:%s (RW:%v)\n",
					m.Source,
					gears.Language.ContainerName,
					m.Destination,
					m.RW,
					)
			}

			var ipAddress string
			for k, n := range c.NetworkSettings.Networks {
				ipAddress += fmt.Sprintf("(%s) %s\n", k, n.IPAddress)
			}

			var state string
			if c.State == ux.StateRunning {
				state = ux.SprintfGreen(c.State)
			} else {
				state = ux.SprintfYellow(c.State)
			}

			t.AppendRow([]interface{}{
				ux.SprintfWhite(name),
				ux.SprintfWhite(gc.Meta.Class),
				state,
				ux.SprintfWhite(c.Image),
				ux.SprintfWhite(ports),
				ux.SprintfWhite(sshPort),
				ux.SprintfWhite(ipAddress),
				ux.SprintfWhite(mounts),
				ux.SprintfWhite(humanize.Bytes(uint64(c.SizeRootFs))),
			})
		}

		gears.State.ClearError()
		count = t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		ux.PrintflnGreen("%d found", count)
		t.Render()
		ux.PrintflnBlue("")
	}

	return count, gears.State
}

func (gears *Gears) PrintContainers(f string) (int, *ux.State) {
	var count int
	if state := gears.IsNil(); state.IsError() {
		return 0, state
	}

	for range onlyOnce {
		ux.PrintfCyan("Installed %s %s: ", gears.Language.AppName, gears.Language.ContainerName)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Name",
			"Class",
			"State",
			"Image",
			"Ports",
			"SSH port",
			"IP Address",
			"Mounts",
			"Size",
		})

		gears.State = gears.Docker.ContainerList(true)
		//gc := toolGear.NewGearConfig(gear.Runtime)
		gc := gearConfig.New(gears.Runtime)
		for _, c := range gears.Array {
			gears.State = gc.ParseJson(c.Container.Summary.Labels["gearbox.json"])
			if gears.State.IsError() {
				continue
			}

			if gc.Meta.Organization != DefaultOrganization {
				continue
			}

			if f != "" {
				if gc.Meta.Name != f {
					continue
				}
			}

			name := strings.TrimPrefix(c.Container.Summary.Names[0], "/")

			sshPort := ""
			var ports string
			for _, p := range c.Container.Summary.Ports {
				if p.PrivatePort == 22 {
					sshPort = fmt.Sprintf("%d", p.PublicPort)
					continue
				}
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
			for _, m := range c.Container.Summary.Mounts {
				// ms += fmt.Sprintf("%s(%s) host:%s => container:%s (RW:%v)\n", m.Name, m.Type, m.Source, m.Destination, m.RW)
				mounts += fmt.Sprintf("host:%s\n\t=> %s:%s (RW:%v)\n",
					m.Source,
					gears.Language.ContainerName,
					m.Destination,
					m.RW,
					)
			}

			var ipAddress string
			for k, n := range c.Container.Summary.NetworkSettings.Networks {
				ipAddress += fmt.Sprintf("(%s) %s\n", k, n.IPAddress)
			}

			var state string
			if c.Container.Details.State.Status == ux.StateRunning {
				state = ux.SprintfGreen(c.Container.Details.State.Status)
			} else {
				state = ux.SprintfYellow(c.Container.Details.State.Status)
			}

			t.AppendRow([]interface{}{
				ux.SprintfWhite(name),
				ux.SprintfWhite(gc.Meta.Class),
				state,
				ux.SprintfWhite(c.Image.Name),
				ux.SprintfWhite(ports),
				ux.SprintfWhite(sshPort),
				ux.SprintfWhite(ipAddress),
				ux.SprintfWhite(mounts),
				ux.SprintfWhite(humanize.Bytes(uint64(*c.Container.Details.SizeRootFs))),
			})
		}

		gears.State.ClearError()
		count = t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		ux.PrintflnGreen("%d found", count)
		t.Render()
		ux.PrintflnBlue("")
	}

	return count, gears.State
}

func (gears *Gears) ContainerSprintf(f string) string {
	var ret string
	if state := gears.IsNil(); state.IsError() {
		ret = ux.SprintfRed("No %s %s found.\n", gears.Language.AppName, gears.Language.ContainerName)
		return ret
	}

	for range onlyOnce {
		gears.State = gears.Docker.ContainerList(true)
		if gears.State.IsNotOk() {
			break
		}

		ret = ux.SprintfCyan("Installed %s %s:\n", gears.Language.AppName, gears.Language.ContainerName)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Name",
			"Class",
			"State",
			"Image",
			"Ports",
			"SSH port",
			"IP Address",
			"Mounts",
			"Size",
		})

		//gc := toolGear.NewGearConfig(gear.Runtime)
		gc := gearConfig.New(gears.Runtime)
		for _, c := range gears.Docker.Containers {
			//c.State = gc.ParseJson(c.Summary.Labels["gearbox.json"])
			//if c.State.IsError() {
			//	break
			//}
			gears.State = gc.ParseJson(c.Labels["gearbox.json"])
			if gears.State.IsError() {
				continue
			}

			if gc.Meta.Organization != DefaultOrganization {
				continue
			}

			if f != "" {
				if gc.Meta.Name != f {
					continue
				}
			}

			name := strings.TrimPrefix(c.Names[0], "/")

			sshPort := ""
			var ports string
			for _, p := range c.Ports {
				if p.PrivatePort == 22 {
					sshPort = fmt.Sprintf("%d", p.PublicPort)
					continue
				}
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
			for _, m := range c.Mounts {
				// ms += fmt.Sprintf("%s(%s) host:%s => container:%s (RW:%v)\n", m.Name, m.Type, m.Source, m.Destination, m.RW)
				mounts += fmt.Sprintf("host:%s\n\t=> %s:%s (RW:%v)\n",
					m.Source,
					gears.Language.ContainerName,
					m.Destination,
					m.RW,
					)
			}

			var ipAddress string
			for k, n := range c.NetworkSettings.Networks {
				ipAddress += fmt.Sprintf("(%s) %s\n", k, n.IPAddress)
			}

			var state string
			if c.State == ux.StateRunning {
				state = ux.SprintfGreen(c.State)
			} else {
				state = ux.SprintfYellow(c.State)
			}

			t.AppendRow([]interface{}{
				ux.SprintfWhite(name),
				ux.SprintfWhite(gc.Meta.Class),
				state,
				ux.SprintfWhite(c.Image),
				ux.SprintfWhite(ports),
				ux.SprintfWhite(sshPort),
				ux.SprintfWhite(ipAddress),
				ux.SprintfWhite(mounts),
				ux.SprintfWhite(humanize.Bytes(uint64(c.SizeRootFs))),
			})
		}

		gears.State.ClearError()
		count := t.Length()
		if count == 0 {
			ret += ux.SprintfYellow("No %s %s found.\n", gears.Language.AppName, gears.Language.ContainerName)
			break
		}
		ret += ux.SprintfGreen("Found %d %s %ss.\n", count, gears.Language.AppName, gears.Language.ContainerName)

		ret += t.Render()
		//ux.PrintflnBlue("")
	}

	return ret
}

func (gears *Gears) FindContainer(gearName string, gearVersion string) (bool, *ux.State) {
	var ok bool
	if state := gears.IsNil(); state.IsError() {
		return false, state
	}

	for range onlyOnce {
		if gearName == "" {
			gears.State.SetError("empty gearname")
			break
		}

		if gearVersion == "" {
			gearVersion = "latest"
		}

		for _, c := range gears.Array {
			ok, _ = MatchContainer(&c.Container.Summary,
				TypeMatchContainer{Organization: DefaultOrganization, Name: gearName, Version: gearVersion})
			if ok {
				gears.Selected = c
				gears.State.RunState = c.Container.Summary.State
				break
			}
		}

		if !ok {
			gears.State.SetWarning("Container '%s-%s' doesn't exist.", gearName, gearVersion)
			break
		}

		gears.State.SetOk("found %s", gears.Language.ContainerName)
	}

	return ok, gears.State
}

//func (gears *Gears) FindContainer(gearName string, gearVersion string) (bool, *ux.State) {
//	var ok bool
//	if state := gears.IsNil(); state.IsError() {
//		return false, state
//	}
//
//	for range onlyOnce {
//		if gearName == "" {
//			gears.State.SetError("empty gearname")
//			break
//		}
//
//		if gearVersion == "" {
//			gearVersion = "latest"
//		}
//
//		gears.State = gears.Docker.ContainerList(false)
//		if gears.State.IsNotOk() {
//			break
//		}
//
//		// Start out with "not found". Will be cleared if found or error occurs.
//		gears.State.SetWarning("Gear '%s:%s' doesn't exist.", gearName, gearVersion)
//		for _, c := range gears.Docker.Containers {
//			var gc *gearConfig.GearConfig
//			ok, gc = MatchContainer(&c,
//				TypeMatchContainer{Organization: DefaultOrganization, Name: gearName, Version: gearVersion})
//			if !ok {
//				continue
//			}
//
//			gears.Selected.Container.Name = gearName
//			//gears.Selected.Container.Name = gc.Meta.Name
//			gears.Selected.Container.Version = gearVersion
//			gears.Selected.Container.GearConfig = gc
//			gears.Selected.Container.Summary = c
//			gears.Selected.Container.ID = c.ID
//			gears.Selected.Container.State = gears.Selected.Container.State.EnsureNotNil()
//			gears.State.SetOk("Found Gear '%s:%s'.", gearName, gearVersion)
//			ok = true
//
//			gears.Selected.Container.Details, gears.State = gears.Docker.ContainerInspect(c.ID)
//			break
//		}
//
//		if gears.State.IsNotOk() {
//			if !ok {
//				gears.State.ClearError()
//			}
//			break
//		}
//
//		//if gears.Selected.Container.Summary == nil {
//		//	break
//		//}
//
//		//ctx2, cancel2 := context.WithTimeout(context.Background(), DefaultTimeout)
//		////noinspection GoDeferInLoop
//		//defer cancel2()
//		//d := types.ContainerJSON{}
//		//d, err = gears.Client.ContainerInspect(ctx2, gears.Container.ID)
//		//if err != nil {
//		//	gears.State.SetError("gear inspect error: %s", err)
//		//	break
//		//}
//		//gears.Container.Details = &d
//	}
//
//	return ok, gears.State
//}

func MatchContainer(m *types.Container, match TypeMatchContainer) (bool, *gearConfig.GearConfig) {
	var ok bool
	gc := gearConfig.New(nil)

	for range onlyOnce {
		if MatchTag("<none>:<none>", m.Names) {
			ok = false
			break
		}

		gc.State = gc.ParseJson(m.Labels["gearbox.json"])
		if gc.State.IsError() {
			ok = false
			break
		}

		if gc.Meta.Organization != DefaultOrganization {
			ok = false
			break
		}

		tagCheck := fmt.Sprintf("%s/%s:%s", match.Organization, match.Name, match.Version)
		if m.Image == tagCheck {
			ok = true
			break
		}

		if gc.Meta.Name != match.Name {
			//if !RunAs.AsLink {
			if gc.Runtime.IsRunningAsFile() {
				ok = false
				break
			}

			cs := gc.MatchCommand(match.Name)
			if cs == nil {
				ok = false
				break
			}

			match.Name = gc.Meta.Name
		}

		if !gc.Versions.HasVersion(match.Version) {
			ok = false
			break
		}

		if match.Version == "latest" {
			gl := gc.Versions.GetLatest()
			if match.Version != "" {
				match.Version = gl
			}
		}

		for range onlyOnce {
			if m.Labels["gearbox.version"] == match.Version {
				ok = true
				break
			}

			if m.Labels["container.majorversion"] == match.Version {
				ok = true
				break
			}

			ok = false
		}
		break
	}

	return ok, gc
}
