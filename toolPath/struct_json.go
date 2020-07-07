package toolPath

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)


func (p *TypeOsPath) MarshalJSON() ([]byte, error) {
	var ret []byte
	var err error
	for range onlyOnce {
		type Alias TypeOsPath

		ret, err = json.Marshal(&struct {
			Path string `json:"path"`
			//Overwrite bool `json:"overwrite"`
			*Alias
		}{
			Path:  p.GetPath(),
			//Overwrite:  p._CanOverwrite,
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
		//type Alias TypeOsPath
		//
		//aux := &struct {
		//	Path string `json:"path"`
		//	*Alias
		//}{
		//	//Path:  p.GetPath(),
		//	Alias: (*Alias)(p),
		//}
		//
		//err = json.Unmarshal(data, &aux)
		//if err != nil {
		//	p.State.SetError(err)
		//	//break
		//}
		//
		//p.SetPath(aux.Path)

		path := string(data)
		path = strings.TrimPrefix(path, "\"")
		path = strings.TrimSuffix(path, "\"")
		p.SetPath(path)
	}

	return err
}


func MapStructureDecodeHook(from reflect.Type, to reflect.Type, ref interface{}) (interface{}, error) {
	var err error
	for range onlyOnce {
		if to.String() != "*toolPath.TypeOsPath" {
			break
		}

		if from.String() != "string" {
			err = errors.New("path is not a string")
			break
		}

		fmt.Printf("from: %s => to: %s\n", from.String(), to.String())
		if to.String() == "*toolPath.TypeOsPath" {
			r := New(nil)
			r.SetPath(ref.(string))
			r.State.SetOk()
			ref = r
		}
	}
	return ref, err
}
