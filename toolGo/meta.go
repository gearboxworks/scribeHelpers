package toolGo

import (
	"errors"
	"fmt"
	"github.com/gearboxworks/scribeHelpers/ux"
)


func (g *TypeGo) GetMetaFile(recurse bool, path ...string) *GoFile {
	var ret *GoFile

	for range onlyOnce {
		state := g.State

		if len(path) == 0 {
			path = defaultVersionFile
		}

		//goFiles := New(nil)
		//if goFiles.State.IsError() {
		//	break
		//}

		if recurse {
			state = g.SetRecursive()
		} else {
			state = g.SetNonRecursive()
		}
		if state.IsError() {
			break
		}

		state = g.Find(path...)
		if state.IsError() {
			break
		}

		state = g.Parse()
		if state.IsError() {
			break
		}

		ret = g.Go.Found
	}

	return ret
}


func (g *TypeGo) GetMeta() *GoMeta {
	return g.Go.Found.meta
}


type GoMeta struct {
	binaryName    Name
	binaryVersion Version
	sourceRepo    Repo
	binaryRepo    Repo

	Valid         bool
}


func (gm *GoMeta) setValue(name string, value string) error {
	var err error
	switch name {
		case BinaryName:
			err = gm.binaryName.Set(value)
		case BinaryVersion:
			err = gm.binaryVersion.Set(value)
		case SourceRepo:
			err = gm.sourceRepo.Set(value)
		case BinaryRepo:
			err = gm.binaryRepo.Set(value)
	}
	return err
}


func (gm *GoMeta) IsValid() bool {
	for range onlyOnce {
		gm.Valid = false
		if gm.binaryName.IsNotValid() {
			break
		}
		if gm.binaryVersion.IsNotValid() {
			break
		}
		if gm.sourceRepo.IsNotValid() {
			break
		}
		if gm.binaryRepo.IsNotValid() {
			break
		}
		gm.Valid = true
	}
	return gm.Valid
}


func (gm *GoMeta) IsNotValid() bool {
	return !gm.IsValid()
}


func (gm *GoMeta) Get(name string) (string, error) {
	var value string
	var err error
	switch name {
		case BinaryName:
			value = gm.binaryName.name
		case BinaryVersion:
			value = gm.binaryVersion.version.String()

		case SourceRepo:
			value = gm.sourceRepo.GetUrl()
		case SourceRepoOwner:
			value = gm.sourceRepo.GetOwner()
		case SourceRepoName:
			value = gm.sourceRepo.GetName()

		case BinaryRepo:
			value = gm.binaryRepo.GetUrl()
		case BinaryRepoOwner:
			value = gm.sourceRepo.GetOwner()
		case BinaryRepoName:
			value = gm.sourceRepo.GetName()
	}
	if value == "" {
		err = errors.New(fmt.Sprintf("Cannot find '%s' constant in src files.", name))
	}
	return value, err
}


func (gm *GoMeta) GetBinaryVersion() *Version {
	return &gm.binaryVersion
}


func (gm *GoMeta) GetBinaryName() *Name {
	return &gm.binaryName
}


func (gm *GoMeta) GetSourceRepo() *Repo {
	return &gm.sourceRepo
}


func (gm *GoMeta) GetBinaryRepo() *Repo {
	return &gm.binaryRepo
}


func (gm *GoMeta) Print(name string) {
	for range onlyOnce {
		if !gm.Valid {
			break
		}

		if name == All {
			fmt.Printf("%v", gm)
			break
		}

		value, err := gm.Get(name)
		if err != nil {
			break
		}
		ux.PrintflnBlue("%s: %s", name, value)
	}
}


func (gm *GoMeta) String() string {
	var ret string
	for range onlyOnce {
		ret += gm.binaryName.String()
		ret += gm.binaryVersion.String()
		ret += ux.SprintfBlue("%s\n", SourceRepo) + gm.sourceRepo.String()
		ret += ux.SprintfBlue("%s\n", BinaryRepo) + gm.binaryRepo.String()
	}
	return ret
}
