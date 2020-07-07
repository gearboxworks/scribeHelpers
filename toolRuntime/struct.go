package toolRuntime

import (
	"github.com/kardianos/osext"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"os/user"
	"path"
	"path/filepath"
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

	WorkingDir     string		`json:"working_dir" mapstructure:"working_dir"`
	ConfigDir      string		`json:"config_dir" mapstructure:"config_dir"`
	CacheDir       string		`json:"cache_dir" mapstructure:"cache_dir"`
	TempDir        string		`json:"temp_dir" mapstructure:"temp_dir"`

	FullArgs       ExecArgs		`json:"full_args" mapstructure:"full_args"`
	Args           ExecArgs		`json:"args" mapstructure:"args"`
	ArgFiles       ExecArgs		`json:"arg_files" mapstructure:"arg_files"`

	Env            ExecEnv		`json:"env" mapstructure:"env"`
	EnvMap         Environment	`json:"env_map" mapstructure:"env_map"`

	TimeStamp      time.Time	`json:"timestamp" mapstructure:"timestamp"`

	GoRuntime      GoRuntime	`json:"go_runtime" mapstructure:"go_runtime"`

	User           User			`json:"user" mapstructure:"user"`

	Debug          bool			`json:"debug" mapstructure:"debug"`
	Verbose        bool			`json:"verbose" mapstructure:"verbose"`
	State          *ux.State	`json:"state" mapstructure:"state"`
}
func (r *TypeRuntime) IsNil() *ux.State {
	return ux.IfNilReturnError(r)
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

type User struct {
	*user.User
}

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

			WorkingDir: ".",
			ConfigDir:  ".",
			CacheDir:   ".",
			TempDir:    ".",

			FullArgs:   os.Args,
			Args:       os.Args[1:],
			ArgFiles:   []string{},

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
			Verbose:    false,
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

		ret.User.User, err = user.Current()
		if err != nil {
			ret.State.SetError(err)
			break
		}

		ret.WorkingDir, err = os.Getwd()
		if err != nil {
			ret.State.SetError(err)
			break
		}

		ret.ConfigDir, err = os.UserConfigDir()
		if err != nil {
			if runtime.GOOS == "windows" {
				ret.ConfigDir = "."
			} else {
				ret.ConfigDir = "."
			}
		}
		//ret.ConfigDir = filepath.Join(ret.ConfigDir, ret.CmdName)

		ret.CacheDir, err = os.UserCacheDir()
		if err != nil {
			if runtime.GOOS == "windows" {
				ret.CacheDir = "."
			} else {
				ret.CacheDir = "."
			}
		}
		ret.CacheDir = filepath.Join(ret.CacheDir, ret.CmdName)

		ret.TempDir = os.TempDir()
		if ret.TempDir == "" {
			if runtime.GOOS == "windows" {
				ret.TempDir = "C:\\tmp"
			} else {
				ret.TempDir = "/tmp"
			}
		}
		//ret.TempDir = filepath.Join(ret.TempDir, ret.CmdName)

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

func (r *TypeRuntime) EnsureNotNil() *TypeRuntime {
	if r == nil {
		return New("binary", "version", false)
	}
	return r
}
