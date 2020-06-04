package main

import (
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/ux"
)

/*
The sole purpose of this file's existence is for unit testing all the package modules.
Additionally, it helps within the GoLand IDE for entity mapping and checking.

This file is not to be used for any production related code.
*/

func main() {
	ux.PrintfBlue("ux.NewState - ")
	State := ux.NewState("testing", true)
	if State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	//ux.PrintfBlue("toolRuntime.New - ")
	//Runtime := toolRuntime.New("test-harness", "1.0.0", true)
	//if Runtime.State.IsNotOk() {
	//	ux.PrintflnError("NOT OK")
	//} else {
	//	ux.PrintflnOk("OK")
	//}


	ux.PrintfBlue("loadTools.New - ")
	Tools := loadTools.New(nil)
	if Tools.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	//ux.PrintfBlue("toolCopy.New - ")
	//Copy := toolCopy.New(Runtime)
	//if Copy.State.IsNotOk() {
	//	ux.PrintflnError("NOT OK")
	//} else {
	//	ux.PrintflnOk("OK")
	//}
	//
	//
	//ux.PrintfBlue("toolDocker.New - ")
	//Docker := toolDocker.New(Runtime)
	//if Docker.State.IsNotOk() {
	//	ux.PrintflnError("NOT OK")
	//} else {
	//	ux.PrintflnOk("OK")
	//}
	//
	//
	//ux.PrintfBlue("toolExec.New - ")
	//Exec := toolExec.New(Runtime)
	//if Exec.State.IsNotOk() {
	//	ux.PrintflnError("NOT OK")
	//} else {
	//	ux.PrintflnOk("OK")
	//}
	//
	//
	//ux.PrintfBlue("toolGit.New - ")
	//Git := toolGit.New(Runtime)
	//if Git.State.IsNotOk() {
	//	ux.PrintflnError("NOT OK")
	//} else {
	//	ux.PrintflnOk("OK")
	//}
	//
	//
	//ux.PrintfBlue("toolGitHub.New - ")
	//Github := toolGitHub.New(Runtime)
	//if Github.State.IsNotOk() {
	//	ux.PrintflnError("NOT OK")
	//} else {
	//	ux.PrintflnOk("OK")
	//}
	//
	//
	//ux.PrintfBlue("toolPath.New - ")
	//Path := toolPath.New(Runtime)
	//if Path.State.IsNotOk() {
	//	ux.PrintflnError("NOT OK")
	//} else {
	//	ux.PrintflnOk("OK")
	//}
	//
	//
	//ux.PrintfBlue("toolPrompt.New - ")
	//Prompt := toolPrompt.New(Runtime)
	//if Prompt.State.IsNotOk() {
	//	ux.PrintflnError("NOT OK")
	//} else {
	//	ux.PrintflnOk("OK")
	//}
	//
	//
	//ux.PrintfBlue("toolService.New - ")
	//Service := toolService.New(Runtime)
	//if Service.State.IsNotOk() {
	//	ux.PrintflnError("NOT OK")
	//} else {
	//	ux.PrintflnOk("OK")
	//}
	//
	//
	//ux.PrintfBlue("toolSystem.New - ")
	//System := toolSystem.New(Runtime)
	//if System.State.IsNotOk() {
	//	ux.PrintflnError("NOT OK")
	//} else {
	//	ux.PrintflnOk("OK")
	//}

}
