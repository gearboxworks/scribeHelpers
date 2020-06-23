package toolGear

import (
	"context"
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
// func ContainerList(f types.ContainerListOptions) error {
func (gear *DockerGear) ContainerList(f string) (int, *ux.State) {
	var count int
	if state := gear.IsNil(); state.IsError() {
		return 0, state
	}

	for range onlyOnce {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		var containers []types.Container
		var err error
		containers, err = gear.Client.ContainerList(ctx, types.ContainerListOptions{Size: true, All: true})
		if err != nil {
			gear.State.SetError("gear list error: %s", err)
			break
		}

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
		for _, c := range containers {
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

func (gear *DockerGear) ContainerSprintf(f string) string {
	var ret string
	if state := gear.IsNil(); state.IsError() {
		ret = ux.SprintfRed("No Gearbox containers found.\n")
		return ret
	}

	for range onlyOnce {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		var containers []types.Container
		var err error
		containers, err = gear.Client.ContainerList(ctx, types.ContainerListOptions{Size: true, All: true})
		if err != nil {
			gear.State.SetError("gear list error: %s", err)
			ret = ux.SprintfRed("Provider error: %s\n", err)
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
		for _, c := range containers {
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

func (gear *DockerGear) FindContainer(gearName string, gearVersion string) (bool, *ux.State) {
	var ok bool
	if state := gear.IsNil(); state.IsError() {
		return false, state
	}

	for range onlyOnce {
		if gearName == "" {
			gear.State.SetError("empty gearname")
			break
		}

		if gearVersion == "" {
			gearVersion = "latest"
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		var containers []types.Container
		var err error
		containers, err = gear.Client.ContainerList(ctx, types.ContainerListOptions{All: true, Limit: 256})
		if err != nil {
			gear.State.SetError("gear list error: %s", err)
			break
		}

		// Start out with "not found". Will be cleared if found or error occurs.
		gear.State.SetWarning("Gear '%s:%s' doesn't exist.", gearName, gearVersion)

		for _, c := range containers {
			//var gc *gearConfig.GearConfig
			ok, gear.Container.GearConfig = MatchContainer(&c,
				TypeMatchContainer{Organization: DefaultOrganization, Name: gearName, Version: gearVersion})
			if !ok {
				continue
			}

			gear.Container.Name = gearName
			gear.Container.Version = gearVersion
			//gear.Container.GearConfig = gc
			gear.Container.Summary = &c
			gear.Container.ID = c.ID
			//gear.Container.Name = gc.Meta.Name
			gear.Container.State = gear.Container.State.EnsureNotNil()
			gear.State.SetOk("Found Gear '%s:%s'.", gearName, gearVersion)
			ok = true
			break
		}

		if gear.State.IsNotOk() {
			if !ok {
				gear.State.ClearError()
			}
			break
		}

		if gear.Container.Summary == nil {
			break
		}

		ctx2, cancel2 := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel2()
		d := types.ContainerJSON{}
		d, err = gear.Client.ContainerInspect(ctx2, gear.Container.ID)
		if err != nil {
			gear.State.SetError("gear inspect error: %s", err)
			break
		}
		gear.Container.Details = &d
	}

	return ok, gear.State
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
