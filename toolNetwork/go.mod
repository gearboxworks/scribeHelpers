module github.com/gearboxworks/scribeHelpers/toolNetwork

go 1.14

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

replace github.com/gearboxworks/scribeHelpers/toolExec => ../toolExec
replace github.com/gearboxworks/scribeHelpers/toolPath => ../toolPath
replace github.com/gearboxworks/scribeHelpers/toolRuntime => ../toolRuntime
replace github.com/gearboxworks/scribeHelpers/toolTypes => ../toolTypes
replace github.com/gearboxworks/scribeHelpers/toolPrompt => ../toolPrompt

require (
	github.com/gearboxworks/scribeHelpers/ux v0.0.0
	github.com/gearboxworks/scribeHelpers/toolExec v0.0.0
	github.com/gearboxworks/scribeHelpers/toolPath v0.0.0
	github.com/gearboxworks/scribeHelpers/toolRuntime v0.0.0
	github.com/gearboxworks/scribeHelpers/toolTypes v0.0.0
	github.com/gearboxworks/scribeHelpers/toolPrompt v0.0.0
)
