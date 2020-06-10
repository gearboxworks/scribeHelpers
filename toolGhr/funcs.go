package toolGhr

import (
	"encoding/json"
	"fmt"
	"github.com/newclarity/scribeHelpers/ux"
	"io"
	"os"
	"strings"
	"time"
)



func ReleaseSync(srcrepo string, binrepo string, version string, path string) *ux.State {
	state := ux.NewState("ReleaseSync", false)

	for range onlyOnce {
		if srcrepo == "" {
			state.SetError("source repo invalid")
			break
		}

		if binrepo == "" {
			state.SetError("binary repo invalid")
			break
		}

		if version == "" {
			state.SetError("version invalid")
			break
		}

		if path == "" {
			state.SetError("cache dir invalid")
			break
		}

		if binrepo == srcrepo {
			// No need to push to binary repo.
			// GoReleaser will handle this.
			break
		}

		// Setup source repo.
		Src := New(nil)
		if Src.State.IsNotOk() {
			state = Src.State
			break
		}
		state = Src.SetAuth(TypeAuth{ Token: "", AuthUser: "" })
		if state.IsNotOk() {
			break
		}
		state = Src.OpenUrl(srcrepo)
		if state.IsNotOk() {
			break
		}
		state = Src.SetTag(version)
		if state.IsNotOk() {
			break
		}

		// Setup destination repo.
		Dest := New(nil)
		if Src.State.IsNotOk() {
			state = Src.State
			break
		}
		state = Dest.IsNil()
		if state.IsNotOk() {
			break
		}
		state = Dest.OpenUrl(binrepo)
		if state.IsNotOk() {
			break
		}
		state = Dest.SetOverwrite(true)
		if state.IsNotOk() {
			break
		}

		// Now sync the release in the destination repo.
		state = Dest.CopyFrom(Src.Repo, path)
	}

	return state
}


/* usually when something goes wrong, github sends something like this back */
type message struct {
	message string        `json:"message"`
	Errors  []GithubError `json:"errors"`
}

type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
}


/* transforms a stream into a message, if it's valid json */
func Tomessage(r io.Reader) (*message, error) {
	var msg message
	if err := json.NewDecoder(r).Decode(&msg); err != nil {
		return nil, err
	}

	return &msg, nil
}


func (m *message) String() string {
	str := fmt.Sprintf("msg: %v, errors: ", m.message)

	errstr := make([]string, len(m.Errors))
	for idx, err := range m.Errors {
		errstr[idx] = fmt.Sprintf("[field: %v, code: %v]",
			err.Field, err.Code)
	}

	return str + strings.Join(errstr, ", ")
}


/* nvls returns the first value in xs that is not empty. */
func nvls(xs ...string) string {
	for _, s := range xs {
		if s != "" {
			return s
		}
	}

	return ""
}


// formats time `t` as `fmt` if it is not nil, otherwise returns `def`
func timeFmtOr(t *time.Time, fmt, def string) string {
	if t == nil {
		return def
	}
	return t.Format(fmt)
}


// isCharDevice returns true if f is a character device (panics if f can't
// be stat'ed).
func isCharDevice(f *os.File) bool {
	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	return (stat.Mode() & os.ModeCharDevice) != 0
}


func Mark(ok bool) string {
	if ok {
		return "✔"
	} else {
		return "✗"
	}
}


// mustCopyN attempts to copy exactly N bytes, if this fails, an error is
// returned.
func (ghr *TypeGhr) mustCopyN(w *os.File, r io.Reader, n int64) *ux.State {
	if state := ghr.IsNil(); state.IsError() {
		return state
	}
	ghr.State.SetFunction()

	for range onlyOnce {
		an, err := io.Copy(w, r)
		if an != n {
			ghr.State.SetError("data did not match content length %d != %d", an, n)
			break
		}
		if err != nil {
			ghr.State.SetError(err)
			break
		}

		ghr.State.SetOk()
	}

	return ghr.State
}


func (ghr *TypeGhr) message(format string, args ...interface{}) {
	ux.PrintfCyan("%s: ", ghr.Repo.GetUrl())
	ux.PrintflnBlue(format, args...)
}
