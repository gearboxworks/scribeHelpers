package gearConfig

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"os/exec"
	"path/filepath"
)


func (gc *GearConfig) ListLinks(version string) *ux.State {
	if state := gc.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var err error

		err = os.Chdir(gc.Runtime.CmdDir)
		if err != nil {
			gc.State.SetError(err)
			break
		}

		isLatest := false
		if version == gc.Versions.GetLatest() {
			isLatest = true
		}

		ux.PrintfCyan("Files for Container: %s-%s\n", gc.Meta.Name, version)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Command",
			"File",
		})

		for name, fileName := range gc.Run.Commands {
			for ver := range gc.Versions {
				if ver != version {
					continue
				}

				if isLatest {
					gc.State = gc.CheckFile("latest", name, fileName)
				} else {
					gc.State = gc.CheckFile(ver, name, fileName)
				}

				if gc.State.IsError() {
					t.AppendRow([]interface{}{
						ux.SprintfRed(name),
						ux.SprintfRed("%s", gc.State.GetError()),
					})
					continue
				}

				if gc.State.IsWarning() {
					t.AppendRow([]interface{}{
						ux.SprintfYellow(name),
						ux.SprintfYellow("%s", gc.State.GetWarning()),
					})
					continue
				}

				if gc.State.IsOk() {
					t.AppendRow([]interface{}{
						ux.SprintfWhite(name),
						ux.SprintfWhite("%s", gc.State.GetOk()),
					})
					continue
				}
			}
		}

		count := t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		t.Render()
		ux.PrintflnGreen("Commands found: %d", count)
		ux.PrintflnBlue("")

		gc.State.SetOk("")
	}

	return gc.State
}


func (gc *GearConfig) CreateLinks(version string) *ux.State {
	if state := gc.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var failed bool
		var err error

		err = os.Chdir(gc.Runtime.CmdDir)
		if err != nil {
			gc.State.SetError(err)
			break
		}

		latest := gc.Versions.GetLatest()
		if version == latest {
			version = "latest"
		}

		for name, fileName := range gc.Run.Commands {
			//var err error
			//var dstFile string
			//var linkStat os.FileInfo
			//
			//if k == "default" {
			//	k = filepath.Base(v)
			//}
			//
			//if version == "latest" {
			//	dstFile, err = filepath.Abs(fmt.Sprintf("%s%c%s", gc.Runtime.CmdDir, filepath.Separator, k))
			//} else {
			//	dstFile, err = filepath.Abs(fmt.Sprintf("%s%c%s-%s", gc.Runtime.CmdDir, filepath.Separator, k, version))
			//}
			//if err != nil {
			//	continue
			//}
			//
			//linkStat, err = os.Lstat(dstFile)
			//if linkStat == nil {
			//	// Symlink doesn't exist - create.
			//	err = os.Symlink(gc.Runtime.CmdFile, dstFile)
			//	if err != nil {
			//		continue
			//	}
			//
			//	//continue
			//	linkStat, err = os.Lstat(dstFile)
			//	if linkStat == nil {
			//		continue
			//	}
			//
			//	links[k] = "linked"
			//}
			////fmt.Printf("'%s' (%s) => '%s'\n", k, dstFile, v)
			////fmt.Printf("\tReadlink() => %s\n", l)
			////fmt.Printf("\tLstat() => %s	%s	%s	%s	%d\n",
			////	linkStat.Name(),
			////	linkStat.IsDir(),
			////	linkStat.Mode().String(),
			////	linkStat.ModTime().String(),
			////	linkStat.Size(),
			////)
			//
			//// Symlink exists - validate.
			//l, _ := os.Readlink(dstFile)
			//if l == defaultBinary {
			//}
			//
			//// @TODO - Since we did the following, then gc.Runtime.CmdFile is actually file.Base() and gc.Runtime.Cmd is the full path.
			//// @TODO - os.Symlink(gc.Runtime.CmdFile, dstFile)
			//
			//fpel, err := filepath.EvalSymlinks(dstFile)
			////fmt.Printf("%s\n", fpel)
			//// @TODO - Confirm that the change from (absolute)gc.Runtime.Cmd to (relative)gc.Runtime.CmdFile works.
			////if fpel != gc.Runtime.Cmd {
			//if fpel != gc.Runtime.CmdFile {
			//	links[k] = "incorrect link"
			//	failed = true
			//}

			gc.State = gc.CheckFile(version, name, fileName)
			//gc.State.PrintResponse()
			if gc.State.GetResponseAsInt() == FileIsNotThere {
				// Symlink doesn't exist - create.
				dstFile := gc.getDstFile(version, name, fileName)
				err = os.Symlink(gc.Runtime.CmdFile, dstFile)
				if err != nil {
					failed = true
					ux.PrintflnError("Failed to create command link %s -> %s", gc.Runtime.CmdFile, dstFile)
					continue
				}

				gc.State = gc.CheckFile(version, name, fileName)
				if gc.State.IsNotOk() {
					failed = true
					gc.State.PrintResponse()
				}

				continue
			}

			if gc.State.IsNotOk() {
				failed = true
				gc.State.PrintResponse()
			}
		}

		if failed {
			gc.State.SetWarning("Failed to add all command links.")
			ux.PrintflnWarning("Failed to add all command links.")
			break
		}

		gc.State.SetOk("Created command links.")
	}

	return gc.State
}


// func (gc *GearConfig) RemoveLinks(c defaults.ExecCommand, name string, version string) *ux.State {
func (gc *GearConfig) RemoveLinks(version string) *ux.State {
	if state := gc.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var failed bool
		var err error

		err = os.Chdir(gc.Runtime.CmdDir)
		if err != nil {
			gc.State.SetError(err)
			break
		}

		latest := gc.Versions.GetLatest()
		if version == latest {
			version = "latest"
		}

		for name, fileName := range gc.Run.Commands {
			//var dstFile string
			//var linkStat os.FileInfo
			//
			//if k == "default" {
			//	k = filepath.Base(v)
			//}
			//
			//if version == "latest" {
			//	dstFile, err = filepath.Abs(fmt.Sprintf("%s%c%s", gc.Runtime.CmdDir, filepath.Separator, k))
			//} else {
			//	dstFile, err = filepath.Abs(fmt.Sprintf("%s%c%s-%s", gc.Runtime.CmdDir, filepath.Separator, k, version))
			//}
			//if err != nil {
			//	continue
			//}
			//
			//linkStat, err = os.Lstat(dstFile)
			//if err != nil {
			//	continue
			//}
			//if linkStat == nil {
			//	// Symlink doesn't exist.
			//	continue
			//}
			//
			//fpel, err := filepath.EvalSymlinks(dstFile)
			////fmt.Printf("%s\n", fpel)
			//if fpel != gc.Runtime.Cmd {
			//	links[k] = "incorrect link"
			//	failed = true
			//	continue
			//}
			//
			//l, _ := os.Readlink(dstFile)
			//if l == defaultBinary {
			//	// Symlink exists - remove.
			//	//if !filepath.IsAbs(l) {
			//	//	l, _ = filepath.Abs(fmt.Sprintf("%s%c%s", c.Dir, filepath.Separator, l))
			//	//}
			//	err = os.Remove(dstFile)
			//}

			gc.State = gc.CheckFile(version, name, fileName)
			if gc.State.GetResponseAsInt() == FileIsUnknown {
				continue
			}
			if gc.State.GetResponseAsInt() == FileIsNotThere {
				continue
			}
			if gc.State.GetResponseAsInt() == FileIsFile {
				ux.PrintflnWarning("Not removing: %s", gc.State.GetError())
				continue
			}
			if gc.State.GetResponseAsInt() == FileIsSymlink {
				ux.PrintflnWarning("Not removing: %s", gc.State.GetError())
				continue
			}

			dstFile := gc.getDstFile(version, name, fileName)
			err = os.Remove(dstFile)
			if err != nil {
				failed = true
				ux.PrintflnError("Failed to remove command link %s -> %s", gc.Runtime.CmdFile, dstFile)
			}
		}

		if failed {
			gc.State.SetWarning("Failed to remove all command links.")
			ux.PrintflnWarning("Failed to remove all command links.")
			break
		}

		gc.State.SetOk("Removed command links.")
	}

	return gc.State
}


const (
	FileIsUnknown		= iota	// Gear file/symlink is unknown.
	FileIsNotThere				// Gear file/symlink doesn't exist - needs to be created.
	FileIsLaunchManaged 		// Gear file/symlink is present and managed by launch.
	FileIsNotInPath				// Gear file/symlink is present and managed by launch, but is not in PATH.
	FileIsNotFirst				// Gear file/symlink is present and managed by launch, but is not FIRST in PATH.
	FileIsFile					// Gear file/symlink is present and a file not managed by launch.
	FileIsSymlink				// Gear file/symlink is present and a symlink not managed by launch.
)
func (gc *GearConfig) getDstFile(version string, name string, fileName string) string {
	var dstFile string

	for range onlyOnce {
		if name == "default" {
			name = filepath.Base(fileName)
		}

		if version == "latest" {
			dstFile, _ = filepath.Abs(fmt.Sprintf("%s%c%s", gc.Runtime.CmdDir, filepath.Separator, name))
		} else {
			dstFile, _ = filepath.Abs(fmt.Sprintf("%s%c%s-%s", gc.Runtime.CmdDir, filepath.Separator, name, version))
		}
	}

	return dstFile
}
func (gc *GearConfig) CheckFile(version string, name string, fileName string) *ux.State {

	for range onlyOnce {
		var err error
		var dstFile string
		var linkStat os.FileInfo

		dstFile = gc.getDstFile(version, name, fileName)
		//if name == "default" {
		//	name = filepath.Base(fileName)
		//}
		//
		//if version == "latest" {
		//	dstFile, err = filepath.Abs(fmt.Sprintf("%s%c%s", gc.Runtime.CmdDir, filepath.Separator, name))
		//} else {
		//	dstFile, err = filepath.Abs(fmt.Sprintf("%s%c%s-%s", gc.Runtime.CmdDir, filepath.Separator, name, version))
		//}
		//if err != nil {
		//	continue
		//}

		linkStat, err = os.Lstat(dstFile)
		if linkStat == nil {
			// Symlink doesn't exist.
			gc.State.SetResponse(FileIsNotThere)
			gc.State.SetWarning("%s is missing.", dstFile)
			break
		}

		var lp string
		var sp string
		var fp string

		// Check for PATH.
		lp, err = exec.LookPath(filepath.Base(dstFile))

		// Check for absolute symlink.
		sp, err = os.Readlink(dstFile)

		// Check for relative symlink.
		fp, err = filepath.EvalSymlinks(dstFile)

		// Check for absolute symlink.
		if sp == defaultBinary {
			if lp == "" {
				gc.State.SetResponse(FileIsNotInPath)
				gc.State.SetWarning("%s is managed, but not in PATH.", dstFile)
				break
			}
			if lp != dstFile {
				gc.State.SetResponse(FileIsNotFirst)
				gc.State.SetWarning("%s is managed, but not FIRST in PATH -> %s.", dstFile, lp)
				break
			}

			gc.State.SetResponse(FileIsLaunchManaged)
			gc.State.SetOk("%s is managed.", dstFile)
			break
		}
		if sp == "" {
			gc.State.SetResponse(FileIsFile)
			gc.State.SetError("%s is a non-managed file.", dstFile)
			break
		}

		// Check for relative symlink.
		if fp == gc.Runtime.Cmd {
			if lp == "" {
				gc.State.SetResponse(FileIsNotInPath)
				gc.State.SetWarning("%s is managed, but not in PATH.", dstFile)
				break
			}
			if lp != dstFile {
				gc.State.SetResponse(FileIsNotFirst)
				gc.State.SetWarning("%s is managed, but not FIRST in PATH -> %s.", dstFile, lp)
				break
			}

			gc.State.SetResponse(FileIsLaunchManaged)
			gc.State.SetOk("%s is managed.", dstFile)
			break
		}

		gc.State.SetResponse(FileIsSymlink)
		gc.State.SetError("%s is a non-managed symlink -> %s.", dstFile, sp)

		if err != nil {

		}
	}

	return gc.State
}

//func (gc *GearConfig) ListFile(version string, name string, fileName string) *ux.State {
//	if state := gc.IsNil(); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		err := os.Chdir(gc.Runtime.CmdDir)
//		if err != nil {
//			gc.State.SetError(err)
//			break
//		}
//
//		ux.PrintfCyan("Files for Container: %s-%s\n", gc.Meta.Name, version)
//		t := table.NewWriter()
//		t.SetOutputMirror(os.Stdout)
//		//t.AppendHeader(table.Row{
//		//	"Command",
//		//	"File",
//		//})
//
//		for k, v := range gc.Run.Commands {
//			for ver := range gc.Versions {
//				if ver != version {
//					continue
//				}
//
//				gc.State = gc.CheckFile(ver, k, v)
//
//				if gc.State.IsError() {
//					t.AppendRow([]interface{}{
//						ux.SprintfRed(k),
//						ux.SprintfRed("%s", gc.State.GetError()),
//					})
//					continue
//				}
//
//				if gc.State.IsWarning() {
//					t.AppendRow([]interface{}{
//						ux.SprintfYellow(k),
//						ux.SprintfYellow("%s", gc.State.GetWarning()),
//					})
//					continue
//				}
//
//				if gc.State.IsOk() {
//					t.AppendRow([]interface{}{
//						ux.SprintfWhite(k),
//						ux.SprintfWhite("%s", gc.State.GetOk()),
//					})
//					continue
//				}
//			}
//		}
//
//		count := t.Length()
//		if count == 0 {
//			ux.PrintfYellow("None found\n")
//			break
//		}
//
//		t.Render()
//		ux.PrintflnGreen("Commands found: %d", count)
//		ux.PrintflnBlue("")
//
//		gc.State.SetOk("")
//	}
//
//	return gc.State
//}
