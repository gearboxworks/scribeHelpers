package scribeLoader

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/newclarity/scribeHelpers/helperCopy"
	"github.com/newclarity/scribeHelpers/helperExec"
	"github.com/newclarity/scribeHelpers/helperGit"
	"github.com/newclarity/scribeHelpers/helperGitHub"
	"github.com/newclarity/scribeHelpers/helperPath"
	"github.com/newclarity/scribeHelpers/helperPrompt"
	"github.com/newclarity/scribeHelpers/helperService"
	"github.com/newclarity/scribeHelpers/helperSystem"
	"github.com/newclarity/scribeHelpers/helperTypes"
	"github.com/newclarity/scribeHelpers/helperUx"
	"github.com/newclarity/scribe/ux"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"text/template"
)


const HelperPrefix = "Helper"

type Helper struct {
	File string
	Function string
	Name string
	Args string
	Return string
}
type Helpers map[string]Helper
type Files map[string]Helpers
type SortedHelpers []Helper


// This method will auto-import exported helper functions within each helper package.
// You need to run `pkgreflect scribe/helpers` after code changes.
func DiscoverHelpers() *ux.State {
	state := ux.NewState(false)
	var tfm template.FuncMap

	for range OnlyOnce {
		// Define additional template functions.
		tfm = sprig.TxtFuncMap()

		//for name, fn := range deploywp.GetHelpers {
		//	tfm[name] = fn
		//}

		for name, fn := range helperCopy.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperExec.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperGit.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperGitHub.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperPath.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperPrompt.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperService.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperSystem.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperTypes.GetHelpers {
			tfm[name] = fn
		}

		for name, fn := range helperUx.GetHelpers {
			tfm[name] = fn
		}
	}

	state.Response = tfm
	return state
}


// @TODO - Add the ability to import from an external package.
// You need to run `pkgreflect scribe/helpers` after code changes.
func AddHelpers(i interface{}) *ux.State {
	state := ux.NewState(false)
	var tfm template.FuncMap

	for range OnlyOnce {
		//for name, fn := range deploywp.GetHelpers {
		//	tfm[name] = fn
		//}
	}

	state.Response = tfm
	return state
}


// This method will print exported helper functions within each helper package.
// You need to run `pkgreflect scribe/helpers` after code changes.
func PrintHelpers() string {
	var ret string

	for range OnlyOnce {
		ret += ux.SprintfCyan("List of defined template functions:\n")

		tfm := DiscoverHelpers()
		if tfm.IsNotOk() {
			ret += ux.SprintfRed("Error discovering helpers.\n")
			break
		}


		files := make(Files)
		for name, fn := range tfm.Response.(template.FuncMap) {
			helper := _GetFunctionInfo(fn)

			if _, ok := files[helper.File]; !ok {
				files[helper.File] = make(Helpers)
			}

			files[helper.File][name] = *helper
			//fmt.Printf("Name[%s]: %s => %s\n", name, helper.Name, helper.Function)
		}

		for fn, fp := range files {
			ret += ux.SprintfWhite("\n# Helper functions within: %s\n", fn)

			// To store the keys in slice in sorted order
			var keys SortedHelpers
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
func (a SortedHelpers) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}


func _GetFunctionInfo(i interface{}) *Helper {
	var helper Helper

	for range OnlyOnce {
		ptr := reflect.ValueOf(i).Pointer()
		ptrs := reflect.ValueOf(i).String()
		ptrn := runtime.FuncForPC(ptr).Name()

		helper.Name = filepath.Ext(ptrn)[1:]
		helper.File = ptrn[0:len(ptrn)-len(helper.Name)-1]
		helper.Name = strings.TrimPrefix(helper.Name, HelperPrefix)

		// ptrs == <func(...interface {}) *helperSystem.TypeReadFile Value>
		helper.Function = strings.Replace(ptrs, "<func", helper.Name, -1)
		helper.Function = strings.TrimSuffix(helper.Function, " Value>")
		// helper.Function == (...interface {}) *helperSystem.TypeReadFile

		helper.Args = strings.Split(ptrs, "(")[1]
		helper.Args = strings.Split(helper.Args, ")")[0]

		helper.Return = strings.TrimSuffix(ptrs, " Value>")
		helper.Return = strings.Split(helper.Return, ")")[1]
		helper.Return = strings.TrimSpace(helper.Return)
		helper.Return = strings.TrimPrefix(helper.Return, "(")

		//if helper.Name == "generateCertificateAuthority" {
		//	fmt.Printf(".")
		//}
	}

	return &helper
}
