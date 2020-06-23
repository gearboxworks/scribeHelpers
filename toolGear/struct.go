package toolGear

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/newclarity/scribeHelpers/toolGear/gearConfig"
	"github.com/newclarity/scribeHelpers/toolGear/gearSsh"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type DockerGear struct {
	Image     *Image
	Container *Container

	gearConfig *gearConfig.GearConfig

	Client    *client.Client
	Ssh       *gearSsh.Ssh

	Runtime   *toolRuntime.TypeRuntime
	State     *ux.State
}


type TypeMatchImage struct {
	Organization string
	Name         string
	Version      string
}
type TypeMatchContainer TypeMatchImage


func New(runtime *toolRuntime.TypeRuntime) *DockerGear {
	var gear DockerGear
	runtime = runtime.EnsureNotNil()

	for range onlyOnce {
		gear.State = ux.NewState(runtime.CmdName, runtime.Debug)
		gear.Runtime = runtime

		gear.Image = NewImage(runtime)
		gear.Container = NewContainer(runtime)

		gear.gearConfig = gearConfig.New(runtime)

		var err error
		gear.Client, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		//cli.DockerClient, err = client.NewEnvClient()
		if err != nil {
			gear.State.SetError("Docker client error: %s", err)
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		_, err = gear.Client.Ping(ctx)
		if err != nil {
			gear.State.SetError("Docker client error: %s", err)
			break
		}

		gear.Image._Parent = &gear
		gear.Container._Parent = &gear
	}

	return &gear
}


func (gear *DockerGear) IsNil() *ux.State {
	if state := ux.IfNilReturnError(gear); state.IsError() {
		return state
	}
	gear.State = gear.State.EnsureNotNil()
	return gear.State
}


func (gear *DockerGear) IsValid() *ux.State {
	if state := ux.IfNilReturnError(gear); state.IsError() {
		return state
	}

	for range onlyOnce {
		gear.State = gear.State.EnsureNotNil()

		if gear.Client == nil {
			gear.State.SetError("docker client is nil")
			break
		}
	}

	return gear.State
}


func (gear *DockerGear) SetSshStatusLine(s bool) {
	gear.Ssh.StatusLine.Enable = s
}


func (gear *DockerGear) SetSshShell(s bool) {
	gear.Ssh.Shell = s
}


func (gear *DockerGear) AddVolume(local string, remote string) bool {
	if gear.Container.VolumeMounts == nil {
		gear.Container.VolumeMounts = make(VolumeMounts)
	}
	return gear.Container.VolumeMounts.Add(local, remote)
}


func (gear *DockerGear) ContainerCreate(gearName string, gearVersion string) *ux.State {
	return gear.Container.ContainerCreate(gearName, gearVersion)
}


func (gear *DockerGear) List(name string) *ux.State {
	if state := ux.IfNilReturnError(gear); state.IsError() {
		return state
	}

	for range onlyOnce {
		_, gear.State = gear.ImageList(name)
		if gear.State.IsError() {
			break
		}

		_, gear.State = gear.ContainerList(name)
		if gear.State.IsError() {
			break
		}

		gear.State = gear.NetworkList(DefaultNetwork)
	}

	return gear.State
}


func (gear *DockerGear) ParseGearConfig(cs string) *ux.State {
	if state := ux.IfNilReturnError(gear); state.IsError() {
		return state
	}

	for range onlyOnce {
		gear.State = gear.gearConfig.ParseJson(cs)
		if gear.State.IsNotOk() {
			break
		}
	}

	return gear.State
}
