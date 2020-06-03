module github.com/newclarity/scribeHelpers/toolCopy

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/toolPath => ../toolPath

replace github.com/newclarity/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/newclarity/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/newclarity/scribeHelpers/toolPath v0.0.0-20200602112526-02c317f22772
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/toolTypes v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/ux v0.0.0-00010101000000-000000000000
	github.com/zloylos/grsync v0.0.0-20200204095520-71a00a7141be
// golang.org/x/crypto v0.0.0-20200602180216-279210d13fed // indirect
)
