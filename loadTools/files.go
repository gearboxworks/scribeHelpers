package loadTools

import (
	"encoding/json"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)


func (at *TypeScribeArgs) LoadJsonFile() *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if at.JsonStruct == nil {
			at.State.SetError("Json structure is nil")
			break
		}

		if at.Json.IsNotOk() {
			break
		}

		at.JsonStruct.JsonFile.SetFileInfo(at.Json.File)
		at.JsonStruct.JsonString = at.Json.String
		at.JsonStruct.JsonString = strings.ReplaceAll(at.JsonStruct.JsonString, "\n", "")
		at.JsonStruct.JsonString = strings.ReplaceAll(at.JsonStruct.JsonString, "\t", "")

		// Process JSON string.
		at.JsonStruct.Json = make(map[string]interface{})
		err := json.Unmarshal([]byte(at.Json.String), &at.JsonStruct.Json)
		if err != nil {
			at.State.SetError("Processing error: %s", err)
			break
		}
	}

	return at.State
}


func (at *TypeScribeArgs) LoadTemplateFile() *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if at.JsonStruct == nil {
			at.State.SetError("Json structure is nil")
			break
		}

		if at.Template.IsNotOk() {
			break
		}


		at.JsonStruct.TemplateFile.SetFileInfo(at.Template.File)

		// Create template instance.
		var t *template.Template
		t, at.State = at.CreateTemplate()
		t.Option("missingkey=error")

		// Do it again - may have to perform recursion here.
		var err error
		at.TemplateRef, err = t.Parse(at.Template.String)
		if err != nil {
			at.State.SetError("Template read error: %s", err)
			break
		}
		at.TemplateRef.Option("missingkey=error")
	}

	return at.State
}


func (at *TypeScribeArgs) LoadOutputFile() *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if at.JsonStruct == nil {
			at.State.SetError("Json structure is nil")
			break
		}

		if at.Output.IsNotOk() {
			break
		}


		if at.Output.Filename == DefaultOutFile {
			at.State = at.Output.File.SetFileHandle(os.Stdout)
			if at.State.IsNotOk() {
				at.State.SetError("Output file error: %s", at.State.GetError())
				break
			}

		} else {
			at.State = at.Output.File.OpenFile()
			if at.State.IsNotOk() {
				at.State.SetError("Output file error: %s", at.State.GetError())
				break
			}
		}


		at.OutputFh, at.State = at.Output.File.GetFileHandle()
		if at.State.IsNotOk() {
			at.State.SetError("Output file error: %s", at.State.GetError())
			break
		}

		at.JsonStruct.OutFile.SetFileInfo(at.Output.File)
	}

	return at.State
}


type TypeArgFile struct {
	File     *toolPath.TypeOsPath

	Filename string
	String   string
	isFile   bool
	isDir    bool
	isString bool

	State    *ux.State
}


func (at *TypeArgFile) IsNil() *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}
	at.State = at.State.EnsureNotNil()
	return at.State
}


func (at *TypeArgFile) IsOk() bool {
	var ok bool

	for range onlyOnce {
		if at.File == nil {
			ok = false
			break
		}

		if at.isFile {
			ok = true
			break
		}

		if at.isString {
			ok = true
			break
		}

		ok = false
	}

	return ok
}


func (at *TypeArgFile) IsNotOk() bool {
	return !at.IsOk()
}


func (at *TypeArgFile) ChangeSuffix(suffix string) {
	s := filepath.Ext(at.Filename)
	at.Filename = at.Filename[:len(at.Filename) - len(s)] + suffix
}


func (at *TypeArgFile) GetPath() string {
	return at.File.GetPath()
}
func (at *TypeArgFile) GetPathAbs() string {
	return at.File.GetPathAbs()
}
func (at *TypeArgFile) GetPathRel() string {
	return at.File.GetPathRel()
}


func (at *TypeArgFile) Exists() bool {
	return at.File.Exists()
}


func (at *TypeArgFile) IsAFile() bool {
	var ok bool
	if at.File.Exists() {
		ok = true
	}
	return ok
}


func (at *TypeArgFile) IsAString() bool {
	var ok bool
	if at.File.NotExists() {
		ok = true
	}
	return ok
}


func (at *TypeArgFile) IsStdFd() bool {
	var ok bool

	for range onlyOnce {
		if at.File.Exists() {
			break
		}

		if at.String != "" {
			break
		}

		ok = true
	}

	return ok
}


func (at *TypeArgFile) IsStdout() bool {
	return at.IsStdFd()
}


func (at *TypeArgFile) IsStdin() bool {
	return at.IsStdFd()
}


// Check if this is a string of characters, instead of filename.
func (at *TypeArgFile) isAString(arg string) bool {
	var ok bool

	for range onlyOnce {
		ok = isAString(arg)
		if ok {
			break
		}

		// Try again with spaces and single quotes removed.
		arg = strings.TrimSpace(arg)
		arg = strings.TrimPrefix(arg, "'")
		arg = strings.TrimSuffix(arg, "'")
		arg = strings.TrimSpace(arg)
		ok = isAString(arg)
		if ok {
			break
		}

		// Try again with spaces and double quotes removed.
		arg = strings.TrimSpace(arg)
		arg = strings.TrimPrefix(arg, "\"")
		arg = strings.TrimSuffix(arg, "\"")
		arg = strings.TrimSpace(arg)
		ok = isAString(arg)
		if ok {
			break
		}

		ok = false
	}

	return ok
}


func isAString(arg string) bool {
	var ok bool

	for range onlyOnce {
		if arg == "" {
			ok = false
			break
		}

		if strings.HasPrefix(arg, "#") {
			ok = true
			break
		}

		if strings.HasPrefix(arg, "{") {
			ok = true
			break
		}

		ok = false
	}

	return ok
}


func (at *TypeArgFile) SetInputFile(file string, remove bool) *ux.State {
	for range onlyOnce {
		if file == "" {
			at.State.SetError("No input file specified.")
			break
		}

		if at.Filename == SelectIgnore {
			at.isFile = false
			at.isString = true
			break
		}

		if at.File == nil {
			at.File = toolPath.ToolNewPath(file)
		}

		if remove {
			at.File.SetRemoveable()
		}

		at.State = at.File.StatPath()
		if at.File.Exists() {
			at.State = at.File.ReadFile()
			if at.State.IsOk() {
				at.isFile = true
				at.isString = false

				at.String = at.File.GetContentString()
				at.Filename = file
				at.State.SetOk("Input file '%s' read OK.", at.File.GetPath())
				break
			}
		}

		if at.isAString(file) {
			at.isFile = false
			at.isString = true

			at.File.SetContents(file)
			at.String = file
			at.Filename = "string"
			at.State.SetOk("Input string set.")
			break
		}

		at.State.SetError("Argument is neither filename nor string.")
		break

	}

	return at.State
}


func (at *TypeArgFile) SetOutputFile(file string, overwrite bool) *ux.State {

	for range onlyOnce {
		if file == "" {
			// Assume STDOUT
			file = DefaultOutFile
		}
		if file == "-" {
			// Assume STDOUT
			file = DefaultOutFile
		}

		if at.File == nil {
			at.File = toolPath.ToolNewPath(file)
			//at.File.State.Clear()	// Special case.
		}
		at.isFile = true

		if file == DefaultOutFile {
			overwrite = true
		}

		if overwrite {
			at.State.SetOk("Output file '%s' set to writeable.", at.Filename)
			at.File.SetOverwriteable()
		}
	}

	return at.State
}


func (at *TypeArgFile) SetWorkingPath(file string, changeDir bool) *ux.State {

	for range onlyOnce {
		if file == "" {
			file = DefaultWorkingPath
		}

		if at.File == nil {
			at.File = toolPath.ToolNewPath(file)
		}
		at.isDir = true

		at.State = at.File.StatPath()
		if at.File.NotExists() {
			at.State.SetError("Error directory '%s' does not exist.")
			break
		}

		if file == DefaultWorkingPath {
			// No need to change dir to "."
			break
		}

		if !changeDir {
			break
		}

		at.State = at.File.Chdir()
		if at.State.IsNotOk() {
			at.State.SetError("Error changing directory: %s")
			break
		}
	}

	return at.State
}
