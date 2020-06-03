package loadHelpers

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/helperExec"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"os/exec"
	"strings"
	"text/template"
)


func (at *TypeScribeArgs) ValidateArgs() *ux.State {

	for range OnlyOnce {
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


		for range OnlyOnce {
			// Validate json and template files/strings.
			if at.Json.Filename == DefaultJsonFile {
				at.Json.Filename = DefaultJsonString
			}
			at.Json.SetInputFile(at.Json.Filename, false)

			if at.Json.Filename == DefaultTemplateFile {
				at.Json.Filename = DefaultTemplateString
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


		at.SetValid()
		at.State.SetOk("Processed arguments.")
	}

	return at.State
}


func (at *TypeScribeArgs) Load() *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		if at.JsonStruct == nil {
			at.JsonStruct = NewJsonStruct(at.Exec.CmdName, at.Exec.CmdVersion, at.Debug)
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

	for range OnlyOnce {
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

			//outFile := helperPath.HelperNewPath(at.OutFile)
			bashFile := at.Output.File
			at.State = bashFile.StatPath()
			if at.State.IsNotOk() {
				//at.State.SetError("Shell script error: %s", err)
				break
			}
			bashFile.Chmod(0755)

			exe := helperExec.New(at.Debug)
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

			at.State = exe.SetPath(path)
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


func (at *TypeScribeArgs) CreateTemplate() (*template.Template, *ux.State) {
	var t *template.Template
	if state := at.IsNil(); state.IsError() {
		return nil, state
	}

	for range OnlyOnce {
		// Define additional template functions.
		at.State = DiscoverHelpers()
		if at.State.IsNotOk() {
			break
		}
		at.LoadHelpers(at.State.Response.(template.FuncMap))

		// Add inbuilt helpers.
		at.Helpers["PrintHelpers"] = PrintHelpers

		t = template.New("JSON").Funcs(at.Helpers)
	}

	return t, at.State
}


// Ability to import from an external package.
// You need to run `pkgreflect scribe/helpers` after code changes.
// func (at *TypeScribeArgs) LoadHelpers(h map[string]reflect.Value) *ux.State {
func (at *TypeScribeArgs) LoadHelpers(h template.FuncMap) *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	for range OnlyOnce {
		for name, fn := range h {
			at.Helpers[name] = fn
		}

		//// Define additional template functions.
		//for name, fn := range h {
		//	// Ignore GetHelpers function.
		//	if name == "GetHelpers" {
		//		continue
		//	}
		//
		//	// Ignore any function that doesn't have a HelperPrefix
		//	if !strings.HasPrefix(name, "Helper") {
		//		continue
		//	}
		//
		//	// Trim HelperPrefix from function template name.
		//	name = strings.TrimPrefix(name, "Helper")
		//	at.Helpers[name] = fn.Interface()
		//}
	}

	return at.State
}


func (at *TypeScribeArgs) PrintHelpers() {
	_, _ = fmt.Fprintf(os.Stderr, PrintHelpers())
}
