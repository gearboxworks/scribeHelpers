package toolGhr

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/toolGhr/github"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
)


type TypeRepo struct {
	Organization string		// `goptions:"-u, --user, description='Github repo user or organisation (required if $GITHUB_USER not set)'"`
	Name         string		// `goptions:"-r, --repo"` // WAS TypeRepo
	Tag          string		// `goptions:"-t, --tag"`
	//Latest       bool		// `goptions:"-l, --latest, description='Download latest release (required if tag is not specified)',mutexgroup='input'"`
	Description  string		// `goptions:"-d, --description, description='Release description, use - for reading a description from stdin (defaults to tag)'"`
	Draft        bool		// `goptions:"--draft, description='The release is a draft'"`
	Prerelease   bool		// `goptions:"-p, --pre-release, description='The release is a pre-release'"`
	Target       string		// `goptions:"-c, --target, description='Commit SHA or branch to create release of (defaults to the repository default branch)'"`

	client       github.Client
	url          string
	urlPrefix    string
	apiUrlPrefix string
	runtime      *toolRuntime.TypeRuntime
	state        *ux.State
}


func NewRepo(runtime *toolRuntime.TypeRuntime) *TypeRepo {
	var repo TypeRepo
	runtime = runtime.EnsureNotNil()

	for range onlyOnce {
		repo = TypeRepo{
			Organization: os.Getenv("GITHUB_USER"),
			Name:         os.Getenv("GITHUB_REPO"),
			Tag:          "latest",
			Description:  "",
			Draft:        false,
			Prerelease:   false,
			Target:       "",

			client:       github.Client{},
			urlPrefix:    DefaultGitHubUrl,
			apiUrlPrefix: DefaultGitHubApiUrl,
			url:          os.Getenv("GITHUB_API"),
			runtime:      runtime,
			state:        ux.NewState(runtime.CmdName, runtime.Debug),
		}
	}
	repo.state.SetPackage("")
	repo.state.SetFunctionCaller()
	return &repo
}


func (repo *TypeRepo) IsNil() *ux.State {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return state
	}
	repo.state = repo.state.EnsureNotNil()
	return repo.state
}


func (repo *TypeRepo) isValid() *ux.State {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return state
	}

	for range onlyOnce {
		repo.state = repo.state.EnsureNotNil()

		if repo.Name == "" {
			repo.state.SetError("repo name is empty")
			break
		}

		if repo.Organization == "" {
			repo.state.SetError("repo org is empty")
			break
		}
	}

	return repo.state
}


func (repo *TypeRepo) isValidTag() *ux.State {
	if State := repo.IsNil(); State.IsError() {
		return State
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if repo.Tag == "" {
			repo.state.SetError("empty tag")
			break
		}

		repo.state.SetOk()
	}

	return repo.state
}


func (repo *TypeRepo) Open(user string, token string) *ux.State {
	if State := repo.IsNil(); State.IsError() {
		return State
	}
	repo.state.SetFunction()

	for range onlyOnce {
		repo.client = github.NewClient(user, token, nil)
		repo.client.SetBaseURL(repo.apiUrlPrefix)

		// 			ghr.State.SetError("Cannot set organization to '%s'", org)
	}

	return repo.state
}


func (repo *TypeRepo) Set(ur TypeRepo) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		repo.state = ur.isValid()
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.SetName(ur.Name)
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.SetTag(ur.Tag)
		if repo.state.IsNotOk() {
			break
		}

		//repo.state = repo.SetDescription(ur.Description)
		//if repo.state.IsNotOk() {
		//	break
		//}

		repo.state = repo.SetDraft(ur.Draft)
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.SetPreRelease(ur.Prerelease)
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.SetTarget(ur.Target)
		if repo.state.IsNotOk() {
			break
		}
	}

	return repo.state
}


func (repo *TypeRepo) SetOrganization(n string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if n == "" {
			repo.state.SetError("invalid repo org")
			break
		}
		repo.Organization = n
		repo.state.SetOk()
	}

	return repo.state
}


func (repo *TypeRepo) SetName(n string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if n == "" {
			repo.state.SetError("invalid repo name")
			break
		}
		repo.Name = n
		repo.state.SetOk()
	}

	return repo.state
}


func (repo *TypeRepo) SetTag(n string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if n == "" {
			repo.state.SetError("invalid repo tag")
			break
		}
		repo.Name = n
		repo.state.SetOk()
	}

	return repo.state
}


func (repo *TypeRepo) SetDescription(n string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if n == "" {
			repo.state.SetError("invalid repo description")
			break
		}
		repo.Description = n
		repo.state.SetOk()
	}

	return repo.state
}


func (repo *TypeRepo) SetDraft(n bool) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		repo.Draft = n
		repo.state.SetOk()
	}

	return repo.state
}


func (repo *TypeRepo) SetPreRelease(n bool) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		repo.Prerelease = n
		repo.state.SetOk()
	}

	return repo.state
}


func (repo *TypeRepo) SetTarget(n string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if n == "" {
			repo.state.SetError("invalid repo target branch")
			break
		}
		repo.Target = n
		repo.state.SetOk()
	}

	return repo.state
}


func (repo *TypeRepo) GetUrl() string {
	var ret string
	if state := repo.IsNil(); state.IsError() {
		return ret
	}
	repo.state.SetFunction()

	for range onlyOnce {
		repo.state = repo.isValid()
		if repo.state.IsNotOk() {
			break
		}

		ret = fmt.Sprintf("%s/%s/%s", repo.urlPrefix, repo.Organization, repo.Name)
		if repo.Tag != "" {
			ret += "/" + repo.Tag
		}
	}

	return ret
}


func (repo *TypeRepo) apiGet(url string, args ...interface{}) *ux.State {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		var ret interface{}

		//u := fmt.Sprintf(github.DefaultBaseURL + url, args...)
		u := fmt.Sprintf(url, args...)
		err := repo.client.Get(u, &ret)
		if err != nil {
			repo.state.SetError(err)
			break
		}

		repo.state.SetOk()
		repo.state.SetResponse(&ret)
	}

	return repo.state
}


func (repo *TypeRepo) GetReleases() (*Releases, *ux.State) {
	var ret Releases
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return &ret, state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		u := fmt.Sprintf(ReleaseListUri, repo.Organization, repo.Name)
		err := repo.client.Get(u, &ret)
		if err != nil {
			repo.state.SetError(err)
			break
		}

		//if repo.state.IsResponseNotOfType("Releases") {
		//	repo.state.SetError("could not get releases")
		//	break
		//}
		//ret = repo.state.GetResponseData().(Releases)

		repo.state.SetOk()
		repo.state.SetResponse(&ret)
	}

	return &ret, repo.state
}


// Get the tags associated with a repo.
func (repo *TypeRepo) GetTags() (*Tags, *ux.State) {
	var tags Tags
	if state := repo.IsNil(); state.IsError() {
		return &Tags{}, state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		u := fmt.Sprintf(TagsUri, repo.Organization, repo.Name)
		err := repo.client.Get(u, &tags)
		if err != nil {
			repo.state.SetError(err)
			break
		}

		//if repo.state.IsResponseNotOfType("[]Tag") {
		//	repo.state.SetError("could not get tags")
		//	break
		//}
		//tags = repo.state.GetResponseData().([]Tag)

		repo.state.SetOk()
		repo.state.SetResponse(&tags)
	}

	return &tags, repo.state
}


// Get an asset associated with a release.
func (repo *TypeRepo) GetAssets(relId int) ([]Asset, *ux.State) {
	var assets []Asset
	if state := repo.IsNil(); state.IsError() {
		return assets, state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		u := fmt.Sprintf(AssetReleaseListUri, repo.Organization, repo.Name, relId)
		err := repo.client.Get(u, &assets)
		if err != nil {
			repo.state.SetError(err)
			break
		}

		//if repo.state.IsResponseNotOfType("[]Asset") {
		//	repo.state.SetError("could not replace asset")
		//	break
		//}
		//assets := repo.state.GetResponseData().([]Asset)

		repo.state.SetOk()
		repo.state.SetResponse(&assets)
	}

	return assets, repo.state
}
