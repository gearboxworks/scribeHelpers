package hosters

import "github.com/newclarity/scribeHelpers/toolCrawler/pages"

type IndexHoster interface {
	Initialize() error
	IndexPage(*pages.Page) bool
}
