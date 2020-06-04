// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package toolService

import "reflect"

var Types = map[string]reflect.Type{
	"ServiceGetter": reflect.TypeOf((*ServiceGetter)(nil)).Elem(),
	"State": reflect.TypeOf((*State)(nil)).Elem(),
	"ToolService": reflect.TypeOf((*ToolService)(nil)).Elem(),
	"TypeService": reflect.TypeOf((*TypeService)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"New": reflect.ValueOf(New),
	"ReflectToolService": reflect.ValueOf(ReflectToolService),
}

var Variables = map[string]reflect.Value{
	"GetTools": reflect.ValueOf(&GetTools),
}

var Consts = map[string]reflect.Value{
	"OnlyOnce": reflect.ValueOf(OnlyOnce),
	"ToolPrefix": reflect.ValueOf(ToolPrefix),
}

