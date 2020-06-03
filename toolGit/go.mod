module github.com/newclarity/scribeHelpers/toolGit

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/toolExec => ../toolExec

replace github.com/newclarity/scribeHelpers/toolPath => ../toolPath

replace github.com/newclarity/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/newclarity/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/newclarity/scribeHelpers/toolExec v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/toolPath v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/toolTypes v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/ux v0.0.0-00010101000000-000000000000
	github.com/tsuyoshiwada/go-gitcmd v0.0.0-20180205145712-5f1f5f9475df
	gopkg.in/src-d/go-git.v4 v4.13.1
)
