package toolGear

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

//"github.com/docker/docker/integration-cli/cli"
// DOCKER_HOST=tcp://macpro:2375

type Docker struct {
	Containers []types.Container
	Images     []types.ImageSummary
	Networks   []types.NetworkResource
	Registry   []registry.SearchResult

	Provider *Provider
	Client   *client.Client

	Runtime *toolRuntime.TypeRuntime
	State   *ux.State
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

func NewDocker(runtime *toolRuntime.TypeRuntime) *Docker {
	var d Docker

	for range onlyOnce {
		runtime = runtime.EnsureNotNil()

		d = Docker{
			Containers: nil,
			Images:     nil,

			Provider: NewProvider(runtime),
			Client:   nil,

			Runtime: runtime,
			State:   ux.NewState(runtime.CmdName, runtime.Debug),
		}

		d.State.SetPackage("")
		d.State.SetFunctionCaller()

		d.State = d.Connect()
		if d.State.IsNotOk() {
			break
		}
	}

	return &d
}

func (d *Docker) Connect() *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//foo := os.Getenv("DOCKER_HOST")
		//fmt.Printf("DOCKER_HOST:%s\n", foo)
		if os.Getenv("DOCKER_HOST") != "" {
			d.Provider.Remote = true
		}

		var err error
		d.Client, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		//cli.DockerClient, err = client.NewEnvClient()
		if d.inspectError(err, "Docker client error") {
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		_, err = d.Client.Ping(ctx)
		if d.inspectError(err, "can not connect to Docker service provider - maybe you haven't set DOCKER_HOST, or Docker not running on this host") {
			break
		}
	}

	return d.State
}

func (d *Docker) IsValid() *ux.State {
	if state := ux.IfNilReturnError(d); state.IsError() {
		return state
	}

	for range onlyOnce {
		d.State = d.State.EnsureNotNil()

		if d.Client == nil {
			d.State.SetError("docker client is nil")
			break
		}
	}

	return d.State
}

func (d *Docker) IsNil() *ux.State {
	if state := ux.IfNilReturnError(d); state.IsError() {
		return state
	}

	for range onlyOnce {
		d.State = d.State.EnsureNotNil()
	}

	return d.State
}

func (d *Docker) inspectError(err error, msg string, args ...interface{}) bool {
	failed := true
	if state := d.IsNil(); state.IsError() {
		//goland:noinspection GoBoolExpressions
		return failed
	}

	for range onlyOnce {
		if err == nil {
			failed = false
			d.State.SetOk()
			break
		}

		msg = fmt.Sprintf(msg, args...)
		d.State.SetError(msg+" (%s)", err)

		if d.State.ErrorHas(ErrorDockerTimeout) {
			d.State.SetError("timeout contacting provider - currently set to %s, trying increasing", d.Provider.Timeout.String())
			break
		}
	}

	return failed
}

func (d *Docker) DetermineTimeout(from int, to int) time.Duration {
	var val time.Duration
	if state := d.IsNil(); state.IsError() {
		return val
	}

	for range onlyOnce {
		var da []time.Duration
		for cd := from; cd < to; cd++ {
			foo := time.Duration(cd) * time.Second
			da = append(da, foo)
		}

		//da := []time.Duration {
		//	1 * time.Second,
		//	2 * time.Second,
		//	3 * time.Second,
		//	4 * time.Second,
		//	5 * time.Second,
		//	6 * time.Second,
		//	7 * time.Second,
		//	8 * time.Second,
		//	9 * time.Second,
		//	10 * time.Second,
		//	11 * time.Second,
		//	12 * time.Second,
		//	13 * time.Second,
		//	14 * time.Second,
		//}

		for _, i := range da {
			//ux.PrintfBlue("\tTesting timeout -> %s", i.String())
			fmt.Printf(".")
			if !d.testIsTimeout(i) {
				val = i
				//ux.PrintflnGreen(" OK")
				break
			}
		}
		fmt.Printf("\n")

		if val == 0 {
			d.State.SetError("Timeout while attempting to connect to server - cannot determine max timeout value.")
			break
		}
	}

	return val
}

func (d *Docker) testIsTimeout(to time.Duration) bool {
	ok := true
	if state := d.IsNil(); state.IsError() {
		//goland:noinspection GoBoolExpressions
		return ok
	}

	for range onlyOnce {
		var err error

		ctx, cancel := context.WithTimeout(context.Background(), to)
		//noinspection GoDeferInLoop
		defer cancel()

		options := &types.ImageListOptions{All: true, Filters: filters.NewArgs()}
		d.Images, err = d.Client.ImageList(ctx, *options)
		if err == nil {
			ok = false
			break
		}

		d.State.SetError(err)

		if d.State.ErrorHas(ErrorDockerTimeout) {
			break
		}

		// Only report timeout errors.
		ok = true
	}

	return ok
}

// ******************************************************************************** //

//func (d *Docker) ImageList(options *types.ImageListOptions) *ux.State {
func (d *Docker) ImageList(filter *filters.Args) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var err error
		if filter == nil {
			f := filters.NewArgs()
			filter = &f
		}

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		options := &types.ImageListOptions{All: true, Filters: *filter}
		d.Images, err = d.Client.ImageList(ctx, *options)
		if d.inspectError(err, "image list error") {
			break
		}

		if len(d.Images) == 0 {
			d.State.SetWarning("no images found")
			break
		}

		d.State.SetOk()
	}

	return d.State
}

func (d *Docker) FindImage(repo string) (types.ImageSummary, types.ImageInspect, *ux.State) {
	var is types.ImageSummary
	var ii types.ImageInspect
	if state := d.IsNil(); state.IsError() {
		return is, ii, state
	}

	for range onlyOnce {
		var err error
		filter := filters.NewArgs()
		//filter.Add("image_name", repo)

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		options := &types.ImageListOptions{All: false, Filters: filter}
		var isa []types.ImageSummary
		isa, err = d.Client.ImageList(ctx, *options)
		if d.inspectError(err, "image list error") {
			break
		}

		if len(isa) == 0 {
			d.State.SetWarning("no images found")
			break
		}

		for _, i := range isa {
			for _, rt := range i.RepoTags {
				if rt == repo {
					is = i
					break
				}
			}
			if is.ID != "" {
				break
			}
		}

		ii, d.State = d.ImageInspectWithRaw(is.ID)
		if d.State.IsNotOk() {
			break
		}

		d.State.SetOk()
	}

	return is, ii, d.State
}

func (d *Docker) ImageInspectWithRaw(imageID string) (types.ImageInspect, *ux.State) {
	var resp types.ImageInspect
	if state := d.IsNil(); state.IsError() {
		return resp, state
	}

	for range onlyOnce {
		var err error

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		resp, _, err = d.Client.ImageInspectWithRaw(ctx, imageID)
		if d.inspectError(err, "error inspecting image") {
			break
		}

		d.State.SetOk()
	}

	return resp, d.State
}

func (d *Docker) ImageSearch(repo string, options ...types.ImageSearchOptions) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var err error

		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		df := filters.NewArgs()
		//df.Add("name", "*")
		if len(options) == 0 {
			options = append(options, types.ImageSearchOptions{Filters: df, Limit: 100})
		}

		d.Registry, err = d.Client.ImageSearch(ctx, repo, (options[0]))
		if d.inspectError(err, "search error") {
			break
		}
		if len(d.Registry) == 0 {
			d.State.SetWarning("'%s' not found in registry", repo)
			break
		}
		sort.Sort(NameSorter(d.Registry))

		d.State.SetResponse(len(d.Registry))
		d.State.SetOk()
	}

	return d.State
}

type NameSorter []registry.SearchResult

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

func (d *Docker) ImageRemove(imageID string, options *types.ImageRemoveOptions) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var err error

		options = &types.ImageRemoveOptions{
			Force:         true,
			PruneChildren: true,
		}

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		_, err = d.Client.ImageRemove(ctx, imageID, *options)
		if d.inspectError(err, "error removing") {
			break
		}

		d.State.SetOk()
	}

	return d.State
}

func (d *Docker) Pull(user string, name string, version string) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var repo string
		if version == "" {
			repo = fmt.Sprintf("%s/%s", user, name)
		} else {
			repo = fmt.Sprintf("%s/%s:%s", user, name, version)
		}

		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
		//noinspection GoDeferInLoop
		defer cancel()

		//df := filters.NewArgs()
		//df.Add("name", "terminus")
		//var results []registry.SearchResult
		//results, err = d.client.ImageSearch(ctx, "", types.ImageSearchOptions{Filters:df})
		//for _, v := range results {
		//	fmt.Printf("%s - %s\n", v.Name, v.Description)
		//}

		var out io.ReadCloser
		var err error
		out, err = d.Client.ImagePull(ctx, repo, types.ImagePullOptions{All: false})
		if d.inspectError(err, "Error pulling Gear %s:%s", name, version) {
			break
		}

		//goland:noinspection GoDeferInLoop
		defer _close(out)

		ux.PrintflnNormal("Pulling Gear %s:%s.", name, version)
		dj := json.NewDecoder(out)
		var event *PullEvent
		for {
			err = dj.Decode(&event)
			if err != nil {
				if err == io.EOF {
					break
				}

				d.State.SetError("Error pulling Gear %s:%s - %s", name, version, err)
				break
			}

			// fmt.Printf("EVENT: %+v\n", event)
			ux.Printf("%+v\r", event.Progress)
		}
		ux.Printf("\n")

		if d.State.IsError() {
			break
		}

		// Latest event for new d
		// EVENT: {Status:Status: Downloaded newer d for busybox:latest Error: Progress:[==================================================>]  699.2kB/699.2kB ProgressDetail:{Current:699243 Total:699243}}
		// Latest event for up-to-date d
		// EVENT: {Status:Status: Image is up to date for busybox:latest Error: Progress: ProgressDetail:{Current:0 Total:0}}
		if event != nil {
			if strings.HasPrefix(event.Status, "Status: Downloaded newer") {
				// new
				ux.PrintfOk("Pulling Gear %s:%s - OK.\n", name, version)
			} else if strings.HasPrefix(event.Status, "Status: Image is up to date for") {
				// up-to-date
				ux.PrintfOk("Pulling Gear %s:%s - updated.\n", name, version)
			} else {
				ux.PrintfWarning("Pulling Gear %s:%s - unknown state.\n", name, version)
			}
		}
		//ux.Printf("\nGear d pull OK: %+v\n", event)
		//ux.Printf("%s\n", event.Status)
	}

	return d.State
}

func (d *Docker) ImageBuild(buildContext io.Reader, options types.ImageBuildOptions) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var err error

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout*40)
		//noinspection GoDeferInLoop
		defer cancel()

		var out types.ImageBuildResponse
		out, err = d.Client.ImageBuild(ctx, buildContext, options)
		if d.inspectError(err, "error building") {
			break
		}
		if out.Body == nil {
			d.State.SetError("Invalid response from Docker build")
			break
		}
		//goland:noinspection GoDeferInLoop
		defer _close(out.Body)

		ux.PrintflnNormal("Building image")
		termFd, isTerm := term.GetFdInfo(os.Stderr)
		err = jsonmessage.DisplayJSONMessagesStream(out.Body, os.Stderr, termFd, isTerm, nil)
		if err != nil {
			d.State.SetError(err)
			break
		}

		////////////////////////////////////////
		//buf := new(bytes.Buffer)
		//var i int64
		//i, err = io.Copy(buf, out.Body)
		//d.State.SetOutput(buf.String())
		//_, _ = io.Copy(os.Stdout, buf)
		//if err != nil {
		//	d.State.SetError(err)
		//	break
		//}
		//fmt.Printf("INT:%d\n", i)

		////////////////////////////////////////
		//dj := json.NewDecoder(out.Body)
		//var event *PullEvent
		//for {
		//	err = dj.Decode(&event)
		//	if err != nil {
		//		if err == io.EOF {
		//			break
		//		}
		//
		//		d.State.SetError("Error creating image - %s", err)
		//		break
		//	}
		//
		//	// fmt.Printf("EVENT: %+v\n", event)
		//	ux.Printf("%+v - %+v - %+v\n",
		//		event.Status,
		//		event.Progress,
		//		event.ProgressDetail,
		//		)
		//}
		//ux.Printf("\n")
		//if d.State.IsError() {
		//	break
		//}
		//// Latest event for new d
		//// EVENT: {Status:Status: Downloaded newer d for busybox:latest Error: Progress:[==================================================>]  699.2kB/699.2kB ProgressDetail:{Current:699243 Total:699243}}
		//// Latest event for up-to-date d
		//// EVENT: {Status:Status: Image is up to date for busybox:latest Error: Progress: ProgressDetail:{Current:0 Total:0}}
		//if event != nil {
		//	if strings.HasPrefix(event.Status, "Status: Downloaded newer") {
		//		// new
		//		ux.PrintflnOk("Creating image - OK.")
		//	} else if strings.HasPrefix(event.Status, "Status: Image is up to date for") {
		//		// up-to-date
		//		ux.PrintflnOk("Creating image - updated.")
		//	} else {
		//		ux.PrintflnWarning("Creating image - unknown state.")
		//	}
		//}
		////ux.Printf("\nGear d pull OK: %+v\n", event)
		////ux.Printf("%s\n", event.Status)

		d.State.SetOk()
	}

	return d.State
}

//func (d *Docker) PullRef(repo string) *ux.State {
//	if state := d.IsNil(); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
//		//noinspection GoDeferInLoop
//		defer cancel()
//
//		var out io.ReadCloser
//		var err error
//		out, err = d.Client.ImagePull(ctx, repo, types.ImagePullOptions{All: false})
//		if d.inspectError(err, "Error pulling image %s", repo) {
//			break
//		}
//
//		//goland:noinspection GoDeferInLoop
//		defer out.Close()
//
//		ux.PrintflnNormal("Pulling image %s.", repo)
//		dj := json.NewDecoder(out)
//		var event *PullEvent
//		for {
//			err = dj.Decode(&event)
//			if err != nil {
//				if err == io.EOF {
//					break
//				}
//
//				d.State.SetError("Error pulling image %s - %s", repo, err)
//				break
//			}
//
//			// fmt.Printf("EVENT: %+v\n", event)
//			ux.Printf("%+v\r", event.Progress)
//		}
//		ux.Printf("\n")
//
//		if d.State.IsError() {
//			break
//		}
//
//		// Latest event for new d
//		// EVENT: {Status:Status: Downloaded newer d for busybox:latest Error: Progress:[==================================================>]  699.2kB/699.2kB ProgressDetail:{Current:699243 Total:699243}}
//		// Latest event for up-to-date d
//		// EVENT: {Status:Status: Image is up to date for busybox:latest Error: Progress: ProgressDetail:{Current:0 Total:0}}
//		if event != nil {
//			if strings.HasPrefix(event.Status, "Status: Downloaded newer") {
//				// new
//				ux.PrintfOk("Pulling image %s - OK.\n", repo)
//			} else if strings.HasPrefix(event.Status, "Status: Image is up to date for") {
//				// up-to-date
//				ux.PrintfOk("Pulling image %s - updated.\n", repo)
//			} else {
//				ux.PrintfWarning("Pulling image %s - unknown state.\n", repo)
//			}
//		}
//		//ux.Printf("\nGear d pull OK: %+v\n", event)
//		//ux.Printf("%s\n", event.Status)
//
//		//buf := new(bytes.Buffer)
//		//_, err = buf.ReadFrom(out)
//		//fmt.Printf("%s", buf.String())
//		//_, _ = io.Copy(os.Stdout, out)
//	}
//
//	return d.State
//}

func (d *Docker) PullRepo(repo string) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var user string
		var name string
		var version string

		re := regexp.MustCompile("([a-zA-Z0-9_-]+)/([a-zA-Z0-9_-]+):([a-zA-Z0-9_.-]+)")
		match := re.FindStringSubmatch(repo)
		if len(match) >= 4 {
			version = match[3]
		}
		if len(match) >= 3 {
			name = match[2]
		}
		if len(match) >= 2 {
			user = match[1]
		}

		d.State = d.Pull(user, name, version)
	}

	return d.State
}

func (d *Docker) Tag(src string, target string) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout*40)
		//noinspection GoDeferInLoop
		defer cancel()

		err := d.Client.ImageTag(ctx, src, target)
		if err != nil {
			d.State.SetError(err)
			break
		}

		d.State.SetOk()
	}

	return d.State
}

// ******************************************************************************** //

func (d *Docker) GetContainerById(containerID string) (types.Container, *ux.State) {
	var c types.Container
	if state := d.IsNil(); state.IsError() {
		return c, state
	}

	for range onlyOnce {
		df := filters.NewArgs()
		df.Add("id", containerID)

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		containers, err := d.Client.ContainerList(ctx, types.ContainerListOptions{All: true, Filters: df})
		if d.inspectError(err, "gear list error") {
			break
		}
		if len(containers) == 0 {
			d.State.SetWarning("no gears found")
			break
		}
		c = containers[0]

		//c.Summary = &containers[0]
		//
		//c.GearConfig = c.GearConfig.EnsureNotNil()
		//c.State = c.GearConfig.ParseJson(c.Summary.Labels["gearbox.json"])
		//if c.State.IsError() {
		//	break
		//}
		//
		//if c.GearConfig.Meta.Organization != DefaultOrganization {
		//	c.State.SetError("not a Gearbox container")
		//	break
		//}
		//
		//d := types.ContainerJSON{}
		//d, err = c._Parent.Client.ContainerInspect(ctx, c.ID)
		//if err != nil {
		//	c.State.SetError("gear inspect error: %s", err)
		//	break
		//}
		//c.Details = &d
		//
		//c.State.SetRunState(c.Details.State.Status)
	}

	return c, d.State
}

func (d *Docker) ContainerList(filter *filters.Args, force bool) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if force {
			d.Containers = nil
		}
		if d.Containers != nil {
			break
		}

		if filter == nil {
			f := filters.NewArgs()
			filter = &f
		}

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		var err error
		options := types.ContainerListOptions{Size: true, All: true, Filters: *filter}
		d.Containers, err = d.Client.ContainerList(ctx, options)
		if d.inspectError(err, "container list error") {
			break
		}

		if len(d.Containers) == 0 {
			d.State.SetWarning("no containers found")
			break
		}
	}

	return d.State
}

func (d *Docker) ContainerStart(containerID string, options *types.ContainerStartOptions) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if d.Containers == nil {
			d.State.SetWarning("no containers found")
			break
		}

		if options == nil {
			options = &types.ContainerStartOptions{}
		}

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		//statusCh, errCh := c._Parent.Client.ContainerWait(ctx, c.ID, "") // container.WaitConditionNotRunning
		//select {
		//	case err := <-errCh:
		//		if err != nil {
		//			c.State.SetError("Docker client error: %s", err)
		//			// fmt.Printf("SC: %s\n", response.Error)
		//			// return false, err
		//		}
		//		break
		//
		//	case status := <-statusCh:
		//		fmt.Printf("status.StatusCode: %#+v\n", status.StatusCode)
		//		break
		//}
		// fmt.Printf("SC: %s\n", status)
		// fmt.Printf("SC: %s\n", err)

		err := d.Client.ContainerStart(ctx, containerID, *options)
		if d.inspectError(err, "Container start error") {
			break
		}
	}

	return d.State
}

func (d *Docker) ContainerStop(containerID string, timeout *time.Duration) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if d.Containers == nil {
			d.State.SetWarning("no containers found")
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		err := d.Client.ContainerStop(ctx, containerID, timeout)
		if d.inspectError(err, "container stop error") {
			break
		}
	}

	return d.State
}

func (d *Docker) ContainerRemove(containerID string, options *types.ContainerRemoveOptions) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if d.Containers == nil {
			d.State.SetWarning("no containers found")
			break
		}

		if options == nil {
			options = &types.ContainerRemoveOptions{}
		}

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		err := d.Client.ContainerRemove(ctx, containerID, *options)
		if d.inspectError(err, "container remove error") {
			break
		}
	}

	return d.State
}

func (d *Docker) ContainerLogs(containerID string, options types.ContainerLogsOptions) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if d.Containers == nil {
			d.State.SetWarning("no containers found")
			break
		}
		options = types.ContainerLogsOptions{ShowStdout: true}

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		// Replace this ID with a container that really exists
		out, err := d.Client.ContainerLogs(ctx, containerID, options)
		if d.inspectError(err, "container logs error") {
			break
		}

		buf := new(bytes.Buffer)
		_, _ = io.Copy(buf, out)
		d.State.SetOutput(buf.String())
		_, _ = io.Copy(os.Stdout, buf)

		d.State.SetOk()
	}

	return d.State
}

func (d *Docker) ContainerCommit() *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if d.Containers == nil {
			d.State.SetWarning("no containers found")
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		createResp, err := d.Client.ContainerCreate(ctx,
			&container.Config{
				Image: "alpine",
				Cmd:   []string{"touch", "/helloworld"},
			},
			nil,
			nil,
			nil, // TODO verify nil is ok
			"")
		if d.inspectError(err, "container create error") {
			break
		}

		err = d.Client.ContainerStart(ctx, createResp.ID, types.ContainerStartOptions{})
		if d.inspectError(err, "container start error") {
			break
		}

		statusCh, errCh := d.Client.ContainerWait(ctx, createResp.ID, container.WaitConditionNotRunning)
		select {
		case err := <-errCh:
			if err != nil {
				//response.State.SetError("gear stop error: %s", err)
				break
			}
		case <-statusCh:
		}

		commitResp, err := d.Client.ContainerCommit(ctx, createResp.ID, types.ContainerCommitOptions{Reference: "helloworld"})
		if d.inspectError(err, "container commit error") {
			break
		}

		fmt.Println(commitResp.ID)

		d.State.SetOk()
	}

	return d.State
}

func (d *Docker) ContainerInspect(containerID string) (*types.ContainerJSON, *ux.State) {
	var ret types.ContainerJSON
	if state := d.IsNil(); state.IsError() {
		return &ret, state
	}

	for range onlyOnce {
		if d.Containers == nil {
			d.State.SetWarning("no containers found")
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		var err error
		ret, err = d.Client.ContainerInspect(ctx, containerID)
		if d.inspectError(err, "container inspect error") {
			break
		}

		d.State.SetOutput(ret)
		d.State.SetOk()
	}

	return &ret, d.State
}

func (d *Docker) ContainerCreate(config *container.Config, hostConfig *container.HostConfig, netConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, *ux.State) {
	var resp container.ContainerCreateCreatedBody
	if state := d.IsNil(); state.IsError() {
		return resp, state
	}

	for range onlyOnce {
		//if d.Containers == nil {
		//	d.State.SetWarning("no containers found")
		//	break
		//}

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		//var resp container.ContainerCreateCreatedBody
		var err error
		resp, err = d.Client.ContainerCreate(ctx,
			config,
			hostConfig,
			netConfig,
			nil, // TODO verify nil is ok
			containerName,
		)
		if d.inspectError(err, "error creating container") {
			break
		}

		if resp.ID == "" {
			d.State.SetError("error creating container")
			break
		}

		d.State.SetOk()
	}

	return resp, d.State
}

// ******************************************************************************** //

func (d *Docker) NetworkList(name string) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//if d.Networks != nil {
		//	break
		//}

		var err error

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		df := filters.NewArgs()
		df.Add("name", name)

		d.Networks, err = d.Client.NetworkList(ctx, types.NetworkListOptions{Filters: df})
		if d.inspectError(err, "error listing networks") {
			break
		}

		d.State.SetResponse(len(d.Networks))

		if len(d.Networks) == 0 {
			d.State.SetWarning("%s network found", name)
			break
		}

		d.State.SetOk()
	}

	return d.State
}

func (d *Docker) NetworkCreate(name string, options types.NetworkCreate) *ux.State {
	if state := d.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		var err error

		ctx, cancel := context.WithTimeout(context.Background(), d.Provider.Timeout)
		//noinspection GoDeferInLoop
		defer cancel()

		resp, err := d.Client.NetworkCreate(ctx, name, options)
		if d.inspectError(err, "error creating network") {
			break
		}

		if resp.ID == "" {
			d.State.SetError("cannot create network")
			break
		}

		d.State.SetOk()
	}

	return d.State
}

func _close(c io.Closer) {
	_ = c.Close()

}
