module github.com/newclarity/scribeHelpers/toolDocker

go 1.14

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20190717161051-705d9623b7c1

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/toolGear => ../toolGear

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v0.0.0-00010101000000-000000000000
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dustin/go-humanize v1.0.0
	github.com/go-openapi/strfmt v0.19.5 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/jedib0t/go-pretty v4.3.0+incompatible
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/newclarity/scribeHelpers/ux v0.0.0-00010101000000-000000000000
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	google.golang.org/grpc v1.29.1 // indirect
	gotest.tools v2.2.0+incompatible // indirect
)
