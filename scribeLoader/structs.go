package scribeLoader

import (
	"github.com/newclarity/scribeHelpers/helperRuntime"
	"os"
	"github.com/newclarity/scribe/ux"
	"text/template"
)

const OnlyOnce = "1"


type ArgTemplate struct {
	Exec            *helperRuntime.Exec

	Json            TypeArgFile
	Template        TypeArgFile
	TemplateRef     *template.Template
	Output          TypeArgFile
	OutputFh        *os.File

	ExecShell      bool // Cmd: "run"
	Chdir          bool // Flag: --chdir
	RemoveTemplate bool // Flag: --rm-tmpl
	ForceOverwrite bool // Flag: --force
	RemoveOutput   bool // Flag: --rm-out
	QuietProgress  bool // Flag: --quiet
	Debug          bool // Flag: --debug

	HelpAll        bool
	HelpFunctions  bool
	HelpVariables  bool
	HelpExamples   bool

	JsonStruct     *jsonStruct

	Helpers        template.FuncMap

	State          *ux.State
	valid          bool
}


func NewArgTemplate(binary string, version string) *ArgTemplate {

	p := ArgTemplate{
		Exec:           helperRuntime.NewExec(binary, version),

		Json:           TypeArgFile{State: ux.NewState(false)},
		Template:       TypeArgFile{State: ux.NewState(false)},
		TemplateRef:    nil,
		Output:         TypeArgFile{State: ux.NewState(false)},
		OutputFh:       nil,

		ExecShell:      false,
		Chdir:          false,
		RemoveTemplate: false,
		ForceOverwrite: false,
		RemoveOutput:   false,
		Debug:          false,

		JsonStruct:     nil,

		Helpers:        make(template.FuncMap),

		State:          ux.NewState(false),
		valid:          false,
	}

	p.State.SetPackage("")
	p.State.SetFunctionCaller()

	return &p
}

func (at *ArgTemplate) IsNil() *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}
	at.State = at.State.EnsureNotNil()
	return at.State
}

func (at *ArgTemplate) IsValid() bool {
	return at.valid
}

func (at *ArgTemplate) SetValid() {
	at.valid = true
}

func (at *ArgTemplate) SetInvalid() {
	at.valid = false
}
