package toolCopy

import (
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/zloylos/grsync"
	"os/exec"
	"reflect"
)


const (
	ConstMethodDefault = ConstMethodRsync
	ConstMethodRsync = "rsync"
	ConstMethodTar = "tar"
	ConstMethodCpio = "cpio"
	ConstMethodSftp = "sftp"
	ConstMethodCp = "cp"
)


// GoLang enums - much better than plain old C type enums!
type TypeCopyMethod struct
{
	Name        string
	Path        string
	AllowRemote bool
	Available   bool
	Function    interface{}
	Options     interface{}

	paths       *TypeOsCopyPaths
	state       *ux.State
}
type TypeCopyMethods struct {
	Selected *TypeCopyMethod
	All      []*TypeCopyMethod
}


func (c *TypeCopyMethod) Run(paths *TypeOsCopyPaths, args ...string) *ux.State {
	for range onlyOnce {
		fn := reflect.ValueOf(c.Function)
		//fmt.Printf("%v\n%v\n%v\n",
		//	fn.Kind(),
		//	fn.String(),
		//	fn.Type(),
		//	)

		rargs := make([]reflect.Value, len(args)+2)
		rargs[0] = reflect.ValueOf(c)
		rargs[1] = reflect.ValueOf(paths)
		for i, a := range args {
			rargs[i+2] = reflect.ValueOf(a)
		}

		ret := fn.Call(rargs)
		c.state = ret[0].Interface().(*ux.State)
	}

	return c.state
}


func _CopyMethodDefault() *TypeCopyMethod {
	return _CopyMethodRsync()
}
func _CopyMethodRsync() *TypeCopyMethod {
	var ret TypeCopyMethod
	{
		path, ok := _ExecExists(ConstMethodRsync)

		opts := grsync.RsyncOptions{
			HardLinks:     true,
			Verbose:       true,
			Archive:       true,
			OneFileSystem: true,
			Progress:      false,
			//RsyncProgramm:     path,
		}

		ret = TypeCopyMethod{
			Name:        ConstMethodRsync,
			Path:        path,
			AllowRemote: true,
			Available:   ok,
			Function:    _CopyRunRsync,
			Options:     opts,
			state:       ux.NewState(ConstMethodRsync, false),
		}
	}
	return &ret
}
func _CopyMethodTar() *TypeCopyMethod {
	var ret TypeCopyMethod
	{
		path, ok := _ExecExists(ConstMethodTar)

		opts := []string{""}

		ret = TypeCopyMethod{
			Name:        ConstMethodTar,
			Path:        path,
			AllowRemote: true,
			Available:   ok,
			Function:    _CopyRunTar,
			Options:     opts,
			state:       ux.NewState(ConstMethodTar, false),
		}
	}
	return &ret
}
func _CopyMethodCpio() *TypeCopyMethod {
	var ret TypeCopyMethod
	{
		path, ok := _ExecExists(ConstMethodCpio)

		opts := []string{""}

		ret = TypeCopyMethod{
			Name:        ConstMethodCpio,
			Path:        path,
			AllowRemote: true,
			Available:   ok,
			Function:    _CopyRunCpio,
			Options:     opts,
			state:       ux.NewState(ConstMethodCpio, false),
		}
	}
	return &ret
}
func _CopyMethodSftp() *TypeCopyMethod {
	var ret TypeCopyMethod
	{
		path, ok := _ExecExists(ConstMethodSftp)

		opts := []string{"-rf"}

		ret = TypeCopyMethod{
			Name:        ConstMethodSftp,
			Path:        path,
			AllowRemote: true,
			Available:   ok,
			Function:    _CopyRunSftp,
			Options:     opts,
			state:       ux.NewState(ConstMethodSftp, false),
		}
	}
	return &ret
}
func _CopyMethodCp() *TypeCopyMethod {
	var ret TypeCopyMethod
	{
		path, ok := _ExecExists(ConstMethodCp)

		opts := []string{"-rip"}

		ret = TypeCopyMethod{
			Name:        ConstMethodCp,
			Path:        path,
			AllowRemote: false,
			Available:   ok,
			Function:    _CopyRunCp,
			Options:     opts,
			state:       ux.NewState(ConstMethodCp, false),
		}
	}
	return &ret
}

func _ExecExists(e string) (string, bool) {
	var path string
	var ok bool

	for range onlyOnce {
		var err error
		path, err = exec.LookPath(e)
		if err != nil {
			break
		}
		ok = true
	}

	return path, ok
}


func NewCopyMethod() *TypeCopyMethods {
	var ret TypeCopyMethods

	for range onlyOnce {
		// Set priority of use.
		ret.All = append(ret.All, _CopyMethodRsync())
		ret.All = append(ret.All, _CopyMethodTar())
		ret.All = append(ret.All, _CopyMethodCpio())
		ret.All = append(ret.All, _CopyMethodSftp())
		ret.All = append(ret.All, _CopyMethodCp())

		for _, m := range ret.All {
			if m.Available {
				ret.Selected = m
				break
			}
		}
	}

	return &ret
}


func (p *TypeCopyMethods) GetSelected() *TypeCopyMethod {
	return p.Selected
}
func (p *TypeCopyMethods) GetOptions() interface{} {
	return p.Selected.Options
}
func (p *TypeCopyMethods) GetName() string {
	return p.Selected.Name
}
func (p *TypeCopyMethods) GetPath() string {
	return p.Selected.Path
}
func (p *TypeCopyMethods) GetAllowRemote() bool {
	return p.Selected.AllowRemote
}
func (p *TypeCopyMethods) GetAvailable() bool {
	return p.Selected.Available
}


func (p *TypeCopyMethods) SelectMethod(method string) bool {
	var ok bool

	for range onlyOnce {
		for _, m := range p.All {
			if m.Name != method {
				continue
			}

			if !m.Available {
				continue
			}

			ok = true
			break
		}
	}

	return ok
}
