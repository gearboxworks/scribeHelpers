// Code generated by github.com/gearboxworks/buildtool DO NOT EDIT.

package loadTools

import "reflect"

var Types = map[string]reflect.Type{
	"Example": reflect.TypeOf((*Example)(nil)).Elem(),
	"Examples": reflect.TypeOf((*Examples)(nil)).Elem(),
	"FileInfo": reflect.TypeOf((*FileInfo)(nil)).Elem(),
	"Files": reflect.TypeOf((*Files)(nil)).Elem(),
	"JsonFile": reflect.TypeOf((*JsonFile)(nil)).Elem(),
	"JsonMap": reflect.TypeOf((*JsonMap)(nil)).Elem(),
	"PkgReflect": reflect.TypeOf((*PkgReflect)(nil)).Elem(),
	"ScribeFile": reflect.TypeOf((*ScribeFile)(nil)).Elem(),
	"SortedTools": reflect.TypeOf((*SortedTools)(nil)).Elem(),
	"TemplateFile": reflect.TypeOf((*TemplateFile)(nil)).Elem(),
	"Tool": reflect.TypeOf((*Tool)(nil)).Elem(),
	"Tools": reflect.TypeOf((*Tools)(nil)).Elem(),
	"TypeArgFile": reflect.TypeOf((*TypeArgFile)(nil)).Elem(),
	"TypeScribeArgs": reflect.TypeOf((*TypeScribeArgs)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"AddTools": reflect.ValueOf(AddTools),
	"ChangeSuffix": reflect.ValueOf(ChangeSuffix),
	"DiscoverTools": reflect.ValueOf(DiscoverTools),
	"New": reflect.ValueOf(New),
	"NewArgFile": reflect.ValueOf(NewArgFile),
	"NewJsonStruct": reflect.ValueOf(NewJsonStruct),
	"PackageReflect": reflect.ValueOf(PackageReflect),
	"PrintTools": reflect.ValueOf(PrintTools),
	"UnescapeString": reflect.ValueOf(UnescapeString),
}

var Variables = map[string]reflect.Value{
}

var Consts = map[string]reflect.Value{
	"CmdBuild": reflect.ValueOf(CmdBuild),
	"CmdConvert": reflect.ValueOf(CmdConvert),
	"CmdLoad": reflect.ValueOf(CmdLoad),
	"CmdPush": reflect.ValueOf(CmdPush),
	"CmdRelease": reflect.ValueOf(CmdRelease),
	"CmdRoot": reflect.ValueOf(CmdRoot),
	"CmdRun": reflect.ValueOf(CmdRun),
	"CmdTools": reflect.ValueOf(CmdTools),
	"DefaultJsonFile": reflect.ValueOf(DefaultJsonFile),
	"DefaultJsonFileSuffix": reflect.ValueOf(DefaultJsonFileSuffix),
	"DefaultJsonString": reflect.ValueOf(DefaultJsonString),
	"DefaultOutFile": reflect.ValueOf(DefaultOutFile),
	"DefaultPkgReflectFile": reflect.ValueOf(DefaultPkgReflectFile),
	"DefaultScribeFile": reflect.ValueOf(DefaultScribeFile),
	"DefaultScribeFileSuffix": reflect.ValueOf(DefaultScribeFileSuffix),
	"DefaultScribeString": reflect.ValueOf(DefaultScribeString),
	"DefaultTemplateFile": reflect.ValueOf(DefaultTemplateFile),
	"DefaultTemplateFileSuffix": reflect.ValueOf(DefaultTemplateFileSuffix),
	"DefaultTemplateString": reflect.ValueOf(DefaultTemplateString),
	"DefaultWorkingPath": reflect.ValueOf(DefaultWorkingPath),
	"FlagChdir": reflect.ValueOf(FlagChdir),
	"FlagConfigFile": reflect.ValueOf(FlagConfigFile),
	"FlagDebug": reflect.ValueOf(FlagDebug),
	"FlagForce": reflect.ValueOf(FlagForce),
	"FlagHelpAll": reflect.ValueOf(FlagHelpAll),
	"FlagHelpExamples": reflect.ValueOf(FlagHelpExamples),
	"FlagHelpFunctions": reflect.ValueOf(FlagHelpFunctions),
	"FlagHelpVariables": reflect.ValueOf(FlagHelpVariables),
	"FlagJsonFile": reflect.ValueOf(FlagJsonFile),
	"FlagOutputFile": reflect.ValueOf(FlagOutputFile),
	"FlagQuiet": reflect.ValueOf(FlagQuiet),
	"FlagRemoveOutput": reflect.ValueOf(FlagRemoveOutput),
	"FlagRemoveTemplate": reflect.ValueOf(FlagRemoveTemplate),
	"FlagScribeFile": reflect.ValueOf(FlagScribeFile),
	"FlagTemplateFile": reflect.ValueOf(FlagTemplateFile),
	"FlagVerbose": reflect.ValueOf(FlagVerbose),
	"FlagVersion": reflect.ValueOf(FlagVersion),
	"FlagWorkingPath": reflect.ValueOf(FlagWorkingPath),
	"SelectConvert": reflect.ValueOf(SelectConvert),
	"SelectDefault": reflect.ValueOf(SelectDefault),
	"SelectFile": reflect.ValueOf(SelectFile),
	"SelectIgnore": reflect.ValueOf(SelectIgnore),
	"SelectStdout": reflect.ValueOf(SelectStdout),
	"SelectString": reflect.ValueOf(SelectString),
	"ToolPrefix": reflect.ValueOf(ToolPrefix),
}

