module github.com/newclarity/scribeHelpers/loadTools

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/newclarity/scribeHelpers/toolCopy => ../toolCopy

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20190717161051-705d9623b7c1

replace github.com/newclarity/scribeHelpers/toolDocker => ../toolDocker

replace github.com/newclarity/scribeHelpers/toolExec => ../toolExec

replace github.com/newclarity/scribeHelpers/toolGear => ../toolGear

replace github.com/newclarity/scribeHelpers/toolGit => ../toolGit

replace github.com/newclarity/scribeHelpers/toolGitHub => ../toolGitHub

replace github.com/newclarity/scribeHelpers/toolPath => ../toolPath

replace github.com/newclarity/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/newclarity/scribeHelpers/toolService => ../toolService

replace github.com/newclarity/scribeHelpers/toolSystem => ../toolSystem

replace github.com/newclarity/scribeHelpers/toolTypes => ../toolTypes

replace github.com/newclarity/scribeHelpers/toolUx => ../toolUx

replace github.com/newclarity/scribeHelpers/toolGhr => ../toolGhr

replace github.com/newclarity/scribeHelpers/toolCobraHelp => ../toolCobraHelp

require (
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/google/uuid v1.1.1 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/newclarity/scribeHelpers/toolCobraHelp v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolCopy v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolExec v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/toolGhr v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolGit v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolGitHub v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolPath v0.0.0-20200621234507-ba6f08c6b68d
	github.com/newclarity/scribeHelpers/toolPrompt v0.0.0-20200621234507-ba6f08c6b68d
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0-20200621234507-ba6f08c6b68d
	github.com/newclarity/scribeHelpers/toolService v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolSystem v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/toolTypes v0.0.0-20200621234507-ba6f08c6b68d
	github.com/newclarity/scribeHelpers/toolUx v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/ux v0.0.0-20200621234507-ba6f08c6b68d
	github.com/spf13/cobra v1.0.0
)
