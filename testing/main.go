/*
The sole purpose of this file's existence is for unit testing all the package modules.
Additionally, it helps within the GoLand IDE for entity mapping and checking.

This file is not to be used for any production related code.
*/
package main

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/toolCopy"
	"github.com/newclarity/scribeHelpers/toolDocker"
	"github.com/newclarity/scribeHelpers/toolExec"
	"github.com/newclarity/scribeHelpers/toolGear"
	"github.com/newclarity/scribeHelpers/toolGhr"
	"github.com/newclarity/scribeHelpers/toolGit"
	"github.com/newclarity/scribeHelpers/toolGitHub"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolPrompt"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/toolService"
	"github.com/newclarity/scribeHelpers/toolSystem"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"strings"
)


const onlyOnce = "1"

func init() {
	_, _ = ux.Open("testing", true)
}

var globalDebug bool

func PrintTestStart(test string) {
	ux.PrintflnWhite("################################################################################")
	ux.PrintflnBlue("# %s - STARTED", test)
}

func PrintTestResult(state *ux.State, test string, sub string, args ...interface{}) {
	ux.PrintfBlue("# %s", test)
	ux.PrintfBlue(".%s - ", fmt.Sprintf(sub, args...))

	for range onlyOnce {
		if state.IsOk() {
			ux.PrintfOk("PASSED - ")
			break
		}

		if state.IsError() {
			ux.PrintfError("FAILED - ")
			break
		}

		if state.IsWarning() {
			ux.PrintfError("FAILED - ")
			break
		}
	}

	r := state.SprintResponse()
	r = strings.TrimSpace(r)
	fmt.Printf("%s\n", r)
}

func PrintTestStop(test string) {
	ux.PrintflnBlue("# %s - STOPPED", test)
	ux.PrintflnWhite("################################################################################")
}


func main() {
	Test_NewState()

	os.Exit(1)

	Test_NewState()
	Test_toolRuntime()

	Test_loadTools()

	Test_toolDocker()
	Test_toolGear()

	Test_toolGit()
	Test_toolGitHub()
	//Test_GhrCopy()
	Test_Ghr()

	Test_toolPath()
	Test_toolPaths()
	Test_toolCopy()

	Test_toolExec()
	Test_NewMultiExec()

	Test_toolPrompt()
	Test_toolService()
	Test_toolSystem()
}


func Test_GhrCopy() {
	test := "toolGhrCopy"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Src := toolGhr.New(nil)
	PrintTestResult(Src.State, test, "New")
	state = Src.IsNil()
	PrintTestResult(state, test, "IsNil()")
	state = Src.SetAuth(toolGhr.TypeAuth{ Token: "", AuthUser: "" })
	PrintTestResult(state, test, "IsNil()")
	//state = Src.Set(toolGhr.TypeRepo{ Organization: "gearboxworks", Name: "launch" })
	state = Src.OpenUrl("https://github.com/newclarity/launch")
	PrintTestResult(state, test, "OpenUrl(\"https://github.com/newclarity/launch\")")


	Dest := toolGhr.New(nil)
	PrintTestResult(Dest.State, test, "New")
	state = Dest.IsNil()
	PrintTestResult(state, test, "IsNil()")
	state = Dest.OpenUrl("https://github.com/mickmake/test")
	PrintTestResult(state, test, "OpenUrl(\"https://github.com/mickmake/test\")")
	state = Dest.SetOverwrite(true)
	PrintTestResult(state, test, "SetOverwrite(true)")

	state = Dest.CopyReleasesFrom(Src.Repo, "dist")
	PrintTestResult(state, test, "CopyFrom(Src.Repo, \"dist\")")

	PrintTestStop(test)
}


func Test_Ghr() {
	test := "toolGhr"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := toolGhr.New(nil)
	PrintTestResult(Test.State, test, "New")

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	state = Test.OpenUrl("mickmake/test")
	PrintTestResult(state, test, "OpenUrl(\"mickmake/test\")")

	state = Test.OpenUrl("https://github.com/mickmake/test")
	PrintTestResult(state, test, "OpenUrl(\"https://github.com/mickmake/test\")")

	state = Test.SetTag("latest")
	PrintTestResult(state, test, "SetTag(\"latest\")")
	state = Test.Info()
	PrintTestResult(state, test, "Info()")

	state = Test.SetTag("1.0")
	PrintTestResult(state, test, "SetTag(\"1.0\")")
	state = Test.Info()
	PrintTestResult(state, test, "Info()")

	count := Test.Repo.CountReleases()
	PrintTestResult(state, test, "Repo.CountReleases() == %d", count)
	rels := Test.Repo.Releases()
	PrintTestResult(state, test, "rels.CountAll: %d", rels.CountAll())
	PrintTestResult(state, test, "len(rels.GetAll()): %d", len(rels.GetAll()))
	PrintTestResult(state, test, "rels.GetLatest: %v", rels.GetLatest())
	PrintTestResult(state, test, "rels.GetSelected: %v", rels.GetSelected())
	PrintTestResult(state, test, "rels.Sprint: %s", rels.Sprint())

	rel := Test.Repo.Release()
	PrintTestResult(state, test, "rels.Sprint: %v", rel)

	state = Test.Delete("1.0.1")
	PrintTestResult(state, test, "Create(\"1.0.1\", true)")

	state = Test.Create( toolGhr.TypeRepo{TagName: "1.0.1", Overwrite: true })
	PrintTestResult(state, test, "Create(\"1.0.1\", true)")

	//state = Test.Upload(true, "testing", "")
	//PrintTestResult(state, test, "Upload(\"testing\", \"\", true)")

	state = Test.UploadMultiple(true, "../testing/Testing2", "pkgreflect.go", "init.go")
	PrintTestResult(state, test, "Upload(\"testing\", \"\", true)")

	state = Test.Download(true, "testing2")
	PrintTestResult(state, test, "Download(\"testing\")")

	relData := toolGhr.TypeRepo{
		Organization: "mickmake",
		Name:         "test",
		TagName:      "2.0.0",
		Description:  "This is a description",
		Draft:        false,
		Prerelease:   false,
		Target:       "",
		Overwrite:    true,
		//Go:        []string{"../testing/testing", "pkgreflect.go", "init.go"},
		Auth:         &toolGhr.TypeAuth{ Token: "", AuthUser: "" },
	}
	state = Test.SetFilePath(".*\\.go", "../testing")
	PrintTestResult(state, test, "Download(\"testing\")")

	state = Test.Create(relData)
	PrintTestResult(state, test, "CreateRelease(relData)")

	state = Test.DeleteAssets("pkgreflect.go", "main.go")
	PrintTestResult(state, test, "CreateRelease(relData)")

	PrintTestStop(test)
}


func Test_NewState() {
	test := "ux"
	PrintTestStart(test)

	Test := ux.NewState("testing", globalDebug)
	PrintTestResult(Test, test, "NewState(\"testing\", %v)", globalDebug)
	var response *ux.TypeResponse


	ux.PrintflnWhite("################################################################################")
	ux.PrintflnBlue("# Testing Test.SetResponse(&[]string{\"one\", \"two\", \"\"})")
	foo := []string{"one", "two", ""}
	Test.SetResponse(&foo)
	response = Test.GetResponse()
	Test_PrintResponse(response, ux.TypeStringArray)


	ux.PrintflnWhite("################################################################################")
	ux.PrintflnBlue("# Testing Test.SetResponse([]string{\"one\", \"two\", \"\"})")
	Test.SetResponse(foo)
	response = Test.GetResponse()
	Test_PrintResponse(response, ux.TypeStringArray)


	ux.PrintflnWhite("################################################################################")
	ux.PrintflnBlue("# Testing Test.SetResponse(&*[]string)")
	var foo2 *[]string
	foo2 = &foo
	Test.SetResponse(&foo2)
	response = Test.GetResponse()
	Test_PrintResponse(response, ux.TypeStringArray)


	ux.PrintflnWhite("################################################################################")
	ux.PrintflnBlue("# Testing Test.SetResponse(&\"hello\")")
	foo3i := "hello"
	foo3 := &foo3i
	Test.SetResponse(foo3)
	response = Test.GetResponse()
	Test_PrintResponse(response, ux.TypeString)


	ux.PrintflnWhite("################################################################################")
	ux.PrintflnBlue("# Testing Test.SetResponse(Test_Func)")
	Test.SetResponse(Test_Func)
	response = Test.GetResponse()
	Test_RunFuncResponse(response, ux.TypeFunc, "")


	ux.PrintflnWhite("################################################################################")
	ux.PrintflnBlue("# Testing Test.SetResponse(Test_FuncVariadic)")
	Test.SetResponse(Test_FuncVariadic)
	response = Test.GetResponse()
	Test_RunFuncResponse(response, ux.TypeFuncVariadic, "", "arg1", "arg2")


	ux.PrintflnWhite("################################################################################")
	ux.PrintflnBlue("# Testing Test.SetResponse(Test_FuncReturn)")
	Test.SetResponse(Test_FuncReturn)
	response = Test.GetResponse()
	Test_RunFuncResponse(response, ux.TypeFuncReturn, ux.TypeString)


	ux.PrintflnWhite("################################################################################")
	ux.PrintflnBlue("# Testing Test.SetResponse(Test_FuncVariadicReturn)")
	Test.SetResponse(Test_FuncVariadicReturn)
	response = Test.GetResponse()
	Test_RunFuncResponse(response, ux.TypeFuncVariadicReturn, ux.TypeStringArray, "arg1", "arg2")


	ux.PrintflnWhite("################################################################################")
	ux.PrintflnBlue("# Testing Test.SetResponse(Test_Func)")
	Test.SetResponse(Test_Func)
	response = Test.GetResponse()
	response = response.AsFuncCall("arg1", "arg2")


	PrintTestStop(test)
}


func Test_PrintResponse(response *ux.TypeResponse, typeCheck string) {
	for range onlyOnce {
		if !response.Valid {
			ux.PrintfError("Response is nil.")
			break
		}

		responseType := response.GetType()
		responsePointer := response.Pointer()
		responseValue := response.Value()

		if response.IsOfType(typeCheck) {
			ux.PrintfOk("Type OK\n")
		} else {
			ux.PrintfError("Type NOT CORRECT: '%v' != '%v'\n", responseType, typeCheck)
		}

		ux.PrintflnGreen("response.GetType() - Name:%s\tString:%s\tKind:%s",
			response.GetType().Name(),
			response.GetType().String(),
			response.GetType().Kind(),
		)

		ux.PrintflnGreen("response.GetType() - Name:%s String:%s",
			responseType.Name(),
			responseType.String(),
		)

		ux.PrintflnGreen("Data returned:")
		ux.PrintflnYellow("responsePointer: %v", responsePointer)
		ux.PrintflnYellow("responseValue: %v", responseValue)
	}
}


func Test_RunFuncResponse(response *ux.TypeResponse, typeCheck string, returnCheck string, args ...interface{}) {
	for range onlyOnce {
		Test_PrintResponse(response, typeCheck)
		ux.PrintflnBlue("# Function response test of type '%s' with return of type '%s'", typeCheck, returnCheck)

		switch {
			case response.IsOfType(ux.TypeFunc):
				callerFunc := response.AsFunc()
				ux.PrintflnBlue("Execute AsFunc(): %v", callerFunc)
				callerFunc()

			case response.IsOfType(ux.TypeFuncReturn):
				callerFunc := response.AsFuncReturn()
				ux.PrintflnBlue("Execute AsFuncReturn(): %v", callerFunc)
				callerResponse := callerFunc()
				Test_PrintResponse(callerResponse, returnCheck)

			case response.IsOfType(ux.TypeFuncVariadic):
				callerFunc := response.AsFuncVariadic()
				ux.PrintflnBlue("Execute AsFuncVariadic(): %v", callerFunc)
				callerFunc(args...)

			case response.IsOfType(ux.TypeFuncVariadicReturn):
				callerFunc := response.AsFuncVariadicReturn()
				ux.PrintflnBlue("Execute AsFuncVariadicReturn(): %v", callerFunc)
				callerResponse := callerFunc(args...)
				Test_PrintResponse(callerResponse, returnCheck)

			default:
				return
		}

		ux.PrintflnBlue("Execute response.AsFuncCall():")
		callResponse := response.AsFuncCall(args)
		ux.PrintflnBlue("# Function returned")
		Test_PrintResponse(callResponse, returnCheck)

		//spew.Dump(response)
	}
}


func Test_Func() {
	ux.PrintflnGreen("Called Test_Func()")

	return
}

func Test_FuncVariadic(args ...interface{}) {
	var str string
	str += "Called Test_FuncVariadic("
	for _, a := range args {
		str += fmt.Sprintf("%s, ", a)
	}
	str += ")"
	ux.PrintflnGreen(str)
}

func Test_FuncReturn() *ux.TypeResponse {
	ux.PrintflnGreen("Called Test_FuncReturn()")

	ret := ux.NewResponse()
	str := "Called Test_FuncReturn()"
	ret.Set(str)
	return ret
}

func Test_FuncVariadicReturn(args ...interface{}) *ux.TypeResponse {
	ux.PrintflnGreen("Called Test_FuncVariadicReturn()")

	ret := ux.NewResponse()
	var str string
	str += "Called Test_FuncVariadicReturn("
	for _, a := range args {
		str += fmt.Sprintf("%s, ", a)
	}
	str += ")"

	ret.Set([]string{str})
	return ret
}


func Test_toolRuntime() {
	test := "toolRuntime"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := toolRuntime.New("test-harness", "1.0.0", globalDebug)
	PrintTestResult(Test.State, test, "New(\"test-harness\", \"1.0.0\", %v)", globalDebug)

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	PrintTestStop(test)
}


func Test_NewMultiExec() {
	test := "toolExec"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := toolExec.NewMultiExec(nil)
	PrintTestResult(Test.State, test, "NewMultiExec")

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	state = Test.ShowProgress()
	PrintTestResult(state, test, "ShowProgress()")

	state = Test.Set("ls", "-l", "-T")
	PrintTestResult(state, test, "Set(\"ls\", \"-l\", \"-T\")")

	state = Test.FindRegex(`go.mod`, "..")
	PrintTestResult(state, test, "FindRegex(`go.mod`, \"..\")")

	state = Test.Run()
	PrintTestResult(state, test, "Run")

	PrintTestStop(test)
}


func Test_loadTools() {
	test := "loadTools"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := loadTools.New("test-harness", "1.0.0", globalDebug)
	PrintTestResult(Test.State, test, "New(\"test-harness\", \"1.0.0\", %v)", globalDebug)

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	state = Test.ValidateArgs()
	PrintTestResult(state, test, "ValidateArgs")

	Test.PrintTools()

	PrintTestStop(test)
}


func Test_toolCopy() {
	test := "toolCopy"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := toolCopy.New(nil)
	PrintTestResult(Test.State, test, "New()")

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	PrintTestStop(test)
}


func Test_toolDocker() {
	test := "toolDocker"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := toolDocker.New(nil)
	PrintTestResult(Test.State, test, "New()")

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	PrintTestStop(test)
}


func Test_toolExec() {
	test := "toolExec"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := toolExec.New(nil)
	PrintTestResult(Test.State, test, "New()")

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	PrintTestStop(test)
}


func Test_toolGear() {
	test := "toolGear"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := toolGear.New(nil)
	PrintTestResult(Test.State, test, "New()")

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	PrintTestStop(test)
}


func Test_toolGit() {
	test := "toolGit"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := toolGit.New(nil)
	PrintTestResult(Test.State, test, "New()")

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	PrintTestStop(test)
}


func Test_toolGitHub() {
	test := "toolGitHub"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := toolGitHub.New(nil)
	PrintTestResult(Test.State, test, "New()")

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	PrintTestStop(test)
}


func Test_toolPath() {
	test := "toolPath"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := toolPath.New(nil)
	PrintTestResult(Test.State, test, "New()")

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	PrintTestStop(test)
}


func Test_toolPaths() {
	test := "toolPath"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := toolPath.NewPaths(nil)
	PrintTestResult(Test.State, test, "NewPaths()")

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	state = Test.FindRegex(`go.mod`, "..")
	PrintTestResult(Test.State, test, "FindRegex(`go.mod`, \"..\")")

	for _, v := range Test.Paths {
		fmt.Printf("%s\n", v.GetDirname())
	}

	PrintTestStop(test)
}


func Test_toolPrompt() {
	test := "toolPrompt"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := toolPrompt.New(nil)
	PrintTestResult(Test.State, test, "New()")

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	PrintTestStop(test)
}


func Test_toolService() {
	test := "toolService"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := toolService.New(nil)
	PrintTestResult(Test.State, test, "New()")

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	PrintTestStop(test)
}


func Test_toolSystem() {
	test := "toolSystem"
	state := ux.NewState(test, globalDebug)
	PrintTestStart(test)

	Test := toolSystem.New(nil)
	PrintTestResult(Test.State, test, "New()")

	state = Test.IsNil()
	PrintTestResult(state, test, "IsNil()")

	PrintTestStop(test)
}
