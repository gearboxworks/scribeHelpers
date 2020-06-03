package helperGear

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/newclarity/scribeHelpers/helperGear/gearConfig"
	"github.com/newclarity/scribeHelpers/helperRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"io"
	"os"
	"strings"
	"time"
)

type Image struct {
	ID         string
	Name       string
	Version    string

	Summary    *types.ImageSummary
	Details    types.ImageInspect
	GearConfig *gearConfig.GearConfig

	_Parent    *DockerGear
	Debug      bool
	State      *ux.State
}


func NewImage(runtime *helperRuntime.TypeRuntime) *Image {
	runtime = runtime.EnsureNotNil()

	i := Image{
		ID:         "",
		Name:       "",
		Version:    "",
		Summary:    nil,
		Details:    types.ImageInspect{},
		GearConfig: gearConfig.New(runtime),
		_Parent:    nil,

		Debug:      runtime.Debug,
		State:      ux.NewState(runtime.CmdName, runtime.Debug),
	}
	i.State.SetPackage("")
	i.State.SetFunctionCaller()
	return &i
}

func (i *Image) EnsureNotNil() *Image {
	for range OnlyOnce {
		if i == nil {
			i = NewImage(nil)
		}
		i.State = i.State.EnsureNotNil()
	}
	return i
}

func (i *Image) IsNil() *ux.State {
	if state := ux.IfNilReturnError(i); state.IsError() {
		return state
	}
	i.State = i.State.EnsureNotNil()
	return i.State
}

func (i *Image) IsValid() *ux.State {
	if state := ux.IfNilReturnError(i); state.IsError() {
		return state
	}

	for range OnlyOnce {
		i.State = i.State.EnsureNotNil()

		if i.ID == "" {
			i.State.SetError("ID is nil")
			break
		}

		if i.Name == "" {
			i.State.SetError("name is nil")
			break
		}

		if i.Version == "" {
			i.State.SetError("version is nil")
			break
		}

		if i._Parent.Client == nil {
			i.State.SetError("docker client is nil")
			break
		}
	}

	return i.State
}


func (i *Image) Status() *ux.State {
	if state := i.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		if i.Summary == nil {
			ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
			//noinspection GoDeferInLoop
			defer cancel()

			df := filters.NewArgs()
			//df.Add("id", i.ID)
			df.Add("reference", fmt.Sprintf("%s/%s:%s", DefaultOrganization, i.Name, i.Version))

			var images []types.ImageSummary
			var err error
			images, err = i._Parent.Client.ImageList(ctx, types.ImageListOptions{All: true, Filters: df})
			if err != nil {
				i.State.SetError("gear list error: %s", err)
				break
			}
			if len(images) == 0 {
				i.State.SetWarning("no gears found")
				break
			}

			i.Summary = &images[0]
			i.Summary.ID = i.ID

			d := types.ImageInspect{}
			d, _, err = i._Parent.Client.ImageInspectWithRaw(ctx, i.ID)
			if err != nil {
				i.State.SetError("gear inspect error: %s", err)
				break
			}
			i.Details = d
		}

		if i.GearConfig == nil {
			i.GearConfig = gearConfig.New(nil)
			i.GearConfig.ParseJson(i.Summary.Labels["gearbox.json"])
			if i.GearConfig.State.IsError() {
				i.State.SetState(i.GearConfig.State)
				break
			}
		}

		if i.GearConfig.Meta.Organization != DefaultOrganization {
			i.State.SetError("not a Gearbox container")
			break
		}
	}

	if i.State.IsError() {
		i.Summary = nil
	}

	return i.State
}


// Pull an image
// Pull an image, like docker pull:
func (i *Image) Pull() *ux.State {
	if state := i.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		var repo string
		if i.Version == "" {
			repo = fmt.Sprintf("gearboxworks/%s", i.Name)
		} else {
			repo = fmt.Sprintf("gearboxworks/%s:%s", i.Name, i.Version)
		}

		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
		//noinspection GoDeferInLoop
		defer cancel()

		//df := filters.NewArgs()
		//df.Add("name", "terminus")
		//var results []registry.SearchResult
		//results, err = i.client.ImageSearch(ctx, "", types.ImageSearchOptions{Filters:df})
		//for _, v := range results {
		//	fmt.Printf("%s - %s\n", v.Name, v.Description)
		//}

		var out io.ReadCloser
		var err error
		out, err = i._Parent.Client.ImagePull(ctx, repo, types.ImagePullOptions{All: false})
		if err != nil {
			i.State.SetError("Error pulling Gear %s:%s - %s", i.Name, i.Version, err)
			break
		}

		//noinspection GoDeferInLoop
		defer out.Close()

		ux.PrintflnNormal("Pulling Gear %s:%s.", i.Name, i.Version)
		d := json.NewDecoder(out)
		var event *PullEvent
		for {
			err := d.Decode(&event)
			if err != nil {
				if err == io.EOF {
					break
				}

				i.State.SetError("Error pulling Gear %s:%s - %s", i.Name, i.Version, err)
				break
			}

			// fmt.Printf("EVENT: %+v\n", event)
			ux.Printf("%+v\r", event.Progress)
		}
		ux.Printf("\n")

		if i.State.IsError() {
			break
		}

		// Latest event for new i
		// EVENT: {Status:Status: Downloaded newer i for busybox:latest Error: Progress:[==================================================>]  699.2kB/699.2kB ProgressDetail:{Current:699243 Total:699243}}
		// Latest event for up-to-date i
		// EVENT: {Status:Status: Image is up to date for busybox:latest Error: Progress: ProgressDetail:{Current:0 Total:0}}
		if event != nil {
			if strings.HasPrefix(event.Status, "Status: Downloaded newer") {
				// new
				ux.PrintfOk("Pulling Gear %s:%s - OK.\n", i.Name, i.Version)
			} else if strings.HasPrefix(event.Status, "Status: Image is up to date for") {
				// up-to-date
				ux.PrintfOk("Pulling Gear %s:%s - updated.\n", i.Name, i.Version)
			} else {
				ux.PrintfWarning("Pulling Gear %s:%s - unknown state.\n", i.Name, i.Version)
			}
		}
		//ux.Printf("\nGear i pull OK: %+v\n", event)
		//ux.Printf("%s\n", event.Status)

		//buf := new(bytes.Buffer)
		//_, err = buf.ReadFrom(out)
		//fmt.Printf("%s", buf.String())
		//_, _ = io.Copy(os.Stdout, out)
	}

	return i.State
}

type PullEvent struct {
	Status         string `json:"status"`
	Error          string `json:"error"`
	Progress       string `json:"progress"`
	ProgressDetail struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"progressDetail"`
}


// Pull an image with authentication
// Pull an image, like docker pull, with authentication:
func (i *Image) ImageAuthPull() *ux.State {
	if state := i.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		authConfig := types.AuthConfig{
			Username: "username",
			Password: "password",
		}

		encodedJSON, err := json.Marshal(authConfig)
		if err != nil {
			i.State.SetError("error pulling gear: %s", err)
			break
		}
		authStr := base64.URLEncoding.EncodeToString(encodedJSON)

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		out, err := i._Parent.Client.ImagePull(ctx, "alpine", types.ImagePullOptions{RegistryAuth: authStr})
		if err != nil {
			i.State.SetError("error pulling gear: %s", err)
			break
		}

		//noinspection GoDeferInLoop
		defer out.Close()

		_, _ = io.Copy(os.Stdout, out)
	}

	return i.State
}


// Remove containers
// Now that you know what containers exist, you can perform operations on them.
// This example stops all running containers.
func (i *Image) Remove() *ux.State {
	if state := i.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		//noinspection GoDeferInLoop
		defer cancel()

		options := types.ImageRemoveOptions {
			Force:         true,
			PruneChildren: true,
		}

		_, err := i._Parent.Client.ImageRemove(ctx, i.ID, options)
		if err != nil {
			i.State.SetError("error removing gear: %s", err)
			break
		}

		i.State.SetOk("removed gear i %s:%s", i.Name, i.Version)
	}

	return i.State
}
