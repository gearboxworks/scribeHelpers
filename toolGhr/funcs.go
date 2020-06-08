package toolGhr

import (
	"encoding/json"
	"fmt"
	"github.com/newclarity/scribeHelpers/toolGhr/github"
	"github.com/newclarity/scribeHelpers/ux"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)


/* usually when something goes wrong, github sends something like this back */
type Message struct {
	Message string        `json:"message"`
	Errors  []GithubError `json:"errors"`
}

type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
}


/* transforms a stream into a Message, if it's valid json */
func ToMessage(r io.Reader) (*Message, error) {
	var msg Message
	if err := json.NewDecoder(r).Decode(&msg); err != nil {
		return nil, err
	}

	return &msg, nil
}


func (m *Message) String() string {
	str := fmt.Sprintf("msg: %v, errors: ", m.Message)

	errstr := make([]string, len(m.Errors))
	for idx, err := range m.Errors {
		errstr[idx] = fmt.Sprintf("[field: %v, code: %v]",
			err.Field, err.Code)
	}

	return str + strings.Join(errstr, ", ")
}


/* nvls returns the first value in xs that is not empty. */
func nvls(xs ...string) string {
	for _, s := range xs {
		if s != "" {
			return s
		}
	}

	return ""
}


// formats time `t` as `fmt` if it is not nil, otherwise returns `def`
func timeFmtOr(t *time.Time, fmt, def string) string {
	if t == nil {
		return def
	}
	return t.Format(fmt)
}


// isCharDevice returns true if f is a character device (panics if f can't
// be stat'ed).
func isCharDevice(f *os.File) bool {
	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	return (stat.Mode() & os.ModeCharDevice) != 0
}


func Mark(ok bool) string {
	if ok {
		return "✔"
	} else {
		return "✗"
	}
}

const (
	TagsUri = "/repos/%s/%s/tags"

	// GET /repos/:owner/:repo/releases/assets/:id
	// DELETE /repos/:owner/:repo/releases/assets/:id
	AssetUri = "/repos/%s/%s/releases/assets/%d"

	// API: https://developer.github.com/v3/repos/releases/#list-assets-for-a-release
	// GET /repos/:owner/:repo/releases/:id/assets
	AssetReleaseListUri = "/repos/%s/%s/releases/%d/assets"

)


type Tag struct {
	Name       string `json:"name"`
	Commit     Commit `json:"commit"`
	ZipBallUrl string `json:"zipball_url"`
	TarBallUrl string `json:"tarball_url"`
}

type Commit struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}


func (t *Tag) String() string {
	return t.Name + " (commit: " + t.Commit.Url + ")"
}


// Get the tags associated with a repo.
func (ghr *TypeGhr) Tags() ([]Tag, *ux.State) {
	var tags []Tag
	if state := ghr.IsNil(); state.IsError() {
		return tags, state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		//client := github.NewClient(ghr.Auth.AuthUser, ghr.Auth.Token, nil)
		//client.SetBaseURL(ghr.urlPrefix)
		err := ghr.Repo.ApiGet(TagsUri, ghr.Repo.Organization, ghr.Repo.Name)
		if err != nil {
			ghr.State.SetError(err)
			break
		}

		if ghr.State.IsResponseNotOfType("[]Tag") {
			ghr.State.SetError("could not get tags")
			break
		}

		tags = ghr.State.GetResponseData().([]Tag)
	}

	return tags, ghr.State
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


// findAsset returns the asset if an asset with name can be found in assets,
// otherwise returns nil.
func findAsset(assets []Asset, name string) *Asset {
	for _, asset := range assets {
		if asset.Name == name {
			return &asset
		}
	}
	return nil
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


// mustCopyN attempts to copy exactly N bytes, if this fails, an error is
// returned.
func (ghr *TypeGhr) mustCopyN(w io.Writer, r io.Reader, n int64) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		an, err := io.Copy(w, r)
		if an != n {
			ghr.State.SetError("data did not match content length %d != %d", an, n)
			break
		}
		ghr.State.SetError(err)
		break
	}

	return ghr.State
}


func (ghr *TypeGhr) renderInfoText(tags []Tag, releases *Releases) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		fmt.Println("tags:")
		for _, tag := range tags {
			fmt.Println("-", &tag)
		}

		fmt.Println("releases:")
		for _, release := range *releases {
			fmt.Println("-", &release)
		}

		ghr.State.SetOk()
	}

	//return nil
	return ghr.State
}


func (ghr *TypeGhr) renderInfoJSON(tags []Tag, releases *Releases) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		out := struct {
			Tags     []Tag
			Releases *Releases
		}{
			Tags:     tags,
			Releases: releases,
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "    ")

		ghr.State.SetOk()
		ghr.State.SetResponse(enc.Encode(&out))
	}

	return ghr.State
}
