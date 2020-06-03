package helperGear

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/newclarity/scribeHelpers/helperGear/gearSsh"
	"github.com/newclarity/scribeHelpers/helperRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type DockerGear struct {
	Image     *Image
	Container *Container

	Client    *client.Client
	Ssh       *gearSsh.Ssh

	Debug     bool
	Runtime   *helperRuntime.TypeRuntime
	State     *ux.State
}


type TypeMatchImage struct {
	Organization string
	Name         string
	Version      string
}
type TypeMatchContainer TypeMatchImage


func New(runtime *helperRuntime.TypeRuntime) *DockerGear {
	var gear DockerGear
	runtime = runtime.EnsureNotNil()

	for range OnlyOnce {
		gear.State = ux.NewState(runtime.CmdName, runtime.Debug)
		gear.Debug = runtime.Debug

		gear.Image = NewImage(runtime)
		gear.Container = NewContainer(runtime)
		gear.Runtime = runtime
		//gear.Image = *gear.Image.EnsureNotNil()
		//gear.State.DebugSet(debugMode)
		//gear.Container = *gear.Container.EnsureNotNil()
		//gear.State.DebugSet(debugMode)

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

		//var result types.Ping
		_, err = gear.Client.Ping(ctx)
		if err != nil {
			gear.State.SetError("Docker client error: %s", err)
			break
		}
		//fmt.Printf("PING: %v", result)

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

	for range OnlyOnce {
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
	return gear.Container.VolumeMounts.Add(local, remote)
}


func (gear *DockerGear) ContainerCreate(gearName string, gearVersion string) *ux.State {
	return gear.Container.ContainerCreate(gearName, gearVersion)
}
