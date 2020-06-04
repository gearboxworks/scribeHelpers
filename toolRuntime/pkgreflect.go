// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package toolRuntime

import "reflect"

var Types = map[string]reflect.Type{
	"Environment": reflect.TypeOf((*Environment)(nil)).Elem(),
	"ExecArgs": reflect.TypeOf((*ExecArgs)(nil)).Elem(),
	"ExecEnv": reflect.TypeOf((*ExecEnv)(nil)).Elem(),
	"TypeRuntime": reflect.TypeOf((*TypeRuntime)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"New": reflect.ValueOf(New),
}

var Variables = map[string]reflect.Value{
	"GetTools": reflect.ValueOf(&GetTools),
}

var Consts = map[string]reflect.Value{
	"OnlyOnce": reflect.ValueOf(OnlyOnce),
	"ToolPrefix": reflect.ValueOf(ToolPrefix),
}

