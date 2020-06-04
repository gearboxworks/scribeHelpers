package toolPath

import (
	"github.com/newclarity/scribeHelpers/toolTypes"
)


// Usage:
//		{{ $str := ReadFile "filename.txt" }}
func ToolReadFile(file ...interface{}) *TypeOsPath {
	ret := New(nil)

	for range OnlyOnce {
		ret.State.SetFunction("")

		f := ReflectPath(file...)
		if f == nil {
			ret.State.SetError("filename empty")
			break
		}

		ret.SetPath(*f)
		ret.State.SetState(ret.ReadFile())
		if ret.State.IsError() {
			break
		}

		// Make available OsPath structure.
		ret.State.Response = ret.GetContentString()
	}

	return ret
}


// Usage:
//		{{ $return := WriteFile .Data.Source 0644 "dir1" "dir2/dir3" "filename.txt" }}
func ToolWriteFile(contents interface{}, perms interface{}, file ...interface{}) *TypeOsPath {
	ret := New(nil)

	for range OnlyOnce {
		ret.State.SetFunction("")

		f := ReflectPath(file...)
		if f == nil {
			ret.State.SetError("filename is nil")
			break
		}
		ret.SetPath(*f)

		c := toolTypes.ReflectString(contents)
		if c == nil {
			ret.State.SetError("content string is nil")
			break
		}
		ret.LoadContents(*c)

		p := ReflectFileMode(perms)
		if p == nil {
			ret.SetMode(0)
		} else {
			ret.SetMode(*p)
		}

		ret.State.SetState(ret.WriteFile())
		if ret.State.IsError() {
			break
		}
	}

	return ret
}


// Usage:
//		{{ $ret := Chmod 0644 "/root" ... }}
//		{{ if $ret.IsOk }}Changed perms of file {{ $ret.Dir }}{{ end }}
func ToolRemoveFile(path ...interface{}) *TypeOsPath {
	ret := New(nil)

	for range OnlyOnce {
		ret.State.SetFunction("")

		f := ReflectPath(path...)
		if f == nil {
			ret.State.SetError("file empty")
			break
		}
		ret.SetPath(*f)

		//if force {
		//	ret.SetRemoveable()
		//}
		ret.State.SetState(ret.RemoveFile())
		if ret.State.IsNotOk() {
			break
		}
		//if ret.Exists() {
		//	ret.State.SetError("file couldn't be removed")
		//	break
		//}
		ret.State.SetOk("file removed OK")
	}

	return ret
}
