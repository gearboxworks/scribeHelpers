package toolGhr

import (
	"fmt"
	"github.com/gearboxworks/scribeHelpers/toolGhr/github"
	"github.com/gearboxworks/scribeHelpers/toolPath"
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
	"net/url"
	"os"
	"strings"
)


type TypeRepo struct {
	Organization string
	Name         string
	TagName      string
	Description  string
	Draft        bool
	Prerelease   bool
	Target       string
	Files        []string
	Overwrite    bool

	Auth         *TypeAuth
	client       *github.Client
	clientValid  bool
	urlPrefix    string
	apiUrlPrefix string

	tags         *tags
	releases     *releases
	assets       *assets
	//files        *TypeFiles

	runtime      *toolRuntime.TypeRuntime
	state        *ux.State
}
func (repo *TypeRepo) IsNil() *ux.State {
	return ux.IfNilReturnError(repo)
}


//type TypeCreateRel struct {
//	Organization string // `goptions:"-u, --user, description='Github repo user or organisation (required if $GITHUB_USER not set)'"`
//	Name         string // `goptions:"-r, --repo"` // WAS TypeRepo
//	TagName      string // `goptions:"-t, --tag"`
//	Description  string // `goptions:"-d, --description, description='Release description, use - for reading a description from stdin (defaults to tag)'"`
//	Draft        bool   // `goptions:"--draft, description='The Release is a draft'"`
//	Prerelease   bool   // `goptions:"-p, --pre-Release, description='The Release is a pre-Release'"`
//	Target       string // `goptions:"-c, --target, description='Commit SHA or branch to create Release of (defaults to the repository default branch)'"`
//}


// Instantiate methods
func NewRepo(runtime *toolRuntime.TypeRuntime) *TypeRepo {
	var repo TypeRepo
	runtime = runtime.EnsureNotNil()

	for range onlyOnce {
		repo = TypeRepo{
			Organization: os.Getenv("GITHUB_USER"),
			Name:         os.Getenv("GITHUB_REPO"),
			TagName:      Latest,
			Description:  "",
			Draft:        false,
			Prerelease:   false,
			Target:       "",
			Files:        []string{},
			Overwrite:      false,

			Auth:         NewAuth(runtime),
			client:       &github.Client{},
			urlPrefix:    DefaultGitHubUrl,
			apiUrlPrefix: os.Getenv("GITHUB_API"),

			tags:         &tags{},
			releases:     &releases{},
			assets:       &assets{},

			runtime:      runtime,
			state:        ux.NewState(runtime.CmdName, runtime.Debug),
		}
	}
	repo.state.SetPackage("")
	repo.state.SetFunctionCaller()

	if repo.apiUrlPrefix == "" {
		repo.apiUrlPrefix = DefaultGitHubApiUrl
	}

	return &repo
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

		repo.state = repo.Auth.IsNil()
		if repo.state.IsNotOk() {
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
		if repo.TagName == "" {
			repo.state.SetError("empty tag")
			break
		}

		repo.state.SetOk()
	}

	return repo.state
}

func (repo *TypeRepo) isAuthChanged(user string, token string) bool {
	var ok bool
	for range onlyOnce {
		if !repo.clientValid {
			ok = true
			break
		}
		if repo.Auth.AuthUser != user {
			ok = true
			break
		}
		if repo.Auth.Token != token {
			ok = true
			break
		}
		ok = false
	}
	return ok
}


// Create/open/set methods
func (repo *TypeRepo) Open(user string, token string) *ux.State {
	if State := repo.IsNil(); State.IsError() {
		return State
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if !repo.isAuthChanged(user, token) {
			break
		}

		if user != "" {
			repo.Auth.AuthUser = user
		}

		if token != "" {
			repo.Auth.Token = token
		}

		repo.client = github.NewClient(repo.Auth.AuthUser, repo.Auth.Token, nil)
		if repo.client == nil {
			break
		}

		repo.client.SetBaseURL(repo.apiUrlPrefix)

		repo.clientValid = true

		//repo.state = repo.FetchReleases()
		//if repo.state.IsNotOk() {
		//	break
		//}
		//
		//repo.state = repo.FetchTags()
		//if repo.state.IsNotOk() {
		//	break
		//}
	}

	return repo.state
}

func (repo *TypeRepo) Set(ur *TypeRepo) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if ur.Auth != nil {
			repo.state = repo.Open(ur.Auth.AuthUser, ur.Auth.Token)
			if repo.state.IsNotOk() {
				break
			}
		}

		repo.state = repo.SetDescription(ur.Description)
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.SetDraft(ur.Draft)
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.SetPrerelease(ur.Prerelease)
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.SetTarget(ur.Target)
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.SetRepo(ur.Organization, ur.Name)
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.SetTag(ur.TagName)
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.SetOverwrite(ur.Overwrite)
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.SetFiles(ur.Files...)
		if repo.state.IsNotOk() {
			break
		}

		repo.state = repo.isValid()
		if repo.state.IsNotOk() {
			break
		}
	}

	return repo.state
}

func (repo *TypeRepo) SetAuth(ur *TypeAuth) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if ur == nil {
			break
		}

		repo.state = repo.Open(ur.AuthUser, ur.Token)
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
		repo.state = repo.Open(repo.Auth.AuthUser, repo.Auth.Token)
		if repo.state.IsNotOk() {
			break
		}

		if org != "" {
			repo.Organization = org
		}
		if name != "" {
			repo.Name = name
		}

		repo.state = repo.Fetch(true)
		if repo.state.IsError() {
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
		repo.state = repo.Open(repo.Auth.AuthUser, repo.Auth.Token)
		if repo.state.IsNotOk() {
			break
		}

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
		switch {
			case len(ua) < 2:
				repo.state.SetError("invalid repo url - %v", err)
				break

			case len(ua) > 2:
				ua = ua[1:3]
				fallthrough
			default:
				repo.state = repo.SetRepo(ua[0], ua[1])
				if repo.state.IsNotOk() {
					break
				}
				repo.state.SetOk()
		}
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
		if repo.TagName == Latest {
			break
		}

		if repo.TagName != "" {
			ret += "/" + repo.TagName
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
		repo.TagName = n

		repo.state = repo.Fetch(true)
		if repo.state.IsError() {
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
			//repo.state.SetError("invalid repo description")
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
	repo.Draft = n
	repo.state.SetOk()
	return repo.state
}

func (repo *TypeRepo) SetPrerelease(n bool) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()
	repo.Prerelease = n
	repo.state.SetOk()
	return repo.state
}

func (repo *TypeRepo) SetOverwrite(n bool) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()
	repo.Overwrite = n
	repo.state.SetOk()
	return repo.state
}

func (repo *TypeRepo) SetTarget(n string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if n == "" {
			//repo.state.SetError("invalid repo target branch")
			break
		}
		repo.Target = n
		repo.state.SetOk()
	}

	return repo.state
}

func (repo *TypeRepo) SetFiles(f ...string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if len(f) == 0 {
			break
		}
		repo.Files = f
		repo.state.SetOk()
	}

	return repo.state
}

func (repo *TypeRepo) SetFilePath(glob string, f ...string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		if len(f) == 0 {
			break
		}

		fp := toolPath.NewPaths(repo.runtime)
		if fp.State.IsNotOk() {
			repo.state = fp.State
			break
		}

		repo.state = fp.FindRegex(glob, f...)
		if repo.state.IsNotOk() {
			break
		}

		if fp.GetLength() == 0 {
			repo.state.SetError("No files with pattern '%s' found in path '%s'", glob, fp.Base.GetPath())
			break
		}

		repo.Files = []string{}
		for _, file := range fp.Paths {
			repo.Files = append(repo.Files, file.GetPathAbs())
		}

		repo.state.SetOk()
	}

	return repo.state
}


// Get the releases associated with a repo.
func (repo *TypeRepo) Fetch(force bool) *ux.State {
	if state := ux.IfNilReturnError(repo); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		repo.state = repo.isValidTag()
		if repo.state.IsError() {
			break
		}

		repo.state = repo.FetchTags(force)
		if repo.state.IsError() {
			break
		}

		repo.state = repo.FetchReleases(force)
		if repo.state.IsError() {
			break
		}

		repo.state = repo.FetchAssets(force)
		if repo.state.IsError() {
			break
		}
	}

	return repo.state
}


func (repo *TypeRepo) message(format string, args ...interface{}) {
	ux.PrintfCyan("%s: ", repo.GetUrl())
	ux.PrintflnBlue(format, args...)
}


func (repo *TypeRepo) messageOk(format string, args ...interface{}) {
	ux.PrintfCyan("%s: ", repo.GetUrl())
	ux.PrintflnGreen(format, args...)
}


func (repo *TypeRepo) messageError(format string, args ...interface{}) {
	ux.PrintfCyan("%s: ", repo.GetUrl())
	ux.PrintflnRed(format, args...)
}
