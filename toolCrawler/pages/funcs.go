package pages

import (
	"github.com/gearboxworks/go-status/only"
	"github.com/gearboxworks/scribeHelpers/toolCrawler/global"
	"github.com/sirupsen/logrus"
	"net/url"
	"strings"
)

func noop(i ...interface{}) interface{} { return i }

func IsIndexable(u global.Url) (ok bool) {
	for range only.Once {
		if len(u) == 0 {
			break
		}
		uu, err := url.Parse(u)
		if err != nil {
			logrus.Warnf("unable to parse URL '%s'", u)
			break
		}
		if len(uu.Path) == 0 {
			break
		}
		if uu.Path[0] == '#' {
			break
		}
		if strings.Contains(uu.Path, "{") && strings.Contains(uu.Path, "}") {
			break
		}
		if uu.Scheme == "tel" {
			break
		}
		if uu.Scheme == "about" {
			break
		}
		if uu.Scheme == "mailto" {
			break
		}
		ok = true
	}
	return ok
}
