package toolTypes

import (
	"github.com/gearboxworks/scribeHelpers/ux"
	"reflect"
	"strings"
)


//func ReflectString(ref interface{}) *string {
//	var s string
//
//	for range onlyOnce {
//		switch ref.(type) {
//			case []byte:
//				s = ref.(string)
//			case string:
//				s = ref.(string)
//			case []string:
//				s = strings.Join(ref.([]string), ux.DefaultSeparator)
//		}
//		//value := reflect.ValueOf(ref)
//		//if value.Kind() == reflect.String {
//		//	st := value.String()
//		//	s = &st
//		//	break
//		//}
//	}
//
//	return &s
//}

func ReflectString(ref interface{}) *string {
	var s *string

	for range onlyOnce {
		//value := reflect.ValueOf(ref)
		//if value.Kind() == reflect.String {
		//	st := value.String()
		//	s = &st
		//	break
		//}
		switch ref.(type) {
			case []byte:
				sr := ref.(string)
				s = &sr
			case string:
				sr := ref.(string)
				s = &sr
			case []string:
				sr := strings.Join(ref.([]string), ux.DefaultSeparator)
				s = &sr
		}
	}

	return s
}

func ReflectStrings(ref ...interface{}) *[]string {
	var sa []string

	for range onlyOnce {
		for _, r := range ref {
			sa = append(sa, *ReflectString(r))
		}
	}

	return &sa
}

func ReflectByteArray(ref interface{}) *[]byte {
	var s []byte

	for range onlyOnce {
		switch ref.(type) {
			case []byte:
				s = ref.([]byte)
			case string:
				s = ref.([]byte)
			case []string:
				s = []byte((strings.Join(ref.([]string), ux.DefaultSeparator)))
		}
		//value := reflect.ValueOf(ref)
		//if value.Kind() != reflect.String {
		//	break
		//}
		//sa := []byte(value.String())
		//s = &sa
	}

	return &s
}

func ReflectBool(ref interface{}) *bool {
	var b *bool

	for range onlyOnce {
		value := reflect.ValueOf(ref)
		if value.Kind() != reflect.Bool {
			break
		}

		ba := value.Bool()
		b = &ba
	}

	return b
}

func ReflectBoolArg(ref interface{}) bool {
	var s bool

	for range onlyOnce {
		value := reflect.ValueOf(ref)
		switch value.Kind() {
			case reflect.Bool:
				s = value.Bool()
			case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
				v := value.Int()
				if v == 0 {
					s = false
				} else {
					s = true
				}
			case reflect.Float32, reflect.Float64:
				v := value.Float()
				if v == 0 {
					s = false
				} else {
					s = true
				}
		}
	}

	return s
}
