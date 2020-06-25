package ux

import "strings"


//func (r *TypeResponse) AsString() *string {
//	var s *string
//	if IsReflectString(r.data) {
//		s = ReflectString(r.data)
//	}
//	//if s == nil {
//	//	var ptr string
//	//	s = &ptr
//	//}
//	return s
//}


func (r *TypeResponse) AsString() *string {
	return r.String
}

func (r *TypeResponse) AsByte() *byte {
	return r.Byte
}


func IsReflectString(i interface{}) bool {
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.String:
	//		return true
	//	default:
	//		return false
	//}
	var ok bool
	switch i.(type) {
	case string, *string:
		ok = true
	}
	return ok
}


func ReflectString(ref interface{}) *string {
	var s string

	for range onlyOnce {
		//value := reflect.ValueOf(ref)
		//if value.Kind() == reflect.String {
		//	st := value.String()
		//	s = &st
		//	break
		//}
		switch ref.(type) {
		case []byte:
			s = ref.(string)
		case string:
			s = ref.(string)
		case []string:
			s = strings.Join(ref.([]string), DefaultSeparator)
		}
	}

	return &s
}
