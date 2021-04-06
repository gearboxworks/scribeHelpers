module github.com/newclarity/scribeHelpers/toolNetwork

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/toolExec => ../toolExec
replace github.com/newclarity/scribeHelpers/toolPath => ../toolPath
replace github.com/newclarity/scribeHelpers/toolRuntime => ../toolRuntime
replace github.com/newclarity/scribeHelpers/toolTypes => ../toolTypes
replace github.com/newclarity/scribeHelpers/toolPrompt => ../toolPrompt

require (
	github.com/newclarity/scribeHelpers/ux v0.0.0
	github.com/newclarity/scribeHelpers/toolExec v0.0.0
	github.com/newclarity/scribeHelpers/toolPath v0.0.0
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0
	github.com/newclarity/scribeHelpers/toolTypes v0.0.0
	github.com/newclarity/scribeHelpers/toolPrompt v0.0.0
)
