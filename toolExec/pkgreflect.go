// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package toolExec

import "reflect"

var Types = map[string]reflect.Type{
	"ToolExecCommand": reflect.TypeOf((*ToolExecCommand)(nil)).Elem(),
	"TypeExecCommand": reflect.TypeOf((*TypeExecCommand)(nil)).Elem(),
	"TypeExecCommandGetter": reflect.TypeOf((*TypeExecCommandGetter)(nil)).Elem(),
	"TypeMultiExecCommand": reflect.TypeOf((*TypeMultiExecCommand)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"New": reflect.ValueOf(New),
	"NewMultiExec": reflect.ValueOf(NewMultiExec),
	"ReflectExecCommand": reflect.ValueOf(ReflectExecCommand),
	"ToolExec": reflect.ValueOf(ToolExec),
	"ToolExecBash": reflect.ValueOf(ToolExecBash),
	"ToolExecCmd": reflect.ValueOf(ToolExecCmd),
	"ToolNewBash": reflect.ValueOf(ToolNewBash),
	"ToolOsExit": reflect.ValueOf(ToolOsExit),
}

var Variables = map[string]reflect.Value{
	"GetTools": reflect.ValueOf(&GetTools),
}

var Consts = map[string]reflect.Value{
	"ToolPrefix": reflect.ValueOf(ToolPrefix),
}

