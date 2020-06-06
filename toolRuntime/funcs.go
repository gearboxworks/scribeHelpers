package toolRuntime

import (
	"github.com/blang/semver"
	"github.com/newclarity/scribeHelpers/ux"
	"path"
	"path/filepath"
	"strings"
)


type VersionValue semver.Version


func (r *TypeRuntime) GetSemVer() *VersionValue {
	v := (VersionValue)(semver.MustParse(r.CmdVersion))
	return &v
}


func (r *TypeRuntime) PrintNameVersion() {
	ux.PrintfBlue("%s ", r.CmdName)
	ux.PrintflnCyan("v%s", r.CmdVersion)
}


func (r *TypeRuntime) TimeStampString() string {
	return r.TimeStamp.Format("2006-01-02T15:04:05-0700")
}


func (r *TypeRuntime) TimeStampEpoch() int64 {
	return r.TimeStamp.Unix()
}


func (r *TypeRuntime) GetEnvMap() *Environment {
	return &r.EnvMap
}


func (r *TypeRuntime) GetArg(index int) string {
	var ret string

	for range onlyOnce {
		if len(r.Args) > index {
			ret = r.Args[index]
		}
	}

	return ret
}


func (r *TypeRuntime) SetArgs(a ...string) error {
	var err error

	for range onlyOnce {
		r.Args = a
	}

	return err
}


func (r *TypeRuntime) GetArgs() []string {
	return r.Args
}


func (r *TypeRuntime) AddArgs(a ...string) error {
	var err error

	for range onlyOnce {
		r.Args = append(r.Args, a...)
	}

	return err
}


func (r *TypeRuntime) SetFullArgs(a ...string) error {
	var err error

	for range onlyOnce {
		r.FullArgs = a
	}

	return err
}


func (r *TypeRuntime) GetFullArgs() []string {
	return r.FullArgs
}


func (r *TypeRuntime) AddFullArgs(a ...string) error {
	var err error

	for range onlyOnce {
		r.FullArgs = append(r.FullArgs, a...)
	}

	return err
}


func (r *TypeRuntime) SetCmd(a ...string) error {
	var err error

	for range onlyOnce {
		r.Cmd, err = filepath.Abs(filepath.Join(a...))
		if err != nil {
			break
		}

		r.CmdDir = path.Dir(r.Cmd)
		r.CmdFile = path.Base(r.Cmd)
	}

	return err
}


func (r *TypeRuntime) IsRunningAs(run string) bool {
	// If OK - running executable file matches the string 'run'.
	//ok, err := regexp.MatchString("^" + run, r.CmdFile)
	ok := strings.HasPrefix(run, r.CmdFile)
	return ok
}
func (r *TypeRuntime) IsRunningAsFile() bool {
	// If OK - running executable file matches the application binary name.
	//ok, err := regexp.MatchString("^" + r.CmdName, r.CmdFile)
	ok := strings.HasPrefix(r.CmdName, r.CmdFile)
	return ok
}
func (r *TypeRuntime) IsRunningAsLink() bool {
	return !r.IsRunningAsFile()
}
