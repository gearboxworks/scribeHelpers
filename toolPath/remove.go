package toolPath

import (
	"github.com/gearboxworks/scribeHelpers/toolPrompt"
	"github.com/gearboxworks/scribeHelpers/ux"
	"os"
)


func (p *TypeOsPath) Remove() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.Clear()

		if !p.IsValid() {
			break
		}

		for range onlyOnce {
			p.StatPath()
			if !p._Exists {
				p.State.SetWarning("path '%s' doesn't exist", p.Path)
				break
			}
			if p._CanRemove {
				break
			}

			if !toolPrompt.ToolUserPromptBool("Remove path '%s'? (Y|N) ", p.Path) {
				p.State.SetWarning("not removing path '%s'", p.Path)
				break
			}
			p.State.Clear()
		}
		if p.State.IsNotOk() {
			break
		}

		err := os.Remove(p.Path)
		if err != nil {
			p.State.SetError(err)
			break
		}

		p.State.SetOk("path '%s' removed OK", p.Path)
	}

	return p.State
}


func (p *TypeOsPath) RemoveFile() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.Clear()

		if !p.IsValid() {
			break
		}

		for range onlyOnce {
			p.StatPath()
			if p._IsDir {
				p.State.SetError("path is a directory")
				break
			}
			if !p._Exists {
				p.State.SetWarning("file '%s' doesn't exist", p.Path)
				break
			}
			if p._CanRemove {
				break
			}

			p.State.Clear()
			if !toolPrompt.ToolUserPromptBool("Remove file '%s'? (Y|N) ", p.Path) {
				p.State.SetWarning("not removing file '%s'", p.Path)
				break
			}
		}
		if p.State.IsNotOk() {
			break
		}

		err := os.Remove(p.Path)
		if err != nil {
			p.State.SetError(err)
			break
		}

		p.State.SetOk("file '%s' removed OK", p.Path)
	}

	return p.State
}


func (p *TypeOsPath) RemoveDir() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.Clear()

		if !p.IsValid() {
			break
		}

		for range onlyOnce {
			p.StatPath()
			if p._IsDir {
				p.State.SetError("path '%s' is a directory", p.Path)
				break
			}
			if !p._Exists {
				p.State.SetWarning("directory '%s' doesn't exist", p.Path)
				break
			}
			if p._CanRemove {
				break
			}

			if !toolPrompt.ToolUserPromptBool("Remove directory '%s'? (Y|N) ", p.Path) {
				p.State.SetWarning("not removing file '%s'", p.Path)
				break
			}
			p.State.Clear()
		}
		if p.State.IsError() {
			break
		}

		err := os.Remove(p.Path)
		if err != nil {
			p.State.SetError(err)
			break
		}

		p.State.SetOk("file '%s' removed OK", p.Path)
	}

	return p.State
}
