package gearConfig

import "fmt"

type GearMeta struct {
	State        string `json:"state"`
	Organization string `json:"organization"`
	Name         string `json:"name"`
	Maintainer   string `json:"maintainer"`
	Class        string `json:"class"`
	Refurl       string `json:"refurl"`
}

type GearBuild struct {
	Ports        GearPorts    `json:"ports"`
	Run          string       `json:"run"`		//
	Args         GearArgs     `json:"args"`		//
	Env          GearEnv      `json:"env"`
	Network      string       `json:"network"`
}

type GearRun struct {
	Ports        GearPorts    `json:"ports"`
	Env          GearEnv      `json:"env"`
	Volumes      string       `json:"volumes"`
	Network      string       `json:"network"`
	Commands     GearCommands `json:"commands"`
}

type GearCommands map[string]string

type GearProject struct {
}

type GearExtensions struct {
}

type GearEnv map[string]string

type GearArgs string

type GearPorts map[string]string

type GearVersion struct {
	MajorVersion string `json:"majorversion"`
	Latest       bool   `json:"latest"`
	Ref          string `json:"ref"`
	Base         string `json:"base"`
}
type GearVersions map[string]GearVersion


func (ports *GearPorts) ToString() string {
	var p string

	for k, v := range *ports {
		p = fmt.Sprintf("%s %s:%s\n", p, k, v)
	}

	return p
}


func (vers *GearVersions) GetLatest() string {
	var v string

	for k, r := range *vers {
		if r.Latest {
			v = k
			break
		}
	}

	return v
}


func (vers *GearVersions) HasVersion(gearVersion string) bool {
	var ok bool

	for range onlyOnce {
		//if gearVersion == "latest" {
		//	gl := vers.GetLatest()
		//	if gl == "" {
		//		break
		//	}
		//}

		for v, r := range *vers {
			if r.Latest && (gearVersion == "latest") {
				ok = true
				break
			}

			if v == gearVersion {
				ok = true
				break
			}

			if r.MajorVersion == gearVersion {
				ok = true
				break
			}
		}
	}

	return ok
}
