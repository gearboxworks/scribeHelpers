module github.com/gearboxworks/scribeHelpers/toolGoReleaser

go 1.14

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

replace github.com/gearboxworks/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/gearboxworks/scribeHelpers/toolPath => ../toolPath

replace github.com/gearboxworks/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/gearboxworks/scribeHelpers/toolExec => ../toolExec

replace github.com/gearboxworks/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/gearboxworks/scribeHelpers/toolExec v0.0.0-20200604000029-dbb313f0fedc
	github.com/gearboxworks/scribeHelpers/toolPath v0.0.0-20200604000029-dbb313f0fedc
	github.com/gearboxworks/scribeHelpers/toolRuntime v0.0.0-20200604000029-dbb313f0fedc
	github.com/gearboxworks/scribeHelpers/ux v0.0.0-20200604000029-dbb313f0fedc
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
)
