package loadTools

import (
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
	"github.com/spf13/cobra"
	"text/template"
)

const onlyOnce = "1"
var onlyTwice = []string{"", ""}


type TypeScribeArgs struct {
	ConfigPath     string		`json:"config_path" mapstructure:"config_path"`
	ConfigDir      string		`json:"config_dir" mapstructure:"config_dir"`
	ConfigFile     string		`json:"config_file" mapstructure:"config_file"`

	//ConfigFile  *TypeArgFile       `json:"config_file" mapstructure:"config_file"`

	Scribe      *ScribeFile        `json:"scribe_file" mapstructure:"config_file"`

	Json        *JsonFile          `json:"data_file" mapstructure:"data_file"`

	Template    *TemplateFile      `json:"template_file" mapstructure:"template_file"`
	TemplateRef *template.Template `json:"-"`

	Output      *TypeArgFile       `json:"output" mapstructure:"output"`

	WorkingPath *TypeArgFile       `json:"working_path" mapstructure:"working_path"`

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

	JsonStruct     *jsonStruct      `json:"data" mapstructure:"data"`

	Tools          template.FuncMap `json:"-"`

	Runtime *toolRuntime.TypeRuntime `json:"-"`
	State   *ux.State                `json:"-"`
	valid   bool                     `json:"-"`
	cmd     *cobra.Command           `json:"-"`
	SelfCmd *cobra.Command
}
func (at *TypeScribeArgs) IsNil() *ux.State {
	return ux.IfNilReturnError(at)
}


func New(binary string, version string, debugFlag bool) *TypeScribeArgs {
	var p TypeScribeArgs

	for range onlyOnce {
		rt := toolRuntime.New(binary, version, debugFlag)

		p = TypeScribeArgs{
			Json: &JsonFile{TypeArgFile: NewArgFile(rt)},

			Scribe: &ScribeFile{TypeArgFile: NewArgFile(rt)},

			Template:    &TemplateFile{TypeArgFile: NewArgFile(rt)},
			TemplateRef: nil,

			Output: NewArgFile(rt), // &TypeArgFile{State: ux.NewState(binary, debugFlag)},

			WorkingPath: NewArgFile(rt), // &TypeArgFile{State: ux.NewState(binary, debugFlag)},

			ExecShell:      false,
			Chdir:          false,
			RemoveTemplate: false,
			ForceOverwrite: false,
			RemoveOutput:   false,
			Debug:          false,
			StripHashBang:  false,
			Verbose:        false,

			JsonStruct: nil,

			Tools: make(template.FuncMap),

			Runtime: rt,
			State:   ux.NewState(binary, debugFlag),
			valid:   false,
			//cmdHelp: toolCobraHelp.New(rt),
		}
		p.State.SetPackage("")
		p.State.SetFunctionCaller()

		p.Scribe.SetDefaults(DefaultScribeFile, DefaultScribeString)
		p.Json.SetDefaults(DefaultJsonFile, DefaultJsonString)
		p.Template.SetDefaults(DefaultTemplateFile, DefaultTemplateString)

		p.State = p.ImportTools(nil)
		if p.State.IsError() {
			break
		}
	}

	return &p
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
