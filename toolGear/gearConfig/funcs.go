package gearConfig

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"path/filepath"
)


func (gc *GearConfig) GetName() string {
	return gc.Meta.Name
}


func (gc *GearConfig) GetCommand(cmd []string) []string {
	var retCmd []string

	for range onlyOnce {
		var cmdExec string
		switch {
		case len(cmd) == 0:
			cmdExec = DefaultCommandName

		case cmd[0] == "":
			cmdExec = DefaultCommandName

		case cmd[0] == gc.Meta.Name:
			cmdExec = DefaultCommandName

		case cmd[0] != "":
			cmdExec = cmd[0]

		default:
			//cmdExec = cmd[0]
			cmdExec = DefaultCommandName
		}

		c := gc.MatchCommand(cmdExec)
		if c == nil {
			retCmd = []string{}
			break
		}

		retCmd = append([]string{*c}, cmd[1:]...)
	}

	return retCmd
}


func (gc *GearConfig) MatchCommand(cmd string) *string {
	var c *string

	for range onlyOnce {
		if c2, ok := gc.Run.Commands[cmd]; ok {
			c = &c2
			break
		}
	}

	return c
}


func (gc *GearConfig) CreateLinks(version string) *ux.State {
	if state := gc.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		links := make(map[string]string)
		var failed bool

		for k, v := range gc.Run.Commands {
			var err error
			var dstFile string
			var linkStat os.FileInfo

			if k == "default" {
				k = filepath.Base(v)
			}

			if version == "latest" {
				dstFile, err = filepath.Abs(fmt.Sprintf("%s%c%s", gc.Runtime.CmdDir, filepath.Separator, k))
			} else {
				dstFile, err = filepath.Abs(fmt.Sprintf("%s%c%s-%s", gc.Runtime.CmdDir, filepath.Separator, k, version))
			}
			if err != nil {
				continue
			}

			linkStat, err = os.Lstat(dstFile)
			if linkStat == nil {
				// Symlink doesn't exist - create.
				err = os.Symlink(gc.Runtime.CmdFile, dstFile)
				if err != nil {
					continue
				}

				//continue
				linkStat, err = os.Lstat(dstFile)
				if linkStat == nil {
					continue
				}

				links[k] = "linked"
			}
			//fmt.Printf("'%s' (%s) => '%s'\n", k, dstFile, v)
			//fmt.Printf("\tReadlink() => %s\n", l)
			//fmt.Printf("\tLstat() => %s	%s	%s	%s	%d\n",
			//	linkStat.Name(),
			//	linkStat.IsDir(),
			//	linkStat.Mode().String(),
			//	linkStat.ModTime().String(),
			//	linkStat.Size(),
			//)

			// Symlink exists - validate.
			l, _ := os.Readlink(dstFile)
			if l == defaultBinary {
			}

			fpel, err := filepath.EvalSymlinks(dstFile)
			//fmt.Printf("%s\n", fpel)
			if fpel != gc.Runtime.Cmd {
				links[k] = "incorrect link"
				failed = true
			}
		}

		if failed {
			gc.State.SetWarning("Failed to add all application links.")
			for k, v := range links {
				if v == "linked" {
					continue
				}
				ux.PrintflnWarning("%s - %s", k, v)
			}
			break
		}

		gc.State.SetOk("Created application links.")
	}

	return gc.State
}


// func (gc *GearConfig) RemoveLinks(c defaults.ExecCommand, name string, version string) *ux.State {
func (gc *GearConfig) RemoveLinks(version string) *ux.State {
	if state := gc.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		links := make(map[string]string)
		var failed bool

		for k, v := range gc.Run.Commands {
			var err error
			var dstFile string
			var linkStat os.FileInfo

			if k == "default" {
				k = filepath.Base(v)
			}

			if version == "latest" {
				dstFile, err = filepath.Abs(fmt.Sprintf("%s%c%s", gc.Runtime.CmdDir, filepath.Separator, k))
			} else {
				dstFile, err = filepath.Abs(fmt.Sprintf("%s%c%s-%s", gc.Runtime.CmdDir, filepath.Separator, k, version))
			}
			if err != nil {
				continue
			}

			linkStat, err = os.Lstat(dstFile)
			if err != nil {
				continue
			}
			if linkStat == nil {
				// Symlink doesn't exist.
				continue
			}

			fpel, err := filepath.EvalSymlinks(dstFile)
			//fmt.Printf("%s\n", fpel)
			if fpel != gc.Runtime.Cmd {
				links[k] = "incorrect link"
				failed = true
				continue
			}

			l, _ := os.Readlink(dstFile)
			if l == defaultBinary {
				// Symlink exists - remove.
				//if !filepath.IsAbs(l) {
				//	l, _ = filepath.Abs(fmt.Sprintf("%s%c%s", c.Dir, filepath.Separator, l))
				//}
				err = os.Remove(dstFile)
			}
		}
		//gc.State.SetDebug("DEBUGIT")

		if failed {
			gc.State.SetWarning("Failed to remove all application links.")
			for k, v := range links {
				ux.PrintflnWarning("%s - %s", k, v)
			}
			break
		}

		gc.State.SetOk("Removed application links.")
	}

	return gc.State
}
