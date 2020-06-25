module github.com/newclarity/scribeHelpers/toolSelfUpdate

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/newclarity/scribeHelpers/toolPath => ../toolPath

replace github.com/newclarity/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/newclarity/scribeHelpers/toolTypes => ../toolTypes

replace github.com/newclarity/scribeHelpers/toolGhr => ../toolGhr

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/go-github/v30 v30.1.0
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0-20200623081955-45abb1cbefe9
	github.com/newclarity/scribeHelpers/ux v0.0.0-20200623081955-45abb1cbefe9
	github.com/rhysd/go-github-selfupdate v1.2.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/tcnksm/go-gitconfig v0.1.2
	github.com/ulikunitz/xz v0.5.7 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sys v0.0.0-20200622214017-ed371f2e16b4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)
