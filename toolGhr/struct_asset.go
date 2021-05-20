package toolGhr

import (
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gearboxworks/scribeHelpers/toolGhr/github"
	"github.com/gearboxworks/scribeHelpers/toolPath"
	"github.com/gearboxworks/scribeHelpers/ux"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// FetchAssets gets the Release assets associated with a repo.
func (repo *TypeRepo) FetchAssets(force bool) *ux.State {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if force {
			repo.assets.all = nil
		}
		if repo.assets.all != nil {
			break
		}
		if repo.releases.selected == nil {
			repo.state.SetWarning("no such release")
			break
		}

		repo.state = repo.ClientGet(&repo.assets.all, releaseIdUri, repo.releases.selected.Id)
		if repo.state.IsNotOk() {
			break
		}

		if repo.assets.all == nil {
			repo.state.SetWarning("no assets found")
			break
		}

		repo.assets.latest = repo.assets.GetLatest()

		repo.Files = []string{}
		for _, file := range repo.assets.all {
			repo.Files = append(repo.Files, file.Name)
		}

		repo.state.SetOk()
		repo.state.SetResponse(&repo.assets)
		// Allows the use of the following in a calling function:
		//if repo.state.IsResponseNotOfType("assets") {
		//	repo.state.SetError("could not get assets")
		//	break
		//}
		//tags := repo.state.GetResponseData().(*assets)
	}

	return repo.state
}

func (repo *TypeRepo) Asset() *Asset {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return nil
	}
	repo.state.SetFunction()
	return repo.assets.GetSelected()
}

func (repo *TypeRepo) Assets() []*Asset {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return nil
	}
	repo.state.SetFunction()
	return repo.assets.GetAll()
}

func (repo *TypeRepo) CountAssets() int {
	if state := repo.IsNil(); state.IsError() {
		return 0
	}
	repo.state.SetFunction()
	return repo.releases.CountAll()
}

func (repo *TypeRepo) PrintAssets() *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()
	repo.assets.Print()
	return repo.state
}

func (repo *TypeRepo) PrintAsset() *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()
	if repo.assets.selected != nil {
		repo.assets.selected.Print()
	}
	return repo.state
}

func (repo *TypeRepo) SelectAsset(label string) *Asset {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return nil
	}
	repo.state.SetFunction()
	label = filepath.Base(label)
	return repo.assets.findAsset(label)
}

func (repo *TypeRepo) SelectRegexpAsset(label string) *Asset {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return nil
	}
	repo.state.SetFunction()
	var asset *Asset

	for range onlyOnce {
		re := regexp.MustCompile(label)
		if re == nil {
			repo.state.SetError("Invalid regular expression.")
			break
		}

		asset = repo.assets.regexpAsset(re)
	}

	return asset
}

// Delete sends a HTTP DELETE request for the given asset to Github. Returns
// nil if the asset was deleted OR there was nothing to delete.
func (repo *TypeRepo) DeleteAsset(label string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if repo.assets.all == nil {
			repo.state.SetError("No assets available")
			break
		}

		ref := repo.SelectAsset(label)
		if ref == nil {
			repo.state.SetError("Release asset '%s' not available", label)
			break
		}

		repo.state = repo.DeleteAssetRef(ref)
	}

	return repo.state
}

func (repo *TypeRepo) DeleteAssetRef(ref *Asset) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if ref == nil {
			repo.messageError("Deleting Release asset FAILED - empty")
			repo.state.SetError("Deleting Release asset FAILED - empty")
			break
		}

		repo.message("Deleting Release asset '%s' ...", ref.Name)
		URL := repo.generateApiUrl(AssetUri, ref.Id)
		resp, err := github.DoAuthRequest("DELETE", URL, "application/json", repo.Auth.Token, nil, nil)
		if err != nil {
			repo.messageError("Deleting Release asset '%s' FAILED", ref.Name)
			repo.state.SetError("failed to delete asset %s (ID: %d), HTTP error: %b", ref.Name, ref.Id, err)
			break
		}
		//noinspection ALL
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			repo.messageError("Deleting Release asset '%s' FAILED", ref.Name)
			repo.state.SetError("failed to delete asset %s (ID: %d), status: %s", ref.Name, ref.Id, resp.Status)
			break
		}

		repo.messageOk("Deleted Release asset '%s' OK", ref.Name)
		repo.state.SetOk()
	}

	return repo.state
}

func (repo *TypeRepo) UploadAsset(overwrite bool, label string, path ...string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if repo.releases.selected == nil {
			repo.state.SetError("no release selected")
			break
		}
		rel := repo.releases.selected

		file := NewFile(repo.runtime)
		if file.state.IsNotOk() {
			repo.state = file.state
			break
		}
		repo.state = file.Set(overwrite, label, path...)
		if repo.state.IsNotOk() {
			break
		}

		for range onlyOnce {
			asset := repo.SelectAsset(file.Label)
			if asset == nil {
				break
			}

			if overwrite == false {
				break
			}

			repo.message("Asset (id: %d) exists in state %s: Removing ...", asset.Id, asset.State)
			repo.state = repo.DeleteAsset(asset.Name)
			if repo.state.IsNotOk() {
				repo.state.SetError("Could not remove asset '%s' prior to upload")
				break
			}
		}

		repo.state = file.OpenRead()
		if repo.state.IsNotOk() {
			repo.state.SetError("file '%s' does not exist", file.Name)
			break
		}
		//noinspection ALL
		defer file.Close()

		v := url.Values{}
		v.Set("name", strings.ToLower(file.Name)) // @TODO - selfupdate lowercase workaround.
		if file.Label != "" {
			v.Set("label", file.Label)
		}

		// Everything set - begin upload.
		repo.message("Uploading Release asset '%s' as label '%s' ...", file.Name, file.Label)
		u := rel.CleanUploadUrl() + "?" + v.Encode()
		resp, err := github.DoAuthRequest("POST", u, "application/octet-stream", repo.Auth.Token, nil, file.fh)
		if err != nil {
			repo.state.SetError("can't create upload request to %v, %v", u, err)
			break
		}
		//noinspection ALL
		defer resp.Body.Close()
		//repo.message("RESPONSE:", resp)

		var r io.Reader = resp.Body
		if repo.runtime.Debug {
			r = io.TeeReader(r, os.Stderr)
		}

		var asset *Asset
		// For HTTP status 201 and 502, Github will return a JSON encoding of
		// the (partially) created asset.
		if resp.StatusCode == http.StatusBadGateway || resp.StatusCode == http.StatusCreated {
			//repo.message("ASSET: ")
			asset = new(Asset)
			if err := json.NewDecoder(r).Decode(&asset); err != nil {
				repo.state.SetError("upload failed (%s), could not unmarshal asset (err: %v)", resp.Status, err)
				break
			}
		} else {
			repo.message("BODY: ")
			if msg, err := Tomessage(r); err == nil {
				repo.state.SetError("could not upload, status code (%s), %v", resp.Status, msg)
				break
			}
			repo.state.SetError("could not upload, status code (%s)", resp.Status)
			break
		}

		if resp.StatusCode == http.StatusBadGateway {
			// 502 means the upload failed, but GitHub still retains metadata
			// (an asset in state "new"). Attempt to delete that now since it
			// would clutter the list of Release assets.
			repo.message("Asset (id: %d) failed to upload and is now in state %s: Removing...", asset.Id, asset.Name)
			repo.state = repo.DeleteAssetRef(asset)
			if repo.state.IsNotOk() {
				repo.state.SetError("Upload failed (%s), could not delete partially uploaded asset (ID: %d, err: %v) in order to cleanly reset GH API state, please try again", resp.Status, asset.Id, err)
				break
			}
			repo.state.SetError("could not upload, status code (%s)", resp.Status)
			break
		}

		repo.messageOk("Uploaded Release asset '%s' OK", file.Label)
		repo.state.SetOk()
	}

	if repo.state.IsNotOk() {
		repo.messageError("Uploading Release asset '%s' FAILED ...", label)
	}
	return repo.state
}

func (repo *TypeRepo) DownloadAsset(overwrite bool, label string, path ...string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if repo.releases.selected == nil {
			repo.state.SetError("no release defined")
			break
		}

		file := NewFile(repo.runtime)
		if file.state.IsNotOk() {
			repo.state = file.state
			break
		}
		repo.state = file.Set(overwrite, label, path...)
		if repo.state.IsNotOk() {
			break
		}
		repo.state = file.Path.StatPath()
		if file.Path.Exists() && !overwrite {
			repo.state.SetOk("Not overwriting file '%s'.", file.Path.GetPathAbs())
			break
		}

		asset := repo.assets.findAsset(label)
		if asset == nil {
			repo.state.SetError("could not find asset named %s", file.Label)
			break
		}

		repo.state = file.OpenWrite()
		if repo.state.IsNotOk() {
			break
		}
		//noinspection ALL
		defer file.Close()

		// Everything set - begin download.
		repo.message("Downloading Release asset '%s' (bytes:%d) ...", asset.Name, asset.Size)
		var resp *http.Response
		var err error
		if repo.Auth.Token == "" {
			// Use the regular github.com site if we don't have a token.
			URL := repo.generateApiUrl(releaseAssetDownload, repo.TagName, asset.Name)
			resp, err = http.Get(URL)
		} else {
			URL := repo.generateApiUrl(AssetUri, asset.Id)
			resp, err = github.DoAuthRequest("GET", URL, "", repo.Auth.Token, map[string]string{"Accept": "application/octet-stream"}, nil)
		}
		//noinspection ALL
		defer resp.Body.Close()

		if err != nil {
			repo.state.SetError("could not fetch releases, %v", err)
			break
		}

		contentLength, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
		if err != nil {
			repo.state.SetError(err)
			break
		}

		if resp.StatusCode != http.StatusOK {
			repo.state.SetError("github did not respond with 200 OK but with %v", resp.Status)
			break
		}

		repo.message("Saving asset '%s' to file '%s' - %v", file.Label, file.Name, resp.Status)
		repo.state = file.Write(resp.Body, contentLength)
		if repo.state.IsNotOk() {
			break
		}

		repo.messageOk("Downloaded Release asset '%s' OK", label)
		repo.state.SetOk()
	}

	if repo.state.IsNotOk() {
		repo.messageError("Downloading Release asset '%s' FAILED ...", label)
	}
	return repo.state
}

func (repo *TypeRepo) DownloadAssets(overwrite bool, path ...string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if repo.releases.selected == nil {
			repo.state.SetError("no release defined")
			break
		}

		// Setup destination path.
		file := toolPath.New(repo.runtime)
		if file.State.IsNotOk() {
			repo.state = file.State
			break
		}
		file.SetPath(path...)
		repo.state = file.StatPath()
		if file.IsAFile() {
			repo.state.SetError("path '%s' is a file, cannot download assets", file.GetPathAbs())
			break
		}
		if file.NotExists() {
			repo.state = file.Mkdir()
			if repo.state.IsNotOk() {
				break
			}
			repo.state = file.StatPath()
		}
		if repo.state.IsNotOk() {
			break
		}

		savedFail := repo.state
		for _, asset := range repo.Assets() {
			repo.state = repo.DownloadAsset(overwrite, asset.Name, filepath.Join(filepath.Join(path...), asset.Name))
			if repo.state.IsNotOk() {
				savedFail = repo.state
			}
		}
		if savedFail.IsNotOk() {
			repo.state = savedFail
			break
		}

		repo.messageOk("Downloaded Release all assets OK")
		repo.state.SetOk()
	}

	if repo.state.IsNotOk() {
		repo.messageError("Downloading Release all asset FAILED")
	}
	return repo.state
}

type Asset struct {
	Url         string    `json:"url"`
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	ContentType string    `json:"content_type"`
	State       string    `json:"state"`
	Size        uint64    `json:"size"`
	Downloads   uint64    `json:"download_count"`
	Created     time.Time `json:"created_at"`
	Published   time.Time `json:"published_at"`
}

func (a *Asset) String() string {
	var ret string

	if a == nil {
		return ret
	}

	//str[idx] = ux.SprintfWhite("\t- artifact: %s, downloads: %d, state: %s, type: %s, size: %s, id: %d",
	ret = ux.SprintfWhite("\t- artifact: %s, downloads: %d, state: %s, size: %s",
		a.Name, a.Downloads, a.State, humanize.Bytes(a.Size))

	return ret
}

func (a *Asset) Print() {
	fmt.Print(a.String())
}

type assets struct {
	all      []*Asset
	selected *Asset
	latest   *Asset
}

func (a *assets) GetAll() []*Asset {
	return a.all
}

func (a *assets) GetSelected() *Asset {
	return a.selected
}

func (a *assets) GetLatest() *Asset {
	var rel *Asset
	for range onlyOnce {
		if a.latest != nil {
			rel = a.latest
			break
		}

		var latestRelIndex = -1
		maxDate := time.Time{}
		for i, asset := range a.all {
			rel = asset
			if relDate := asset.Published; relDate.After(maxDate) {
				maxDate = relDate
				latestRelIndex = i
			}
		}
		if latestRelIndex == -1 {
			break
		}

		rel = a.all[latestRelIndex]
	}
	return rel
}

func (a *assets) CountAll() int {
	return len(a.all)
}

func (a *assets) Sprint() string {
	var ret string
	switch {
	case a.all == nil:
		ret += ux.SprintfWarning("No assets found.")

	case a.selected == nil:
		// Print all assets.
		ret += ux.SprintfWarning("Found %d assets.", a.CountAll())
		for _, release := range a.all {
			ret += fmt.Sprintf("\n####\n%v", release)
		}

	default:
		// Print selected Release.
		ret += fmt.Sprintf("\n####\n%v", a.selected.Name)
	}
	return ret
}

func (a *assets) Print() {
	fmt.Print(a.Sprint())
}

func (a *assets) findAsset(label string) *Asset {
	for range onlyOnce {
		a.selected = nil

		if label == Latest {
			a.selected = a.GetLatest()
			break
		}

		for _, asset := range a.all {
			if asset.Name == label {
				a.selected = asset
			}
		}
	}

	return a.selected
}

func (a *assets) regexpAsset(label *regexp.Regexp) *Asset {
	for range onlyOnce {
		a.selected = nil

		for _, asset := range a.all {
			if label.MatchString(asset.Name) {
				a.selected = asset
			}
		}
	}

	return a.selected
}

func (a *assets) String() string {
	var ret string

	ret += ux.SprintfWhite("assets: (%d)\n", len(a.all))

	str := make([]string, len(a.all)+1)
	for idx, asset := range a.all {
		//str[idx] = ux.SprintfWhite("\t- artifact: %s, downloads: %d, state: %s, type: %s, size: %s, id: %d",
		str[idx] = fmt.Sprintf("%v", asset)
	}
	ret += strings.Join(str, "\n")

	return ret
}
