package toolGear

import "time"

//goland:noinspection SpellCheckingInspection
const (
	DefaultTimeout = time.Second * 10
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

const (
	ErrorDockerTimeout = "context deadline exceeded"
	DefaultMinTimeout = 5
	DefaultMaxTimeout = 30
)

type ExecCommand struct {
	Dir string
	File string
	FullPath string
	AsLink bool
}
var RunAs ExecCommand
