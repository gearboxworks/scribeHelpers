package cmd

import (
	"github.com/gearboxworks/scribeHelpers/toolCrawler/global"
	"github.com/spf13/cobra"
)

// @see https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery/

var RootCmd = &cobra.Command{
	Use:   "website-indexer",
	Short: "Populate an Algolia or Elastic App Search index for a given website domain.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return err
	},
}

func init() {
	pf := RootCmd.PersistentFlags()
	pf.BoolVarP(&global.NoCache, "no-cache", "", false, "Disable caching")
}
