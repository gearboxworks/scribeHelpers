// Code generated by github.com/newclarity/PackageReflect DO NOT EDIT.

package toolDocker

import "reflect"

var Types = map[string]reflect.Type{
	"Container": reflect.TypeOf((*Container)(nil)).Elem(),
	"Containers": reflect.TypeOf((*Containers)(nil)).Elem(),
	"TypeDocker": reflect.TypeOf((*TypeDocker)(nil)).Elem(),
	"Image": reflect.TypeOf((*Image)(nil)).Elem(),
	"PullEvent": reflect.TypeOf((*PullEvent)(nil)).Elem(),
	"TypeMatchContainer": reflect.TypeOf((*TypeMatchContainer)(nil)).Elem(),
	"TypeMatchImage": reflect.TypeOf((*TypeMatchImage)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"MatchTag": reflect.ValueOf(MatchTag),
	"New": reflect.ValueOf(New),
	"NewContainer": reflect.ValueOf(NewContainer),
	"NewImage": reflect.ValueOf(NewImage),
	"ParseHostURL": reflect.ValueOf(ParseHostURL),
}

var Variables = map[string]reflect.Value{
	"GetTools": reflect.ValueOf(&GetTools),
}

var Consts = map[string]reflect.Value{
	"DefaultNetwork": reflect.ValueOf(DefaultNetwork),
	"DefaultPathNone": reflect.ValueOf(DefaultPathNone),
	"DefaultProject": reflect.ValueOf(DefaultProject),
	"DefaultTimeout": reflect.ValueOf(DefaultTimeout),
	"ToolPrefix": reflect.ValueOf(ToolPrefix),
}

