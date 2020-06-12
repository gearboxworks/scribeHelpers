package toolCopy

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/zloylos/grsync"
	"path/filepath"
	"time"
)

type MethodFunc interface {
	_CopyRunDefault(ref *TypeCopyMethod, paths *TypeOsCopyPaths, args ...string) *ux.State
}

func _CopyRunDefault(ref *TypeCopyMethod, paths *TypeOsCopyPaths, args ...string) *ux.State {
	return _CopyRunRsync(ref, paths, args...)
}
func _CopyRunRsync(ref *TypeCopyMethod, paths *TypeOsCopyPaths, args ...string) *ux.State {
	if state := ux.IfNilReturnError(ref); state.IsError() {
		return state
	}

	for range onlyOnce {
		if paths == nil {
			ref.state.SetError("No paths defined.")
			break
		}

		ref.state = paths.Source.StatPath()
		if !paths.Source.Exists() {
			//ref.state.SetError("src path not found")
			break
		}

		paths.Destination.SetOverwriteable()

		for range onlyOnce {
			ref.state = paths.Destination.StatPath()
			if paths.Destination.NotExists() {
				ref.state.SetOk()
				break
			}
			if paths.Destination.IsOverwriteable() {
				break
			}
			ref.state.SetError("cannot overwrite destination '%s'", paths.Destination.GetPath())
		}
		if ref.state.IsError() {
			break
		}

		// Adjust src and dst paths.
		src := ""
		dst := ""
		if paths.Source.IsADir() {
			src = fmt.Sprintf("%s%c", paths.Source.GetPathAbs(), filepath.Separator)
			dst = fmt.Sprintf("%s%c", paths.Destination.GetPathAbs(), filepath.Separator)
		} else {
			src = paths.Source.GetPathAbs()
			dst = fmt.Sprintf("%s%c", paths.Destination.GetPathAbs(), filepath.Separator)
		}

		task := grsync.NewTask(
			src,
			dst,
			ref.Options.(grsync.RsyncOptions),
		)

		loop := true
		go func() {
			ux.Printf("\n")
			for ;loop; {
				state := task.State()
				ux.PrintfGreen(
					"Copy progress: %.2f / rem. %d / tot. %d / sp. %s \n",
					state.Progress,
					state.Remain,
					state.Total,
					state.Speed,
				)
				time.Sleep(time.Second)
			}
			ux.Printf("\n")
		}()

		err := task.Run()
		loop = false
		ref.state.SetOutput(task.Log().Stdout)
		//l := task.Log().Stdout
		//fmt.Print("%s\n", l)
		ref.state.SetError(err)
		if ref.state.IsError() {
			break
		}


		//opts := []string{}
		////opts = append(opts, paths.RsyncOptions...)
		//opts = append(opts, paths.Source.GetPath())
		//opts = append(opts, paths.Destination.GetPath())
		//cmd := exepaths.Command("rsync", opts...)
		//out, err := cmd.CombinedOutput()
		//ref.state.SetOutput(out)
		//ref.state.SetError(err)
		//
		//if ref.state.IsError() {
		//	if exitError, ok := err.(*exepaths.ExitError); ok {
		//		waitStatus := exitError.Sys().(syscall.WaitStatus)
		//		ref.state.ExitCode = waitStatus.ExitStatus()
		//	}
		//
		//	//fmt.Printf("%s\n", ret.PrintError())
		//	break
		//}
		//waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
		//ref.state.ExitCode = waitStatus.ExitStatus()

		ref.state.SetOk("Files copied with %s OK", ref.Name)
	}

	return ref.state
}
func _CopyRunTar(ref *TypeCopyMethod, paths *TypeOsCopyPaths, args ...string) *ux.State {
	if state := ux.IfNilReturnError(ref); state.IsError() {
		return state
	}

	return ref.state
}
func _CopyRunCpio(ref *TypeCopyMethod, paths *TypeOsCopyPaths, args ...string) *ux.State {
	if state := ux.IfNilReturnError(ref); state.IsError() {
		return state
	}

	return ref.state
}
func _CopyRunSftp(ref *TypeCopyMethod, paths *TypeOsCopyPaths, args ...string) *ux.State {
	if state := ux.IfNilReturnError(ref); state.IsError() {
		return state
	}

	return ref.state
}
func _CopyRunCp(ref *TypeCopyMethod, paths *TypeOsCopyPaths, args ...string) *ux.State {
	if state := ux.IfNilReturnError(ref); state.IsError() {
		return state
	}

	for range onlyOnce {
		if paths == nil {
			ref.state.SetError("No paths defined.")
			break
		}

		ref.state = paths.Source.StatPath()
		if !paths.Source.Exists() {
			//ref.state.SetError("src path not found")
			break
		}

		paths.Destination.SetOverwriteable()

		for range onlyOnce {
			ref.state = paths.Destination.StatPath()
			if paths.Destination.NotExists() {
				ref.state.SetOk()
				break
			}
			if paths.Destination.CanOverwrite() {
				break
			}
			ref.state.SetError("cannot overwrite destination '%s'", paths.Destination.GetPath())
		}
		if ref.state.IsError() {
			break
		}

		// Adjust src and dst paths.
		src := ""
		dst := ""
		if paths.Source.IsADir() {
			src = fmt.Sprintf("%s%c", paths.Source.GetPathAbs(), filepath.Separator)
			dst = fmt.Sprintf("%s%c", paths.Destination.GetPathAbs(), filepath.Separator)
		}

		task := grsync.NewTask(
			src,
			dst,
			ref.Options.(grsync.RsyncOptions),
		)

		loop := true
		go func() {
			ux.Printf("\n")
			for ;loop; {
				state := task.State()
				ux.PrintfGreen(
					"Copy progress: %.2f / rem. %d / tot. %d / sp. %s \n",
					state.Progress,
					state.Remain,
					state.Total,
					state.Speed,
				)
				time.Sleep(time.Second)
			}
			ux.Printf("\n")
		}()

		err := task.Run()
		loop = false
		ref.state.SetOutput(task.Log().Stdout)
		//l := task.Log().Stdout
		//fmt.Print("%s\n", l)
		ref.state.SetError(err)
		if ref.state.IsError() {
			break
		}


		//opts := []string{}
		////opts = append(opts, paths.RsyncOptions...)
		//opts = append(opts, paths.Source.GetPath())
		//opts = append(opts, paths.Destination.GetPath())
		//cmd := exepaths.Command("rsync", opts...)
		//out, err := cmd.CombinedOutput()
		//ref.state.SetOutput(out)
		//ref.state.SetError(err)
		//
		//if ref.state.IsError() {
		//	if exitError, ok := err.(*exepaths.ExitError); ok {
		//		waitStatus := exitError.Sys().(syscall.WaitStatus)
		//		ref.state.ExitCode = waitStatus.ExitStatus()
		//	}
		//
		//	//fmt.Printf("%s\n", ret.PrintError())
		//	break
		//}
		//waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
		//ref.state.ExitCode = waitStatus.ExitStatus()

		ref.state.SetOk("Files copied with %s OK", ref.Name)
	}

	return ref.state
}
