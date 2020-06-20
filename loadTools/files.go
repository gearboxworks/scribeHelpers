package loadTools

import (
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"path/filepath"
	"strings"
)


type TypeArgFile struct {
	File          string
	DefaultString string
	DefaultFile   string
	valid         bool
	*toolPath.TypeOsPath

	//String   string
	//isDir    bool
	//isFile   bool
	//isString bool
	//
	//State    *ux.State
}


func NewArgFile(rt *toolRuntime.TypeRuntime) *TypeArgFile {
	rt = rt.EnsureNotNil()

	af := TypeArgFile{
		File:          "",
		DefaultString: "",
		DefaultFile:   "",
		valid:         false,
		TypeOsPath:    toolPath.New(rt),
	}

	return &af
}


func (at *TypeArgFile) IsNil() *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}
	at.State = at.State.EnsureNotNil()
	return at.State
}


func (at *TypeArgFile) SetDefaults(file string, str string) *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}
	at.DefaultFile = file
	at.DefaultString = str
	return at.State
}


func (at *TypeArgFile) IsSet() bool {
	var ok bool
	for range onlyOnce {
		if !at.valid {
			break
		}

		if at.File == "" {
			break
		}

		if at.File == SelectIgnore {
			break
		}

		if at.File == SelectDefault {
			break
		}

		if at.File == SelectStdout {
			break
		}

		ok = true
	}
	return ok
}
func (at *TypeArgFile) IsNotSet() bool {
	return !at.IsSet()
}


func (at *TypeArgFile) IsFileSet() bool {
	var ok bool
	for range onlyOnce {
		if at.File == SelectFile {
			ok = true
			break
		}
	}
	return ok
}
func (at *TypeArgFile) IsNotFileSet() bool {
	return !at.IsFileSet()
}


func (at *TypeArgFile) IsOk() bool {
	return at.valid
}
func (at *TypeArgFile) IsNotOk() bool {
	return !at.valid
}


func (at *TypeArgFile) ChangeDir() bool {
	var ok bool
	for range onlyOnce {
		if at.File != SelectFile {
			break
		}

		if at.Chdir().IsNotOk() {
			break
		}

		ok = true
	}

	return ok
}


func (at *TypeArgFile) IsStdFd() bool {
	if at.File == SelectStdout {
		return true
	}
	return false
}


func (at *TypeArgFile) IsStdout() bool {
	return at.IsStdFd()
}


func (at *TypeArgFile) IsStdin() bool {
	return at.IsStdFd()
}


func (at *TypeArgFile) Ignore() *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}
	return at.SetInputFile(SelectIgnore)
}


func (at *TypeArgFile) IsIgnore() bool {
	if state := at.IsNil(); state.IsError() {
		return true
	}
	if at.File == SelectIgnore {
		return true
	}
	return false
}


func (at *TypeArgFile) IsNotIgnore() bool {
	return !at.IsIgnore()
}


func (at *TypeArgFile) SetInputFile(fileName string) *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}

	for range onlyOnce {
		if fileName == "" {
			at.State.SetError("No input file specified.")
			break
		}

		if fileName == SelectIgnore {
			at.valid = true
			at.File = SelectIgnore
			at.SetContents(at.DefaultString)
			at.State.SetOk("Ignore file.")
			break
		}

		if fileName == SelectDefault {
			at.valid = true
			at.File = SelectDefault
			at.SetContents(at.DefaultString)
			at.State.SetOk("Input string set.")
			break
		}

		if fileName == at.DefaultFile {
			at.valid = true
			at.File = SelectDefault
			at.SetContents(at.DefaultString)
			at.State.SetOk("Input string set.")
			break
		}

		at.SetPath(fileName)
		at.State = at.StatPath()
		if at.Exists() {
			at.State = at.ReadFile()
			if at.State.IsOk() {
				at.valid = true
				at.File = SelectFile
				at.State.SetOk("Input file '%s' read OK.", at.GetPath())
				break
			}
		}

		if at.isAString(fileName) {
			at.valid = true
			at.File = SelectString
			at.SetContents(fileName)
			at.State.SetOk("Input string set.")
			break
		}

		at.valid = false
		at.State.SetError("Argument is neither filename nor string.")
	}

	return at.State
}


func (at *TypeArgFile) SetOutputFile(fileName string, overwrite bool) *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}

	for range onlyOnce {
		if fileName == SelectStdout {
			fileName = DefaultOutFile
		} else if fileName == "" {
			fileName = DefaultOutFile
		} else if fileName == "-" {
			fileName = DefaultOutFile
		}

		if fileName == DefaultOutFile {
			at.valid = true
			at.File = SelectDefault
			at.State.SetOk()
			at.SetOverwriteable()
			at.SetFileHandle(os.Stdout)
			break
		}

		at.SetPath(fileName)
		at.State = at.StatPath()
		if overwrite {
			at.SetOverwriteable()
			//at.State.SetOk()
			//break
		}

		//if at.Exists() {
		//	at.State.SetError("Output file '%s' exists.", at.GetPath())
		//	break
		//}

		at.State = at.OpenFile()
		if at.State.IsNotOk() {
			break
		}
	}

	return at.State
}


func (at *TypeArgFile) SetInputString(fileString string) *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}

	for range onlyOnce {
		if fileString == "" {
			at.State.SetError("No input string specified.")
			break
		}

		at.valid = true
		at.File = SelectString
		at.SetContents(fileString)
		at.State.SetOk()
	}

	return at.State
}


func (at *TypeArgFile) SetInputStringArray(stringArray []string) *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}

	for range onlyOnce {
		if len(stringArray) == 0 {
			at.State.SetError("No input string specified.")
			break
		}

		at.valid = true
		at.File = SelectString
		at.SetContents(stringArray)
		at.State.SetOk()
	}

	return at.State
}


func (at *TypeArgFile) SetDefaultString() *ux.State {
	return at.SetInputString(at.DefaultString)
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


func ChangeSuffix(file string, suffix string) string {
	s := filepath.Ext(file)
	file = file[:len(file) - len(s)] + suffix
	return file
}
