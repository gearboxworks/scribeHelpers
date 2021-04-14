package toolGhr

import (
	"fmt"
	"github.com/gearboxworks/scribeHelpers/ux"
)

const (
	repoBaseUrl         = "%s/repos/%s/%s"

	releaseListUri		= "/releases"
	releaseLatestUri	= "/releases/" + Latest
	releaseTagUri 		= "/releases/%s"
	releaseAssetDownload = "/releases/download/%s/%s"
	releaseIdUri		= "/releases/%d"

	releaseDateFormat	= "02/01/2006 at 15:04"
	Latest				= "latest"

	tagsUri				= "/tags"
	tagLatestUri		= "/tags/" + Latest
	tagUri				= "/tags/%s"

	// DELETE /repos/:owner/:repo/git/refs/:ref
	tagRef				= "/git/refs/tags/%s"

	// GET /repos/:owner/:repo/releases/assets/:id
	// DELETE /repos/:owner/:repo/releases/assets/:id
	AssetUri			= "/releases/assets/%d"

	// API: https://developer.github.com/v3/repos/releases/#list-assets-for-a-release
	// GET /repos/:owner/:repo/releases/:id/assets
	assetsUri			= "/releases/%d/assets"
)

func (repo *TypeRepo) ClientGet(ref interface{}, uri string, args ...interface{}) *ux.State {
	if State := ux.IfNilReturnError(repo); State.IsError() {
		return State
	}

	for range onlyOnce {
		URL := repo.generateApiUrl(uri, args...)
		err := repo.client.Get(URL, ref)
		if err != nil {
			repo.state.SetError(err)
			break
		}
		if ref == nil {
			repo.state.SetWarning("no results found")
			break
		}
	}

	return repo.state
}

func (repo *TypeRepo) generateApiUrl(format string, args ...interface{}) string {
	if state := repo.IsNil(); state.IsError() {
		return ""
	}
	repo.state.SetFunction()
	base := ""
	if repo.Auth.Token == "" {
		base = repo.urlPrefix
	} else {
		base = repo.apiUrlPrefix
	}
	prefix := fmt.Sprintf(repoBaseUrl, base, repo.Organization, repo.Name)
	return prefix + fmt.Sprintf(format, args...)
}
