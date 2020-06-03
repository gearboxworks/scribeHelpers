package toolGear

import (
	"github.com/newclarity/scribeHelpers/toolRuntime"
	"github.com/newclarity/scribeHelpers/ux"
	"net/url"
	"os"
)

const (
	ProviderDocker = "docker"
)

type Provider struct {
	Name    string  `json:"name"`
	Host    string  `json:"host"`
	Port    string  `json:"port"`
	Url     url.URL `json:"url"`
	Project string  `json:"project"`

	runtime *toolRuntime.TypeRuntime
	State   *ux.State
}


func NewProvider(runtime *toolRuntime.TypeRuntime) *Provider {
	runtime = runtime.EnsureNotNil()

	p := &Provider{
		Name:    ProviderDocker,
		Host:    "",
		Port:    "",
		Url:     url.URL{},
		Project: "",
		runtime: runtime,
		State:   ux.NewState(runtime.CmdName, runtime.Debug),
	}
	p.State.SetPackage("")
	p.State.SetFunction("")
	return p
}


func (p *Provider) SetProvider(provider string) *ux.State {
	switch provider {
		case ProviderDocker:
			p.Name = provider

		default:
			p.State.SetError("Unknown provider '%s'", provider)
	}

	return p.State
}


func (p *Provider) SetHost(host string, port string) *ux.State {
	switch p.Name {
		case ProviderDocker:
			if host == "" {
				p.State.SetOk()	// Don't error - default is local host via socket.
				break
			}
			p.Host = host

			if port == "" {
				port = "2376"
			}
			p.Port = port

			var urlString *url.URL
			var err error
			//urlString, err = client.ParseHostURL(fmt.Sprintf("tcp://%s:%s", p.Host, p.Port))
			urlString, err = ParseHostURL("tcp://%s:%s", p.Host, p.Port)
			if err != nil {
				p.State.SetError(err)
				break
			}

			err = os.Setenv("DOCKER_HOST", urlString.String())
			if err != nil {
				p.State.SetError(err)
				break
			}
			p.State.SetOk()

		default:
			p.State.SetError("Unknown provider '%s'", p.Name)
	}

	return p.State
}
