package loader

import (
	"github.com/newclarity/scribeHelpers/helperPath"
	"github.com/newclarity/scribeHelpers/helperRuntime"
	"github.com/newclarity/scribe/ux"
	"strings"
)


type jsonStruct struct {
	Exec            *helperRuntime.Exec

	TemplateFile    FileInfo
	JsonFile        FileInfo
	OutFile         FileInfo
	//Env             helperSystem.Environment
	Env             *helperRuntime.Environment

	JsonString      string
	CreationEpoch   int64
	CreationDate    string
	CreationInfo    string
	CreationWarning string

	Json            map[string]interface{}

	state           *ux.State
}

func (at *jsonStruct) IsNil() *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}
	at.state = at.state.EnsureNotNil()
	return at.state
}


func NewJsonStruct(binary string, version string) *jsonStruct {
	js := jsonStruct {
		Exec:            helperRuntime.NewExec(binary, version),
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


type FileInfo struct {
	Dir           string
	Name          string
	CreationEpoch int64
	CreationDate  string

	State         *ux.State
}


func (fi *FileInfo) SetFileInfo(path *helperPath.TypeOsPath) {
	fi.Dir = path.GetDirname()
	fi.Name = path.GetFilename()
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
