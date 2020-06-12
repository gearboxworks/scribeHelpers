package toolRuntime

import "strings"

func (r *ExecArgs) ToString() string {
	return strings.Join(*r, " ")
}

func (r *ExecArgs) String() string {
	return strings.Join(*r, " ")
}
