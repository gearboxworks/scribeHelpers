module github.com/gearboxworks/scribeHelpers/toolGitHub

go 1.14

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

replace github.com/gearboxworks/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/gearboxworks/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/gearboxworks/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/google/go-github/v31 v31.0.0
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/gearboxworks/scribeHelpers/toolPrompt v0.0.0-20200604000029-dbb313f0fedc
	github.com/gearboxworks/scribeHelpers/toolRuntime v0.0.0-20200604000029-dbb313f0fedc
	github.com/gearboxworks/scribeHelpers/toolTypes v0.0.0-20200604000029-dbb313f0fedc
	github.com/gearboxworks/scribeHelpers/ux v0.0.0-20200604000029-dbb313f0fedc
	github.com/nsf/termbox-go v0.0.0-20200418040025-38ba6e5628f1 // indirect
	golang.org/x/crypto v0.0.0-20200602180216-279210d13fed // indirect
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980 // indirect
)
