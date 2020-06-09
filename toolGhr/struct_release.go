package toolGhr

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
	"time"
)


const (
	ReleaseListUri    = "/repos/%s/%s/releases"
	ReleaseLatestUri  = "/repos/%s/%s/releases/latest"
	ReleaseDateFormat = "02/01/2006 at 15:04"
)


type Release struct {
	Url         string     `json:"url"`
	PageUrl     string     `json:"html_url"`
	UploadUrl   string     `json:"upload_url"`
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"body"`
	TagName     string     `json:"tag_name"`
	Draft       bool       `json:"draft"`
	Prerelease  bool       `json:"prerelease"`
	Created     *time.Time `json:"created_at"`
	Published   *time.Time `json:"published_at"`
	Assets      []Asset    `json:"assets"`
}
type Releases struct {
	All []*Release
	Selected *Release
}


// findRelease returns the release if a release with name can be found in releases,
// otherwise returns nil.
func (r *Releases) findRelease(name string) *Release {
	for range onlyOnce {
		r.Selected = nil

		if name == "latest" {
			// Latest will always be first... Maybe... @TODO - TO BE CHECKED
			r.Selected = r.All[0]
			break
		}

		for _, release := range r.All {
			if release.Name == name {
				r.Selected = release
			}
		}
	}

	return r.Selected
}


func (r *Release) CleanUploadUrl() string {
	bracket := strings.Index(r.UploadUrl, "{")

	if bracket == -1 {
		return r.UploadUrl
	}

	return r.UploadUrl[0:bracket]
}


func (r *Release) String() string {
	var ret string

	ret = ux.SprintfWhite("Name: %s\n", r.Name)
	ret += ux.SprintfWhite("Published: %s\n", r.Published.Format("2006-01-02T15:04:05-0700"))
	ret += ux.SprintfWhite("Tag: %s\n", r.TagName)
	ret += ux.SprintfWhite("Url: %s\n", r.PageUrl)
	ret += ux.SprintfWhite("Draft: %s\tPre-release: %s\n", Mark(r.Draft), Mark(r.Prerelease))
	ret += ux.SprintfWhite("Assets: (%d)\n", len(r.Assets))

	str := make([]string, len(r.Assets)+1)
	for idx, asset := range r.Assets {
		//str[idx] = ux.SprintfWhite("\t- artifact: %s, downloads: %d, state: %s, type: %s, size: %s, id: %d",
		str[idx] = ux.SprintfWhite("\t- artifact: %s, downloads: %d, state: %s, size: %s",
			asset.Name, asset.Downloads, asset.State, humanize.Bytes(asset.Size))
	}
	ret += strings.Join(str, "\n")

	return ret
}


func (r *Release) Print() string {
	str := make([]string, len(r.Assets)+1)
	str[0] = fmt.Sprintf(
		"%s, name: '%s', description: '%s', id: %d, tagged: %s, published: %s, draft: %v, prerelease: %v",
		r.TagName, r.Name, r.Description, r.Id,
		timeFmtOr(r.Created, ReleaseDateFormat, ""),
		timeFmtOr(r.Published, ReleaseDateFormat, ""),
		Mark(r.Draft), Mark(r.Prerelease))

	for idx, asset := range r.Assets {
		str[idx+1] = fmt.Sprintf("  - artifact: %s, downloads: %d, state: %s, type: %s, size: %s, id: %d",
			asset.Name, asset.Downloads, asset.State, asset.ContentType,
			humanize.Bytes(asset.Size), asset.Id)
	}

	return strings.Join(str, "\n")
}


type ReleaseCreate struct {
	TagName         string `json:"tag_name"`
	TargetCommitish string `json:"target_commitish,omitempty"`
	Name            string `json:"name"`
	Body            string `json:"body"`
	Draft           bool   `json:"draft"`
	Prerelease      bool   `json:"prerelease"`
}


//func (ghr *TypeGhr) GetReleases() (Releases, *ux.State) {
//	var ret Releases
//	if state := ux.IfNilReturnError(ghr); state.IsError() {
//		return ret, state
//	}
//	ghr.State.SetFunction()
//
//	for range onlyOnce {
//		//c := github.NewClient(ghr.Auth.AuthUser, ghr.Auth.Token, nil)
//		//c.SetBaseURL(ghr.urlPrefix)
//		ghr.State = ghr.Repo.apiGet(ReleaseListUri, ghr.Repo.Organization, ghr.Repo.Name)
//		if ghr.State.IsNotOk() {
//			break
//		}
//
//		if ghr.State.IsResponseNotOfType("Releases") {
//			ghr.State.SetError("could not get releases")
//			break
//		}
//
//		ret = ghr.State.GetResponseData().(Releases)
//	}
//
//	return ret, nil
//}


func (ghr *TypeGhr) latestReleaseApi() (*Release, *ux.State) {
	var ret Release
	if state := ux.IfNilReturnError(ghr); state.IsError() {
		return &ret, state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		//c := github.NewClient(ghr.Auth.AuthUser, ghr.Auth.Token, nil)
		//c.SetBaseURL(ghr.urlPrefix)
		err := ghr.Repo.apiGet(ReleaseLatestUri, ghr.Repo.Organization, ghr.Repo.Name)
		if err != nil {
			ghr.State.SetError(err)
			break
		}

		if ghr.State.IsResponseNotOfType("Release") {
			ghr.State.SetError("could not get latest release")
			break
		}

		ret = ghr.State.GetResponseData().(Release)
	}

	return &ret, ghr.State
}


func (ghr *TypeGhr) LatestRelease() (*Release, *ux.State) {
	var ret *Release
	if state := ux.IfNilReturnError(ghr); state.IsError() {
		return ret, state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		// If latestReleaseApi DOESN'T give an error, return the release.
		ret, ghr.State = ghr.latestReleaseApi()
		if ghr.State.IsNotOk() {
			break
		}

		// The enterprise urlPrefix doesnt support the latest release endpoint. Get
		// all releases and compare the published date to get the latest.
		ghr.State = ghr.Repo.GetReleases()
		if ghr.State.IsNotOk() {
			break
		}

		var latestRelIndex = -1
		maxDate := time.Time{}
		for i, release := range *ghr.Repo.releases {
			if relDate := *release.Published; relDate.After(maxDate) {
				maxDate = relDate
				latestRelIndex = i
			}
		}
		if latestRelIndex == -1 {
			ghr.State.SetError("could not find the latest release")
			break
		}

		ret = (*ghr.Repo.releases)[latestRelIndex]
		ux.PrintflnBlue("Found %d releases, latest release is %s", len(*ghr.Repo.releases), ret.Name)
	}

	return ret, ghr.State
}


func (ghr *TypeGhr) GetRelease() *ux.State {
	if state := ux.IfNilReturnError(ghr); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.Repo.GetReleases()
		if ghr.State.IsNotOk() {
			ghr.State.SetError("could not find the rel corresponding to tag %s", ghr.Repo.Tag)
			break
		}

		if len(*ghr.Repo.releases) == 0 {
			ghr.State.SetError("could not find any releases")
		}

		if ghr.Repo.Tag == "latest" {
			ghr.Repo.Tag = (*ghr.Repo.releases)[0].Name
		}

		for _, rel := range *ghr.Repo.releases {
			if rel.TagName == ghr.Repo.Tag {
				ghr.Repo.release = rel
				break
			}
		}
	}

	return ghr.State
}


/* find the release-id of the specified tag */
func (ghr *TypeGhr) idOfTag() (int, *ux.State) {
	var ret int
	if state := ux.IfNilReturnError(ghr); state.IsError() {
		return ret, state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		var rel *Release

		rel, ghr.State = ghr.releaseOfTag()
		if ghr.State.IsNotOk() {
			break
		}

		ret = rel.Id
	}

	return ret, nil
}
