package toolGit

import (
	"github.com/newclarity/scribeHelpers/ux"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)


// Usage:
//		{{- $cmd := $git.GetBranch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GetBranch() (string, *ux.State) {
	var ret string
	if state := g.IsNil(); state.IsError() {
		return ret, state
	}
	g.State.SetFunction()

	for range onlyOnce {
		g.State = g.Exec("symbolic-ref", "--short", "HEAD")
		g.State.OutputTrim()
		ret = g.State.Output
		g.State.SetResponse(&ret)
	}
	return ret, g.State
}


// Usage:
//		{{- $cmd := $git.GetBranch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) BranchExists(branch string) (bool, *ux.State) {
	var ret bool
	if state := g.IsNil(); state.IsError() {
		return ret, state
	}
	g.State.SetFunction()

	for range onlyOnce {
		//t := toolTypes.ReflectString(branch)
		//if t == nil {
		//	g.State.SetError("branch is nil")
		//	break
		//}

		g.State = g.Exec("branch", "--list", branch)
		if g.State.IsError() {
			break
		}

		if g.State.Output == branch {
			ret = true
			g.State.SetResponse(&ret)
		}
	}
	return ret, g.State
}


// Usage:
//		{{- $cmd := $git.GetTags }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GetTags() ([]string, *ux.State) {
	var ret []string
	if state := g.IsNil(); state.IsError() {
		return ret, state
	}
	g.State.SetFunction()

	for range onlyOnce {
		// git show-ref --tag
		//
		// 	tagrefs, err := r.Tags()
		//	CheckIfError(err)
		//	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		//		fmt.Println(t)
		//		return nil
		//	})

		g.State.SetSeparator(",")
		//g.State = g.Exec("log", "-1", "--decorate=short", "--pretty=format:%D")
		g.State = g.Exec("tag", "-l")
		if g.State.IsError() {
			break
		}
		g.State.OutputArrayTrim()

		//var tags []string
		//tags = make([]string, 0)
		//for _, t := range g.State.GetOutputArray() {
		//	if t[:5] != " tag:" {
		//		continue
		//	}
		//	tags = append(tags, t[6:])
		//}
		//g.State.SetResponse(tags)
		ret = g.State.GetOutputArray()
		g.State.SetResponse(&ret)
	}

	return ret, g.State
}


// Usage:
//		{{- $cmd := $git.CreateTag "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) CreateTag(tag string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		//t := toolTypes.ReflectString(tag)
		//if t == nil {
		//	g.State.SetError("tag is nil")
		//	break
		//}

		g.State = g.Exec("tag", tag)
		if g.State.IsError() {
			break
		}
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.RemoveTag "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) RemoveTag(tag string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction()

	for range onlyOnce {
		//t := toolTypes.ReflectString(tag)
		//if t == nil {
		//	g.State.SetError("tag is nil")
		//	break
		//}

		g.State = g.Exec("tag", "-d", tag)
		if g.State.IsError() {
			break
		}
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.TagExists "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) TagExists(tag string) (bool, *ux.State) {
	var ret bool
	if state := g.IsNil(); state.IsError() {
		return ret, state
	}
	g.State.SetFunction()

	for range onlyOnce {
		//t := toolTypes.ReflectString(tag)
		//if t == nil {
		//	g.State.SetError("tag is nil")
		//	break
		//}

		g.State = g.Exec("tag", "-l", tag)
		if g.State.IsError() {
			break
		}

		if g.State.Output == tag {
			ret = true
			//noinspection ALL
			g.State.SetResponse(&ret)
		}
	}

	return ret, g.State
}


// Usage:
//		{{- $cmd := $git.GetTagObject "1.0" }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *TypeGit) GetTagObject(tag string) (*object.Tag, *ux.State) {
	var ret *object.Tag
	if state := g.IsNil(); state.IsError() {
		return ret, state
	}
	g.State.SetFunction()

	for range onlyOnce {
		//t := toolTypes.ReflectString(tag)
		//if t == nil {
		//	g.State.SetError("tag is nil")
		//	break
		//}

		//var r *Reference
		//r, g.State.Error = g.repository.Tag(*t)
		r, err := g.repository.Tag(tag)
		g.State.SetError(err)
		if g.State.IsError() {
			break
		}

		//var to *Tag
		//to, g.State.Error = g.repository.TagObject(r.Hash())
		ret, err = g.repository.TagObject(r.Hash())
		g.State.SetError(err)
		if g.State.IsError() {
			break
		}

		g.State.SetResponse(&ret)
	}

	return ret, g.State
}
