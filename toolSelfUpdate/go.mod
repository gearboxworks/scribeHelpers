module github.com/gearboxworks/scribeHelpers/toolSelfUpdate

go 1.14

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

replace github.com/gearboxworks/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/gearboxworks/scribeHelpers/toolPath => ../toolPath

replace github.com/gearboxworks/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/gearboxworks/scribeHelpers/toolTypes => ../toolTypes

replace github.com/gearboxworks/scribeHelpers/toolGhr => ../toolGhr

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/gearboxworks/scribeHelpers/toolRuntime v0.0.0-20200623081955-45abb1cbefe9
	github.com/gearboxworks/scribeHelpers/ux v0.0.0-20200623081955-45abb1cbefe9
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/go-github/v30 v30.1.0
	github.com/rhysd/go-github-selfupdate v1.2.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/tcnksm/go-gitconfig v0.1.2
	github.com/ulikunitz/xz v0.5.7 // indirect
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)
