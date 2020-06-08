// This is a "fork and modify" of the codebase from https://github.com/github-release/github-release
// "github-release" is a stale repo - so the need to fork.
package toolGhr

import (
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type GhrGetter interface {
}


type State ux.State
func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}
func ReflectToolGhr(e *TypeGhr) *ToolGhr {
	return (*ToolGhr)(e)
}


type TypeGhr struct {
	Path      *toolPath.TypeOsPath

	Auth     *TypeAuth
	Repo     *TypeRepo
	File     *TypeFile

	urlPrefix string
	//gh    github.Client
	release   Release
	releases  []Release

	runtime  *toolRuntime.TypeRuntime
	State    *ux.State

	// Download
	//Token    string `goptions:"-s, --security-token, description='Github token ($GITHUB_TOKEN if set). required if repo is private.'"`
	//User     string `goptions:"-u, --user, description='Github repo user or organisation (required if $GITHUB_USER not set)'"`
	//AuthUser string `goptions:"-a, --auth-user, description='Username for authenticating to the API (falls back to $GITHUB_AUTH_USER or $GITHUB_USER)'"`
	//TypeRepo     string `goptions:"-r, --repo, description='Github repo (required if $GITHUB_REPO not set)'"`
	//Tag      string `goptions:"-t, --tag, description='Git tag to download from (required if latest is not specified)', mutexgroup='input',obligatory"`
	//Name     string `goptions:"-n, --name, description='Name of the file', obligatory"`
	//Latest   bool   `goptions:"-l, --latest, description='Download latest release (required if tag is not specified)',mutexgroup='input'"`
	//
	// Upload
	//Token    string   `goptions:"-s, --security-token, description='Github token (required if $GITHUB_TOKEN not set)'"`
	//User     string   `goptions:"-u, --user, description='Github repo user or organisation (required if $GITHUB_USER not set)'"`
	//AuthUser string   `goptions:"-a, --auth-user, description='Username for authenticating to the API (falls back to $GITHUB_AUTH_USER or $GITHUB_USER)'"`
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
	//Tag        string `goptions:"-t, --tag, obligatory, description='Git tag to create a release from'"`
	//Name       string `goptions:"-n, --name, description='Name of the release (defaults to tag)'"`
	//Desc       string `goptions:"-d, --description, description='Release description, use - for reading a description from stdin (defaults to tag)'"`
	//Draft      bool   `goptions:"--draft, description='The release is a draft'"`
	//Prerelease bool   `goptions:"-p, --pre-release, description='The release is a pre-release'"`
	//Target     string `goptions:"-c, --target, description='Commit SHA or branch to create release of (defaults to the repository default branch)'"`
	//
	// Edit
	//Token      string `goptions:"-s, --security-token, description='Github token (required if $GITHUB_TOKEN not set)'"`
	//User       string `goptions:"-u, --user, description='Github repo user or organisation (required if $GITHUB_USER not set)'"`
	//AuthUser   string `goptions:"-a, --auth-user, description='Username for authenticating to the API (falls back to $GITHUB_AUTH_USER or $GITHUB_USER)'"`
	//TypeRepo       string `goptions:"-r, --repo, description='Github repo (required if $GITHUB_REPO not set)'"`
	//Tag        string `goptions:"-t, --tag, obligatory, description='Git tag to edit the release of'"`
	//Name       string `goptions:"-n, --name, description='New name of the release (defaults to tag)'"`
	//Desc       string `goptions:"-d, --description, description='New release description, use - for reading a description from stdin (defaults to tag)'"`
	//Draft      bool   `goptions:"--draft, description='The release is a draft'"`
	//Prerelease bool   `goptions:"-p, --pre-release, description='The release is a pre-release'"`
	//
	// Delete
	//Token    string `goptions:"-s, --security-token, description='Github token (required if $GITHUB_TOKEN not set)'"`
	//User     string `goptions:"-u, --user, description='Github repo user or organisation (required if $GITHUB_USER not set)'"`
	//AuthUser string `goptions:"-a, --auth-user, description='Username for authenticating to the API (falls back to $GITHUB_AUTH_USER or $GITHUB_USER)'"`
	//TypeRepo     string `goptions:"-r, --repo, description='Github repo (required if $GITHUB_REPO not set)'"`
	//Tag      string `goptions:"-t, --tag, obligatory, description='Git tag of release to delete'"`
	//
	// Info
	//Token    string `goptions:"-s, --security-token, description='Github token ($GITHUB_TOKEN if set). required if repo is private.'"`
	//User     string `goptions:"-u, --user, description='Github repo user or organisation (required if $GITHUB_USER not set)'"`
	//AuthUser string `goptions:"-a, --auth-user, description='Username for authenticating to the API (falls back to $GITHUB_AUTH_USER or $GITHUB_USER)'"`
	//TypeRepo     string `goptions:"-r, --repo, description='Github repo (required if $GITHUB_REPO not set)'"`
	//Tag      string `goptions:"-t, --tag, description='Git tag to query (optional)'"`
	//JSON     bool   `goptions:"-j, --json, description='Emit info as JSON instead of text'"`
}


func New(runtime *toolRuntime.TypeRuntime) *TypeGhr {
	var ghr TypeGhr
	runtime = runtime.EnsureNotNil()

	for range onlyOnce {
		ghr = TypeGhr{
			Path:     toolPath.New(runtime),

			Auth: NewAuth(runtime),
			Repo: NewRepo(runtime),
			File: NewFile(runtime),

			urlPrefix: DefaultGitHubUrl,
			//gh:    github.Client{},

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

		ghr.State = ghr.Auth.isValid()
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


//func (ghr *TypeGhr) isValidTag(latest bool) *ux.State {
//	if State := ghr.IsNil(); State.IsError() {
//		return State
//	}
//	ghr.State.SetFunction()
//
//	for range onlyOnce {
//		ghr.State = ghr.Repo.isValid()
//		if ghr.State.IsNotOk() {
//			break
//		}
//
//		if ghr.Repo.Tag == "" && !latest {
//			ghr.State.SetError("empty tag")
//			break
//		}
//
//		ghr.State.SetOk()
//	}
//
//	return ghr.State
//}


//func (ghr *TypeGhr) validateCredentials() *ux.State {
//	if State := ghr.IsNil(); State.IsError() {
//		return State
//	}
//	ghr.State.SetFunction()
//
//	for range onlyOnce {
//		ghr.State = ghr.Repo.isValidTag(false)
//		if ghr.State.IsNotOk() {
//			break
//		}
//
//		if ghr.Auth.Token == "" {
//			ghr.State.SetError("empty token")
//			break
//		}
//
//		ghr.State.SetOk()
//	}
//
//	// return nil
//	return ghr.State
//}


func (ghr *TypeGhr) Open(org string, repo string) *ux.State {
	if State := ghr.IsNil(); State.IsError() {
		return State
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		ghr.State = ghr.Repo.Open(ghr.Auth.AuthUser, ghr.Auth.Token)
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State = ghr.Repo.SetOrganization(org)
		if ghr.State.IsNotOk() {
			ghr.State.SetError("Cannot set organization to '%s'", org)
			break
		}

		ghr.State = ghr.Repo.SetName(repo)
		if ghr.State.IsNotOk() {
			ghr.State.SetError("Cannot set repo to '%s'", repo)
			break
		}

		var rels *Releases
		rels, ghr.State = ghr.Repo.GetReleases()
		if ghr.State.IsNotOk() {
			ghr.State.SetError("Cannot connect to repo '%s'", ghr.Repo.GetUrl())
			break
		}
		ghr.State.SetOk("Found %d releases at repo '%s'", len(*rels), ghr.Repo.GetUrl())
	}

	return ghr.State
}


func (ghr *TypeGhr) setAuth(a TypeAuth) *ux.State {
	if State := ghr.IsNil(); State.IsError() {
		return State
	}
	ghr.State.SetFunction()
	ghr.State = ghr.Auth.Set(a)
	return ghr.State
}


func (ghr *TypeGhr) setRepo(r TypeRepo) *ux.State {
	if State := ghr.IsNil(); State.IsError() {
		return State
	}
	ghr.State.SetFunction()
	ghr.State = ghr.Repo.Set(r)
	return ghr.State
}


func (ghr *TypeGhr) setFile(f TypeFile) *ux.State {
	if State := ghr.IsNil(); State.IsError() {
		return State
	}
	ghr.State.SetFunction()
	ghr.State = ghr.File.Set(f)
	return ghr.State
}


func (ghr *TypeGhr) GetReleases() *ux.State {
	if State := ghr.IsNil(); State.IsError() {
		return State
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		var rels *Releases
		rels, ghr.State = ghr.Repo.GetReleases()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State.SetOk("Found %d releases at repo '%s'", len(*rels), ghr.Repo.GetUrl())
		ghr.State.SetResponse(&rels)
	}

	return ghr.State
}


func (ghr *TypeGhr) GetTags() *ux.State {
	if State := ghr.IsNil(); State.IsError() {
		return State
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		var tags *Tags
		tags, ghr.State = ghr.Repo.GetTags()
		if ghr.State.IsNotOk() {
			break
		}

		ghr.State.SetOk("Found %d releases at repo '%s'", len(*tags), ghr.Repo.GetUrl())
		ghr.State.SetResponse(&tags)
	}

	return ghr.State
}


func (ghr *TypeGhr) SetTag(n string) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()
	ghr.State = ghr.Repo.SetName(n)
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


func (ghr *TypeGhr) SetPreRelease(n bool) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()
	ghr.State = ghr.Repo.SetPreRelease(n)
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
