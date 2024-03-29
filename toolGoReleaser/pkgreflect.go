// Code generated by github.com/gearboxworks/buildtool DO NOT EDIT.

package toolGoReleaser

import "reflect"

var Types = map[string]reflect.Type{
	"GoReleaserGetter": reflect.TypeOf((*GoReleaserGetter)(nil)).Elem(),
	"State": reflect.TypeOf((*State)(nil)).Elem(),
	"ToolGoReleaser": reflect.TypeOf((*ToolGoReleaser)(nil)).Elem(),
	"TypeGoReleaser": reflect.TypeOf((*TypeGoReleaser)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"New": reflect.ValueOf(New),
	"ReflectState": reflect.ValueOf(ReflectState),
	"ReflectToolGoReleaser": reflect.ValueOf(ReflectToolGoReleaser),
}

var Variables = map[string]reflect.Value{
	"GetTools": reflect.ValueOf(&GetTools),
}

var Consts = map[string]reflect.Value{
	"DefaultCmd": reflect.ValueOf(DefaultCmd),
	"DefaultFile": reflect.ValueOf(DefaultFile),
	"ToolPrefix": reflect.ValueOf(ToolPrefix),
}

