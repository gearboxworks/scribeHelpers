package toolGhr

import (
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"io"
	"os"
	"path/filepath"
)


func (repo *TypeRepo) FileRead(file string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		repo.state = repo.file.OpenRead(file)
		if repo.file.Path.NotExists() {
			break
		}
	}

	return repo.state
}

func (repo *TypeRepo) SetFile(file string, label string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()

	for range onlyOnce {
		repo.file.Name = file
		repo.file.Label = label
		if label == "" {
			repo.file.Label = filepath.Base(repo.file.Name)
		}
		repo.state.SetOk()
	}

	return repo.state
}

func (repo *TypeRepo) SetFileLabel(label string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()
	repo.file.Label = label
	if repo.file.Label == "" {
		repo.file.Label = repo.file.Name
	}
	return repo.state
}

func (repo *TypeRepo) SetFileName(file string) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()
	repo.file.Name = file
	return repo.state
}

func (repo *TypeRepo) FileOpenWrite(file string, overwrite bool) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()
	return repo.file.OpenWrite(file, overwrite)
}

func (repo *TypeRepo) FileWrite(data io.Reader, n int64) *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()
	return repo.file.Write(data, n)
}

func (repo *TypeRepo) FileClose() *ux.State {
	if state := repo.IsNil(); state.IsError() {
		return state
	}
	repo.state.SetFunction()
	return repo.file.Close()
}


type TypeFile struct {
	Handle   *os.File
	Name     string					// `goptions:"-n, --name, description='Name of the file', obligatory"`
	Label    string					// `goptions:"-l, --label, description='Label (description) of the file'"`
	Path     *toolPath.TypeOsPath	// WAS Name - *os.TypeFile `goptions:"-f, --file, description='TypeFile to upload (use - for stdin)', rdonly, obligatory"`

	//Latest   bool					// `goptions:"-l, --latest, description='Download latest Release (required if tag is not specified)',mutexgroup='input'"`
	//Replace  bool					// `goptions:"-R, --replace, description='Replace asset with same name if it already exists (WARNING: not atomic, failure to upload will remove the original asset too)'"`
	//JSON     bool					// `goptions:"-j, --json, description='Emit info as JSON instead of text'"`

	runtime  *toolRuntime.TypeRuntime
	state    *ux.State
}

func NewFile(runtime *toolRuntime.TypeRuntime) *TypeFile {
	var f TypeFile
	runtime = runtime.EnsureNotNil()

	for range onlyOnce {
		f = TypeFile{
			Name:    "",
			//Latest:  false,
			Label:   "",
			Path:    toolPath.New(runtime),
			//Replace: false,
			//JSON:    false,

			runtime: runtime,
			state:   ux.NewState(runtime.CmdName, runtime.Debug),
		}
	}
	f.state.SetPackage("")
	f.state.SetFunctionCaller()
	return &f
}

func (file *TypeFile) IsNil() *ux.State {
	if state := ux.IfNilReturnError(file); state.IsError() {
		return state
	}
	file.state = file.state.EnsureNotNil()
	return file.state
}

func (file *TypeFile) isValid() *ux.State {
	if state := ux.IfNilReturnError(file); state.IsError() {
		return state
	}

	for range onlyOnce {
		file.state = file.state.EnsureNotNil()

		if file.Name == "" {
			file.state.SetError("filename is empty")
			break
		}
	}

	return file.state
}

func (file *TypeFile) OpenRead(f string) *ux.State {
	if state := file.IsNil(); state.IsError() {
		return state
	}
	file.state.SetFunction()

	for range onlyOnce {
		file.Name = f

		file.state = file.isValid()
		if file.state.IsNotOk() {
			break
		}

		if !file.Path.SetPath(f) {
			file.state.SetError("provided file was not valid")
			break
		}

		file.state = file.Path.StatPath()
		if file.state.IsNotOk() {
			break
		}

		if file.Label == "" {
			file.Label = file.Path.GetFilename()
		}

		if file.Name == "" {
			file.Name = file.Path.GetFilename()
		}

		file.state = file.Path.OpenFileHandle()
		if file.state.IsNotOk() {
			break
		}

		file.Handle = file.Path.FileHandle

		file.state.SetOk()
	}

	return file.state
}

func (file *TypeFile) OpenWrite(f string, overwrite bool) *ux.State {
	if state := file.IsNil(); state.IsError() {
		return state
	}
	file.state.SetFunction()

	for range onlyOnce {
		file.Name = f

		file.state = file.isValid()
		if file.state.IsNotOk() {
			break
		}

		if !file.Path.SetPath(f) {
			file.state.SetError("provided file was not valid")
			break
		}

		file.state = file.Path.StatPath()
		if file.state.IsNotOk() {
			break
		}

		if overwrite {
			file.Path.SetOverwriteable()
		}

		if file.Label == "" {
			file.Label = file.Path.GetFilename()
		}

		if file.Name == "" {
			file.Name = file.Path.GetFilename()
		}

		file.state = file.Path.OpenFile()
		if file.state.IsNotOk() {
			break
		}

		file.Handle = file.Path.FileHandle

		file.state.SetOk()
	}

	return file.state
}

func (file *TypeFile) Write(data io.Reader, n int64) *ux.State {
	if state := file.IsNil(); state.IsError() {
		return state
	}
	file.state.SetFunction()

	for range onlyOnce {
		an, err := io.Copy(file.Handle, data)
		if an != n {
			file.state.SetError("data did not match content length %d != %d", an, n)
			break
		}
		if err != nil {
			file.state.SetError(err)
			break
		}

		file.state.SetOk()
	}

	return file.state
}

func (file *TypeFile) Close() *ux.State {
	if state := file.IsNil(); state.IsError() {
		return state
	}
	file.state.SetFunction()

	for range onlyOnce {
		file.state = file.isValid()
		if file.state.IsNotOk() {
			break
		}

		file.state = file.Path.CloseFile()
		if file.state.IsNotOk() {
			break
		}

		file.state.SetOk()
	}

	return file.state
}
