package toolSelfUpdate

import "github.com/blang/semver"

type stringValue string
type versionValue semver.Version
type flagValue bool


func ReflectStringValue(ref interface{}) *stringValue {
	var ret *stringValue
	switch ref.(type) {
		case *[]byte:
			ret = ref.(*stringValue)
		case *string:
			ret = ref.(*stringValue)
		case []byte:
			ret = ref.(*stringValue)
		case string:
			ret = ref.(*stringValue)
	}
	return ret
}


func ReflectVersionValue(ref interface{}) *versionValue {
	var ret versionValue
	switch ref.(type) {
		case *[]byte:
			ret = versionValue(semver.MustParse(*(ref.(*string))))
		case *string:
			ret = versionValue(semver.MustParse(*(ref.(*string))))
		case []byte:
			ret = versionValue(semver.MustParse((ref.(string))))
		case string:
			ret = versionValue(semver.MustParse((ref.(string))))
	}
	return &ret
}


func ReflectFlagValue(ref interface{}) *flagValue {
	var ret *flagValue
	switch ref.(type) {
		case *bool:
			ret = ref.(*flagValue)
	}
	return ret
}


func (v *versionValue) ToString() string {
	return (semver.Version)(*v).String()
}
func toVersionValue(version string) *versionValue {
	v := versionValue(semver.MustParse(version))
	return &v
}
func toSemVer(version string) semver.Version {
	return semver.MustParse(version)
}
func (v *versionValue) ToSemVer() semver.Version {
	return semver.Version(*v)
}
func (v *versionValue) IsValid() bool {
	var ok bool
	for range onlyOnce {
		if v == nil {
			break
		}

		err := (semver.Version)(*v).Validate()
		if err != nil {
			break
		}

		ok = true
	}
	return ok
}
func (v *versionValue) IsNotValid() bool {
	return !v.IsValid()
}
func (v *versionValue) IsLatest() bool {
	return (semver.Version)(*v).String() == LatestVersion
}


func (v *stringValue) ToString() string {
	return string(*v)
}
func toStringValue(s string) *stringValue {
	v := stringValue(s)
	return &v
}

func toBoolValue(b bool) *flagValue {
	v := flagValue(b)
	return &v
}


func (v *stringValue) IsValid() bool {
	var ok bool
	for range onlyOnce {
		if v == nil {
			break
		}
		if *v == "" {
			break
		}
		ok = true
	}
	return ok
}
func (v *stringValue) IsNotValid() bool {
	return !v.IsValid()
}


func (v *stringValue) IsNil() bool {
	if v == nil {
		return true
	}
	return false
}
func (v *stringValue) IsNotNil() bool {
	return !v.IsNil()
}


func (v *stringValue) IsEmpty() bool {
	if v == nil {
		return true
	}
	return false
}
func (v *stringValue) IsNotEmpty() bool {
	return !v.IsEmpty()
}
