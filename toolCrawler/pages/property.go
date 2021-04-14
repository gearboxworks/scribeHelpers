package pages

import "github.com/gearboxworks/scribeHelpers/toolCrawler/global"

type PropertyName = string
type PropertyMap map[PropertyName]*Property
type Properties []*Property
type Property struct {
	Name  PropertyName
	Value global.Content
}
