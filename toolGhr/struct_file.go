package toolGhr

import (
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"io"
	"os"
	"path/filepath"
)


type TypeFile struct {
	Name  string				// `goptions:"-n, --name, description='Name of the file', obligatory"`
	Label string				// `goptions:"-l, --label, description='Label (description) of the file'"`
	Path  *toolPath.TypeOsPath	// WAS Name - *os.TypeFile `goptions:"-f, --file, description='TypeFile to upload (use - for stdin)', rdonly, obligatory"`

	fh       *os.File
	runtime  *toolRuntime.TypeRuntime
	state    *ux.State
}
func (f *TypeFile) IsNil() *ux.State {
	return ux.IfNilReturnError(f)
}

func NewFile(runtime *toolRuntime.TypeRuntime) *TypeFile {
	var f TypeFile
	runtime = runtime.EnsureNotNil()

	for range onlyOnce {
		f = TypeFile{
			Name:    "",
			Label:   "",
			Path:    toolPath.New(runtime),

			fh:      nil,
			runtime: runtime,
			state:   ux.NewState(runtime.CmdName, runtime.Debug),
		}
	}
	f.state.SetPackage("")
	f.state.SetFunctionCaller()
	return &f
}

func (f *TypeFile) isValid() *ux.State {
	if state := ux.IfNilReturnError(f); state.IsError() {
		return state
	}

	for range onlyOnce {
		f.state = f.state.EnsureNotNil()

		if f.Name == "" {
			f.state.SetError("filename is empty")
			break
		}
	}

	return f.state
}

func (f *TypeFile) Set(overwrite bool, label string, path ...string) *ux.State {
	if state := f.IsNil(); state.IsError() {
		return state
	}
	f.state.SetFunction()

	for range onlyOnce {
		if f == nil {
			f.state.SetError("Provided file was not valid")
			break
		}
		if label == "" {
			f.state.SetError("Require a label to upload.")
			break
		}

		if len(path) == 0 {
			path = []string{label}
		}

		if !f.Path.SetPath(path...) {
			f.state.SetError("provided file was not valid")
			break
		}

		//label = strings.ToLower(filepath.Base(label))	// @TODO - selfupdate lowercase workaround.
		label = filepath.Base(label)
		f.Label = label
		f.Name = label
		if overwrite {
			f.Path.SetOverwriteable()
		}

		f.state.SetOk()
	}

	return f.state
}

func (f *TypeFile) OpenRead() *ux.State {
	if state := f.IsNil(); state.IsError() {
		return state
	}
	f.state.SetFunction()

	for range onlyOnce {
		f.state = f.isValid()
		if f.state.IsNotOk() {
			break
		}

		f.state = f.Path.StatPath()
		if f.state.IsNotOk() {
			break
		}

		if f.Label == "" {
			f.Label = f.Path.GetFilename()
		}

		if f.Name == "" {
			f.Name = f.Path.GetFilename()
		}

		f.state = f.Path.OpenFileHandle()
		if f.state.IsNotOk() {
			break
		}

		f.fh = f.Path.FileHandle

		f.state.SetOk()
	}

	return f.state
}

func (f *TypeFile) OpenWrite() *ux.State {
	if state := f.IsNil(); state.IsError() {
		return state
	}
	f.state.SetFunction()

	for range onlyOnce {
		f.state = f.isValid()
		if f.state.IsNotOk() {
			break
		}

		if f.Label == "" {
			f.Label = f.Path.GetFilename()
		}

		if f.Name == "" {
			f.Name = f.Path.GetFilename()
		}

		//f.state = f.Base.StatPath()
		//if f.Base.NotExists() {
		//	break
		//}

		f.state = f.Path.OpenFile()
		if f.state.IsNotOk() {
			break
		}

		f.fh = f.Path.FileHandle

		f.state.SetOk()
	}

	return f.state
}

func (f *TypeFile) Write(data io.Reader, n int64) *ux.State {
	if state := f.IsNil(); state.IsError() {
		return state
	}
	f.state.SetFunction()

	for range onlyOnce {
		an, err := io.Copy(f.fh, data)
		if an != n {
			f.state.SetError("data did not match content length %d != %d", an, n)
			break
		}
		if err != nil {
			f.state.SetError(err)
			break
		}

		f.state.SetOk()
	}

	return f.state
}

func (f *TypeFile) Close() *ux.State {
	if state := f.IsNil(); state.IsError() {
		return state
	}
	f.state.SetFunction()

	for range onlyOnce {
		f.state = f.isValid()
		if f.state.IsNotOk() {
			break
		}

		f.state = f.Path.CloseFile()
		if f.state.IsNotOk() {
			break
		}

		f.state.SetOk()
	}

	return f.state
}
