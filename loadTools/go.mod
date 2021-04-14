module github.com/gearboxworks/scribeHelpers/loadTools

go 1.14

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

replace github.com/gearboxworks/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/gearboxworks/scribeHelpers/toolCopy => ../toolCopy

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20190717161051-705d9623b7c1

replace github.com/gearboxworks/scribeHelpers/toolDocker => ../toolDocker

replace github.com/gearboxworks/scribeHelpers/toolExec => ../toolExec

replace github.com/gearboxworks/scribeHelpers/toolGear => ../toolGear

replace github.com/gearboxworks/scribeHelpers/toolGit => ../toolGit

replace github.com/gearboxworks/scribeHelpers/toolGitHub => ../toolGitHub

replace github.com/gearboxworks/scribeHelpers/toolPath => ../toolPath

replace github.com/gearboxworks/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/gearboxworks/scribeHelpers/toolService => ../toolService

replace github.com/gearboxworks/scribeHelpers/toolSystem => ../toolSystem

replace github.com/gearboxworks/scribeHelpers/toolTypes => ../toolTypes

replace github.com/gearboxworks/scribeHelpers/toolUx => ../toolUx

replace github.com/gearboxworks/scribeHelpers/toolGhr => ../toolGhr

//replace github.com/gearboxworks/scribeHelpers/toolCobraHelp => ../toolCobraHelp

require (
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/google/uuid v1.1.1 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/gearboxworks/scribeHelpers/toolCopy v0.0.0-00010101000000-000000000000
	github.com/gearboxworks/scribeHelpers/toolExec v0.0.0-20200604000029-dbb313f0fedc
	github.com/gearboxworks/scribeHelpers/toolGhr v0.0.0-00010101000000-000000000000
	github.com/gearboxworks/scribeHelpers/toolGit v0.0.0-00010101000000-000000000000
	github.com/gearboxworks/scribeHelpers/toolGitHub v0.0.0-00010101000000-000000000000
	github.com/gearboxworks/scribeHelpers/toolPath v0.0.0-20200621234507-ba6f08c6b68d
	github.com/gearboxworks/scribeHelpers/toolPrompt v0.0.0-20200621234507-ba6f08c6b68d
	github.com/gearboxworks/scribeHelpers/toolRuntime v0.0.0-20200621234507-ba6f08c6b68d
	github.com/gearboxworks/scribeHelpers/toolService v0.0.0-00010101000000-000000000000
	github.com/gearboxworks/scribeHelpers/toolSystem v0.0.0-00010101000000-000000000000
	github.com/gearboxworks/scribeHelpers/toolTypes v0.0.0-20200621234507-ba6f08c6b68d
	github.com/gearboxworks/scribeHelpers/toolUx v0.0.0-00010101000000-000000000000
	github.com/gearboxworks/scribeHelpers/ux v0.0.0-20200621234507-ba6f08c6b68d
	github.com/spf13/cobra v1.0.0
)
