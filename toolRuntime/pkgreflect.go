// Code generated by github.com/gearboxworks/buildtool DO NOT EDIT.

package toolRuntime

import "reflect"

var Types = map[string]reflect.Type{
	"Environment": reflect.TypeOf((*Environment)(nil)).Elem(),
	"ExecArgs": reflect.TypeOf((*ExecArgs)(nil)).Elem(),
	"ExecEnv": reflect.TypeOf((*ExecEnv)(nil)).Elem(),
	"GoRuntime": reflect.TypeOf((*GoRuntime)(nil)).Elem(),
	"State": reflect.TypeOf((*State)(nil)).Elem(),
	"ToolRuntime": reflect.TypeOf((*ToolRuntime)(nil)).Elem(),
	"TypeRuntime": reflect.TypeOf((*TypeRuntime)(nil)).Elem(),
	"User": reflect.TypeOf((*User)(nil)).Elem(),
	"VersionValue": reflect.TypeOf((*VersionValue)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"New": reflect.ValueOf(New),
	"ReflectState": reflect.ValueOf(ReflectState),
	"ReflectToolRuntime": reflect.ValueOf(ReflectToolRuntime),
}

var Variables = map[string]reflect.Value{
	"GetTools": reflect.ValueOf(&GetTools),
}

var Consts = map[string]reflect.Value{
	"ToolPrefix": reflect.ValueOf(ToolPrefix),
}

