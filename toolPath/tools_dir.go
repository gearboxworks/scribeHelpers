package toolPath


// Usage:
//		{{ $ret := Chdir "/root" }}
//		{{ if $ret.IsOk }}OK{{ end }}
func ToolChdir(dir ...interface{}) *TypeOsPath {
	ret := New(nil)

	for range onlyOnce {
		ret.State.SetFunction()

		f := ReflectPath(dir...)
		if f == nil {
			ret.State.SetError("directory is empty")
			break
		}
		ret.SetPath(*f)

		ret.State = ret.Chdir()
	}

	return ret
}


// Usage:
//		{{ $ret := GetCwd }}
//		{{ if $ret.IsOk }}Current directory is {{ $ret.Dir }}{{ end }}
func ToolGetCwd() *TypeOsPath {
	ret := New(nil)

	for range onlyOnce {
		ret.State.SetFunction()

		state := ret.GetCwd()
		ret.State = state
		if ret.State.IsError() {
			break
		}
	}

	return ret
}


// Usage:
//		{{ $ret := GetCwd }}
//		{{ if $ret.IsOk }}Current directory is {{ $ret.Dir }}{{ end }}
func ToolIsCwd() *TypeOsPath {
	ret := New(nil)

	for range onlyOnce {
		ret.State.SetFunction()

		ok := ret.IsCwd()
		ret.State.SetResponse(&ok)
	}

	return ret
}


// Usage:
//		{{ $ret := Chmod 0644 "/root" ... }}
//		{{ if $ret.IsOk }}Changed perms of file {{ $ret.Dir }}{{ end }}
func ToolCreateDir(path ...interface{}) *TypeOsPath {
	ret := New(nil)

	for range onlyOnce {
		ret.State.SetFunction()

		f := ReflectPath(path...)
		if f == nil {
			ret.State.SetError("path empty")
			break
		}
		ret.SetPath(*f)

		ret.State = ret.Mkdir()
		if ret.State.IsError() {
			break
		}
	}

	return ret
}


// Usage:
//		{{ $ret := Chmod 0644 "/root" ... }}
//		{{ if $ret.IsOk }}Changed perms of file {{ $ret.Dir }}{{ end }}
func ToolRemoveDir(path ...interface{}) *TypeOsPath {
	ret := New(nil)

	for range onlyOnce {
		ret.State.SetFunction()

		f := ReflectPath(path...)
		if f == nil {
			ret.State.SetError("path empty")
			break
		}
		ret.SetPath(*f)

		//if force {
		//	ret.SetRemoveable()
		//}
		ret.State = ret.RemoveDir()
		if ret.State.IsError() {
			break
		}
	}

	return ret
}
