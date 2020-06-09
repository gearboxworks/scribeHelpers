// This is a "fork and modify" of the codebase from https://github.com/github-release/github-release
// "github-Release" is a stale repo - so the need to fork.
package toolGhr

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type GhrGetter interface {
}

//var _ ux.ResponseGetter = (*TypeResponse)(nil)
//type TypeResponse ux.TypeResponse
//func (t TypeResponse) GetResponse() interface{} {
//	panic("implement me")
//}


type State ux.State
func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}
func ReflectToolGhr(e *TypeGhr) *ToolGhr {
	return (*ToolGhr)(e)
}


type TypeGhr struct {
	Repo     *TypeRepo

	runtime  *toolRuntime.TypeRuntime
	State    *ux.State

	// Download
	//Token    string `goptions:"-s, --security-token, description='Github token ($GITHUB_TOKEN if set). required if repo is private.'"`
	//User     string `goptions:"-u, --user, description='Github repo user or organisation (required if $GITHUB_USER not set)'"`
	//AuthUser string `goptions:"-a, --Auth-user, description='Username for authenticating to the API (falls back to $GITHUB_AUTH_USER or $GITHUB_USER)'"`
	//TypeRepo     string `goptions:"-r, --repo, description='Github repo (required if $GITHUB_REPO not set)'"`
	//Tag      string `goptions:"-t, --tag, description='Git tag to download from (required if latest is not specified)', mutexgroup='input',obligatory"`
	//Name     string `goptions:"-n, --name, description='Name of the file', obligatory"`
	//Latest   bool   `goptions:"-l, --latest, description='Download latest Release (required if tag is not specified)',mutexgroup='input'"`
	//
	// Upload
	//Token    string   `goptions:"-s, --security-token, description='Github token (required if $GITHUB_TOKEN not set)'"`
	//User     string   `goptions:"-u, --user, description='Github repo user or organisation (required if $GITHUB_USER not set)'"`
	//AuthUser string   `goptions:"-a, --Auth-user, description='Username for authenticating to the API (falls back to $GITHUB_AUTH_USER or $GITHUB_USER)'"`
	//TypeRepo     string   `goptions:"-r, --repo, description='Github repo (required if $GITHUB_REPO not set)'"`
	//Tag      string   `goptions:"-t, --tag, description='Git tag to upload to', obligatory"`
	//Name     string   `goptions:"-n, --name, description='Name of the file', obligatory"`
	//Label    string   `goptions:"-l, --label, description='Label (description) of the file'"`
	//TypeFile     *os.TypeFile `goptions:"-f, --file, description='TypeFile to upload (use - for stdin)', rdonly, obligatory"`
	//Replace  bool     `goptions:"-R, --replace, description='Replace asset with same name if it already exists (WARNING: not atomic, failure to upload will remove the original asset too)'"`
	//
	// Release
	//Token      string `goptions:"-s, --security-token, description='Github token (required if $GITHUB_TOKEN not set)'"`
	//User       string `goptions:"-u, --user, description='Github repo user or organisation (required if $GITHUB_USER not set)'"`
	//TypeRepo       string `goptions:"-r, --repo, description='Github repo (required if $GITHUB_REPO not set)'"`
	//Tag        string `goptions:"-t, --tag, obligatory, description='Git tag to create a Release from'"`
	//Name       string `goptions:"-n, --name, description='Name of the Release (defaults to tag)'"`
	//Desc       string `goptions:"-d, --description, description='Release description, use - for reading a description from stdin (defaults to tag)'"`
	//Draft      bool   `goptions:"--draft, description='The Release is a draft'"`
	//Prerelease bool   `goptions:"-p, --pre-Release, description='The Release is a pre-Release'"`
	//Target     string `goptions:"-c, --target, description='Commit SHA or branch to create Release of (defaults to the repository default branch)'"`
	//
	// Edit
	//Token      string `goptions:"-s, --security-token, description='Github token (required if $GITHUB_TOKEN not set)'"`
	//User       string `goptions:"-u, --user, description='Github repo user or organisation (required if $GITHUB_USER not set)'"`
	//AuthUser   string `goptions:"-a, --Auth-user, description='Username for authenticating to the API (falls back to $GITHUB_AUTH_USER or $GITHUB_USER)'"`
	//TypeRepo       string `goptions:"-r, --repo, description='Github repo (required if $GITHUB_REPO not set)'"`
	//Tag        string `goptions:"-t, --tag, obligatory, description='Git tag to edit the Release of'"`
	//Name       string `goptions:"-n, --name, description='New name of the Release (defaults to tag)'"`
	//Desc       string `goptions:"-d, --description, description='New Release description, use - for reading a description from stdin (defaults to tag)'"`
	//Draft      bool   `goptions:"--draft, description='The Release is a draft'"`
	//Prerelease bool   `goptions:"-p, --pre-Release, description='The Release is a pre-Release'"`
	//
	// Delete
	//Token    string `goptions:"-s, --security-token, description='Github token (required if $GITHUB_TOKEN not set)'"`
	//User     string `goptions:"-u, --user, description='Github repo user or organisation (required if $GITHUB_USER not set)'"`
	//AuthUser string `goptions:"-a, --Auth-user, description='Username for authenticating to the API (falls back to $GITHUB_AUTH_USER or $GITHUB_USER)'"`
	//TypeRepo     string `goptions:"-r, --repo, description='Github repo (required if $GITHUB_REPO not set)'"`
	//Tag      string `goptions:"-t, --tag, obligatory, description='Git tag of Release to delete'"`
	//
	// Info
	//Token    string `goptions:"-s, --security-token, description='Github token ($GITHUB_TOKEN if set). required if repo is private.'"`
	//User     string `goptions:"-u, --user, description='Github repo user or organisation (required if $GITHUB_USER not set)'"`
	//AuthUser string `goptions:"-a, --Auth-user, description='Username for authenticating to the API (falls back to $GITHUB_AUTH_USER or $GITHUB_USER)'"`
	//TypeRepo     string `goptions:"-r, --repo, description='Github repo (required if $GITHUB_REPO not set)'"`
	//Tag      string `goptions:"-t, --tag, description='Git tag to query (optional)'"`
	//JSON     bool   `goptions:"-j, --json, description='Emit info as JSON instead of text'"`
}

func New(runtime *toolRuntime.TypeRuntime) *TypeGhr {
	var ghr TypeGhr
	runtime = runtime.EnsureNotNil()

	for range onlyOnce {
		ghr = TypeGhr{
			//Path:     toolPath.New(runtime),

			Repo: NewRepo(runtime),
			//File: NewFile(runtime),

			runtime:  runtime,
			State:    ux.NewState(runtime.CmdName, runtime.Debug),
		}
	}
	ghr.State.SetPackage("")
	ghr.State.SetFunctionCaller()
	return &ghr
}

func (ghr *TypeGhr) IsNil() *ux.State {
	if State := ux.IfNilReturnError(ghr); State.IsError() {
		return State
	}
	ghr.State = ghr.State.EnsureNotNil()
	return ghr.State
}

func (ghr *TypeGhr) isValid() *ux.State {
	if State := ux.IfNilReturnError(ghr); State.IsError() {
		return State
	}

	for range onlyOnce {
		ghr.State = ghr.State.EnsureNotNil()

		ghr.State = ghr.Repo.Auth.isValid()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State = ghr.Repo.isValid()
		if ghr.State.IsNotOk() {
			break
		}

		//ghr.State = ghr.File.isValid()
		//if ghr.State.IsNotOk() {
		//	break
		//}
	}

	return ghr.State
}

func (ghr *TypeGhr) Open(org string, repo string) *ux.State {
	if State := ghr.IsNil(); State.IsError() {
		return State
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.Repo.Open(ghr.Repo.Auth.AuthUser, ghr.Repo.Auth.Token)
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State = ghr.Repo.SetRepo(org, repo)
		if ghr.State.IsNotOk() {
			break
		}

		//ghr.State = ghr.Repo.Releases()
		//if ghr.State.IsNotOk() {
		//	ghr.State.SetError("Cannot connect to repo '%s'", ghr.Repo.GetUrl())
		//	break
		//}
		ghr.State.SetOk("Found %d releases at repo '%s'", ghr.Repo.CountReleases(), ghr.Repo.GetUrl())
	}

	return ghr.State
}

func (ghr *TypeGhr) OpenUrl(repoUrl string) *ux.State {
	if State := ghr.IsNil(); State.IsError() {
		return State
	}
	ghr.State.SetFunction()
	return ghr.Repo.SetUrl(repoUrl)
}

func (ghr *TypeGhr) Set(n TypeRepo) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()
	ghr.State = ghr.Repo.Set(n)
	return ghr.State
}

func (ghr *TypeGhr) SetTag(n string) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()
	ghr.State = ghr.Repo.SetTag(n)
	return ghr.State
}

func (ghr *TypeGhr) SetDescription(n string) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()
	ghr.State = ghr.Repo.SetDescription(n)
	return ghr.State
}

func (ghr *TypeGhr) SetDraft(n bool) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()
	ghr.State = ghr.Repo.SetDraft(n)
	return ghr.State
}

func (ghr *TypeGhr) SetPrerelease(n bool) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()
	ghr.State = ghr.Repo.SetPrerelease(n)
	return ghr.State
}

func (ghr *TypeGhr) SetTarget(n string) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()
	ghr.State = ghr.Repo.SetTarget(n)
	return ghr.State
}
