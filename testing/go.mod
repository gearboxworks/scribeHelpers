module testing

go 1.14

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20190717161051-705d9623b7c1

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

replace github.com/gearboxworks/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/gearboxworks/scribeHelpers/loadTools => ../loadTools

replace github.com/gearboxworks/scribeHelpers/toolCopy => ../toolCopy

replace github.com/gearboxworks/scribeHelpers/toolDocker => ../toolDocker

replace github.com/gearboxworks/scribeHelpers/toolExec => ../toolExec

replace github.com/gearboxworks/scribeHelpers/toolGit => ../toolGit

replace github.com/gearboxworks/scribeHelpers/toolGear => ../toolGear

replace github.com/gearboxworks/scribeHelpers/toolGitHub => ../toolGitHub

replace github.com/gearboxworks/scribeHelpers/toolPath => ../toolPath

replace github.com/gearboxworks/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/gearboxworks/scribeHelpers/toolService => ../toolService

replace github.com/gearboxworks/scribeHelpers/toolSystem => ../toolSystem

replace github.com/gearboxworks/scribeHelpers/toolTypes => ../toolTypes

replace github.com/gearboxworks/scribeHelpers/toolGhr => ../toolGhr

replace github.com/gearboxworks/scribeHelpers/toolUx => ../toolUx

require (
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/gearboxworks/scribeHelpers/loadTools v0.0.0-00010101000000-000000000000 // indirect
	github.com/gearboxworks/scribeHelpers/toolDocker v0.0.0-00010101000000-000000000000 // indirect
	github.com/gearboxworks/scribeHelpers/toolGear v0.0.0-00010101000000-000000000000 // indirect
	github.com/gearboxworks/scribeHelpers/toolGhr v0.0.0-00010101000000-000000000000 // indirect
)
