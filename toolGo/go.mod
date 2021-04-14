module github.com/gearboxworks/scribeHelpers/toolGo

go 1.14

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

replace github.com/gearboxworks/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/gearboxworks/scribeHelpers/toolPath => ../toolPath

replace github.com/gearboxworks/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/gearboxworks/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/fatih/astrewrite v0.0.0-20191207154002-9094e544fcef
	github.com/fatih/camelcase v1.0.0
	github.com/fatih/structtag v1.2.0
	github.com/gearboxworks/scribeHelpers/toolPath v0.0.0-20200612064705-ff77857fcb54
	github.com/gearboxworks/scribeHelpers/toolPrompt v0.0.0-20200612064705-ff77857fcb54 // indirect
	github.com/gearboxworks/scribeHelpers/toolRuntime v0.0.0-20200612064705-ff77857fcb54
	github.com/gearboxworks/scribeHelpers/toolTypes v0.0.0-20200612064705-ff77857fcb54 // indirect
	github.com/gearboxworks/scribeHelpers/ux v0.0.0-20200612064705-ff77857fcb54
	golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
)
