package toolGear

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/newclarity/scribeHelpers/toolGear/gearConfig"
	"github.com/newclarity/scribeHelpers/toolGear/gearSsh"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
	"time"
)


type Container struct {
	ID           string
	Name         string
	Version      string
	IsLatest     bool

	VolumeMounts VolumeMounts
	SshfsMounts  SshfsMounts
	Ssh          *gearSsh.Ssh

	Summary      types.Container
	Details      *types.ContainerJSON
	GearConfig   *gearConfig.GearConfig

	Docker       *Docker
	runtime      *toolRuntime.TypeRuntime
	State        *ux.State
}
type Containers []Container


func NewContainer(runtime *toolRuntime.TypeRuntime) *Container {
	runtime = runtime.EnsureNotNil()

	c := &Container{
		ID:         "",
		Name:       "",
		Version:    "",
		IsLatest:   false,

		VolumeMounts: make(VolumeMounts),
		SshfsMounts:  make(SshfsMounts),
		Ssh:        nil,

		Summary:    types.Container{},
		Details:    &types.ContainerJSON{},
		GearConfig: gearConfig.New(runtime),

		Docker:    nil,
		runtime:    runtime,
		State:      ux.NewState(runtime.CmdName, runtime.Debug),
	}
	c.State.SetPackage("")
	c.State.SetFunctionCaller()
	return c
}

func (c *Container) EnsureNotNil() *Container {
	for range onlyOnce {
		if c == nil {
			c = NewContainer(nil)
		}
		c.State = c.State.EnsureNotNil()
	}
	return c
}

func (c *Container) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}

func (c *Container) IsValid() bool {
	var ok bool
	if state := ux.IfNilReturnError(c); state.IsError() {
		return ok
	}

	for range onlyOnce {
		c.State = c.State.EnsureNotNil()

		if c.ID == "" {
			c.State.SetError("gear ID is nil")
			break
		}

		if c.Name == "" {
			c.State.SetError("gear name is nil")
			break
		}

		if c.Version == "" {
			c.State.SetError("gear version is nil")
			break
		}

		if c.Docker.Client == nil {
			c.State.SetError("docker client is nil")
			break
		}

		ok = true
	}

	return ok
}
func (c *Container) IsNotValid() bool {
	return !c.IsValid()
}

func (c *Container) Refresh() *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		c.Summary, c.State = c.Docker.GetContainerById(c.ID)
		if c.State.IsNotOk() {
			break
		}

		c.Details, c.State = c.Docker.ContainerInspect(c.ID)
		if c.State.IsNotOk() {
			break
		}

		c.GearConfig = c.GearConfig.EnsureNotNil()
		c.State = c.GearConfig.ParseJson(c.Summary.Labels["gearbox.json"])
		if c.State.IsError() {
			break
		}

		c.State.RunState = c.Summary.State
	}

	return c.State
}

func (c *Container) WaitForState(s string, t time.Duration) *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		until := time.Now()
		until.Add(t)

		for now := time.Now(); until.Before(now); now = time.Now() {
			c.State = c.Refresh()
			if c.State.IsError() {
				break
			}

			if c.State.RunStateEquals(s) {
				break
			}
		}
	}

	return c.State
}

// Run a container in the background
// You can also run containers in the background, the equivalent of typing docker run -d bfirsh/reticulate-splines:
func (c *Container) Start() *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		c.State = c.Refresh()
		if c.State.IsError() {
			break
		}

		if c.State.IsRunning() {
			break
		}

		if !c.State.IsCreated() && !c.State.IsExited() {
			break
		}

		c.State = c.Docker.ContainerStart(c.ID, nil)

		c.State = c.WaitForState(ux.StateRunning, DefaultTimeout)
		if c.State.IsError() {
			break
		}

		if !c.State.IsRunning() {
			c.State.SetError("Gear '%s:%s' failed to start.", c.Name, c.Version)
			break
		}

		c.State = c.Refresh()
	}

	return c.State
}

// Stop all running containers
// Now that you know what containers exist, you can perform operations on them.
// This example stops all running containers.
func (c *Container) Stop() *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		c.State = c.Docker.ContainerStop(c.ID, nil)
		if c.State.IsNotOk() {
			break
		}

		c.State = c.Refresh()
	}

	return c.State
}

// Remove containers
// Now that you know what containers exist, you can perform operations on them.
// This example stops all running containers.
func (c *Container) Remove() *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		options := types.ContainerRemoveOptions{
			RemoveVolumes: false,
			RemoveLinks:   false,
			Force:         false,
		}
		c.State = c.Docker.ContainerRemove(c.ID, &options)
		if c.State.IsNotOk() {
			break
		}

		c.State.SetOk("OK")
	}

	return c.State
}

// Print the logs of a specific container
// You can also perform actions on individual containers.
// This example prints the logs of a container given its ID.
// You need to modify the code before running it to change the hard-coded ID of the container to print the logs for.
func (c *Container) Logs() *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		options := types.ContainerLogsOptions{ShowStdout: true}

		c.State = c.Docker.ContainerLogs(c.ID, options)
		if c.State.IsNotOk() {
			break
		}

		//c.State = c.Refresh()
	}

	return c.State
}

// Commit a container
// Commit a container to create an image from its contents:
func (c *Container) Commit() *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		c.State = c.Docker.ContainerCommit()
	}

	return c.State
}

func (c *Container) GetContainerSsh() (string, *ux.State) {
	var port string
	if state := c.IsNil(); state.IsError() {
		return port, state
	}

	for range onlyOnce {
		var found bool
		for _, p := range c.Summary.Ports {
			if p.PrivatePort == 22 {
				port = fmt.Sprintf("%d", p.PublicPort)
				found = true
				break
			}
		}

		if !found {
			c.State.SetError("no SSH port")
		}
	}

	return port, c.State
}

func (c *Container) GetName() string {
	return strings.TrimPrefix(c.Summary.Names[0], "/")
}

func (c *Container) GetVersion() string {
	return c.Summary.Labels["gearbox.version"]
}

func (c *Container) GetMounts() []types.MountPoint {
	return c.Summary.Mounts
}

func (c *Container) GetNetworks() map[string]*network.EndpointSettings {
	return c.Summary.NetworkSettings.Networks
}

func (c *Container) GetState() string {
	return c.Summary.State
}

func (c *Container) GetSize() uint64 {
	return uint64(c.Summary.SizeRootFs)
}

func (c *Container) GetLabels() map[string]string {
	return c.Summary.Labels
}

func (c *Container) GetVolumeMounts() []string {
	var ret []string
	if state := c.IsNil(); state.IsError() {
		return ret
	}

	for range onlyOnce {
		for k, v := range c.VolumeMounts {
			ret = append(ret, fmt.Sprintf("%s:%s", k, v))
		}
	}

	return ret
}

func (c *Container) GetFixedPorts() nat.PortMap {
	if state := c.IsNil(); state.IsError() {
		return nil
	}
	return c.GearConfig.GetFixedPorts()
}


type Port struct {
	Name string
	types.Port
}
type Ports map[uint16]*Port
func (c *Container) GetPorts() Ports {
	ports := make(Ports)

	for range onlyOnce {
		//if c.Summary == nil {
		//	break
		//}
		var found bool
		for _, p := range c.Summary.Ports {
			ports[p.PrivatePort] = &Port {
				Name: "",
				Port: p,
			}
		}

		if !found {
			c.State.SetError("no ports")
		}
	}

	return ports
}
