package toolGhr

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/toolGhr/github"
	"github.com/newclarity/scribeHelpers/ux"
	"net/http"
	"time"
)

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
type Assets struct {
	All []*Asset
	Selected *Asset
}


// findAsset returns the asset if an asset with name can be found in assets,
// otherwise returns nil.
func (a *Assets) findAsset(name string) *Asset {
	for range onlyOnce {
		a.Selected = nil

		if name == "latest" {
			// Latest will always be first... Maybe... @TODO - TO BE CHECKED
			a.Selected = a.All[0]
			break
		}

		for _, asset := range a.All {
			if asset.Name == name {
				a.Selected = asset
			}
		}
	}

	return a.Selected
}


// Delete sends a HTTP DELETE request for the given asset to Github. Returns
// nil if the asset was deleted OR there was nothing to delete.
func (ghr *TypeGhr) DeleteAsset(a *Asset) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		//URL := nvls(ghr.urlPrefix, github.DefaultBaseURL) + fmt.Sprintf(AssetUri, ghr.Repo.Organization, ghr.Repo.Name, a.Id)
		URL := fmt.Sprintf(AssetUri, ghr.Repo.Organization, ghr.Repo.Name, a.Id)
		resp, err := github.DoAuthRequest("DELETE", URL, "application/json", ghr.Auth.Token, nil, nil)
		if err != nil {
			ghr.State.SetError("failed to delete asset %s (ID: %d), HTTP error: %b", a.Name, a.Id, err)
			break
		}
		//noinspection ALL
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusNoContent {
			ghr.State.SetError("failed to delete asset %s (ID: %d), status: %s", a.Name, a.Id, resp.Status)
			break
		}
	}
	return nil
}


//func (a *Asset) Delete() error {
//	{
//		URL := nvls(ghr.urlPrefix, github.DefaultBaseURL) + fmt.Sprintf(AssetUri, ghr.Auth.User, ghr.Repo.Name, a.Id)
//		resp, err := github.DoAuthRequest("DELETE", URL, "application/json", ghr.Auth.Token, nil, nil)
//		if err != nil {
//			return fmt.Errorf("failed to delete asset %s (ID: %d), HTTP error: %b", a.Name, a.Id, err)
//		}
//		//noinspection ALL
//		defer resp.Body.Close()
//		if resp.StatusCode != http.StatusNoContent {
//			return fmt.Errorf("failed to delete asset %s (ID: %d), status: %s", a.Name, a.Id, resp.Status)
//		}
//	}
//	return nil
//}
