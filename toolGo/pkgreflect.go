// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package toolGo

import "reflect"

var Types = map[string]reflect.Type{
	"GoGetter": reflect.TypeOf((*GoGetter)(nil)).Elem(),
	"State": reflect.TypeOf((*State)(nil)).Elem(),
	"ToolGo": reflect.TypeOf((*ToolGo)(nil)).Elem(),
	"TypeGo": reflect.TypeOf((*TypeGo)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"New": reflect.ValueOf(New),
	"ReflectToolGo": reflect.ValueOf(ReflectToolGo),
}

var Variables = map[string]reflect.Value{
	"GetTools": reflect.ValueOf(&GetTools),
}

var Consts = map[string]reflect.Value{
	"ToolPrefix": reflect.ValueOf(ToolPrefix),
}

