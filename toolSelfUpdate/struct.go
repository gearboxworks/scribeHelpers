package toolSelfUpdate

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"runtime"
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

	config     *selfupdate.Config
	ref        *selfupdate.Updater

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


func New(rt *toolRuntime.TypeRuntime) *TypeSelfUpdate {
	rt = rt.EnsureNotNil()

	te := TypeSelfUpdate{
		name:       toStringValue(rt.CmdName),
		version:    toVersionValue(rt.CmdVersion),
		sourceRepo: toStringValue(stripUrlPrefix(rt.CmdSourceRepo)),
		binaryRepo: toStringValue(stripUrlPrefix(rt.CmdBinaryRepo)),
		logging:    toBoolValue(rt.Debug),

		config:     &selfupdate.Config{
			APIToken:            "",
			EnterpriseBaseURL:   "",
			EnterpriseUploadURL: "",
			Validator:           nil, 	// &MyValidator{},
			Filters:             []string{},
		},

		useRepo:    "",

		runtime: rt,
		State:   ux.NewState(rt.CmdName, rt.Debug),
	}
	te.State.SetPackage("")
	te.State.SetFunctionCaller()

	// Workaround for selfupdate not being flexible enough to support variable asset names
	// Should enable a template similar to GoReleaser.
	// EG: {{ .ProjectName }}-{{ .Os }}_{{ .Arch }}
	//var asset string
	//asset, te.State = toolGhr.GetAsset(rt.CmdBinaryRepo, "latest")
	//te.config.Filters = append(te.config.Filters, asset)

	// Ignore the above and just make sure all filenames are lowercase.
	te.config.Filters = append(te.config.Filters, addFilters(rt.CmdName, runtime.GOOS, runtime.GOARCH)...)
	te.ref, _ = selfupdate.NewUpdater(*te.config)
	if *te.logging {
		selfupdate.EnableLog()
	}

	return &te
}


//type MyValidator struct {
//}
//func (v *MyValidator) Validate(release, asset []byte) error {
//	calculatedHash := fmt.Sprintf("%x", sha256.Sum256(release))
//	hash := fmt.Sprintf("%s", asset[:sha256.BlockSize])
//	if calculatedHash != hash {
//		return fmt.Errorf("sha2: validation failed: hash mismatch: expected=%q, got=%q", calculatedHash, hash)
//	}
//	return nil
//}
//func (v *MyValidator) Suffix() string {
//	return ".gz"
//}


func addFilters(Binary string, Os string, Arch string) []string {
	var ret []string
	ret = append(ret, fmt.Sprintf("(?i)%s_.*_%s_%s.*", Binary, Os, Arch))
	ret = append(ret, fmt.Sprintf("(?i)%s_%s_%s.*", Binary, Os, Arch))
	ret = append(ret, fmt.Sprintf("(?i)%s-.*_%s_%s.*", Binary, Os, Arch))
	ret = append(ret, fmt.Sprintf("(?i)%s-%s_%s.*", Binary, Os, Arch))
	if Arch == "amd64" {
		// This is recursive - so be careful what you place in the "Arch" argument.
		ret = append(ret, addFilters(Binary, Os, "x86_64.*")...)
		ret = append(ret, addFilters(Binary, Os, "64.*")...)
		ret = append(ret, addFilters(Binary, Os, "64bit.*")...)
	}
	return ret
}


func (su *TypeSelfUpdate) IsValid() *ux.State {
	for range onlyOnce {
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
