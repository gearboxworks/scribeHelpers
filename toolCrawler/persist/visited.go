package persist

import (
	"github.com/newclarity/scribeHelpers/toolCrawler/global"
)

type Visited struct {
	Id           SqlId
	ResourceHash Hash
	Timestamp    global.UnixTime
	Headers      string
	Body         string
	Cookies      string
}
