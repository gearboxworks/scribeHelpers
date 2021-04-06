package toolGear

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/registry"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
	"github.com/newclarity/scribeHelpers/toolGear/gearConfig"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"sort"
	"strings"
)


type Gears struct {
	Language    Language
	Array		map[string]*Gear
	Selected    *Gear
	Build		*gearConfig.GearConfig

	Docker      *Docker

	Runtime     *toolRuntime.TypeRuntime
	State       *ux.State
}
type Language struct {
	AppName string
	ImageName string
	ContainerName string
}


func NewGears(runtime *toolRuntime.TypeRuntime) Gears {
	var gears Gears

	for range onlyOnce {
		runtime = runtime.EnsureNotNil()

		l := Language {
			AppName:       "Gearbox",
			ImageName:     "Gear Image",
			ContainerName: "Gear",
		}

		gears = Gears {
			Language:   l,
			Array:      make(map[string]*Gear),
			Selected:   nil,

			Docker:     NewDocker(runtime),

			Runtime:    runtime,
			State:      ux.NewState(runtime.CmdName, runtime.Debug),
		}
		gears.State.SetPackage("")
		gears.State.SetFunctionCaller()

		//_ = gears.Get()
		//gears.State = gears.Get()
		//if gears.State.IsNotOk() {
		//	break
		//}
	}

	return gears
}

func NewGearConfig(runtime *toolRuntime.TypeRuntime) *gearConfig.GearConfig {
	return gearConfig.New(runtime)
}


func (gears *Gears) IsValid() *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.State.EnsureNotNil()

		if gears.Docker.Client == nil {
			gears.State.SetError("docker client is nil")
			break
		}
	}

	return gears.State
}

func (gears *Gears) IsNil() *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.State.EnsureNotNil()
	}

	return gears.State
}


// ******************************************************************************** //

func (gears *Gears) SetLanguage(appName string, imageName string, containerName string) *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.Language = Language {
			AppName:       appName,
			ImageName:     imageName,
			ContainerName: containerName,
		}
	}

	return gears.State
}

func (gears *Gears) SetProvider(provider string) *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.Docker.Provider.SetProvider(provider)
		if gears.State.IsNotOk() {
			break
		}
	}

	return gears.State
}

func (gears *Gears) SetProviderUrl(Url string) *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.Docker.Provider.SetUrl(Url)
		if gears.State.IsNotOk() {
			break
		}

		gears.State = gears.Docker.Connect()
	}

	return gears.State
}

func (gears *Gears) SetProviderHost(host string, port string) *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.Docker.Provider.SetHost(host, port)
		if gears.State.IsNotOk() {
			break
		}

		gears.State = gears.Docker.Connect()
	}

	return gears.State
}

func (gears *Gears) Get() *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.GetImages("")
		if gears.State.IsNotOk() {
			break
		}

		gears.State = gears.GetContainers("")
		if gears.State.IsNotOk() {
			break
		}
	}

	return gears.State
}

func (gears *Gears) Refresh() *ux.State {
	state := ux.EnsureStateNotNil(nil)

	for _, v := range gears.Array {
		state = v.Refresh()
		if state.IsNotOk() {
			break
		}
	}

	return state
}

func (gears *Gears) ListContainers(name string) *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.GetContainers("")
		if gears.State.IsNotOk() {
			break
		}

		ux.PrintfCyan("%s %ss: ", gears.Language.AppName, gears.Language.ContainerName)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Name",
			"Class",
			"State",
			"Image",
			"Ports",
			"IP Address",
			"Mounts",
			"Size",
		})

		for _, gear := range gears.Array {
			if gear.Container == nil {
				continue
			}
			//if gear.Container.Summary == nil {
			//	continue
			//}
			if gear.Container.ID == "" {
				continue
			}
			name := gear.Container.Name

			sshPort := ""
			var ports string
			for _, p := range gear.Container.GetPorts() {
				if p.PrivatePort == 22 {
					sshPort = fmt.Sprintf("%d", p.PublicPort)
					continue
				}
				//ports += fmt.Sprintf("%s://%s:%d => %d\n", p.Type, p.IP, p.PublicPort, p.PrivatePort)
				if p.IP == "0.0.0.0" {
					ports += fmt.Sprintf("%d => %d\n", p.PublicPort, p.PrivatePort)
				} else {
					ports += fmt.Sprintf("%s://%s:%d => %d\n", p.Type, p.IP, p.PublicPort, p.PrivatePort)
				}
			}
			if sshPort == "0" {
				sshPort = "none"
			}

			var mounts string
			for _, m := range gear.Container.GetMounts() {
				// ms += fmt.Sprintf("%s(%s) host:%s => container:%s (RW:%v)\n", m.Name, m.Type, m.Source, m.Destination, m.RW)
				mounts += fmt.Sprintf("host:%s\n\t=> %s:%s (RW:%v)\n",
					m.Source,
					gears.Language.ContainerName,
					m.Destination,
					m.RW,
				)
			}

			var ipAddress string
			for k, n := range gear.Container.GetNetworks() {
				ipAddress += fmt.Sprintf("(%s) %s\n", k, n.IPAddress)
			}

			state := gear.Container.GetState()
			if state == ux.StateRunning {
				state = ux.SprintfGreen(state)
			} else {
				state = ux.SprintfYellow(state)
			}

			t.AppendRow([]interface{}{
				ux.SprintfWhite(name),
				ux.SprintfWhite(gear.GearConfig.GetClass()),
				state,
				ux.SprintfWhite(gear.Image.GetName()),
				ux.SprintfWhite(ports),
				ux.SprintfWhite(ipAddress),
				ux.SprintfWhite(mounts),
				ux.SprintfWhite(humanize.Bytes(gear.Container.GetSize())),
			})
		}

		gears.State.ClearError()
		count := t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		ux.PrintflnGreen("%d found", count)
		t.Render()
		gears.State.SetResponse(count)

		ux.PrintflnBlue("")
	}

	return gears.State
}

func (gears *Gears) Ls(name string) *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.GetContainers(name)
		if gears.State.IsError() {
			break
		}

		gears.State = gears.NetworkList(DefaultNetwork)
	}

	return gears.State
}

func (gears *Gears) AddGears(gc *gearConfig.GearConfig) *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	for range onlyOnce {
		for k := range gc.Versions {
			var ok bool
			gears.State = gears.FindImage(gc.Meta.Name, k)
			//if gears.State.IsNotOk() {
			//	break
			//}
			ok = gears.State.GetResponseAsBool()
			if ok {
				gears.Selected.BuildFlag = true
				continue
			}

			gear := NewGear(gears.Runtime, gears.Docker)
			gear.GearConfig = gc
			gear.BuildFlag = true
			gears.Array[gc.Meta.Name + "-" + k] = gear
		}
	}

	return gears.State
}


// ******************************************************************************** //

func (gears *Gears) SelectedPull(version string) *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	return gears.Selected.Pull(version)
}

func (gears *Gears) SelectedSsh(interactive bool, statusLine bool, mountPath string, cmdArgs []string) *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	return gears.Selected.ContainerSsh(interactive, statusLine, mountPath, cmdArgs)
}

func (gears *Gears) SelectedStart() *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	return gears.Selected.Start()
}

func (gears *Gears) SelectedStop() *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	return gears.Selected.Stop()
}

func (gears *Gears) SelectedLogs() *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	return gears.Selected.Logs()
}

func (gears *Gears) SelectedRemove() *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	return gears.Selected.Remove()
}

func (gears *Gears) SelectedImageRemove() *ux.State {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return state
	}

	return gears.Selected.ImageRemove()
}

func (gears *Gears) SelectedAddVolume(local string, remote string) bool {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return false
	}

	return gears.Selected.AddVolume(local, remote)
}

func (gears *Gears) SelectedVersions() *gearConfig.GearVersions {
	if state := ux.IfNilReturnError(gears); state.IsError() {
		return &gearConfig.GearVersions{}
	}

	return gears.Selected.GetVersions()
}


// ******************************************************************************** //

func (gears *Gears) ListImages(f string) *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.GetImages("")
		if gears.State.IsNotOk() {
			break
		}

		ux.PrintfCyan("Downloaded Gearbox images: ")
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Class", "Image", "Ports", "Size"})

		//sort.Sort(ArraySorter(gears.Array))
		for _, gear := range gears.Array {
			// foo := fmt.Sprintf("%s/%s", gc.Organization, gc.Name)
			t.AppendRow([]interface{}{
				ux.SprintfWhite(gear.GearConfig.GetClass()),
				//ux.SprintfWhite(gc.Meta.State),
				ux.SprintfWhite(gear.Image.GetName()),
				ux.SprintfWhite("%s", gear.GearConfig.GetPorts()),
				ux.SprintfWhite(humanize.Bytes(gear.Image.GetSize())),
			})
		}

		gears.State.ClearError()
		count := t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		ux.PrintflnGreen("%d found", count)
		t.Render()
		gears.State.SetResponse(count)

		ux.PrintflnBlue("")
	}

	return gears.State
}

//type ArraySorter map[string]*Gear
//func (a ArraySorter) Len() int           { return len(a) }
//func (a ArraySorter) Swap(i, j string)      { a[i], a[j] = a[j], a[i] }
//func (a ArraySorter) Less(i, j string) bool { return a[i].GearConfig.Meta.Name < a[j].GearConfig.Meta.Name }

func (gears *Gears) GetImages(name string) *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.Docker.ImageList(nil)
		if gears.State.IsNotOk() {
			break
		}

		for _, c := range gears.Docker.Images {
			gear := NewGear(gears.Runtime, gears.Docker)
			//gear.Docker = gears.Docker

			if _, ok := c.Labels["gearbox.json"]; !ok {
				continue
			}

			gears.State = gear.GearConfig.ParseJson(c.Labels["gearbox.json"])
			if gears.State.IsError() {
				continue
			}

			if gear.GearConfig.Meta.Organization != DefaultOrganization {
				continue
			}

			if len(c.RepoTags) == 0 {
				continue
			}

			if c.RepoTags[0] == "<none>:<none>" {
				continue
			}

			if name != "" {
				if gear.GearConfig.Meta.Name != name {
					continue
				}
			}

			gear.Image.ID = c.ID
			gear.Image.Name = gear.GearConfig.GetName()
			gear.Image.Version = c.Labels["gearbox.version"]
			gear.Image.Summary = c
			gear.Image.Docker = gear.Docker
			gear.Container.Docker = gear.Docker
			gear.Image.GearConfig = gear.GearConfig
			gear.Image.Details, _ = gear.Docker.ImageInspectWithRaw(c.ID)

			gears.Array[c.ID] = gear
		}
	}

	return gears.State
}

func (gears *Gears) FindImage(gearName string, gearVersion string) *ux.State {
	var ok bool
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if gearName == "" {
			gears.State.SetError("empty gear name")
			break
		}

		if gearVersion == "" {
			gearVersion = gearConfig.LatestName
		}

		for _, i := range gears.Array {
			ok, _ = MatchImage(&i.Image.Summary,
				TypeMatchImage{Organization: DefaultOrganization, Name: gearName, Version: gearVersion})
			if ok {
				gears.Selected = i
				break
			}
		}

		if !ok {
			gears.State.SetWarning("Container image '%s:%s' doesn't exist.", gearName, gearVersion)
		}

		gears.State.SetOk("found image")
	}

	gears.State.SetResponse(ok)
	return gears.State
}

func (gears *Gears) CreateImage(gearName string, gearVersion string) *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if gearName == "" {
			gears.State.SetError("empty gear name")
			break
		}

		if gearVersion == "" {
			gearVersion = gearConfig.LatestName
		}

		var found bool
		gears.State = gears.FindImage(gearName, gearVersion)
		if gears.State.IsError() {
			break
		}
		found = gears.State.GetResponseAsBool()
		if found {
			//gears.State = gears.Selected.Remove()
			//if gears.State.IsError() {
			//	break
			//}
			gears.State.SetWarning("Already exists.")
			break
		}

		//gears.State = gears.Docker.Pull("gearboxworks", gearName, gearVersion)
		//if gears.State.IsNotOk() {
		//	break
		//}

		vers := gears.Selected.GetVersion(gearVersion)
		//var DockerArgs string
		if vers.IsBaseRef() {
			//DockerArgs = "--squash"
		} else {
			// docker pull "${GB_REF}"
			run := gears.Selected.GetBuildRun()
			if run == "" {
				//GEARBOX_ENTRYPOINT="$(docker inspect --format '{{ with }}{{ else }}{{ with .ContainerConfig.Entrypoint}}{{ index . 0 }}{{ end }}' "${GB_REF}")"
				//export GEARBOX_ENTRYPOINT
				//GEARBOX_ENTRYPOINT_ARGS="$(docker inspect --format '{{ join .ContainerConfig.Entrypoint " " }}' "${GB_REF}")"
				//export GEARBOX_ENTRYPOINT_ARGS
			} else {
				//GEARBOX_ENTRYPOINT="${GB_RUN}"
				//export GEARBOX_ENTRYPOINT
				//GEARBOX_ENTRYPOINT_ARGS="${GB_ARGS}"
				//export GEARBOX_ENTRYPOINT_ARGS

				//
			}

			// docker build -t ${GB_IMAGENAME}:${GB_VERSION} -f ${GB_DOCKERFILE} --build-arg GEARBOX_ENTRYPOINT
			// --build-arg GEARBOX_ENTRYPOINT_ARGS ${DOCKER_ARGS} .

			//if [ "${GB_MAJORVERSION}" != "" ]
			//then
			//	docker tag ${GB_IMAGENAME}:${GB_VERSION} ${GB_IMAGENAME}:${GB_MAJORVERSION}
			//fi
			//
			//if [ "${GB_LATEST}" == "true" ]
			//then
			//	docker tag ${GB_IMAGENAME}:${GB_VERSION} ${GB_IMAGENAME}:latest
			//fi
		}

		//if gears.Selected.GearConfig.Versions.HasVersion(gearVersion) == "" {
		//	//
		//}
	}

	return gears.State
}

// Search for an image in remote registry.
func (gears *Gears) Search(gearName string, gearVersion string) *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var repo string
		if gearVersion == "" {
			repo = fmt.Sprintf("gearboxworks/%s", gearName)
		} else {
			repo = fmt.Sprintf("gearboxworks/%s:%s", gearName, gearVersion)
		}
		repo = gearName

		var resp []registry.SearchResult
		resp, gears.State = gears.Docker.ImageSearch(repo, nil)

		sort.Sort(NameSorter(resp))
		for _, v := range resp {
			if !strings.HasPrefix(v.Name, "gearboxworks/") {
				continue
			}
			ux.PrintfCyan("%s", v.Name)
			ux.PrintfWhite("\t- ")
			ux.PrintfBlue("%s\n", v.Description)
		}
	}

	return gears.State
}

type NameSorter []registry.SearchResult
func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }


func MatchImage(m *types.ImageSummary, match TypeMatchImage) (bool, *gearConfig.GearConfig) {
	var ok bool
	gc := gearConfig.New(nil)

	for range onlyOnce {
		if MatchTag("<none>:<none>", m.RepoTags) {
			ok = false
			break
		}

		gc.State = gc.ParseJson(m.Labels["gearbox.json"])
		if gc.State.IsError() {
			ok = false
			break
		}

		if gc.Meta.Organization != DefaultOrganization {
			ok = false
			break
		}

		//tagCheck := fmt.Sprintf("%s/%s:%s", match.Organization, match.Name, match.Version)
		//if !MatchTag(tagCheck, m.RepoTags) {
		//	ok = false
		//	break
		//}
		//
		//if _, ok2 := m.Labels["container.latest"]; ok2 {
		//	tagCheck := fmt.Sprintf("%s/%s:latest", match.Organization, match.Name)
		//	if !MatchTag(tagCheck, m.RepoTags) {
		//		ok = false
		//		break
		//	}
		//}

		if gc.Meta.Name != match.Name {
			if !RunAs.AsLink {
				ok = false
				break
			}

			cs := gc.MatchCommand(match.Name)
			if cs == nil {
				ok = false
				break
			}

			match.Name = gc.Meta.Name
		}

		if !gc.Versions.HasVersion(match.Version) {
			ok = false
			break
		}

		if match.Version == gearConfig.LatestName {
			gl := gc.Versions.GetLatest()
			if match.Version != "" {
				match.Version = gl
			}
		}

		for range onlyOnce {
			if m.Labels["gearbox.version"] == match.Version {
				ok = true
				break
			}

			if m.Labels["container.majorversion"] == match.Version {
				ok = true
				break
			}

			ok = false
		}
		break
	}

	return ok, gc
}

func MatchTag(match string, tags []string) bool {
	var ok bool

	for _, s := range tags {
		if s == match {
			ok = true
			break
		}
	}

	return ok
}

//func (gear *Gear) updateImage(image *types.ImageSummary, id string) (*ux.State) {
//	if state := gear.IsNil(); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		if _, ok := image.Labels["gearbox.json"]; !ok {
//			continue
//		}
//
//		gear.State = gear.GearConfig.ParseJson(image.Labels["gearbox.json"])
//		if gear.State.IsError() {
//			continue
//		}
//
//		if gear.GearConfig.Meta.Organization != DefaultOrganization {
//			continue
//		}
//
//		if len(image.RepoTags) == 0 {
//			continue
//		}
//
//		if image.RepoTags[0] == "<none>:<none>" {
//			continue
//		}
//
//		gear.Image.ID = image.ID
//		gear.Image.Name = gear.GearConfig.GetName()
//		gear.Image.Version = image.Labels["gearbox.version"]
//		gear.Image.Summary = *image
//		gear.Image.Docker = gear.Docker
//		gear.Image.GearConfig = gear.GearConfig
//		gear.Image.Details, _ = gear.Docker.ImageInspectWithRaw(image.ID)
//
//		gear.Container.Docker = gear.Docker
//		gear.Container.GearConfig = gear.GearConfig
//	}
//
//	return gear.State
//}


// ******************************************************************************** //

func (gears *Gears) ContainerCreate(gearName string, gearVersion string) *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//if c.runtime.Debug {
		//	fmt.Printf("DEBUG: ContainerCreate(%s, %s)\n", gearName, gearVersion)
		//}

		if gearName == "" {
			gears.State.SetError("empty gearname")
			break
		}

		if gearVersion == "" {
			gearVersion = gearConfig.LatestName
		}

		var ok bool
		ok, gears.State = gears.FindContainer(gearName, gearVersion)
		if gears.State.IsError() {
			break
		}
		if !ok {
			// Find Gear image since we don't have a container.
			for range onlyOnce {
				gears.State = gears.FindImage(gearName, gearVersion)
				if gears.State.IsError() {
					ok = false
					break
				}
				ok = gears.State.GetResponseAsBool()
				if ok {
					break
				}

				ux.PrintflnNormal("Downloading Gear '%s:%s'.", gearName, gearVersion)
				newGear := NewGear(gears.Runtime, gears.Docker)

				// Pull Gear image.
				newGear.Image.ID = gearName
				newGear.Image.Name = gearName
				newGear.Image.Version = gearVersion
				gears.State = newGear.Image.Pull()
				if gears.State.IsError() {
					gears.State.SetError("no such gear '%s'", gearName)
					break
				}

				ok, newGear.GearConfig = MatchImage(&newGear.Image.Summary,
					TypeMatchImage{Organization: DefaultOrganization, Name: gearName, Version: gearVersion})
				if ok {
					break
				}
				newGear.Image.GearConfig = newGear.GearConfig
				newGear.Container.GearConfig = newGear.GearConfig

				gears.State = newGear.Image.Refresh()
				if gears.State.IsNotOk() {
					break
				}

				gears.Array[newGear.Image.ID] = newGear
				gears.Selected = gears.Array[newGear.Image.ID]

				ok = true

				//gears.State = gears.Docker.ImageList(&types.ImageListOptions{All: true})
				//if gears.State.IsNotOk() {
				//	break
				//}
				//
				//// Confirm it's there.
				//ok, gears.State = gears.FindImage(gearName, gearVersion)
				//if gears.State.IsError() {
				//	ok = false
				//	break
				//}

			}
			if !ok {
				gears.State.SetError("Cannot install Gear image '%s:%s' - %s.", gearName, gearVersion, gears.State.GetError())
				break
			}

			//c.State.Clear()
		}

		//c.Selected.Container.ID = c.Selected.Image.ID
		//c.Selected.Container.Name = c.Selected.Image.Name
		//c.Selected.Container.Version = c.Selected.Image.Version

		// c.Image.Details.Container = "gearboxworks/golang:1.14"
		// tag := fmt.Sprintf("", c.Image.Name, c.Image.Version)
		tag := fmt.Sprintf("gearboxworks/%s:%s", gearName, gearVersion)
		gn := fmt.Sprintf("%s-%s", gearName, gearVersion)

		binds := gears.Selected.GetVolumeMounts()

		config := container.Config {
			// Hostname:        "",
			// Domainname:      "",
			User:            "root",
			// AttachStdin:     false,
			AttachStdout:    true,
			AttachStderr:    true,
			ExposedPorts:    nil,
			Tty:             false,
			OpenStdin:       false,
			StdinOnce:       false,
			Env:             nil,
			Cmd:             []string{"/init"},
			// Healthcheck:     nil,
			// ArgsEscaped:     false,
			Image:           tag,
			// Volumes:         nil,
			// WorkingDir:      "",
			// Entrypoint:      nil,
			// NetworkDisabled: false,
			// MacAddress:      "",
			// OnBuild:         nil,
			// Labels:          nil,
			// StopSignal:      "",
			// StopTimeout:     nil,
			// Shell:           nil,
		}

		netConfig := network.NetworkingConfig {}

		// DockerMount
		// ms := mount.Mount {
		// 	Type:          "bind",
		// 	Source:        "/Users/mick/Documents/GitHub/containers/docker-golang",
		// 	Target:        "/foo",
		// 	ReadOnly:      false,
		// 	Consistency:   "",
		// 	BindOptions:   nil,
		// 	VolumeOptions: nil,
		// 	TmpfsOptions:  nil,
		// }

		//ports := gears.Selected.GetFixedPorts()
		//var ports nat.PortMap
		//if len(gears.Selected.GearConfig.Build.FixedPorts) > 0 {
		//	ports = make(nat.PortMap)
		//	for k, v := range gears.Selected.GearConfig.Build.FixedPorts {
		//		fmt.Printf("%s => %v\n", k, v)
		//		var bind []nat.PortBinding
		//		bind = append(bind, nat.PortBinding {
		//			HostIP: "0.0.0.0",
		//			HostPort: v,
		//		})
		//		ports[(nat.Port)(v + "/tcp")] = bind
		//	}
		//} else {
		//	ports = nil
		//}

		hostConfig := container.HostConfig {
			Binds:           binds,
			ContainerIDFile: "",
			LogConfig:       container.LogConfig{
				Type:   "",
				Config: nil,
			},
			NetworkMode:     DefaultNetwork,
			PortBindings:    gears.Selected.GetFixedPortBindings(),
			RestartPolicy:   container.RestartPolicy {
				Name:              "",
				MaximumRetryCount: 0,
			},
			AutoRemove:      false,
			VolumeDriver:    "",
			VolumesFrom:     nil,
			CapAdd:          nil,
			CapDrop:         nil,
			//Capabilities:    nil,
			//CgroupnsMode:    "",
			DNS:             []string{},
			DNSOptions:      []string{},
			DNSSearch:       []string{},
			ExtraHosts:      nil,
			GroupAdd:        nil,
			IpcMode:         "",
			Cgroup:          "",
			Links:           nil,
			OomScoreAdj:     0,
			PidMode:         "",
			Privileged:      true,
			PublishAllPorts: true,
			ReadonlyRootfs:  false,
			SecurityOpt:     nil,
			StorageOpt:      nil,
			Tmpfs:           nil,
			UTSMode:         "",
			UsernsMode:      "",
			ShmSize:         0,
			Sysctls:         nil,
			Runtime:         "runc",
			ConsoleSize:     [2]uint{},
			Isolation:       "",
			Resources:       container.Resources{},
			Mounts:          []mount.Mount{},
			//MaskedPaths:     nil,
			//ReadonlyPaths:   nil,
			Init:            nil,
		}

		var resp container.ContainerCreateCreatedBody
		resp, gears.State = gears.Docker.ContainerCreate(&config, &hostConfig, &netConfig, gn)
		if gears.State.IsNotOk() {
			break
		}

		gears.Selected.Container.ID = resp.ID
		gears.Selected.Container.Name = gearName
		gears.Selected.Container.Version = gearVersion
		gears.Selected.Container.Docker = gears.Docker
		if gearVersion == gearConfig.LatestName {
			gears.Selected.Container.IsLatest = true
		}

		// var response Response
		gears.State = gears.Selected.Refresh()
		if gears.State.IsError() {
			break
		}

		if gears.State.IsCreated() {
			break
		}

		//if c.State.IsRunning() {
		//	break
		//}
		//
		//if c.State.IsPaused() {
		//	break
		//}
		//
		//if c.State.IsRestarting() {
		//	break
		//}
	}

	return gears.State
}

func (gears *Gears) GetContainers(name string) *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.Docker.ContainerList(nil,true)
		if gears.State.IsNotOk() {
			break
		}

		var c types.Container
		for _, c = range gears.Docker.Containers {
			if _, ok := c.Labels["gearbox.json"]; !ok {
				continue
			}

			gear := NewGear(gears.Runtime, gears.Docker)
			//gear.Docker = gears.Docker
			gears.State = gear.GearConfig.ParseJson(c.Labels["gearbox.json"])
			if gears.State.IsError() {
				continue
			}

			if gear.GearConfig.Meta.Organization != DefaultOrganization {
				continue
			}

			if name != "" {
				if gear.GearConfig.Meta.Name != name {
					continue
				}
			}

			gear.Container.ID = c.ID
			gear.Container.Name = gear.GearConfig.GetName()
			gear.Container.Version = c.Labels["gearbox.version"]
			gear.Container.Summary = c
			gear.Container.Docker = gear.Docker
			gear.Container.GearConfig = gear.GearConfig
			gear.Container.Details, _ = gear.Docker.ContainerInspect(c.ID)
			//gear.State.RunState = c.State
			gear.State.RunState = c.State
			if c.Labels["container.latest"] == "true" {
				gear.Container.IsLatest = true
			}

			//gear.Image.ID = gear.Container.Summary.ImageID
			if _, ok := gears.Array[c.ImageID]; ok {
				gears.Array[c.ImageID].Container = gear.Container
			} else {
				gears.Array[c.ID] = gear
			}
		}

	}

	return gears.State
}

func (gears *Gears) ContainerListFiles(name string) (int, *ux.State) {
	var count int
	if state := gears.IsNil(); state.IsError() {
		return 0, state
	}

	for range onlyOnce {
		gears.State = gears.Docker.ContainerList(nil,true)

		ux.PrintfCyan("Installed %s %s: ", gears.Language.AppName, gears.Language.ContainerName)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Name",
			"Class",
			"State",
			"Image",
			"Ports",
			"SSH port",
			"IP Address",
			"Mounts",
			"Size",
		})


		//gc := toolGear.NewGearConfig(gear.Runtime)
		gc := gearConfig.New(gears.Runtime)
		for _, c := range gears.Docker.Containers {
			gears.State = gc.ParseJson(c.Labels["gearbox.json"])
			if gears.State.IsError() {
				continue
			}

			if gc.Meta.Organization != DefaultOrganization {
				continue
			}

			if name != "" {
				if gc.Meta.Name != name {
					continue
				}
			}

			name := strings.TrimPrefix(c.Names[0], "/")

			sshPort := ""
			var ports string
			for _, p := range c.Ports {
				if p.PrivatePort == 22 {
					sshPort = fmt.Sprintf("%d", p.PublicPort)
					continue
				}
				//ports += fmt.Sprintf("%s://%s:%d => %d\n", p.Type, p.IP, p.PublicPort, p.PrivatePort)
				if p.IP == "0.0.0.0" {
					ports += fmt.Sprintf("%d => %d\n", p.PublicPort, p.PrivatePort)
				} else {
					ports += fmt.Sprintf("%s://%s:%d => %d\n", p.Type, p.IP, p.PublicPort, p.PrivatePort)
				}
			}
			if sshPort == "0" {
				sshPort = "none"
			}

			var mounts string
			for _, m := range c.Mounts {
				// ms += fmt.Sprintf("%s(%s) host:%s => container:%s (RW:%v)\n", m.Name, m.Type, m.Source, m.Destination, m.RW)
				mounts += fmt.Sprintf("host:%s\n\t=> %s:%s (RW:%v)\n",
					m.Source,
					gears.Language.ContainerName,
					m.Destination,
					m.RW,
				)
			}

			var ipAddress string
			for k, n := range c.NetworkSettings.Networks {
				ipAddress += fmt.Sprintf("(%s) %s\n", k, n.IPAddress)
			}

			var state string
			if c.State == ux.StateRunning {
				state = ux.SprintfGreen(c.State)
			} else {
				state = ux.SprintfYellow(c.State)
			}

			t.AppendRow([]interface{}{
				ux.SprintfWhite(name),
				ux.SprintfWhite(gc.Meta.Class),
				state,
				ux.SprintfWhite(c.Image),
				ux.SprintfWhite(ports),
				ux.SprintfWhite(sshPort),
				ux.SprintfWhite(ipAddress),
				ux.SprintfWhite(mounts),
				ux.SprintfWhite(humanize.Bytes(uint64(c.SizeRootFs))),
			})
		}

		gears.State.ClearError()
		count = t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		ux.PrintflnGreen("%d found", count)
		t.Render()
		ux.PrintflnBlue("")
	}

	return count, gears.State
}

func (gears *Gears) PrintContainers(name string) (int, *ux.State) {
	var count int
	if state := gears.IsNil(); state.IsError() {
		return 0, state
	}

	for range onlyOnce {
		ux.PrintfCyan("Installed %s %s: ", gears.Language.AppName, gears.Language.ContainerName)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Name",
			"Class",
			"State",
			"Image",
			"Ports",
			"SSH port",
			"IP Address",
			"Mounts",
			"Size",
		})

		gears.State = gears.Docker.ContainerList(nil,true)
		//gc := toolGear.NewGearConfig(gear.Runtime)
		gc := gearConfig.New(gears.Runtime)
		for _, c := range gears.Array {
			gears.State = gc.ParseJson(c.Container.Summary.Labels["gearbox.json"])
			if gears.State.IsError() {
				continue
			}

			if gc.Meta.Organization != DefaultOrganization {
				continue
			}

			if name != "" {
				if gc.Meta.Name != name {
					continue
				}
			}

			name := strings.TrimPrefix(c.Container.Summary.Names[0], "/")

			sshPort := ""
			var ports string
			for _, p := range c.Container.Summary.Ports {
				if p.PrivatePort == 22 {
					sshPort = fmt.Sprintf("%d", p.PublicPort)
					continue
				}
				//ports += fmt.Sprintf("%s://%s:%d => %d\n", p.Type, p.IP, p.PublicPort, p.PrivatePort)
				if p.IP == "0.0.0.0" {
					ports += fmt.Sprintf("%d => %d\n", p.PublicPort, p.PrivatePort)
				} else {
					ports += fmt.Sprintf("%s://%s:%d => %d\n", p.Type, p.IP, p.PublicPort, p.PrivatePort)
				}
			}
			if sshPort == "0" {
				sshPort = "none"
			}

			var mounts string
			for _, m := range c.Container.Summary.Mounts {
				// ms += fmt.Sprintf("%s(%s) host:%s => container:%s (RW:%v)\n", m.Name, m.Type, m.Source, m.Destination, m.RW)
				mounts += fmt.Sprintf("host:%s\n\t=> %s:%s (RW:%v)\n",
					m.Source,
					gears.Language.ContainerName,
					m.Destination,
					m.RW,
				)
			}

			var ipAddress string
			for k, n := range c.Container.Summary.NetworkSettings.Networks {
				ipAddress += fmt.Sprintf("(%s) %s\n", k, n.IPAddress)
			}

			var state string
			if c.Container.Details.State.Status == ux.StateRunning {
				state = ux.SprintfGreen(c.Container.Details.State.Status)
			} else {
				state = ux.SprintfYellow(c.Container.Details.State.Status)
			}

			t.AppendRow([]interface{}{
				ux.SprintfWhite(name),
				ux.SprintfWhite(gc.Meta.Class),
				state,
				ux.SprintfWhite(c.Image.Name),
				ux.SprintfWhite(ports),
				ux.SprintfWhite(sshPort),
				ux.SprintfWhite(ipAddress),
				ux.SprintfWhite(mounts),
				ux.SprintfWhite(humanize.Bytes(uint64(*c.Container.Details.SizeRootFs))),
			})
		}

		gears.State.ClearError()
		count = t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		ux.PrintflnGreen("%d found", count)
		t.Render()
		ux.PrintflnBlue("")
	}

	return count, gears.State
}

func (gears *Gears) ContainerSprintf(name string) string {
	var ret string
	if state := gears.IsNil(); state.IsError() {
		ret = ux.SprintfRed("No %s %s found.\n", gears.Language.AppName, gears.Language.ContainerName)
		return ret
	}

	for range onlyOnce {
		gears.State = gears.Docker.ContainerList(nil,true)
		if gears.State.IsNotOk() {
			break
		}

		ret = ux.SprintfCyan("Installed %s %s:\n", gears.Language.AppName, gears.Language.ContainerName)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Name",
			"Class",
			"State",
			"Image",
			"Ports",
			"SSH port",
			"IP Address",
			"Mounts",
			"Size",
		})

		//gc := toolGear.NewGearConfig(gear.Runtime)
		gc := gearConfig.New(gears.Runtime)
		for _, c := range gears.Docker.Containers {
			//c.State = gc.ParseJson(c.Summary.Labels["gearbox.json"])
			//if c.State.IsError() {
			//	break
			//}
			gears.State = gc.ParseJson(c.Labels["gearbox.json"])
			if gears.State.IsError() {
				continue
			}

			if gc.Meta.Organization != DefaultOrganization {
				continue
			}

			if name != "" {
				if gc.Meta.Name != name {
					continue
				}
			}

			name := strings.TrimPrefix(c.Names[0], "/")

			sshPort := ""
			var ports string
			for _, p := range c.Ports {
				if p.PrivatePort == 22 {
					sshPort = fmt.Sprintf("%d", p.PublicPort)
					continue
				}
				//ports += fmt.Sprintf("%s://%s:%d => %d\n", p.Type, p.IP, p.PublicPort, p.PrivatePort)
				if p.IP == "0.0.0.0" {
					ports += fmt.Sprintf("%d => %d\n", p.PublicPort, p.PrivatePort)
				} else {
					ports += fmt.Sprintf("%s://%s:%d => %d\n", p.Type, p.IP, p.PublicPort, p.PrivatePort)
				}
			}
			if sshPort == "0" {
				sshPort = "none"
			}

			var mounts string
			for _, m := range c.Mounts {
				// ms += fmt.Sprintf("%s(%s) host:%s => container:%s (RW:%v)\n", m.Name, m.Type, m.Source, m.Destination, m.RW)
				mounts += fmt.Sprintf("host:%s\n\t=> %s:%s (RW:%v)\n",
					m.Source,
					gears.Language.ContainerName,
					m.Destination,
					m.RW,
				)
			}

			var ipAddress string
			for k, n := range c.NetworkSettings.Networks {
				ipAddress += fmt.Sprintf("(%s) %s\n", k, n.IPAddress)
			}

			var state string
			if c.State == ux.StateRunning {
				state = ux.SprintfGreen(c.State)
			} else {
				state = ux.SprintfYellow(c.State)
			}

			t.AppendRow([]interface{}{
				ux.SprintfWhite(name),
				ux.SprintfWhite(gc.Meta.Class),
				state,
				ux.SprintfWhite(c.Image),
				ux.SprintfWhite(ports),
				ux.SprintfWhite(sshPort),
				ux.SprintfWhite(ipAddress),
				ux.SprintfWhite(mounts),
				ux.SprintfWhite(humanize.Bytes(uint64(c.SizeRootFs))),
			})
		}

		gears.State.ClearError()
		count := t.Length()
		if count == 0 {
			ret += ux.SprintfYellow("No %s %s found.\n", gears.Language.AppName, gears.Language.ContainerName)
			break
		}
		ret += ux.SprintfGreen("Found %d %s %ss.\n", count, gears.Language.AppName, gears.Language.ContainerName)

		ret += t.Render()
		//ux.PrintflnBlue("")
	}

	return ret
}

func (gears *Gears) FindContainer(gearName string, gearVersion string) (bool, *ux.State) {
	var ok bool
	if state := gears.IsNil(); state.IsError() {
		return false, state
	}

	for range onlyOnce {
		if gearName == "" {
			gears.State.SetError("empty gearname")
			break
		}

		if gearVersion == "" {
			gearVersion = gearConfig.LatestName
		}

		for _, c := range gears.Array {
			ok, _ = MatchContainer(&c.Container.Summary,
				TypeMatchContainer{Organization: DefaultOrganization, Name: gearName, Version: gearVersion})
			if ok {
				gears.Selected = c
				gears.State.RunState = c.Container.Summary.State
				break
			}
		}

		if !ok {
			gears.State.SetWarning("Container '%s-%s' doesn't exist.", gearName, gearVersion)
			break
		}

		gears.State.SetOk("found %s", gears.Language.ContainerName)
	}

	return ok, gears.State
}

func (gears *Gears) FindContainers(gearName string) (map[string]*Gear, *ux.State) {
	ret := make(map[string]*Gear)
	if state := gears.IsNil(); state.IsError() {
		return ret, state
	}

	for range onlyOnce {
		if gearName == "" {
			gears.State.SetError("empty gearname")
			break
		}

		gearVersion := "all"

		var ok bool
		for k, v := range gears.Array {
			ok, _ = MatchContainer(&v.Container.Summary,
				TypeMatchContainer{Organization: DefaultOrganization, Name: gearName, Version: gearVersion})
			if ok {
				ret[k] = v
				//break
			}
		}

		if !ok {
			gears.State.SetWarning("Containers '%s' don't exist.", gearName)
			break
		}

		gears.State.SetOk("found %s", gears.Language.ContainerName)
	}

	return ret, gears.State
}

func MatchContainer(m *types.Container, match TypeMatchContainer) (bool, *gearConfig.GearConfig) {
	var ok bool
	gc := gearConfig.New(nil)

	for range onlyOnce {
		if MatchTag("<none>:<none>", m.Names) {
			ok = false
			break
		}

		gc.State = gc.ParseJson(m.Labels["gearbox.json"])
		if gc.State.IsError() {
			ok = false
			break
		}

		if gc.Meta.Organization != DefaultOrganization {
			ok = false
			break
		}

		tagCheck := fmt.Sprintf("%s/%s:%s", match.Organization, match.Name, match.Version)
		if m.Image == tagCheck {
			ok = true
			break
		}

		if gc.Meta.Name != match.Name {
			//if !RunAs.AsLink {
			if gc.Runtime.IsRunningAsFile() {
				ok = false
				break
			}

			cs := gc.MatchCommand(match.Name)
			if cs == nil {
				ok = false
				break
			}

			match.Name = gc.Meta.Name
		}

		if match.Version == "all" {
			ok = true
			break
		}

		if !gc.Versions.HasVersion(match.Version) {
			ok = false
			break
		}

		if match.Version == gearConfig.LatestName {
			gl := gc.Versions.GetLatest()
			if match.Version != "" {
				match.Version = gl
			}
		}

		for range onlyOnce {
			if m.Labels["gearbox.version"] == match.Version {
				ok = true
				break
			}

			if m.Labels["container.majorversion"] == match.Version {
				ok = true
				break
			}

			ok = false
		}
		break
	}

	return ok, gc
}

//func (gears *Gears) updateContainer(c *types.Container, id string) *ux.State {
//	if state := gears.IsNil(); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		if c == nil {
//			c, gears.State = gears.Docker.GetContainerById(id)
//			if gears.State.IsNotOk() {
//				break
//			}
//		}
//
//		if _, ok := c.Labels["gearbox.json"]; !ok {
//			continue
//		}
//
//		gear := NewGear(gears.Runtime, gears.Docker)
//		//gear.Docker = gears.Docker
//		gears.State = gear.GearConfig.ParseJson(c.Labels["gearbox.json"])
//		if gears.State.IsError() {
//			continue
//		}
//
//		if gear.GearConfig.Meta.Organization != DefaultOrganization {
//			continue
//		}
//
//		gear.Container.ID = c.ID
//		gear.Container.Name = gear.GearConfig.GetName()
//		gear.Container.Version = c.Labels["gearbox.version"]
//		gear.Container.Summary = *c
//		gear.Container.Docker = gear.Docker
//		gear.Container.GearConfig = gear.GearConfig
//		gear.Container.Details, _ = gear.Docker.ContainerInspect(c.ID)
//		//gear.State.RunState = c.State
//		gear.State.RunState = c.State
//		if c.Labels["container.latest"] == "true" {
//			gear.Container.IsLatest = true
//		}
//
//		//gear.Image.ID = gear.Container.Summary.ImageID
//		if _, ok := gears.Array[c.ImageID]; ok {
//			gears.Array[c.ImageID].Container = gear.Container
//		} else {
//			gears.Array[c.ID] = gear
//		}
//	}
//
//	return gears.State
//}


// ******************************************************************************** //

func (gears *Gears) NetworkList(name string) *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.Docker.NetworkList(name)
		if gears.State.IsNotOk() {
			break
		}

		ux.PrintflnCyan("\nConfigured %s networks:", gears.Language.AppName)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Name",
			"Driver",
			"Subnet",
		})

		for _, c := range gears.Docker.Networks {
			n := ""
			if len(c.IPAM.Config) > 0 {
				n = c.IPAM.Config[0].Subnet
			}

			t.AppendRow([]interface{}{
				ux.SprintfWhite(c.Name),
				ux.SprintfWhite(c.Driver),
				ux.SprintfWhite(n),
			})
		}

		t.Render()
	}

	return gears.State
}

func (gears *Gears) FindNetwork(name string) *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if name == "" {
			gears.State.SetError("empty %s name", gears.Language.ContainerName)
			break
		}

		gears.State = gears.Docker.NetworkList(name)
		if gears.State.IsNotOk() {
			break
		}

		for _, c := range gears.Docker.Networks {
			if c.Name == name {
				gears.State.SetOk("found")
				break
			}
		}
	}

	return gears.State
}

func (gears *Gears) NetworkCreate(netName string) *ux.State {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.FindNetwork(netName)
		if gears.State.IsError() {
			break
		}
		if gears.State.IsOk() {
			break
		}

		options := types.NetworkCreate {
			CheckDuplicate: true,
			Driver:         "bridge",
			Scope:          "local",
			EnableIPv6:     false,
			IPAM:           &network.IPAM {
				Driver:  "default",
				Options: nil,
				Config:  []network.IPAMConfig {
					{
						Subnet: "172.42.0.0/24",
						Gateway: "172.42.0.1",
					},
				},
			},
			Internal:       false,
			Attachable:     false,
			Ingress:        false,
			ConfigOnly:     false,
			ConfigFrom:     nil,
			Options:        nil,
			Labels:         nil,
		}

		gears.State = gears.Docker.NetworkCreate(netName, options)
	}

	return gears.State
}
