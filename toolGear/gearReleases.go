package toolGear

import (
	"fmt"
	"github.com/cavaliercoder/grab"
	"github.com/gearboxworks/scribeHelpers/toolPath"
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
	"github.com/google/go-github/github"
	"golang.org/x/net/context"
	"os"
	"strings"
	"time"
)


//goland:noinspection SpellCheckingInspection
const (
	Brandname = "Gearbox"
)

type GitHubRepo struct {
	Map      ReleasesMap
	Latest   *Release
	Selected *Release
	BaseDir  *toolPath.TypeOsPath

	runtime  *toolRuntime.TypeRuntime
	State    *ux.State
}
type ReleasesMap map[Version]*Release
type Version string

type Release struct {
	Version       Version
	File          *toolPath.TypeOsPath
	Size          int64
	Url           string
	Instance      *github.RepositoryRelease
	DlIndex       int
	IsDownloading bool

	//runtime       *toolRuntime.TypeRuntime
	State         *ux.State
}

type ReleaseSelector struct {
	// These are considered to be AND-ed together.
	FromDate        time.Time
	UntilDate       time.Time
	SpecificVersion string
	RegexpVersion   string
	Latest			*bool

	//runtime         *toolRuntime.TypeRuntime
	State           *ux.State
}


func NewRepo(runtime *toolRuntime.TypeRuntime) *GitHubRepo {
	runtime = runtime.EnsureNotNil()

	repo := GitHubRepo{
		Map:      make(ReleasesMap),
		Latest:   nil,
		Selected: nil,
		BaseDir:  toolPath.ToolNewPath(runtime.CmdDir, "iso"),

		runtime:  runtime,
		State:    ux.NewState(runtime.CmdName, runtime.Debug),
	}
	//repo.State = ret.UpdateReleases()

	return &repo
}

//goland:noinspection SpellCheckingInspection
func (ghr *GitHubRepo) ShowReleases() *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		ux.Printf("Latest release: %v\n\n", ghr.Latest)
		for _, release := range ghr.Map {
			ux.Printf("Assets for release:	%v\n", release.Instance.GetName())
			ux.Printf("UploadURL: 			%v\n", release.Instance.GetUploadURL())
			ux.Printf("ZipballURL: 			%v\n", release.Instance.GetZipballURL())
			ux.Printf("TarballURL: 			%v\n", release.Instance.GetTarballURL())
			ux.Printf("Body: 				%v\n", release.Instance.GetBody())
			ux.Printf("AssetsURL: 			%v\n", release.Instance.GetAssetsURL())
			ux.Printf("URL: 				%v\n", release.Instance.GetURL())
			ux.Printf("HTMLURL:				%v\n", release.Instance.GetHTMLURL())

			for _, asset := range release.Instance.Assets {
				ux.Printf("	Name:				%v\n", asset.GetName())
				ux.Printf("	ID:					%v\n", asset.GetID())
				ux.Printf("	URL:					%v\n", asset.GetURL())
				ux.Printf("	Size:				%v\n", asset.GetSize())
				ux.Printf("	CreatedAt:			%v\n", asset.GetCreatedAt())
				ux.Printf("	UpdatedAt:			%v\n", asset.GetUpdatedAt())
				ux.Printf("	BrowserDownloadURL:	%v\n", asset.GetBrowserDownloadURL())
				ux.Printf("	State:				%v\n", asset.GetState())
				ux.Printf("	ContentType:			%v\n", asset.GetContentType())
				ux.Printf("	DownloadCount:		%v\n", asset.GetDownloadCount())
				ux.Printf("	NodeID:				%v\n", asset.GetNodeID())
			}
		}

		//eblog.Debug(entity.VmBoxEntityName, "Showing all ISO releases. Latest == %s", ghr.Latest)
	}

	return ghr.State
}


func (r *Release) ShowRelease() *ux.State {
	if state := r.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if r.Instance.Name == nil {
			r.State.SetError("no release version specified")
			break
		}

		ux.Printf("Assets for release:	%v\n", *r.Instance.Name)
		for _, asset := range r.Instance.Assets {
			ux.Printf("	Name:				%v\n", asset.GetName())
			ux.Printf("	ID:					%v\n", asset.GetID())
			ux.Printf("	URL:					%v\n", asset.GetURL())
			ux.Printf("	Size:				%v\n", asset.GetSize())
			ux.Printf("	CreatedAt:			%v\n", asset.GetCreatedAt())
			ux.Printf("	UpdatedAt:			%v\n", asset.GetUpdatedAt())
			ux.Printf("	BrowserDownloadURL:	%v\n", asset.GetBrowserDownloadURL())
			ux.Printf("	State:				%v\n", asset.GetState())
			ux.Printf("	ContentType:			%v\n", asset.GetContentType())
			ux.Printf("	DownloadCount:		%v\n", asset.GetDownloadCount())
			ux.Printf("	NodeID:				%v\n", asset.GetNodeID())
		}
	}

	return r.State
}


func (ghr *GitHubRepo) UpdateReleases() *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//goland:noinspection SpellCheckingInspection
		if ghr.BaseDir == nil {
			ghr.BaseDir = toolPath.ToolNewPath("")
			_ = ghr.BaseDir.AppendPath("iso")
			//p := ospaths.New("")
			//ghr.BaseDir = p.UserConfigDir.AddToPath("iso")
		}

		var rm = make(ReleasesMap)
		ghr.Map = rm

		client := github.NewClient(nil)
		//ctx := context.Background()
		opt := &github.ListOptions{}

		//goland:noinspection SpellCheckingInspection
		releases, _, err := client.Repositories.ListReleases(context.Background(), "gearboxworks", "docker-os", opt)
		if err != nil {
			ghr.State.SetError("can't fetch GitHub releases")
			break
		}

		findFirst := true
		for _, rel := range releases {
			if rel == nil {
				continue
			}

			name := Version(rel.GetName())

			release := Release{
				Version: name,
				Url: "",
				Instance: rel,
			}

			// rm[name].Url/TypeFile - Find the first ISO asset.
			for _, asset := range rel.Assets {
				if strings.HasSuffix(asset.GetBrowserDownloadURL(), ".iso") {
					// Return the first ISO found.
					release.Url = asset.GetBrowserDownloadURL()
					//release.TypeFile = ghr.BaseDir.AddFileToPath(asset.GetName())
					_ = release.File.SetPath(ghr.BaseDir.GetPath(), asset.GetName())
					release.Size = int64(asset.GetSize())
					break
				}
			}

			// rm[name].Version - Copy version name.
			rm[name] = &release

			// rm.Latest - Find first version and select as 'latest'.
			if findFirst {
				ghr.Latest = &release
				findFirst = false
			}
		}

		//if findFirst == true {
		//	// If we never found a "first", then there was no data.
		//	// So don't update the map.
		//}

		ghr.Map = rm
	}

	return ghr.State
}


/*
Updates the following:
   me.VmIsoVersion    string
   me.VmIsoFile       string
   me.VmIsoUrl 		string
   me.VmIsoRelease    Release
*/
//noinspection GoUnusedParameter
func (ghr *GitHubRepo) SelectRelease(selector ReleaseSelector) *Release {
	var r *Release
	if state := ghr.IsNil(); state.IsError() {
		return &Release{ State: state }
	}

	for range onlyOnce {
		//err = ghr.UpdateReleases()
		//if err != nil {
		//	break
		//}

		// For now just select the latest.
		ghr.Selected = ghr.Latest
		r = ghr.Selected
	}

	return r
}


func (r *Release) GetIso() *ux.State {
	if state := r.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if r.File.GetPath() == "" {
			r.State.SetError(fmt.Sprintf("no %s OS iso file defined VmIsoUrl:%s VmIsoFile:%s", Brandname, r.Url, r.File.GetPath()))
			break
		}

		if r.Url == "" {
			r.State.SetError(fmt.Sprintf("no %s OS iso url defined VmIsoUrl:%s VmIsoFile:%s", Brandname, r.Url, r.File.GetPath()))
			break
		}


		var numb int
		numb, r.State = r.IsIsoFilePresent()
		if numb != IsoFileNeedsToDownload {
			break
		}


		// Start download
		r.DlIndex = 0
		r.IsDownloading = true
		client := grab.NewClient()
		req, _ := grab.NewRequest(r.File.GetPath(), r.Url)
		ux.Printf("downloading ISO from URL %s", req.URL().String())
		resp := client.Do(req)
		// fmt.Printf("  %v\n", resp.HTTPResponse.Status)
		ux.Printf("%s VM: Downloading ISO from '%s' to '%s'. Size:%d\n",
			Brandname,
			r.Url,
			r.File.GetPath(),
			resp.Size)


		// start UI loop
		t := time.NewTicker(500 * time.Millisecond)
		//noinspection GoDeferInLoop
		defer t.Stop()

		Loop:
			for {
				select {
					case <-t.C:
						r.DlIndex = int(100*resp.Progress())
						//r.publishDownloadState()
						//fmt.Printf("Downloading '%s' transferred %v / %v bytes (%d%%)\n", r.TypeFile.GetPath(), resp.BytesComplete(), resp.Size, r.DlIndex)
						fmt.Printf("%s VM: Downloading ISO - %d%% complete.\r",
							Brandname,
							r.DlIndex)

					case <-resp.Done:
						// download is complete
						break Loop
				}
			}

		// check for errors
		if err := resp.Err(); err != nil {
			ux.PrintfError("\nDownload failed\n")
			r.State.SetError(fmt.Sprintf("ISO download failed VmIsoUrl:%s VmIsoFile:%s", r.Url, r.File.GetPath()))
			break
		}
		ux.PrintfOk("%s VM: Downloaded ISO completed OK.\n",
			Brandname,
		)


		r.DlIndex = 100
		//r.publishDownloadState()
		r.IsDownloading = false
	}

	return r.State
}


//func (me *Release) publishDownloadState() {
//
//	client := messages.MessageAddress(entity.VmUpdateEntityName)
//	state := states.New(&client, &client, entity.VmBoxEntityName)
//	r.State.SetWant("100%")
//	r.State.SetCurrent(states.State(fmt.Sprintf("%d%%", me.DlIndex)))
//
//	f := messages.MessageAddress(states.ActionUpdate)
//	msg := f.ConstructMessage(entity.BroadcastEntityName, states.ActionStatus, r.State.ToMessageText())
//	_ = me.channels.Publish(msg)
//}


const IsoFileNeedsToDownload	= 0
const IsoFileIsDownloading		= 1
const IsoFileDownloaded			= 2
func (r *Release) IsIsoFilePresent() (int, *ux.State) {
	var ret int
	if state := r.IsNil(); state.IsError() {
		return 0, state
	}

	for range onlyOnce {
		if r.File.GetPath() == "" {
			r.State.SetError( fmt.Sprintf("no %s OS iso file defined VmIsoUrl:%s VmIsoFile:%s", Brandname, r.Url, r.File.GetPath()))
			break
		}

		stat, err := os.Stat(r.File.GetPath())
		if os.IsNotExist(err) {
			r.State.SetError("ISO file needs to download from GitHub VmIsoUrl:%s VmIsoFile:%s", r.Url, r.File.GetPath())
			ret = IsoFileNeedsToDownload
			break
		}

		if r.IsDownloading {
			r.State.SetError("ISO file still downloading VmIsoUrl:%s VmIsoFile:%s Percent:%d", r.Url, r.File.GetPath(), r.DlIndex)
			ret = IsoFileIsDownloading
			break
		}

		if stat.Size() != r.Size {
			r.State.SetError("ISO file needs to re-download from GitHub VmIsoUrl:%s VmIsoFile:%s", r.Url, r.File.GetPath())
			ret = IsoFileNeedsToDownload
			break
		}

		//if r.DlIndex < 100 {
		//	err = errors.New("ISO file needs to re-download from GitHub VmIsoUrl:%s VmIsoFile:%s", r.Url, r.TypeFile.GetPath())
		//	ret = IsoFileNeedsToDownload
		//	break
		//}

		ret = IsoFileDownloaded
		r.DlIndex = 100
	}

	return ret, r.State
}


func (ghr *GitHubRepo) IsNil() *ux.State {
	if state := ux.IfNilReturnError(ghr); state.IsError() {
		return state
	}
	ghr.State = ghr.State.EnsureNotNil()
	return ghr.State
}

func (ghr *GitHubRepo) IsValid() *ux.State {
	if state := ux.IfNilReturnError(ghr); state.IsError() {
		return state
	}

	for range onlyOnce {
		ghr.State = ghr.State.EnsureNotNil()
	}

	return ghr.State
}


func (r *Release) IsNil() *ux.State {
	if state := ux.IfNilReturnError(r); state.IsError() {
		return state
	}
	r.State = r.State.EnsureNotNil()
	return r.State
}


func (r *Release) IsValid() *ux.State {
	if state := ux.IfNilReturnError(r); state.IsError() {
		return state
	}

	for range onlyOnce {
		r.State = r.State.EnsureNotNil()
	}

	return r.State
}


//type ReleaseAsset struct {
//	ID                 *int64     `json:"id,omitempty"`
//	URL                *string    `json:"url,omitempty"`
//	Name               *string    `json:"name,omitempty"`
//	Label              *string    `json:"label,omitempty"`
//	State              *string    `json:"state,omitempty"`
//	ContentType        *string    `json:"content_type,omitempty"`
//	Size               *int       `json:"size,omitempty"`
//	DownloadCount      *int       `json:"download_count,omitempty"`
//	CreatedAt          *Timestamp `json:"created_at,omitempty"`
//	UpdatedAt          *Timestamp `json:"updated_at,omitempty"`
//	BrowserDownloadURL *string    `json:"browser_download_url,omitempty"`
//	Uploader           *User      `json:"uploader,omitempty"`
//	NodeID             *string    `json:"node_id,omitempty"`
//}
//
//type RepositoryRelease struct {
//	ID              *int64         `json:"id,omitempty"`
//	TagName         *string        `json:"tag_name,omitempty"`
//	TargetCommitish *string        `json:"target_commitish,omitempty"`
//	Name            *string        `json:"name,omitempty"`
//	Body            *string        `json:"body,omitempty"`
//	Draft           *bool          `json:"draft,omitempty"`
//	Prerelease      *bool          `json:"prerelease,omitempty"`
//	CreatedAt       *Timestamp     `json:"created_at,omitempty"`
//	PublishedAt     *Timestamp     `json:"published_at,omitempty"`
//	URL             *string        `json:"url,omitempty"`
//	HTMLURL         *string        `json:"html_url,omitempty"`
//	AssetsURL       *string        `json:"assets_url,omitempty"`
//	Assets          []ReleaseAsset `json:"assets,omitempty"`
//	UploadURL       *string        `json:"upload_url,omitempty"`
//	ZipballURL      *string        `json:"zipball_url,omitempty"`
//	TarballURL      *string        `json:"tarball_url,omitempty"`
//	Author          *User          `json:"author,omitempty"`
//	NodeID          *string        `json:"node_id,omitempty"`
//}
//
//
//data returned:
//
//release.ID=0xc000289538
//release.TagName=0xc0002964c0
//release.TargetCommitish=0xc0002964d0
//release.Name=0xc0002964e0
//release.Body=0xc000296770
//release.Draft=0xc00028955b
//release.Prerelease=0xc00028957d
//release.CreatedAt=2019-05-23 02:34:10 +0000 UTC
//release.PublishedAt=2019-05-23 02:43:04 +0000 UTC
//release.URL=0xc000296470
//release.HTMLURL=0xc0002964a0
//release.AssetsURL=0xc000296480
//release.Assets=[github.ReleaseAsset{
//	ID:12825393,
//	URL:"https://api.github.com/repos/gearboxworks/gearbox-os/releases/assets/12825393",
//	Name:"gearbox-0.5.0.iso",
//	State:"uploaded",
//	ContentType:"application/octet-stream",
//	Size:67108864,
//	DownloadCount:0,
//	CreatedAt:github.Timestamp{2019-05-23 02:37:48 +0000 UTC},
//	UpdatedAt:github.Timestamp{2019-05-23 02:42:56 +0000 UTC},
//	BrowserDownloadURL:"https://github.com/gearboxworks/gearbox-os/releases/download/0.5.0/gearbox-0.5.0.iso",
//	Uploader:github.User{
//		Login:"MickMake",
//		ID:17118367,
//		NodeID:"MDQ6VXNlcjE3MTE4MzY3",
//		AvatarURL:"https://avatars0.githubusercontent.com/u/17118367?v=4",
//		HTMLURL:"https://github.com/MickMake",
//		GravatarID:"",
//		Type:"User",
//		SiteAdmin:false,
//		URL:"https://api.github.com/users/MickMake",
//		EventsURL:"https://api.github.com/users/MickMake/events{/privacy}",
//		FollowingURL:"https://api.github.com/users/MickMake/following{/other_user}",
//		FollowersURL:"https://api.github.com/users/MickMake/followers",
//		GistsURL:"https://api.github.com/users/MickMake/gists{/gist_id}",
//		OrganizationsURL:"https://api.github.com/users/MickMake/orgs",
//		ReceivedEventsURL:"https://api.github.com/users/MickMake/received_events",
//		ReposURL:"https://api.github.com/users/MickMake/repos",
//		StarredURL:"https://api.github.com/users/MickMake/starred{/owner}{/repo}",
//		SubscriptionsURL:"https://api.github.com/users/MickMake/subscriptions"
//		},
//	NodeID:"MDEyOlJlbGVhc2VBc3NldDEyODI1Mzkz"
//	}]
//release.UploadURL=0xc000296490
//release.ZipballURL=0xc000296760
//release.TarballURL=0xc000296750
//release.Author=github.User{Login:"MickMake", ID:17118367, NodeID:"MDQ6VXNlcjE3MTE4MzY3", AvatarURL:"https://avatars0.githubusercontent.com/u/17118367?v=4", HTMLURL:"https://github.com/MickMake", GravatarID:"", Type:"User", SiteAdmin:false, URL:"https://api.github.com/users/MickMake", EventsURL:"https://api.github.com/users/MickMake/events{/privacy}", FollowingURL:"https://api.github.com/users/MickMake/following{/other_user}", FollowersURL:"https://api.github.com/users/MickMake/followers", GistsURL:"https://api.github.com/users/MickMake/gists{/gist_id}", OrganizationsURL:"https://api.github.com/users/MickMake/orgs", ReceivedEventsURL:"https://api.github.com/users/MickMake/received_events", ReposURL:"https://api.github.com/users/MickMake/repos", StarredURL:"https://api.github.com/users/MickMake/starred{/owner}{/repo}", SubscriptionsURL:"https://api.github.com/users/MickMake/subscriptions"}
//release.NodeID=0xc0002964b0
//
//
//type Release struct {
//	Name string
//	UploadURL string
//	ZipballURL string
//	TarballURL string
//	Body string
//	AssetsURL string
//	URL string
//	HTMLURL string
//	Name string
//    Assets
//}
//type GetReleases []Release
//
//type Asset struct {
//      Name
//      ID
//      URL
//      Size
//      CreatedAt
//      UpdatedAt
//      BrowserDownloadURL
//      State
//      ContentType
//      DownloadCount
//      NodeID
//}
//type Assets []Asset
//
//
//
//
//   Assets for release:	0.5.0
//   UploadURL: 			https://uploads.github.com/repos/gearboxworks/gearbox-os/releases/17531887/assets{?name,label}
//   ZipballURL: 			https://api.github.com/repos/gearboxworks/gearbox-os/zipball/0.5.0
//   TarballURL: 			https://api.github.com/repos/gearboxworks/gearbox-os/tarball/0.5.0
//   Body: 				0.5.0 pre-release
//   AssetsURL: 			https://api.github.com/repos/gearboxworks/gearbox-os/releases/17531887/assets
//   URL: 				https://api.github.com/repos/gearboxworks/gearbox-os/releases/17531887
//   HTMLURL:				https://github.com/gearboxworks/gearbox-os/releases/tag/0.5.0
//   foo: 				0.5.0
//   Name:				gearbox-0.5.0.iso
//   ID:					12825393
//   URL:					https://api.github.com/repos/gearboxworks/gearbox-os/releases/assets/12825393
//   Size:				67108864
//   CreatedAt:			2019-05-23 02:37:48 +0000 UTC
//   UpdatedAt:			2019-05-23 02:42:56 +0000 UTC
//   BrowserDownloadURL:	https://github.com/gearboxworks/gearbox-os/releases/download/0.5.0/gearbox-0.5.0.iso
//   State:				uploaded
//   ContentType:			application/octet-stream
//   DownloadCount:		0
//   NodeID:				MDEyOlJlbGVhc2VBc3NldDEyODI1Mzkz
//
//	fmt.Printf(`
//		release.ID=%v
//		release.TagName=%v
//		release.TargetCommitish=%v
//		release.Name=%v
//		release.Body=%v
//		release.Draft=%v
//		release.Prerelease=%v
//		release.CreatedAt=%v
//		release.PublishedAt=%v
//		release.URL=%v
//		release.HTMLURL=%v
//		release.AssetsURL=%v
//		release.Assets=%v
//		release.UploadURL=%v
//		release.ZipballURL=%v
//		release.TarballURL=%v
//		release.Author=%v
//		release.NodeID=%v\n`,
//		release.ID,
//		release.TagName,
//		release.TargetCommitish,
//		release.Name,
//		release.Body,
//		release.Draft,
//		release.Prerelease,
//		release.CreatedAt,
//		release.PublishedAt,
//		release.URL,
//		release.HTMLURL,
//		release.AssetsURL,
//		release.Assets,
//		release.UploadURL,
//		release.ZipballURL,
//		release.TarballURL,
//		release.Author,
//		release.NodeID,
//		)

