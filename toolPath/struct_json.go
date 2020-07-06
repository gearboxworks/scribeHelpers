package toolPath

import "encoding/json"


func (p *TypeOsPath) MarshalJSON() ([]byte, error) {
	var ret []byte
	var err error
	for range onlyOnce {
		type Alias TypeOsPath

		ret, err = json.Marshal(&struct {
			Path string `json:"path"`
			*Alias
		}{
			Path:  p.GetPath(),
			Alias: (*Alias)(p),
		})

		if err != nil {
			p.State.SetError(err)
			break
		}

		p.State.SetOk()
	}

	return ret, err
}

// Since aux is a pointer to a struct, you don't need to do Unmarshal(data, &aux). Drop one of the two ampersands (probably the first one) to remove an unneeded indirection.
// See: https://play.golang.org/p/CTmJhTiGAM


func (p *TypeOsPath) UnmarshalJSON(data []byte) error {
	var err error
	for range onlyOnce {
		type Alias TypeOsPath

		aux := &struct {
			Path string `json:"path"`
			*Alias
		}{
			Alias: (*Alias)(p),
		}

		err = json.Unmarshal(data, &aux)
		if err != nil {
			p.State.SetError(err)
			break
		}

		p.SetPath(aux.Path)
	}

	return err
}
