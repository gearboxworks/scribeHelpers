package toolSelfUpdate

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

func (su *TypeSelfUpdate) Update() *ux.State {
	for range onlyOnce {
		su.State = su.IsValid()
		if su.State.IsNotOk() {
			break
		}


		ux.PrintflnBlue("Checking '%s' for version greater than v%s", su.useRepo, su.version.ToString())
		previous := su.version.ToSemVer()
		latest, err := selfupdate.UpdateSelf(previous, su.useRepo)
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


func (su *TypeSelfUpdate) GetVersion(version *versionValue) *selfupdate.Release {
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
				release, ok, err = selfupdate.DetectLatest(su.useRepo)

			default:
				release, ok, err = selfupdate.DetectVersion(su.useRepo, version.ToString())
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


func (su *TypeSelfUpdate) PrintVersion(version *versionValue) *ux.State {
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


		current := su.GetVersion(su.version)
		if su.State.IsNotOk() {
			break
		}

		latest := su.GetVersion(nil)
		if su.State.IsNotOk() {
			break
		}

		if current.Version.Equals(latest.Version) {
			su.State.SetOk("%s is up to date at v%s.",
				su.name.ToString(),
				su.version.ToString())
			if print {
				ux.PrintflnOk("%s", su.State.GetOk())
				fmt.Printf(printVersion(current))
				break
			}
			break
		}

		if current.Version.LE(latest.Version) {
			su.State.SetWarning("%s can be updated to v%s.",
				su.name.ToString(),
				su.version.ToString())
			if print {
				ux.PrintflnWarning("%s", su.State.GetOk())
				ux.PrintflnBlue("Current version (v)", current.Version.String())
				fmt.Printf(printVersion(current))
				ux.PrintflnBlue("Updated version (v)", latest.Version.String())
				fmt.Printf(printVersion(latest))
				break
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
				ux.PrintflnBlue("Current version (v)", current.Version.String())
				fmt.Printf(printVersion(current))
				ux.PrintflnBlue("Updated version (v)", latest.Version.String())
				fmt.Printf(printVersion(latest))
				break
			}
			break
		}
	}

	return su.State
}


func (su *TypeSelfUpdate) Set(s SelfUpdateArgs) *ux.State {
	if s.name != nil {
		su.name = (*stringValue)(s.name)
	}

	if s.version != nil {
		su.version = toVersionValue(*s.version)
	}

	if s.binaryRepo != nil {
		su.binaryRepo = (*stringValue)(s.binaryRepo)
	}

	if s.sourceRepo != nil {
		su.sourceRepo = (*stringValue)(s.sourceRepo)
	}

	if s.logging != nil {
		su.logging = (*flagValue)(s.logging)
	} else {
		su.logging = &defaultFalse
	}

	su.State = su.IsValid()

	return su.State
}


func (su *TypeSelfUpdate) SetDebug(value bool) *ux.State {
	su.logging = (*flagValue)(&value)
	su.State = su.IsValid()
	return su.State
}


func (su *TypeSelfUpdate) SetName(value string) *ux.State {
	su.name = (*stringValue)(&value)
	su.State = su.IsValid()
	return su.State
}


func (su *TypeSelfUpdate) SetVersion(value string) *ux.State {
	su.version = toVersionValue(value)
	su.State = su.IsValid()
	return su.State
}


func (su *TypeSelfUpdate) SetSourceRepo(value string) *ux.State {
	su.sourceRepo = (*stringValue)(&value)
	su.State = su.IsValid()
	return su.State
}


func (su *TypeSelfUpdate) SetBinaryRepo(value string) *ux.State {
	su.binaryRepo = (*stringValue)(&value)
	su.State = su.IsValid()
	return su.State
}
