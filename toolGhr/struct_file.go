package toolGhr

import (
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
)


type TypeFile struct {
	Name     string					// `goptions:"-n, --name, description='Name of the file', obligatory"`
	//Latest   bool					// `goptions:"-l, --latest, description='Download latest release (required if tag is not specified)',mutexgroup='input'"`
	Label    string					// `goptions:"-l, --label, description='Label (description) of the file'"`
	Path     *toolPath.TypeOsPath	// WAS Name - *os.TypeFile `goptions:"-f, --file, description='TypeFile to upload (use - for stdin)', rdonly, obligatory"`
	Replace  bool					// `goptions:"-R, --replace, description='Replace asset with same name if it already exists (WARNING: not atomic, failure to upload will remove the original asset too)'"`
	JSON     bool					// `goptions:"-j, --json, description='Emit info as JSON instead of text'"`

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
			Replace: false,
			JSON:    false,

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


func (file *TypeFile) Set(f TypeFile) *ux.State {
	if state := file.IsNil(); state.IsError() {
		return state
	}
	file.state.SetFunction()

	for range onlyOnce {
		file.state = f.isValid()
		if file.state.IsNotOk() {
			break
		}

		//
	}

	return file.state
}
