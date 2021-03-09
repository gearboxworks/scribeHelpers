package gearConfig

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/ux"
)

type GearMeta struct {
	State        string `json:"state"`
	Organization string `json:"organization"`
	Name         string `json:"name"`
	Maintainer   string `json:"maintainer"`
	Class        string `json:"class"`
	Refurl       string `json:"refurl"`
}

func (gm *GearMeta) String() string {
	var ret string
	//if state := ux.IfNilReturnError(gm); state.IsError() {
	//	return ux.SprintfRed("GearConfig is nil!\n")
	//}

	for range onlyOnce {
		ret += ux.SprintfCyan("# GearMeta\n")
		ret += ux.SprintfBlue("\tOrganization: %s\n", gm.Organization)
		ret += ux.SprintfBlue("\tName:         %s\n", gm.Name)
		ret += ux.SprintfBlue("\tMaintainer:   %s\n", gm.Maintainer)
		ret += ux.SprintfBlue("\tClass:        %s\n", gm.Class)
		ret += ux.SprintfBlue("\tRefurl:       %s\n", gm.Refurl)
	}

	return ret
}


type GearBuild struct {
	Ports        GearPorts    `json:"ports"`
	Run          string       `json:"run"`		//
	Args         GearArgs     `json:"args"`		//
	Env          GearEnv      `json:"env"`
	Network      string       `json:"network"`
}

func (b *GearBuild) String() string {
	var ret string
	//if state := ux.IfNilReturnError(b); state.IsError() {
	//	return ux.SprintfRed("GearConfig is nil!\n")
	//}

	for range onlyOnce {
		ret += ux.SprintfCyan("# GearBuild\n")
		ret += ux.SprintfBlue("\tPorts:   %s\n", b.Ports.String())
		ret += ux.SprintfBlue("\tRun:     %s\n", b.Run)
		ret += ux.SprintfBlue("\tArgs:    %s\n", b.Args.String())
		ret += ux.SprintfBlue("\tNetwork: %s\n", b.Network)
		ret += ux.SprintfBlue("\tEnv:     %s\n", b.Env.String())
	}

	return ret
}


type GearRun struct {
	Ports        GearPorts    `json:"ports"`
	Env          GearEnv      `json:"env"`
	Volumes      string       `json:"volumes"`
	Network      string       `json:"network"`
	Commands     GearCommands `json:"commands"`
}

func (r *GearRun) String() string {
	var ret string
	//if state := ux.IfNilReturnError(r); state.IsError() {
	//	return ux.SprintfRed("GearConfig is nil!\n")
	//}

	for range onlyOnce {
		ret += ux.SprintfCyan("# GearRun\n")
		ret += ux.SprintfBlue("\tPorts:    %s\n", r.Ports.String())
		ret += ux.SprintfBlue("\tEnv:      %s\n", r.Env.String())
		ret += ux.SprintfBlue("\tVolumes:  %s\n", r.Volumes)
		ret += ux.SprintfBlue("\tNetwork:  %s\n", r.Network)
		ret += ux.SprintfBlue("\tCommands: %s\n", r.Commands.String())
	}

	return ret
}


type GearCommands map[string]string

func (c *GearCommands) String() string {
	var ret string
	//if state := ux.IfNilReturnError(c); state.IsError() {
	//	return ux.SprintfRed("GearConfig is nil!\n")
	//}

	//ret += ux.SprintfCyan("# GearCommands\n")
	ret += "\n"
	for k, v := range *c {
		ret += ux.SprintfBlue("\t\t%-16s => %s\n", k, v)
	}

	return ret
}


type GearProject struct {
}

func (p *GearProject) String() string {
	var ret string
	//if state := ux.IfNilReturnError(p); state.IsError() {
	//	return ux.SprintfRed("GearConfig is nil!\n")
	//}

	for range onlyOnce {
		ret += ux.SprintfCyan("# GearProject\n")
		//ret += ux.SprintfBlue("\tfoo: %v", p.Meta)
		//ret += ux.SprintfBlue("\tfoo: %v", p.Build)
		//ret += ux.SprintfBlue("\tfoo: %v", p.Run)
		//ret += ux.SprintfBlue("\tfoo: %v", p.Project)
		//ret += ux.SprintfBlue("\tfoo: %v", p.Extensions)
		//ret += ux.SprintfBlue("\tfoo: %v", p.Versions)
	}

	return ret
}


type GearExtensions struct {
}

func (ge *GearExtensions) String() string {
	var ret string
	//if state := ux.IfNilReturnError(ge); state.IsError() {
	//	return ux.SprintfRed("GearConfig is nil!\n")
	//}

	for range onlyOnce {
		ret += ux.SprintfCyan("# GearExtensions\n")
	}

	return ret
}


type GearEnv map[string]string

func (ge *GearEnv) String() string {
	var ret string
	//if state := ux.IfNilReturnError(ge); state.IsError() {
	//	return ux.SprintfRed("GearConfig is nil!\n")
	//}

	//ret += ux.SprintfCyan("# GearEnv\n")
	ret += "\n"
	for k, v := range *ge {
		ret += ux.SprintfBlue("\t\t%25s='%s'\n", k, v)
	}

	return ret
}


type GearArgs string

func (ga *GearArgs) String() string {
	var ret string
	//if state := ux.IfNilReturnError(ga); state.IsError() {
	//	return ux.SprintfRed("GearConfig is nil!\n")
	//}

	for range onlyOnce {
		//ret += ux.SprintfCyan("# GearArgs: ")
		ret += ux.SprintfBlue("%s", *ga)
	}

	return ret
}


type GearPorts map[string]string

func (ports *GearPorts) String() string {
	var ret string
	//if state := ux.IfNilReturnError(ports); state.IsError() {
	//	return ux.SprintfRed("GearConfig is nil!\n")
	//}

	//ret += ux.SprintfCyan("# GearPorts: ")
	for k, v := range *ports {
		ret += ux.SprintfBlue("%s:%s ", k, v)
	}

	return ret
}


type GearVersion struct {
	//Version 	 string `json:"version"`
	MajorVersion string `json:"majorversion"`
	Latest       bool   `json:"latest"`
	Ref          string `json:"ref"`
	Base         string `json:"base"`
}

func (v *GearVersion) String() string {
	var ret string
	//if state := ux.IfNilReturnError(v); state.IsError() {
	//	return ux.SprintfRed("GearConfig is nil!\n")
	//}

	for range onlyOnce {
		//ret += ux.SprintfCyan("# GearVersion\n")
		//ret += ux.SprintfBlue("Version: %-8s", v)
		ret += ux.SprintfBlue("MajorVersion: %-8s", v.MajorVersion)
		ret += ux.SprintfBlue("\tLatest: %v", v.Latest)
		ret += ux.SprintfBlue("\tRef: %-16s", v.Ref)
		ret += ux.SprintfBlue("\tBase: %s", v.Base)
	}

	return ret
}


type GearVersions map[string]GearVersion

func (vers *GearVersions) String() string {
	var ret string
	//if state := ux.IfNilReturnError(vers); state.IsError() {
	//	return ux.SprintfRed("GearConfig is nil!\n")
	//}

	ret += ux.SprintfCyan("# GearVersions\n")
	for k, v := range *vers {
		ret += ux.SprintfBlue("\t%-16s %s\n", k, v.String())
	}

	return ret
}


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
