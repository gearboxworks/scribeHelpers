package toolDocker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"strings"
)


// List all images
// List the images on your Engine, similar to docker image ls:
// func ImageList(f types.ImageListOptions) error {
func (d *TypeDocker) ImageList(f string) (int, *ux.State) {
	var count int
	if state := d.IsNil(); state.IsError() {
		return 0, state
	}

	for range onlyOnce {
		df := filters.NewArgs()
		//if f != "" {
		//	df.Add("label", f)
		//}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		images, err := d.Client.ImageList(ctx, types.ImageListOptions{All: true, Filters: df})
		if err != nil {
			d.State.SetError("image list error: %s", err)
			break
		}

		ux.PrintfCyan("Downloaded TypeDocker images: ")
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Image", "Size"})

		for _, i := range images {
			if len(i.RepoTags) == 0 {
				continue
			}

			if i.RepoTags[0] == "<none>:<none>" {
				continue
			}

			t.AppendRow([]interface{}{
				ux.SprintfWhite(i.RepoTags[0]),
				ux.SprintfWhite(humanize.Bytes(uint64(i.Size))),
			})
		}

		d.State.ClearError()
		count = t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		ux.PrintflnGreen("%d found", count)
		t.Render()
		ux.PrintflnBlue("")
	}

	return count, d.State
}


func (d *TypeDocker) FindImage(org string, name string, version string) (bool, *ux.State) {
	var ok bool
	if state := d.IsNil(); state.IsError() {
		return false, state
	}

	for range onlyOnce {
		if name == "" {
			d.State.SetError("empty name")
			break
		}

		if version == "" {
			version = "latest"
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		images, err := d.Client.ImageList(ctx, types.ImageListOptions{All: true})
		if err != nil {
			d.State.SetError("image list error: %s", err)
			break
		}

		if len(images) == 0 {
			break
		}

		tagCheck := fmt.Sprintf("%s/%s:%s", org, name, version)
		// Start out with "not found". Will be cleared if found or error occurs.
		d.State.SetWarning("Image '%s' doesn't exist.", tagCheck)

		for _, i := range images {
			if !MatchTag(tagCheck, i.RepoTags) {
				continue
			}

			d.Image.Name = name
			d.Image.Version = version
			d.Image.Summary = &i
			d.Image.ID = i.ID
			d.Image.State = d.Image.State.EnsureNotNil()
			ok = true
			d.State.SetOk("Found image '%s'.", tagCheck)
			break
		}

		if d.State.IsNotOk() {
			if !ok {
				d.State.ClearError()
			}
			break
		}

		if d.Image.Summary == nil {
			break
		}

		ctx2, cancel2 := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel2()
		d.Image.Details, _, err = d.Client.ImageInspectWithRaw(ctx2, d.Image.ID)
		if err != nil {
			d.State.SetError("error inspecting d: %s", err)
			break
		}

		d.State.SetOk("found image")
	}

	return ok, d.State
}


// Search for an image in remote registry.
func (d *TypeDocker) Search(org string, name string, version string) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var repo string
		if version == "" {
			repo = fmt.Sprintf("%s/%s", org, name)
		} else {
			repo = fmt.Sprintf("%s/%s:%s", org, name, version)
		}

		ctx := context.Background()
		//ctx, cancel := context.WithTimeout(context.Background(), Timeout * 1000)
		//defer cancel()

		df := filters.NewArgs()
		repo = name

		images, err := d.Client.ImageSearch(ctx, repo, types.ImageSearchOptions{Filters: df, Limit: 100})
		if err != nil {
			d.State.SetError("image search error: %s", err)
			break
		}

		for _, v := range images {
			if !strings.HasPrefix(v.Name, org + "/") {
				continue
			}
			fmt.Printf("%s - %s\n", v.Name, v.Description)
		}
	}

	return d.State
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
