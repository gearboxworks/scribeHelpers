package toolGear

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/jedib0t/go-pretty/table"
	//"launch/defaults"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
)

// List and manage containers
// You can use the API to list containers that are running, just like using docker ps:
// func ContainerList(f types.ContainerListOptions) error {
func (gear *Gears) NetworkList(name string) *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gear.State = gear.Docker.NetworkList(name)
		if gear.State.IsNotOk() {
			break
		}

		ux.PrintflnCyan("\nConfigured Gearbox networks:")
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Name",
			"Driver",
			"Subnet",
		})

		for _, c := range gear.Docker.Networks {
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

	return gear.State
}


func (gear *Gears) FindNetwork(name string) *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if name == "" {
			gear.State.SetError("empty container name")
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

		gear.State = gear.Docker.NetworkList(name)
		if gear.State.IsNotOk() {
			break
		}

		for _, c := range gear.Docker.Networks {
			if c.Name == name {
				gear.State.SetOk("found")
				break
			}
		}
	}

	return gear.State
}


func (gear *Gears) NetworkCreate(netName string) *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gear.State = gear.FindNetwork(netName)
		if gear.State.IsError() {
			break
		}
		if gear.State.IsOk() {
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

		gear.State = gear.Docker.NetworkCreate(netName, options)

		//ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		////noinspection GoDeferInLoop
		//defer cancel()
		//
		//resp, err := gear.Docker.Client.NetworkCreate(ctx, netName, netConfig)
		//gear.State.SetError(err)
		//if gear.State.IsError() {
		//	break
		//}
		//
		//if resp.ID == "" {
		//	gear.State.SetError("cannot create network")
		//	break
		//}
	}

	return gear.State
}
