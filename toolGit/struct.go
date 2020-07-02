package toolGit

import (
	"github.com/newclarity/scribeHelpers/toolExec"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/tsuyoshiwada/go-gitcmd"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"net/url"
)


var _ toolExec.TypeExecCommandGetter = (*TypeExecCommand)(nil)
type TypeExecCommand toolExec.TypeExecCommand


type TypeGit struct {
	Url          *url.URL
	Base         *toolPath.TypeOsPath

	GitConfig    *gitcmd.Config
	GitOptions   []string

	skipDirCheck bool

	client       gitcmd.Client
	repository   *git.Repository

	runtime      *toolRuntime.TypeRuntime
	State        *ux.State
}


func New(runtime *toolRuntime.TypeRuntime) *TypeGit {
	runtime = runtime.EnsureNotNil()

	p := TypeGit{
		Url:          &url.URL{},
		Base:         toolPath.New(runtime),
		GitConfig:    nil,
		GitOptions:   nil,
		skipDirCheck: false,
		client:       gitcmd.New(&gitcmd.Config{}),
		repository:   nil,

		runtime:      runtime,
		State:        ux.NewState(runtime.CmdName, runtime.Debug),
	}
	p.State.SetPackage("")
	p.State.SetFunctionCaller()
	return &p
}


type State ux.State
func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}
//func ReflectState(p *ux.State) *ux.State {
//	return (*State)(p)
//}
func ReflectToolGit(g *TypeGit) *ToolGit {
	return (*ToolGit)(g)
}


func (g *TypeGit) IsOk() bool {
	var ok bool
	if state := g.IsNil(); state.IsError() {
		return false
	}

	for range onlyOnce {
		if !g.IsAvailable() {
			break
		}
		if g.IsNilRepository() {
			break
		}
		g.State.Clear()
		ok = true
	}

	return ok
}
func (g *TypeGit) IsNotOk() bool {
	return !g.IsOk()
}


func (g *TypeGit) IsNil() *ux.State {
	if state := ux.IfNilReturnError(g); state.IsError() {
		return state
	}
	g.State = g.State.EnsureNotNil()
	return g.State
}


func (g *TypeGit) EnsureNotNil() *TypeGit {
	if g == nil {
		return New(nil)
	}
	return g
}


func (g *TypeGit) IsNilRepository() bool {
	ok := true

	for range onlyOnce {
		if g.repository == nil {
			g.State.SetError("repository not open")
			break
		}
		g.State.Clear()
		ok = false
	}

	return ok
}


func (g *TypeGit) IsAvailable() bool {
	ok := false

	for range onlyOnce {
		err := g.client.CanExec()
		if err != nil {
			g.State.SetError("`git` does not exist or its command file is not executable: %s", err)
			break
		}
		g.State.Clear()
		ok = true
	}

	return ok
}
func (g *TypeGit) IsNotAvailable() bool {
	return !g.IsAvailable()
}


func (g *TypeGit) SetConfig(config gitcmd.Config) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		g.GitConfig = &config
		g.client = gitcmd.New(&config)
	}

	return g.State
}


func responseToObjectTag(r *ux.TypeResponse) *object.Tag {
	var o *object.Tag

	for range onlyOnce {
		if !r.IsOfType("object.Tag") {
			break
		}
		o = r.Pointer().(*object.Tag)
	}

	return o
}


type (
	Dir          = string
	Url          = *url.URL
	Filepath     = string
	Filepaths    []Filepath
	ReadableName = string
	Tagname      = string
)

type (
	PullOptions  = git.PullOptions
	LogOptions   = git.LogOptions
	Tag          = object.Tag
	Reference    = plumbing.Reference
	CloneOptions = git.CloneOptions
	Status       = git.Status
)
