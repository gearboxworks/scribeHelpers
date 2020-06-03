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
	return p._Path
}
func (p *TypeOsPath) GetPathAbs() string {
	var s string

	for range OnlyOnce {
		var err error
		s, err = filepath.Abs(p._Path)
		if err != nil {
			s = p._Path
			break
		}
	}

	return s
}
func (p *TypeOsPath) GetPathRel() string {
	var s string

	for range OnlyOnce {
		var err error

		s, err = os.Getwd()
		if err != nil {
			s = p._Path
			break
		}

		s, err = filepath.Rel(s, p._Path)
		if err != nil {
			s = p._Path
			break
		}
	}

	return s
}
func (p *TypeOsPath) SetPath(path ...string) bool {
	p._Path = ""
	return p.AppendPath(path...)
}
func (p *TypeOsPath) AppendPath(path ...string) bool {
	var ok bool

	for range OnlyOnce {
		if p._IsRemotePath(p._Path) {
			ok = p._AppendRemotePath(path...)
			break
		}
		if p._IsRemotePath(path...) {
			ok = p._AppendRemotePath(path...)
			break
		}

		ok = p._AppendLocalPath(path...)
	}

	return ok
}
func (p *TypeOsPath) _AppendLocalPath(path ...string) bool {
	for range OnlyOnce {
		p._Valid = false
		p._Path = _GetAbsPath(path...)
		if p._Path == "" {
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
	}

	return p._Valid
}
func (p *TypeOsPath) _AppendRemotePath(path ...string) bool {
	for range OnlyOnce {
		p._Valid = false
		// @TODO - May have to change this logic to:
		// @TODO - p._Path = strings.Join(path, "")
		p._Path = filepath.Join(path...)
		if p._Path == "" {
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


//func (p *TypeOsPath) SetDirname(dirname string) {
//	// @TODO - Add in extra logic to convert dirname to path.
//	p._Dirname = dirname
//}
func (p *TypeOsPath) GetDirname() string {
	return p._Dirname
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

	for range OnlyOnce {
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

	for range OnlyOnce {
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

	for range OnlyOnce {
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
	p.State.Clear()
}
func (p *TypeOsPath) IsAFile() bool {
	return p._IsFile
}


func (p *TypeOsPath) ThisIsADir() {
	p._IsFile = false
	p._IsDir = true
	p.State.Clear()
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
	for range OnlyOnce {
		//if !p._Valid {
		//	p.State.SetError("path not valid")
		//	break
		//}

		if p._Path == "" {
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

	for range OnlyOnce {
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

	for range OnlyOnce {
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

	for range OnlyOnce {
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

	for range OnlyOnce {
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
