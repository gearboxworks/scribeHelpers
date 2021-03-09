package toolGear

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
	"github.com/newclarity/scribeHelpers/toolGear/gearConfig"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"strings"
)


// List and manage containers
// You can use the API to list containers that are running, just like using docker ps:

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

func (gear *Gears) ContainerListFiles(f string) (int, *ux.State) {
	var count int
	if state := gear.IsNil(); state.IsError() {
		return 0, state
	}

	for range onlyOnce {
		gear.State = gear.Docker.ContainerList(true)

		ux.PrintfCyan("Installed Gearbox gears: ")
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
		gc := gearConfig.New(gear.Runtime)
		for _, c := range gear.Docker.Containers {
			gear.State = gc.ParseJson(c.Labels["gearbox.json"])
			if gear.State.IsError() {
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
				mounts += fmt.Sprintf("host:%s\n\t=> container:%s (RW:%v)\n", m.Source, m.Destination, m.RW)
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

		gear.State.ClearError()
		count = t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		ux.PrintflnGreen("%d found", count)
		t.Render()
		ux.PrintflnBlue("")
	}

	return count, gear.State
}


// func ContainerList(f types.ContainerListOptions) error {
func (gear *Gears) PrintContainers(f string) (int, *ux.State) {
	var count int
	if state := gear.IsNil(); state.IsError() {
		return 0, state
	}

	for range onlyOnce {
		ux.PrintfCyan("Installed Gearbox gears: ")
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

		gear.State = gear.Docker.ContainerList(true)
		//gc := toolGear.NewGearConfig(gear.Runtime)
		gc := gearConfig.New(gear.Runtime)
		for _, c := range gear.Array {
			gear.State = gc.ParseJson(c.Container.Summary.Labels["gearbox.json"])
			if gear.State.IsError() {
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
				mounts += fmt.Sprintf("host:%s\n\t=> container:%s (RW:%v)\n", m.Source, m.Destination, m.RW)
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

		gear.State.ClearError()
		count = t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		ux.PrintflnGreen("%d found", count)
		t.Render()
		ux.PrintflnBlue("")
	}

	return count, gear.State
}

func (gear *Gears) ContainerSprintf(f string) string {
	var ret string
	if state := gear.IsNil(); state.IsError() {
		ret = ux.SprintfRed("No Gearbox containers found.\n")
		return ret
	}

	for range onlyOnce {
		gear.State = gear.Docker.ContainerList(true)
		if gear.State.IsNotOk() {
			break
		}

		ret = ux.SprintfCyan("Installed Gearbox gears:\n")
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
		gc := gearConfig.New(gear.Runtime)
		for _, c := range gear.Docker.Containers {
			//c.State = gc.ParseJson(c.Summary.Labels["gearbox.json"])
			//if c.State.IsError() {
			//	break
			//}
			gear.State = gc.ParseJson(c.Labels["gearbox.json"])
			if gear.State.IsError() {
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
				mounts += fmt.Sprintf("host:%s\n\t=> container:%s (RW:%v)\n", m.Source, m.Destination, m.RW)
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

		gear.State.ClearError()
		count := t.Length()
		if count == 0 {
			ret += ux.SprintfYellow("No Gearbox containers found.\n")
			break
		}
		ret += ux.SprintfGreen("Found %d Gearbox containers.\n", count)

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

		gears.State = gears.Docker.ContainerList(false)

		// Start out with "not found". Will be cleared if found or error occurs.
		gears.State.SetWarning("Gear '%s:%s' doesn't exist.", gearName, gearVersion)
		for _, c := range gears.Docker.Containers {
			var gc *gearConfig.GearConfig
			ok, gc = MatchContainer(&c,
				TypeMatchContainer{Organization: DefaultOrganization, Name: gearName, Version: gearVersion})
			if !ok {
				continue
			}

			gears.Selected.Container.Name = gearName
			//gears.Selected.Container.Name = gc.Meta.Name
			gears.Selected.Container.Version = gearVersion
			gears.Selected.Container.GearConfig = gc
			gears.Selected.Container.Summary = c
			gears.Selected.Container.ID = c.ID
			gears.Selected.Container.State = gears.Selected.Container.State.EnsureNotNil()
			gears.State.SetOk("Found Gear '%s:%s'.", gearName, gearVersion)
			ok = true

			gears.Selected.Container.Details, gears.State = gears.Docker.ContainerInspect(c.ID)
			break
		}

		if gears.State.IsNotOk() {
			if !ok {
				gears.State.ClearError()
			}
			break
		}

		//if gears.Selected.Container.Summary == nil {
		//	break
		//}

		//ctx2, cancel2 := context.WithTimeout(context.Background(), DefaultTimeout)
		////noinspection GoDeferInLoop
		//defer cancel2()
		//d := types.ContainerJSON{}
		//d, err = gears.Client.ContainerInspect(ctx2, gears.Container.ID)
		//if err != nil {
		//	gears.State.SetError("gear inspect error: %s", err)
		//	break
		//}
		//gears.Container.Details = &d
	}

	return ok, gears.State
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
			for k, v := range gear.GearConfig.Build.Ports {
				if k == p.Name {
					fmt.Printf("HEY1")
				}
				if v == p.Name {
					fmt.Printf("HEY2")
				}
			}
		}
	}

	return ports, gear.State
}

func (gear *Gear) ListContainerPorts() *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//var err error

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


//func MatchContainer(m *types.Container, gearOrg string, gearName string, gearVersion string) (bool, *gearConfig.GearConfig) {
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
