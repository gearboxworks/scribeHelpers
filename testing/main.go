package main

import (
	"github.com/newclarity/scribeHelpers/helperCopy"
	"github.com/newclarity/scribeHelpers/helperDocker"
	"github.com/newclarity/scribeHelpers/helperExec"
	"github.com/newclarity/scribeHelpers/helperGit"
	"github.com/newclarity/scribeHelpers/helperGitHub"
	"github.com/newclarity/scribeHelpers/helperPath"
	"github.com/newclarity/scribeHelpers/helperPrompt"
	"github.com/newclarity/scribeHelpers/helperRuntime"
	"github.com/newclarity/scribeHelpers/helperService"
	"github.com/newclarity/scribeHelpers/helperSystem"
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


	ux.PrintfBlue("helperCopy.New - ")
	Copy := helperCopy.New(true)
	if Copy.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("helperDocker.New - ")
	Docker := helperDocker.New(true)
	if Docker.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("helperExec.New - ")
	Exec := helperExec.New(true)
	if Exec.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("helperGit.New - ")
	Git := helperGit.New(true)
	if Git.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("helperGitHub.New - ")
	Github := helperGitHub.New(true)
	if Github.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("helperPath.New - ")
	Path := helperPath.New(true)
	if Path.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("helperPrompt.New - ")
	Prompt := helperPrompt.New(true)
	if Prompt.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("helperRuntime.New - ")
	Runtime := helperRuntime.New("test-harness", "1.0.0", true)
	if Runtime.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("helperService.New - ")
	Service := helperService.New(true)
	if Service.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}


	ux.PrintfBlue("helperSystem.New - ")
	System := helperSystem.New(true)
	if System.State.IsNotOk() {
		ux.PrintflnError("NOT OK")
	} else {
		ux.PrintflnOk("OK")
	}

}
