module github.com/newclarity/scribeHelpers/helperPath

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/helperPrompt => ../helperPrompt

replace github.com/newclarity/scribeHelpers/helperTypes => ../helperTypes

replace github.com/newclarity/scribeHelpers/helperRuntime => ../helperRuntime

require (
	github.com/newclarity/scribeHelpers/helperPrompt v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/helperRuntime v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/helperTypes v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/ux v0.0.0-00010101000000-000000000000
)
