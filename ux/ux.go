// Other possibilities:
//
// https://github.com/nsf/termbox-go
// https://github.com/jroimartin/gocui
// https://github.com/marcusolsson/tui-go
// https://github.com/rivo/tview
// https://github.com/gizak/termui
// https://github.com/logrusorgru/aurora
package ux

import (
	"errors"
	"fmt"
	"github.com/logrusorgru/aurora"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"text/template"
)


type UxGetter interface {
	Open() error
	Close()
	PrintfWhite(format string, args ...interface{})
	PrintfCyan(format string, args ...interface{})
	PrintfYellow(format string, args ...interface{})
	PrintfRed(format string, args ...interface{})
	PrintfGreen(format string, args ...interface{})
	PrintfBlue(format string, args ...interface{})
	PrintfMagenta(format string, args ...interface{})
	SprintfWhite(format string, args ...interface{}) string
	SprintfCyan(format string, args ...interface{}) string
	SprintfYellow(format string, args ...interface{}) string
	SprintfRed(format string, args ...interface{}) string
	SprintfGreen(format string, args ...interface{}) string
	SprintfBlue(format string, args ...interface{}) string
	SprintfMagenta(format string, args ...interface{}) string
	Sprintf(format string, args ...interface{}) string
	Printf(format string, args ...interface{})
	SprintfOk(format string, args ...interface{}) string
	PrintfOk(format string, args ...interface{})
	SprintfWarning(format string, args ...interface{}) string
	PrintfWarning(format string, args ...interface{})
	SprintfError(format string, args ...interface{}) string
	PrintfError(format string, args ...interface{})
	SprintError(err error) string
	PrintError(err error)
	GetTerminalSize() (int, int, error)
}

type Ux struct {
}


//noinspection GoUnusedGlobalVariable
type typeColours struct {
	Ref           aurora.Aurora
	Defined       bool
	Name          string
	EnableColours bool
	TemplateRef   *template.Template
	TemplateFuncs template.FuncMap
	Prefix        string
}
var colours typeColours


func Open(name string, enable bool) (*typeColours, error) {
	var err error

	for range onlyOnce {
		if name == "" {
			name = "Gearbox"
		}
		name += ": "

		colours.Ref = aurora.NewAurora(enable)
		colours.Name = name
		colours.EnableColours = enable
		colours.Defined = true
		colours.Prefix = fmt.Sprintf("%s", aurora.BrightCyan(colours.Name).Bold())

		//err = termui.Init();
		//if err != nil {
		//	fmt.Printf("failed to initialize termui: %v", err)
		//	break
		//}

		err = CreateTemplate()
	}

	return &colours, err
}


func Close() {
	if colours.Defined {
		//termui.Close()
	}
}


func DisableColours() {
	colours.EnableColours = false
}
func EnableColours() {
	colours.EnableColours = true
}



func PrintfWhite(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightWhite(inline))
	}
	_, _ = fmt.Fprint(os.Stdout, inline)
}

func PrintfCyan(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightCyan(inline))
	}
	_, _ = fmt.Fprint(os.Stdout, inline)
}

func PrintfYellow(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightYellow(inline))
	}
	_, _ = fmt.Fprint(os.Stdout, inline)
}

func PrintfRed(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightRed(inline))
	}
	_, _ = fmt.Fprint(os.Stdout, inline)
}

func PrintfGreen(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightGreen(inline))
	}
	_, _ = fmt.Fprint(os.Stdout, inline)
}

func PrintfBlue(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightBlue(inline))
	}
	_, _ = fmt.Fprint(os.Stdout, inline)
}

func PrintfMagenta(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightMagenta(inline))
	}
	_, _ = fmt.Fprint(os.Stdout, inline)
}

func PrintflnWhite(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightWhite(inline))
	}
	_, _ = fmt.Fprint(os.Stdout, inline + "\n")
}

func PrintflnCyan(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightCyan(inline))
	}
	_, _ = fmt.Fprint(os.Stdout, inline + "\n")
}

func PrintflnYellow(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightYellow(inline))
	}
	_, _ = fmt.Fprint(os.Stdout, inline + "\n")
}

func PrintflnRed(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightRed(inline))
	}
	_, _ = fmt.Fprint(os.Stdout, inline + "\n")
}

func PrintflnGreen(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightGreen(inline))
	}
	_, _ = fmt.Fprint(os.Stdout, inline + "\n")
}

func PrintflnBlue(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightBlue(inline))
	}
	_, _ = fmt.Fprint(os.Stdout, inline + "\n")
}

func PrintflnMagenta(format string, args ...interface{}) {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightMagenta(inline))
	}
	_, _ = fmt.Fprint(os.Stdout, inline + "\n")
}


func SprintfWhite(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightWhite(inline))
	}
	return inline
}

func SprintfCyan(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightCyan(inline))
	}
	return inline
}

func SprintfYellow(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightYellow(inline))
	}
	return inline
}

func SprintfRed(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightRed(inline))
	}
	return inline
}

func SprintfGreen(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightGreen(inline))
	}
	return inline
}

func SprintfBlue(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightBlue(inline))
	}
	return inline
}

func SprintfMagenta(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s", aurora.BrightMagenta(inline))
	}
	return inline
}


func Sprintf(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = fmt.Sprintf("%s%s", colours.Prefix, inline)
	}
	return inline
}
func Printf(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, Sprintf(format, args...))
}


func SprintfNormal(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	if colours.EnableColours {
		inline = Sprintf("%s", aurora.BrightBlue(inline))
	}
	return inline
}

func PrintfNormal(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfNormal(format, args...))
}

func PrintflnNormal(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfNormal(format + "\n", args...))
}


func SprintfInfo(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	return Sprintf("%s", aurora.BrightBlue(inline))
}

func PrintfInfo(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfInfo(format, args...))
}

func PrintflnInfo(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfInfo(format + "\n", args...))
}


func SprintfOk(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	return Sprintf("%s", aurora.BrightGreen(inline))
}

func PrintfOk(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfOk(format, args...))
}

func PrintflnOk(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfOk(format + "\n", args...))
}


func SprintfDebug(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func PrintfDebug(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf(format + "\n", args...))
}


func SprintfWarning(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	return Sprintf("%s", aurora.BrightYellow(inline))
}

func PrintfWarning(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfWarning(format, args...))
}

func PrintflnWarning(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, SprintfWarning(format + "\n", args...))
}


func SprintfError(format string, args ...interface{}) string {
	inline := fmt.Sprintf(format, args...)
	return Sprintf("%s", aurora.BrightRed(inline))
}

func PrintfError(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, SprintfError(format, args...))
}

func PrintflnError(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, SprintfError(format + "\n", args...))
}


func SprintError(err error) string {
	var s string

	for range onlyOnce {
		if err == nil {
			break
		}

		s = Sprintf("%s%s\n", aurora.BrightRed("ERROR: ").Framed(), aurora.BrightRed(err).Framed().SlowBlink().BgBrightWhite())
	}

	return s
}

func PrintError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, SprintError(err))
}


func GetTerminalSize() (int, int, error) {
	var width int
	var height int
	var err error

	fileDescriptor := int(os.Stdin.Fd())
	if terminal.IsTerminal(fileDescriptor) {
		width, height, err = terminal.GetSize(fileDescriptor)
	} else {
		err = errors.New("not a terminal")
	}

	return width, height, err
}
