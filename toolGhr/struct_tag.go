package toolGhr

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/toolGhr/github"
	"github.com/newclarity/scribeHelpers/ux"
	"net/http"
	"strings"
)


// Get the tags associated with a repo.
func (repo *TypeRepo) FetchTags(force bool) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if force {
			repo.tags.all = nil
		}
		if repo.tags.all != nil {
			break
		}

		repo.state = repo.ClientGet(&repo.tags.all, tagsUri)
		if repo.state.IsNotOk() {
			repo.state.SetError("no tags found")
			break
		}
		//URL := repo.generateApiUrl(tagsUri)
		//err := repo.client.Get(URL, &repo.tags.all)
		//if err != nil {
		//	repo.state.SetError(err)
		//	break
		//}
		//if repo.tags.all == nil {
		//	repo.state.SetError("no tags found")
		//	break
		//}

		// @TODO - figure out how to do this.
		//// Sometimes we can't second guess what the "latest" is based on date alone.
		//u = fmt.Sprintf(tagUri, repo.Organization, repo.Name, Latest)
		//err = repo.client.Get(u, &repo.tags.latest)
		//if err != nil {
		//	repo.state.SetError(err)
		//	break
		//}
		repo.tags.latest = repo.tags.GetLatest()

		if repo.tags.findTag(repo.TagName) == nil {
			repo.state.SetWarning("no tag '%s' found", repo.TagName)
			break
		}

		repo.state.SetOk()
		repo.state.SetResponse(&repo.tags)
		// Allows the use of the following in a calling function:
		//if repo.state.IsResponseNotOfType("tags") {
		//	repo.state.SetError("could not get tags")
		//	break
		//}
		//tags := repo.state.GetResponseData().(*tags)
	}

	return repo.state
}

func (repo *TypeRepo) Tag() *Tag {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return nil
	}
	repo.state.SetFunction()
	return repo.tags.GetSelected()
}

func (repo *TypeRepo) Tags() []*Tag {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return nil
	}
	repo.state.SetFunction()
	return repo.tags.GetAll()
}

func (repo *TypeRepo) CountTags() int {
	if state := repo.IsNil(); state.IsError() {
		return 0
	}
	repo.state.SetFunction()
	return repo.tags.CountAll()
}

func (repo *TypeRepo) PrintTags() *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()
	repo.tags.Print()
	return repo.state
}

func (repo *TypeRepo) PrintTag() *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()
	if repo.tags.selected != nil {
		repo.tags.selected.Print()
	}
	return repo.state
}

func (repo *TypeRepo) SelectTag(tag string) *Tag {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return nil
	}
	repo.state.SetFunction()
	return repo.tags.findTag(tag)
}

// Delete sends a HTTP DELETE request for the given asset to Github. Returns
// nil if the asset was deleted OR there was nothing to delete.
func (repo *TypeRepo) DeleteTag(tag string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if repo.tags.all == nil {
			repo.state.SetError("No tags available")
			break
		}

		ref := repo.SelectTag(tag)
		if ref == nil {
			repo.state.SetError("Release tag '%s' not available", tag)
			break
		}

		repo.state = repo.DeleteTagRef(ref)
	}

	return repo.state
}

func (repo *TypeRepo) DeleteTagRef(ref *Tag) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if ref == nil {
			repo.messageError("Deleting Release tag FAILED - empty")
			repo.state.SetError("Deleting Release tag FAILED - empty")
			break
		}

		repo.message("Deleting Release tag '%s' ...", ref.Name)
		URL := repo.generateApiUrl(tagRef, ref.Name)
		resp, err := github.DoAuthRequest("DELETE", URL, "application/json", repo.Auth.Token, nil, nil)
		if err != nil {
			repo.messageError("Deleting Release tag '%s' FAILED", ref.Name)
			repo.state.SetError("Release deletion failed: %v", err)
			break
		}
		//noinspection ALL
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			repo.messageError("Deleting Release tag '%s' FAILED", ref.Name)
			repo.state.SetError("failed to delete tag %s - status: %s", ref.Name, resp.Status)
			break
		}

		repo.messageOk("Deleted Release tag '%s' OK", ref.Name)
		repo.state.SetOk()
	}

	return repo.state
}


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
	if t == nil {
		return ""
	}
	return t.Name + " (commit: " + t.Commit.Url + ")"
}

func (t *Tag) Print() {
	fmt.Print(t.String())
}


type tags struct {
	all      []*Tag
	selected *Tag
	latest   *Tag
}

func (t *tags) GetAll() []*Tag {
	return t.all
}

func (t *tags) GetSelected() *Tag {
	return t.selected
}

func (t *tags) GetLatest() *Tag {
	var rel *Tag
	for range onlyOnce {
		if t.latest != nil {
			rel = t.latest
			break
		}

		// Latest will always be first... Maybe... @TODO - TO BE CHECKED
		rel = t.all[0]
	}
	return rel
}

func (t *tags) CountAll() int {
	return len(t.all)
}

func (t *tags) Sprint() string {
	var ret string
	switch {
		case t.all == nil:
			ret += ux.SprintfWarning("No repo tags found")

		case t.selected == nil:
			// Print all tags.
			ret += ux.SprintfWarning("Found %d tags.", t.CountAll())
			var s []string
			for _, tag := range t.all {
				s = append(s, tag.Name)
			}
			ret += ux.SprintfWarning("Repo tags: %s", strings.Join(s, ", "))

		default:
			// Print selected tag.
			ret += ux.SprintfWarning("Repo tag: %s", t.selected.Name)
	}
	return ret
}

func (t *tags) Print() {
	fmt.Println(t.Sprint())
}

func (t *tags) findTag(name string) *Tag {
	for range onlyOnce {
		t.selected = nil

		if name == Latest {
			t.selected = t.GetLatest()
			break
		}

		for _, tag := range t.all {
			if tag.Name == name {
				t.selected = tag
			}
		}
	}

	return t.selected
}
