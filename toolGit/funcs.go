package toolGit

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
)


// Usage:
func (g *TypeGit) Push(comment string, args ...interface{}) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range onlyOnce {
		c := fmt.Sprintf(comment, args...)
		if c == "" {
			g.State.SetError("Missing comment to Git commit.")
			break
		}

		ux.PrintflnBlue("Adding files to repo...")
		g.State = g.AddFiles(".")
		if g.State.IsNotOk() {
			break
		}
		resp := g.State.GetResponse()
		files := resp.GetStringArray()
		ux.PrintflnBlue("Added %d files to repo.", len(*files))


		g.State = g.ChangedFiles()
		if g.State.IsNotOk() {
			break
		}


		resp = g.State.GetResponse()
		files = resp.GetStringArray()
		if len(*files) > 0 {
			ux.PrintflnBlue("Changed %d files in repo.", len(*files))
			g.State = g.Exec(gitCommandCommit, "-m", c, ".")
			if g.State.IsNotOk() {
				break
			}
		}


		//args = append([]string{"--porcelain"}, args...)
		ux.PrintflnBlue("Pushing repo.")
		g.State = g.Exec(gitCommandPush, "--porcelain")
		if g.State.IsNotOk() {
			break
		}

		//var fps Filepaths
		//fps = make(Filepaths, len(g.State.OutputArray))
		//for i, fp := range g.State.OutputArray {
		//	s := strings.Fields(fp)
		//	if len(s) > 1 {
		//		fps[i] = s[1]
		//	}
		//}
		//
		//g.State.SetResponse(&fps)
	}

	return g.State
}


func (g *TypeGit) AddTag(version string, comment string, args ...interface{}) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

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
		g.State = g.Exec("tag", "-a", version, "-m", c)
		if g.State.IsNotOk() {
			break
		}

		ux.PrintflnBlue("Pushing to origin...")
		g.State = g.Exec("push", "origin", version)
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
	g.State.SetFunction("")

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
		g.State = g.Exec("tag", "-d", version)
		if g.State.IsNotOk() {
			break
		}

		ux.PrintflnBlue("Pushing to origin...")
		g.State = g.Exec("push", "--delete", "origin", version)
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
	g.State.SetFunction("")

	for range onlyOnce {
		//ux.PrintflnBlue("Checking tag %s in repo...", version)
		g.State = g.Exec("tag", "-l", version)
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
