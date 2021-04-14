package toolRuntime

import (
	"github.com/blang/semver"
	"github.com/gearboxworks/scribeHelpers/ux"
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


func (r *TypeRuntime) SetArgs(a ...string) {
	r.Args.Set(a...)
	//var err error
	//
	//for range onlyOnce {
	//	r.Args = a
	//}
	//
	//return err
}


func (r *TypeRuntime) AddArgs(a ...string) {
	r.Args.Append(a...)
	//var err error
	//
	//for range onlyOnce {
	//	r.Args = append(r.Args, a...)
	//}
	//
	//return err
}


func (r *TypeRuntime) GetArgs() []string {
	return r.Args.GetAll()
}


func (r *TypeRuntime) GetArg(index int) string {
	return r.Args.Get(index)
}

func (r *TypeRuntime) GetArgRange(lower int, upper int) []string {
	return r.Args.Range(lower, upper)
}

func (r *TypeRuntime) SprintfArgRange(lower int, upper int) string {
	return r.Args.SprintfRange(lower, upper)
}

func (r *TypeRuntime) SprintfArgsFrom(lower int) string {
	return r.Args.SprintfFrom(lower)
}

func (r *TypeRuntime) GetNargs(begin int, size int) []string {
	return r.Args.GetFromSize(begin, size)
}

func (r *TypeRuntime) SprintfNargs(lower int, upper int) string {
	return r.Args.SprintfFromSize(lower, upper)
}


func (r *TypeRuntime) SetFullArgs(a ...string) {
	r.FullArgs.Set(a...)
	//var err error
	//
	//for range onlyOnce {
	//	r.FullArgs = a
	//}
	//
	//return err
}


func (r *TypeRuntime) AddFullArgs(a ...string) {
	r.FullArgs.Append(a...)
	//var err error
	//
	//for range onlyOnce {
	//	r.FullArgs = append(r.FullArgs, a...)
	//}
	//
	//return err
}


func (r *TypeRuntime) GetFullArgs() []string {
	return r.FullArgs.GetAll()
}


func (r *TypeRuntime) SetCmd(a ...string) error {
	var err error

	for range onlyOnce {
		r.Cmd, err = filepath.Abs(filepath.Join(a...))
		if err != nil {
			break
		}

		r.CmdDir = filepath.Dir(r.Cmd)
		r.CmdFile = filepath.Base(r.Cmd)
	}

	return err
}


func (r *TypeRuntime) IsRunningAs(run string) bool {
	var ok bool
	// If OK - running executable file matches the string 'run'.
	//ok, err := regexp.MatchString("^" + run, r.CmdFile)

	if r.IsWindows() {
		//fmt.Printf("DEBUG: WINDOWS!\n")
		ok = strings.HasPrefix(run, strings.TrimSuffix(r.CmdFile, ".exe"))
		//run = strings.TrimSuffix(run, ".exe")
	} else {
		ok = strings.HasPrefix(run, r.CmdFile)
	}
	//fmt.Printf("DEBUG: Cmd.Runtime.IsRunningAs?? %s\n", ok)
	//fmt.Printf("DEBUG: run: %s\n", run)
	//fmt.Printf("DEBUG: r.CmdName: %s\n", r.CmdName)
	//fmt.Printf("DEBUG: r.CmdFile: %s\n", r.CmdFile)
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


//func foo() {
//	fmt.Printf("Go runs OK!\n")
//	fmt.Printf("PPID: %d -> PID:%d\n", os.Getppid(), os.Getpid())
//	fmt.Printf("Compiler: %s v%s\n", runtime.Compiler, runtime.Version())
//	fmt.Printf("Architecture: %s v%s\n", runtime.GOARCH, runtime.GOOS)
//	fmt.Printf("GOROOT: %s\n", runtime.GOROOT())
//}
