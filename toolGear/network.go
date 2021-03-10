package toolGear

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/jedib0t/go-pretty/table"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
)


// List and manage containers
// You can use the API to list containers that are running, just like using docker ps:
// func ContainerList(f types.ContainerListOptions) error {
func (gears *Gears) NetworkList(name string) *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.Docker.NetworkList(name)
		if gears.State.IsNotOk() {
			break
		}

		ux.PrintflnCyan("\nConfigured %s networks:", gears.Language.AppName)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Name",
			"Driver",
			"Subnet",
		})

		for _, c := range gears.Docker.Networks {
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

	return gears.State
}


func (gears *Gears) FindNetwork(name string) *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if name == "" {
			gears.State.SetError("empty %s name", gears.Language.ContainerName)
			break
		}

		//df := filters.NewArgs()
		//df.Add("name", netName)
		//
		//ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		////noinspection GoDeferInLoop
		//defer cancel()
		//
		//nets, err := gear.Docker.Client.NetworkList(ctx, types.NetworkListOptions{Filters: df})
		//if err != nil {
		//	gear.State.SetError("gear image search error: %s", err)
		//	break
		//}

		gears.State = gears.Docker.NetworkList(name)
		if gears.State.IsNotOk() {
			break
		}

		for _, c := range gears.Docker.Networks {
			if c.Name == name {
				gears.State.SetOk("found")
				break
			}
		}
	}

	return gears.State
}


func (gears *Gears) NetworkCreate(netName string) *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.FindNetwork(netName)
		if gears.State.IsError() {
			break
		}
		if gears.State.IsOk() {
			break
		}

		options := types.NetworkCreate {
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

		gears.State = gears.Docker.NetworkCreate(netName, options)
	}

	return gears.State
}
