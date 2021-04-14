package toolCopy

import (
	"github.com/gearboxworks/scribeHelpers/ux"
)


func (c *TypeOsCopy) Copy(args ...string) *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		c.State = c.Paths.Source.StatPath()

		if !c.Paths.Source.Exists() {
			c.State.SetError("src path not found")
			break
		}

		if c.State.IsNotOk() {
			c.State.SetError("src path not found")
			break
		}

		for range onlyOnce {
			c.State = c.Paths.Destination.StatPath()
			if c.Paths.Destination.NotExists() {
				c.State.SetOk()	// Destination path can actually not exist.
				break
			}
			if c.Paths.Destination.CanOverwrite() {
				break
			}
			c.State.SetError("cannot overwrite destination")
		}

		if c.State.IsError() {
			break
		}

		// @TODO Should be a better way.
		method := c.Method.GetSelected()
		name := c.Method.GetName()
		switch name {
			case ConstMethodRsync:
				c.State = method.Run(&c.Paths, args...)
		}
	}

	return c.State
}
