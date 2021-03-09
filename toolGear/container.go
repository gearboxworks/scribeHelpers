package toolGear

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	//"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/newclarity/scribeHelpers/toolGear/gearConfig"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
	"time"
)


type Container struct {
	ID           string
	Name         string
	Version      string

	VolumeMounts VolumeMounts
	SshfsMounts  SshfsMounts

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

		VolumeMounts: make(VolumeMounts),
		SshfsMounts:  make(SshfsMounts),

		Summary:    types.Container{},
		Details:    nil,
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


func (c *Container) IsValid() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
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
	}

	return c.State
}


func (c *Container) Status() *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//df := filters.NewArgs()
		//df.Add("id", c.ID)
		//
		//ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		////noinspection GoDeferInLoop
		//defer cancel()
		//
		//var containers []types.Container
		//var err error
		//containers, err = c._Parent.Client.ContainerList(ctx, types.ContainerListOptions{All: true, Filters: df})
		//if err != nil {
		//	c.State.SetError("gear list error: %s", err)
		//	break
		//}
		//if len(containers) == 0 {
		//	c.State.SetWarning("no gears found")
		//	break
		//}

		//containers := c.Docker.GetContainerById()
		//
		//c.Summary = containers[0]

		c.GearConfig = c.GearConfig.EnsureNotNil()
		c.State = c.GearConfig.ParseJson(c.Summary.Labels["gearbox.json"])
		if c.State.IsError() {
			break
		}

		if c.GearConfig.Meta.Organization != DefaultOrganization {
			c.State.SetError("not a Gearbox container")
			break
		}

		//d := types.ContainerJSON{}
		//d, err = c.Docker.Client.ContainerInspect(ctx, c.ID)
		//if err != nil {
		//	c.State.SetError("gear inspect error: %s", err)
		//	break
		//}

		c.Details, c.State = c.Docker.ContainerInspect(c.ID)

		c.State.SetRunState(c.Details.State.Status)
	}

	if c.State.IsError() {
		c.Details = nil
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
			c.State = c.Status()
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
		c.State = c.Status()
		if c.State.IsError() {
			break
		}

		if c.State.IsRunning() {
			break
		}

		if !c.State.IsCreated() && !c.State.IsExited() {
			break
		}

		c.State = c.Docker.ContainerStart(c.ID, types.ContainerStartOptions{})

		//ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		////noinspection GoDeferInLoop
		//defer cancel()
		//
		//err := c.Docker.Client.ContainerStart(ctx, c.ID, types.ContainerStartOptions{})
		//if err != nil {
		//	c.State.SetError("Gear '%s:%s' start error - %s", c.Name, c.Version, err)
		//	break
		//}

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

		c.State = c.WaitForState(ux.StateRunning, DefaultTimeout)
		if c.State.IsError() {
			break
		}
		if !c.State.IsRunning() {
			c.State.SetError("Gear '%s:%s' failed to start.", c.Name, c.Version)
			break
		}
	}

	return c.State
}


// Run a container
// This first example shows how to run a container using the Docker API.
// On the command line, you would use the docker run command, but this is just as easy to do from your own apps too.
// This is the equivalent of typing docker run alpine echo hello world at the command prompt:
func (c *Gears) ContainerCreate(gearName string, gearVersion string) *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//if c.runtime.Debug {
		//	fmt.Printf("DEBUG: ContainerCreate(%s, %s)\n", gearName, gearVersion)
		//}

		if gearName == "" {
			c.State.SetError("empty gearname")
			break
		}

		if gearVersion == "" {
			gearVersion = "latest"
		}

		var ok bool
		ok, c.State = c.FindContainer(gearName, gearVersion)
		if c.State.IsError() {
			break
		}
		if !ok {
			// Find Gear image since we don't have a container.
			for range onlyOnce {
				ok, c.State = c.FindImage(gearName, gearVersion)
				if c.State.IsError() {
					ok = false
					break
				}
				if ok {
					break
				}

				ux.PrintflnNormal("Downloading Gear '%s:%s'.", gearName, gearVersion)

				// Pull Gear image.
				c.Selected.Image.ID = gearName
				c.Selected.Image.Name = gearName
				c.Selected.Image.Version = gearVersion
				c.State = c.Selected.Image.Pull()
				if c.State.IsError() {
					c.State.SetError("no such gear '%s'", gearName)
					break
				}

				// Confirm it's there.
				ok, c.State = c.FindImage(gearName, gearVersion)
				if c.State.IsError() {
					ok = false
					break
				}
			}
			if !ok {
				c.State.SetError("Cannot install Gear image '%s:%s' - %s.", gearName, gearVersion, c.State.GetError())
				break
			}
			//c.State.Clear()
		}

		c.Selected.Container.ID = c.Selected.Image.ID
		c.Selected.Container.Name = c.Selected.Image.Name
		c.Selected.Container.Version = c.Selected.Image.Version

		// c.Image.Details.Container = "gearboxworks/golang:1.14"
		// tag := fmt.Sprintf("", c.Image.Name, c.Image.Version)
		tag := fmt.Sprintf("gearboxworks/%s:%s", c.Selected.Image.Name, c.Selected.Image.Version)
		gn := fmt.Sprintf("%s-%s", c.Selected.Image.Name, c.Selected.Image.Version)

		var binds []string
		for k, v := range c.Selected.Container.VolumeMounts {
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

		hostConfig := container.HostConfig {
			Binds:           binds,
			ContainerIDFile: "",
			LogConfig:       container.LogConfig{
				Type:   "",
				Config: nil,
			},
			NetworkMode:     DefaultNetwork,
			PortBindings:    nil,						// @TODO
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
		resp, c.State = c.Docker.ContainerCreate(&config, &hostConfig, &netConfig, gn)

		c.Selected.Container.ID = resp.ID
		//c.Container.Name = c.Image.Name
		//c.Container.Version = c.Image.Version

		// var response Response
		c.State = c.Status()
		if c.State.IsError() {
			break
		}

		if c.State.IsCreated() {
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
		//ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		////noinspection GoDeferInLoop
		//defer cancel()
		//
		//err := c.Docker.Client.ContainerStop(ctx, c.ID, nil)
		//if err != nil {
		//	c.State.SetError("gear stop error: %s", err)
		//	break
		//}

		c.State = c.Docker.ContainerStop(c.ID, nil)
		if c.State.IsNotOk() {
			break
		}

		c.State = c.Status()
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
		c.State = c.Docker.ContainerRemove(c.ID, options)

		//ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		////noinspection GoDeferInLoop
		//defer cancel()
		//
		//err := c.Docker.Client.ContainerRemove(ctx, c.ID, options)
		//if err != nil {
		//	c.State.SetError("gear remove error: %s", err)
		//	break
		//}

		//c.State = c.State()
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

		//ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		////noinspection GoDeferInLoop
		//defer cancel()
		//
		//// Replace this ID with a container that really exists
		//out, err := c.Docker.Client.ContainerLogs(ctx, c.ID, options)
		//if err != nil {
		//	c.State.SetError("gear logs error: %s", err)
		//	break
		//}

		c.State = c.Docker.ContainerLogs(c.ID, options)

		//_, _ = io.Copy(os.Stdout, out)

		c.State = c.Status()
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


func (c *Container) GetName() (string) {
	return strings.TrimPrefix(c.Summary.Names[0], "/")
}

func (c *Container) GetVersion() (string) {
	return c.Summary.Labels["gearbox.version"]
}

func (c *Container) GetMounts() ([]types.MountPoint) {
	return c.Summary.Mounts
}

func (c *Container) GetNetworks() (map[string]*network.EndpointSettings) {
	return c.Summary.NetworkSettings.Networks
}

func (c *Container) GetState() (string) {
	return c.Summary.State
}

func (c *Container) GetSize() (uint64) {
	return uint64(c.Summary.SizeRootFs)
}

func (c *Container) GetLabels() (map[string]string) {
	return c.Summary.Labels
}


type Port struct {
	Name string
	types.Port
}
type Ports map[uint16]Port
func (c *Container) GetPorts() (Ports) {
	ports := make(Ports)

	for range onlyOnce {
		//if c.Summary == nil {
		//	break
		//}
		var found bool
		for _, p := range c.Summary.Ports {
			ports[p.PrivatePort] = Port {
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
