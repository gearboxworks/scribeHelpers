module github.com/newclarity/scribeHelpers/toolGit

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux
replace github.com/newclarity/scribeHelpers/toolRuntime => ../toolRuntime
replace github.com/newclarity/scribeHelpers/toolExec => ../toolExec
replace github.com/newclarity/scribeHelpers/toolPath => ../toolPath
replace github.com/newclarity/scribeHelpers/toolTypes => ../toolTypes
replace github.com/newclarity/scribeHelpers/toolPrompt => ../toolPrompt

require (
	github.com/newclarity/scribeHelpers/toolExec v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/toolPath v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/toolRuntime v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/toolTypes v0.0.0-20200604000029-dbb313f0fedc
	github.com/newclarity/scribeHelpers/ux v0.0.0-20200604000029-dbb313f0fedc
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/tsuyoshiwada/go-gitcmd v0.0.0-20180205145712-5f1f5f9475df
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	gopkg.in/src-d/go-git.v4 v4.13.1
)
