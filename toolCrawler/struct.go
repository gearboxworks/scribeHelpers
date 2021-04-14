package toolCrawler

import (
	"github.com/gearboxworks/scribeHelpers/toolPath"
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
)

type CrawlerGetter interface {
}


type TypeCrawler struct {
	name    string
	path    *toolPath.TypeOsPath

	runtime *toolRuntime.TypeRuntime
	State   *ux.State
}
func (e *TypeCrawler) IsNil() *ux.State {
	return ux.IfNilReturnError(e)
}


func New(runtime *toolRuntime.TypeRuntime) *TypeCrawler {
	runtime = runtime.EnsureNotNil()

	te := TypeCrawler{
		name:    "",
		path:    toolPath.New(runtime),

		runtime: runtime,
		State:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
	te.State.SetPackage("")
	te.State.SetFunctionCaller()
	return &te
}

