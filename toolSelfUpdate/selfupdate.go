package toolSelfUpdate

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"strings"
)

func (su *TypeSelfUpdate) Update() *ux.State {
	for range onlyOnce {
		su.State = su.IsValid()
		if su.State.IsNotOk() {
			break
		}


		ux.PrintflnBlue("Checking '%s' for version greater than v%s", su.useRepo, su.version.ToString())
		previous := su.version.ToSemVer()
		latest, err := su.ref.UpdateSelf(previous, su.useRepo)
		if err != nil {
			su.State.SetError(err)
			break
		}

		if previous.Equals(latest.Version) {
			ux.PrintflnOk("%s is up to date: v%s", su.name.ToString(), su.version.ToString())
		} else {
			ux.PrintflnOk("%s updated to v%s", su.name.ToString(), latest.Version)
			if latest.ReleaseNotes != "" {
				ux.PrintflnOk("%s %s Release Notes:\n%s", su.name.ToString(), latest.Version, latest.ReleaseNotes)
			}
		}
	}

	return su.State
}


func (su *TypeSelfUpdate) GetVersion(version *VersionValue) *selfupdate.Release {
	var release *selfupdate.Release

	for range onlyOnce {
		su.State = su.IsValid()
		if su.State.IsNotOk() {
			break
		}

		var ok bool
		var err error

		switch {
			case version.IsNotValid():
				fallthrough
			case version.IsLatest():
				release, ok, err = su.ref.DetectLatest(su.useRepo)

			default:
				v := version.ToString()
				if !strings.HasPrefix(v, "v") {
					v = "v" + v
				}
				release, ok, err = su.ref.DetectVersion(su.useRepo, v)
		}

		if !ok {
			su.State.SetWarning(errorNoVersion)
			break
		}
		if err != nil {
			su.State.SetWarning("%s - %s", errorNoVersion, err)
			break
		}

		su.State.SetOutput(release)
	}

	return release
}


func (su *TypeSelfUpdate) PrintVersion(version *VersionValue) *ux.State {
	for range onlyOnce {
		su.State = su.IsValid()
		if su.State.IsNotOk() {
			break
		}

		release := su.GetVersion(version)
		if su.State.IsNotOk() {
			break
		}

		fmt.Printf(printVersion(release))
	}

	return su.State
}


func (su *TypeSelfUpdate) IsUpdated(print bool) *ux.State {
	for range onlyOnce {
		su.State = su.IsValid()
		if su.State.IsNotOk() {
			break
		}

		latest := su.GetVersion(nil)
		if su.State.IsNotOk() {
			break
		}

		current := su.GetVersion(su.version)

		if current == nil {
			su.State.SetWarning("%s can be updated to v%s.",
				su.name.ToString(),
				su.version.ToString())
			if print {
				ux.PrintflnWarning("Current version info unknown.")
				ux.PrintflnBlue("Current version (v%s)\n", su.version.ToString())
				ux.PrintflnBlue("Updated version (v%s)", latest.Version.String())
				fmt.Printf(printVersion(latest))
			}
			su.State.Clear()
			break
		}

		if current.Version.Equals(latest.Version) {
			su.State.SetOk("%s is up to date at v%s.",
				su.name.ToString(),
				su.version.ToString())
			if print {
				ux.PrintflnOk("%s", su.State.GetOk())
				fmt.Printf(printVersion(current))
			}
			break
		}

		if current.Version.LE(latest.Version) {
			su.State.SetWarning("%s can be updated to v%s.",
				su.name.ToString(),
				su.version.ToString())
			if print {
				ux.PrintflnWarning("%s", su.State.GetOk())
				ux.PrintflnBlue("Current version (v%s)", current.Version.String())
				fmt.Printf(printVersion(current))
				ux.PrintflnBlue("Updated version (v%s)", latest.Version.String())
				fmt.Printf(printVersion(latest))
			}
			break
		}

		if current.Version.GT(latest.Version) {
			su.State.SetWarning("%s is more recent at v%s, (latest is %s).",
				su.name.ToString(),
				su.version.ToString(),
				latest.Version.String())
			if print {
				ux.PrintflnWarning("%s", su.State.GetOk())
				ux.PrintflnBlue("Current version (v%s)", current.Version.String())
				fmt.Printf(printVersion(current))
				ux.PrintflnBlue("Updated version (v%s)", latest.Version.String())
				fmt.Printf(printVersion(latest))
			}
			break
		}
	}

	return su.State
}


func (su *TypeSelfUpdate) Set(s SelfUpdateArgs) *ux.State {
	if s.name != nil {
		su.name = (*StringValue)(s.name)
	}

	if s.version != nil {
		su.version = toVersionValue(*s.version)
	}

	if s.binaryRepo != nil {
		su.binaryRepo = (*StringValue)(s.binaryRepo)
	}

	if s.sourceRepo != nil {
		su.sourceRepo = (*StringValue)(s.sourceRepo)
	}

	if s.logging != nil {
		su.logging = (*FlagValue)(s.logging)
	} else {
		su.logging = &defaultFalse
	}

	su.State = su.IsValid()

	return su.State
}


func (su *TypeSelfUpdate) SetDebug(value bool) *ux.State {
	su.logging = (*FlagValue)(&value)
	su.State = su.IsValid()
	return su.State
}


func (su *TypeSelfUpdate) SetName(value string) *ux.State {
	su.name = (*StringValue)(&value)
	su.State = su.IsValid()
	return su.State
}


func (su *TypeSelfUpdate) SetVersion(value string) *ux.State {
	su.version = toVersionValue(value)
	su.State = su.IsValid()
	return su.State
}


func (su *TypeSelfUpdate) SetSourceRepo(value string) *ux.State {
	su.sourceRepo = (*StringValue)(&value)
	su.State = su.IsValid()
	return su.State
}


func (su *TypeSelfUpdate) SetBinaryRepo(value string) *ux.State {
	su.binaryRepo = (*StringValue)(&value)
	su.State = su.IsValid()
	return su.State
}
