package toolGear

import (
	"github.com/newclarity/scribeHelpers/toolGear/gearSsh"
	"github.com/newclarity/scribeHelpers/ux"
	"golang.org/x/crypto/ssh"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"time"
)


func (gear *DockerGear) ContainerSsh(interactive bool, statusLine bool, mountPath string, cmdArgs []string) *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		// Get Docker container SSH port.
		var clientPort string
		clientPort, gear.State = gear.Container.GetContainerSsh()
		if gear.State.IsError() {
			break
		}
		if clientPort == "" {
			gear.State.SetError("no SSH port in gear")
			break
		}

		u := url.URL{}
		err := u.UnmarshalBinary([]byte(gear.Client.DaemonHost()))
		if err != nil {
			gear.State.SetError("error finding SSH port: %s", err)
			break
		}


		// Create SSH client config.
		// fmt.Printf("Connect to %s:%s\n", u.Hostname(), port)
		gear.Ssh = gearSsh.NewSshClient(gearSsh.SshClientArgs {
			ClientAuth: &gearSsh.SshAuth {
				Host:      u.Hostname(),
				Port:      clientPort,
				Username:  gearSsh.DefaultUsername,
				Password:  gearSsh.DefaultPassword,
			},
			StatusLine: gearSsh.StatusLine {
				Enable: statusLine,
			},
			Shell:       interactive,
			GearName:    gear.Container.Name,
			GearVersion: gear.Container.Version,
			CmdArgs:     cmdArgs,
			State:       ux.NewState(gear.Runtime.CmdName, gear.Runtime.Debug),
		})


		// @TODO - Add remote host capability here!
		// Run server for SSHFS if required.
		gear.State = gear.SetMountPath(mountPath)
		if gear.State.IsOk() {
			err = gear.Ssh.InitServer()
			if err == nil {
				//noinspection GoUnhandledErrorResult
				go gear.Ssh.StartServer()

				// GEARBOX_MOUNT_HOST=10.0.5.57
				// GEARBOX_MOUNT_PATH=/Users/mick/.gearbox
				// GEARBOX_MOUNT_PORT=49410
				//time.Sleep(time.Second * 5)
				//for ; gear.Ssh.ServerAuth == nil; {
				//	time.Sleep(time.Second)
				//}

				err = os.Setenv("GEARBOX_MOUNT_HOST", gear.Ssh.ServerAuth.Host)
				err = os.Setenv("GEARBOX_MOUNT_PORT", gear.Ssh.ServerAuth.Port)
				err = os.Setenv("GEARBOX_MOUNT_USER", gear.Ssh.ServerAuth.Username)
				err = os.Setenv("GEARBOX_MOUNT_PASSWORD", gear.Ssh.ServerAuth.Password)
				err = os.Setenv("GEARBOX_MOUNT_PATH", gear.Ssh.FsMount)
			}
		}


		// Process env
		gear.State = gear.Ssh.GetEnv()
		if err != nil {
			break
		}


		// Connect to container SSH - retry 5 times.
		for i := 0; i < 5; i++ {
			gear.State.ClearError()
			err = gear.Ssh.Connect()
			if err == nil {
				break
			}

			switch v := err.(type) {
				case *ssh.ExitError:
					gear.State.SetExitCode(v.Waitmsg.ExitStatus())
					if len(cmdArgs) == 0 {
						gear.State.SetError("Command exited with error code %d", v.Waitmsg.ExitStatus())
					} else {
						gear.State.SetError("Command '%s' exited with error code %d", cmdArgs[0], v.Waitmsg.ExitStatus())
					}
					i = 5
					continue

				default:
					gear.State.SetError("SSH to Gear %s:%s failed.", gear.Container.Name, gear.Container.Version)
			}
			time.Sleep(time.Second)
		}
	}

	return gear.State
}


func (gear *DockerGear) SetMountPath(mp string) *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var err error
		var cwd string

		if mp == DefaultPathNone {
			break
		}

		switch {
			case mp == DefaultPathEmpty:
				fallthrough
			case mp == DefaultPathCwd:
				cwd, err = os.Getwd()
				if err != nil {
					gear.State.SetError(err)
					break
				}
				gear.State.SetOk()
				gear.Ssh.FsMount = cwd

			case mp == DefaultPathHome:
				var u *user.User
				u, err = user.Current()
				if err != nil {
					gear.State.SetError(err)
					break
				}
				gear.State.SetOk()
				gear.Ssh.FsMount = u.HomeDir

			default:
				mp, err = filepath.Abs(mp)
				if err != nil {
					gear.State.SetError(err)
					break
				}
				gear.State.SetOk()
				gear.Ssh.FsMount = mp
		}
	}

	return gear.State
}


func (gear *DockerGear) AddMount(local string, remote string) bool {
	if gear.Container.SshfsMounts == nil {
		gear.Container.SshfsMounts = make(SshfsMounts)
	}
	return gear.Container.SshfsMounts.Add(local, remote)
}
