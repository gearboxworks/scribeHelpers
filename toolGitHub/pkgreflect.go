// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package toolGitHub

import "reflect"

var Types = map[string]reflect.Type{
	"TypeGetRepositories": reflect.TypeOf((*TypeGetRepositories)(nil)).Elem(),
	"TypeGetRepository": reflect.TypeOf((*TypeGetRepository)(nil)).Elem(),
	"TypeGitHub": reflect.TypeOf((*TypeGitHub)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"New": reflect.ValueOf(New),
	"ToolGitHubGetBranch": reflect.ValueOf(ToolGitHubGetBranch),
	"ToolGitHubGetOrganization": reflect.ValueOf(ToolGitHubGetOrganization),
	"ToolGitHubLogin": reflect.ValueOf(ToolGitHubLogin),
}

var Variables = map[string]reflect.Value{
	"GetTools": reflect.ValueOf(&GetTools),
}

var Consts = map[string]reflect.Value{
	"onlyOnce": reflect.ValueOf(onlyOnce),
	"ToolPrefix": reflect.ValueOf(ToolPrefix),
}

