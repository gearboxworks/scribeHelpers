package toolGhr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/newclarity/scribeHelpers/toolGhr/github"
	"github.com/newclarity/scribeHelpers/ux"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)


func (ghr *TypeGhr) Info() *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}


		ghr.message("Getting release tag info ...")
		// Find regular git tags.
		var tags *Tags
		tags, ghr.State = ghr.Repo.GetTags()
		if ghr.State.IsNotOk() {
			ghr.State.SetError("could not fetch tags, %v", ghr.State.GetError())
			break
		}
		if len(*tags) == 0 {
			ghr.State.SetError("no tags available for %s", ghr.Repo.GetUrl())
			break
		}

		if ghr.Repo.Tag == "latest" {
			ghr.Repo.Tag = (*tags)[0].Name
		}

		if ghr.Repo.Tag == "" {
			// Return all tags.
			var t []string
			for _, tag := range *tags {
				t = append(t, tag.Name)
			}
			ghr.message("Tags: %s", strings.Join(t, ", "))

		} else {
			// If the user only requested one tag, filter out the rest.
			for _, t := range *tags {
				if t.Name == ghr.Repo.Tag {
					ghr.message("Tag: %s", ghr.Repo.Tag)
					break
				}
			}
		}


		ghr.message("Getting release info ...")
		// List releases + assets.
		if ghr.Repo.Tag == "" {
			var releases *Releases
			// Get all releases.
			releases, ghr.State = ghr.Repo.GetReleases()
			if ghr.State.IsNotOk() {
				break
			}
			ghr.message("Found %d releases.", len(*releases))
			for _, release := range *releases {
				fmt.Printf("\n####\n%v", release)
			}

		} else {
			var rel *Release
			// Get only one release.
			rel, ghr.State = ghr.releaseOfTag()
			if ghr.State.IsNotOk() {
				break
			}
			ghr.message("Found 1 release.")
			ghr.message("- %v", rel)
		}


		ghr.State.SetOk()
	}

	//return renderer(tags, releases)
	return ghr.State
}


func (ghr *TypeGhr) Upload() *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State = ghr.Repo.isValidTag()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.message("Uploading release asset ...")

		if ghr.File.Path == nil {
			ghr.State.SetError("provided file was not valid")
			break
		}
		//noinspection ALL
		defer ghr.File.Path.CloseFile()

		//ghr.State = ghr.validateCredentials()
		//if ghr.State.IsNotOk() {
		//	break
		//}

		// Find the release corresponding to the entered tag, if any.
		var rel *Release
		rel, ghr.State = ghr.releaseOfTag()
		if ghr.State.IsNotOk() {
			break
		}


		// If the user has attempted to upload this asset before, someone could
		// expect it to be present in the release struct (rel.Assets). However,
		// we have to separately ask for the specific assets of this release.
		// Reason: the assets in the Release struct do not contain incomplete
		// uploads (which regrettably happen often using the Github API). See
		// issue #26.
		var assets []Asset
		assets, ghr.State = ghr.Repo.GetAssets(rel.Id)
		if ghr.State.IsNotOk() {
			break
		}
		//if ghr.State.IsResponseNotOfType("[]Asset") {
		//	ghr.State.SetError("could not replace asset")
		//	break
		//}
		//assets = ghr.State.GetResponseData().([]Asset)


		// Incomplete (failed) uploads will have their state set to new. These
		// assets are (AFAIK) useless in all cases. The only thing they will do
		// is prevent the upload of another asset of the same name. To work
		// around this GH API weirdness, let's just delete assets if:
		//
		// 1. Their state is new.
		// 2. The user explicitly asked to delete/replace the asset with -R.
		if asset := findAsset(assets, ghr.File.Name); asset != nil && (asset.State == "new" || ghr.File.Replace) {
			ghr.message("Asset (id: %d) exists in state %s: Removing ...", asset.Id, asset.Name)
			ghr.State = ghr.DeleteAsset(asset)
			if ghr.State.IsNotOk() {
				ghr.State.SetError("could not replace asset: %v", ghr.State.GetError())
				break
			}
		}

		v := url.Values{}
		v.Set("name", ghr.File.Name)
		if ghr.File.Label != "" {
			v.Set("label", ghr.File.Label)
		}

		u := rel.CleanUploadUrl() + "?" + v.Encode()

		ghr.State = ghr.File.Path.OpenFileHandle()
		if ghr.State.IsNotOk() {
			break
		}

		var resp *http.Response
		var err error
		resp, err = github.DoAuthRequest("POST", u, "application/octet-stream", ghr.Auth.Token, nil, ghr.File.Path.FileHandle)
		if err != nil {
			ghr.State.SetError("can't create upload request to %v, %v", u, err)
			break
		}
		//noinspection ALL
		defer resp.Body.Close()
		ghr.message("RESPONSE:", resp)

		var r io.Reader = resp.Body
		if ghr.runtime.Debug {
			r = io.TeeReader(r, os.Stderr)
		}
		var asset *Asset
		// For HTTP status 201 and 502, Github will return a JSON encoding of
		// the (partially) created asset.
		if resp.StatusCode == http.StatusBadGateway || resp.StatusCode == http.StatusCreated {
			ghr.message("ASSET: ")
			asset = new(Asset)
			if err := json.NewDecoder(r).Decode(&asset); err != nil {
				ghr.State.SetError("upload failed (%s), could not unmarshal asset (err: %v)", resp.Status, err)
				break
			}
		} else {
			ghr.message("BODY: ")
			if msg, err := Tomessage(r); err == nil {
				ghr.State.SetError("could not upload, status code (%s), %v", resp.Status, msg)
				break
			}
			ghr.State.SetError("could not upload, status code (%s)", resp.Status)
			break
		}

		if resp.StatusCode == http.StatusBadGateway {
			// 502 means the upload failed, but GitHub still retains metadata
			// (an asset in state "new"). Attempt to delete that now since it
			// would clutter the list of release assets.
			ghr.message("Asset (id: %d) failed to upload and is now in state %s: Removing...", asset.Id, asset.Name)
			ghr.State = ghr.DeleteAsset(asset)
			if ghr.State.IsNotOk() {
				ghr.State.SetError("Upload failed (%s), could not delete partially uploaded asset (ID: %d, err: %v) in order to cleanly reset GH API state, please try again", resp.Status, asset.Id, err)
				break
			}
			ghr.State.SetError("could not upload, status code (%s)", resp.Status)
			break
		}

		ghr.State.SetOk()
	}

	// return nil
	return ghr.State
}


func (ghr *TypeGhr) Download(file string) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State = ghr.Repo.isValidTag()
		if ghr.State.IsNotOk() {
			break
		}

		if file == "" {
			ghr.State.SetError("no asset to download")
			break
		}
		ghr.File.Name = file

		ghr.message("Downloading release asset ...")

		// Find the release corresponding to the entered tag, if any.
		var rel *Release
		//if ghr.Repo.Tag == "latest" {
		//	rel, ghr.State = ghr.LatestRelease()
		//} else {
			rel, ghr.State = ghr.releaseOfTag()
		//}
		if ghr.State.IsNotOk() {
			break
		}

		asset := findAsset(rel.Assets, ghr.File.Name)
		if asset == nil {
			ghr.State.SetError("could not find asset named %s", ghr.File.Name)
			break
		}

		var resp *http.Response
		var err error
		if ghr.Auth.Token == "" {
			// Use the regular github.com site if we don't have a token.
			resp, err = http.Get(DefaultGitHubUrl + fmt.Sprintf("/%s/%s/releases/download/%s/%s", ghr.Repo.Organization, ghr.Repo.Name, ghr.Repo.Tag, ghr.File.Name))
		} else {
			//u := nvls(ghr.urlPrefix, github.DefaultBaseURL) + fmt.Sprintf(AssetUri, ghr.Repo.Organization, ghr.Repo.Name, asset.Id)
			u := fmt.Sprintf(AssetUri, ghr.Repo.Organization, ghr.Repo.Name, asset.Id)
			resp, err = github.DoAuthRequest("GET", u, "", ghr.Auth.Token, map[string]string{
				"Accept": "application/octet-stream",
			}, nil)
		}
		//noinspection ALL
		defer resp.Body.Close()

		if err != nil {
			ghr.State.SetError("could not fetch releases, %v", err)
			break
		}

		ghr.message("GET", resp.Request.URL, "->", resp)

		contentLength, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
		if err != nil {
			ghr.State.SetError(err)
			break
		}

		if resp.StatusCode != http.StatusOK {
			ghr.State.SetError("github did not respond with 200 OK but with %v", resp.Status)
			break
		}

		out := os.Stdout // Pipe the asset to stdout by default.
		if isCharDevice(out) {
			// If stdout is a char device, assume it's a TTY (terminal). In this
			// case, don't pipe th easset to stdout, but create it as a file in
			// the current working folder.
			if out, err = os.Create(ghr.File.Name); err != nil {
				ghr.State.SetError("could not create file %s", ghr.File.Name)
				break
			}
			//noinspection ALL
			defer out.Close()
		}

		ghr.State.SetOk()
		ghr.State.SetResponse(ghr.mustCopyN(out, resp.Body, contentLength))
	}

	//return mustCopyN(out, resp.Body, contentLength)
	return ghr.State
}


func (ghr *TypeGhr) Release() *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State = ghr.Repo.isValidTag()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.message("Releasing '%s' ...", ghr.Repo.Tag)

		//ghr.State = ghr.validateCredentials()
		//if ghr.State.IsNotOk() {
		//	break
		//}

		// Check if we need to read the description from stdin.
		if ghr.Repo.Description == "-" {
			b, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				ghr.State.SetError("could not read description from stdin: %v", err)
				break
			}
			ghr.Repo.Description = string(b)
		}

		params := ReleaseCreate{
			TagName:         ghr.Repo.Tag,
			TargetCommitish: ghr.Repo.Target,
			Name:            ghr.File.Name,
			Body:            ghr.Repo.Description,
			Draft:           ghr.Repo.Draft,
			Prerelease:      ghr.Repo.Prerelease,
		}

		/* encode params as json */
		payload, err := json.Marshal(params)
		if err != nil {
			ghr.State.SetError("can't encode release creation params, %v", err)
			break
		}
		reader := bytes.NewReader(payload)

		//URL := nvls(ghr.urlPrefix, github.DefaultBaseURL) + fmt.Sprintf("/repos/%s/%s/releases", ghr.Repo.Organization, ghr.Repo.Name)
		URL := fmt.Sprintf("/repos/%s/%s/releases", ghr.Repo.Organization, ghr.Repo.Name)
		resp, err := github.DoAuthRequest("POST", URL, "application/json", ghr.Auth.Token, nil, reader)
		if err != nil {
			ghr.State.SetError("while submitting %v, %v", string(payload), err)
			break
		}
		//noinspection ALL
		defer resp.Body.Close()

		ghr.message("RESPONSE:", resp)
		if resp.StatusCode != http.StatusCreated {
			if resp.StatusCode == 422 {
				ghr.State.SetError("github returned %v (this is probably because the release already exists)", resp.Status)
				break
			}
			ghr.State.SetError("github returned %v", resp.Status)
			break
		}

		if ghr.runtime.Debug {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				ghr.State.SetError("error while reading response, %v", err)
				break
			}
			ghr.message("BODY:", string(body))
		}

		ghr.State.SetOk()
	}

	// return nil
	return ghr.State
}


func (ghr *TypeGhr) Update() *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State = ghr.Repo.isValidTag()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.message("Updating release '%s' ...", ghr.Repo.Tag)

		//ghr.State = ghr.validateCredentials()
		//if ghr.State.IsNotOk() {
		//	break
		//}

		var id int
		id, ghr.State = ghr.idOfTag()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.message("Release %s has id %d", ghr.Repo.Tag, id)
		// Check if we need to read the description from stdin.
		if ghr.Repo.Description == "-" {
			b, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				ghr.State.SetError("could not read description from stdin: %v", err)
				break
			}
			ghr.Repo.Description = string(b)
		}

		/* the release create struct works for editing releases as well */
		params := ReleaseCreate{
			TagName:    ghr.Repo.Tag,
			Name:       ghr.File.Name,
			Body:       ghr.Repo.Description,
			Draft:      ghr.Repo.Draft,
			Prerelease: ghr.Repo.Prerelease,
		}

		/* encode the parameters as JSON, as required by the github API */
		payload, err := json.Marshal(params)
		if err != nil {
			ghr.State.SetError("can't encode release creation params, %v", err)
			break
		}

		//URL := nvls(ghr.urlPrefix, github.DefaultBaseURL) + fmt.Sprintf("/repos/%s/%s/releases/%d", ghr.Repo.Organization, ghr.Repo.Name, id)
		URL := fmt.Sprintf("/repos/%s/%s/releases/%d", ghr.Repo.Organization, ghr.Repo.Name, id)
		resp, err := github.DoAuthRequest("PATCH", URL, "application/json", ghr.Auth.Token, nil, bytes.NewReader(payload))
		if err != nil {
			ghr.State.SetError("while submitting %v, %v", string(payload), err)
			break
		}
		//noinspection ALL
		defer resp.Body.Close()

		ghr.message("RESPONSE:", resp)
		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == 422 {
				ghr.State.SetError("github returned %v (this is probably because the release already exists)", resp.Status)
				break
			}
			ghr.State.SetError("github returned unexpected status code %v", resp.Status)
			break
		}

		if ghr.runtime.Debug {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				ghr.State.SetError("error while reading response, %v", err)
				break
			}
			ghr.message("BODY:", string(body))
		}

		ghr.State.SetOk()
	}

	// return nil
	return ghr.State
}


func (ghr *TypeGhr) Delete() *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}
		ghr.message("Deleting release '%s' ...", ghr.Repo.Tag)

		var id int
		id, ghr.State = ghr.idOfTag()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.message("Release %v has id %v", ghr.Repo.Tag, id)

		baseURL := nvls(ghr.urlPrefix, github.DefaultBaseURL)
		resp, err := github.DoAuthRequest("DELETE", baseURL+fmt.Sprintf("/repos/%s/%s/releases/%d",
			ghr.Repo.Organization, ghr.Repo.Name, id), "application/json", ghr.Auth.Token, nil, nil)
		if err != nil {
			ghr.State.SetError("release deletion failed: %v", err)
			break
		}
		//noinspection ALL
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			ghr.State.SetError("could not delete the release corresponding to tag %s on repo %s/%s", ghr.Repo.Tag, ghr.Repo.Organization, ghr.Repo.Name)
			break
		}

		ghr.State.SetOk()
	}

	// return nil
	return ghr.State
}
