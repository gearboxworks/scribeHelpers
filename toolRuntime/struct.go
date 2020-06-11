package toolRuntime

import (
	"github.com/kardianos/osext"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)


type TypeRuntime struct {
	CmdName        string		`json:"cmd_name" mapstructure:"cmd_name"`
	CmdVersion     string		`json:"cmd_version" mapstructure:"cmd_version"`
	CmdSourceRepo  string		`json:"cmd_source_repo" mapstructure:"cmd_source_repo"`
	CmdBinaryRepo  string		`json:"cmd_binary_repo" mapstructure:"cmd_binary_repo"`

	Cmd            string		`json:"cmd" mapstructure:"cmd"`
	CmdDir         string		`json:"cmd_dir" mapstructure:"cmd_dir"`
	CmdFile        string		`json:"cmd_file" mapstructure:"cmd_file"`

	FullArgs       ExecArgs		`json:"full_args" mapstructure:"full_args"`
	Args           ExecArgs		`json:"args" mapstructure:"args"`

	Env            ExecEnv		`json:"env" mapstructure:"env"`
	EnvMap         Environment	`json:"env_map" mapstructure:"env_map"`

	TimeStamp      time.Time	`json:"timestamp" mapstructure:"timestamp"`

	GoRuntime      GoRuntime	`json:"go_runtime" mapstructure:"go_runtime"`

	Debug          bool			`json:"debug" mapstructure:"debug"`
	State          *ux.State	`json:"state" mapstructure:"state"`
}

type ExecArgs []string
type ExecEnv []string
type Environment map[string]string
type GoRuntime struct {
	Os string
	Arch string
	Root string
	Version string
	Compiler string
	NumCpus int
}

//type ExecCommand struct {
//	Dir string
//	TypeFile string
//	FullPath string
//	AsLink bool
//}
// var RunAs ExecCommand


// Instead of creating every time, let's cache the initial result in a global variable.
var globalRuntime *TypeRuntime


func New(binary string, version string, debugFlag bool) *TypeRuntime {
	var ret *TypeRuntime

	for range onlyOnce {
		if globalRuntime != nil {
			// Instead of creating every time, let's cache the initial result in a global variable.
			//globalRuntime.TimeStamp = time.Now()
			ret = globalRuntime
			break
		}

		ret = &TypeRuntime{
			CmdName:    binary,
			CmdVersion: version,
			Cmd:        "",
			CmdDir:     "",
			CmdFile:    "",
			FullArgs:   os.Args[1:],
			Args:       os.Args[1:],
			Env:        os.Environ(),
			EnvMap:     make(Environment),
			TimeStamp:  time.Now(),

			GoRuntime: GoRuntime{
				Os:       runtime.GOOS,
				Arch:     runtime.GOARCH,
				Root:     runtime.GOROOT(),
				Version:  runtime.Version(),
				Compiler: runtime.Compiler,
				NumCpus:  runtime.NumCPU(),
			},

			Debug:      debugFlag,
			State:      ux.NewState(binary, debugFlag),
		}

		for _, item := range os.Environ() {
			s := strings.SplitN(item, "=", 2)
			ret.EnvMap[s[0]] = s[1]
		}

		var err error
		var exe string
		//ret.Cmd, err = os.Executable()
		//if err != nil {
		//	ret.State.SetError(err)
		//	break
		//}
		//ret.Cmd, err = filepath.Abs(ret.Cmd)
		//if err != nil {
		//	ret.State.SetError(err)
		//	break
		//}
		exe, err = osext.Executable()
		if err != nil {
			ret.State.SetError(err)
			break
		}
		ret.Cmd =     exe
		ret.CmdDir =  path.Dir(exe)
		ret.CmdFile = path.Base(exe)

		ret.State.SetPackage("")
		ret.State.SetFunction()

		// Instead of creating every time, let's cache the initial result in a global variable.
		globalRuntime = ret
	}

	return ret
}


func (r *TypeRuntime) SetRepos(source string, binary string) *ux.State {
	if state := ux.IfNilReturnError(r); state.IsError() {
		return state
	}
	r.CmdSourceRepo = source
	r.CmdBinaryRepo = binary

	return r.State
}


func (r *TypeRuntime) IsNil() *ux.State {
	if state := ux.IfNilReturnError(r); state.IsError() {
		return state
	}
	r.State = r.State.EnsureNotNil()
	return r.State
}


func (r *TypeRuntime) EnsureNotNil() *TypeRuntime {
	if r == nil {
		return New("binary", "version", false)
	}
	return r
}
