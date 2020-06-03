package loadHelpers

import (
	"github.com/newclarity/scribeHelpers/helperRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"text/template"
)

const OnlyOnce = "1"


type TypeScribeArgs struct {
	Json            *TypeArgFile
	Template        *TypeArgFile
	TemplateRef     *template.Template
	Output          *TypeArgFile
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

	Runtime        *helperRuntime.TypeRuntime
	State          *ux.State
	valid          bool
}


func New(binary string, version string, debugFlag bool) *TypeScribeArgs {

	p := TypeScribeArgs{
		Json:           &TypeArgFile{State: ux.NewState(debugFlag)},
		Template:       &TypeArgFile{State: ux.NewState(debugFlag)},
		TemplateRef:    nil,
		Output:         &TypeArgFile{State: ux.NewState(debugFlag)},
		OutputFh:       nil,

		ExecShell:      false,
		Chdir:          false,
		RemoveTemplate: false,
		ForceOverwrite: false,
		RemoveOutput:   false,
		Debug:          false,

		JsonStruct:     nil,

		Helpers:        make(template.FuncMap),

		Runtime:        helperRuntime.New(binary, version, debugFlag),
		State:          ux.NewState(debugFlag),
		valid:          false,
	}

	p.State.SetPackage("")
	p.State.SetFunctionCaller()

	return &p
}

func (at *TypeScribeArgs) IsNil() *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}
	at.State = at.State.EnsureNotNil()
	return at.State
}

func (at *TypeScribeArgs) IsValid() bool {
	return at.valid
}

func (at *TypeScribeArgs) SetValid() {
	at.valid = true
}

func (at *TypeScribeArgs) SetInvalid() {
	at.valid = false
}
