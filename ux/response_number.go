package ux

import "reflect"


//func (r *TypeResponse) AsInt() *int64 {
//	var s *int64
//	if IsReflectArray(r.data) {
//		s = ReflectInt(r.data)
//	}
//	//if s == nil {
//	//	var ptr int64
//	//	s = &ptr
//	//}
//	return s
//}
//
//func (r *TypeResponse) AsUint() *uint64 {
//	var s *uint64
//	if IsReflectArray(r.data) {
//		s = ReflectUint(r.data)
//	}
//	//if s == nil {
//	//	var ptr uint64
//	//	s = &ptr
//	//}
//	return s
//}
//
//func (r *TypeResponse) AsFloat() *float64 {
//	var s *float64
//	if IsReflectArray(r.data) {
//		s = ReflectFloat(r.data)
//	}
//	//if s == nil {
//	//	var ptr float64
//	//	s = &ptr
//	//}
//	return s
//}


func (r *TypeResponse) AsInt() *int {
	return r.Int
}

func (r *TypeResponse) AsInt8() *int8 {
	return r.Int8
}

func (r *TypeResponse) AsInt16() *int16 {
	return r.Int16
}

func (r *TypeResponse) AsInt32() *int32 {
	return r.Int32
}

func (r *TypeResponse) AsInt64() *int64 {
	return r.Int64
}

func (r *TypeResponse) AsUint() *uint {
	return r.Uint
}

func (r *TypeResponse) AsUint8() *uint8 {
	return r.Uint8
}

func (r *TypeResponse) AsUint16() *uint16 {
	return r.Uint16
}

func (r *TypeResponse) AsUint32() *uint32 {
	return r.Uint32
}

func (r *TypeResponse) AsUint64() *uint64 {
	return r.Uint64
}

func (r *TypeResponse) AsUintptr() *uintptr {
	return r.Uintptr
}

func (r *TypeResponse) AsFloat32() *float32 {
	return r.Float32
}

func (r *TypeResponse) AsFloat64() *float64 {
	return r.Float64
}

func (r *TypeResponse) AsComplex64() *complex64 {
	return r.Complex64
}

func (r *TypeResponse) AsComplex128() *complex128 {
	return r.Complex128
}


func IsReflectInt(i interface{}) bool {
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
	//		return true
	//	default:
	//		return false
	//}
	var ok bool
	switch i.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		ok = true
	}
	return ok
}

func IsReflectUint(i interface{}) bool {
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
	//		return true
	//	default:
	//		return false
	//}
	var ok bool
	switch i.(type) {
	case uint, uint8, uint16, uint32, uint64:
		ok = true
	}
	return ok
}

func IsReflectFloat(i interface{}) bool {
	//v := reflect.ValueOf(i)
	//switch v.Kind() {
	//	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
	//		return true
	//	default:
	//		return false
	//}
	var ok bool
	switch i.(type) {
	case float32, float64:
		ok = true
	}
	return ok
}


func ReflectInt(ref interface{}) *int64 {
	var s int64

	for range onlyOnce {
		value := reflect.ValueOf(ref)
		switch value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
			s = value.Int()
		default:
			s = 0
		}
	}

	return &s
}

func ReflectUint(ref interface{}) *uint64 {
	var s uint64

	value := reflect.ValueOf(ref)
	switch value.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		s = value.Uint()
	default:
		s = 0
	}

	return &s
}

func ReflectFloat(ref interface{}) *float64 {
	var s float64

	for range onlyOnce {
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
