package ux

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/logrusorgru/aurora"
	"strings"
	"text/template"
)


const esc        = "\033["
const clear      = esc + "0m"

func CreateTemplate() error {
	var err error

	for range onlyOnce {
		LoadFuncs()
		colours.TemplateRef = template.New("colours")
		if colours.TemplateRef == nil {
			err = errors.New("Template error - cannot init.")
			break
		}

		colours.TemplateRef = colours.TemplateRef.Funcs(colours.TemplateFuncs)
		if colours.TemplateRef == nil {
			err = errors.New("Template error - cannot load tools.")
			break
		}

		colours.TemplateRef = colours.TemplateRef.Option("missingkey=error")
		if colours.TemplateRef == nil {
			err = errors.New("Template error - cannot set options.")
			break
		}
	}

	return err
}


func TemplatePrintf(format string, args ...interface{}) {
	str, err := TemplateSprintf(format, args...)
	if err != nil {
		fmt.Printf("%s", err.Error())
	} else {
		fmt.Printf("%s", str)
	}
	//return err
}


func TemplateSprintf(format string, args ...interface{}) (string, error) {
	var err error
	var ret string

	for range onlyOnce {
		if colours.TemplateRef == nil {
			err = CreateTemplate()
			if err != nil {
				break
			}
		}

		format = fmt.Sprintf(format, args...)

		var tmplParsed *template.Template
		tmplParsed, err = colours.TemplateRef.Parse(format)
		if err != nil {
			err = errors.New(fmt.Sprintf("Template error - cannot parse - %v", err))
			break
		}

		tmplParsed = tmplParsed.Option("missingkey=error")
		if tmplParsed == nil {
			err = errors.New("Template error - cannot set options.")
			break
		}

		var out bytes.Buffer
		var str2 string
		err = tmplParsed.Execute(&out, &str2)
		if err != nil {
			err = errors.New(fmt.Sprintf("Error processing template: %s", err))
			break
		}

		ret = out.String()
	}

	if err != nil {
		ret = err.Error() + " - " + ret
	}

	return ret, err
}


////////////////////////////////////////////////////////////////////////////////
func LoadFuncs() {
	for range onlyOnce {
		colours.TemplateFuncs = make(template.FuncMap)
		colours.TemplateFuncs["Prefix"] = templatePrefix

		colours.TemplateFuncs["Reset"] = templateReset
		colours.TemplateFuncs["Bold"] = templateBold
		colours.TemplateFuncs["Faint"] = templateFaint
		colours.TemplateFuncs["DoublyUnderline"] = templateDoublyUnderline
		colours.TemplateFuncs["Fraktur"] = templateFraktur
		colours.TemplateFuncs["Italic"] = templateItalic
		colours.TemplateFuncs["Underline"] = templateUnderline
		colours.TemplateFuncs["SlowBlink"] = templateSlowBlink
		colours.TemplateFuncs["RapidBlink"] = templateRapidBlink
		colours.TemplateFuncs["Blink"] = templateBlink
		colours.TemplateFuncs["Reverse"] = templateReverse
		colours.TemplateFuncs["Inverse"] = templateInverse
		colours.TemplateFuncs["Conceal"] = templateConceal
		colours.TemplateFuncs["Hidden"] = templateHidden
		colours.TemplateFuncs["CrossedOut"] = templateCrossedOut
		colours.TemplateFuncs["StrikeThrough"] = templateStrikeThrough
		colours.TemplateFuncs["Framed"] = templateFramed
		colours.TemplateFuncs["Encircled"] = templateEncircled
		colours.TemplateFuncs["Overlined"] = templateOverlined
		colours.TemplateFuncs["Black"] = templateBlack
		colours.TemplateFuncs["Red"] = templateRed
		colours.TemplateFuncs["Green"] = templateGreen
		colours.TemplateFuncs["Yellow"] = templateYellow
		colours.TemplateFuncs["Brown"] = templateBrown
		colours.TemplateFuncs["Blue"] = templateBlue
		colours.TemplateFuncs["Magenta"] = templateMagenta
		colours.TemplateFuncs["Cyan"] = templateCyan
		colours.TemplateFuncs["White"] = templateWhite
		colours.TemplateFuncs["BrightBlack"] = templateBrightBlack
		colours.TemplateFuncs["BrightRed"] = templateBrightRed
		colours.TemplateFuncs["BrightGreen"] = templateBrightGreen
		colours.TemplateFuncs["BrightYellow"] = templateBrightYellow
		colours.TemplateFuncs["BrightBlue"] = templateBrightBlue
		colours.TemplateFuncs["BrightMagenta"] = templateBrightMagenta
		colours.TemplateFuncs["BrightCyan"] = templateBrightCyan
		colours.TemplateFuncs["BrightWhite"] = templateBrightWhite
		colours.TemplateFuncs["BgBlack"] = templateBgBlack
		colours.TemplateFuncs["BgRed"] = templateBgRed
		colours.TemplateFuncs["BgGreen"] = templateBgGreen
		colours.TemplateFuncs["BgYellow"] = templateBgYellow
		colours.TemplateFuncs["BgBrown"] = templateBgBrown
		colours.TemplateFuncs["BgBlue"] = templateBgBlue
		colours.TemplateFuncs["BgMagenta"] = templateBgMagenta
		colours.TemplateFuncs["BgCyan"] = templateBgCyan
		colours.TemplateFuncs["BgWhite"] = templateBgWhite
		colours.TemplateFuncs["BgBrightBlack"] = templateBgBrightBlack
		colours.TemplateFuncs["BgBrightRed"] = templateBgBrightRed
		colours.TemplateFuncs["BgBrightGreen"] = templateBgBrightGreen
		colours.TemplateFuncs["BgBrightYellow"] = templateBgBrightYellow
		colours.TemplateFuncs["BgBrightBlue"] = templateBgBrightBlue
		colours.TemplateFuncs["BgBrightMagenta"] = templateBgBrightMagenta
		colours.TemplateFuncs["BgBrightCyan"] = templateBgBrightCyan
		colours.TemplateFuncs["BgBrightWhite"] = templateBgBrightWhite
	}
}

func templatePrefix() string {
	return colours.Prefix
}

func templateReset() string {
	// Don't remove the "clear" escape sequence.
	return aurora.Black("").String()
}

func templateBold() string {
	return strings.TrimSuffix(aurora.Bold("").String(), clear)
}

func templateFaint() string {
	return strings.TrimSuffix(aurora.Faint("").String(), clear)
}

func templateDoublyUnderline() string {
	return strings.TrimSuffix(aurora.DoublyUnderline("").String(), clear)
}

func templateFraktur() string {
	return strings.TrimSuffix(aurora.Fraktur("").String(), clear)
}

func templateItalic() string {
	return strings.TrimSuffix(aurora.Italic("").String(), clear)
}

func templateUnderline() string {
	return strings.TrimSuffix(aurora.Underline("").String(), clear)
}

func templateSlowBlink() string {
	return strings.TrimSuffix(aurora.SlowBlink("").String(), clear)
}

func templateRapidBlink() string {
	return strings.TrimSuffix(aurora.RapidBlink("").String(), clear)
}

func templateBlink() string {
	return strings.TrimSuffix(aurora.Blink("").String(), clear)
}

func templateReverse() string {
	return strings.TrimSuffix(aurora.Reverse("").String(), clear)
}

func templateInverse() string {
	return strings.TrimSuffix(aurora.Inverse("").String(), clear)
}

func templateConceal() string {
	return strings.TrimSuffix(aurora.Conceal("").String(), clear)
}

func templateHidden() string {
	return strings.TrimSuffix(aurora.Hidden("").String(), clear)
}

func templateCrossedOut() string {
	return strings.TrimSuffix(aurora.CrossedOut("").String(), clear)
}

func templateStrikeThrough() string {
	return strings.TrimSuffix(aurora.StrikeThrough("").String(), clear)
}

func templateFramed() string {
	return strings.TrimSuffix(aurora.Framed("").String(), clear)
}

func templateEncircled() string {
	return strings.TrimSuffix(aurora.Encircled("").String(), clear)
}

func templateOverlined() string {
	return strings.TrimSuffix(aurora.Overlined("").String(), clear)
}

func templateBlack() string {
	return strings.TrimSuffix(aurora.Black("").String(), clear)
}

func templateRed() string {
	return strings.TrimSuffix(aurora.Red("").String(), clear)
}

func templateGreen() string {
	return strings.TrimSuffix(aurora.Green("").String(), clear)
}

func templateYellow() string {
	return strings.TrimSuffix(aurora.Yellow("").String(), clear)
}

func templateBrown() string {
	return strings.TrimSuffix(aurora.Brown("").String(), clear)
}

func templateBlue() string {
	return strings.TrimSuffix(aurora.Blue("").String(), clear)
}

func templateMagenta() string {
	return strings.TrimSuffix(aurora.Magenta("").String(), clear)
}

func templateCyan() string {
	return strings.TrimSuffix(aurora.Cyan("").String(), clear)
}

func templateWhite() string {
	return strings.TrimSuffix(aurora.White("").String(), clear)
}

func templateBrightBlack() string {
	return strings.TrimSuffix(aurora.BrightBlack("").String(), clear)
}

func templateBrightRed() string {
	return strings.TrimSuffix(aurora.BrightRed("").String(), clear)
}

func templateBrightGreen() string {
	return strings.TrimSuffix(aurora.BrightGreen("").String(), clear)
}

func templateBrightYellow() string {
	return strings.TrimSuffix(aurora.BrightYellow("").String(), clear)
}

func templateBrightBlue() string {
	return strings.TrimSuffix(aurora.BrightBlue("").String(), clear)
}

func templateBrightMagenta() string {
	return strings.TrimSuffix(aurora.BrightMagenta("").String(), clear)
}

func templateBrightCyan() string {
	return strings.TrimSuffix(aurora.BrightCyan("").String(), clear)
}

func templateBrightWhite() string {
	return strings.TrimSuffix(aurora.BrightWhite("").String(), clear)
}

func templateBgBlack() string {
	return strings.TrimSuffix(aurora.BgBlack("").String(), clear)
}

func templateBgRed() string {
	return strings.TrimSuffix(aurora.BgRed("").String(), clear)
}

func templateBgGreen() string {
	return strings.TrimSuffix(aurora.BgGreen("").String(), clear)
}

func templateBgYellow() string {
	return strings.TrimSuffix(aurora.BgYellow("").String(), clear)
}

func templateBgBrown() string {
	return strings.TrimSuffix(aurora.BgBrown("").String(), clear)
}

func templateBgBlue() string {
	return strings.TrimSuffix(aurora.BgBlue("").String(), clear)
}

func templateBgMagenta() string {
	return strings.TrimSuffix(aurora.BgMagenta("").String(), clear)
}

func templateBgCyan() string {
	return strings.TrimSuffix(aurora.BgCyan("").String(), clear)
}

func templateBgWhite() string {
	return strings.TrimSuffix(aurora.BgWhite("").String(), clear)
}

func templateBgBrightBlack() string {
	return strings.TrimSuffix(aurora.BgBrightBlack("").String(), clear)
}

func templateBgBrightRed() string {
	return strings.TrimSuffix(aurora.BgBrightRed("").String(), clear)
}

func templateBgBrightGreen() string {
	return strings.TrimSuffix(aurora.BgBrightGreen("").String(), clear)
}

func templateBgBrightYellow() string {
	return strings.TrimSuffix(aurora.BgBrightYellow("").String(), clear)
}

func templateBgBrightBlue() string {
	return strings.TrimSuffix(aurora.BgBrightBlue("").String(), clear)
}

func templateBgBrightMagenta() string {
	return strings.TrimSuffix(aurora.BgBrightMagenta("").String(), clear)
}

func templateBgBrightCyan() string {
	return strings.TrimSuffix(aurora.BgBrightCyan("").String(), clear)
}

func templateBgBrightWhite() string {
	return strings.TrimSuffix(aurora.BgBrightWhite("").String(), clear)
}
