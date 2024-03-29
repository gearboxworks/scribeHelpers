package toolTypes

import (
	"encoding/json"
	"fmt"
	"github.com/gearboxworks/scribeHelpers/ux"
	"regexp"
	"strings"
)

type TypeGenericString struct {
	Valid bool
	Error error
	String string
}

type TypeGenericStringArray struct {
	Valid bool
	Error error
	Array []string
}


// Usage:
//		{{ if IsString $output }}YES{{ end }}
func ToolIsString(i interface{}) bool {
	return ux.IsReflectString(i)
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.String:
	//		return true
	//	default:
	//		return false
	//}
}


// Usage:
//		{{ $str := ToUpper "lowercase" }}
func ToolToUpper(i interface{}) string {
	return strings.ToUpper(*ux.ReflectString(i))
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.String:
	//		return strings.ToUpper(i.(string))
	//	default:
	//		return ""
	//}
}


// Usage:
//		{{ $str := ToLower "UPPERCASE" }}
func ToolToLower(i interface{}) string {
	return strings.ToLower(*ux.ReflectString(i))
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.String:
	//		return strings.ToLower(i.(string))
	//	default:
	//		return ""
	//}
}


// Usage:
//		{{ $str := ToString .Json.array }}
func ToolToString(i interface{}) string {
	ret := ""
	var j []byte
	var err error
	j, err = json.Marshal(i)
	if err == nil {
		ret = string(j)
	}
	return ret
}


// Usage:
//		{{ if ExecParseOutput $output "uid=%s" "mick" ... }}YES{{ end }}
func ToolContains(s interface{}, substr interface{}) bool {
	var ret bool

	for range onlyOnce {
		sp := ReflectString(s)
		if sp == nil {
			break
		}

		ssp := ReflectString(substr)
		if ssp == nil {
			break
		}

		ret = strings.Contains(*sp, *ssp)
	}

	return ret
}


// Usage:
//		{{ Sprintf "uid=%s" "mick" ... }}
func ToolSprintf(format interface{}, a ...interface{}) string {
	var ret string

	for range onlyOnce {
		p := ReflectString(format)
		if p == nil {
			break
		}
		ret = fmt.Sprintf(*p, a...)
	}

	return ret
}


// Usage:
//		{{ Grep .This.output "uid=%s" "mick" ... }}
func ToolGrepArray(str interface{}, format interface{}, a ...interface{}) []string {
	var ret []string

	for range onlyOnce {
		s := ReflectString(str)
		if s == nil {
			break
		}
		text := strings.Split(*s, "\n")

		f := ReflectString(format)
		if f == nil {
			break
		}

		res := fmt.Sprintf(*f, a...)
		re := regexp.MustCompile(res)

		for _, line := range text {
			if re.MatchString(line) {
				ret = append(ret, line)
			}
		}
	}

	return ret
}


// Usage:
//		{{ Grep .This.output "uid=%s" "mick" ... }}
func ToolGrep(str interface{}, format interface{}, a ...interface{}) string {
	var ret string

	for range onlyOnce {
		sa := ToolGrepArray(str, format, a...)

		ret = strings.Join(sa, "\n")
	}

	return ret
}
