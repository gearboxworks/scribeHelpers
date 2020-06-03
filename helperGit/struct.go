package helperGit

import (
	"github.com/tsuyoshiwada/go-gitcmd"
	"github.com/newclarity/scribeHelpers/helperExec"
	"github.com/newclarity/scribeHelpers/helperPath"
	"github.com/newclarity/scribeHelpers/ux"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)


var _ helperExec.TypeExecCommandGetter = (*TypeExecCommand)(nil)
type TypeExecCommand helperExec.TypeExecCommand


type TypeGit struct {
	Url          string
	Base         *helperPath.TypeOsPath

	GitConfig    *gitcmd.Config
	GitOptions   []string

	skipDirCheck bool

	client       gitcmd.Client
	repository   *git.Repository

	Debug        bool
	State        *ux.State
}


func New(debugFlag bool) *TypeGit {
	p := TypeGit{
		Url:          "",
		Base:         helperPath.New(debugFlag),
		GitConfig:    nil,
		GitOptions:   nil,
		skipDirCheck: false,
		client:       nil,
		repository:   nil,

		Debug:        debugFlag,
		State:        ux.NewState(debugFlag),
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
func ReflectHelperGit(g *TypeGit) *HelperGit {
	return (*HelperGit)(g)
}


func (g *TypeGit) IsOk() bool {
	var ok bool
	if state := g.IsNil(); state.IsError() {
		return false
	}

	for range OnlyOnce {
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
		return New(true)
	}
	return g
}


func (g *TypeGit) IsNilRepository() bool {
	ok := true

	for range OnlyOnce {
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

	for range OnlyOnce {
		g.State.SetError(g.client.CanExec())
		if g.State.IsError() {
			g.State.SetError("`git` does not exist or its command file is not executable: %s", g.State.GetError())
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


type (
	Dir          = string
	Url          = string
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
