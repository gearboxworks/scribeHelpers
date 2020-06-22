package toolSelfUpdate

const (
	errorNoRepo = "repo is not defined - selfupdate disabled"
	errorNoVersion = "no versions in repo - selfupdate disabled"
	LatestVersion = "latest"

	CmdSelfUpdate		= "selfupdate"

	CmdVersion 			= "version"
	CmdVersionInfo		= "info"
	CmdVersionList		= "list"
	CmdVersionLatest	= "latest"
	CmdVersionCheck		= "check"
	CmdVersionUpdate	= "update"

	FlagVersion 		= "version"
)

var defaultFalse = FlagValue(false)
