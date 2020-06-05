package toolGitHub

import (
	"context"
	"github.com/google/go-github/v31/github"
	"github.com/newclarity/scribeHelpers/toolTypes"
	"github.com/newclarity/scribeHelpers/ux"
)


type TypeGetRepository struct {
	Data *github.Repository

	Valid bool
	State *ux.State
}

// Usage:
//		{{ $git := GitHubLogin }}
//		{{ $repos := $git.GetRepository "gearboxworks" "docker-template" }}
func (gh *TypeGitHub) GetRepository(owner interface{}, repo interface{}) *TypeGetRepository {
	var ret TypeGetRepository

	for range onlyOnce {
		op := toolTypes.ReflectString(owner)
		if op == nil {
			break
		}

		rp := toolTypes.ReflectString(repo)
		if rp == nil {
			break
		}

		var err error
		ctx := context.Background()
		ret.Data, _, err = gh.Client.Repositories.Get(ctx, *op, *rp)

		ret.State.SetError(err)
		if ret.State.IsError() {
			break
		}

		ret.Valid = true
	}

	return &ret
}

func (me *TypeGetRepository) GetName() toolTypes.TypeGenericStringArray {
	var ret toolTypes.TypeGenericStringArray

	for range onlyOnce {
		ret.Array = append(ret.Array, *me.Data.Name)
		ret.Valid = true
	}

	return ret
}

func (me *TypeGetRepository) GetFullName() toolTypes.TypeGenericStringArray {
	var ret toolTypes.TypeGenericStringArray

	for range onlyOnce {
		ret.Array = append(ret.Array, *me.Data.FullName)
		ret.Valid = true
	}

	return ret
}

func (me *TypeGetRepository) GetUrl() toolTypes.TypeGenericStringArray {
	var ret toolTypes.TypeGenericStringArray

	for range onlyOnce {
		ret.Array = append(ret.Array, *me.Data.URL)
		ret.Valid = true
	}

	return ret
}

// Usage: {{ $branch := GetHeadBranch }}
func (me *TypeGetRepository) GetHeadBranch() toolTypes.TypeGenericString {
	var ret toolTypes.TypeGenericString

	ret.String = me.Data.GetDefaultBranch()

	//for range onlyOnce {
	//	ret.data = me.data.GetDefaultBranch()
	//
	//	branchRefs, ret.Error = me.data.Branches()
	//	if ret.Error != nil {
	//		break
	//	}
	//
	//	headRef, ret.Error = repository.Head()
	//	if ret.Error != nil {
	//		break
	//	}
	//
	//	var currentBranchName string
	//	ret.Error = branchRefs.ForEach(func(branchRef *plumbing.Reference) error {
	//		if branchRef.Hash() == headRef.Hash() {
	//			ret.data = branchRef.Name().output()
	//
	//			return nil
	//		}
	//
	//		return nil
	//	})
	//
	//	if ret.Error != nil {
	//		break
	//	}
	//}

	return ret
}

func (me *TypeGetRepository) GetCurrentCommitFromRepository() toolTypes.TypeGenericString {
	var ret toolTypes.TypeGenericString

	//for range onlyOnce {
	//	headRef, ret.Error = repository.Head()
	//	if ret.Error != nil {
	//		break
	//	}
	//
	//	ret.data = headRef.Hash().output()
	//}

	return ret
}

func (me *TypeGetRepository) GetLatestTagFromRepository() toolTypes.TypeGenericString {
	var ret toolTypes.TypeGenericString

	//for range onlyOnce {
	//	tagRefs, ret.Error = repository.Tags()
	//	if ret.Error != nil {
	//		break
	//	}
	//
	//	var latestTagCommit *object.Commit
	//	err = tagRefs.ForEach(func(tagRef *plumbing.Reference) error {
	//		revision := plumbing.Revision(tagRef.Name().output())
	//		tagCommitHash, ret.Error = repository.ResolveRevision(revision)
	//		if ret.Error != nil {
	//			return err
	//		}
	//
	//		commit, ret.Error = repository.CommitObject(*tagCommitHash)
	//		if ret.Error != nil {
	//			return err
	//		}
	//
	//		if latestTagCommit == nil {
	//			latestTagCommit = commit
	//			ret.data = tagRef.Name().output()
	//		}
	//
	//		if commit.Committer.When.After(latestTagCommit.Committer.When) {
	//			latestTagCommit = commit
	//			ret.data = tagRef.Name().output()
	//		}
	//
	//		return nil
	//	})
	//
	//	if ret.Error != nil {
	//		break
	//	}
	//}

	return ret
}


//////////////////////////////////////////////////////////////////////

type TypeGetRepositories struct {
	Valid bool
	Error error
	Data []*github.Repository
}


// Usage:
//		{{ $git := GitHubLogin }}
//		{{ $repos := $git.GetRepositories "gearboxworks" }}
func (gh *TypeGitHub) GetRepositories(owner interface{}) *TypeGetRepositories {
	var ret TypeGetRepositories

	for range onlyOnce {
		op := toolTypes.ReflectString(owner)
		if op == nil {
			break
		}

		ctx := context.Background()
		ret.Data, _, ret.Error = gh.Client.Repositories.List(ctx, *op, nil)

		if ret.Error != nil {
			break
		}

		//fmt.Printf("%+v\n", pack)
		ret.Valid = true
	}

	return &ret
}

// Usage:
//		{{ $git := GitHubLogin }}
//		{{ $repos := $git.GetRepositories "gearboxworks" }}
//		{{ $names := $repos.GetNames }}
func (me *TypeGetRepositories) GetNames() toolTypes.TypeGenericStringArray {
	var ret toolTypes.TypeGenericStringArray

	for range onlyOnce {
		for _, v := range me.Data {
			ret.Array = append(ret.Array, *v.Name)
		}
		ret.Valid = true
	}

	return ret
}

// Usage:
//		{{ $git := GitHubLogin }}
//		{{ $repos := $git.GetRepositories "gearboxworks" }}
//		{{ $names := $repos.GetFullNames }}
func (me *TypeGetRepositories) GetFullNames() toolTypes.TypeGenericStringArray {
	var ret toolTypes.TypeGenericStringArray

	for range onlyOnce {
		for _, v := range me.Data {
			ret.Array = append(ret.Array, *v.FullName)
		}
		ret.Valid = true
	}

	return ret
}

// Usage:
//		{{ $git := GitHubLogin }}
//		{{ $repos := $git.GetRepositories "gearboxworks" }}
//		{{ $urls := $repos.GetUrls }}
func (me *TypeGetRepositories) GetUrls() toolTypes.TypeGenericStringArray {
	var ret toolTypes.TypeGenericStringArray

	for range onlyOnce {
		for _, v := range me.Data {
			ret.Array = append(ret.Array, *v.URL)
		}
		ret.Valid = true
	}

	return ret
}
