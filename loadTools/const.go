package loadTools

const (
	SelectStdout = "select:stdout"
	SelectConvert = "select:convert"
	SelectIgnore = "select:ignore"

	DefaultJsonString 		= "{}"
	DefaultTemplateString 	= "{{ .Exec.CmdName }}"

	DefaultJsonFile 		= "scribe.json"
	DefaultTemplateFile 	= "scribe.tmpl"
	DefaultOutFile 			= "/dev/stdout"
	DefaultWorkingPath	    = "."

	DefaultJsonFileSuffix 		= ".json"
	DefaultTemplateFileSuffix 	= ".tmpl"

	CmdRun 				= "run"
	CmdLoad 			= "load"
	CmdConvert 			= "convert"
	CmdTools 			= "tools"
	CmdVersion 			= "version"
	CmdBuild 			= "build"
	CmdPush 			= "push"
	CmdRelease 			= "release"
	CmdSelfUpdate		= "selfupdate"

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

	FlagVersion 		= "version"
	FlagHelpFunctions	= "help-functions"
	FlagHelpVariables	= "help-variables"
	FlagHelpExamples	= "help-examples"
	FlagHelpAll			= "help-all"
)
