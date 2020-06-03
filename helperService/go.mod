module github.com/newclarity/scribeHelpers/helperService

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux
replace github.com/newclarity/scribeHelpers/helperRuntime => ../helperRuntime

require (
	github.com/newclarity/scribeHelpers/helperPath v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/helperPrompt v0.0.0-20200603025545-971efd0cb59a // indirect
	github.com/newclarity/scribeHelpers/helperRuntime v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/helperTypes v0.0.0-20200603025545-971efd0cb59a // indirect
	golang.org/x/crypto v0.0.0-20200602180216-279210d13fed // indirect
)
