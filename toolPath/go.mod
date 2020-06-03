module github.com/newclarity/scribeHelpers/toolPath

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/newclarity/scribeHelpers/toolTypes => ../toolTypes

replace github.com/newclarity/scribeHelpers/toolRuntime => ../toolRuntime

require (
	github.com/newclarity/scribeHelpers/toolPrompt v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/toolTypes v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/ux v0.0.0-00010101000000-000000000000
)
