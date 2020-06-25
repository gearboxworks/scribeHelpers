package ux

import "reflect"


func IsReflectMap(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Map:
		return true
	default:
		return false
	}
}
