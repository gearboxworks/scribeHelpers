module github.com/gearboxworks/scribeHelpers/toolCrawler

go 1.14

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

replace github.com/gearboxworks/scribeHelpers/toolRuntime => ../toolRuntime

replace github.com/gearboxworks/scribeHelpers/toolPath => ../toolPath

replace github.com/gearboxworks/scribeHelpers/toolPrompt => ../toolPrompt

replace github.com/gearboxworks/scribeHelpers/toolTypes => ../toolTypes

require (
	github.com/algolia/algoliasearch-client-go v2.25.0+incompatible
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.2.4 // indirect
	github.com/gearboxworks/go-status v0.0.0-20190623205420-467f07bd7e0e
	github.com/gearboxworks/scribeHelpers/toolPath v0.0.0-20200701071225-c7db504f92c9
	github.com/gearboxworks/scribeHelpers/toolPrompt v0.0.0-20200701071225-c7db504f92c9 // indirect
	github.com/gearboxworks/scribeHelpers/toolRuntime v0.0.0-20200701071225-c7db504f92c9
	github.com/gearboxworks/scribeHelpers/toolTypes v0.0.0-20200701071225-c7db504f92c9 // indirect
	github.com/gearboxworks/scribeHelpers/ux v0.0.0-20200701071225-c7db504f92c9
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly v1.2.0
	github.com/hashicorp/golang-lru v0.5.4
	github.com/jtacoma/uritemplates v1.0.0
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/mattn/go-sqlite3 v1.14.0
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/temoto/robotstxt v1.1.1 // indirect
	golang.org/x/net v0.7.0
)
