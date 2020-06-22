module github.com/newclarity/scribeHelpers/toolCobraHelp

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/newclarity/scribeHelpers/toolPath => ../toolPath

replace github.com/newclarity/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/newclarity/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/newclarity/scribeHelpers/toolPath v0.0.0-20200621234507-ba6f08c6b68d
	github.com/newclarity/scribeHelpers/toolPrompt v0.0.0-20200621234507-ba6f08c6b68d // indirect
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0-20200621234507-ba6f08c6b68d
	github.com/newclarity/scribeHelpers/toolTypes v0.0.0-20200621234507-ba6f08c6b68d // indirect
	github.com/newclarity/scribeHelpers/ux v0.0.0-20200621234507-ba6f08c6b68d
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
	golang.org/x/sys v0.0.0-20200620081246-981b61492c35 // indirect
	golang.org/x/text v0.3.3 // indirect
)
