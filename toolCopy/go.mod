module github.com/gearboxworks/scribeHelpers/toolCopy

go 1.14

replace github.com/gearboxworks/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/gearboxworks/scribeHelpers/toolPath => ../toolPath

replace github.com/gearboxworks/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/gearboxworks/scribeHelpers/toolTypes => ../toolTypes

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

require (
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/gearboxworks/scribeHelpers/toolPath v0.0.0-20200611181056-b2e5f7fd5978
	github.com/gearboxworks/scribeHelpers/toolPrompt v0.0.0-20200611181056-b2e5f7fd5978 // indirect
	github.com/gearboxworks/scribeHelpers/toolRuntime v0.0.0-20200611181056-b2e5f7fd5978
	github.com/gearboxworks/scribeHelpers/toolTypes v0.0.0-20200611181056-b2e5f7fd5978
	github.com/gearboxworks/scribeHelpers/ux v0.0.0-20200611181056-b2e5f7fd5978
	github.com/nsf/termbox-go v0.0.0-20200418040025-38ba6e5628f1 // indirect
	github.com/zloylos/grsync v0.0.0-20200204095520-71a00a7141be
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
	golang.org/x/sys v0.0.0-20200610111108-226ff32320da // indirect
)
