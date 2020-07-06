package toolPath

import (
	"bytes"
	"github.com/newclarity/scribeHelpers/toolPrompt"
	"github.com/newclarity/scribeHelpers/ux"
	"io/ioutil"
	"os"
	"strings"
)


const DefaultSeparator = "\n"


func (p *TypeOsPath) LoadContents(data ...interface{}) {
	for range onlyOnce {
		p._String = ""
		p._Array = []string{}

		p.AppendContents(data...)
	}
}
func (p *TypeOsPath) SetContents(data ...interface{}) {
	p.LoadContents(data...)
}


func (p *TypeOsPath) AppendContents(data ...interface{}) {
	for range onlyOnce {
		if p._Separator == "" {
			p._Separator = DefaultSeparator
		}

		for _, d := range data {
			//value := reflect.ValueOf(d)
			//switch value.Kind() {
			//	case reflect.output:
			//		p._Array = append(p._Array, value.output())
			//	case reflect.Array:
			//		p._Array = append(p._Array, d.([]string)...)
			//	case reflect.Slice:
			//		p._Array = append(p._Array, d.([]string)...)
			//}

			var sa []string
			switch d.(type) {
				case []string:
					for _, s := range d.([]string) {
						if s != "" {
							s = removeDupeEol(s)
							sa = append(sa, strings.Split(s, p._Separator)...)
						}
					}

				case string:
					if d.(string) != "" {
						s := removeDupeEol(d.(string))
						sa = append(sa, strings.Split(s, p._Separator)...)
					}

				case []byte:
					s := removeDupeEol(string(d.([]byte)))
					//s := removeDupeEol(d.(string))
					sa = append(sa, strings.Split(s, p._Separator)...)

				case *bytes.Buffer:
					b := d.(*bytes.Buffer).String()
					s := removeDupeEol(b)
					sa = append(sa, strings.Split(s, p._Separator)...)
			}

			p._Array = append(p._Array, sa...)
		}
		p._String = strings.Join(p._Array, p._Separator)
	}
}

func removeDupeEol(s string) string {
	s = strings.ReplaceAll(s, `\n\n`, `\n`)
	s = strings.ReplaceAll(s, `\r\n\r\n`, `\r\n`)

	// @TODO better way to do this.
	s = strings.ReplaceAll(s, `\n\n`, `\n`)
	s = strings.ReplaceAll(s, `\r\n\r\n`, `\r\n`)

	s = strings.TrimSpace(s)

	return s
}


func (p *TypeOsPath) GetContentString() string {
	if p._Separator == "" {
		p._Separator = DefaultSeparator
	}

	return strings.Join(p._Array, p._Separator)
}


func (p *TypeOsPath) GetContentLength() int {
	return len(p._String)
}


func (p *TypeOsPath) GetContentArray() []string {
	return p._Array
}


func (p *TypeOsPath) GetContentByteArray() []byte {
	return []byte(strings.Join(p._Array, p._Separator))
}


func (p *TypeOsPath) SetSeparator(separator string) {
	for range onlyOnce {
		p._Separator = separator
		p._Array = strings.Split(p._String, p._Separator)
	}
}


func (p *TypeOsPath) GetSeparator() string {
	return p._Separator
}


func (p *TypeOsPath) Foo() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.SetOk()

		if !p.IsValid() {
			p.State.SetWarning("path is invalid")
			break
		}

		p.StatPath()
		if p.State.IsError() {
			break
		}
		if !p._Exists {
			p.State.SetError("file '%s' not found", p.Path)
			break
		}
		if p._IsDir {
			p.State.SetError("path '%s' is a directory", p.Path)
			break
		}



		//f, err := os.Open("/tmp/dat")
		//
		//b1 := make([]byte, 5)
		//n1, err := f.Read(b1)
		//
		//fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
		//o3, err := f.Seek(6, 0)
		//
		//b3 := make([]byte, 2)
		//n3, err := io.ReadAtLeast(f, b3, 2)
		//
		//fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
		//
		//r4 := bufio.NewReader(f)
		//b4, err := r4.Peek(5)
		//check(err)
		//fmt.Printf("5 bytes: %s\n", string(b4))



		var d []byte
		var err error
		d, err = ioutil.ReadFile(p.Path)
		if err != nil {
			p.State.SetError(err)
			break
		}

		p.LoadContents(d)
		p.State.SetOk("file '%s' read OK", p.Path)
	}

	return p.State
}


func (p *TypeOsPath) ReadFile() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.SetOk()

		if !p.IsValid() {
			p.State.SetWarning("path is invalid")
			break
		}

		p.StatPath()
		if p.State.IsError() {
			break
		}
		if !p._Exists {
			p.State.SetError("file '%s' not found", p.Path)
			break
		}
		if p._IsDir {
			p.State.SetError("path '%s' is a directory", p.Path)
			break
		}

		var d []byte
		var err error
		d, err = ioutil.ReadFile(p.Path)
		if err != nil {
			p.State.SetError(err)
			break
		}

		p.LoadContents(d)
		p.State.SetOk("file '%s' read OK", p.Path)
	}

	return p.State
}


func (p *TypeOsPath) WriteFile() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.SetOk()

		if !p.IsValid() {
			p.State.SetWarning("path is invalid")
			break
		}

		if p._String == "" {
			p.State.SetError("content string is nil")
			break
		}

		for range onlyOnce {
			p.StatPath()
			if p._IsDir {
				p.State.SetError("path '%s' is a directory", p.Path)
				break
			}
			if p.NotExists() {
				p.State.SetOk()
				break
			}
			if p._CanOverwrite {
				break
			}

			if !toolPrompt.ToolUserPromptBool("Overwrite file '%s'? (Y|N) ", p.Path) {
				p.State.SetWarning("not overwriting file '%s'", p.Path)
				break
			}
			p.State.SetOk()
		}
		if p.State.IsNotOk() {
			break
		}


		if p._Mode == 0 {
			p._Mode = 0644
		}

		err := ioutil.WriteFile(p.Path, []byte(p._String), p._Mode)
		if err != nil {
			p.State.SetError(err)
			break
		}

		p.State.SetOk("file '%s' written OK", p.Path)
	}

	return p.State
}


func (p *TypeOsPath) OpenFile() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.SetOk()

		if !p.IsValid() {
			p.State.SetWarning("path is invalid")
			break
		}

		for range onlyOnce {
			p.StatPath()
			if p._IsDir {
				p.State.SetError("path '%s' is a directory", p.Path)
				break
			}
			if p.NotExists() {
				p.State.SetOk()
				break
			}
			if p._CanOverwrite {
				break
			}

			if !toolPrompt.ToolUserPromptBool("Overwrite file '%s'? (Y|N) ", p.Path) {
				p.State.SetWarning("not overwriting file '%s'", p.Path)
				break
			}
			p.State.SetOk()
		}
		if p.State.IsNotOk() {
			break
		}


		if p._Mode == 0 {
			p._Mode = 0644
		}


		var err error
		p.FileHandle, err = os.Create(p.Path)
		if err != nil {
			p.State.SetError("Cannot open file '%s' for writing - %s", p.Path, err)
			break
		}

		p.State.SetResponse(p.FileHandle)

		p.State.SetOk("TypeFile '%s' opened OK", p.Path)
	}

	return p.State
}


func (p *TypeOsPath) OpenFileHandle() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.SetOk()

		if !p.IsValid() {
			p.State.SetWarning("path is invalid")
			break
		}

		p.StatPath()
		if p._IsDir {
			p.State.SetError("path '%s' is a directory", p.Path)
			break
		}
		if p.NotExists() {
			p.State.SetOk()
			break
		}


		var err error
		p.FileHandle, err = os.Open(p.Path)
		if err != nil {
			p.State.SetError("Cannot open file '%s' for writing - %s", p.Path, err)
			break
		}

		p.State.SetResponse(p.FileHandle)
		p.State.SetOk("TypeFile '%s' opened OK", p.Path)
	}

	return p.State
}


func (p *TypeOsPath) SetFileHandle(fh *os.File) *ux.State {
	for range onlyOnce {
		p.FileHandle = fh
		p.SetPath(p.FileHandle.Name())
		p.State.SetOk()
	}

	return p.State
}


func (p *TypeOsPath) GetFileHandle() (*os.File, *ux.State) {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.SetOk()

		if !p.IsValid() {
			p.State.SetWarning("path is invalid")
			break
		}

		p.State.SetResponse(p.FileHandle)
	}

	return p.FileHandle, p.State
}


func (p *TypeOsPath) CloseFile() *ux.State {
	for range onlyOnce {
		p.State.SetFunction()
		p.State.SetOk()

		var err error
		err = p.FileHandle.Sync()
		if err != nil {
			p.State.SetWarning("Error when syncing file '%s' - ", p.Path, err)
		}

		err = p.FileHandle.Close()
		if err != nil {
			p.State.SetWarning("Error when closing file '%s' - ", p.Path, err)
		}

		p.State.SetOk("TypeFile '%s' closed OK", p.Path)
	}

	return p.State
}
