package toolPath

import (
	"github.com/newclarity/scribeHelpers/ux"
	"os"
)


func (p *TypeOsPath) Chdir() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.Clear()

		if !p.IsValid() {
			break
		}

		p.State = p.StatPath()
		if p.State.IsError() {
			break
		}
		if !p._Exists {
			p.State.SetError("directory not found")
			break
		}
		if p._IsFile {
			err := os.Chdir(p._Dirname)
			p.State.SetError(err)
			if p.State.IsError() {
				break
			}
			//p.State.SetError("directory is a file")
			break
		}

		// @TODO - If we change dir and it's relative, we will lose the path.
		// @TODO - This can be both good or bad.

		err := os.Chdir(p._Path)
		p.State.SetError(err)
		if p.State.IsError() {
			break
		}

		p.State.SetOk("chdir OK")
	}

	return p.State
}


func (p *TypeOsPath) GetCwd() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.Clear()

		if !p.IsValid() {
			break
		}


		var cwd string
		//p.State.SetResponse(&cwd)
		var err error
		cwd, err = os.Getwd()
		p.State.SetError(err)
		if p.State.IsError() {
			break
		}

		p.State.SetResponse(&cwd)
		p.State.Clear()
	}

	return p.State
}


func (p *TypeOsPath) IsCwd() bool {
	var ok bool

	for range onlyOnce {
		p.State.SetFunction()

		cwd, err := os.Getwd()
		p.State.SetError(err)
		if p.State.IsError() {
			break
		}

		if cwd != p._Path {
			break
		}

		ok = true
	}

	return ok
}


func (p *TypeOsPath) Mkdir() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.Clear()

		if !p.IsValid() {
			break
		}


		if p._Mode == 0 {
			p._Mode = 0755
		}

		ok := false
		p.State.SetResponse(&ok)
		var err error

		if p._Dirname != "" {
			err = os.Mkdir(p._Dirname, p._Mode)
		} else {
			err = os.Mkdir(p._Path, p._Mode)
		}

		p.State.SetError(err)
		if p.State.IsError() {
			break
		}

		ok = true
		p.State.SetResponse(&ok)
		p.State.Clear()
		p.State.SetOk("mkdir OK")
	}

	return p.State
}
