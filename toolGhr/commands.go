package toolGhr

import (
	"bytes"
	"encoding/json"
	"github.com/newclarity/scribeHelpers/toolGhr/github"
	"github.com/newclarity/scribeHelpers/ux"
	"io/ioutil"
	"net/http"
	"os"
)


// Show repo information
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

		ghr.State = ghr.Repo.Fetch(true)
		if ghr.State.IsError() {
			break
		}

		ghr.message("Getting repo tag info ...")
		ghr.State = ghr.Repo.PrintTags()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.message("Getting repo Release info ...")
		ghr.State = ghr.Repo.PrintReleases()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State.SetOk()
	}

	return ghr.State
}


// Upload multiple files to a repo Release.
func (ghr *TypeGhr) UploadMultiple(replace bool, files ...string) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}

		for _, file := range files {
			ghr.State = ghr.Upload(replace, file, "")
			if ghr.State.IsNotOk() {
				break
			}
		}
	}

	return ghr.State
}


// Upload a file to a repo Release.
func (ghr *TypeGhr) Upload(replace bool, file string, label string) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.Repo.SetFile(file, label)

		ghr.State = ghr.Repo.Fetch(true)
		if ghr.State.IsError() {
			break
		}


		//// Incomplete (failed) uploads will have their state set to new. These
		//// assets are (AFAIK) useless in all cases. The only thing they will do
		//// is prevent the upload of another asset of the same name. To work
		//// around this GH API weirdness, let's just delete assets if:
		////
		//// 1. Their state is new.
		//// 2. The user explicitly asked to delete/replace the asset with -R.
		for range onlyOnce {
			asset := ghr.Repo.SelectAsset(file)
			if asset == nil {
				break
			}

			//if asset.State != "new" {
			//	break
			//}

			if replace == false {
				break
			}

			ghr.message("Asset (id: %d) exists in state %s: Removing ...", asset.Id, asset.State)
			ghr.State = ghr.Repo.DeleteAsset(asset)
			if ghr.State.IsNotOk() {
				ghr.State.SetError("could not replace asset: %v", ghr.State.GetError())
				break
			}
		}

		ghr.State = ghr.Repo.UploadAsset(file)
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State.SetOk()
	}

	return ghr.State
}


// Download a file from a repo Release.
func (ghr *TypeGhr) Download(file string, overwrite bool) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.Repo.SetFile(file, "")

		ghr.State = ghr.Repo.Fetch(true)
		if ghr.State.IsError() {
			break
		}

		ghr.State = ghr.Repo.DownloadAsset(file, overwrite)
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State.SetOk()
	}

	return ghr.State
}


// Create a repo Release.
func (ghr *TypeGhr) CreateRelease(n TypeRepo) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.Set(n)
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}


		// Create release
		ghr.State = ghr.Repo.Fetch(true)
		if ghr.State.IsError() {
			break
		}

		for range onlyOnce {
			rel := ghr.Repo.SelectRelease(ghr.Repo.TagName)
			if rel == nil {
				break
			}

			if ghr.Repo.Replace == false {
				break
			}

			ghr.message("Release (id: %d) exists in state %s: Removing ...", rel.Id, rel.Name)
			ghr.State = ghr.Repo.DeleteRelease(rel)
			if ghr.State.IsNotOk() {
				ghr.State.SetError("could not replace release: %v", rel.Name)
				break
			}

			ghr.State = ghr.Repo.Fetch(true)
			if ghr.State.IsError() {
				break
			}
		}

		ghr.State = ghr.Repo.CreateRelease(ghr.Repo.TagName)
		if ghr.State.IsNotOk() {
			ghr.State.SetError("could not create release '%s'", ghr.Repo.TagName)
			break
		}

		ghr.State = ghr.Repo.Fetch(true)
		if ghr.State.IsError() {
			break
		}


		// Upload files
		for _, file := range ghr.Repo.Files {
			for range onlyOnce {
				asset := ghr.Repo.SelectAsset(file)
				if asset == nil {
					break
				}

				if ghr.Repo.Replace == false {
					break
				}

				ghr.message("Asset (id: %d) exists in state %s: Removing ...", asset.Id, asset.State)
				ghr.State = ghr.Repo.DeleteAsset(asset)
				if ghr.State.IsNotOk() {
					ghr.State.SetError("could not replace asset: %v", ghr.State.GetError())
					break
				}
			}

			ghr.State = ghr.Repo.UploadAsset(file)
			if ghr.State.IsNotOk() {
				break
			}
		}
	}

	return ghr.State
}


// Create a repo Release.
func (ghr *TypeGhr) Create(tag string, replace bool) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State = ghr.SetTag(tag)
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State = ghr.Repo.Fetch(true)
		if ghr.State.IsError() {
			break
		}

		for range onlyOnce {
			rel := ghr.Repo.SelectRelease(tag)
			if rel == nil {
				break
			}

			if replace == false {
				break
			}

			ghr.message("Release (id: %d) exists. Removing ...", rel.Id)
			ghr.State = ghr.Repo.DeleteRelease(rel)
			if ghr.State.IsNotOk() {
				ghr.State.SetError("could not replace release: %s", rel.Name)
				break
			}
		}

		ghr.State = ghr.Repo.CreateRelease(tag)
		if ghr.State.IsNotOk() {
			ghr.State.SetError("could not create release '%s'", tag)
			break
		}

		ghr.State.SetOk()
	}

	return ghr.State
}


// Update a repo Release.
func (ghr *TypeGhr) Update(tag string) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State = ghr.Repo.Fetch(true)
		if ghr.State.IsError() {
			break
		}

		rel := ghr.Repo.SelectRelease(tag)
		if rel == nil {
			ghr.State.SetError("no valid release found")
			break
		}


		ghr.message("Release %s has id %d", ghr.Repo.TagName, rel.Id)
		// Check if we need to read the description from stdin.
		if ghr.Repo.Description == "-" {
			b, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				ghr.State.SetError("could not read description from stdin: %v", err)
				break
			}
			ghr.Repo.Description = string(b)
		}

		/* the Release create struct works for editing releases as well */
		params := releaseCreate{
			TagName:    ghr.Repo.TagName,
			Name:       ghr.Repo.file.Name,
			Body:       ghr.Repo.Description,
			Draft:      ghr.Repo.Draft,
			Prerelease: ghr.Repo.Prerelease,
		}

		/* encode the parameters as JSON, as required by the github API */
		payload, err := json.Marshal(params)
		if err != nil {
			ghr.State.SetError("can't encode Release creation params, %v", err)
			break
		}


		ghr.message("Updating Release '%s' ...", ghr.Repo.TagName)
		URL := ghr.Repo.generateApiUrl(releaseIdUri, rel.Id)
		resp, err := github.DoAuthRequest("PATCH", URL, "application/json", ghr.Repo.Auth.Token, nil, bytes.NewReader(payload))
		if err != nil {
			ghr.State.SetError("while submitting %v, %v", string(payload), err)
			break
		}
		//noinspection ALL
		defer resp.Body.Close()

		ghr.message("RESPONSE:", resp)
		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == 422 {
				//ghr.State.SetError("github returned %v (this is probably because the Release already exists)", resp.Status)
				ghr.State.SetError("release '%s' already exists")
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

	return ghr.State
}


// Delete a repo Release.
func (ghr *TypeGhr) Delete(tag string) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State = ghr.SetTag(tag)
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State = ghr.Repo.Fetch(true)
		if ghr.State.IsError() {
			break
		}

		rel := ghr.Repo.SelectRelease(tag)
		if rel == nil {
			ghr.State.SetError("no valid release found")
			break
		}

		ghr.State = ghr.Repo.DeleteRelease(rel)

		//ghr.message("Deleting Release '%s' ...", ghr.Repo.Tag)
		//URL := ghr.Repo.generateApiUrl(releaseIdUri, rel.Id)
		//resp, err := github.DoAuthRequest("DELETE", URL, "application/json", ghr.Repo.Auth.Token, nil, nil)
		//if err != nil {
		//	ghr.State.SetError("Release deletion failed: %v", err)
		//	break
		//}
		////noinspection ALL
		//defer resp.Body.Close()
		//
		//if resp.StatusCode != http.StatusNoContent {
		//	ghr.State.SetError("could not delete the Release corresponding to tag %s on repo %s/%s", ghr.Repo.Tag, ghr.Repo.Organization, ghr.Repo.Name)
		//	break
		//}
		//
		//
		//ghr.message("Deleting Release Tag '%s' ...", ghr.Repo.Tag)
		//URL = ghr.Repo.generateApiUrl(tagRef, ghr.Repo.Tag)
		//resp, err = github.DoAuthRequest("DELETE", URL, "application/json", ghr.Repo.Auth.Token, nil, nil)
		//if err != nil {
		//	ghr.State.SetError("Release deletion failed: %v", err)
		//	break
		//}
		////noinspection ALL
		//defer resp.Body.Close()
		//
		//if resp.StatusCode != http.StatusNoContent {
		//	ghr.State.SetError("could not delete the Release corresponding to tag %s on repo %s/%s", ghr.Repo.Tag, ghr.Repo.Organization, ghr.Repo.Name)
		//	break
		//}
	}

	return ghr.State
}
