module github.com/gearboxworks/scribeHelpers/toolGhr

go 1.14

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

replace github.com/gearboxworks/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/gearboxworks/scribeHelpers/toolPath => ../toolPath

replace github.com/gearboxworks/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/gearboxworks/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/inconshreveable/log15 v0.0.0-20200109203555-b30bc20e4fd1 // indirect
	github.com/kevinburke/rest v0.0.0-20200429221318-0d2892b400f8
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/gearboxworks/scribeHelpers/toolPath v0.0.0-20200606063537-e5c648daf391
	github.com/gearboxworks/scribeHelpers/toolPrompt v0.0.0-20200606063537-e5c648daf391 // indirect
	github.com/gearboxworks/scribeHelpers/toolRuntime v0.0.0-20200606063537-e5c648daf391
	github.com/gearboxworks/scribeHelpers/toolTypes v0.0.0-20200606063537-e5c648daf391 // indirect
	github.com/gearboxworks/scribeHelpers/ux v0.0.0-20200606063537-e5c648daf391
	github.com/tomnomnom/linkheader v0.0.0-20180905144013-02ca5825eb80
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
)
