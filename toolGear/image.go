package toolGear

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/gearboxworks/scribeHelpers/toolGear/gearConfig"
	"github.com/gearboxworks/scribeHelpers/toolNetwork"
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
	"github.com/jedib0t/go-pretty/table"
	"io"
	"os"
	"strconv"
	"strings"
)


type Image struct {
	Language   *Language
	ID         string
	Name       string
	Version    string

	Summary    types.ImageSummary
	Details    types.ImageInspect
	GearConfig *gearConfig.GearConfig

	Docker    *Docker

	runtime    *toolRuntime.TypeRuntime
	State      *ux.State
}


func NewImage(runtime *toolRuntime.TypeRuntime) *Image {
	runtime = runtime.EnsureNotNil()

	i := Image {
		ID:         "",
		Name:       "",
		Version:    "",
		Summary:    types.ImageSummary{},
		Details:    types.ImageInspect{},
		GearConfig: gearConfig.New(runtime),
		Docker:    nil,

		runtime:    runtime,
		State:      ux.NewState(runtime.CmdName, runtime.Debug),
	}
	i.State.SetPackage("")
	i.State.SetFunctionCaller()
	return &i
}

func (i *Image) EnsureNotNil() *Image {
	for range onlyOnce {
		if i == nil {
			//goland:noinspection ALL
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

	for range onlyOnce {
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

		if i.Docker.Client == nil {
			i.State.SetError("docker client is nil")
			break
		}
	}

	return i.State
}


func (i *Image) Refresh() *ux.State {
	if state := i.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if i.Summary.ID == "" {
			ctx, cancel := context.WithTimeout(context.Background(), i.Docker.Provider.Timeout)
			//noinspection GoDeferInLoop
			defer cancel()

			df := filters.NewArgs()
			//df.Add("id", i.ID)
			df.Add("reference", fmt.Sprintf("%s/%s:%s", DefaultOrganization, i.Name, i.Version))

			var images []types.ImageSummary
			var err error
			images, err = i.Docker.Client.ImageList(ctx, types.ImageListOptions{All: true, Filters: df})
			if err != nil {
				i.State.SetError("gear list error: %s", err)
				break
			}
			if len(images) == 0 {
				i.State.SetWarning("no gears found")
				break
			}

			i.Summary = images[0]
			i.ID = i.Summary.ID

			i.Details, _, err = i.Docker.Client.ImageInspectWithRaw(ctx, i.ID)
			if err != nil {
				i.State.SetError("gear inspect error: %s", err)
				break
			}
		}

		if i.GearConfig.IsNotValid() {
			//i.GearConfig = gearConfig.New(nil)
			i.GearConfig.ParseJson(i.Summary.Labels["gearbox.json"])
			if i.GearConfig.State.IsError() {
				i.State = i.GearConfig.State
				break
			}
		}

		if i.GearConfig.Meta.Organization != DefaultOrganization {
			i.State.SetError("not a valid image")
			break
		}
	}

	return i.State
}


// Pull an image
// Pull an image, like docker pull:
func (i *Image) Pull() *ux.State {
	if state := i.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//goland:noinspection SpellCheckingInspection
		i.State = i.Docker.Pull("gearboxworks", i.Name, i.Version)
	}

	return i.State
}


// Pull an image with authentication
// Pull an image, like docker pull, with authentication:
func (i *Image) ImageAuthPull() *ux.State {
	if state := i.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
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

		ctx, cancel := context.WithTimeout(context.Background(), i.Docker.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		out, err := i.Docker.Client.ImagePull(ctx, "alpine", types.ImagePullOptions{RegistryAuth: authStr})
		if err != nil {
			i.State.SetError("error pulling gear: %s", err)
			break
		}

		//goland:noinspection ALL
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

	for range onlyOnce {
		//options := types.ImageRemoveOptions {
		//	Force:         true,
		//	PruneChildren: true,
		//}

		i.State = i.Docker.ImageRemove(i.ID, nil)
		if i.State.IsNotOk() {
			break
		}

		i.State.SetOk("removed image i %s:%s", i.Name, i.Version)
	}

	return i.State
}


func (i *Image) GetName() string {
	if i.Details.RepoTags == nil {
		return ""
	}

	if len(i.Details.RepoTags) == 0 {
		return ""
	}

	return strings.TrimPrefix(i.Details.RepoTags[0], "/")
}

func (i *Image) GetVersion() string {
	if i.Summary.Labels == nil {
		return ""
	}

	if len(i.Summary.Labels) == 0 {
		return ""
	}

	return i.Summary.Labels["gearbox.version"]
}

func (i *Image) GetSize() uint64 {
	return uint64(i.Summary.Size)
}

func (i *Image) GetLabels() map[string]string {
	return i.Summary.Labels
}

//type GearPort struct {
//	Name string
//	Value types.Port
//	Available bool
//}
//type GearPortMap map[string]GearPort

func (i *Image) GetPorts() Ports {
	ports := make(Ports)

	for range onlyOnce {
		scan := toolNetwork.New()
		if scan.State.IsNotOk() {
			i.State = scan.State
			break
		}
		i.State = scan.GetPorts()
		if i.State.IsNotOk() {
			break
		}

		for k, v := range *i.GearConfig.GetFixedPorts() {
			p, err := strconv.Atoi(v)
			if err != nil {
				continue
			}

			ports[uint16(p)] = &Port {
				Name: k,
				Available: scan.IsAvailable(uint16(p)),
				Port: types.Port {
					IP:          "0.0.0.0",
					PrivatePort: 0,
					PublicPort:  uint16(p),
					Type:        "tcp",
				},
			}
		}

		//for p := range i.Details.ContainerConfig.ExposedPorts {
		//	ports[uint16(p.Int())] = &Port {
		//		Name: "",
		//		Available: scan.IsAvailable(uint16(p.Int())),
		//		Port: types.Port{
		//			IP:          "",
		//			PrivatePort: 0,
		//			PublicPort:  0,
		//			Type:        "",
		//		},
		//	}
		//}

	}

	return ports
}

func (i *Image) ListPorts() *ux.State {
	if state := i.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//var err error
		//if gear.IsNotRunning() {
		//	break
		//}

		ports := i.GetPorts()
		if len(ports) == 0 {
			break
		}

		ux.PrintfCyan("Open ports for Image: %s:%s\n", i.Name, i.Version)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"Container",
			"Port Name",
			"Host Port",
			"Free",
		})

		used := 0
		for _, v := range ports {
			if v.PrivatePort == 22 {
				t.AppendRow([]interface{} {
					ux.SprintfYellow("%s:%s\n", i.Name, i.Version),
					ux.SprintfYellow("ssh"),
					ux.SprintfYellow("%s:%d", v.IP, v.PublicPort),
					ux.SprintfYellow("Yes"),
				})
				continue
			}

			if v.Available {
				t.AppendRow([]interface{} {
					ux.SprintfGreen("%s:%s\n", i.Name, i.Version),
					ux.SprintfGreen(v.Name),
					ux.SprintfGreen("%s:%d", v.IP, v.PublicPort),
					ux.SprintfGreen("Yes"),
				})
				continue
			}

			t.AppendRow([]interface{} {
				ux.SprintfRed("%s:%s\n", i.Name, i.Version),
				ux.SprintfRed(v.Name),
				ux.SprintfRed("%s:%d", v.IP, v.PublicPort),
				ux.SprintfRed("No"),
			})
			used++
		}

		count := t.Length()
		if count == 0 {
			ux.PrintfYellow("None found\n")
			break
		}

		t.Render()
		ux.PrintflnGreen("Ports found: %d", count)

		if used > 0 {
			ux.PrintflnRed("Warning: There are ports that are being used.")
		}
		ux.PrintflnBlue("")

		i.State.SetOk("")
	}

	return i.State
}
