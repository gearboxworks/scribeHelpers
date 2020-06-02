package helperTypes

import (
	"github.com/newclarity/scribe/ux"
	"reflect"
)


func HelperIsInt(i interface{}) bool {
	return ux.IsReflectInt(i)
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
	//		return true
	//	default:
	//		return false
	//}
}


func ReflectInt(ref interface{}) *int64 {
	var s int64

	for range OnlyOnce {
		value := reflect.ValueOf(ref)
		switch value.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
				s = value.Int()
			default:
				s = 0
		}
	}

	return &s
}


func ReflectInt32(ref interface{}) *int32 {
	var s int32

	for range OnlyOnce {
		value := reflect.ValueOf(ref)
		switch value.Kind() {
			case reflect.Int, reflect.Int32, reflect.Uint32:
				s = int32(value.Int())
			default:
				s = 0
		}
	}

	return &s
}


func ReflectFloat(ref interface{}) *float64 {
	var s float64

	for range OnlyOnce {
		value := reflect.ValueOf(ref)
		switch value.Kind() {
			case reflect.Float32, reflect.Float64:
				s = value.Float()
			default:
				s = 0
		}
	}

	return &s
}
