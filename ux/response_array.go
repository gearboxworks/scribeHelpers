package ux

import (
	"reflect"
	"strings"
)


//func (r *TypeResponse) GetStringArray() *[]string {
//	if r.IsOfType("[]string") {
//		return (r.data).(*[]string)
//	}
//
//	return &[]string{}
//}
//
//func (r *TypeResponse) AsStringArray() *[]string {
//	var s *[]string
//	if IsReflectArray(r.data) {
//		s = ReflectStringArray(r.data)
//	}
//	//if s == nil {
//	//	s = &[]string{}
//	//}
//	return s
//}
//
//func (r *TypeResponse) AsByteArray() *[]byte {
//	var s *[]byte
//	if IsReflectArray(r.data) {
//		s = ReflectByteArray(r.data)
//	}
//	//if s == nil {
//	//	s = &[]byte{}
//	//}
//	return s
//}

func (r *TypeResponse) AsStringArray() *[]string {
	return r.StringArray
}

func (r *TypeResponse) AsByteArray() *[]byte {
	return r.ByteArray
}


func IsReflectArray(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Array:
		return true
	default:
		return false
	}
}

func IsReflectByteArray(i interface{}) bool {
	var ok bool
	switch i.(type) {
	case []byte:
		ok = true
	}
	return ok
}

func IsReflectSlice(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Slice:
		return true
	default:
		return false
	}
}


func ReflectStringArray(ref ...interface{}) *[]string {
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
		//value := reflect.ValueOf(ref)
		//if value.Kind() != reflect.String {
		//	break
		//}
		//sa := []byte(value.String())
		//s = &sa
		switch ref.(type) {
		case []byte:
			s = ref.([]byte)
		case string:
			s = ref.([]byte)
		case []string:
			s = []byte((strings.Join(ref.([]string), DefaultSeparator)))
		}
	}

	return &s
}
