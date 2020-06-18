package toolGo

import (
	"errors"
	"fmt"
	"github.com/blang/semver"
	"github.com/newclarity/scribeHelpers/ux"
	"go/ast"
	"net/url"
	"path/filepath"
	"strings"
)


const (
	All           = "All"

	BinaryName    = "BinaryName"
	BinaryVersion = "BinaryVersion"

	SourceRepo    = "SourceRepo"
	SourceRepoOwner    = "SourceRepoOwner"
	SourceRepoName    = "SourceRepoName"

	BinaryRepo    = "BinaryRepo"
	BinaryRepoOwner    = "BinaryRepoOwner"
	BinaryRepoName    = "BinaryRepoName"
)
var defaultVersionFile = []string{"defaults", "version.go"}


func (g *TypeGo) GetMeta(recurse bool, path ...string) *GoMeta {
	var ret *GoMeta

	for range onlyOnce {
		if len(path) == 0 {
			path = defaultVersionFile
		}

		//goFiles := New(nil)
		//if goFiles.State.IsError() {
		//	break
		//}
		//
		//goFiles.SetNonRecursive()
		//if goFiles.State.IsError() {
		//	break
		//}

		g.Find(path...)
		if g.State.IsError() {
			break
		}

		g.Parse()
		if g.State.IsError() {
			break
		}

		ret = g.Files.GetMeta()
	}

	return ret
}


func GetDefaultFile() string {
	return filepath.Join(defaultVersionFile...)
}


type GoMeta struct {
	binaryName    Name
	binaryVersion Version
	sourceRepo    Repo
	binaryRepo    Repo

	Valid         bool
	//State         *ux.State
}


type Repo struct {
	url *url.URL
}
func (r *Repo) GetOwner() string {
	value, _ := r.Get()
	return value
}
func (r *Repo) GetName() string {
	_, value := r.Get()
	return value
}
func (r *Repo) Get() (string, string) {
	var owner string
	var name string
	for range onlyOnce {
		pa := strings.Split(r.url.Path, "/")
		switch len(pa) {
			case 0:
			case 1:
			case 2:
				owner = pa[1]
				name = ""

			default:
				owner = pa[1]
				name = pa[2]
		}
	}
	return owner, name
}
func (r *Repo) GetUrl() string {
	return r.url.String()
}
func (r *Repo) String() string {
	var ret string
	for range onlyOnce {
		ret += ux.SprintfBlue("Repo URL: ")   + ux.SprintfCyan("%v\n", r.url)
		ret += ux.SprintfBlue("Repo owner: ") + ux.SprintfCyan("%s\n", r.GetOwner())
		ret += ux.SprintfBlue("Repo name: ")  + ux.SprintfCyan("%s\n", r.GetName())
	}
	return ret
}
func (r *Repo) IsValid() bool {
	var ok bool
	for range onlyOnce {
		if r == nil {
			break
		}
		if r.url == nil {
			break
		}
		ok = true
	}
	return ok
}
func (r *Repo) IsNotValid() bool {
	return !r.IsValid()
}


type Version struct {
	version *semver.Version
}
func (v *Version) String() string {
	var ret string
	for range onlyOnce {
		ret += ux.SprintfBlue("%s: ", BinaryVersion)
		ret += ux.SprintfCyan("%v\n", v.version)
	}
	return ret
}
func (v *Version) Get() Version {
	return *v
}
func (v *Version) IsValid() bool {
	var ok bool
	for range onlyOnce {
		if v == nil {
			break
		}
		if v.version == nil {
			break
		}
		if err := v.version.Validate(); err != nil {
			break
		}
		ok = true
	}
	return ok
}
func (v *Version) IsNotValid() bool {
	return !v.IsValid()
}


type Name struct {
	name string
}
func (n *Name) String() string {
	var ret string
	for range onlyOnce {
		ret += ux.SprintfBlue("%s: ", BinaryName)
		ret += ux.SprintfCyan("%v\n", n.name)
	}
	return ret
}
func (n *Name) Get() string {
	return n.name
}
func (n *Name) IsValid() bool {
	var ok bool
	for range onlyOnce {
		if n == nil {
			break
		}
		if n.name == "" {
			break
		}
		ok = true
	}
	return ok
}
func (n *Name) IsNotValid() bool {
	return !n.IsValid()
}


func (gm *GoMeta) Set(name string, value string) bool {
	var ok bool
	switch name {
		case BinaryName:
			gm.binaryName.name = value
			ok = true
		case BinaryVersion:
			v, err := semver.Parse(value)
			if err == nil {
				gm.binaryVersion.version = &v
				ok = true
			}
		case SourceRepo:
			v, err := url.Parse(addPrefix(value))
			if err == nil {
				gm.sourceRepo.url = v
				ok = true
			}
		case BinaryRepo:
			v, err := url.Parse(addPrefix(value))
			if err == nil {
				gm.binaryRepo.url = v
				ok = true
			}
	}
	return ok
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


func addPrefix(u string) string {
	for range onlyOnce {
		if strings.HasPrefix(u, "http") {
			// We have a full URL - no change.
			break
		}

		if strings.HasPrefix(u, "github.com") {
			// We have a github.com specific string.
			u = "https://" + u
			break
		}

		ua := strings.Split(u, "/")
		if len(ua) == 0 {
			// Dunno, leave as is.
			break
		}

		if strings.Contains(ua[0], ".") {
			// We have a host defined in the first segment.
			u = "https://" + u
			break
		}

		// We probably just have a "owner/repo_name" style URL.
		u = "https://github.com/" + u
	}

	return u
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


func (gf *GoFiles) GetMeta() *GoMeta {
	var ret *GoMeta

	for range onlyOnce {
		for _, f := range *gf {
			//ux.PrintflnBlue("# %s", f.Path.GetPath())
			//ux.PrintflnCyan("%v", f.Ast.Name)
			ret = f.GetMeta()
			if ret != nil {
				break
			}
		}
	}

	return ret
}


func (gf *GoFile) GetMeta() *GoMeta {
	for range onlyOnce {
		var ok bool
		for _, object := range gf.Ast.Decls {
			if ok {
				break
			}

			//ux.PrintflnBlue("%s => %v", name, object)
			switch decl := object.(type) {
				case *ast.FuncDecl:
					//ux.PrintflnBlue("Func")
				case *ast.GenDecl:
					for _, spec := range decl.Specs {
						ok = gf.getGenDecl(spec)
					}
				default:
					ux.PrintflnBlue("Unknown declaration @\n", decl.Pos())
			}
		}
		if ok {
			gf.meta.Valid = true
			break
		}
	}

	return gf.meta
}


func (gf *GoFile) getGenDecl(spec ast.Spec) bool {
	var ok bool

	switch spec := spec.(type) {
		case *ast.ImportSpec:
			//ux.PrintflnBlue("Import", spec.Path.Value)
		case *ast.TypeSpec:
			//ux.PrintflnBlue("Type", spec.Name.String())
		case *ast.ValueSpec:
			for _, id := range spec.Names {
				name, value := getValueSpec(id)
				if name == "" {
					continue
				}
				gf.meta.Set(name, value)
				ok = true
			}
		default:
			ux.PrintflnBlue("Unknown token type\n")	// : %s\n", decl.Tok)
	}

	return ok
}


func getValueSpec(id *ast.Ident) (string, string) {
	var name string
	var value string
	switch id.Name {
		case BinaryName:
			fallthrough
		case BinaryVersion:
			fallthrough
		case SourceRepo:
			fallthrough
		case BinaryRepo:
			name = id.Name
	}

	expr := id.Obj.Decl.(*ast.ValueSpec).Values[0]
	value = getExpr(expr)
	//fmt.Printf("%s => %s", name, value)
	//fmt.Printf("\n")

	//ux.PrintflnBlue("Var %s: %v", id.Name, id.Obj.Decl.(*ast.ValueSpec).Values[0].(*ast.BasicLit).Value)
	return name, value
}


func getExpr(expr ast.Expr) string {
	var ret string
	switch expr.(type) {
		case *ast.BasicLit:
			ret = expr.(*ast.BasicLit).Value
			ret = strings.TrimPrefix(ret, "\"")
			ret = strings.TrimSuffix(ret, "\"")
		case *ast.BinaryExpr:
			ret = getExpr(expr.(*ast.BinaryExpr).X) + getExpr(expr.(*ast.BinaryExpr).Y)
		case *ast.Ident:
			_, ret = getValueSpec(expr.(*ast.Ident))
			//id := expr.(*ast.BinaryExpr).Y.(*ast.Object).Decl
			//switch id.(type) {
			//	case *ast.ValueSpec:
			//		_, ret = getValueSpec(id.(*ast.ValueSpec))
			//}
	}
	return ret
}
