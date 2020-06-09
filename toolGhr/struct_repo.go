package toolGhr

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/toolGhr/github"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"net/url"
	"os"
	"strings"
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

	client       *github.Client
	url          string
	urlPrefix    string
	apiUrlPrefix string

	Tags         *Tags
	Releases     *Releases
	Release      *Release

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

			client:       &github.Client{},
			urlPrefix:    DefaultGitHubUrl,
			apiUrlPrefix: DefaultGitHubApiUrl,
			url:          os.Getenv("GITHUB_API"),

			Tags:         &Tags{},
			Releases:     &Releases{},
			Release:      &Release{},

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
		if repo.client == nil {
			break
		}

		repo.client.SetBaseURL(repo.apiUrlPrefix)

		//repo.state = repo.GetReleases()
		//if repo.state.IsNotOk() {
		//	break
		//}
		//
		//repo.state = repo.GetTags()
		//if repo.state.IsNotOk() {
		//	break
		//}
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

		repo.state = repo.SetRepo(ur.Organization, ur.Name)
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


func (repo *TypeRepo) SetRepo(org string, name string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if org == "" {
			repo.state.SetError("invalid repo org")
			break
		}
		repo.Organization = org

		if name == "" {
			repo.state.SetError("invalid repo name")
			break
		}
		repo.Name = name

		repo.state = repo.GetReleases()
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.GetTags()
		if repo.state.IsNotOk() {
			break
		}

		repo.state.SetOk()
	}

	return repo.state
}

func (repo *TypeRepo) SetUrl(repoUrl string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if repoUrl == "" {
			repo.state.SetError("invalid repo url")
			break
		}


		u, err := url.Parse(repoUrl)
		if err != nil {
			repo.state.SetError("invalid repo url - %v", err)
			break
		}

		ua := strings.Split(u.Path, "/")
		if len(ua) < 2 {
			repo.state.SetError("invalid repo url - %v", err)
			break
		}

		repo.state = repo.SetRepo(ua[0], ua[1])
		if repo.state.IsNotOk() {
			break
		}

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
		if repo.Tag == "latest" {
			break
		}

		if repo.Tag != "" {
			ret += "/" + repo.Tag
		}
	}

	return ret
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

		repo.state = repo.GetReleases()
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.GetTags()
		if repo.state.IsNotOk() {
			break
		}

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


func (repo *TypeRepo) apiGetDISABLED(url string, args ...interface{}) *ux.State {
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


// Get the releases associated with a repo.
func (repo *TypeRepo) GetReleases() *ux.State {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		u := fmt.Sprintf(ReleaseListUri, repo.Organization, repo.Name)
		err := repo.client.Get(u, &repo.Releases)
		if err != nil {
			repo.state.SetError(err)
			break
		}

		if repo.Releases.findRelease(repo.Tag) == nil {
			repo.state.SetError("no release '%s' found", repo.Tag)
		}

		repo.state.SetOk()
		repo.state.SetResponse(&repo.Releases)
		// Allows the use of the following in a calling function:
		//if repo.state.IsResponseNotOfType("Releases") {
		//	repo.state.SetError("could not get releases")
		//	break
		//}
		//releases := repo.state.GetResponseData().(*Releases)
	}

	return repo.state
}

func (repo *TypeRepo) CountReleases() int {
	if state := repo.IsNil(); state.IsError() {
		return 0
	}
	repo.state.SetFunction()
	return len(repo.Releases.All)
}

func (repo *TypeRepo) PrintReleases() *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	switch {
		case repo.Releases.All == nil:
			repo.message("No repo tags found")

		case repo.Releases.Selected == nil:
			// Print all releases.
			repo.message("Found %d releases.", repo.CountReleases())
			for _, release := range repo.Releases.All {
				fmt.Printf("\n####\n%v", release)
			}

		default:
			// Print selected tag.
			fmt.Printf("\n####\n%v", repo.Releases.Selected.Name)
	}

	return repo.state
}


// Get the tags associated with a repo.
func (repo *TypeRepo) GetTags() *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		u := fmt.Sprintf(TagsUri, repo.Organization, repo.Name)
		err := repo.client.Get(u, &repo.Tags)
		if err != nil {
			repo.state.SetError(err)
			break
		}

		if repo.Tags.findTag(repo.Tag) == nil {
			repo.state.SetError("no tag '%s' found", repo.Tag)
		}

		repo.state.SetOk()
		repo.state.SetResponse(&repo.Tags)
		// Allows the use of the following in a calling function:
		//if repo.state.IsResponseNotOfType("Tags") {
		//	repo.state.SetError("could not get tags")
		//	break
		//}
		//tags := repo.state.GetResponseData().(*Tags)
	}

	return repo.state
}

func (repo *TypeRepo) CountTags() int {
	if state := repo.IsNil(); state.IsError() {
		return 0
	}
	repo.state.SetFunction()
	return len(repo.Tags.All)
}

func (repo *TypeRepo) PrintTags() *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	switch {
		case repo.Tags.All == nil:
			repo.message("No repo tags found")

		case repo.Tags.Selected == nil:
			// Print all tags.
			repo.message("Found %d tags.", repo.CountTags())
			var t []string
			for _, tag := range repo.Tags.All {
				t = append(t, tag.Name)
			}
			repo.message("Repo tags: %s", strings.Join(t, ", "))

		default:
			// Print selected tag.
			repo.message("Repo tag: %s", repo.Tags.Selected.Name)
	}

	return repo.state
}


func (repo *TypeRepo) message(format string, args ...interface{}) {
	ux.PrintfCyan("%s: ", repo.GetUrl())
	ux.PrintflnBlue(format, args...)
}


//// Get an asset associated with a release.
//func (repo *TypeRepo) GetAssets(relId int) *ux.State {
//	if state := repo.IsNil(); state.IsError() {
//		return state
//	}
//	repo.state.SetFunction()
//
//	for range onlyOnce {
//		u := fmt.Sprintf(AssetReleaseListUri, repo.Organization, repo.Name, relId)
//		err := repo.client.Get(u, &repo.Assets)
//		if err != nil {
//			repo.state.SetError(err)
//			break
//		}
//
//		repo.state.SetOk()
//		repo.state.SetResponse(&repo.Assets)
//		// Allows the use of the following in a calling function:
//		//if repo.state.IsResponseNotOfType("Assets") {
//		//	repo.state.SetError("could not replace asset")
//		//	break
//		//}
//		//assets := repo.state.GetResponseData().(*Assets)
//	}
//
//	return repo.state
//}
