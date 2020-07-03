package ux

import (
	"fmt"
	"runtime"
)

// We really, really, really want to avoid to ever call these functions.
// However, despite trying to write self-healing code, all the best intentions...

type RuntimeDebug struct {
	Enabled  bool
	File     string
	Line     int
	Function string
	Valid    bool
}
type Callers []*RuntimeDebug

func (p *RuntimeDebug) fetchRuntimeDebug(level int) {
	for range onlyOnce {
		if p == nil {
			break
		}
		if level == 0 {
			level = 1
		}

		// Discover package name.
		var pc uintptr
		pc, p.File, p.Line, p.Valid = runtime.Caller(level)
		if p.Valid {
			details := runtime.FuncForPC(pc)
			p.Function = details.Name()
			//f, l := details.FileLine(pc)
			//fmt.Printf("%s:%d - %s:%d\n",
			//	p.TypeFile,
			//	p.Line,
			//	f,
			//	l,
			//	)
		}
		//fmt.Printf("DEBUG => %s:%d [%s]\n", p.TypeFile, p.Line, p.Function)
	}
}

func fetchRuntimeDebug(level int) *RuntimeDebug {
	var p RuntimeDebug
	for range onlyOnce {
		if level == 0 {
			level = 1
		}

		// Discover package name.
		var pc uintptr
		pc, p.File, p.Line, p.Valid = runtime.Caller(level)
		if p.Valid {
			details := runtime.FuncForPC(pc)
			p.Function = details.Name()
			//f, l := details.FileLine(pc)
			//fmt.Printf("%s:%d - %s:%d\n",
			//	p.TypeFile,
			//	p.Line,
			//	f,
			//	l,
			//	)
		}
		//fmt.Printf("DEBUG => %s:%d [%s]\n", p.TypeFile, p.Line, p.Function)
	}
	return &p
}


const (
	PanicErrorPrefix = "SW ERROR: "
	PanicErrorNotGivenAPointer = PanicErrorPrefix + "Not given a pointer to structure '%s'"
	PanicErrorGivenANilFunction = PanicErrorPrefix + "This is a nil function"
)

func Panic(format string, args ...interface{}) {
	err := fmt.Sprintf(format, args...)
	PrintfError(err)
	PanicDump()
	panic(err)
}

func StatePanic(state *State) {
	Panic(state.SprintError())
}

func GetCallers(from int, until int) Callers {
	var ret Callers

	if from == 0 {
		from = 1
	}

	for i := from; i < until; i++ {
		c := fetchRuntimeDebug(i)
		if !c.Valid {
			break
		}
		ret = append(ret, c)
	}

	return ret
}

func (p RuntimeDebug) String() string {
	ret, _ := TemplateSprintf("{{ BrightWhite }}File: {{ BrightCyan }}%s{{ Reset }}:{{ BrightYellow }}%d{{ BrightWhite }} Function: {{ BrightBlue }}%s{{ Reset }}\n",
		p.File,
		p.Line,
		p.Function)
	return ret
}

func (cs Callers) String() string {
	var ret string
	for _, c := range cs {
		ret += fmt.Sprintf("%s", c)
	}
	return ret
}

func PanicDump() {
	cs := GetCallers(2, 50)
	PrintflnYellow("\nPanic Dump:")
	fmt.Printf("%s", cs.String())
}