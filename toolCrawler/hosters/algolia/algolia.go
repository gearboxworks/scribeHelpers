package algolia

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/gearboxworks/scribeHelpers/toolCrawler/config"
	"github.com/gearboxworks/scribeHelpers/toolCrawler/global"
	"github.com/gearboxworks/scribeHelpers/toolCrawler/hosters"
	"github.com/gearboxworks/scribeHelpers/toolCrawler/pages"
	"github.com/sirupsen/logrus"
)

var _ hosters.IndexHoster = (*Algolia)(nil)

type Algolia struct {
	Index  global.Index
	Client algoliasearch.Client
	Config *config.Config
}

func NewAlgolia(c *config.Config) *Algolia {
	a := Algolia{}
	a.Config = c
	client := algoliasearch.NewClient(c.AppId, c.ApiKey)
	a.Index = client.InitIndex(c.IndexName)
	return &a
}

func (me *Algolia) Initialize() error {
	settings := algoliasearch.Map{
		"searchableAttributes": me.Config.SearchAttrs,
	}
	_, err := me.Index.SetSettings(settings)
	if err != nil {
		logrus.Fatalf("Unable to set index settings: %s", err)
	}
	return err
}

func (me *Algolia) IndexPage(p *pages.Page) bool {
	var err error
	//for range only.Once {
	//	o := hosters.NewObject(global.Object{
	//		"objectID": p.Id.String(),
	//		"title":    p.Title,
	//		"url":      p.Url,
	//		"body":     p.Body.String(),
	//	})
	//
	//	o.AppendProps(p.HeaderMap.ExtractStringMap())
	//
	//	o.AppendProps(p.ElementsMap.ExtractStringMap())
	//
	//	ps := me.Config.UrlPatterns.ExtractStringMap(p.Url)
	//	o.AppendProps(ps)
	//	_, err = me.Index.AddObject(o.Object)
	//	if err != nil {
	//		me.Config.OnFailedVisit(err, p.Url, "adding page to index")
	//		break
	//	}
	//}
	return err == nil
}

func noop(i ...interface{}) interface{} { return i }
