module github.com/gearboxworks/scribeHelpers/toolCobraHelp

go 1.14

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

replace github.com/gearboxworks/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/gearboxworks/scribeHelpers/toolPath => ../toolPath

replace github.com/gearboxworks/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/gearboxworks/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/gearboxworks/scribeHelpers/toolRuntime v0.0.0-20200621234507-ba6f08c6b68d
	github.com/gearboxworks/scribeHelpers/ux v0.0.0-20200621234507-ba6f08c6b68d
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.3.3 // indirect
)
