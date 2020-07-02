package pages

import "github.com/newclarity/scribeHelpers/toolCrawler/global"

type PropertyName = string
type PropertyMap map[PropertyName]*Property
type Properties []*Property
type Property struct {
	Name  PropertyName
	Value global.Content
}
