package toolGear

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/newclarity/scribeHelpers/toolGear/gearConfig"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
)


// List all images
// List the images on your Engine, similar to docker image ls:
func (gears *Gears) GetImages(name string) (*ux.State) {
	if state := gears.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		gears.State = gears.Docker.ImageList(&types.ImageListOptions{All: true})
		if gears.State.IsNotOk() {
			break
		}

		for _, c := range gears.Docker.Images {
			if _, ok := c.Labels["gearbox.json"]; !ok {
				continue
			}

			gear := NewGear(gears.Runtime)
			gear.Docker = gears.Docker

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


func (gears *Gears) FindImage(gearName string, gearVersion string) (bool, *ux.State) {
	var ok bool
	if state := gears.IsNil(); state.IsError() {
		return false, state
	}

	for range onlyOnce {
		if gearName == "" {
			gears.State.SetError("empty gear name")
			break
		}

		if gearVersion == "" {
			gearVersion = "latest"
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

	return ok, gears.State
}

//func (gear *Gears) FindImage(gearName string, gearVersion string) (bool, *ux.State) {
//	var ok bool
//	if state := gear.IsNil(); state.IsError() {
//		return false, state
//	}
//
//	for range onlyOnce {
//		if gearName == "" {
//			gear.State.SetError("empty gear name")
//			break
//		}
//
//		if gearVersion == "" {
//			gearVersion = "latest"
//		}
//
//		gear.State = gear.Docker.ImageList(types.ImageListOptions{})
//		if gear.State.IsNotOk() {
//			break
//		}
//		if len(gear.Docker.Images) == 0 {
//			break
//		}
//
//		// Start out with "not found". Will be cleared if found or error occurs.
//		gear.State.SetWarning("Gear image '%s:%s' doesn't exist.", gearName, gearVersion)
//		for _, i := range gear.Docker.Images {
//			//var gc *gearConfig.GearConfig
//			ok, gear.Selected.GearConfig = MatchImage(&i,
//				TypeMatchImage{Organization: DefaultOrganization, Name: gearName, Version: gearVersion})
//			if !ok {
//				continue
//			}
//
//			gear.Selected.Image.Name = gearName
//			gear.Selected.Image.Version = gearVersion
//			//gear.Selected.Image.GearConfig = gc
//			gear.Selected.Image.Summary = i
//			gear.Selected.Image.ID = i.ID
//			gear.Selected.Image.State = gear.Selected.Image.State.EnsureNotNil()
//			//gear.Image.client = gear.DockerClient
//			ok = true
//			gear.State.SetOk("Found Gear image '%s:%s'.", gearName, gearVersion)
//			break
//		}
//
//		if gear.State.IsNotOk() {
//			if !ok {
//				gear.State.ClearError()
//			}
//			break
//		}
//
//		if gear.Selected.Image.Summary.ID == "" {
//			break
//		}
//
//		gear.Selected.Image.Details, gear.State = gear.Docker.ImageInspectWithRaw(gear.Selected.Image.ID)
//
//		gear.State.SetOk("found image")
//	}
//
//	return ok, gear.State
//}


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

		//repo = gearName
		//ctx := context.Background()
		////ctx, cancel := context.WithTimeout(context.Background(), Timeout * 1000)
		////defer cancel()
		//df := filters.NewArgs()
		////df.Add("name", "terminus")
		//images, err := gear.Docker.Client.ImageSearch(ctx, repo, types.ImageSearchOptions{Filters: df, Limit: 100})
		//if err != nil {
		//	gear.State.SetError("gear image search error: %s", err)
		//	break
		//}
		var resp []registry.SearchResult
		resp, gears.State = gears.Docker.ImageSearch(repo, nil)

		for _, v := range resp {
			if !strings.HasPrefix(v.Name, "gearboxworks/") {
				continue
			}
			fmt.Printf("%s - %s\n", v.Name, v.Description)
		}
	}

	return gears.State
}


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

		tagCheck := fmt.Sprintf("%s/%s:%s", match.Organization, match.Name, match.Version)
		if !MatchTag(tagCheck, m.RepoTags) {
			ok = false
			break
		}

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

		if match.Version == "latest" {
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
