package hosters

import "github.com/gearboxworks/scribeHelpers/toolCrawler/pages"

type IndexHoster interface {
	Initialize() error
	IndexPage(*pages.Page) bool
}
