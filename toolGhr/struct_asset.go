package toolGhr

import (
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/newclarity/scribeHelpers/toolGhr/github"
	"github.com/newclarity/scribeHelpers/ux"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)


// Get the Release assets associated with a repo.
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

		URL := repo.generateApiUrl(releaseIdUri, repo.releases.selected.Id)
		err := repo.client.Get(URL, &repo.assets.all)
		if err != nil {
			repo.state.SetError(err)
			break
		}

		if repo.assets.all == nil {
			repo.state.SetWarning("no assets found")
			break
		}

		// @TODO - figure out how to do this.
		// Sometimes we can't second guess what the "latest" is based on date alone.
		//u = fmt.Sprintf(assetsUri, repo.Organization, repo.Name, repo.releases.selected.Id)
		//err = repo.client.Get(u, &repo.tags.latest)
		//if err != nil {
		//	repo.state.SetError(err)
		//	break
		//}
		repo.assets.latest = repo.assets.GetLatest()

		if repo.file.Name != "" {
			if repo.assets.findAsset(repo.file.Name) == nil {
				repo.state.SetWarning("asset '%s' not found", repo.file.Name)
				break
			}
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
	repo.releases.Print()
	return repo.state
}

func (repo *TypeRepo) SelectAsset(file string) *Asset {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return nil
	}
	repo.state.SetFunction()
	return repo.assets.findAsset(file)
}

// Delete sends a HTTP DELETE request for the given asset to Github. Returns
// nil if the asset was deleted OR there was nothing to delete.
func (repo *TypeRepo) DeleteAsset(a *Asset) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		URL := repo.generateApiUrl(AssetUri, a.Id)
		resp, err := github.DoAuthRequest("DELETE", URL, "application/json", repo.Auth.Token, nil, nil)
		if err != nil {
			repo.state.SetError("failed to delete asset %s (ID: %d), HTTP error: %b", a.Name, a.Id, err)
			break
		}

		//noinspection ALL
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			repo.state.SetError("failed to delete asset %s (ID: %d), status: %s", a.Name, a.Id, resp.Status)
			break
		}
	}

	return repo.state
}

func (repo *TypeRepo) UploadAsset(file string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		repo.state = repo.SetFile(file ,"")
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.FileRead(file)
		if repo.state.IsNotOk() {
			repo.state.SetError("file '%s' does not exist", file)
			break
		}
		//noinspection ALL
		defer repo.FileClose()

		v := url.Values{}
		v.Set("name", repo.file.Name)
		if repo.file.Label != "" {
			v.Set("label", repo.file.Label)
		}

		if repo.releases.selected == nil {
			repo.state.SetError("no release defined")
			break
		}
		rel := repo.releases.selected

		repo.message("Uploading Release asset '%s' ...", file)
		u := rel.CleanUploadUrl() + "?" + v.Encode()
		resp, err := github.DoAuthRequest("POST", u, "application/octet-stream", repo.Auth.Token, nil, repo.file.Handle)
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
			repo.state = repo.DeleteAsset(asset)
			if repo.state.IsNotOk() {
				repo.state.SetError("Upload failed (%s), could not delete partially uploaded asset (ID: %d, err: %v) in order to cleanly reset GH API state, please try again", resp.Status, asset.Id, err)
				break
			}
			repo.state.SetError("could not upload, status code (%s)", resp.Status)
			break
		}

		repo.state.SetOk()
	}

	return repo.state
}

func (repo *TypeRepo) DownloadAsset(file string, overwrite bool) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		repo.state = repo.SetFile(file ,"")
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.FileOpenWrite(file, overwrite)
		if repo.state.IsNotOk() {
			break
		}
		//noinspection ALL
		defer repo.FileClose()

		asset := repo.assets.findAsset(file)
		if asset == nil {
			repo.state.SetError("could not find asset named %s", repo.file.Name)
			break
		}

		repo.message("Downloading Release asset ...")
		var resp *http.Response
		var err error
		if repo.Auth.Token == "" {
			// Use the regular github.com site if we don't have a token.
			URL := repo.generateApiUrl(releaseAssetDownload, repo.TagName, repo.file.Name)
			resp, err = http.Get(URL)
		} else {
			URL := repo.generateApiUrl(AssetUri, asset.Id)
			resp, err = github.DoAuthRequest("GET", URL, "", repo.Auth.Token, map[string]string{"Accept": "application/octet-stream",}, nil)
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

		repo.message("Saving file '%s' - %v", file, resp.Status)
		repo.state = repo.FileWrite(resp.Body, contentLength)
		if repo.state.IsNotOk() {
			break
		}

		repo.state.SetOk()
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

func (a *assets) findAsset(name string) *Asset {
	for range onlyOnce {
		a.selected = nil

		if name == Latest {
			a.selected = a.GetLatest()
			break
		}

		for _, asset := range a.all {
			if asset.Name == name {
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
