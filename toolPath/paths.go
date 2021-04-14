package toolPath

import (
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)


type TypeOsPaths struct {
	Paths   []*TypeOsPath
	Base    *TypeOsPath

	recurse bool                     `json:"-"`
	Runtime *toolRuntime.TypeRuntime `json:"-"`
	State   *ux.State                `json:"-"`
}


func NewPaths(runtime *toolRuntime.TypeRuntime) *TypeOsPaths {
	runtime = runtime.EnsureNotNil()

	p := &TypeOsPaths{
		Paths:   make([]*TypeOsPath, 0),
		Base:    New(runtime),

		recurse: true,

		Runtime: runtime,
		State:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
	p.State.SetPackage("")
	p.State.SetFunctionCaller()
	return p
}


//func (p *TypeOsPaths) IsNil() *ux.State {
//	if state := ux.IfNilReturnError(p); state.IsError() {
//		return state
//	}
//	p.State = p.State.EnsureNotNil()
//	return p.State
//}
// 	if state := ux.IfNilReturnError(at); state.IsError() {
//		return state
//	}


func (p *TypeOsPaths) SetBasePath(path ...string) *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}

	for range onlyOnce {
		if len(path) == 0 {
			//p.State.SetError("no path specified")
			break
		}

		p.Base.SetPath(path...)

		p.State = p.Base.StatPath()
		if p.State.IsNotOk() {
			break
		}
	}

	return p.State
}


func (p *TypeOsPaths) AddRelPath(path ...string) *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
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


func (p *TypeOsPaths) SetRecursive() *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}
	p.recurse = true
	return p.State
}
func (p *TypeOsPaths) SetNonRecursive() *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}
	p.recurse = false
	return p.State
}


func (p *TypeOsPaths) Find(path ...string) *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
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

		files := p.find()
		if p.State.IsNotOk() {
			break
		}

		for _, f := range files {
			p.appendPath(f.Path)
		}
	}

	return p.State
}


func (p *TypeOsPaths) FindRegex(re string, path ...string) *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
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

		files := p.find()
		if p.State.IsNotOk() {
			break
		}

		match := regexp.MustCompile(re)
		for _, f := range files {
			if match == nil {
				p.appendPath(f.Path)
				continue
			}
			if match.MatchString(f.Path) {
				p.appendPath(f.Path)
				continue
			}
		}
	}

	return p.State
}


func (p *TypeOsPaths) FindByExt(ext string, path ...string) *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
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

		files := p.find()
		if p.State.IsNotOk() {
			break
		}

		re := ".*\\." + ext + "$"
		match := regexp.MustCompile(re)
		for _, f := range files {
			if match == nil {
				p.appendPath(f.Path)
				continue
			}
			if match.MatchString(f.Path) {
				p.appendPath(f.Path)
				continue
			}
		}
	}

	return p.State
}


func (p *TypeOsPaths) GetLength() int {
	return len(p.Paths)
}

type ffiles struct {
	FileInfo os.FileInfo
	Path     string
}

func (p *TypeOsPaths) find() []ffiles {
	var files []ffiles
	if state := ux.IfNilReturnError(p); state.IsError() {
		return files
	}

	for range onlyOnce {
		var err error
		if p.recurse {
			// Always operate on the GetDirname - TypeOsPath will always properly resolve this for us.
			err = filepath.Walk(p.Base.GetDirname(),
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					files = append(files, ffiles{
						FileInfo: info,
						Path:     path,
					})
					return nil
				})

		} else {
			// Always operate on the GetDirname - TypeOsPath will always properly resolve this for us.
			var f2 []os.FileInfo
			f2, err = ioutil.ReadDir(p.Base.GetDirname())
			for _, f := range f2 {
				files = append(files, ffiles{
					FileInfo: f,
					Path:     filepath.Join(p.Base._Dirname, f.Name()),
				})
			}
		}

		if err != nil {
			p.State.SetError(err)
			break
		}

		if len(files) == 0 {
			p.State.SetWarning("no files found")
			break
		}
	}

	return files
}


//func (p *TypeOsPaths) find(re string) *ux.State {
//	if state := p.IsNil(); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		var match *regexp.Regexp
//		if re != "" {
//			match = regexp.MustCompile(re)
//		}
//
//		// Always operate on the GetDirname - TypeOsPath will always properly resolve this for us.
//		files, err := ioutil.ReadDir(p.Base.GetDirname())
//		if err != nil {
//			p.State.SetError(err)
//			break
//		}
//
//		for _, f := range files {
//			if match == nil {
//				p.appendPath(f.Name())
//				continue
//			}
//			if match.MatchString(f.Name()) {
//				p.appendPath(f.Name())
//				continue
//			}
//		}
//
//		if p.GetLength() == 0 {
//			p.State.SetWarning("no files found")
//			break
//		}
//	}
//
//	return p.State
//}
//
//
//func (p *TypeOsPaths) findRecursive(re string) *ux.State {
//	if state := p.IsNil(); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		var match *regexp.Regexp
//		if re != "" {
//			match = regexp.MustCompile(re)
//		}
//
//		// Always operate on the GetDirname - TypeOsPath will always properly resolve this for us.
//		err := filepath.Walk(p.Base.GetDirname(),
//			func(path string, info os.FileInfo, err error) error {
//				if err != nil {
//					return err
//				}
//				if match == nil {
//					p.appendPath(path)
//					return nil
//				}
//				if match.MatchString(path) {
//					p.appendPath(path)
//					return nil
//				}
//				return nil
//			})
//
//		if err != nil {
//			p.State.SetError(err)
//			break
//		}
//
//		if p.GetLength() == 0 {
//			p.State.SetWarning("no files found")
//			break
//		}
//	}
//
//	return p.State
//}


func (p *TypeOsPaths) appendPath(path ...string) *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
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
