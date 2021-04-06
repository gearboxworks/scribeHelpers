package toolGear

import "time"


const (
	DefaultTimeout = time.Second * 2
	DefaultOrganization = "gearboxworks"
	DefaultPathNone = "none"
	DefaultPathCwd = "cwd"
	DefaultPathHome = "home"
	DefaultPathEmpty = ""
	DefaultProvider = "docker"

	DefaultBrandName = "Gearbox"
	DefaultProject = "/home/gearbox/projects/default"
	DefaultTmpDir = "/tmp"
	DefaultNetwork = "gearboxnet"
	DefaultUnitTestCmd = "/etc/gearbox/unit-tests/run.sh"
	DefaultCommandName = "default"

	//LatestName = "latest"
)

type ExecCommand struct {
	Dir string
	File string
	FullPath string
	AsLink bool
}
var RunAs ExecCommand
