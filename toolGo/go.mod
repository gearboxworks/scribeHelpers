module github.com/newclarity/scribeHelpers/toolGo

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/newclarity/scribeHelpers/toolPath => ../toolPath

replace github.com/newclarity/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/newclarity/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/newclarity/scribeHelpers/toolPath v0.0.0-20200612064705-ff77857fcb54
	github.com/newclarity/scribeHelpers/toolPrompt v0.0.0-20200612064705-ff77857fcb54 // indirect
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0-20200612064705-ff77857fcb54
	github.com/newclarity/scribeHelpers/toolTypes v0.0.0-20200612064705-ff77857fcb54 // indirect
	github.com/newclarity/scribeHelpers/ux v0.0.0-20200612064705-ff77857fcb54
)
