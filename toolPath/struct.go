package toolPath

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"time"
)


type OsPathGetter interface {
}


type TypeOsPath struct {
	Path      string `json:"path"`
	_Filename string
	_Dirname  string
	_IsDir    bool
	_IsFile   bool
	_Exists   bool
	_ModTime  time.Time
	_Name     string
	_Mode     os.FileMode
	_Size     int64

	_String       string
	_Array        []string
	_Separator    string
	FileHandle    *os.File `json:"-"`

	_Valid        bool
	_CanOverwrite bool
	_CanRemove    bool
	_Remote       bool

	State         *ux.State    `json:"-"`
}
func (p *TypeOsPath) IsNil() *ux.State {
	return ux.IfNilReturnError(p)
}


func New(runtime *toolRuntime.TypeRuntime) *TypeOsPath {
	runtime = runtime.EnsureNotNil()

	p := &TypeOsPath{
		Path:      "",
		_Filename: "",
		_Dirname:  "",
		_IsDir:    false,
		_IsFile:   false,
		_Exists:   false,
		_ModTime:  time.Time{},
		_Mode:     0,
		_Size:     0,
		_String:   "",
		_Array:    nil,
		_Separator:    DefaultSeparator,
		_Valid:        false,
		_CanRemove:    false,
		_CanOverwrite: false,

		State:         ux.NewState(runtime.CmdName, runtime.Debug),
	}
	p.State.SetPackage("")
	p.State.SetFunctionCaller()
	return p
}


//func (p *TypeOsPath) IsNil() *ux.State {
//	if state := ux.IfNilReturnError(p); state.IsError() {
//		return state
//	}
//	p.State = p.State.EnsureNotNil()
//	return p.State
//}

// Replaced with:
// 	if state := ux.IfNilReturnError(at); state.IsError() {
//		return state
//	}

func (p *TypeOsPath) EnsureNotNil() *TypeOsPath {
	if p == nil {
		return New(nil)
	}
	return p
}
