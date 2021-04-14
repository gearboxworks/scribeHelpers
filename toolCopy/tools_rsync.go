package toolCopy

import (
	"github.com/gearboxworks/scribeHelpers/toolPath"
	"github.com/gearboxworks/scribeHelpers/ux"
	"github.com/zloylos/grsync"
	"time"
)


// Usage:
//		{{ $return := WriteFile "filename.txt" .data.Source 0644 }}
func ToolCopyRsync(src interface{}, dest interface{}, exclude ...interface{}) *ux.State {
	c := New(nil)

	for range onlyOnce {
		s := toolPath.ReflectAbsPath(src)
		if s == nil {
			break
		}
		if !c.Paths.Source.SetPath(*s) {
			break
		}
		c.State = c.Paths.Source.StatPath()
		if !c.Paths.Source.Exists() {
			//c.State.SetError("src path not found")
			break
		}


		c.Paths.Destination.SetOverwriteable()

		d := toolPath.ReflectAbsPath(dest)
		if d == nil {
			break
		}
		if !c.Paths.Destination.SetPath(*d) {
			break
		}
		for range onlyOnce {
			c.State = c.Paths.Destination.StatPath()
			if c.Paths.Destination.NotExists() {
				c.State.Clear()
				break
			}
			if c.Paths.Destination.CanOverwrite() {
				break
			}
			c.State.SetError("cannot overwrite destination '%s'", c.Paths.Destination.GetPath())
		}
		if c.State.IsError() {
			break
		}


		if !c.Method.SelectMethod(ConstMethodRsync) {
			c.State.SetError("rsync method unavailable")
			break
		}


		task := grsync.NewTask(
			c.Paths.Source.GetPath(),
			c.Paths.Destination.GetPath(),
			c.Method.Selected.Options.(grsync.RsyncOptions),
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
		c.State.SetOutput(task.Log().Stdout)
		//l := task.Log().Stdout
		//fmt.Print("%s\n", l)
		c.State.SetError(err)
		if c.State.IsError() {
			break
		}


		//opts := []string{}
		////opts = append(opts, c.RsyncOptions...)
		//opts = append(opts, c.goFiles.Source.GetPath())
		//opts = append(opts, c.goFiles.Destination.GetPath())
		//cmd := exec.Command("rsync", opts...)
		//out, err := cmd.CombinedOutput()
		//c.State.SetOutput(out)
		//c.State.SetError(err)
		//
		//if c.State.IsError() {
		//	if exitError, ok := err.(*exec.ExitError); ok {
		//		waitStatus := exitError.Sys().(syscall.WaitStatus)
		//		c.State.ExitCode = waitStatus.ExitStatus()
		//	}
		//
		//	//fmt.Printf("%s\n", ret.PrintError())
		//	break
		//}
		//waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
		//c.State.ExitCode = waitStatus.ExitStatus()

		c.State.SetOk("%s", c.State.Output)
	}

	return c.State
}


//func ToolRsync(src interface{}, dest interface{}, options interface{}, exclude ...interface{}) *ToolOsCopy {
//	ret := NewOsCopy()
//
//	for range onlyOnce {
//		s := toolTypes.ReflectString(src)
//		if s == nil {
//			ret.State.SetError("rsync source empty")
//			break
//		}
//		if ret.Source.SetPath(*s) {
//			ret.State.SetError("rsync source empty")
//		}
//
//
//		d := toolTypes.ReflectString(dest)
//		if d == nil {
//			ret.State.SetError("rsync destination empty")
//			break
//		}
//		if ret.Source.SetPath(*s) {
//			ret.State.SetError("rsync destination empty")
//		}
//
//
//		o := toolTypes.ReflectString(options)
//		switch {
//			case o == nil:
//				fallthrough
//			case *o == "":
//				ret.RsyncOptions = []string{"-HvaxPn"}
//			default:
//				ret.RsyncOptions = []string{*o}
//		}
//
//		e := toolTypes.ReflectStrings(exclude...)
//		if e == nil {
//			break
//		}
//		ret.ExcludeFiles = *e
//
//		for _, es := range ret.ExcludeFiles {
//			ret.RsyncOptions = append(ret.RsyncOptions, fmt.Sprintf("--exclude='%s'", es))
//		}
//	}
//
//	return (*ToolOsCopy)(ret)
//}
