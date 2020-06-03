package toolCopy

import (
	"github.com/newclarity/scribeHelpers/ux"
)


func (c *TypeOsCopy) Copy() *ux.State {
	if state := c.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		if !c.Source.Exists() {
			c.State.SetError("src path not found")
			break
		}

		for range OnlyOnce {
			if c.Destination.NotExists() {
				break
			}
			if c.Destination.CanOverwrite() {
				break
			}
			c.State.SetError("cannot overwrite destination")
		}

		if c.State.IsError() {
			break
		}

		// @TODO - do copying of files here

		c.State.SetOk("chdir OK")
	}

	return c.State
}
