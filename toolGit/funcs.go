package toolGit

import (
	"fmt"
	"github.com/gearboxworks/scribeHelpers/ux"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"strings"
)


func GitOpen(path ...string) *TypeGit {
	repo := New(nil)

	for range onlyOnce {
		if len(path) == 0 {
			path = []string{"."}
		}

		repo.State = repo.SetPath(path...)
		if repo.State.IsNotOk() {
			break
		}

		ux.PrintflnBlue("Opening Git repo at path '%s'", repo.Base.GetPath())
		repo.State = repo.Open()
		if repo.State.IsNotOk() {
			break
		}
	}

	return repo
}


func GitClone(url string, path ...string) *TypeGit {
	repo := New(nil)

	for range onlyOnce {
		if len(path) == 0 {
			path = []string{"."}
		}

		repo.State = repo.SetPath(path...)
		if repo.State.IsNotOk() {
			break
		}

		repo.State = repo.SetUrl(url)
		if repo.State.IsNotOk() {
			break
		}

		ux.PrintflnBlue("Cloning Git repo '%s' into path '%s'", repo.Url, repo.Base.GetPath())
		repo.State = repo.Clone()
		if repo.State.IsNotOk() {
			break
		}
	}

	return repo
}


// Usage:
//		{{- $cmd := $git.Clone }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
// func (me *ToolGit) Clone(url interface{}, dir ...interface{}) *TypeExecCommand {
func (g *TypeGit) Remove() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		g.State = g.Base.StatPath()
		if g.State.IsError() {
			break
		}

		// @TODO - TO BE IMPLEMENTED

		//g.Base.Exists
		//ps := toolSystem.ResolveAbsPath(*d)
		//if ps.IsFile {
		//	break
		//}
		//if ps.IsDir {
		//	if ps.Exists {
		//		g.State.SetError("Repository exists for directory '%s'", ps.Filename)
		//		g.Cmd.Exit = 1
		//		break
		//	}
		//}
		//
		//g.SetUrl(*u)
		//g.Base = ps
		//ux.PrintfWhite("Cloning %s into %s\n", g.Url, g.Base.Filename)
		//
		//g.skipDirCheck = true
		//g.Cmd = (*toolTypes.TypeExecCommand)(g.Exec(gitCommandClone, g.Url, g.Base.Filename))
		//g.skipDirCheck = false
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.Lock }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) Lock() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		var ot *object.Tag
		ot, g.State = g.GetTagObject(LockTag)
		if g.State.IsError() {
			break
		}

		//ot := responseToObjectTag(g.State.GetResponse())
		//if ot == nil {
		//	g.State.SetError("Error tags empty.")
		//	break
		//}

		g.State.SetResponse(ot.ID())
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.GetStatus }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GetStatus() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		var wt *git.Worktree
		var err error
		wt, err = g.repository.Worktree()
		g.State.SetError(err)
		if g.State.IsError() {
			break
		}

		var sts git.Status
		sts, err = wt.Status()
		g.State.SetError(err)
		if g.State.IsError() {
			break
		}

		g.State.SetResponse(&sts)
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.Open }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) Pull(opts ...*PullOptions) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		if len(opts) == 0 {
			opts = []*PullOptions{}
		}

		var wt *git.Worktree
		var err error
		wt, err = g.repository.Worktree()
		g.State.SetError(err)
		if g.State.IsError() {
			break
		}

		err = wt.Pull(opts[0])
		g.State.SetError(err)
		if g.State.IsError() {
			break
		}
	}

	return g.State
}


// Usage:
func (g *TypeGit) Commit(paths []string, comment string, args ...interface{}) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		if paths == nil {
			paths = []string{"."}
		}
		if len(paths) == 0 {
			paths = []string{"."}
		}

		c := fmt.Sprintf(comment, args...)
		if c == "" {
			g.State.SetError("Missing comment to Git commit.")
			break
		}

		ux.PrintflnBlue("Adding files to repo...")
		g.State = g.AddFiles(paths...)
		if g.State.IsNotOk() {
			break
		}
		files := g.State.GetResponse().StringArray
		ux.PrintflnBlue("Added %d files to repo.", len(*files))

		g.State = g.ChangedFiles()
		if g.State.IsNotOk() {
			break
		}
		files = g.State.GetResponse().StringArray
		if len(*files) > 0 {
			//if (*files)[0] == "" {
			//	// SIGH... Need to fix the empty strings of GetResponse()
			//	break
			//}
			ux.PrintflnBlue("%d files changed in repo.", len(*files))
			g.State = g.Exec(gitCommandCommit, "-m", c, ".")
			if g.State.IsNotOk() {
				break
			}
		}
	}

	return g.State
}


// Usage:
func (g *TypeGit) Push() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		ux.PrintflnBlue("Pushing to repo '%s'.", g.Url)
		g.State = g.Exec(gitCommandPush, "--porcelain")
		if g.State.IsNotOk() {
			break
		}
	}

	return g.State
}


func (g *TypeGit) AddTag(version string, comment string, args ...interface{}) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		if version == "" {
			g.State.SetError("Missing tag version.")
			break
		}
		if !strings.HasPrefix(version, "v") {
			version = "v" + version
		}

		if g.IsTagExisting(version) {
			g.State.SetOk()
			break
		}


		c := fmt.Sprintf(comment, args...)
		if c == "" {
			c = fmt.Sprintf("Release %s", version)
		}

		ux.PrintflnBlue("Tagging version %s in repo...", version)
		g.State = g.Exec(gitCommandTag, "-a", version, "-m", c)
		if g.State.IsNotOk() {
			break
		}

		ux.PrintflnBlue("Pushing to origin...")
		g.State = g.Exec(gitCommandPush, "origin", version)
		if g.State.IsNotOk() {
			break
		}
	}

	return g.State
}


func (g *TypeGit) DelTag(version string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		if version == "" {
			g.State.SetError("Missing tag version.")
			break
		}
		if !strings.HasPrefix(version, "v") {
			version = "v" + version
		}

		if !g.IsTagExisting(version) {
			g.State.SetOk()
			break
		}


		ux.PrintflnBlue("Removing version tag in repo...")
		g.State = g.Exec(gitCommandTag, "-d", version)
		if g.State.IsNotOk() {
			break
		}

		ux.PrintflnBlue("Pushing to origin...")
		g.State = g.Exec(gitCommandPush, "--delete", "origin", version)
		if g.State.IsNotOk() {
			break
		}

		g.State.SetOk()
	}

	return g.State
}


func (g *TypeGit) IsTagExisting(version string) bool {
	var ok bool
	if state := g.IsNil(); state.IsError() {
		return false
	}
	g.State.SetFunction()

	for range onlyOnce {
		//ux.PrintflnBlue("Checking tag %s in repo...", version)
		g.State = g.Exec(gitCommandTag, "-l", version)
		if g.State.IsNotOk() {
			break
		}

		if !strings.HasPrefix(version, "v") {
			version = "v" + version
		}

		ok = false
		for _, t := range g.State.OutputArray {
			if t == version {
				ok = true
				break
			}
		}
	}

	return ok
}
