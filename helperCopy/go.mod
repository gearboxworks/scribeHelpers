module github.com/newclarity/scribeHelpers/helperCopy

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/helperPath => ../helperPath

replace github.com/newclarity/scribeHelpers/helperPrompt => ../helperPrompt

replace github.com/newclarity/scribeHelpers/helperTypes => ../helperTypes

require (
	github.com/newclarity/scribeHelpers/helperPath v0.0.0-20200602112526-02c317f22772
	github.com/newclarity/scribeHelpers/helperRuntime v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/helperTypes v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/ux v0.0.0-00010101000000-000000000000
	github.com/zloylos/grsync v0.0.0-20200204095520-71a00a7141be
// golang.org/x/crypto v0.0.0-20200602180216-279210d13fed // indirect
)
