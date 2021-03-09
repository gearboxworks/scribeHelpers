package toolPath

import (
	"github.com/newclarity/scribeHelpers/toolTypes"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)


func (p *TypeOsPath) GetPath() string {
	return p.Path
}
func (p *TypeOsPath) GetPathAbs() string {
	var s string

	for range onlyOnce {
		var err error
		s, err = filepath.Abs(p.Path)
		if err != nil {
			s = p.Path
			break
		}
	}

	return s
}
func (p *TypeOsPath) GetPathRel() string {
	var s string

	for range onlyOnce {
		var err error

		s, err = os.Getwd()
		if err != nil {
			s = p.Path
			break
		}

		s, err = filepath.Rel(s, p.Path)
		if err != nil {
			s = p.Path
			break
		}
	}

	return s
}
func (p *TypeOsPath) SetPath(path ...string) bool {
	p.Path = ""
	return p.AppendPath(path...)
}

func (p *TypeOsPath) SetFile(path ...string) bool {
	p.Path = ""
	return p.AppendFile(path...)
}

func (p *TypeOsPath) SetDir(path ...string) bool {
	p.Path = ""
	return p.AppendDir(path...)
}

func (p *TypeOsPath) AppendFile(path ...string) bool {
	var ok bool
	for range onlyOnce {
		ok = p.AppendPath(path...)
		if !ok {
			break
		}
		p.ThisIsAFile()
	}
	return ok
}

func (p *TypeOsPath) AppendDir(path ...string) bool {
	var ok bool
	for range onlyOnce {
		ok = p.AppendPath(path...)
		if !ok {
			break
		}
		p.ThisIsADir()
	}
	return ok
}

func (p *TypeOsPath) AppendPath(path ...string) bool {
	var ok bool

	for range onlyOnce {
		if p._IsRemotePath(p.Path) {
			ok = p._AppendRemotePath(path...)
			break
		}

		if p._IsRemotePath(path...) {
			ok = p._AppendRemotePath(path...)
			break
		}

		if p._IsRelativeLocalPath(path...) {
			ok = p._AppendRelativeLocalPath(path...)
			break
		}

		ok = p._AppendLocalPath(path...)
	}

	return ok
}
func (p *TypeOsPath) _AppendLocalPath(path ...string) bool {
	for range onlyOnce {
		p._Valid = false
		p.Path = _GetAbsPath(path...)
		if p.Path == "" {
			p.State.SetError("src path empty")
			break
		}
		//p._Valid = true
		p._Remote = false

		// Reset these until a later StatPath()
		p._Dirname = ""
		p._Filename = ""
		p._IsDir = false
		p._IsFile = false
		p._Exists = false
		p._Valid = true
	}

	return p._Valid
}
func (p *TypeOsPath) _AppendRelativeLocalPath(path ...string) bool {
	for range onlyOnce {
		p._Valid = false
		p.Path = filepath.Join(p.Path, filepath.Join(path...))
		if p.Path == "" {
			p.State.SetError("src path empty")
			break
		}
		//p._Valid = true
		p._Remote = false

		// Reset these until a later StatPath()
		p._Dirname = ""
		p._Filename = ""
		p._IsDir = false
		p._IsFile = false
		p._Exists = false
		p._Valid = true
	}

	return p._Valid
}
func (p *TypeOsPath) _AppendRemotePath(path ...string) bool {
	for range onlyOnce {
		p._Valid = false
		p.Path = filepath.Join(path...)
		if p.Path == "" {
			p.State.SetError("src path empty")
			break
		}
		p._Valid = true
		p._Remote = true
	}

	return p._Valid
}
func (p *TypeOsPath) _IsRemotePath(path ...string) bool {
	return strings.ContainsAny(strings.Join(path, ""), ":@")
}
func (p *TypeOsPath) _IsRelativeLocalPath(path ...string) bool {
	return !filepath.IsAbs(filepath.Join(path...))
}


func (p *TypeOsPath) IsAbs() bool {
	return filepath.IsAbs(p.Path)
}
func (p *TypeOsPath) IsRelative() bool {
	return !p.IsAbs()
}


//func (p *TypeOsPath) SetRemote() {
//	// @TODO - Add in extra logic to convert filename to path.
//	p._Remote = true
//}
func (p *TypeOsPath) IsRemote() bool {
	return p._Remote
}


//func (p *TypeOsPath) SetFilename(filename string) {
//	// @TODO - Add in extra logic to convert filename to path.
//	p._Filename = filename
//}
func (p *TypeOsPath) GetFilename() string {
	return p._Filename
}


func (p *TypeOsPath) GetParentDir() string {
	var ret string
	for range onlyOnce {
		// Filesystems are funny, especially when you symlink a directory.
		// The parent, (".."), doesn't always point to the direct parent.
		d1 := p.GetDirnameAbs()
		d2 := p.GetParentDirAbs()
		s, err := filepath.Rel(d1, d2)
		if err == nil {
			ret = s
		}
		//var s string
		//var s2 string
		//var err error
		//ret = filepath.Join(p._Dirname, "..")
		//s, err = filepath.Abs(ret)
		//if err != nil {
		//	break
		//}
		//s2, err = filepath.Abs(p._Dirname)
		//if err != nil {
		//	break
		//}
		//s, err = filepath.Rel(s2, s)
		//if err == nil {
		//	ret = s
		//}
	}
	return ret
}
func (p *TypeOsPath) GetParentDirAbs() string {
	var ret string
	for range onlyOnce {
		// Filesystems are funny, especially when you symlink a directory.
		// The parent, (".."), doesn't always point to the direct parent.
		var s string
		var err error
		ret = filepath.Join(p._Dirname, "..")
		s, err = filepath.Abs(ret)
		if err != nil {
			break
		}
		ret = s
	}
	return ret
}


func (p *TypeOsPath) GetDirname() string {
	return p._Dirname
}
func (p *TypeOsPath) GetDirnameAbs() string {
	var s string
	for range onlyOnce {
		var err error
		s, err = filepath.Abs(p._Dirname)
		if err != nil {
			s = p._Dirname
		}
	}
	return s
}


func (p *TypeOsPath) SetModTime(time time.Time) {
	p._ModTime = time
}
func (p *TypeOsPath) GetModTime() time.Time {
	return p._ModTime
}
func (p *TypeOsPath) GetModTimeString() string {
	return p._ModTime.Format("2006-01-02T15:04:05-0700")
}
func (p *TypeOsPath) GetModTimeEpoch() int64 {
	return p._ModTime.Unix()
}


func (p *TypeOsPath) SetMode(mode os.FileMode) {
	p._Mode = mode
}
func (p *TypeOsPath) GetMode() os.FileMode {
	return p._Mode
}


//func (p *TypeOsPath) SetSize(size int64) {
//	p._Size = size
//}
func (p *TypeOsPath) GetSize() int64 {
	return p._Size
}


//func (p *TypeOsPath) SetExists() {
//	p._Exists = true
//}
func (p *TypeOsPath) Exists() bool {
	var ok bool

	for range onlyOnce {
		if !p.IsValid() {
			break
		}
		if !p._Exists {
			p.State.SetError("path does not exist")
			break
		}
		p.State.SetOk("path exists")
		ok = p._Exists
	}

	return ok
}
func (p *TypeOsPath) NotExists() bool {
	return !p.Exists()
}
func (p *TypeOsPath) FileExists() bool {
	var ok bool

	for range onlyOnce {
		if !p.IsValid() {
			break
		}
		if !p._Exists {
			p.State.SetError("file does not exist")
			break
		}
		if !p._IsFile {
			p.State.SetError("file is a dir")
			break
		}
		p.State.SetOk("file exists")
		ok = p._Exists
	}

	return ok
}
func (p *TypeOsPath) DirExists() bool {
	var ok bool

	for range onlyOnce {
		if !p.IsValid() {
			break
		}
		if !p._Exists {
			p.State.SetError("dir does not exist")
			break
		}
		if !p._IsDir {
			p.State.SetError("dir is a file")
			break
		}
		p.State.SetOk("dir exists")
		ok = p._Exists
	}

	return ok
}


func (p *TypeOsPath) ThisIsAFile() {
	p._IsFile = true
	p._IsDir = false
	p.State.SetOk()
}
func (p *TypeOsPath) IsAFile() bool {
	return p._IsFile
}


func (p *TypeOsPath) ThisIsADir() {
	p._IsFile = false
	p._IsDir = true
	p.State.SetOk()
}
func (p *TypeOsPath) IsADir() bool {
	return p._IsDir
}


func (p *TypeOsPath) _SetValid() {
	p._Valid = true
}
func (p *TypeOsPath) _SetInvalid() {
	p._Valid = false
}
func (p *TypeOsPath) IsValid() bool {
	for range onlyOnce {
		//if !p._Valid {
		//	p.State.SetError("path not valid")
		//	break
		//}

		if p.Path == "" {
			p.State.SetError("path not set")
			break
		}

		p._Valid = true
	}

	return p._Valid
}
func (p *TypeOsPath) IsInvalid() bool {
	return !p.IsValid()
}
func (p *TypeOsPath) IsNotValid() bool {
	return !p.IsValid()
}


func (p *TypeOsPath) SetOverwriteable() {
	p._CanOverwrite = true
}
func (p *TypeOsPath) CanOverwrite() bool {
	return p._CanOverwrite
}
func (p *TypeOsPath) IsOverwriteable() bool {
	return p._CanOverwrite
}


func (p *TypeOsPath) SetRemoveable() {
	p._CanRemove = true
}
func (p *TypeOsPath) CanRemove() bool {
	return p._CanRemove
}
func (p *TypeOsPath) IsRemoveable() bool {
	return p._CanRemove
}


func ReflectFileMode(ref interface{}) *os.FileMode {
	var fm os.FileMode

	for range onlyOnce {
		value := reflect.ValueOf(ref)
		if value.Kind() != reflect.Uint32 {
			break
		}

		fm = os.FileMode(value.Uint())
	}

	return &fm
}


func ReflectPath(ref ...interface{}) *string {
	var fp string

	for range onlyOnce {
		var path []string
		for _, r := range ref {
			// Sometimes we can have dirs within each string slice.
			// EG: [0] = "dir1/dir2" OR [0] = "dir1\dir2"
			// This handles paths across O/S sanely.
			p := filepath.SplitList(*toolTypes.ReflectString(r))

			path = append(path, p...)
		}
		fp = filepath.Join(path...)
	}

	return &fp
}


func ReflectAbsPath(ref ...interface{}) *string {
	var fp string

	for range onlyOnce {
		path := ReflectPath(ref...)

		var err error
		fp, err = filepath.Abs(*path)
		if err != nil {
			fp = *path
		}
	}

	return &fp
}


func _GetAbsPath(p ...string) string {
	var ret string

	for range onlyOnce {
		ret = filepath.Join(p...)

		if filepath.IsAbs(ret) {
			break
		}

		var err error
		ret, err = filepath.Abs(ret)
		if err != nil {
			ret = ""
			break
		}
	}

	return ret
}
