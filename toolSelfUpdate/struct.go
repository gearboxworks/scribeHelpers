package toolSelfUpdate

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

type SelfUpdateGetter interface {
}

type SelfUpdateArgs struct {
	name       *string
	version    *string
	sourceRepo *string
	binaryRepo *string

	logging    *bool
}

type TypeSelfUpdate struct {
	name       *StringValue
	version    *VersionValue
	sourceRepo *StringValue
	binaryRepo *StringValue
	logging    *FlagValue

	useRepo    string

	runtime    *toolRuntime.TypeRuntime
	State      *ux.State
}


type state ux.State
func (p *state) Reflect() *ux.State {
	return (*ux.State)(p)
}
func ReflectToolSelfUpdate(p *TypeSelfUpdate) *ToolSelfUpdate {
	return (*ToolSelfUpdate)(p)
}

type ToolSelfUpdate TypeSelfUpdate
func (su *ToolSelfUpdate) Reflect() *TypeSelfUpdate {
	return (*TypeSelfUpdate)(su)
}

func (su *TypeSelfUpdate) IsNil() *ux.State {
	if state := ux.IfNilReturnError(su); state.IsError() {
		return state
	}
	su.State = su.State.EnsureNotNil()
	return su.State
}


func New(runtime *toolRuntime.TypeRuntime) *TypeSelfUpdate {
	runtime = runtime.EnsureNotNil()

	te := TypeSelfUpdate{
		name:       toStringValue(runtime.CmdName),
		version:    toVersionValue(runtime.CmdVersion),
		sourceRepo: toStringValue(stripUrlPrefix(runtime.CmdSourceRepo)),
		binaryRepo: toStringValue(stripUrlPrefix(runtime.CmdBinaryRepo)),
		logging:    toBoolValue(runtime.Debug),
		runtime:    runtime,
		State:      ux.NewState(runtime.CmdName, runtime.Debug),
	}
	te.State.SetPackage("")
	te.State.SetFunctionCaller()

	return &te
}


func (su *TypeSelfUpdate) IsValid() *ux.State {
	for range onlyOnce {
		if *su.logging {
			selfupdate.EnableLog()
		}

		if su.name.IsNotValid() {
			su.State.SetWarning("binary name is not defined - selfupdate disabled")
			break
		}

		if su.version.IsNotValid() {
			su.State.SetWarning("binary version is not defined - selfupdate disabled")
			break
		}

		// Refer to binary repo definition first.
		if su.binaryRepo.IsValid() {
			su.useRepo = su.binaryRepo.ToString()
			su.State.SetOk()
			break
		}

		// If binary repo is not set, use source repo.
		if su.sourceRepo.IsValid() {
			su.useRepo = su.sourceRepo.ToString()
			su.State.SetOk()
			break
		}

		su.State.SetWarning(errorNoRepo)
	}

	return su.State
}


func (su *TypeSelfUpdate) getRepo() string {
	var ret string

	for range onlyOnce {
		if su.binaryRepo.IsValid() {
			ret = su.binaryRepo.ToString()
			break
		}
		if su.sourceRepo.IsValid() {
			ret = su.sourceRepo.ToString()
			break
		}
	}

	return ret
}
