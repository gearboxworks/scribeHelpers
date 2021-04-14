package toolCrawler

import "github.com/gearboxworks/scribeHelpers/ux"


type ToolCrawler TypeCrawler
func (e *ToolCrawler) IsNil() *ux.State {
	return ux.IfNilReturnError(e)
}

func (e *ToolCrawler) Reflect() *TypeCrawler {
	return (*TypeCrawler)(e)
}

func (e *TypeCrawler) Reflect() *ToolCrawler {
	return (*ToolCrawler)(e)
}


type State ux.State

func (s *State) Reflect() *ux.State {
	return (*ux.State)(s)
}

func ReflectState(p *ux.State) *State {
	return (*State)(p)
}

func ReflectToolCrawler(e *TypeCrawler) *ToolCrawler {
	return (*ToolCrawler)(e)
}
