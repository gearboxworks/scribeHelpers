module github.com/newclarity/scribeHelpers/toolService

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux
replace github.com/newclarity/scribeHelpers/toolRuntime => ../toolRuntime

require (
	github.com/newclarity/scribeHelpers/toolPath v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/toolPrompt v0.0.0-20200603025545-971efd0cb59a // indirect
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/toolTypes v0.0.0-20200603025545-971efd0cb59a // indirect
	golang.org/x/crypto v0.0.0-20200602180216-279210d13fed // indirect
)
