package toolDocker

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type TypeDocker struct {
	Image     *Image
	Container *Container
	Client    *client.Client

	Runtime   *toolRuntime.TypeRuntime
	State     *ux.State
}
func (d *TypeDocker) IsNil() *ux.State {
	return ux.IfNilReturnError(d)
}


func New(runtime *toolRuntime.TypeRuntime) *TypeDocker {
	var gear TypeDocker
	runtime = runtime.EnsureNotNil()

	for range onlyOnce {
		gear.State = ux.NewState(runtime.CmdName, runtime.Debug)

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
			gear.State.SetError("TypeDocker client error: %s", err)
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		//var result types.Ping
		_, err = gear.Client.Ping(ctx)
		if err != nil {
			gear.State.SetError("TypeDocker client error: %s", err)
			break
		}
		//fmt.Printf("PING: %v", result)

		gear.Image._Parent = &gear
		gear.Container._Parent = &gear
	}

	return &gear
}


func (d *TypeDocker) IsValid() *ux.State {
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


type TypeMatchImage struct {
	Organization string
	Name         string
	Version      string
}
type TypeMatchContainer TypeMatchImage
