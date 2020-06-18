package toolGo

import (
	"errors"
	"fmt"
	"github.com/blang/semver"
	"github.com/newclarity/scribeHelpers/ux"
)


type Version struct {
	version *semver.Version
}


func (v *Version) Set(version string) error {
	var err error
	for range onlyOnce {
		sver, err := semver.Parse(version)
		if err == nil {
			v.version = &sver
			break
		}
		err = errors.New(fmt.Sprintf("Invalid version '%s' - %v", version, err))
	}
	return err
}


func (v *Version) String() string {
	var ret string
	for range onlyOnce {
		ret += ux.SprintfBlue("%s: ", BinaryVersion)
		ret += ux.SprintfCyan("%v\n", v.version)
	}
	return ret
}


func (v *Version) Name() string {
	return v.version.String()
}


func (v *Version) Get() Version {
	return *v
}


func (v *Version) IsValid() bool {
	var ok bool
	for range onlyOnce {
		if v == nil {
			break
		}
		if v.version == nil {
			break
		}
		if err := v.version.Validate(); err != nil {
			break
		}
		ok = true
	}
	return ok
}


func (v *Version) IsNotValid() bool {
	return !v.IsValid()
}


func (v *Version) IsSame(compare *Version) bool {
	var ok bool
	for range onlyOnce {
		if v.version.String() != compare.version.String() {
			break
		}
		ok = true
	}
	return ok
}


func (v *Version) IsNotSame(compare *Version) bool {
	return !v.IsSame(compare)
}
