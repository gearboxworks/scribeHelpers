package toolRuntime

import "strings"

func (me *ExecArgs) ToString() string {
	return strings.Join(*me, " ")
}

