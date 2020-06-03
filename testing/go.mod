module testing

go 1.14

require (
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/go-openapi/strfmt v0.19.5 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/jedib0t/go-pretty v4.3.0+incompatible // indirect
	github.com/newclarity/scribeHelpers/helperCopy v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperDocker v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperExec v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/helperGit v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperGitHub v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperPath v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperPrompt v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperRuntime v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/helperService v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperSystem v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperTypes v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/helperUx v0.0.0-00010101000000-000000000000 // indirect
	github.com/newclarity/scribeHelpers/loadHelpers v0.0.0-00010101000000-000000000000
	github.com/newclarity/scribeHelpers/ux v0.0.0-00010101000000-000000000000
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/sftp v1.11.0 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	google.golang.org/grpc v1.29.1 // indirect
)

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20190717161051-705d9623b7c1

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/loadHelpers => ../loadHelpers

replace github.com/newclarity/scribeHelpers/helperCopy => ../helperCopy

replace github.com/newclarity/scribeHelpers/helperDocker => ../helperDocker

replace github.com/newclarity/scribeHelpers/helperExec => ../helperExec

replace github.com/newclarity/scribeHelpers/helperGit => ../helperGit

replace github.com/newclarity/scribeHelpers/helperGitHub => ../helperGitHub

replace github.com/newclarity/scribeHelpers/helperPath => ../helperPath

replace github.com/newclarity/scribeHelpers/helperPrompt => ../helperPrompt

replace github.com/newclarity/scribeHelpers/helperService => ../helperService

replace github.com/newclarity/scribeHelpers/helperSystem => ../helperSystem

replace github.com/newclarity/scribeHelpers/helperTypes => ../helperTypes

replace github.com/newclarity/scribeHelpers/helperUx => ../helperUx

replace github.com/newclarity/scribeHelpers/helperRuntime => ../helperRuntime
