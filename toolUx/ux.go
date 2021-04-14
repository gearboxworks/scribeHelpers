package toolUx

import (
	"github.com/gearboxworks/scribeHelpers/ux"
)


// Was going to implement as a Getter interface, but no need.

//var _ ux.UxGetter = (*Ux)(nil)
//type Ux ux.Ux
//
//func (u *Ux) Open() error {
//	panic("implement me")
//}
//
//func (u *Ux) Close() {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfWhite(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfCyan(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfYellow(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfRed(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfGreen(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfBlue(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfMagenta(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfWhite(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfCyan(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfYellow(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfRed(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfGreen(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfBlue(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfMagenta(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) Sprintf(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) Printf(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfOk(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfOk(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfWarning(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfWarning(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfError(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfError(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) SprintError(err error) string {
//	panic("implement me")
//}
//
//func (u *Ux) PrintError(err error) {
//	panic("implement me")
//}
//
//func (u *Ux) GetTerminalSize() (int, int, error) {
//	panic("implement me")
//}


/////////////////////////////////////

func ToolPrintfWhite(format string, args ...interface{}) string {
	return ux.SprintfWhite(format, args...)
}

func ToolPrintfCyan(format string, args ...interface{}) string {
	return ux.SprintfCyan(format, args...)
}

func ToolPrintfYellow(format string, args ...interface{}) string {
	return ux.SprintfYellow(format, args...)
}

func ToolPrintfRed(format string, args ...interface{}) string {
	return ux.SprintfRed(format, args...)
}

func ToolPrintfGreen(format string, args ...interface{}) string {
	return ux.SprintfGreen(format, args...)
}

func ToolPrintfBlue(format string, args ...interface{}) string {
	return ux.SprintfBlue(format, args...)
}

func ToolPrintfMagenta(format string, args ...interface{}) string {
	return ux.SprintfMagenta(format, args...)
}

func ToolPrintf(format string, args ...interface{}) string {
	return ux.Sprintf(format, args...)
}

func ToolPrintfOk(format string, args ...interface{}) string {
	return ux.SprintfOk(format, args...)
}

func ToolPrintfWarning(format string, args ...interface{}) string {
	return ux.SprintfWarning(format, args...)
}

func ToolPrintfError(format string, args ...interface{}) string {
	return ux.SprintfError(format, args...)
}

func ToolPrintln() string {
	return ux.SprintfNormal("\n")
}


func ToolPrintflnWhite(format string, args ...interface{}) string {
	return ux.SprintfWhite(format, args...) + "\n"
}

func ToolPrintflnCyan(format string, args ...interface{}) string {
	return ux.SprintfCyan(format, args...) + "\n"
}

func ToolPrintflnYellow(format string, args ...interface{}) string {
	return ux.SprintfYellow(format, args...) + "\n"
}

func ToolPrintflnRed(format string, args ...interface{}) string {
	return ux.SprintfRed(format, args...) + "\n"
}

func ToolPrintflnGreen(format string, args ...interface{}) string {
	return ux.SprintfGreen(format, args...) + "\n"
}

func ToolPrintflnBlue(format string, args ...interface{}) string {
	return ux.SprintfBlue(format, args...) + "\n"
}

func ToolPrintflnMagenta(format string, args ...interface{}) string {
	return ux.SprintfMagenta(format, args...) + "\n"
}

func ToolPrintfln(format string, args ...interface{}) string {
	return ux.Sprintf(format, args...) + "\n"
}

func ToolPrintflnOk(format string, args ...interface{}) string {
	return ux.SprintfOk(format, args...) + "\n"
}

func ToolPrintflnWarning(format string, args ...interface{}) string {
	return ux.SprintfWarning(format, args...) + "\n"
}

func ToolPrintflnError(format string, args ...interface{}) string {
	return ux.SprintfError(format, args...) + "\n"
}


func ToolPrintError(err error) string {
	return ux.SprintError(err)
}
