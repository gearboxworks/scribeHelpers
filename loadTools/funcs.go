package loadTools

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/newclarity/scribeHelpers/toolCopy"
	"github.com/newclarity/scribeHelpers/toolExec"
	"github.com/newclarity/scribeHelpers/toolGhr"
	"github.com/newclarity/scribeHelpers/toolGit"
	"github.com/newclarity/scribeHelpers/toolGitHub"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolPrompt"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/toolService"
	"github.com/newclarity/scribeHelpers/toolSystem"
	"github.com/newclarity/scribeHelpers/toolTypes"
	"github.com/newclarity/scribeHelpers/toolUx"
	"github.com/newclarity/scribeHelpers/ux"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)


func (at *TypeScribeArgs) ProcessArgs(cmd string, args []string) *ux.State {
	for range onlyOnce {
		at.State = at.CheckArgs(cmd, args...)
		if at.State.IsNotOk() {
			break
		}

		at.State = at.ValidateArgs()
		if at.State.IsNotOk() {
			break
		}
	}

	return at.State
}


func (at *TypeScribeArgs) CheckArgs(cmd string, args ...string) *ux.State {
	for range onlyTwice {
		if len(args) >= 1 {
			ext := filepath.Ext(args[0])

			if ext == ".scribe" {
				if at.Scribe.IsNotIgnore() {
					at.PrintflnNotify("Setting scribe file '%s'", args[0])
					at.Template.File = args[0]
					args = args[1:]
				}
			}

			if ext == ".json" {
				if at.Json.IsNotIgnore() {
					at.PrintflnNotify("Setting JSON file '%s'", args[0])
					at.Json.File = args[0]
					args = args[1:]
				}
				continue
			}

			if ext == ".tmpl" {
				if at.Template.IsNotIgnore() {
					at.PrintflnNotify("Setting template file '%s'", args[0])
					at.Template.File = args[0]
					args = args[1:]
				}
			}
		}
	}

	for range onlyOnce {
		err := at.Runtime.SetArgs(cmd)
		if err != nil {
			at.State.SetError(err)
			break
		}

		err = at.Runtime.AddArgs(args...)
		if err != nil {
			at.State.SetError(err)
			break
		}
	}

	return at.State
}


func (at *TypeScribeArgs) ProcessInputFiles() *ux.State {
	for range onlyOnce {
		// Fetch input files.
		for range onlyOnce {
			// Validate scribe file OR string.
			at.State = at.Scribe.SetInputFile(at.Scribe.File)
			if at.State.IsNotOk() {
				break
			}
			// scribe:OK
			if at.Scribe.IsSet() {
				// Add {{ and }} to scribe file.
				ca := at.Scribe.GetContentArray()
				if len(ca) == 0 {
					break
				}

				for l := range ca {
					ca[l] = fmt.Sprintf("{{ %s }}", ca[l])
				}

				at.PrintflnNotify("Using scribe file '%s' of %d bytes.", at.Scribe.GetPath(), at.Scribe.GetContentLength())
				at.Template.SetInputStringArray(ca)
				at.Json.SetDefaultString()
				break
			}

			// Validate json file OR string.
			at.State = at.Json.SetInputFile(at.Json.File)
			if at.State.IsNotOk() {
				break
			}
			at.PrintflnNotify("Using JSON file '%s' of %d bytes.", at.Json.GetPath(), at.Json.GetContentLength())

			// Validate template file OR string.
			at.State = at.Template.SetInputFile(at.Template.File)
			if at.State.IsNotOk() {
				break
			}
			if at.RemoveTemplate {
				at.PrintflnNotify("Will remove template file '%s' afterwards.", at.Template.GetPath())
				at.Template.SetRemoveable()
			}
			at.PrintflnNotify("Using template file '%s' of %d bytes.", at.Template.GetPath(), at.Template.GetContentLength())
		}
		if at.State.IsNotOk() {
			break
		}


		// If JSON set and template not set, try and use the JSON filename with a tmpl extension.
		for range onlyOnce {
			if at.Json.IsNotSet() {
				break
			}

			if at.Template.IsSet() {
				break
			}

			newFile := ChangeSuffix(at.Json.GetPath(), DefaultTemplateFileSuffix)
			at.State = at.Template.SetInputFile(newFile)
			if at.Template.IsNotOk() {
				at.State.SetError("Template not provided.")
				break
			}
		}
		if at.State.IsNotOk() {
			break
		}


		// If Template set and JSON not set, try and use the template filename with a json extension.
		for range onlyOnce {
			if at.Template.IsNotSet() {
				break
			}

			if at.Json.IsSet() {
				break
			}

			newFile := ChangeSuffix(at.Template.GetPath(), DefaultJsonFileSuffix)
			at.State = at.Json.SetInputFile(newFile)
			if at.Json.IsNotOk() {
				at.State.SetError("Json not provided.")
				break
			}
		}
		if at.State.IsNotOk() {
			break
		}


		// If Template set and JSON not set, try and use the template filename with a json extension.
		if at.Json.IsNotOk() && at.Template.IsNotOk() {
			at.State.SetError("No input files provided.")
			break
		}


		// Strip out #! at start of template.
		for range onlyOnce {
			if !at.StripHashBang {
				break
			}

			ca := at.Template.GetContentArray()
			if len(ca) == 0 {
				break
			}

			if !strings.HasPrefix(ca[0], "#!") {
				break
			}

			at.PrintflnNotify("Stripping '#!' from template file '%s'.", at.Template.GetPath())
			at.Template.SetContents(ca[1:])
		}


		// Set output file.
		for range onlyOnce {
			if at.Output.File != SelectConvert {
				break
			}

			if at.Template.IsNotFileSet() {
				at.Output.File = SelectStdout
				break
			}

			at.Output.File = strings.TrimSuffix(at.Template.GetPathAbs(), DefaultTemplateFileSuffix)
		}
		at.State = at.Output.SetOutputFile(at.Output.File, at.ForceOverwrite)
		if at.State.IsNotOk() {
			break
		}


		at.State.SetOk("Processed input files.")
	}

	return at.State
}


// if at.Chdir { Attempt to change directories to any of the input files. }
func (at *TypeScribeArgs) ChangeDir() *ux.State {
	for range onlyOnce {
		if !at.Chdir {
			break
		}

		if at.Scribe.ChangeDir() {
			at.PrintflnNotify("Changed to scribe directory '%s'", at.Scribe.GetDirname())
			at.State.SetOk()
			break
		}

		if at.Json.ChangeDir() {
			at.PrintflnNotify("Changed to JSON directory '%s'", at.Json.GetDirname())
			at.State.SetOk()
			break
		}

		if at.Template.ChangeDir() {
			at.PrintflnNotify("Changed to template directory '%s'", at.Template.GetDirname())
			at.State.SetOk()
			break
		}
	}

	return at.State
}


func (at *TypeScribeArgs) ValidateArgs() *ux.State {

	for range onlyOnce {
		at.SetInvalid()		// Start with invalid.

		// Debug mode.
		if at.Debug {
			at.ForceOverwrite = false
			at.RemoveOutput = false
			at.RemoveTemplate = false
		} else {
			//at.OverWriteOutput = false
			//at.RemoveOutput = false
			//at.RemoveTemplate = false
		}


		// Change to working path first.
		at.WorkingPath.SetPath(at.WorkingPath.File)
		at.State = at.WorkingPath.Chdir()
		if at.State.IsNotOk() {
			break
		}


		// Process and load all input files.
		at.State = at.ProcessInputFiles()
		if at.State.IsNotOk() {
			break
		}


		// Attempt to change directories to any of the input files.
		at.State = at.ChangeDir()
		if at.State.IsNotOk() {
			break
		}


		at.SetValid()
		at.State.SetOk("Processed arguments.")
	}

	return at.State
}


func (at *TypeScribeArgs) Load() *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		at.PrintflnNotify("Loading template file '%s'", at.Template.GetPath())

		if at.JsonStruct == nil {
			at.JsonStruct = NewJsonStruct(at.Runtime.CmdName, at.Runtime.CmdVersion, at.Debug)
		}

		// Historic reasons...
		at.JsonStruct.CreationEpoch = at.JsonStruct.Exec.TimeStampEpoch()
		at.JsonStruct.CreationDate = at.JsonStruct.Exec.TimeStampString()
		at.JsonStruct.Env = at.JsonStruct.Exec.GetEnvMap()

		at.State = at.JsonStruct.LoadJsonFile(at.Json)
		if at.State.IsNotOk() {
			break
		}

		at.State = at.JsonStruct.LoadTemplateFile(at.Template)
		if at.State.IsNotOk() {
			break
		}

		at.JsonStruct.CreationInfo = fmt.Sprintf("Created on %s, using template:%s and json:%s", at.JsonStruct.CreationDate, at.JsonStruct.TemplateFile.Name, at.JsonStruct.JsonFile.Name)
		at.JsonStruct.CreationWarning = "WARNING: This file has been auto-generated. DO NOT EDIT: WARNING"

		at.State = at.CreateTemplate()
		if at.State.IsError() {
			break
		}

		at.State.SetOk()
	}

	return at.State
}


func (at *TypeScribeArgs) Run() *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		at.PrintflnNotify("Processing template file '%s'. Output sent to '%s'", at.Template.GetPath(), at.Output.GetPath())

		err := at.TemplateRef.Execute(at.Output.FileHandle, &at.JsonStruct)
		//err := at.TemplateRef.Execute(os.Stdout, &at.JsonStruct)
		if err != nil {
			at.State.SetError("Error processing template: %s", err)
			break
		}

		at.State = at.Output.CloseFile()
		if at.State.IsNotOk() {
			at.State.SetError("Error closing output: %s", err)
			break
		}

		// Are we treating this as a shell script?
		if at.ExecShell {
			at.PrintflnNotify("Executing file '%s'", at.Output.GetPath())

			//outFile := toolPath.ToolNewPath(at.OutFile)
			//bashFile := at.Output
			at.State = at.Output.StatPath()
			if at.State.IsNotOk() {
				//at.State.SetError("Shell script error: %s", err)
				break
			}
			at.Output.Chmod(0755)

			exe := toolExec.New(at.Runtime)
			at.State = exe.State
			if at.State.IsError() {
				at.State.SetError("Shell script error: %s", at.State.GetError())
				break
			}

			path, err := exec.LookPath("bash")
			if err != nil {
				at.State.SetError("Shell script error: %s", err)
				break
			}

			at.State = exe.SetCmd(path)
			if at.State.IsError() {
				at.State.SetError("Shell script error: %s", at.State.GetError())
				break
			}

			at.State = exe.SetArgs(at.Output.GetPath())
			if at.State.IsError() {
				at.State.SetError("Shell script error: %s", at.State.GetError())
				break
			}

			if at.QuietProgress {
				exe.SilenceProgress()
			} else {
				exe.ShowProgress()
			}

			at.State = exe.Run()
			if at.State.IsError() {
				at.State.SetError("Shell script error: %s", at.State.GetError())
				break
			}

			//if at.QuietProgress {
			//	fmt.Printf("# STDOUT from script: %s\n", at.Output.GetPath())
			//	fmt.Printf("%s\n", exe.GetStdoutString())
			//
			//	fmt.Printf("# STDERR from script: %s\n", at.Output.GetPath())
			//	fmt.Printf("%s\n", exe.GetStderrString())
			//
			//	fmt.Printf("# Exit code from script: %s\n", at.Output.GetPath())
			//	fmt.Printf("%d\n", exe.GetExitCode())
			//}

			if at.RemoveOutput {
				at.PrintflnNotify("Removing output file '%s'.", at.Output.GetPath())
				at.State = at.Output.RemoveFile()
				if at.State.IsError() {
					at.State.SetError("Shell script error: %s", at.State.GetError())
					break
				}
			}
		}

		if at.RemoveTemplate {
			at.PrintflnNotify("Removing template file '%s'.", at.Template.GetPath())
			at.State = at.Template.RemoveFile()
			if at.State.IsNotOk() {
				break
			}
		}
	}

	return at.State
}


func (at *TypeScribeArgs) CreateTemplate() *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//if at.Tools == nil {
			at.State = at.ImportTools(nil)
			if at.State.IsError() {
				break
			}
		//}

		at.TemplateRef = template.New("JSON")
		if at.TemplateRef == nil {
			at.State.SetError("Template error - cannot init.")
			break
		}

		at.TemplateRef = at.TemplateRef.Funcs(at.Tools)
		if at.TemplateRef == nil {
			at.State.SetError("Template error - cannot load tools.")
			break
		}

		at.TemplateRef = at.TemplateRef.Option("missingkey=error")
		if at.TemplateRef == nil {
			at.State.SetError("Template error - cannot set options.")
			break
		}

		// Do it again - may have to perform recursion here.
		var err error
		at.TemplateRef, err = at.TemplateRef.Parse(at.Template.GetContentString())
		if err != nil {
			at.State.SetError("Template error - cannot parse - %v", err)
			break
		}

		at.TemplateRef = at.TemplateRef.Option("missingkey=error")
		if at.TemplateRef == nil {
			at.State.SetError("Template error - cannot set options.")
			break
		}
	}

	return at.State
}


// Ability to import from an external package.
// You need to run `buildtool pkgreflect scribe/tools` after code changes.
// OR add a "go:generate buildtool pkgreflect scribe/tools" comment to main.go.
// func (at *TypeScribeArgs) ImportTools(h map[string]reflect.Value) *ux.State {
func (at *TypeScribeArgs) ImportTools(h *template.FuncMap) *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		at.Tools = make(template.FuncMap)
		// Define external template functions.
		at.Tools = sprig.TxtFuncMap()

		if h != nil {
			for name, fn := range *h {
				at.Tools[name] = fn
			}
		}
		for name, fn := range toolCopy.GetTools {
			at.Tools[name] = fn
		}

		//for name, fn := range toolDocker.GetTools {
		//	at.Tools[name] = fn
		//}

		for name, fn := range toolExec.GetTools {
			at.Tools[name] = fn
		}

		//for name, fn := range toolGear.GetTools {
		//	at.Tools[name] = fn
		//}

		for name, fn := range toolGhr.GetTools {
			at.Tools[name] = fn
		}

		for name, fn := range toolGit.GetTools {
			at.Tools[name] = fn
		}

		for name, fn := range toolGitHub.GetTools {
			at.Tools[name] = fn
		}

		//for name, fn := range toolGoReleaser.GetTools {
		//	at.Tools[name] = fn
		//}

		for name, fn := range toolPath.GetTools {
			at.Tools[name] = fn
		}

		for name, fn := range toolPrompt.GetTools {
			at.Tools[name] = fn
		}

		for name, fn := range toolRuntime.GetTools {
			at.Tools[name] = fn
		}

		//for name, fn := range toolSelfUpdate.GetTools {
		//	at.Tools[name] = fn
		//}

		for name, fn := range toolService.GetTools {
			at.Tools[name] = fn
		}

		for name, fn := range toolSystem.GetTools {
			at.Tools[name] = fn
		}

		for name, fn := range toolTypes.GetTools {
			at.Tools[name] = fn
		}

		for name, fn := range toolUx.GetTools {
			at.Tools[name] = fn
		}


		// Add inbuilt Tools.
		at.Tools["PrintTools"] = PrintTools
	}

	return at.State
}


func (at *TypeScribeArgs) PrintTools() {
	for range onlyOnce {
		var ret string

		ret += ux.SprintfCyan("List of defined template functions:\n")

		files := make(Files)
		for name, fn := range at.Tools {
			Tool := _GetFunctionInfo(fn)

			if _, ok := files[Tool.File]; !ok {
				files[Tool.File] = make(Tools)
			}

			files[Tool.File][name] = *Tool
		}

		for fn, fp := range files {
			ret += ux.SprintfWhite("\n# Tool functions within: %s\n", fn)

			// To store the keys in slice in sorted order
			var keys SortedTools
			for _, k := range fp {
				keys = append(keys, k)
			}
			sort.Slice(keys, keys.Less)

			for _, hp := range keys {
				ret += fmt.Sprintf("%s( %s )\t=> ( %s )\n",
					ux.SprintfGreen(hp.Name),
					ux.SprintfCyan(hp.Args),
					ux.SprintfYellow(hp.Return),
				)
			}
		}

		ret += ux.SprintfBlue("\nSee http://masterminds.github.io/sprig/ for additional functions...\n")

		fmt.Print(ret)
	}
}
