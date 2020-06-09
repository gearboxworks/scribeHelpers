package toolGhr

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
type Tags struct {
	All []*Tag
	Selected *Tag
}

type Commit struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}


func (t *Tag) String() string {
	return t.Name + " (commit: " + t.Commit.Url + ")"
}


// findTag returns the tag if a tag with name can be found in tags,
// otherwise returns nil.
func (t *Tags) findTag(name string) *Tag {
	for range onlyOnce {
		t.Selected = nil

		if name == "latest" {
			// Latest will always be first... Maybe... @TODO - TO BE CHECKED
			t.Selected = t.All[0]
			break
		}

		for _, tag := range t.All {
			if tag.Name == name {
				t.Selected = tag
			}
		}
	}

	return t.Selected
}
