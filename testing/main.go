package main

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/toolCopy"
	"github.com/newclarity/scribeHelpers/toolDocker"
	"github.com/newclarity/scribeHelpers/toolExec"
	"github.com/newclarity/scribeHelpers/toolGear"
	"github.com/newclarity/scribeHelpers/toolGit"
	"github.com/newclarity/scribeHelpers/toolGitHub"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolPrompt"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/toolService"
	"github.com/newclarity/scribeHelpers/toolSystem"
	"github.com/newclarity/scribeHelpers/ux"
	"strings"
)

/*
The sole purpose of this file's existence is for unit testing all the package modules.
Additionally, it helps within the GoLand IDE for entity mapping and checking.

This file is not to be used for any production related code.
*/

const onlyOnce = "1"

var globalDebug bool

func main() {
	Test_NewState()
	Test_toolRuntime()

	Test_loadTools()

	Test_toolDocker()
	Test_toolGear()

	Test_toolGit()
	Test_toolGitHub()

	Test_toolPath()
	Test_toolPaths()
	Test_toolCopy()

	Test_toolExec()
	Test_NewMultiExec()

	Test_toolPrompt()
	Test_toolService()
	Test_toolSystem()
}


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


func Test_NewState() {
	test := "ux"
	PrintTestStart(test)

	Test := ux.NewState("testing", globalDebug)
	PrintTestResult(Test, test, "NewState(\"testing\", %v)", globalDebug)

	PrintTestStop(test)
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
