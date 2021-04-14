package toolPrompt

import (
	"fmt"
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"

	// "golang.org/x/crypto/ssh/terminal"
)

// @TODO - This is a workaround for the duplicates that appear with pkgreflect.

type TypePrompt struct {
	 string

	 Debug bool
	 State *ux.State
}


func UserPrompt(prompt string, args ...interface{}) string {
	var p TypePrompt
	p.Set(prompt, args...)
	return p.UserPrompt()
}
func (p *TypePrompt) Set(prompt string, args ...interface{}) {
	p.string = fmt.Sprintf(prompt, args...)
}


func UserPromptHidden(prompt string, args ...interface{}) string {
	var p TypePrompt
	p.Set(prompt, args)
	return p.UserPromptHidden()
}


// Dummy New()
func New(runtime *toolRuntime.TypeRuntime) *TypePrompt {
	runtime = runtime.EnsureNotNil()

	p := TypePrompt{
		string: "",
		Debug:  runtime.Debug,
		State:  ux.NewState(runtime.CmdName, runtime.Debug),
	}
	p.State.SetPackage("")
	p.State.SetFunctionCaller()
	return &p
}


func (p *TypePrompt) IsNil() *ux.State {
	if state := ux.IfNilReturnError(p); state.IsError() {
		return state
	}
	p.State = p.State.EnsureNotNil()
	return p.State
}


func (p *TypePrompt) EnsureNotNil() *TypePrompt {
	if p == nil {
		return New(nil)
	}
	return p
}
