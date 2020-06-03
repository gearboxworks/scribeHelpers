// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package loadHelpers

import "reflect"

var Types = map[string]reflect.Type{
	"FileInfo": reflect.TypeOf((*FileInfo)(nil)).Elem(),
	"Files": reflect.TypeOf((*Files)(nil)).Elem(),
	"Helper": reflect.TypeOf((*Helper)(nil)).Elem(),
	"Helpers": reflect.TypeOf((*Helpers)(nil)).Elem(),
	"PkgReflect": reflect.TypeOf((*PkgReflect)(nil)).Elem(),
	"SortedHelpers": reflect.TypeOf((*SortedHelpers)(nil)).Elem(),
	"TypeArgFile": reflect.TypeOf((*TypeArgFile)(nil)).Elem(),
	"TypeScribeArgs": reflect.TypeOf((*TypeScribeArgs)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"AddHelpers": reflect.ValueOf(AddHelpers),
	"DiscoverHelpers": reflect.ValueOf(DiscoverHelpers),
	"New": reflect.ValueOf(New),
	"NewJsonStruct": reflect.ValueOf(NewJsonStruct),
	"PackageReflect": reflect.ValueOf(PackageReflect),
	"PrintHelpers": reflect.ValueOf(PrintHelpers),
	"UnescapeString": reflect.ValueOf(UnescapeString),
}

var Variables = map[string]reflect.Value{
}

var Consts = map[string]reflect.Value{
	"CmdBuild": reflect.ValueOf(CmdBuild),
	"CmdConvert": reflect.ValueOf(CmdConvert),
	"CmdHelpers": reflect.ValueOf(CmdHelpers),
	"CmdLoad": reflect.ValueOf(CmdLoad),
	"CmdPush": reflect.ValueOf(CmdPush),
	"CmdRelease": reflect.ValueOf(CmdRelease),
	"CmdRun": reflect.ValueOf(CmdRun),
	"CmdVersion": reflect.ValueOf(CmdVersion),
	"DefaultJsonFile": reflect.ValueOf(DefaultJsonFile),
	"DefaultJsonFileSuffix": reflect.ValueOf(DefaultJsonFileSuffix),
	"DefaultJsonString": reflect.ValueOf(DefaultJsonString),
	"DefaultOutFile": reflect.ValueOf(DefaultOutFile),
	"DefaultTemplateFile": reflect.ValueOf(DefaultTemplateFile),
	"DefaultTemplateFileSuffix": reflect.ValueOf(DefaultTemplateFileSuffix),
	"DefaultTemplateString": reflect.ValueOf(DefaultTemplateString),
	"FlagChdir": reflect.ValueOf(FlagChdir),
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
	"FlagTemplateFile": reflect.ValueOf(FlagTemplateFile),
	"FlagVersion": reflect.ValueOf(FlagVersion),
	"HelperPrefix": reflect.ValueOf(HelperPrefix),
	"OnlyOnce": reflect.ValueOf(OnlyOnce),
	"SelectConvert": reflect.ValueOf(SelectConvert),
	"SelectStdout": reflect.ValueOf(SelectStdout),
}

