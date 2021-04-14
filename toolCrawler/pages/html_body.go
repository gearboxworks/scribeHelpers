package pages

import (
	"github.com/gearboxworks/scribeHelpers/toolCrawler/global"
	"strings"
)

type HtmlBody global.Strings

func (me HtmlBody) String() string {
	return strings.Join(me, "\n")
}
