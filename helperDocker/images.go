package helperDocker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
	"github.com/newclarity/scribeHelpers/helperGear/gearConfig"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"strings"
)


// List all images
// List the images on your Engine, similar to docker image ls:
// func ImageList(f types.ImageListOptions) error {
func (gear *DockerGear) ImageList(f string) (int, *ux.State) {
	var count int
	if state := gear.IsNil(); state.IsError() {
		return 0, state
	}

	for range OnlyOnce {
		df := filters.NewArgs()
		//if f != "" {
		//	df.Add("label", f)
		//}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		images, err := gear.Client.ImageList(ctx, types.ImageListOptions{All: true, Filters: df})
		if err != nil {
			gear.State.SetError("gear image list error: %s", err)
			break
		}

		ux.PrintfCyan("Downloaded Gearbox images: ")
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Class", "Image", "Ports", "Size"})

		gc := gearConfig.New(gear.Runtime)
		for _, i := range images {
			gear.State = gc.ParseJson(i.Labels["gearbox.json"])
			if gear.State.IsError() {
				continue
			}

			if gc.Meta.Organization != DefaultOrganization {
				continue
			}

			if len(i.RepoTags) == 0 {
				continue
			}

			if i.RepoTags[0] == "<none>:<none>" {
				continue
			}

			if f != "" {
				if gc.Meta.Name != f {
					continue
				}
			}

			// foo := fmt.Sprintf("%s/%s", gc.Organization, gc.Name)
			t.AppendRow([]interface{}{
				ux.SprintfWhite(gc.Meta.Class),
				//ux.SprintfWhite(gc.Meta.State),
				ux.SprintfWhite(i.RepoTags[0]),
				ux.SprintfWhite(gc.Build.Ports.ToString()),
				ux.SprintfWhite(humanize.Bytes(uint64(i.Size))),
			})
		}

		gear.State.ClearError()
		count = t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		ux.PrintflnGreen("%d found", count)
		t.Render()
		ux.PrintflnBlue("")
	}

	return count, gear.State
}


func (gear *DockerGear) FindImage(gearName string, gearVersion string) (bool, *ux.State) {
	var ok bool
	if state := gear.IsNil(); state.IsError() {
		return false, state
	}

	for range OnlyOnce {
		if gearName == "" {
			gear.State.SetError("empty gear name")
			break
		}

		if gearVersion == "" {
			gearVersion = "latest"
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		images, err := gear.Client.ImageList(ctx, types.ImageListOptions{All: true})
		if err != nil {
			gear.State.SetError("gear image list error: %s", err)
			break
		}

		if len(images) == 0 {
			break
		}

		// Start out with "not found". Will be cleared if found or error occurs.
		gear.State.SetWarning("Gear image '%s:%s' doesn't exist.", gearName, gearVersion)

		for _, i := range images {
			var gc *gearConfig.GearConfig
			ok, gc = MatchImage(&i,
				TypeMatchImage{Organization: DefaultOrganization, Name: gearName, Version: gearVersion})
			if !ok {
				continue
			}

			gear.Image.Name = gearName
			gear.Image.Version = gearVersion
			gear.Image.GearConfig = gc
			gear.Image.Summary = &i
			gear.Image.ID = i.ID
			gear.Image.State = gear.Image.State.EnsureNotNil()
			//gear.Image.client = gear.DockerClient
			ok = true

			break
		}

		if gear.State.IsNotOk() {
			if !ok {
				gear.State.ClearError()
			}
			break
		}

		if gear.Image.Summary == nil {
			break
		}

		ctx2, cancel2 := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel2()
		gear.Image.Details, _, err = gear.Client.ImageInspectWithRaw(ctx2, gear.Image.ID)
		if err != nil {
			gear.State.SetError("error inspecting gear: %s", err)
			break
		}

		gear.State.SetOk("found image")
	}

	return ok, gear.State
}


// Search for an image in remote registry.
func (gear *DockerGear) Search(gearName string, gearVersion string) *ux.State {
	if state := gear.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		var repo string
		if gearVersion == "" {
			repo = fmt.Sprintf("gearboxworks/%s", gearName)
		} else {
			repo = fmt.Sprintf("gearboxworks/%s:%s", gearName, gearVersion)
		}

		ctx := context.Background()
		//ctx, cancel := context.WithTimeout(context.Background(), Timeout * 1000)
		//defer cancel()

		df := filters.NewArgs()
		//df.Add("name", "terminus")
		repo = gearName

		images, err := gear.Client.ImageSearch(ctx, repo, types.ImageSearchOptions{Filters: df, Limit: 100})
		if err != nil {
			gear.State.SetError("gear image search error: %s", err)
			break
		}

		for _, v := range images {
			if !strings.HasPrefix(v.Name, "gearboxworks/") {
				continue
			}
			fmt.Printf("%s - %s\n", v.Name, v.Description)
		}
	}

	return gear.State
}


//func MatchImage(m *types.ImageSummary, gearOrg string, gearName string, gearVersion string) (bool, *gearConfig.GearConfig) {
func MatchImage(m *types.ImageSummary, match TypeMatchImage) (bool, *gearConfig.GearConfig) {
	var ok bool
	gc := gearConfig.New(gear.Runtime)

	for range OnlyOnce {
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

		for range OnlyOnce {
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
