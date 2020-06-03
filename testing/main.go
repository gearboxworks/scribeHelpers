package main

import (
	"github.com/newclarity/scribeHelpers/toolCopy"
	"github.com/newclarity/scribeHelpers/toolDocker"
	"github.com/newclarity/scribeHelpers/toolExec"
	"github.com/newclarity/scribeHelpers/toolGit"
	"github.com/newclarity/scribeHelpers/toolGitHub"
	"github.com/newclarity/scribeHelpers/toolPath"
	"github.com/newclarity/scribeHelpers/toolPrompt"
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/toolService"
	"github.com/newclarity/scribeHelpers/toolSystem"
	"github.com/newclarity/scribeHelpers/loadHelpers"
	"github.com/newclarity/scribeHelpers/ux"
)

/*
The sole purpose of this file's existence is for unit testing all the package modules.
Additionally, it helps within the GoLand IDE for entity mapping and checking.

This file is not to be used for any production related code.
*/

func main() {
	ux.PrintfBlue("ux.NewState - ")
	State := ux.NewState(true)
	if State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("ux.NewState - ")
	Helpers := loadHelpers.New("test-harness", "1.0.0", true)
	if Helpers.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("toolCopy.New - ")
	Copy := toolCopy.New(true)
	if Copy.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("toolDocker.New - ")
	Docker := toolDocker.New(true)
	if Docker.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("toolExec.New - ")
	Exec := toolExec.New(true)
	if Exec.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("toolGit.New - ")
	Git := toolGit.New(true)
	if Git.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("toolGitHub.New - ")
	Github := toolGitHub.New(true)
	if Github.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("toolPath.New - ")
	Path := toolPath.New(true)
	if Path.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("toolPrompt.New - ")
	Prompt := toolPrompt.New(true)
	if Prompt.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("toolRuntime.New - ")
	Runtime := toolRuntime.New("test-harness", "1.0.0", true)
	if Runtime.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("toolService.New - ")
	Service := toolService.New(true)
	if Service.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("toolSystem.New - ")
	System := toolSystem.New(true)
	if System.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}

}
