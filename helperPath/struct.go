package helperPath

import (
	"github.com/newclarity/scribeHelpers/helperRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"time"
)


type OsPathGetter interface {
}


type TypeOsPath struct {
	State         *ux.State

	_Path         string
	_Filename     string
	_Dirname      string
	_IsDir        bool
	_IsFile       bool
	_Exists       bool
	_ModTime      time.Time
	_Name         string
	_Mode         os.FileMode
	_Size         int64

	_String       string
	_Array        []string
	_Separator    string
	fileHandle    *os.File

	_Valid        bool
	_CanOverwrite bool
	_CanRemove    bool
	_Remote       bool
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}
func ReflectHelperOsPath(p *TypeOsPath) *HelperOsPath {
	return (*HelperOsPath)(p)
}


func New(runtime *helperRuntime.TypeRuntime) *TypeOsPath {
	runtime = runtime.EnsureNotNil()

	p := &TypeOsPath{
		_Path:         "",
		_Filename:     "",
		_Dirname:      "",
		_IsDir:        false,
		_IsFile:       false,
		_Exists:       false,
		_ModTime:      time.Time{},
		_Mode:         0,
		_Size:         0,
		_String:       "",
		_Array:        nil,
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


func (p *TypeOsPath) IsNil() *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}
	p.State = p.State.EnsureNotNil()
	return p.State
}


func (p *TypeOsPath) EnsureNotNil() *TypeOsPath {
	if p == nil {
		return New(nil)
	}
	return p
}
