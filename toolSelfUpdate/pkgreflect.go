// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package toolSelfUpdate

import "reflect"

var Types = map[string]reflect.Type{
	"flagValue": reflect.TypeOf((*flagValue)(nil)).Elem(),
	"SelfUpdateArgs": reflect.TypeOf((*SelfUpdateArgs)(nil)).Elem(),
	"SelfUpdateGetter": reflect.TypeOf((*SelfUpdateGetter)(nil)).Elem(),
	"state": reflect.TypeOf((*state)(nil)).Elem(),
	"stringValue": reflect.TypeOf((*stringValue)(nil)).Elem(),
	"ToolSelfUpdate": reflect.TypeOf((*ToolSelfUpdate)(nil)).Elem(),
	"TypeSelfUpdate": reflect.TypeOf((*TypeSelfUpdate)(nil)).Elem(),
	"versionValue": reflect.TypeOf((*versionValue)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"New": reflect.ValueOf(New),
	"ReflectFlagValue": reflect.ValueOf(ReflectFlagValue),
	"ReflectStringValue": reflect.ValueOf(ReflectStringValue),
	"ReflectToolSelfUpdate": reflect.ValueOf(ReflectToolSelfUpdate),
	"ReflectVersionValue": reflect.ValueOf(ReflectVersionValue),
}

var Variables = map[string]reflect.Value{
	"GetTools": reflect.ValueOf(&GetTools),
}

var Consts = map[string]reflect.Value{
	"LatestVersion": reflect.ValueOf(LatestVersion),
	"onlyOnce": reflect.ValueOf(onlyOnce),
	"ToolPrefix": reflect.ValueOf(ToolPrefix),
}

