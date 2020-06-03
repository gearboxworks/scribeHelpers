module github.com/newclarity/scribeHelpers/toolSystem

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux
replace github.com/newclarity/scribeHelpers/toolRuntime => ../toolRuntime
replace github.com/newclarity/scribeHelpers/toolPath => ../toolPath
replace github.com/newclarity/scribeHelpers/toolPrompt => ../toolPrompt
replace github.com/newclarity/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/newclarity/scribeHelpers/toolPath v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/toolPrompt v0.0.0-20200603025545-971efd0cb59a // indirect
	github.com/newclarity/scribeHelpers/toolTypes v0.0.0-20200603025545-971efd0cb59a
	github.com/newclarity/scribeHelpers/ux v0.0.0-00010101000000-000000000000
	github.com/shirou/gopsutil v2.20.5+incompatible
	github.com/stretchr/testify v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20200602180216-279210d13fed // indirect
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980 // indirect
)
