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
func (ghr *TypeGhr) mustCopyN(w io.Writer, r io.Reader, n int64) *ux.State {
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


//func (ghr *TypeGhr) renderInfoText(tags Tags, releases *Releases) *ux.State {
//	if state := ghr.IsNil(); state.IsError() {
//		return state
//	}
//	ghr.State.SetFunction()
//
//	for range onlyOnce {
//		var t []string
//		for _, tag := range tags {
//			t = append(t, tag.Name)
//		}
//		ghr.message("Tags: %s", strings.Join(t, ", "))
//
//		ghr.message("Releases")
//		for _, release := range *releases {
//			ghr.message("- %v", release)
//		}
//
//		ghr.State.SetOk()
//	}
//
//	//return nil
//	return ghr.State
//}


//func (ghr *TypeGhr) renderInfoJSON(tags Tags, releases *Releases) *ux.State {
//	if state := ghr.IsNil(); state.IsError() {
//		return state
//	}
//	ghr.State.SetFunction()
//
//	for range onlyOnce {
//		out := struct {
//			Tags     Tags
//			Releases *Releases
//		}{
//			Tags:     tags,
//			Releases: releases,
//		}
//
//		enc := json.NewEncoder(os.Stdout)
//		enc.SetIndent("", "    ")
//
//		ghr.State.SetOk()
//		ghr.State.SetResponse(enc.Encode(&out))
//	}
//
//	return ghr.State
//}
