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

		err := os.Chdir(p.Path)
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

		if cwd != p.Path {
			break
		}

		ok = true
	}

	return ok
}


func (p *TypeOsPath) Mkdir() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.SetOk()

		if !p.IsValid() {
			break
		}

		if p._Mode == 0 {
			p._Mode = 0755
		}

		ok := false
		p.State.SetResponse(&ok)

		dir := ""
		for range onlyOnce {
			if p._Dirname != "" {
				dir = p._Dirname
				break
			}
			if p.Path != "" {
				dir = p.Path
				break
			}
			p.State.SetError("Path is empty")
		}
		if p.State.IsNotOk() {
			break
		}

		p.StatPath()
		if p.Exists() {
			p.State.SetOk()	// Ignore errors.
			break
		}

		err := os.Mkdir(dir, p._Mode)
		if err != nil {
			p.State.SetError(err)
			break
		}

		p.State = p.StatPath()
		if p.State.IsNotOk() {
			break
		}

		ok = true
		p.State.SetResponse(&ok)
		//p.State.Clear()
		p.State.SetOk("mkdir OK")
	}

	return p.State
}


func (p *TypeOsPath) MkdirAll() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.SetOk()

		if !p.IsValid() {
			break
		}

		if p._Mode == 0 {
			p._Mode = 0755
		}

		ok := false
		p.State.SetResponse(&ok)

		dir := ""
		for range onlyOnce {
			if p._Dirname != "" {
				dir = p._Dirname
				break
			}
			if p.Path != "" {
				dir = p.Path
				break
			}
			p.State.SetError("Path is empty")
		}
		if p.State.IsNotOk() {
			break
		}

		p.StatPath()
		if p.Exists() {
			p.State.SetOk()	// Ignore errors.
			break
		}

		err := os.MkdirAll(dir, p._Mode)
		if err != nil {
			p.State.SetError(err)
			break
		}

		p.State = p.StatPath()
		if p.State.IsNotOk() {
			break
		}

		ok = true
		p.State.SetResponse(&ok)
		//p.State.Clear()
		p.State.SetOk("mkdir OK")
	}

	return p.State
}
