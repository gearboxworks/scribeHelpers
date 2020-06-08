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
type Releases []*Release

func (r *Release) CleanUploadUrl() string {
	bracket := strings.Index(r.UploadUrl, "{")

	if bracket == -1 {
		return r.UploadUrl
	}

	return r.UploadUrl[0:bracket]
}

func (r *Release) String() string {
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
//		ghr.State = ghr.Repo.ApiGet(ReleaseListUri, ghr.Repo.Organization, ghr.Repo.Name)
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
		err := ghr.Repo.ApiGet(ReleaseLatestUri, ghr.Repo.Organization, ghr.Repo.Name)
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
		var rels *Releases
		rels, ghr.State = ghr.Repo.GetReleases()
		if ghr.State.IsNotOk() {
			break
		}

		var latestRelIndex = -1
		maxDate := time.Time{}
		for i, release := range *rels {
			if relDate := *release.Published; relDate.After(maxDate) {
				maxDate = relDate
				latestRelIndex = i
			}
		}
		if latestRelIndex == -1 {
			ghr.State.SetError("could not find the latest release")
			break
		}

		ret = (*rels)[latestRelIndex]
		ux.PrintflnBlue("Found %d releases, latest release is %v", len(*rels), (*rels)[latestRelIndex])
	}

	return ret, ghr.State
}


func (ghr *TypeGhr) ReleaseOfTag() (*Release, *ux.State) {
	var ret *Release
	if state := ux.IfNilReturnError(ghr); state.IsError() {
		return &Release{}, state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		var rels *Releases
		rels, ghr.State = ghr.Repo.GetReleases()
		if ghr.State.IsNotOk() {
			ghr.State.SetError("could not find the rel corresponding to tag %s", ghr.Repo.Tag)
			break
		}

		for _, rel := range *rels {
			if rel.TagName == ghr.Repo.Tag {
				ret = rel
				break
			}
		}
	}

	return ret, ghr.State
}


/* find the release-id of the specified tag */
func (ghr *TypeGhr) IdOfTag() (int, *ux.State) {
	var ret int
	if state := ux.IfNilReturnError(ghr); state.IsError() {
		return ret, state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		var rel *Release

		rel, ghr.State = ghr.ReleaseOfTag()
		if ghr.State.IsNotOk() {
			break
		}

		ret = rel.Id
	}

	return ret, nil
}
