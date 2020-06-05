// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package toolGear

import "reflect"

var Types = map[string]reflect.Type{
	"Container": reflect.TypeOf((*Container)(nil)).Elem(),
	"Containers": reflect.TypeOf((*Containers)(nil)).Elem(),
	"DockerGear": reflect.TypeOf((*DockerGear)(nil)).Elem(),
	"ExecCommand": reflect.TypeOf((*ExecCommand)(nil)).Elem(),
	"Gear": reflect.TypeOf((*Gear)(nil)).Elem(),
	"GitHubRepo": reflect.TypeOf((*GitHubRepo)(nil)).Elem(),
	"Image": reflect.TypeOf((*Image)(nil)).Elem(),
	"Provider": reflect.TypeOf((*Provider)(nil)).Elem(),
	"PullEvent": reflect.TypeOf((*PullEvent)(nil)).Elem(),
	"Release": reflect.TypeOf((*Release)(nil)).Elem(),
	"ReleaseSelector": reflect.TypeOf((*ReleaseSelector)(nil)).Elem(),
	"ReleasesMap": reflect.TypeOf((*ReleasesMap)(nil)).Elem(),
	"SshfsMounts": reflect.TypeOf((*SshfsMounts)(nil)).Elem(),
	"TypeMatchContainer": reflect.TypeOf((*TypeMatchContainer)(nil)).Elem(),
	"TypeMatchImage": reflect.TypeOf((*TypeMatchImage)(nil)).Elem(),
	"Version": reflect.TypeOf((*Version)(nil)).Elem(),
	"VolumeMounts": reflect.TypeOf((*VolumeMounts)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"MatchContainer": reflect.ValueOf(MatchContainer),
	"MatchImage": reflect.ValueOf(MatchImage),
	"MatchTag": reflect.ValueOf(MatchTag),
	"New": reflect.ValueOf(New),
	"NewContainer": reflect.ValueOf(NewContainer),
	"NewGear": reflect.ValueOf(NewGear),
	"NewGearConfig": reflect.ValueOf(NewGearConfig),
	"NewImage": reflect.ValueOf(NewImage),
	"NewProvider": reflect.ValueOf(NewProvider),
	"NewRepo": reflect.ValueOf(NewRepo),
	"ParseHostURL": reflect.ValueOf(ParseHostURL),
}

var Variables = map[string]reflect.Value{
	"GetTools": reflect.ValueOf(&GetTools),
	"RunAs": reflect.ValueOf(&RunAs),
}

var Consts = map[string]reflect.Value{
	"Brandname": reflect.ValueOf(Brandname),
	"DefaultBrandName": reflect.ValueOf(DefaultBrandName),
	"DefaultCommandName": reflect.ValueOf(DefaultCommandName),
	"DefaultNetwork": reflect.ValueOf(DefaultNetwork),
	"DefaultOrganization": reflect.ValueOf(DefaultOrganization),
	"DefaultPathCwd": reflect.ValueOf(DefaultPathCwd),
	"DefaultPathEmpty": reflect.ValueOf(DefaultPathEmpty),
	"DefaultPathHome": reflect.ValueOf(DefaultPathHome),
	"DefaultPathNone": reflect.ValueOf(DefaultPathNone),
	"DefaultProject": reflect.ValueOf(DefaultProject),
	"DefaultProvider": reflect.ValueOf(DefaultProvider),
	"DefaultTimeout": reflect.ValueOf(DefaultTimeout),
	"DefaultTmpDir": reflect.ValueOf(DefaultTmpDir),
	"DefaultUnitTestCmd": reflect.ValueOf(DefaultUnitTestCmd),
	"IsoFileDownloaded": reflect.ValueOf(IsoFileDownloaded),
	"IsoFileIsDownloading": reflect.ValueOf(IsoFileIsDownloading),
	"IsoFileNeedsToDownload": reflect.ValueOf(IsoFileNeedsToDownload),
	"onlyOnce": reflect.ValueOf(onlyOnce),
	"ProviderDocker": reflect.ValueOf(ProviderDocker),
	"ToolPrefix": reflect.ValueOf(ToolPrefix),
}

