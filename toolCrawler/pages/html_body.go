package pages

import (
	"github.com/newclarity/scribeHelpers/toolCrawler/global"
	"strings"
)

type HtmlBody global.Strings

func (me HtmlBody) String() string {
	return strings.Join(me, "\n")
}
