package toolDocker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"strings"
)


// List and manage containers
// You can use the API to list containers that are running, just like using docker ps:
// func ContainerList(f types.ContainerListOptions) error {
func (d *Docker) ContainerList(f string) (int, *ux.State) {
	var count int
	if state := d.IsNil(); state.IsError() {
		return 0, state
	}

	for range onlyOnce {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		var containers []types.Container
		var err error
		containers, err = d.Client.ContainerList(ctx, types.ContainerListOptions{Size: true, All: true})
		if err != nil {
			d.State.SetError("d list error: %s", err)
			break
		}

		ux.PrintfCyan("Installed Docker containers: ")
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Name",
			"State",
			"Image",
			"Ports",
			"SSH port",
			"IP Address",
			"Mounts",
			"Size",
		})


		for _, c := range containers {
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
				state,
				ux.SprintfWhite(c.Image),
				ux.SprintfWhite(ports),
				ux.SprintfWhite(sshPort),
				ux.SprintfWhite(ipAddress),
				ux.SprintfWhite(mounts),
				ux.SprintfWhite(humanize.Bytes(uint64(c.SizeRootFs))),
			})
		}

		d.State.ClearError()
		count = t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		ux.PrintflnGreen("%d found", count)
		t.Render()
		ux.PrintflnBlue("")
	}

	return count, d.State
}


func (d *Docker) FindContainer(org string, name string, version string) (bool, *ux.State) {
	var ok bool
	if state := d.IsNil(); state.IsError() {
		return false, state
	}

	for range onlyOnce {
		if name == "" {
			d.State.SetError("empty name")
			break
		}

		if version == "" {
			version = "latest"
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		var containers []types.Container
		var err error
		containers, err = d.Client.ContainerList(ctx, types.ContainerListOptions{All: true, Limit: 256})
		if err != nil {
			d.State.SetError("list error: %s", err)
			break
		}

		tagCheck := fmt.Sprintf("%s/%s:%s", org, name, version)
		// Start out with "not found". Will be cleared if found or error occurs.
		d.State.SetWarning("Container '%s' doesn't exist.", tagCheck)

		for _, c := range containers {
			if c.Image != tagCheck {
				continue
			}

			d.Container.Name = name
			d.Container.Version = version
			d.Container.Summary = &c
			d.Container.ID = c.ID
			d.Container.State = d.Container.State.EnsureNotNil()
			ok = true
			d.State.SetOk("Found container '%s'.", tagCheck)
			break
		}

		if d.State.IsNotOk() {
			if !ok {
				d.State.ClearError()
			}
			break
		}

		if d.Container.Summary == nil {
			break
		}

		ctx2, cancel2 := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel2()

		dc := types.ContainerJSON{}
		dc, err = d.Client.ContainerInspect(ctx2, d.Container.ID)
		if err != nil {
			d.State.SetError("d inspect error: %s", err)
			break
		}
		d.Container.Details = &dc
	}

	return ok, d.State
}
