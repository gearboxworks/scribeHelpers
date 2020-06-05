package toolPath

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"path/filepath"
	"regexp"
)


type TypeOsPaths struct {
	Paths   []*TypeOsPath
	Base    *TypeOsPath

	Runtime *toolRuntime.TypeRuntime
	State   *ux.State
}


func NewPaths(runtime *toolRuntime.TypeRuntime) *TypeOsPaths {
	runtime = runtime.EnsureNotNil()

	p := &TypeOsPaths{
		Paths:   make([]*TypeOsPath, 0),
		Base:    New(runtime),

		Runtime: runtime,
		State:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
	p.State.SetPackage("")
	p.State.SetFunctionCaller()
	return p
}


func (p *TypeOsPaths) IsNil() *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}
	p.State = p.State.EnsureNotNil()
	return p.State
}


func (p *TypeOsPaths) SetBasePath(path ...string) *ux.State {
	if state := p.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if len(path) == 0 {
			//p.State.SetError("no path specified")
			break
		}

		if !p.Base.SetPath(path...) {
			p.State.SetError("no path specified")
			break
		}

		p.State = p.Base.StatPath()
		if p.State.IsNotOk() {
			break
		}
	}

	return p.State
}


func (p *TypeOsPaths) Find(path ...string) *ux.State {
	if state := p.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if len(path) != 0 {
			// Allows setting path before this function.
			// Good for multiple runs.
			p.State = p.SetBasePath(path...)
			if p.State.IsNotOk() {
				break
			}
		}

		p.State = p.find("")
		if p.State.IsNotOk() {
			break
		}
	}

	return p.State
}


func (p *TypeOsPaths) FindRegex(re string, path ...string) *ux.State {
	if state := p.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if len(path) != 0 {
			// Allows setting path before this function.
			// Good for multiple runs.
			p.State = p.SetBasePath(path...)
			if p.State.IsNotOk() {
				break
			}
		}

		p.State = p.find(re)
		if p.State.IsNotOk() {
			break
		}
	}

	return p.State
}


func (p *TypeOsPaths) GetLength() int {
	return len(p.Paths)
}


func (p *TypeOsPaths) find(re string) *ux.State {
	if state := p.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var match *regexp.Regexp
		if re != "" {
			match = regexp.MustCompile(re)
		}

		// Always operate on the GetDirname - TypeOsPath will always properly resolve this for us.
		err := filepath.Walk(p.Base.GetDirname(),
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if re != "" {
					if !match.MatchString(path) {
						return nil
					}
				}

				p.appendPath(path)
				return nil
			})

		if err != nil {
			p.State.SetError(err)
			break
		}

		if p.GetLength() == 0 {
			p.State.SetWarning("no files found")
			break
		}
	}

	return p.State
}


func (p *TypeOsPaths) appendPath(path ...string) *ux.State {
	if state := p.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		p2 := New(nil)
		p2.SetPath(path...)
		p.State = p2.StatPath()
		if p.State.IsNotOk() {
			break
		}
		p.Paths = append(p.Paths, p2)
	}

	return p.State
}
