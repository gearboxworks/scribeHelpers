package toolGear

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"io"
	"os"
	"time"
)

//"github.com/docker/docker/integration-cli/cli"
// DOCKER_HOST=tcp://macpro:2375


type Docker struct {
	Containers []types.Container
	Images     []types.ImageSummary
	Networks   []types.NetworkResource

	Client    	*client.Client

	Runtime     *toolRuntime.TypeRuntime
	State       *ux.State
}


func NewDocker(runtime *toolRuntime.TypeRuntime) *Docker {
	var d Docker

	for range onlyOnce {
		runtime = runtime.EnsureNotNil()

		d = Docker{
			Containers:     nil,
			Images: nil,

			Client:         nil,

			Runtime:        runtime,
			State:          ux.NewState(runtime.CmdName, runtime.Debug),
		}

		//foo := os.Getenv("DOCKER_HOST")
		//fmt.Printf("DOCKER_HOST:%s\n", foo)

		var err error
		d.Client, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		//cli.DockerClient, err = client.NewEnvClient()
		if err != nil {
			d.State.SetError("Docker client error: %s", err)
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		_, err = d.Client.Ping(ctx)
		if err != nil {
			//gears.State.SetError("Docker client error: %s", err)
			d.State.SetError("can not connect to Docker service provider - maybe you haven't set DOCKER_HOST, or Docker not running on this host")
			break
		}

		d.State.SetPackage("")
		d.State.SetFunctionCaller()

		if d.State.IsNotOk() {
			d.State.SetError("can not connect to Docker service provider - maybe you haven't set DOCKER_HOST, or Docker not running on this host")
			//gear.State = gear.Docker.State
			break
		}
	}

	return &d
}


func (d *Docker) IsValid() *ux.State {
	if state := ux.IfNilReturnError(d); state.IsError() {
		return state
	}

	for range onlyOnce {
		d.State = d.State.EnsureNotNil()

		if d.Client == nil {
			d.State.SetError("docker client is nil")
			break
		}
	}

	return d.State
}


func (d *Docker) IsNil() *ux.State {
	if state := ux.IfNilReturnError(d); state.IsError() {
		return state
	}

	for range onlyOnce {
		d.State = d.State.EnsureNotNil()
	}

	return d.State
}


func (d *Docker) GetContainerById(containerID string) (types.Container, *ux.State) {
	var container types.Container
	if state := d.IsNil(); state.IsError() {
		return container, state
	}

	for range onlyOnce {
		df := filters.NewArgs()
		df.Add("id", containerID)

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		containers, err := d.Client.ContainerList(ctx, types.ContainerListOptions{All: true, Filters: df})
		if err != nil {
			d.State.SetError("gear list error: %s", err)
			break
		}
		if len(containers) == 0 {
			d.State.SetWarning("no gears found")
			break
		}
		container = containers[0]

		//c.Summary = &containers[0]
		//
		//c.GearConfig = c.GearConfig.EnsureNotNil()
		//c.State = c.GearConfig.ParseJson(c.Summary.Labels["gearbox.json"])
		//if c.State.IsError() {
		//	break
		//}
		//
		//if c.GearConfig.Meta.Organization != DefaultOrganization {
		//	c.State.SetError("not a Gearbox container")
		//	break
		//}
		//
		//d := types.ContainerJSON{}
		//d, err = c._Parent.Client.ContainerInspect(ctx, c.ID)
		//if err != nil {
		//	c.State.SetError("gear inspect error: %s", err)
		//	break
		//}
		//c.Details = &d
		//
		//c.State.SetRunState(c.Details.State.Status)
	}

	return container, d.State
}


func (d *Docker) ContainerList(force bool) (*ux.State) {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if force {
			d.Containers = nil
		}

		if d.Containers != nil {
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		//var containers []types.Container
		var err error
		d.Containers, err = d.Client.ContainerList(ctx, types.ContainerListOptions{Size: true, All: true})
		if err != nil {
			d.State.SetError("container list error: %s", err)
			break
		}

		if len(d.Containers) == 0 {
			d.State.SetWarning("no containers found")
			break
		}
	}

	return d.State
}


func (d *Docker) ContainerStart(containerID string, options *types.ContainerStartOptions) (*ux.State) {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if d.Containers == nil {
			d.State.SetWarning("no containers found")
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		//statusCh, errCh := c._Parent.Client.ContainerWait(ctx, c.ID, "") // container.WaitConditionNotRunning
		//select {
		//	case err := <-errCh:
		//		if err != nil {
		//			c.State.SetError("Docker client error: %s", err)
		//			// fmt.Printf("SC: %s\n", response.Error)
		//			// return false, err
		//		}
		//		break
		//
		//	case status := <-statusCh:
		//		fmt.Printf("status.StatusCode: %#+v\n", status.StatusCode)
		//		break
		//}
		// fmt.Printf("SC: %s\n", status)
		// fmt.Printf("SC: %s\n", err)

		err := d.Client.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
		if err != nil {
			d.State.SetError("Container start error - %s", err)
			break
		}
	}

	return d.State
}


func (d *Docker) ContainerStop(containerID string, timeout *time.Duration) (*ux.State) {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if d.Containers == nil {
			d.State.SetWarning("no containers found")
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		err := d.Client.ContainerStop(ctx, containerID, timeout)
		if err != nil {
			d.State.SetError("container stop error: %s", err)
			break
		}
	}

	return d.State
}


func (d *Docker) ContainerRemove(containerID string, options *types.ContainerRemoveOptions) (*ux.State) {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if d.Containers == nil {
			d.State.SetWarning("no containers found")
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		err := d.Client.ContainerRemove(ctx, containerID, *options)
		if err != nil {
			d.State.SetError("container remove error: %s", err)
			break
		}
	}

	return d.State
}


func (d *Docker) ContainerLogs(containerID string, options types.ContainerLogsOptions) (*ux.State) {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if d.Containers == nil {
			d.State.SetWarning("no containers found")
			break
		}
		options = types.ContainerLogsOptions{ShowStdout: true}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		// Replace this ID with a container that really exists
		out, err := d.Client.ContainerLogs(ctx, containerID, options)
		if err != nil {
			d.State.SetError("container logs error: %s", err)
			break
		}
		_, _ = io.Copy(os.Stdout, out)

		d.State.SetOutput(out)
		d.State.SetOk()
	}

	return d.State
}


func (d *Docker) ContainerCommit() (*ux.State) {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if d.Containers == nil {
			d.State.SetWarning("no containers found")
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		createResp, err := d.Client.ContainerCreate(ctx, &container.Config{
			Image: "alpine",
			Cmd:   []string{"touch", "/helloworld"},
		}, nil, nil, "")
		if err != nil {
			break
		}

		if err := d.Client.ContainerStart(ctx, createResp.ID, types.ContainerStartOptions{}); err != nil {
			break
		}

		statusCh, errCh := d.Client.ContainerWait(ctx, createResp.ID, container.WaitConditionNotRunning)
		select {
		case err := <-errCh:
			if err != nil {
				//response.State.SetError("gear stop error: %s", err)
				break
			}
		case <-statusCh:
		}

		commitResp, err := d.Client.ContainerCommit(ctx, createResp.ID, types.ContainerCommitOptions{Reference: "helloworld"})
		if err != nil {
			break
		}

		fmt.Println(commitResp.ID)

		d.State.SetOk()
	}

	return d.State
}


func (d *Docker) ContainerInspect(containerID string) (*types.ContainerJSON, *ux.State) {
	var ret types.ContainerJSON
	if state := d.IsNil(); state.IsError() {
		return &ret, state
	}

	for range onlyOnce {
		if d.Containers == nil {
			d.State.SetWarning("no containers found")
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		var err error
		ret, err = d.Client.ContainerInspect(ctx, containerID)
		if err != nil {
			d.State.SetError("container inspect error: %s", err)
			break
		}

		d.State.SetOutput(ret)
		d.State.SetOk()
	}

	return &ret, d.State
}


func (d *Docker) ContainerCreate(config *container.Config, hostConfig *container.HostConfig, netConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, *ux.State) {
	var resp container.ContainerCreateCreatedBody
	if state := d.IsNil(); state.IsError() {
		return resp, state
	}

	for range onlyOnce {
		//if d.Containers == nil {
		//	d.State.SetWarning("no containers found")
		//	break
		//}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		//var resp container.ContainerCreateCreatedBody
		var err error
		resp, err = d.Client.ContainerCreate(ctx, config, hostConfig, netConfig, containerName)
		if err != nil {
			d.State.SetError("error creating container: %s", err)
			break
		}

		if resp.ID == "" {
			d.State.SetError("error creating container")
			break
		}

		d.State.SetOk()
	}

	return resp, d.State
}


func (d *Docker) ImageList(options *types.ImageListOptions) (*ux.State) {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var err error
		df := filters.NewArgs()

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		d.Images, err = d.Client.ImageList(ctx, types.ImageListOptions{All: true, Filters: df})
		if err != nil {
			d.State.SetError("image list error: %s", err)
			break
		}

		if len(d.Images) == 0 {
			d.State.SetWarning("no images found")
			break
		}

		d.State.SetOk()
	}

	return d.State
}


func (d *Docker) ImageInspectWithRaw(imageID string) (types.ImageInspect, *ux.State) {
	var resp types.ImageInspect
	if state := d.IsNil(); state.IsError() {
		return resp, state
	}

	for range onlyOnce {
		var err error

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		resp, _, err = d.Client.ImageInspectWithRaw(ctx, imageID)
		if err != nil {
			d.State.SetError("error inspecting image: %s", err)
			break
		}

		d.State.SetOk()
	}

	return resp, d.State
}


func (d *Docker) ImageSearch(repo string, options *types.ImageSearchOptions) ([]registry.SearchResult, *ux.State) {
	var resp []registry.SearchResult
	if state := d.IsNil(); state.IsError() {
		return resp, state
	}

	for range onlyOnce {
		var err error

		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		df := filters.NewArgs()
		//df.Add("name", "terminus")
		options = &types.ImageSearchOptions{Filters: df, Limit: 100}

		resp, err = d.Client.ImageSearch(ctx, repo, *options)
		if err != nil {
			d.State.SetError("gear image search error: %s", err)
			break
		}

		d.State.SetOk()
	}

	return resp, d.State
}


func (d *Docker) NetworkList(name string) (*ux.State) {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//if d.Networks != nil {
		//	break
		//}

		var err error

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		df := filters.NewArgs()
		df.Add("name", name)

		d.Networks, err = d.Client.NetworkList(ctx, types.NetworkListOptions{Filters: df})
		if err != nil {
			d.State.SetError("error listing networks")
			break
		}

		d.State.SetOk()
	}

	return d.State
}


func (d *Docker) NetworkCreate(name string, options types.NetworkCreate) (*ux.State) {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var err error

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		resp, err := d.Client.NetworkCreate(ctx, name, options)
		if err != nil {
			d.State.SetError(err)
			break
		}

		if resp.ID == "" {
			d.State.SetError("cannot create network")
			break
		}

		d.State.SetOk()
	}

	return d.State
}


func (d *Docker) ImageRemove(imageID string, options *types.ImageRemoveOptions) (*ux.State) {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var err error

		options := types.ImageRemoveOptions {
			Force:         true,
			PruneChildren: true,
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		_, err = d.Client.ImageRemove(ctx, imageID, options)
		if err != nil {
			d.State.SetError("error removing: %s", err)
			break
		}

		d.State.SetOk()
	}

	return d.State
}
