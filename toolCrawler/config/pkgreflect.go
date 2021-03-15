// Code generated by github.com/newclarity/PackageReflect DO NOT EDIT.

package config

import "reflect"

var Types = map[string]reflect.Type{
	"Config": reflect.TypeOf((*Config)(nil)).Elem(),
	"DurationString": reflect.TypeOf((*DurationString)(nil)).Elem(),
	"UrlPattern": reflect.TypeOf((*UrlPattern)(nil)).Elem(),
	"UrlPatterns": reflect.TypeOf((*UrlPatterns)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"DefaultJson": reflect.ValueOf(DefaultJson),
	"LoadConfig": reflect.ValueOf(LoadConfig),
}

var Variables = map[string]reflect.Value{
}

var Consts = map[string]reflect.Value{
	"DefaultRevisit": reflect.ValueOf(DefaultRevisit),
	"Dir": reflect.ValueOf(Dir),
	"Filename": reflect.ValueOf(Filename),
	"InitialPause": reflect.ValueOf(InitialPause),
	"PauseIncrease": reflect.ValueOf(PauseIncrease),
	"TimeoutErr": reflect.ValueOf(TimeoutErr),
	"UnsuccessfulErr": reflect.ValueOf(UnsuccessfulErr),
}
