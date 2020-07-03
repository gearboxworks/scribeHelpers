// High level helper functions available within templates - general file related.
package toolPath


// Usage:
//		{{ $str := ReadFile "filename.txt" }}
//func ToolNewPath(file ...interface{}) *ToolOsPath {
func ToolNewPath(file ...interface{}) *TypeOsPath {
	ret := New(nil)

	for range onlyOnce {
		ret.State.SetFunction()

		f := ReflectPath(file...)
		if f == nil {
			ret.State.SetOk("path empty")
			break
		}

		if !ret.SetPath(*f) {
			ret.State.SetError("path error")
			break
		}

		ret.State = ret.StatPath()
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
		ret.State.SetFunction()

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

		ret.State = ret.Chmod(*m)
		if ret.State.IsError() {
			break
		}
	}

	return ret
}
