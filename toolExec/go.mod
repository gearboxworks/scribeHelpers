module github.com/newclarity/scribeHelpers/toolExec

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/newclarity/scribeHelpers/toolPath => ../toolPath

replace github.com/newclarity/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/newclarity/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/newclarity/scribeHelpers/toolPath v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/toolTypes v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/ux v0.0.0-20200604000029-dbb313f0fedc
)
