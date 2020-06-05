// High level helper functions available within templates - general file related.
package toolPath

import (
	"github.com/newclarity/scribeHelpers/ux"
)


type ToolOsPath TypeOsPath
func (g *ToolOsPath) Reflect() *TypeOsPath {
	return (*TypeOsPath)(g)
}
func (p *TypeOsPath) Reflect() *ToolOsPath {
	return (*ToolOsPath)(p)
}

func (c *ToolOsPath) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}


// Usage:
//		{{ $str := ReadFile "filename.txt" }}
//func ToolNewPath(file ...interface{}) *ToolOsPath {
func ToolNewPath(file ...interface{}) *TypeOsPath {
	ret := New(nil)

	for range onlyOnce {
		ret.State.SetFunction("")

		f := ReflectPath(file...)
		if f == nil {
			ret.State.SetOk("path empty")
			break
		}

		if !ret.SetPath(*f) {
			ret.State.SetError("path error")
			break
		}

		ret.State.SetState(ret.StatPath())
		if ret.State.IsError() {
			break
		}
	}

	//return ReflectToolOsPath(ret)
	return ret
}


// Usage:
//		{{ $ret := Chmod 0644 "/root" ... }}
//		{{ if $ret.IsOk }}Changed perms of file {{ $ret.Dir }}{{ end }}
func ToolChmod(mode interface{}, path ...interface{}) *TypeOsPath {
	ret := New(nil)

	for range onlyOnce {
		ret.State.SetFunction("")

		f := ReflectPath(path...)
		if f == nil {
			ret.State.SetError("path empty")
			break
		}
		ret.SetPath(*f)

		m := ReflectFileMode(mode)
		if m == nil {
			break
		}

		ret.State.SetState(ret.Chmod(*m))
		if ret.State.IsError() {
			break
		}
	}

	return ret
}
