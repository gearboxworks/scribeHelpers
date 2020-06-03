package toolDocker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/jedib0t/go-pretty/table"
	//"launch/defaults"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
)

// List and manage containers
// You can use the API to list containers that are running, just like using docker ps:
// func ContainerList(f types.ContainerListOptions) error {
func (d *Docker) NetworkList(f string) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		df := filters.NewArgs()
		df.Add("name", f)

		nets, err := d.Client.NetworkList(ctx, types.NetworkListOptions{Filters: df})
		if err != nil {
			d.State.SetError("error listing networks")
			break
		}

		ux.PrintflnCyan("\nConfigured Docker networks:")
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Name",
			"Driver",
			"Subnet",
		})

		for _, c := range nets {
			n := ""
			if len(c.IPAM.Config) > 0 {
				n = c.IPAM.Config[0].Subnet
			}

			t.AppendRow([]interface{}{
				ux.SprintfWhite(c.Name),
				ux.SprintfWhite(c.Driver),
				ux.SprintfWhite(n),
			})
		}

		t.Render()
	}

	return d.State
}


func (d *Docker) FindNetwork(netName string) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		if netName == "" {
			d.State.SetError("empty name")
			break
		}

		df := filters.NewArgs()
		df.Add("name", netName)

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		nets, err := d.Client.NetworkList(ctx, types.NetworkListOptions{Filters: df})
		if err != nil {
			d.State.SetError("image search error: %s", err)
			break
		}

		for _, c := range nets {
			if c.Name == netName {
				d.State.SetOk("found")
				break
			}
		}
	}

	return d.State
}


func (d *Docker) NetworkCreate(netName string) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		d.State = d.FindNetwork(netName)
		if d.State.IsError() {
			break
		}
		if d.State.IsOk() {
			break
		}

		netConfig := types.NetworkCreate {
			CheckDuplicate: true,
			Driver:         "bridge",
			Scope:          "local",
			EnableIPv6:     false,
			IPAM:           &network.IPAM {
				Driver:  "default",
				Options: nil,
				Config:  []network.IPAMConfig {
					{
						Subnet: "172.42.0.0/24",
						Gateway: "172.42.0.1",
					},
				},
			},
			Internal:       false,
			Attachable:     false,
			Ingress:        false,
			ConfigOnly:     false,
			ConfigFrom:     nil,
			Options:        nil,
			Labels:         nil,
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		resp, err := d.Client.NetworkCreate(ctx, netName, netConfig)
		d.State.SetError(err)
		if d.State.IsError() {
			break
		}

		if resp.ID == "" {
			d.State.SetError("cannot create network")
			break
		}
	}

	return d.State
}
