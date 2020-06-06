package ux

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

//
// Built-in string type:
// string
//
// Built-in boolean type:
// bool
//
// Built-in numeric types:
// int8
// uint8 (byte)
// int16
// uint16
// int32 (rune)
// uint32
// int64
// uint64
// int
// uint
// uintptr
// float32
// float64
// complex64
// complex128


type TypeResponse struct {
	ofType reflect.Type
	data   interface{}
}
func newResponse() TypeResponse {
	return TypeResponse{ ofType: nil, data: nil }
}
func (state *State) GetResponse() *TypeResponse {
	return &state.response
}
func (state *State) SetResponse(r interface{}) bool {
	var ok bool

	for range onlyOnce {
		v := reflect.ValueOf(r)
		if state.debug.Enabled {
			PrintflnBlue("SetResponse() interface{} is a '%s' kind of type '%s'", v.Kind().String(), v.String())
		}

		if v.Kind() != reflect.Ptr {
			PrintflnError("SetResponse requires a pointer type, but is a '%s' kind of type '%s'", v.Kind().String(), v.String())
			PrintflnError("Example: State.SetResponse(&xyzzy)")
			panic("ABORTING")
			break
		}

		s := v.Elem()
		state.response = TypeResponse{
			ofType: s.Type(),
			data:   r,
		}
		if state.debug.Enabled {
			PrintflnBlue("Type: '%s' - '%s'", s.Type().String(), s.Type().Name())
		}

		ok = true
	}

	return ok
}
func (r *TypeResponse) IsOfType(t string) bool {
	var ok bool

	//fmt.Printf("%s - %s\n", r.ofType.String(), r.ofType.Name())
	for range onlyOnce {
		if r.ofType.String() == t {
			ok = true
			break
		}

		if r.ofType.Name() == t {
			ok = true
			break
		}

		ok = false
	}

	return ok
}
func (state *State) GetResponseData() interface{} {
	return state.response.data
}
func (r *TypeResponse) GetData() interface{} {
	return r.data
}
func (state *State) GetResponseType() reflect.Type {
	return state.response.ofType
}
func (r *TypeResponse) GetType() reflect.Type {
	return r.ofType
}
func (r *TypeResponse) GetStringArray() *[]string {
	if r.IsOfType("[]string") {
		return (r.data).(*[]string)
	}

	return &[]string{}
}


func (r *TypeResponse) ResponseToString() *string {
	var s *string
	if IsReflectString(r.data) {
		s = ReflectString(r.data)
	}
	//if s == nil {
	//	var ptr string
	//	s = &ptr
	//}
	return s
}
func (r *TypeResponse) ResponseToArray() *[]string {
	var s *[]string
	if IsReflectArray(r.data) {
		s = ReflectStringArray(r.data)
	}
	//if s == nil {
	//	s = &[]string{}
	//}
	return s
}
func (r *TypeResponse) ResponseToByteArray() *[]byte {
	var s *[]byte
	if IsReflectArray(r.data) {
		s = ReflectByteArray(r.data)
	}
	//if s == nil {
	//	s = &[]byte{}
	//}
	return s
}
func (r *TypeResponse) ResponseToInt() *int64 {
	var s *int64
	if IsReflectArray(r.data) {
		s = ReflectInt(r.data)
	}
	//if s == nil {
	//	var ptr int64
	//	s = &ptr
	//}
	return s
}
func (r *TypeResponse) ResponseToUint() *uint64 {
	var s *uint64
	if IsReflectArray(r.data) {
		s = ReflectUint(r.data)
	}
	//if s == nil {
	//	var ptr uint64
	//	s = &ptr
	//}
	return s
}
func (r *TypeResponse) ResponseToFloat() *float64 {
	var s *float64
	if IsReflectArray(r.data) {
		s = ReflectFloat(r.data)
	}
	//if s == nil {
	//	var ptr float64
	//	s = &ptr
	//}
	return s
}


//type TypeReflect struct {
//	Invalid bool
//
//	Bool *bool
//	Int *int
//	Int8 *int8
//	Int16 *int16
//	Int32 *int32
//	Int64 *int64
//	Uint *uint
//	Uint8 *uint8
//	Uint16 *uint16
//	Uint32 *uint32
//	Uint64 *uint64
//	Uintptr *uintptr
//	Float32 *float32
//	Float64 *float64
//	Complex64 *complex64
//	Complex128 *complex128
//	Array *array
//	Chan *Chan
//	Func *Func
//	Interface *interface{}
//	Map *Map
//	Ptr *ptr
//	Slice *slice
//	String *String
//	Struct *Struct
//	UnsafePointer *UnsafePointer
//}
//func Reflect(ref interface{}) *TypeReflect {
//	var tr TypeReflect
//
//	// 	Invalid
//	//	Bool
//	//	Int
//	//	Int8
//	//	Int16
//	//	Int32
//	//	Int64
//	//	Uint
//	//	Uint8
//	//	Uint16
//	//	Uint32
//	//	Uint64
//	//	Uintptr
//	//	Float32
//	//	Float64
//	//	Complex64
//	//	Complex128
//	//	Array
//	//	Chan
//	//	Func
//	//	Interface
//	//	Map
//	//	Ptr
//	//	Slice
//	//	String
//	//	Struct
//	//	UnsafePointer
//
//	v := reflect.ValueOf(ref)
//	switch v.Kind() {
//		case reflect.Invalid:
//
//		case reflect.Bool:
//
//		case reflect.Int:
//			fallthrough
//		case reflect.Int8:
//			fallthrough
//		case reflect.Int16:
//			fallthrough
//		case reflect.Int32:
//			fallthrough
//		case reflect.Int64:
//
//		case reflect.Uint:
//			fallthrough
//		case reflect.Uint8:
//			fallthrough
//		case reflect.Uint16:
//			fallthrough
//		case reflect.Uint32:
//			fallthrough
//		case reflect.Uint64:
//
//		case reflect.Uintptr:
//
//		case reflect.Float32:
//			fallthrough
//		case reflect.Float64:
//
//		case reflect.Complex64:
//		case reflect.Complex128:
//
//		case reflect.Array:
//
//		case reflect.Chan:
//
//		case reflect.Func:
//
//		case reflect.Interface:
//
//		case reflect.Map:
//
//		case reflect.Ptr:
//
//		case reflect.Slice:
//
//		case reflect.String:
//
//		case reflect.Struct:
//
//		case reflect.UnsafePointer:
//	}
//
//	for range onlyOnce {
//		if IsReflectString(ref) {
//
//		}
//		switch ref.(type) {
//			case []byte:
//				s = ref.(string)
//				ref.()
//			case string:
//				s = ref.(string)
//			case []string:
//				s = strings.Join(ref.([]string), DefaultSeparator)
//		}
//	}
//
//	return &tr
//}


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
func IsReflectByteArray(i interface{}) bool {
	var ok bool
	switch i.(type) {
		case []byte:
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


func foo() {
	fmt.Printf("Go runs OK!\n")
	fmt.Printf("PPID: %d -> PID:%d\n", os.Getppid(), os.Getpid())
	fmt.Printf("Compiler: %s v%s\n", runtime.Compiler, runtime.Version())
	fmt.Printf("Architecture: %s v%s\n", runtime.GOARCH, runtime.GOOS)
	fmt.Printf("GOROOT: %s\n", runtime.GOROOT())
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


func IsReflectArray(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
		case reflect.Array:
			return true
		default:
			return false
	}
}


func IsReflectMap(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
		case reflect.Map:
			return true
		default:
			return false
	}
}
