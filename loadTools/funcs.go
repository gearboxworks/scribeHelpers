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
	"sort"
	"strings"
	"text/template"
)


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


		////////////////////////////////////////////////////
		// Fetch input files.
		for range onlyOnce {
			// Validate json and template files/strings.
			if at.Json.Filename == DefaultJsonFile {
				at.Json.Filename = DefaultJsonString
			} else if at.Json.Filename == SelectIgnore {
				at.Json.Filename = DefaultJsonString
			}
			at.Json.SetInputFile(at.Json.Filename, false)

			if at.Template.Filename == DefaultTemplateFile {
				at.Template.Filename = DefaultTemplateString
			} else if at.Template.Filename == SelectIgnore {
				at.Template.Filename = DefaultTemplateString
			}
			at.Template.SetInputFile(at.Template.Filename, at.RemoveTemplate)

			at.State.Clear()

			// json:empty && tmpl:empty
			if at.Json.IsNotOk() && at.Template.IsNotOk() {
				at.Json.SetInputFile(DefaultJsonFile, false)
				if at.Json.IsNotOk() {
					at.State.SetError("Neither template nor json provided.")
					break
				}

				at.Template.Filename = at.Json.Filename
				at.Template.ChangeSuffix(DefaultTemplateFileSuffix)
				at.Template.SetInputFile(at.Template.Filename, false)
				if at.Json.IsNotOk() {
					at.State.SetError("Neither template nor json provided.")
					break
				}
			}

			// json:OK && tmpl:OK
			if at.Json.IsOk() && at.Template.IsOk() {
				at.State.SetOk()
				break
			}

			// json:OK && tmpl:empty
			if at.Json.IsOk() && at.Template.IsNotOk() {
				at.Template.Filename = at.Json.Filename
				at.Template.ChangeSuffix(DefaultTemplateFileSuffix)
				at.Template.SetInputFile(at.Template.Filename, false)
				if at.Template.IsNotOk() {
					at.State.SetError("Template not provided.")
					break
				}
			}

			// json:empty && tmpl:OK
			if at.Json.IsNotOk() && at.Template.IsOk() {
				at.Json.Filename = at.Template.Filename
				at.Json.ChangeSuffix(DefaultJsonFileSuffix)
				at.Json.SetInputFile(at.Json.Filename, false)
				if at.Json.IsNotOk() {
					at.State.SetError("Json not provided.")
					break
				}
			}

			// json:empty && tmpl:empty
			if at.Json.IsNotOk() && at.Template.IsNotOk() {
				at.State.SetError("Neither template nor json provided.")
				break
			}
		}

		if at.State.IsNotOk() {
			break
		}


		////////////////////////////////////////////////////
		// Strip out #! at start of template.
		for range onlyOnce {
			if !at.StripHashBang {
				break
			}

			ca := at.Template.File.GetContentArray()
			if len(ca) == 0 {
				break
			}

			if !strings.HasPrefix(ca[0], "#!") {
				break
			}

			at.Template.File.SetContents(ca[1:])
		}


		////////////////////////////////////////////////////
		// Add {{ and }} to template file.
		for range onlyOnce {
			if !at.AddBrackets {
				break
			}

			ca := at.Template.File.GetContentArray()
			if len(ca) == 0 {
				break
			}

			for l := range ca {
				ca[l] = fmt.Sprintf("{{- %s }}", ca[l])
			}

			at.Template.File.SetContents(ca)
		}


		////////////////////////////////////////////////////
		// Output file.
		if at.Output.Filename == SelectStdout {
			at.Output.Filename = DefaultOutFile
			at.ForceOverwrite = true
		} else if at.Output.Filename == SelectConvert {
			at.Output.Filename = strings.TrimSuffix(at.Template.Filename, DefaultTemplateFileSuffix)
		}
		at.State = at.Output.SetOutputFile(at.Output.Filename, at.ForceOverwrite)
		if at.State.IsNotOk() {
			break
		}


		////////////////////////////////////////////////////
		// Chdir.
		if at.Chdir {
			at.State = at.Json.File.Chdir()
			if at.State.IsNotOk() {
				at.State.SetError("Error changing directory: %s")
				break
			}
		}


		////////////////////////////////////////////////////
		// WorkingPath
		at.State = at.WorkingPath.SetWorkingPath(at.WorkingPath.Filename, true)
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
		if at.JsonStruct == nil {
			at.JsonStruct = NewJsonStruct(at.Runtime.CmdName, at.Runtime.CmdVersion, at.Debug)
		}

		// Historic reasons...
		at.JsonStruct.CreationEpoch = at.JsonStruct.Exec.TimeStampEpoch()
		at.JsonStruct.CreationDate = at.JsonStruct.Exec.TimeStampString()
		at.JsonStruct.Env = at.JsonStruct.Exec.GetEnvMap()

		at.State = at.LoadJsonFile()
		if at.State.IsNotOk() {
			at.State.SetError("Json error: %s", at.State.GetError())
			break
		}

		at.State = at.LoadTemplateFile()
		if at.State.IsNotOk() {
			at.State.SetError("Template error: %s", at.State.GetError())
			break
		}

		at.JsonStruct.CreationInfo = fmt.Sprintf("Created on %s, using template:%s and json:%s", at.JsonStruct.CreationDate, at.JsonStruct.TemplateFile.Name, at.JsonStruct.JsonFile.Name)
		at.JsonStruct.CreationWarning = "WARNING: This file has been auto-generated. DO NOT EDIT: WARNING"

		at.State.Clear()
	}

	return at.State
}


func (at *TypeScribeArgs) Run() *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		if at.Output.Filename == "" {
			at.State.SetError("No output file specified.")
			break
		}

		at.State = at.LoadOutputFile()
		if at.State.IsNotOk() {
			break
		}

		err := at.TemplateRef.Execute(at.OutputFh, &at.JsonStruct)
		//err := at.TemplateRef.Execute(os.Stdout, &at.JsonStruct)
		if err != nil {
			at.State.SetError("Error processing template: %s", err)
			break
		}

		at.State = at.Output.File.CloseFile()
		if at.State.IsNotOk() {
			at.State.SetError("Error closing output: %s", err)
			break
		}

		// Are we treating this as a shell script?
		if at.ExecShell {
			ux.PrintflnOk("Executing file '%s'", at.Output.Filename)

			//outFile := toolPath.ToolNewPath(at.OutFile)
			bashFile := at.Output.File
			at.State = bashFile.StatPath()
			if at.State.IsNotOk() {
				//at.State.SetError("Shell script error: %s", err)
				break
			}
			bashFile.Chmod(0755)

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

			at.State = exe.SetArgs(bashFile.GetPath())
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

			if at.QuietProgress {
				fmt.Printf("# STDOUT from script: %s\n", bashFile.GetPath())
				fmt.Printf("%s\n", exe.GetStdoutString())

				fmt.Printf("# STDERR from script: %s\n", bashFile.GetPath())
				fmt.Printf("%s\n", exe.GetStderrString())

				fmt.Printf("# Exit code from script: %s\n", bashFile.GetPath())
				fmt.Printf("%d\n", exe.GetExitCode())
			}

			if at.RemoveOutput {
				at.State = bashFile.RemoveFile()
				if at.State.IsError() {
					at.State.SetError("Shell script error: %s", at.State.GetError())
					break
				}
			}
		}

		if at.RemoveTemplate {
			at.State = at.Template.File.RemoveFile()
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
		//// Define additional template functions.
		//at.State = DiscoverTools()
		//if at.State.IsNotOk() {
		//	break
		//}
		//
		//tfm := responseToFuncMap(at.State.GetResponse())
		//at.State = at.ImportTools(tfm)
		//if at.State.IsNotOk() {
		//	break
		//}
		//
		//// Add inbuilt Tools.
		//at.Tools["PrintTools"] = PrintTools

		if at.Tools == nil {
			at.State = at.ImportTools(nil)
			if at.State.IsError() {
				break
			}
		}

		t := template.New("JSON").Funcs(at.Tools)
		if t == nil {
			at.State.SetError("Template creation error.")
			break
		}

		t.Option("missingkey=error")

		// Do it again - may have to perform recursion here.
		var err error
		at.TemplateRef, err = t.Parse(at.Template.String)
		if err != nil {
			at.State.SetError("Template read error: %s", err)
			break
		}
		if at.TemplateRef == nil {
			at.State.SetError("Template creation error.")
			break
		}

		at.TemplateRef.Option("missingkey=error")
	}

	return at.State
}


// Ability to import from an external package.
// You need to run `pkgreflect scribe/tools` after code changes.
// func (at *TypeScribeArgs) ImportTools(h map[string]reflect.Value) *ux.State {
func (at *TypeScribeArgs) ImportTools(h *template.FuncMap) *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		//if h == nil {
		//	at.State.SetError("Error importing Tools - empty list.")
		//	break
		//}

		at.Tools = make(template.FuncMap)
		// Define external template functions.
		at.Tools = sprig.TxtFuncMap()

		for name, fn := range *h {
			at.Tools[name] = fn
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


//func (at *TypeScribeArgs) ImportTools(h *template.FuncMap) *ux.State {
//	if state := at.IsNil(); state.IsError() {
//		return state
//	}
//
//	for range onlyOnce {
//		if h == nil {
//			at.State.SetError("Error importing Tools - empty list.")
//			break
//		}
//
//		for name, fn := range *h {
//			at.Tools[name] = fn
//		}
//
//		//// Define additional template functions.
//		//for name, fn := range h {
//		//	// Ignore GetTools function.
//		//	if name == "GetTools" {
//		//		continue
//		//	}
//		//
//		//	// Ignore any function that doesn't have a ToolPrefix
//		//	if !strings.HasPrefix(name, "Tool") {
//		//		continue
//		//	}
//		//
//		//	// Trim ToolPrefix from function template name.
//		//	name = strings.TrimPrefix(name, "Tool")
//		//	at.Tools[name] = fn.Interface()
//		//}
//	}
//
//	return at.State
//}


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
			//fmt.Printf("Name[%s]: %s => %s\n", name, Tool.Name, Tool.Function)
		}

		for fn, fp := range files {
			ret += ux.SprintfWhite("\n# Tool functions within: %s\n", fn)

			// To store the keys in slice in sorted order
			var keys SortedTools
			for _, k := range fp {
				keys = append(keys, k)
			}
			sort.Slice(keys, keys.Less)

			//for _, hp := range fp {
			for _, hp := range keys {
				ret += fmt.Sprintf("%s( %s )\t=> ( %s )\n",
					ux.SprintfGreen(hp.Name),
					ux.SprintfCyan(hp.Args),
					ux.SprintfYellow(hp.Return),
				)

				// fmt.Printf("%s\n\targs: %s\n\tReturn: %s\n", hp.Function, hp.args, hp.Return)
			}
		}

		ret += ux.SprintfBlue("\nSee http://masterminds.github.io/sprig/ for additional functions...\n")

		fmt.Print(ret)
	}
}
