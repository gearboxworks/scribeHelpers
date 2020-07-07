package loadTools

import (
	"encoding/json"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
)


type jsonStruct struct {
	Exec            *toolRuntime.TypeRuntime	`json:"Exec" mapstructure:"Exec"`

	TemplateFile    FileInfo	`json:"TemplateFile" mapstructure:"TemplateFile"`
	JsonFile        FileInfo	`json:"JsonFile" mapstructure:"JsonFile"`
	OutFile         FileInfo	`json:"OutFile" mapstructure:"OutFile"`
	//Env             toolSystem.Environment
	Env             *toolRuntime.Environment	`json:"Env" mapstructure:"Env"`

	JsonString      string	`json:"JsonString" mapstructure:"JsonString"`
	CreationEpoch   int64	`json:"CreationEpoch" mapstructure:"CreationEpoch"`
	CreationDate    string	`json:"CreationDate" mapstructure:"CreationDate"`
	CreationInfo    string	`json:"CreationInfo" mapstructure:"CreationInfo"`
	CreationWarning string	`json:"CreationWarning" mapstructure:"CreationWarning"`

	Json            map[string]interface{}	`json:"Json" mapstructure:"Json"`

	state           *ux.State	`json:"-"`
}

type JsonMap map[string]interface{}


func NewJsonStruct(binary string, version string, debugFlag bool) *jsonStruct {
	js := jsonStruct {
		Exec:            toolRuntime.New(binary, version, debugFlag),
		TemplateFile:    FileInfo{},
		JsonFile:        FileInfo{},
		OutFile:         FileInfo{},
		//Env:             nil,
		JsonString:      "",
		CreationEpoch:   0,
		CreationDate:    "",
		CreationInfo:    "",
		CreationWarning: "",
		Json:            nil,
	}

	return &js
}


func (at *jsonStruct) IsNil() *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}
	at.state = at.state.EnsureNotNil()
	return at.state
}


func (at *jsonStruct) LoadJsonFile(jsonRef *TypeArgFile) *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if jsonRef.IsNotOk() {
			break
		}

		at.JsonFile.SetFileInfo(jsonRef.TypeOsPath)
		at.JsonString = jsonRef.GetContentString()
		at.JsonString = strings.ReplaceAll(at.JsonString, "\n", "")
		at.JsonString = strings.ReplaceAll(at.JsonString, "\t", "")

		// Process JSON string.
		at.Json = make(map[string]interface{})
		err := json.Unmarshal([]byte(jsonRef.GetContentString()), &at.Json)
		if err != nil {
			at.state.SetError("Json processing error: %s", err)
			break
		}

		at.state.SetOk()
	}

	return at.state
}


func (at *jsonStruct) LoadTemplateFile(templateRef *TypeArgFile) *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if templateRef.IsNotOk() {
			at.state.SetError("Template processing error")
			break
		}

		at.TemplateFile.SetFileInfo(templateRef.TypeOsPath)
		at.state.SetOk()
	}

	return at.state
}


func (at *jsonStruct) LoadOutputFile(templateRef *TypeArgFile) *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//if at.Output.File == DefaultOutFile {
		//	at.State = at.Output.SetFileHandle(os.Stdout)
		//	if at.State.IsNotOk() {
		//		at.State.SetError("Output file error: %s", at.State.GetError())
		//		break
		//	}
		//
		//} else {
		//	at.State = at.Output.OpenFile()
		//	if at.State.IsNotOk() {
		//		at.State.SetError("Output file error: %s", at.State.GetError())
		//		break
		//	}
		//}
		//
		//
		//at.OutputFh, at.State = at.Output.GetFileHandle()
		//if at.State.IsNotOk() {
		//	at.State.SetError("Output file error: %s", at.State.GetError())
		//	break
		//}

		at.OutFile.SetFileInfo(templateRef.TypeOsPath)
	}

	return at.state
}


type FileInfo struct {
	Dir           string
	Name          string
	CreationEpoch int64
	CreationDate  string

	State         *ux.State
}


func (fi *FileInfo) SetFileInfo(path *toolPath.TypeOsPath) {
	fi.Dir = path.GetDirname()
	if fi.Dir == "" {
		fi.Dir = "."
	}

	fi.Name = path.GetFilename()
	if fi.Name == "" {
		fi.Name = path.Path
	}

	fi.CreationDate = path.GetModTimeString()
	fi.CreationEpoch = path.GetModTimeEpoch()
	fi.State = path.State
}


func UnescapeString(s string) string {

	// \a	Alert or bell
	// \b	Backspace
	// \\	Backslash
	// \t	Horizontal tab
	// \n	Line feed or newline
	// \f	Form feed
	// \r	Carriage return
	// \v	Vertical tab
	// \'	Single quote (only in rune literals)
	// \"	Double quote (only in string literals)

	s = strings.ReplaceAll(s, `\a`, "\a")
	s = strings.ReplaceAll(s, `\b`, "\b")
	s = strings.ReplaceAll(s, `\\`, "\\")
	s = strings.ReplaceAll(s, `\t`, "\t")
	s = strings.ReplaceAll(s, `\n`, "\n")
	s = strings.ReplaceAll(s, `\f`, "\f")
	s = strings.ReplaceAll(s, `\r`, "\r")
	s = strings.ReplaceAll(s, `\v`, "\v")
	s = strings.ReplaceAll(s, `\'`, `'`)
	s = strings.ReplaceAll(s, `\"`, `"`)

	return s
}
