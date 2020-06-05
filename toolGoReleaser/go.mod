module github.com/newclarity/scribeHelpers/toolGoReleaser

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/newclarity/scribeHelpers/toolPath => ../toolPath

replace github.com/newclarity/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/newclarity/scribeHelpers/toolExec => ../toolExec

replace github.com/newclarity/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/newclarity/scribeHelpers/toolExec v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/toolPath v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/ux v0.0.0-20200604000029-dbb313f0fedc
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
)
