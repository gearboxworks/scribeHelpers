package toolGhr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gearboxworks/scribeHelpers/toolGhr/github"
	"github.com/gearboxworks/scribeHelpers/toolPath"
	"github.com/gearboxworks/scribeHelpers/ux"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)


// Copy repo release from source to destination.
func CopyReleases(srcRepoUrl string, srcTag string, destRepoUrl string, cacheDir ...string) *ux.State {
	state := ux.NewState("CopyReleases", false)

	for range onlyOnce {
		// Setup source repo.
		Src := New(nil)
		if Src.State.IsNotOk() {
			state = Src.State
			break
		}
		//state = Src.SetAuth(TypeAuth{ Token: "", AuthUser: "" })
		//if state.IsNotOk() {
		//	break
		//}
		state = Src.OpenUrl(srcRepoUrl)
		if state.IsNotOk() {
			break
		}


		// Setup destination repo.
		Dest := New(nil)
		if Src.State.IsNotOk() {
			state = Src.State
			break
		}
		//state = Src.SetAuth(TypeAuth{ Token: "", AuthUser: "" })
		//if state.IsNotOk() {
		//	break
		//}
		state = Dest.OpenUrl(destRepoUrl)
		if state.IsNotOk() {
			break
		}
		state = Dest.SetOverwrite(true)
		if state.IsNotOk() {
			break
		}


		// Now sync the release in the destination repo.
		state = Dest.CopyReleasesFrom(Src.Repo, srcTag, cacheDir...)
	}

	return state
}


// Return asset for current architecture.
func GetAsset(repoUrl string, tag string) (string, *ux.State) {
	state := ux.NewState("GetAsset", false)
	var ret string

	for range onlyOnce {
		// Setup source repo.
		repo := New(nil)
		if repo.State.IsNotOk() {
			state = repo.State
			break
		}
		//state = repo.SetAuth(TypeAuth{ Token: "", AuthUser: "" })
		//if state.IsNotOk() {
		//	break
		//}
		state = repo.OpenUrl(repoUrl)
		if state.IsNotOk() {
			break
		}

		state = repo.SetTag(tag)
		if state.IsNotOk() {
			break
		}

		ret, state = repo.GetAsset()
		if state.IsNotOk() {
			break
		}

		state.SetOk()
	}

	return ret, state
}


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


// Show repo information
func (ghr *TypeGhr) CopyReleasesFrom(srcRepo *TypeRepo, srcTag string, cacheDir ...string) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}


		// Setup cache dir.
		if len(cacheDir) == 0 {
			cacheDir = []string{"dist"}
		}
		ghr.message("Setting up cache directory '%s' ...", filepath.Join(cacheDir...))
		dir := toolPath.New(ghr.runtime)
		if ghr.State.IsNotOk() {
			break
		}
		dir.SetPath(cacheDir...)
		ghr.State = dir.StatPath()
		if ghr.State.IsError() {
			break
		}


		// Setup src repo.
		ghr.message("Setting up source repo '%s' ...", srcRepo.GetUrl())
		ghr.State = srcRepo.Fetch(true)
		if ghr.State.IsError() {
			break
		}
		if srcTag != "" {
			ghr.State = srcRepo.SetTag(srcTag)
			if ghr.State.IsError() {
				ghr.State.SetError("No source release found")
				break
			}
		}
		srcRef := srcRepo.Release()
		if srcRef == nil {
			ghr.State.SetError("No source release found")
			break
		}


		// Setup destination repo.
		dstRepo := ghr.Repo
		ghr.message("Setup destination repo '%s' ...", dstRepo.GetUrl())
		ghr.State = dstRepo.Fetch(true)
		if ghr.State.IsError() {
			break
		}

		//dstRepo.Organization = "",
		//dstRepo.Name =         "",
		//dstRepo.Auth =         dstRepo.Auth,
		dstRepo.TagName     = srcRef.TagName
		dstRepo.Description = srcRef.Description
		dstRepo.Draft       = srcRef.Draft
		dstRepo.Prerelease  = srcRef.Prerelease
		dstRepo.Target      = srcRepo.Target
		//dstRepo.Go       = srcRepo.Go
		//dstRepo.Overwrite     = srcRepo.Overwrite

		dstRepo.Files = []string{}
		for _, file := range srcRepo.Files {
			dstRepo.Files = append(dstRepo.Files, filepath.Join(filepath.Join(cacheDir...), file))
		}


		// Copy src files to cache.
		ghr.message("Download %d files from source repo to cache dir '%s' ...", len(srcRepo.Files), dir.GetPath())
		ghr.State = srcRepo.DownloadAssets(false, cacheDir...)
		if ghr.State.IsError() {
			break
		}


		// Create release on destination repo.
		ghr.message("Creating release on destination repo...")
		ghr.State = dstRepo.CreateRelease(nil)
		if ghr.State.IsNotOk() {
			ghr.State.SetError("could not create release '%s'", ghr.Repo.TagName)
			break
		}


		// Upload files
		ghr.message("Uploading assets to destination repo...")
		for _, file := range dstRepo.Files {
			ghr.State = ghr.Upload(dstRepo.Overwrite, file)
			if ghr.State.IsNotOk() {
				// Retry same file again if failed.
				ghr.State = ghr.Upload(dstRepo.Overwrite, file)
				break
			}
		}


		ghr.State = dstRepo.Fetch(true)
		if ghr.State.IsError() {
			break
		}


		ghr.message("Destination repo now in sync for Release '%s'.", dstRepo.TagName)
		//srcRepo.PrintRelease()
		dstRepo.PrintRelease()

		ghr.State.SetOk()
	}

	return ghr.State
}


// Upload multiple files to a repo Release.
func (ghr *TypeGhr) UploadMultiple(overwrite bool, files ...string) *ux.State {
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
			ghr.State = ghr.Upload(overwrite, file)
			if ghr.State.IsNotOk() {
				// Retry same file again if failed.
				ghr.State = ghr.Upload(overwrite, file)
				break
			}
		}
	}

	return ghr.State
}


// Upload a file to a repo Release.
func (ghr *TypeGhr) Upload(overwrite bool, label string, path ...string) *ux.State {
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

		ghr.State = ghr.Repo.UploadAsset(overwrite, label, path...)
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State.SetOk()
	}

	return ghr.State
}


// Download a file from a repo Release.
func (ghr *TypeGhr) Download(overwrite bool, name string, path ...string) *ux.State {
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

		ghr.State = ghr.Repo.DownloadAsset(overwrite, name, path...)
		if ghr.State.IsError() {
			break
		}

		ghr.State.SetOk()
	}

	return ghr.State
}


// Upload multiple files to a repo Release.
func (ghr *TypeGhr) DeleteAssets(labels ...string) *ux.State {
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

		for _, file := range labels {
			ghr.State = ghr.Repo.DeleteAsset(file)
			if ghr.State.IsNotOk() {
				break
			}
		}
	}

	return ghr.State
}


// Show repo information
func (ghr *TypeGhr) GetAsset() (string, *ux.State) {
	var ret string
	if state := ghr.IsNil(); state.IsError() {
		return ret, state
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

		label := fmt.Sprintf("(?i).*%s_%s.*",
			ghr.runtime.GoRuntime.Os,
			ghr.runtime.GoRuntime.Arch,
			)

		asset := ghr.Repo.SelectRegexpAsset(label)
		if asset == nil {
			break
		}

		ret = asset.Name
		ghr.State.SetOk()
	}

	return ret, ghr.State
}


// Create a repo Release.
func (ghr *TypeGhr) Create(rel TypeRepo) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.Repo.Fetch(true)
		if ghr.State.IsError() {
			break
		}

		ghr.State = ghr.Set(rel)
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State = ghr.isValid()
		if ghr.State.IsNotOk() {
			break
		}


		// Create release
		ghr.State = ghr.Repo.CreateRelease(&rel)
		if ghr.State.IsNotOk() {
			ghr.State.SetError("could not create release '%s'", ghr.Repo.TagName)
			break
		}

		ghr.State = ghr.Repo.Fetch(true)
		if ghr.State.IsError() {
			break
		}

		// Upload files
		ghr.State = ghr.UploadMultiple(ghr.Repo.Overwrite, ghr.Repo.Files...)
		if ghr.State.IsError() {
			break
		}
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
			Name:       ghr.Repo.TagName,
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

		ghr.State = ghr.Repo.DeleteRelease(tag)
	}

	return ghr.State
}
