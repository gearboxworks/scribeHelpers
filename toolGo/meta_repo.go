package toolGo

import (
	"errors"
	"fmt"
	"github.com/gearboxworks/scribeHelpers/ux"
	"net/url"
	"strings"
)


type Repo struct {
	url *url.URL
}


func (r *Repo) Set(repo string) error {
	var err error
	for range onlyOnce {
		v, err := url.Parse(addPrefix(repo))
		if err == nil {
			r.url = v
			break
		}
		err = errors.New(fmt.Sprintf("Invalid repo URL '%s' - %v", repo, err))
	}
	return err
}


func (r *Repo) GetOwner() string {
	value, _ := r.Get()
	return value
}


func (r *Repo) GetName() string {
	_, value := r.Get()
	return value
}


func (r *Repo) Get() (string, string) {
	var owner string
	var name string
	for range onlyOnce {
		pa := strings.Split(r.url.Path, "/")
		switch len(pa) {
		case 0:
		case 1:
		case 2:
			owner = pa[1]
			name = ""

		default:
			owner = pa[1]
			name = pa[2]
		}
	}
	return owner, name
}


func (r *Repo) GetUrl() string {
	return r.url.String()
}


func (r *Repo) String() string {
	var ret string
	for range onlyOnce {
		ret += ux.SprintfBlue("Repo URL: ")   + ux.SprintfCyan("%v\n", r.url)
		ret += ux.SprintfBlue("Repo owner: ") + ux.SprintfCyan("%s\n", r.GetOwner())
		ret += ux.SprintfBlue("Repo name: ")  + ux.SprintfCyan("%s\n", r.GetName())
	}
	return ret
}


func (r *Repo) IsValid() bool {
	var ok bool
	for range onlyOnce {
		if r == nil {
			break
		}
		if r.url == nil {
			break
		}
		ok = true
	}
	return ok
}


func (r *Repo) IsNotValid() bool {
	return !r.IsValid()
}


func (r *Repo) IsSame(compare *Repo) bool {
	var ok bool
	for range onlyOnce {
		if r.GetUrl() != compare.GetUrl() {
			break
		}
		ok = true
	}
	return ok
}


func (r *Repo) IsNotSame(compare *Repo) bool {
	return !r.IsSame(compare)
}


func addPrefix(u string) string {
	for range onlyOnce {
		if strings.HasPrefix(u, "http") {
			// We have a full URL - no change.
			break
		}

		if strings.HasPrefix(u, "github.com") {
			// We have a github.com specific string.
			u = "https://" + u
			break
		}

		ua := strings.Split(u, "/")
		if len(ua) == 0 {
			// Dunno, leave as is.
			break
		}

		if strings.Contains(ua[0], ".") {
			// We have a host defined in the first segment.
			u = "https://" + u
			break
		}

		// We probably just have a "owner/repo_name" style URL.
		u = "https://github.com/" + u
	}

	return u
}
