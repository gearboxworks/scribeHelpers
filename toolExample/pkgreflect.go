// Code generated by github.com/gearboxworks/buildtool DO NOT EDIT.

package toolExample

import "reflect"

var Types = map[string]reflect.Type{
	"ExampleGetter": reflect.TypeOf((*ExampleGetter)(nil)).Elem(),
	"State": reflect.TypeOf((*State)(nil)).Elem(),
	"ToolExample": reflect.TypeOf((*ToolExample)(nil)).Elem(),
	"TypeExample": reflect.TypeOf((*TypeExample)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"New": reflect.ValueOf(New),
	"ReflectState": reflect.ValueOf(ReflectState),
	"ReflectToolExample": reflect.ValueOf(ReflectToolExample),
}

var Variables = map[string]reflect.Value{
	"GetTools": reflect.ValueOf(&GetTools),
}

var Consts = map[string]reflect.Value{
	"ToolPrefix": reflect.ValueOf(ToolPrefix),
}

