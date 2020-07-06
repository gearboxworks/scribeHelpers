package toolPath

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)


func (p *TypeOsPath) StatPath() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.SetOk()

		if p.Path == "" {
			p.State.SetError("path is empty")
			break
		}

		if p._Remote {
			// @TODO - Maybe add in some remote checks?
			p._Valid = true
			p._Exists = true
			p.State.SetOk("path is remote")
			break
		}

		if strings.HasPrefix(p.Path, "~/") {
			u, err := user.Current()
			if err != nil {
				p.State.SetError(err)
				p._Valid = false
				p._Exists = false
				break
			}
			p.Path = strings.TrimPrefix(p.Path, "~/")
			p.Path = filepath.Join(u.HomeDir, p.Path)
		}

		var stat os.FileInfo
		var err error
		stat, err = os.Stat(p.Path)
		if os.IsNotExist(err) {
			p.State.SetError("path does not exist - %s", err)
			p._Exists = false
			break
		}
		p.State.SetError(err)
		if p.State.IsError() {
			break
		}

		p._Valid = true
		p._Exists = true
		p._ModTime = stat.ModTime()
		p._Name = stat.Name()
		p._Mode = stat.Mode()
		p._Size = stat.Size()

		if stat.IsDir() {
			p._IsDir = true
			p._IsFile = false
			p._Dirname = fmt.Sprintf("%s%c", p.Path, filepath.Separator)
			p.Path = p._Dirname
			p._Filename = ""

		} else {
			p._IsDir = false
			p._IsFile = true
			p._Dirname = fmt.Sprintf("%s%c", filepath.Dir(p.Path), filepath.Separator)
			p._Filename = filepath.Base(p.Path)
		}

		p.State.SetOk("stat OK")
	}

	return p.State
}


func (p *TypeOsPath) Chmod(m os.FileMode) *ux.State {
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

		var err error
		err = os.Chmod(p.Path, m)
		p.State.SetError(err)
		if p.State.IsError() {
			break
		}

		p.State = p.StatPath()
		if p.State.IsError() {
			break
		}

		p.State.SetOk("chmod OK")
	}

	return p.State
}


func (p *TypeOsPath) ChangeExtension(ext string) {
	s := filepath.Ext(p.Path)
	p.Path = p.Path[:len(p.Path) - len(s)] + ext
}


func (p *TypeOsPath) ChangeSuffix(ext string) {
	p.ChangeExtension(ext)
}
