package toolUx

import "encoding/json"

type colors struct {
	Bg struct {
		Black    string `json:"black"`
		Blue     string `json:"blue"`
		Cyan     string `json:"cyan"`
		Dwhite   string `json:"dwhite"`
		Gray     string `json:"gray"`
		Green    string `json:"green"`
		Lblue    string `json:"lblue"`
		Lcyan    string `json:"lcyan"`
		Lgreen   string `json:"lgreen"`
		Lmagenta string `json:"lmagenta"`
		Lpink    string `json:"lpink"`
		Lred     string `json:"lred"`
		Lwhite   string `json:"lwhite"`
		Lyellow  string `json:"lyellow"`
		Magenta  string `json:"magenta"`
		Mwhite   string `json:"mwhite"`
		Orange   string `json:"orange"`
		Pink     string `json:"pink"`
		Red      string `json:"red"`
		White    string `json:"white"`
		Yellow   string `json:"yellow"`
	} `json:"bg"`
	Blink struct {
		Fast string `json:"fast"`
		Off  string `json:"off"`
		Slow string `json:"slow"`
	} `json:"blink"`
	Bold struct {
		Off string `json:"off"`
		On  string `json:"on"`
	} `json:"bold"`
	Circle struct {
		Off string `json:"off"`
		On  string `json:"on"`
	} `json:"circle"`
	Clear struct {
		Line struct {
			All   string `json:"all"`
			End   string `json:"end"`
			Start string `json:"start"`
		} `json:"line"`
		Screen struct {
			All   string `json:"all"`
			End   string `json:"end"`
			Start string `json:"start"`
		} `json:"screen"`
	} `json:"clear"`
	Conceal struct {
		Off string `json:"off"`
		On  string `json:"on"`
	} `json:"conceal"`
	Cursor struct {
		Down      string `json:"down"`
		Left      string `json:"left"`
		Next_line string `json:"next-line"`
		Prev_line string `json:"prev-line"`
		Restore   string `json:"restore"`
		Right     string `json:"right"`
		Save      string `json:"save"`
		Up        string `json:"up"`
	} `json:"cursor"`
	Dim struct {
		Off string `json:"off"`
		On  string `json:"on"`
	} `json:"dim"`
	Fg struct {
		Black    string `json:"black"`
		Blue     string `json:"blue"`
		Cyan     string `json:"cyan"`
		Dwhite   string `json:"dwhite"`
		Gray     string `json:"gray"`
		Green    string `json:"green"`
		Lblue    string `json:"lblue"`
		Lcyan    string `json:"lcyan"`
		Lgreen   string `json:"lgreen"`
		Lmagenta string `json:"lmagenta"`
		Lpink    string `json:"lpink"`
		Lred     string `json:"lred"`
		Lwhite   string `json:"lwhite"`
		Lyellow  string `json:"lyellow"`
		Magenta  string `json:"magenta"`
		Mwhite   string `json:"mwhite"`
		Orange   string `json:"orange"`
		Pink     string `json:"pink"`
		Red      string `json:"red"`
		White    string `json:"white"`
		Yellow   string `json:"yellow"`
	} `json:"fg"`
	Frame struct {
		Off string `json:"off"`
		On  string `json:"on"`
	} `json:"frame"`
	Inverse struct {
		Off string `json:"off"`
		On  string `json:"on"`
	} `json:"inverse"`
	Italic struct {
		Off string `json:"off"`
		On  string `json:"on"`
	} `json:"italic"`
	Overline struct {
		Off string `json:"off"`
		On  string `json:"on"`
	} `json:"overline"`
	Overscore struct {
		Off string `json:"off"`
		On  string `json:"on"`
	} `json:"overscore"`
	Reset     string `json:"reset"`
	Underline struct {
		Off string `json:"off"`
		On  string `json:"on"`
	} `json:"underline"`
}

func init() {
	_ = json.Unmarshal(([]byte)(ColorString), &Color)
}

var Color colors

var ColorString = `
{
	"reset": "\u001b[0m",
	"bold": {
		"on": "\u001b[1m",
		"off": "\u001b[21m"
	},
	"dim": {
		"on": "\u001b[2m",
		"off": "\u001b[22m"
	},
	"italic": {
		"on": "\u001b[3m",
		"off": "\u001b[23m"
	},
	"underline": {
		"on": "\u001b[4m",
		"off": "\u001b[24m"
	},
	"blink": {
		"slow": "\u001b[5m",
		"fast": "\u001b[6m",
		"off": "\u001b[25m"
	},
	"inverse": {
		"on": "\u001b[7m",
		"off": "\u001b[27m"
	},
	"conceal": {
		"on": "\u001b[8m",
		"off": "\u001b[28m"
	},
	"overscore": {
		"on": "\u001b[9m",
		"off": "\u001b[29m"
	},
	"frame": {
		"on": "\u001b[51m",
		"off": "\u001b[54m"
	},
	"circle": {
		"on": "\u001b[52m",
		"off": "\u001b[54m"
	},
	"overline": {
		"on": "\u001b[53m",
		"off": "\u001b[55m"
	},

	"clear": {
		"line": {
			"end": "\u001b[0K",
			"start": "\u001b[1K",
			"all": "\u001b[2K"
		},
		"screen": {
			"end": "\u001b[0J",
			"start": "\u001b[1J",
			"all": "\u001b[2J"
		}
	},

	"cursor": {
		"save": "\u001b[{s}",
		"restore": "\u001b[{u}",

		"up": "\u001b[1A",
		"down": "\u001b[1B",
		"right": "\u001b[1C",
		"left": "\u001b[1D",
		"prev-line": "\u001b[1E",
		"next-line": "\u001b[1F"
	},

	"fg": {
		"black": "\u001b[30m",
		"red": "\u001b[31m",
		"green": "\u001b[32m",
		"yellow": "\u001b[33m",
		"blue": "\u001b[34m",
		"magenta": "\u001b[35m", "pink": "\u001b[35m",
		"cyan": "\u001b[36m",
		"mwhite": "\u001b[37m",

		"gray": "\u001b[30;1m", "dwhite": "\u001b[30;1m",
		"lred": "\u001b[31;1m", "orange": "\u001b[31;1m",
		"lgreen": "\u001b[32;1m",
		"lyellow": "\u001b[33;1m",
		"lblue": "\u001b[34;1m",
		"lmagenta": "\u001b[35;1m", "lpink": "\u001b[35;1m",
		"lcyan": "\u001b[36;1m",
		"white": "\u001b[37;1m", "lwhite": "\u001b[37;1m"
	},

	"bg": {
		"black": "\u001b[40m",
		"red": "\u001b[41m",
		"green": "\u001b[42m",
		"yellow": "\u001b[43m",
		"blue": "\u001b[44m",
		"magenta": "\u001b[45m", "pink": "\u001b[45m",
		"cyan": "\u001b[46m",
		"mwhite": "\u001b[47m",

		"gray": "\u001b[40;1m", "dwhite": "\u001b[40;1m",
		"lred": "\u001b[41;1m", "orange": "\u001b[41;1m",
		"lgreen": "\u001b[42;1m",
		"lyellow": "\u001b[43;1m",
		"lblue": "\u001b[44;1m",
		"lmagenta": "\u001b[45;1m", "lpink": "\u001b[45;1m",
		"lcyan": "\u001b[46;1m",
		"white": "\u001b[47;1m", "lwhite": "\u001b[47;1m"
	}
}
`
