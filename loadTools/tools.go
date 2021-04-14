package loadTools

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/gearboxworks/scribeHelpers/toolCopy"
	"github.com/gearboxworks/scribeHelpers/toolExec"
	"github.com/gearboxworks/scribeHelpers/toolGit"
	"github.com/gearboxworks/scribeHelpers/toolGitHub"
	"github.com/gearboxworks/scribeHelpers/toolPath"
	"github.com/gearboxworks/scribeHelpers/toolPrompt"
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/toolService"
	"github.com/gearboxworks/scribeHelpers/toolSystem"
	"github.com/gearboxworks/scribeHelpers/toolTypes"
	"github.com/gearboxworks/scribeHelpers/toolUx"
	"github.com/gearboxworks/scribeHelpers/ux"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"text/template"
)


const ToolPrefix = "Tool"
type Tool struct {
	File string
	Function string
	Name string
	Args string
	Return string
}
type Tools map[string]Tool
type Files map[string]Tools
type SortedTools []Tool


// This method will auto-import exported Tool functions within each Tool package.
// You need to run `pkgreflect scribe/Tools` after code changes.
func DiscoverTools() *ux.State {
	state := ux.NewState("scribe", false)
	var tfm template.FuncMap

	for range onlyOnce {
		// Define additional template functions.
		tfm = sprig.TxtFuncMap()

		for name, fn := range toolCopy.GetTools {
			tfm[name] = fn
		}

		//for name, fn := range toolDocker.GetTools {
		//	tfm[name] = fn
		//}

		for name, fn := range toolExec.GetTools {
			tfm[name] = fn
		}

		//for name, fn := range toolGear.GetTools {
		//	tfm[name] = fn
		//}

		for name, fn := range toolGit.GetTools {
			tfm[name] = fn
		}

		for name, fn := range toolGitHub.GetTools {
			tfm[name] = fn
		}

		for name, fn := range toolPath.GetTools {
			tfm[name] = fn
		}

		for name, fn := range toolPrompt.GetTools {
			tfm[name] = fn
		}

		for name, fn := range toolRuntime.GetTools {
			tfm[name] = fn
		}

		for name, fn := range toolService.GetTools {
			tfm[name] = fn
		}

		for name, fn := range toolSystem.GetTools {
			tfm[name] = fn
		}

		for name, fn := range toolTypes.GetTools {
			tfm[name] = fn
		}

		for name, fn := range toolUx.GetTools {
			tfm[name] = fn
		}
	}

	state.SetResponse(&tfm)
	return state
}


// @TODO - Add the ability to import from an external package.
// You need to run `pkgreflect scribe/Tools` after code changes.
func AddTools(i interface{}) *ux.State {
	state := ux.NewState("scribe", false)
	var tfm template.FuncMap

	for range onlyOnce {
		//for name, fn := range deploywp.GetTools {
		//	tfm[name] = fn
		//}
	}

	state.SetResponse(&tfm)
	return state
}


// This method will print exported Tool functions within each Tool package.
// You need to run `pkgreflect scribe/Tools` after code changes.
func PrintTools() string {
	var ret string

	for range onlyOnce {
		ret += ux.SprintfCyan("List of defined template functions:\n")

		state := DiscoverTools()
		if state.IsNotOk() {
			ret += ux.SprintfRed("Error discovering Tools.\n")
			break
		}

		tfm := responseToFuncMap(state.GetResponse())
		if tfm == nil {
			ret += ux.SprintfRed("Error discovering Tools.\n")
		}

		files := make(Files)
		for name, fn := range *tfm {
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
	}

	return ret
}
func (a SortedTools) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}


func responseToFuncMap(r *ux.TypeResponse) *template.FuncMap {
	var tfm *template.FuncMap

	for range onlyOnce {
		if !r.IsOfType("template.FuncMap") {
			break
		}
		tfm = r.Pointer().(*template.FuncMap)
	}

	return tfm
}


func _GetFunctionInfo(i interface{}) *Tool {
	var Tool Tool

	for range onlyOnce {
		ptr := reflect.ValueOf(i).Pointer()
		ptrs := reflect.ValueOf(i).String()
		ptrn := runtime.FuncForPC(ptr).Name()

		Tool.Name = filepath.Ext(ptrn)[1:]
		Tool.File = ptrn[0:len(ptrn)-len(Tool.Name)-1]
		Tool.Name = strings.TrimPrefix(Tool.Name, ToolPrefix)

		// ptrs == <func(...interface {}) *toolSystem.TypeReadFile Value>
		Tool.Function = strings.Replace(ptrs, "<func", Tool.Name, -1)
		Tool.Function = strings.TrimSuffix(Tool.Function, " Value>")
		// Tool.Function == (...interface {}) *toolSystem.TypeReadFile

		Tool.Args = strings.Split(ptrs, "(")[1]
		Tool.Args = strings.Split(Tool.Args, ")")[0]

		Tool.Return = strings.TrimSuffix(ptrs, " Value>")
		Tool.Return = strings.Split(Tool.Return, ")")[1]
		Tool.Return = strings.TrimSpace(Tool.Return)
		Tool.Return = strings.TrimPrefix(Tool.Return, "(")

		//if Tool.Name == "generateCertificateAuthority" {
		//	fmt.Printf(".")
		//}
	}

	return &Tool
}
