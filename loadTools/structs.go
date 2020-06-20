package loadTools

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"text/template"
)

const onlyOnce = "1"
var onlyTwice = []string{"", ""}


type TypeScribeArgs struct {
	Scribe          *TypeArgFile

	Json            *TypeArgFile
	Template        *TypeArgFile
	TemplateRef     *template.Template

	Output          *TypeArgFile

	WorkingPath     *TypeArgFile

	ExecShell      bool // Cmd: "run"
	Chdir          bool // Flag: --chdir
	RemoveTemplate bool // Flag: --rm-tmpl
	ForceOverwrite bool // Flag: --force
	RemoveOutput   bool // Flag: --rm-out
	QuietProgress  bool // Flag: --quiet
	Verbose        bool // Flag: --verbose
	Debug          bool // Flag: --debug
	StripHashBang  bool // If set, strips #! at the start of the template file.
	//AddBrackets    bool // If set, adds {{ and }} to the template file.

	HelpAll        bool
	HelpFunctions  bool
	HelpVariables  bool
	HelpExamples   bool

	JsonStruct     *jsonStruct

	Tools          template.FuncMap

	Runtime        *toolRuntime.TypeRuntime
	State          *ux.State
	valid          bool
}


func New(binary string, version string, debugFlag bool) *TypeScribeArgs {
	rt := toolRuntime.New(binary, version, debugFlag)

	p := TypeScribeArgs{
		Json:           NewArgFile(rt),		// &TypeArgFile{State: ux.NewState(binary, debugFlag)},

		Scribe:         NewArgFile(rt),		// &TypeArgFile{State: ux.NewState(binary, debugFlag)},

		Template:       NewArgFile(rt),		// &TypeArgFile{State: ux.NewState(binary, debugFlag)},
		TemplateRef:    nil,

		Output:         NewArgFile(rt),		// &TypeArgFile{State: ux.NewState(binary, debugFlag)},

		WorkingPath:    NewArgFile(rt),		// &TypeArgFile{State: ux.NewState(binary, debugFlag)},

		ExecShell:      false,
		Chdir:          false,
		RemoveTemplate: false,
		ForceOverwrite: false,
		RemoveOutput:   false,
		Debug:          false,
		StripHashBang:  false,
		Verbose:        false,

		JsonStruct:     nil,

		Tools:          make(template.FuncMap),

		Runtime:        rt,
		State:          ux.NewState(binary, debugFlag),
		valid:          false,
	}
	p.State.SetPackage("")
	p.State.SetFunctionCaller()

	p.Scribe.SetDefaults(DefaultScribeFile, DefaultScribeString)
	p.Json.SetDefaults(DefaultJsonFile, DefaultJsonString)
	p.Template.SetDefaults(DefaultTemplateFile, DefaultTemplateString)

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


func (at *TypeScribeArgs) PrintflnOk(format string, args ...interface{}) {
	if state := at.IsNil(); state.IsError() {
		return
	}
	if at.Verbose {
		ux.PrintflnOk(format, args...)
	}
}


func (at *TypeScribeArgs) PrintflnNotify(format string, args ...interface{}) {
	if state := at.IsNil(); state.IsError() {
		return
	}
	if at.Verbose {
		ux.PrintflnBlue(format, args...)
	}
}
