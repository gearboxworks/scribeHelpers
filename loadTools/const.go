package loadTools

const (
	SelectStdout	= "select:stdout"
	SelectConvert	= "select:convert"
	SelectIgnore	= "select:ignore"
	SelectString	= "select:string"
	SelectFile		= "select:file"
	SelectDefault	= "select:default"

	DefaultScribeFile 			= "default.scribe"
	DefaultScribeFileSuffix 	= ".scribe"
	DefaultScribeString			= ".Exec.CmdName"

	DefaultJsonFile 			= "default.json"
	DefaultJsonFileSuffix 		= ".json"
	DefaultJsonString 			= "{}"

	DefaultTemplateFile 		= "default.tmpl"
	DefaultTemplateFileSuffix 	= ".tmpl"
	DefaultTemplateString 		= "{{ .Exec.CmdName }}"

	DefaultOutFile 				= "/dev/stdout"
	DefaultWorkingPath	    	= "."


	CmdRun 				= "run"
	CmdLoad 			= "load"
	CmdConvert 			= "convert"
	CmdTools 			= "tools"
	CmdBuild 			= "build"
	CmdPush 			= "push"
	CmdRelease 			= "release"

	CmdSelfUpdate		= "selfupdate"
	CmdVersion 			= "version"
	CmdVersionInfo		= "info"
	CmdVersionList		= "list"
	CmdVersionLatest	= "latest"
	CmdVersionCheck		= "check"
	CmdVersionUpdate	= "update"

	FlagScribeFile     	= "scribe"
	FlagJsonFile     	= "json"
	FlagTemplateFile	= "template"
	FlagOutputFile	    = "out"
	FlagWorkingPath	    = "path"

	FlagChdir       	= "chdir"
	FlagForce 			= "force"
	FlagRemoveTemplate	= "rm-tmpl"
	FlagRemoveOutput	= "rm-out"
	FlagDebug 			= "debug"
	FlagQuiet			= "quiet"
	FlagVerbose			= "verbose"

	FlagVersion 		= "version"
	FlagHelpFunctions	= "help-functions"
	FlagHelpVariables	= "help-variables"
	FlagHelpExamples	= "help-examples"
	FlagHelpAll			= "help-all"
)
