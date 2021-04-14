package toolNetwork

import (
	"github.com/gearboxworks/scribeHelpers/toolExec"
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
	"net"
	"strconv"
	"strings"
)

//goland:noinspection ALL
type TypeNetwork struct {
	Listeners      Listeners
	ListenersMap   ListenersMap

	Runtime        *toolRuntime.TypeRuntime
	State          *ux.State
}
func (r *TypeNetwork) IsNil() *ux.State {
	return ux.IfNilReturnError(r)
}

type Listener struct {
	Proto string
	Host  string
	Port  uint16
	//Available bool
}
type Listeners []Listener
type ListenersMap map[uint16]*Listener


func New() *TypeNetwork {
	var ret *TypeNetwork

	for range onlyOnce {
		ret = &TypeNetwork {
			Listeners: Listeners{},

			State: ux.NewState("", false),
		}
	}

	return ret
}

func (r *TypeNetwork) EnsureNotNil() *TypeNetwork {
	if r == nil {
		return New()
	}

	return r
}

func (r *TypeNetwork) GetPorts() *ux.State {
	if state := r.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		exe := toolExec.New(r.Runtime)
		r.State = exe.Exec("netstat", "-an")
		if r.State.IsNotOk() {
			break
		}

		r.Listeners = make(Listeners, 0)
		r.ListenersMap = make(ListenersMap)

		out := exe.GetStdoutArray("\n")
		for _, v := range out {
			col := strings.Fields(v)
			if len(col) == 0 {
				continue
			}

			if !strings.HasPrefix(col[0],"tcp") && !strings.HasPrefix(col[0],"udp") {
				continue
			}

			if strings.HasPrefix(col[0],"tcp") && !strings.HasPrefix(col[5],"LISTEN") {
				continue
			}

			// 0          1      2  3                      4                      5
			// tcp4       0      0  *.111                  *.*                    LISTEN

			e := Listener {
				Proto: col[0],
				Host:  col[3],
				Port:  0,
			}

			// Replace last "." in response.
			i := strings.LastIndex(e.Host, ".")
			e.Host = e.Host[:i] + strings.Replace(e.Host[i:], ".", ":", 1)

			var err error
			var p string
			e.Host, p, err = net.SplitHostPort(e.Host)
			if err != nil {
				//r.State.SetError(err)
				//break
				continue
			}
			if e.Host == "*" {
				e.Host = "0.0.0.0"
			}

			var pi int
			pi, err = strconv.Atoi(p)
			if err != nil {
				//r.State.SetError(err)
				//break
				continue
			}
			e.Port = uint16(pi)

			r.Listeners = append(r.Listeners, e)
			r.ListenersMap[e.Port] = &e
		}
	}

	return r.State
}
