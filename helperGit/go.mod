module github.com/newclarity/scribeHelpers/helperGit

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/helperExec => ../helperExec

replace github.com/newclarity/scribeHelpers/helperPath => ../helperPath

replace github.com/newclarity/scribeHelpers/helperPrompt => ../helperPrompt

replace github.com/newclarity/scribeHelpers/helperTypes => ../helperTypes

require (
	github.com/newclarity/scribeHelpers/helperExec v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/helperPath v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/helperRuntime v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/helperTypes v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/ux v0.0.0-00010101000000-000000000000
	github.com/tsuyoshiwada/go-gitcmd v0.0.0-20180205145712-5f1f5f9475df
	gopkg.in/src-d/go-git.v4 v4.13.1
)
