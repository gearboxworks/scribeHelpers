package ux

import (
	"reflect"
)


func (r *TypeResponse) AsBool() *bool {
	return r.Bool
}


func IsReflectBool(i interface{}) bool {
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.Bool:
	//		return true
	//	default:
	//		return false
	//}
	var ok bool
	switch i.(type) {
	case bool:
		ok = true
	}
	return ok
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


func (r *TypeResponse) AsArray() *[]interface{} {
	return r.InterfaceArray
}

func (r *TypeResponse) AsFunc() *func() {
	return r.Func
}

func (r *TypeResponse) AsInterface() interface{} {
	return r.Interface
}

func (r *TypeResponse) AsMap() *map[interface{}]interface{} {
	return r.Map
}

func (r *TypeResponse) AsPtr() *interface{} {
	return r.Ptr
}

func (r *TypeResponse) AsSlice() *[]interface{} {
	return r.Slice
}

func (r *TypeResponse) AsStruct() *struct{} {
	return r.Struct
}

