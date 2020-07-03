package toolGit

import (
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/tsuyoshiwada/go-gitcmd"
)


// Usage:
//		{{ $git := NewGit }}
func ToolNewGit(path ...interface{}) *ToolGit {
	ret := New(nil)

	for range onlyOnce {
		p := toolPath.ReflectAbsPath(path...)
		if p == nil {
			break
		}
		if ret.Base.SetPath(*p) {
			state := ret.Base.StatPath()
			ret.State = state
			if ret.Base.Exists() {

			}
			if ret.State.IsError() {
				break
			}

			// Can now set it after.
			//ret.State.SetError("%s destination empty", *p)
			//break
		}

		//ret.Cmd = toolExec.NewExecCommand(false)
		ret.client = gitcmd.New(ret.GitConfig)

		if ret.IsNotAvailable() {
			break
		}
	}

	return ReflectToolGit(ret)
}
