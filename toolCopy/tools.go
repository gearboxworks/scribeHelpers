// High level helper functions available within templates - file copy.
package toolCopy

import (
	"fmt"
	"github.com/gearboxworks/scribeHelpers/toolTypes"
	"github.com/gearboxworks/scribeHelpers/ux"
	"os/exec"
	"strings"
	"syscall"
)


// Alias of Rsync || Tar || whatever - basically determine what tool to use based on availability.
// @TODO - To be implemented.
// Usage:
//		{{ $copy := CopyFiles }}
func ToolCopyFiles() *ToolOsCopy {
	ret := New(nil)

	for range onlyOnce {
		ret.State.Clear()
	}

	return (*ToolOsCopy)(ret)
}


// Usage:
//		{{ $copy := CopyFiles }}
//		{{ $state := SetSourcePath "filename.txt" }}
func (c *ToolOsCopy) SetSourcePath(src ...interface{}) *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}
	c.State.SetFunction()

	for range onlyOnce {
		p := toolTypes.ReflectStrings(src...)
		if p == nil {
			c.State.SetError("%s source empty", c.Method.GetName())
			break
		}
		if !c.Paths.Source.SetPath(*p...) {
			c.State.SetError("%s source empty", c.Method.GetName())
			break
		}
		c.State.Clear()
	}

	return c.State
}
func (c *ToolOsCopy) SetSource(dest ...interface{}) *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}
	c.State.SetFunction()
	return c.SetSourcePath(dest...)
}


// Usage:
//		{{ $copy := CopyFiles }}
//		{{ $state := SetDestinationPath "filename.txt" }}
func (c *ToolOsCopy) SetDestinationPath(dest ...interface{}) *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}
	c.State.SetFunction()

	for range onlyOnce {
		p := toolTypes.ReflectStrings(dest...)
		if p == nil {
			c.State.SetError("%s destination empty", c.Method.GetName())
			break
		}
		if !c.Paths.Destination.SetPath(*p...) {
			c.State.SetError("%s destination empty", c.Method.GetName())
			break
		}
		c.State.Clear()
	}

	return c.State
}
func (c *ToolOsCopy) SetTarget(dest ...interface{}) *ux.State {
	return c.SetDestinationPath(dest...)
}


// Usage:
//		{{ $copy := CopyFiles }}
//		{{ $state := SetSourcePath "filename.txt" }}
func (c *ToolOsCopy) SetExcludePaths(exclude ...interface{}) *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}
	c.State.SetFunction()

	for range onlyOnce {
		e := toolTypes.ReflectStrings(exclude...)
		if e == nil {
			break
		}
		if !c.Paths.Exclude.SetPaths(*e...) {
			// Do nothing. Allow empty exclude paths.
		}
		c.State.Clear()
	}

	return c.State
}


// Usage:
//		{{ $copy := CopyFiles }}
//		{{ $state := SetSourcePath "filename.txt" }}
func (c *ToolOsCopy) SetIncludePaths(include ...interface{}) *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}
	c.State.SetFunction()

	for range onlyOnce {
		i := toolTypes.ReflectStrings(include...)
		if i == nil {
			break
		}
		if !c.Paths.Include.SetPaths(*i...) {
			// Do nothing. Allow empty exclude paths.
		}
		c.State.Clear()
	}

	return c.State
}


// Usage:
//		{{ $return := WriteFile "filename.txt" .data.Source 0644 }}
func (c *ToolOsCopy) Run() *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}
	c.State.SetFunction()

	for range onlyOnce {
		c.State = c.Paths.Source.StatPath()
		if c.State.IsError() {
			break
		}

		opts := []string{}
		//opts = append(opts, c.RsyncOptions...)
		opts = append(opts, c.Paths.Source.GetPath())
		opts = append(opts, c.Paths.Destination.GetPath())

		cmd := exec.Command("rsync", opts...)

		out, err := cmd.CombinedOutput()
		c.State.SetOutput(out)
		c.State.SetError(err)

		if c.State.IsError() {
			if exitError, ok := err.(*exec.ExitError); ok {
				waitStatus := exitError.Sys().(syscall.WaitStatus)
				c.State.ExitCode = waitStatus.ExitStatus()
			}

			//fmt.Printf("%s\n", ret.PrintError())
			break
		}

		waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
		c.State.ExitCode = waitStatus.ExitStatus()

		fmt.Printf("\nrsync %s\n", strings.Join(opts, " "))
		c.State.SetOk("%s", c.State.Output)
	}

	return c.State
}


//// Usage:
////		{{ $copy := CopyFiles }}
////		{{ $state := SetSourcePath "filename.txt" }}
//func (c *ToolOsCopy) SetOptions(src interface{}) *ux.State {
//	for range onlyOnce {
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
//	return (*State)(c.State)
//}
//
