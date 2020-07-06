package toolGit

import (
	"encoding/json"
)


func (g *TypeGit) MarshalJSON() ([]byte, error) {
	var ret []byte
	var err error
	for range onlyOnce {
		type Alias TypeGit

		ret, err = json.Marshal(&struct {
			Url string `json:"url"`
			Base string `json:"path"`
			*Alias
		}{
			Url:   g.Url.String(),
			Base:   g.Base.GetPath(),
			Alias: (*Alias)(g),
		})

		if err != nil {
			g.State.SetError(err)
			break
		}

		g.State.SetOk()
	}

	return ret, err
}

// Since aux is a pointer to a struct, you don't need to do Unmarshal(data, &aux). Drop one of the two ampersands (probably the first one) to remove an unneeded indirection.
// See: https://play.golang.org/p/CTmJhTiGAM


func (g *TypeGit) UnmarshalJSON(data []byte) error {
	var err error
	for range onlyOnce {
		type Alias TypeGit

		aux := &struct {
			Url string `json:"url"`
			Base string `json:"path"`
			*Alias
		}{
			Alias: (*Alias)(g),
		}

		err = json.Unmarshal(data, &aux)
		if err != nil {
			g.State.SetError(err)
			break
		}

		g.State = g.SetUrl(aux.Url)
		if g.State.IsNotOk() {
			err = g.State.GetError()
			break
		}

		g.State = g.SetPath(aux.Base)
		if g.State.IsNotOk() {
			err = g.State.GetError()
			break
		}
	}

	return err
}
