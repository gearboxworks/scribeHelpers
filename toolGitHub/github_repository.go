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

	for range OnlyOnce {
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

	for range OnlyOnce {
		ret.Array = append(ret.Array, *me.Data.Name)
		ret.Valid = true
	}

	return ret
}

func (me *TypeGetRepository) GetFullName() toolTypes.TypeGenericStringArray {
	var ret toolTypes.TypeGenericStringArray

	for range OnlyOnce {
		ret.Array = append(ret.Array, *me.Data.FullName)
		ret.Valid = true
	}

	return ret
}

func (me *TypeGetRepository) GetUrl() toolTypes.TypeGenericStringArray {
	var ret toolTypes.TypeGenericStringArray

	for range OnlyOnce {
		ret.Array = append(ret.Array, *me.Data.URL)
		ret.Valid = true
	}

	return ret
}

// Usage: {{ $branch := GetHeadBranch }}
func (me *TypeGetRepository) GetHeadBranch() toolTypes.TypeGenericString {
	var ret toolTypes.TypeGenericString

	ret.String = me.Data.GetDefaultBranch()

	//for range OnlyOnce {
	//	ret.Data = me.Data.GetDefaultBranch()
	//
	//	branchRefs, ret.Error = me.Data.Branches()
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
	//			ret.Data = branchRef.Name().output()
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

	//for range OnlyOnce {
	//	headRef, ret.Error = repository.Head()
	//	if ret.Error != nil {
	//		break
	//	}
	//
	//	ret.Data = headRef.Hash().output()
	//}

	return ret
}

func (me *TypeGetRepository) GetLatestTagFromRepository() toolTypes.TypeGenericString {
	var ret toolTypes.TypeGenericString

	//for range OnlyOnce {
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
	//			ret.Data = tagRef.Name().output()
	//		}
	//
	//		if commit.Committer.When.After(latestTagCommit.Committer.When) {
	//			latestTagCommit = commit
	//			ret.Data = tagRef.Name().output()
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

	for range OnlyOnce {
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

	for range OnlyOnce {
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

	for range OnlyOnce {
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

	for range OnlyOnce {
		for _, v := range me.Data {
			ret.Array = append(ret.Array, *v.URL)
		}
		ret.Valid = true
	}

	return ret
}
