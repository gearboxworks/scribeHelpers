package toolGhr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/newclarity/scribeHelpers/toolGhr/github"
	"github.com/newclarity/scribeHelpers/ux"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)


// Get the releases associated with a repo.
func (repo *TypeRepo) FetchReleases(force bool) *ux.State {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if force {
			repo.releases.all = nil
		}
		if repo.releases.all != nil {
			break
		}

		repo.state = repo.ClientGet(&repo.releases.all, releaseListUri)
		if repo.state.IsNotOk() {
			break
		}

		if repo.releases.all == nil {
			repo.state.SetWarning("no releases found")
			break
		}

		// Sometimes we can't second guess what the "latest" is based on date alone.
		repo.state = repo.ClientGet(&repo.releases.latest, releaseTagUri, Latest)
		if repo.state.IsNotOk() {
			repo.state.SetWarning("no latest release found")
			break
		}

		if repo.releases.findRelease(repo.TagName) == nil {
			repo.state.SetWarning("no Release '%s' found", repo.TagName)
			break
		}

		repo.state.SetOk()
		repo.state.SetResponse(&repo.releases)
		// Allows the use of the following in a calling function:
		//if repo.state.IsResponseNotOfType("releases") {
		//	repo.state.SetError("could not get releases")
		//	break
		//}
		//releases := repo.state.GetResponseData().(*releases)
	}

	return repo.state
}

func (repo *TypeRepo) Latest() *Release {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return nil
	}
	repo.state.SetFunction()
	return repo.releases.GetLatest()
}

func (repo *TypeRepo) Release() *Release {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return nil
	}
	repo.state.SetFunction()
	return repo.releases.GetSelected()
}

func (repo *TypeRepo) Releases() *releases {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return nil
	}
	repo.state.SetFunction()
	return repo.releases
}

func (repo *TypeRepo) CountReleases() int {
	if state := repo.IsNil(); state.IsError() {
		return 0
	}
	repo.state.SetFunction()
	return repo.releases.CountAll()
}

func (repo *TypeRepo) PrintReleases() *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()
	repo.releases.Print()
	return repo.state
}

func (repo *TypeRepo) PrintRelease() *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()
	if repo.releases.selected != nil {
		repo.releases.selected.Print()
	}
	return repo.state
}

func (repo *TypeRepo) SelectRelease(tag string) *Release {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return nil
	}
	repo.state.SetFunction()
	rel := repo.releases.findRelease(tag)
	if rel != nil {
		repo.TagName = rel.Name
	}
	return rel
}

// Delete sends a HTTP DELETE request for the given asset to Github. Returns
// nil if the asset was deleted OR there was nothing to delete.
func (repo *TypeRepo) DeleteRelease(tag string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if repo.releases.all == nil {
			repo.state.SetError("No releases available")
			break
		}

		ref := repo.SelectRelease(tag)
		if ref == nil {
			repo.state.SetError("Release '%s' not available", tag)
			break
		}

		repo.state = repo.DeleteReleaseRef(ref)
	}

	return repo.state
}

func (repo *TypeRepo) DeleteReleaseRef(ref *Release) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if ref == nil {
			repo.messageError("Deleting Release FAILED - empty")
			repo.state.SetError("Deleting Release FAILED - empty")
			break
		}

		repo.message("Deleting Release '%s' ...", ref.TagName)
		URL := repo.generateApiUrl(releaseIdUri, ref.Id)
		resp, err := github.DoAuthRequest("DELETE", URL, "application/json", repo.Auth.Token, nil, nil)
		if err != nil {
			repo.messageError("Deleting Release '%s' FAILED", ref.TagName)
			repo.state.SetError("failed to delete release %s (ID: %d), HTTP error: %b", ref.TagName, ref.Id, err)
			break
		}
		//noinspection ALL
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			repo.messageError("Deleting Release asset '%s' FAILED", ref.TagName)
			repo.state.SetError("failed to delete release %s (ID: %d) - status: %s", ref.TagName, ref.Id, resp.Status)
			break
		}
		repo.messageOk("Deleted Release '%s' OK", ref.TagName)

		repo.state = repo.DeleteTag(ref.TagName)
		if repo.state.IsNotOk() {
			break
		}

		repo.state.SetOk()
	}

	return repo.state
}

func (repo *TypeRepo) CreateRelease(ref *TypeRepo) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if ref != nil {
			repo.state = repo.Set(ref)
			if repo.state.IsNotOk() {
				break
			}
		}

		if repo.Overwrite {
			repo.state = repo.deleteIfReleaseExist(repo.TagName)
			if repo.state.IsNotOk() {
				break
			}
		}


		repo.message("Creating release '%s' ...", repo.TagName)
		// Check if we need to read the description from stdin.
		if repo.Description == "-" {
			b, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				repo.state.SetError("could not read description from stdin: %v", err)
				break
			}
			repo.Description = string(b)
		}

		params := releaseCreate{
			TagName:         repo.TagName,
			TargetCommitish: repo.Target,
			Name:            repo.TagName,
			Body:            repo.Description,
			Draft:           repo.Draft,
			Prerelease:      repo.Prerelease,
		}

		/* encode params as json */
		payload, err := json.Marshal(params)
		if err != nil {
			repo.state.SetError("can't encode Release creation params, %v", err)
			break
		}
		reader := bytes.NewReader(payload)

		URL := repo.generateApiUrl(releaseListUri)
		resp, err := github.DoAuthRequest("POST", URL, "application/json", repo.Auth.Token, nil, reader)
		if err != nil {
			repo.state.SetError("while submitting %v, %v", string(payload), err)
			break
		}
		//noinspection ALL
		defer resp.Body.Close()

		//ghr.message("RESPONSE:", resp)
		if resp.StatusCode != http.StatusCreated {
			if resp.StatusCode == 422 {
				//repo.state.SetError("github returned %v (this is probably because the Release already exists)", resp.Status)
				repo.state.SetError("Release '%s' already exists", repo.TagName)
				break
			}
			repo.state.SetError("github returned %v", resp.Status)
			break
		}

		if repo.runtime.Debug {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				repo.state.SetError("error while reading response, %v", err)
				break
			}
			repo.message("BODY:", string(body))
		}

		repo.messageOk("Created Release '%s' OK", repo.TagName)
		repo.state.SetOk()
	}

	if repo.state.IsNotOk() {
		repo.messageError("Creating Release FAILED - %s", repo.state.GetError())
	}
	return repo.state
}

func (repo *TypeRepo) UpdateRelease(rel *TypeRepo) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		r := repo.SelectRelease(rel.TagName)
		if r == nil {
			repo.state.SetError("Release '%s' does not exist.", rel.TagName)
			break
		}

		repo.state = repo.Set(rel)
		if repo.state.IsNotOk() {
			break
		}


		// Check if we need to read the description from stdin.
		if repo.Description == "-" {
			b, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				repo.state.SetError("could not read description from stdin: %v", err)
				break
			}
			repo.Description = string(b)
		}

		/* the Release create struct works for editing releases as well */
		params := releaseCreate{
			TagName:    repo.TagName,
			Name:       repo.TagName,
			Body:       repo.Description,
			Draft:      repo.Draft,
			Prerelease: repo.Prerelease,
		}

		/* encode the parameters as JSON, as required by the github API */
		payload, err := json.Marshal(params)
		if err != nil {
			repo.state.SetError("can't encode Release creation params, %v", err)
			break
		}


		repo.message("Updating Release '%s' ...", repo.TagName)
		URL := repo.generateApiUrl(releaseIdUri, r.Id)
		resp, err := github.DoAuthRequest("PATCH", URL, "application/json", repo.Auth.Token, nil, bytes.NewReader(payload))
		if err != nil {
			repo.state.SetError("while submitting %v, %v", string(payload), err)
			break
		}
		//noinspection ALL
		defer resp.Body.Close()

		//repo.message("RESPONSE:", resp)
		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == 422 {
				//repo.state.SetError("github returned %v (this is probably because the Release already exists)", resp.Status)
				repo.state.SetError("Release '%s' already exists", repo.TagName)
				break
			}
			repo.state.SetError("github returned unexpected status code %v", resp.Status)
			break
		}

		if repo.runtime.Debug {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				repo.state.SetError("error while reading response, %v", err)
				break
			}
			repo.message("BODY:", string(body))
		}

		repo.messageOk("Updating release '%s' OK", repo.TagName)
		repo.state.SetOk()
	}

	if repo.state.IsNotOk() {
		repo.messageError("Updating Release FAILED - %s", repo.state.GetError())
	}
	return repo.state
}

func (repo *TypeRepo) deleteIfReleaseExist(tag string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		rel := repo.SelectRelease(tag)
		if rel == nil {
			break
		}

		//if rel.Replace == false {
		//	break
		//}

		repo.message("Release '%s' exists. Removing ...", rel.TagName)
		repo.state = repo.DeleteRelease(rel.TagName)
		if repo.state.IsNotOk() {
			repo.state.SetError("Could not replace release '%s' prior to creating.", rel.TagName)
			break
		}
		repo.message("Release '%s' removed OK ...", rel.TagName)

		repo.state = repo.Fetch(true)
		if repo.state.IsError() {
			break
		}

		repo.state.SetOk()
	}

	return repo.state
}


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

func (r *Release) CleanUploadUrl() string {
	bracket := strings.Index(r.UploadUrl, "{")

	if bracket == -1 {
		return r.UploadUrl
	}

	return r.UploadUrl[0:bracket]
}

func (r *Release) String() string {
	var ret string

	if r == nil {
		return ret
	}

	ret = ux.SprintfWhite("Name: %s\n", r.Name)
	ret += ux.SprintfWhite("Published: %s\n", r.Published.Format("2006-01-02T15:04:05-0700"))
	ret += ux.SprintfWhite("Tag: %s\n", r.TagName)
	ret += ux.SprintfWhite("Url: %s\n", r.PageUrl)
	ret += ux.SprintfWhite("Draft: %s\tPre-Release: %s\n", Mark(r.Draft), Mark(r.Prerelease))
	ret += ux.SprintfWhite("assets: (%d)\n", len(r.Assets))

	str := make([]string, len(r.Assets)+1)
	for idx, asset := range r.Assets {
		//str[idx] = ux.SprintfWhite("\t- artifact: %s, downloads: %d, state: %s, type: %s, size: %s, id: %d",
		str[idx] = ux.SprintfWhite("\t- artifact: %s, downloads: %d, state: %s, size: %s",
			asset.Name, asset.Downloads, asset.State, humanize.Bytes(asset.Size))
	}
	ret += strings.Join(str, "\n")

	return ret
}

func (r *Release) Print() {
	fmt.Print(r.String())
}


type releases struct {
	all      []*Release
	selected *Release
	latest   *Release
}

func (r *releases) GetAll() []*Release {
	return r.all
}

func (r *releases) GetSelected() *Release {
	return r.selected
}

func (r *releases) GetLatest() *Release {
	var rel *Release
	for range onlyOnce {
		if r.latest != nil {
			rel = r.latest
			break
		}

		var latestRelIndex = -1
		maxDate := time.Time{}
		for i, release := range r.all {
			rel = release
			if relDate := *release.Published; relDate.After(maxDate) {
				maxDate = relDate
				latestRelIndex = i
			}
		}
		if latestRelIndex == -1 {
			break
		}

		rel = r.all[latestRelIndex]
	}
	return rel
}

func (r *releases) CountAll() int {
	return len(r.all)
}

func (r *releases) Sprint() string {
	var ret string
	switch {
		case r.all == nil:
			ret += ux.SprintfWarning("No releases found.")

		case r.selected == nil:
			// Print all releases.
			ret += ux.SprintfWarning("Found %d releases.", r.CountAll())
			for _, release := range r.all {
				ret += fmt.Sprintf("\n####\n%v", release)
			}

		default:
			// Print selected Release.
			ret += fmt.Sprintf("\n####\n%v", r.selected)
	}
	return ret
}

func (r *releases) Print() {
	fmt.Print(r.Sprint())
}

func (r *releases) findRelease(name string) *Release {
	for range onlyOnce {
		r.selected = nil

		if name == Latest {
			r.selected = r.GetLatest()
			break
		}

		for _, release := range r.all {
			if release.Name == name {
				r.selected = release
			}
			if release.TagName == name {
				r.selected = release
			}
		}
	}

	return r.selected
}


type releaseCreate struct {
	TagName         string `json:"tag_name"`
	TargetCommitish string `json:"target_commitish,omitempty"`
	Name            string `json:"name"`
	Body            string `json:"body"`
	Draft           bool   `json:"draft"`
	Prerelease      bool   `json:"prerelease"`
}
