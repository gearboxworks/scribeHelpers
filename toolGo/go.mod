module github.com/newclarity/scribeHelpers/toolGo

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/newclarity/scribeHelpers/toolPath => ../toolPath

replace github.com/newclarity/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/newclarity/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/fatih/astrewrite v0.0.0-20191207154002-9094e544fcef
	github.com/fatih/camelcase v1.0.0
	github.com/fatih/structtag v1.2.0
	github.com/newclarity/scribeHelpers/toolPath v0.0.0-20200612064705-ff77857fcb54
	github.com/newclarity/scribeHelpers/toolPrompt v0.0.0-20200612064705-ff77857fcb54 // indirect
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0-20200612064705-ff77857fcb54
	github.com/newclarity/scribeHelpers/toolTypes v0.0.0-20200612064705-ff77857fcb54 // indirect
	github.com/newclarity/scribeHelpers/ux v0.0.0-20200612064705-ff77857fcb54
	golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
)
