module github.com/gearboxworks/scribeHelpers/toolGear

go 1.14

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20190717161051-705d9623b7c1

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

replace github.com/gearboxworks/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/gearboxworks/scribeHelpers/toolGear/gearConfig => ./gearConfig

replace github.com/gearboxworks/scribeHelpers/toolPath => ../toolPath

replace github.com/gearboxworks/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/gearboxworks/scribeHelpers/toolTypes => ../toolTypes

replace github.com/gearboxworks/scribeHelpers/toolNetwork => ../toolNetwork

replace github.com/gearboxworks/scribeHelpers/toolExec => ../toolExec


require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200428143746-21a406dcc535 // indirect
	//github.com/cakturk/go-netstat v0.0.0-20200220111822-e5b49efee7a5
	github.com/cavaliercoder/grab v2.0.0+incompatible
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dustin/go-humanize v1.0.0
	github.com/fatih/color v1.9.0
	github.com/go-openapi/errors v0.19.4 // indirect
	github.com/go-openapi/strfmt v0.19.5 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jedib0t/go-pretty v4.3.0+incompatible
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mitchellh/mapstructure v1.3.1
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/gearboxworks/scribeHelpers/toolNetwork v0.0.0
	github.com/gearboxworks/scribeHelpers/toolPath v0.0.0-20200604000029-dbb313f0fedc
	github.com/gearboxworks/scribeHelpers/toolRuntime v0.0.0-20200604000029-dbb313f0fedc
	github.com/gearboxworks/scribeHelpers/ux v0.0.0-20200604000029-dbb313f0fedc
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/sftp v1.11.0
	github.com/sirupsen/logrus v1.6.0 // indirect
	go.mongodb.org/mongo-driver v1.3.4 // indirect
	golang.org/x/crypto v0.0.0-20200602180216-279210d13fed
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	google.golang.org/genproto v0.0.0-20200603110839-e855014d5736 // indirect
	google.golang.org/grpc v1.29.1 // indirect
	gotest.tools v2.2.0+incompatible // indirect
)
