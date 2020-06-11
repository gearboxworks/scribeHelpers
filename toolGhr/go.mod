module github.com/newclarity/scribeHelpers/toolGhr

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/newclarity/scribeHelpers/toolPath => ../toolPath

replace github.com/newclarity/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/newclarity/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/inconshreveable/log15 v0.0.0-20200109203555-b30bc20e4fd1 // indirect
	github.com/kevinburke/rest v0.0.0-20200429221318-0d2892b400f8
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/newclarity/scribeHelpers/toolPath v0.0.0-20200606063537-e5c648daf391
	github.com/newclarity/scribeHelpers/toolPrompt v0.0.0-20200606063537-e5c648daf391 // indirect
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0-20200606063537-e5c648daf391
	github.com/newclarity/scribeHelpers/toolTypes v0.0.0-20200606063537-e5c648daf391 // indirect
	github.com/newclarity/scribeHelpers/ux v0.0.0-20200606063537-e5c648daf391
	github.com/tomnomnom/linkheader v0.0.0-20180905144013-02ca5825eb80
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
)
