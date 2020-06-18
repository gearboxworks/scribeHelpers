package toolGo

import (
	"errors"
	"fmt"
	"github.com/newclarity/scribeHelpers/ux"
)


type Name struct {
	name string
}


func (n *Name) Set(name string) error {
	var err error
	for range onlyOnce {
		if name != "" {
			n.name = name
			break
		}
		err = errors.New(fmt.Sprintf("Invalid repo name '%s' - %v", name, err))
	}
	return err
}


func (n *Name) String() string {
	var ret string
	for range onlyOnce {
		ret += ux.SprintfBlue("%s: ", BinaryName)
		ret += ux.SprintfCyan("%v\n", n.name)
	}
	return ret
}


func (n *Name) Get() string {
	return n.name
}


func (n *Name) IsValid() bool {
	var ok bool
	for range onlyOnce {
		if n == nil {
			break
		}
		if n.name == "" {
			break
		}
		ok = true
	}
	return ok
}


func (n *Name) IsNotValid() bool {
	return !n.IsValid()
}


func (n *Name) IsSame(compare *Name) bool {
	var ok bool
	for range onlyOnce {
		if n.name != compare.name {
			break
		}
		ok = true
	}
	return ok
}


func (n *Name) IsNotSame(compare *Name) bool {
	return !n.IsSame(compare)
}
