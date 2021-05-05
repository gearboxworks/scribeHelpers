package toolGear

import (
	"fmt"
	"github.com/gearboxworks/scribeHelpers/toolGear/gearSsh"
	"github.com/gearboxworks/scribeHelpers/ux"
	"golang.org/x/crypto/ssh"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)


func (gear *Gear) ContainerSsh(interactive bool, statusLine bool, mountPath string, cmdArgs []string) *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	return gear.Container.ContainerSsh(interactive, statusLine, mountPath, cmdArgs)
}

func (gear *Gear) SetMountPath(mp string) *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	return gear.Container.SetMountPath(mp)
}

func (gear *Gear) AddMount(local string, remote string) bool {
	if state := gear.IsNil(); state.IsError() {
		return false
	}

	return gear.Container.AddMount(local, remote)
}

func (gear *Gear) SetSshStatusLine(s bool) {
	if state := gear.IsNil(); state.IsError() {
		return
	}

	gear.Container.SetSshStatusLine(s)
}

func (gear *Gear) SetSshShell(s bool) {
	if state := gear.IsNil(); state.IsError() {
		return
	}

	gear.Container.SetSshShell(s)
}

//func (gear *Gear) SetDebug(s bool) {
//	if state := gear.IsNil(); state.IsError() {
//		return
//	}
//
//	//gear.Container.Ssh.Debug = s
//	//gear.Container.SetDebug(s)
//}


func (c *Container) ContainerSsh(interactive bool, statusLine bool, mountPath string, cmdArgs []string) *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		// Get Docker container SSH port.
		var clientPort string
		clientPort, c.State = c.GetContainerSsh()
		if c.State.IsError() {
			break
		}
		if clientPort == "" {
			c.State.SetError("no SSH port in gear")
			break
		}

		u := url.URL{}
		err := u.UnmarshalBinary([]byte(c.Docker.Client.DaemonHost()))
		if err != nil {
			c.State.SetError("error finding SSH port: %s", err)
			break
		}


		// Create SSH client config.
		// fmt.Printf("Connect to %s:%s\n", u.Hostname(), port)
		c.Ssh = gearSsh.NewSshClient(gearSsh.SshClientArgs {
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
			GearName:    c.Name,
			GearVersion: c.Version,
			CmdArgs:     cmdArgs,
			FsRemote:    c.Docker.Provider.Remote,

			Debug:       c.runtime.Debug,
			State:       ux.NewState(c.runtime.CmdName, c.runtime.Debug),
		})


		// @TODO - Add remote host capability here!
		// Run server for SSHFS if required.
		c.State = c.SetMountPath(mountPath)
		if c.State.IsOk() {
			err = c.Ssh.InitServer()
			if err == nil {
				//noinspection GoUnhandledErrorResult
				go c.Ssh.StartServer()

				// GEARBOX_MOUNT_HOST=10.0.5.57
				// GEARBOX_MOUNT_PATH=/Users/mick/.gearbox
				// GEARBOX_MOUNT_PORT=49410
				//time.Sleep(time.Second * 5)
				//for ; gear.Ssh.ServerAuth == nil; {
				//	time.Sleep(time.Second)
				//}

				err = os.Setenv("GEARBOX_MOUNT_HOST", c.Ssh.ServerAuth.Host)
				err = os.Setenv("GEARBOX_MOUNT_PORT", c.Ssh.ServerAuth.Port)
				err = os.Setenv("GEARBOX_MOUNT_USER", c.Ssh.ServerAuth.Username)
				err = os.Setenv("GEARBOX_MOUNT_PASSWORD", c.Ssh.ServerAuth.Password)
				err = os.Setenv("GEARBOX_MOUNT_PATH", c.Ssh.FsMount)
			}
		}


		// Process env
		c.State = c.Ssh.GetEnv()
		if err != nil {
			break
		}

		if c.runtime.Debug {
			fmt.Printf("DEBUG: c.runtime.Args == %s\r\n", strings.Join(c.runtime.Args, " "))
			fmt.Printf("DEBUG: c.runtime.FullArgs == %s\r\n", strings.Join(c.runtime.FullArgs, " "))
		}


		// Connect to container SSH - retry 5 times.
		for i := 0; i < 5; i++ {
			c.State.ClearError()
			err = c.Ssh.Connect()
			if err == nil {
				break
			}

			switch v := err.(type) {
				case *ssh.ExitError:
					c.State.SetExitCode(v.Waitmsg.ExitStatus())
					if len(cmdArgs) == 0 {
						c.State.SetError("Command exited with error code %d", v.Waitmsg.ExitStatus())
					} else {
						c.State.SetError("Command '%s' exited with error code %d", cmdArgs[0], v.Waitmsg.ExitStatus())
					}
					i = 5
					continue

				default:
					c.State.SetError("SSH to Gear %s:%s failed.", c.Name, c.Version)
			}
			time.Sleep(time.Second)
		}
	}

	return c.State
}

func (c *Container) SetMountPath(mp string) *ux.State {
	if state := c.IsNil(); state.IsError() {
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
					c.State.SetError(err)
					break
				}
				c.State.SetOk()
				c.Ssh.FsMount = cwd

			case mp == DefaultPathHome:
				var u *user.User
				u, err = user.Current()
				if err != nil {
					c.State.SetError(err)
					break
				}
				c.State.SetOk()
				c.Ssh.FsMount = u.HomeDir

			default:
				mp, err = filepath.Abs(mp)
				if err != nil {
					c.State.SetError(err)
					break
				}
				c.State.SetOk()
				c.Ssh.FsMount = mp
		}

		c.Ssh.FsRemote = c.Docker.Provider.Remote
		if c.Docker.Provider.IsRemote() {
		}
	}

	return c.State
}

func (c *Container) AddMount(local string, remote string) bool {
	if c.SshfsMounts == nil {
		c.SshfsMounts = make(SshfsMounts)
	}
	return c.SshfsMounts.Add(local, remote)
}

func (c *Container) SetSshStatusLine(s bool) {
	c.Ssh.StatusLine.Enable = s
}

func (c *Container) SetSshShell(s bool) {
	c.Ssh.Shell = s
}
