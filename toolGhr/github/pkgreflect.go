// Code generated by github.com/newclarity/PackageReflect DO NOT EDIT.

package github

import "reflect"

var Types = map[string]reflect.Type{
	"Client": reflect.TypeOf((*Client)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"DoAuthRequest": reflect.ValueOf(DoAuthRequest),
	"GetFileSize": reflect.ValueOf(GetFileSize),
	"NewClient": reflect.ValueOf(NewClient),
}

var Variables = map[string]reflect.Value{
	"VERBOSITY": reflect.ValueOf(&VERBOSITY),
}

var Consts = map[string]reflect.Value{
	"DefaultBaseURL": reflect.ValueOf(DefaultBaseURL),
}

