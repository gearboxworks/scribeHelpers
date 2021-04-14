module github.com/gearboxworks/scribeHelpers/toolService

go 1.14

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

replace github.com/gearboxworks/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/gearboxworks/scribeHelpers/toolPath => ../toolPath

replace github.com/gearboxworks/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/gearboxworks/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/gearboxworks/scribeHelpers/toolPath v0.0.0-20200604000029-dbb313f0fedc
	github.com/gearboxworks/scribeHelpers/toolRuntime v0.0.0-20200604000029-dbb313f0fedc
	github.com/gearboxworks/scribeHelpers/ux v0.0.0-20200604000029-dbb313f0fedc
)
