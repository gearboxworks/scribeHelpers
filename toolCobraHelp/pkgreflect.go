// Code generated by github.com/newclarity/PackageReflect DO NOT EDIT.

package toolCobraHelp

import "reflect"

var Types = map[string]reflect.Type{
	"Cmds": reflect.TypeOf((*Cmds)(nil)).Elem(),
	"CobraGetter": reflect.TypeOf((*CobraGetter)(nil)).Elem(),
	"Example": reflect.TypeOf((*Example)(nil)).Elem(),
	"Examples": reflect.TypeOf((*Examples)(nil)).Elem(),
	"State": reflect.TypeOf((*State)(nil)).Elem(),
	"ToolCobra": reflect.TypeOf((*ToolCobra)(nil)).Elem(),
	"TypeCommands": reflect.TypeOf((*TypeCommands)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"New": reflect.ValueOf(New),
	"ReflectState": reflect.ValueOf(ReflectState),
	"ReflectToolCobra": reflect.ValueOf(ReflectToolCobra),
}

var Variables = map[string]reflect.Value{
	"GetTools": reflect.ValueOf(&GetTools),
}

var Consts = map[string]reflect.Value{
	"FlagHelpAll": reflect.ValueOf(FlagHelpAll),
	"FlagHelpExamples": reflect.ValueOf(FlagHelpExamples),
	"ToolPrefix": reflect.ValueOf(ToolPrefix),
}

