module github.com/newclarity/scribeHelpers/helperGitHub

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/helperPrompt => ../helperPrompt

replace github.com/newclarity/scribeHelpers/helperTypes => ../helperTypes

require (
	github.com/google/go-github/v31 v31.0.0
	github.com/newclarity/scribeHelpers/helperPrompt v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/helperRuntime v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/helperTypes v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/ux v0.0.0-00010101000000-000000000000
)
