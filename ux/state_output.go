package ux

import (
	"fmt"
	"regexp"
	"strings"
)


func (state *State) SetRunState(rs string) {
	state.RunState = rs
}
func (state *State) GetRunState() string {
	return state.RunState
}

func (state *State) RunStateEquals(format string, args ...interface{}) bool {
	var ret bool

	for range OnlyOnce {
		s := fmt.Sprintf(format, args...)
		if strings.Compare(state.RunState, s) == 0 {
			ret = true
		}
	}

	return ret
}


func (state *State) SetOutput(data ...interface{}) {
	for range OnlyOnce {
		state.Output = ""
		state.OutputArray = []string{}

		state.OutputAppend(data...)
	}
}

func (state *State) OutputAppend(data ...interface{}) {
	for range OnlyOnce {
		if state._Separator == "" {
			state._Separator = DefaultSeparator
		}

		for _, d := range data {
			//value := reflect.ValueOf(d)
			//switch value.Kind() {
			//	case reflect._Output:
			//		state._Array = append(state._Array, value._Output())
			//	case reflect.Array:
			//		state._Array = append(state._Array, d.([]string)...)
			//	case reflect.Slice:
			//		state._Array = append(state._Array, d.([]string)...)
			//}

			var sa []string
			switch d.(type) {
				case []string:
					for _, s := range d.([]string) {
						sa = append(sa, strings.Split(s, state._Separator)...)
					}
				case string:
					sa = append(sa, strings.Split(d.(string), state._Separator)...)
				case []byte:
					ts := d.([]byte)
					sa = append(sa, strings.Split(string(ts), state._Separator)...)
			}

			state.OutputArray = append(state.OutputArray, sa...)
		}
		state.Output = strings.Join(state.OutputArray, state._Separator)
	}
}

func (state *State) GetOutput() string {
	if state._Separator == "" {
		state._Separator = DefaultSeparator
	}

	return strings.Join(state.OutputArray, state._Separator)
}

func (state *State) GetOutputArray() []string {
	return state.OutputArray
}

func (state *State) SetSeparator(separator string) {
	for range OnlyOnce {
		state._Separator = separator
		state.OutputArray = strings.Split(state.Output, state._Separator)
	}
}

func (state *State) GetSeparator() string {
	return state._Separator
}

func (state *State) OutputTrim() {
	for range OnlyOnce {
		state.Output = strings.TrimSpace(state.Output)
		state.OutputArray = strings.Split(state.Output, state._Separator)
	}
}

func (state *State) OutputArrayTrim() {
	for range OnlyOnce {
		for _, s := range state.OutputArray {
			state.OutputArray = append(state.OutputArray, strings.Split(s, state._Separator)...)
		}
		state.Output = strings.Join(state.OutputArray, state._Separator)
	}
}

func (state *State) OutputEquals(format string, args ...interface{}) bool {
	var ret bool

	for range OnlyOnce {
		s := fmt.Sprintf(format, args...)
		if strings.Compare(state.Output, s) == 0 {
			ret = true
		}
	}

	return ret
}

func (state *State) OutputParse(format string, args ...interface{}) bool {
	var ret bool

	for range OnlyOnce {
		s := fmt.Sprintf(format, args...)

		ret = strings.Contains(state.Output, s)
	}

	return ret
}

func (state *State) OutputArrayGrep(format string, a ...interface{}) []string {
	var ret []string

	for range OnlyOnce {
		if len(state.OutputArray) == 0 {
			break
		}

		res := fmt.Sprintf(format, a...)
		re := regexp.MustCompile(res)
		for _, line := range state.OutputArray {
			if re.MatchString(line) {
				ret = append(ret, line)
			}
		}
	}

	return ret
}

func (state *State) OutputGrep(format string, a ...interface{}) string {
	var ret string

	for range OnlyOnce {
		if state.Output == "" {
			break
		}

		res := fmt.Sprintf(format, a...)
		re := regexp.MustCompile(res)
		if re.MatchString(state.Output) {
			ret = state.Output
		}
	}

	return ret
}
