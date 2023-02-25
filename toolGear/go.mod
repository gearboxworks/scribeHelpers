module github.com/gearboxworks/scribeHelpers/toolGear

go 1.14

// replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20190717161051-705d9623b7c1
// replace github.com/docker/docker => github.com/docker/engine v1.13.1
// replace github.com/docker/docker => github.com/docker/engine v17.12.0-ce-rc1.0.20200916142827-bd33bbf0497b+incompatible

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

replace github.com/gearboxworks/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/gearboxworks/scribeHelpers/toolGear/gearConfig => ./gearConfig

replace github.com/gearboxworks/scribeHelpers/toolPath => ../toolPath

replace github.com/gearboxworks/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/gearboxworks/scribeHelpers/toolTypes => ../toolTypes

replace github.com/gearboxworks/scribeHelpers/toolExec => ../toolExec

replace github.com/gearboxworks/scribeHelpers/toolNetwork => ../toolNetwork

require (
	github.com/Microsoft/hcsshim v0.8.17 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200428143746-21a406dcc535 // indirect
	//github.com/cakturk/go-netstat v0.0.0-20200220111822-e5b49efee7a5
	github.com/cavaliercoder/grab v2.0.0+incompatible
	github.com/containerd/containerd v1.5.2 // indirect
	github.com/docker/docker v20.10.6+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/dustin/go-humanize v1.0.0
	github.com/fatih/color v1.9.0
	github.com/gearboxworks/scribeHelpers/toolNetwork v0.0.0-00010101000000-000000000000
	github.com/gearboxworks/scribeHelpers/toolPath v0.0.0
	github.com/gearboxworks/scribeHelpers/toolRuntime v0.0.0
	github.com/gearboxworks/scribeHelpers/ux v0.0.0
	github.com/go-openapi/errors v0.19.4 // indirect
	github.com/go-openapi/strfmt v0.19.5 // indirect
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jedib0t/go-pretty v4.3.0+incompatible
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mitchellh/mapstructure v1.3.1
	github.com/moby/sys/mount v0.2.0 // indirect
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/pkg/sftp v1.11.0
	github.com/sirupsen/logrus v1.8.1 // indirect
	go.mongodb.org/mongo-driver v1.3.4 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/net v0.7.0
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
)
