package ux

import (
	//"fmt"
	"reflect"
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


type TypeReflect struct {
	Invalid       bool

	Bool          *bool

	Int           *int
	Int8          *int8
	Int16         *int16
	Int32         *int32
	Int64         *int64

	Uint          *uint
	Uint8         *uint8
	Uint16        *uint16
	Uint32        *uint32
	Uint64        *uint64

	Uintptr       *uintptr

	Float32       *float32
	Float64       *float64

	Complex64     *complex64
	Complex128    *complex128

	InterfaceArray *[]interface{}

	Func          *func()

	Interface     interface{}

	Map           *map[interface{}]interface{}

	Ptr           *interface{}

	Slice         *[]interface{}

	String        *string
	StringArray   *[]string

	Byte          *byte
	ByteArray     *[]byte

	Struct        *struct{}

	//UnsafePointer *UnsafePointer
	//Chan          *chan()
}

// type TypeReflect struct {
//	Invalid bool
//
//	Bool          *bool
//	Int           *int
//	Int8          *int8
//	Int16         *int16
//	Int32         *int32
//	Int64         *int64
//	Uint          *uint
//	Uint8         *uint8
//	Uint16        *uint16
//	Uint32        *uint32
//	Uint64        *uint64
//	Uintptr       *uintptr
//	Float32       *float32
//	Float64       *float64
//	Complex64     *complex64
//	Complex128    *complex128
//	Array         *[]interface{}
//	Chan          *chan()
//	Func          *func()
//	Interface     *interface{}
//	Map           *map[interface{}]interface{}
//	Ptr           *interface{}
//	Slice         *[]interface{}
//	String        *string
//	Struct        *struct{}
//	UnsafePointer *UnsafePointer
//}


type TypeResponse struct {
	ofType reflect.Type
	data   interface{}
	TypeReflect
}


type ResponseGetter interface {
	GetResponse() interface{}
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
		if state.debug.Enabled {
			v := reflect.ValueOf(r)
			PrintflnBlue("SetResponse() interface{} is a '%s' kind of type '%s'", v.Kind().String(), v.String())
		}

		ok = state.response.Set(r)

		if state.debug.Enabled {
			PrintflnBlue("Type: '%s' - '%s'", state.response.ofType.String(), state.response.ofType.Name())
		}
	}

	return ok
}


func (state *State) GetResponseType() reflect.Type {
	return state.response.ofType
}

func (state *State) GetResponseData() interface{} {
	return state.response.data
}


//func InspectStructV(val reflect.Value) {
//	if val.Kind() == reflect.Interface && !val.IsNil() {
//		elm := val.Elem()
//		if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
//			val = elm
//		}
//	}
//	if val.Kind() == reflect.Ptr {
//		val = val.Elem()
//	}
//
//	//for i := 0; i < val.NumField(); i++ {
//	//	valueField := val.Field(i)
//	//	typeField := val.Type().Field(i)
//		valueField := val
//		typeField := val.Type()
//		address := "not-addressable"
//
//		if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
//			elm := valueField.Elem()
//			if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
//				valueField = elm
//			}
//		}
//
//		if valueField.Kind() == reflect.Ptr {
//			valueField = valueField.Elem()
//
//		}
//		if valueField.CanAddr() {
//			address = fmt.Sprintf("0x%X", valueField.Addr().Pointer())
//		}
//
//		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Address: %v\t, Field type: %v\t, Field kind: %v\n",
//			typeField.Name,
//			valueField.Interface(),
//			address,
//			typeField,
//			valueField.Kind())
//
//		//if valueField.Kind() == reflect.Struct {
//		//	InspectStructV(valueField)
//		//}
//	//}
//}


func (r *TypeResponse) Set(ref interface{}) bool {
	var ok bool

	for range onlyOnce {
		v := reflect.ValueOf(ref)

		if v.Kind() != reflect.Ptr {
			PrintflnError("SetResponse requires a pointer type, but is a '%s' kind of type '%s'", v.Kind().String(), v.String())
			PrintflnError("Example: State.SetResponse(&xyzzy)")
			panic("ABORTING")
			break
		}

		s := v.Elem()
		r.ofType = s.Type()
		r.data = ref

		// @TODO MICKMAKE
		//fmt.Printf("String: %s\t", r.ofType.String())
		//fmt.Printf("Name: %s\t", r.ofType.Name())
		//fmt.Printf("Kind: %s\n", r.ofType.Kind())
		// @TODO MICKMAKE

		// If we have a pointer, then call again with the value of that pointer.
		if r.ofType.Kind().String() == "ptr" {
			ok = r.Set(s.Addr().Elem().Interface())
			break
		}

		r.TypeReflect = TypeReflect{}
		switch r.ofType.String() {
			case "bool":
				r.Bool = ref.(*bool)

			case "int":
				r.Int = ref.(*int)
			case "int8":
				r.Int8 = ref.(*int8)
			case "int16":
				r.Int16 = ref.(*int16)
			case "int32":
				r.Int32 = ref.(*int32)
			case "int64":
				r.Int64 = ref.(*int64)

			case "uint":
				r.Uint = ref.(*uint)
			//case uint8:
			//	r.Uint8 = ref.(*uint8)
			case "uint16":
				r.Uint16 = ref.(*uint16)
			case "uint32":
				r.Uint32 = ref.(*uint32)
			case "uint64":
				r.Uint64 = ref.(*uint64)

			case "uintptr":
				r.Uintptr = ref.(*uintptr)

			case "float32":
				r.Float32 = ref.(*float32)
			case "float64":
				r.Float64 = ref.(*float64)

			case "complex64":
				r.Complex64 = ref.(*complex64)
			case "complex128":
				r.Complex128 = ref.(*complex128)

			case "[]interface{}":
				r.InterfaceArray = ref.(*[]interface{})

			case "func()":
				r.Func = ref.(*func())

			case "map[interface{}]interface{}":
				r.Map = ref.(*map[interface{}]interface{})

			case "interface{}":
				*r.Ptr = ref.(*interface{})

			case "string":
				r.String = ref.(*string)
			case "[]string":
				r.StringArray = ref.(*[]string)

			case "byte":
				r.Byte = ref.(*byte)
			case "[]byte":
				r.ByteArray = ref.(*[]byte)

			case "struct{}":
				r.Struct = ref.(*struct{})
		}

		//switch ref.(type) {
		//	case bool:
		//		*r.Bool = ref.(bool)
		//
		//	case int:
		//		*r.Int = ref.(int)
		//	case int8:
		//		*r.Int8 = ref.(int8)
		//	case int16:
		//		*r.Int16 = ref.(int16)
		//	case int32:
		//		*r.Int32 = ref.(int32)
		//	case int64:
		//		*r.Int64 = ref.(int64)
		//
		//	case uint:
		//		*r.Uint = ref.(uint)
		//	//case uint8:
		//	//	*r.Uint8 = ref.(uint8)
		//	case uint16:
		//		*r.Uint16 = ref.(uint16)
		//	case uint32:
		//		*r.Uint32 = ref.(uint32)
		//	case uint64:
		//		*r.Uint64 = ref.(uint64)
		//
		//	case uintptr:
		//		*r.Uintptr = ref.(uintptr)
		//
		//	case float32:
		//		*r.Float32 = ref.(float32)
		//	case float64:
		//		*r.Float64 = ref.(float64)
		//
		//	case complex64:
		//		*r.Complex64 = ref.(complex64)
		//	case complex128:
		//		*r.Complex128 = ref.(complex128)
		//
		//	case []interface{}:
		//		*r.Array = ref.([]interface{})
		//
		//	case func():
		//		*r.Func = ref.(func())
		//
		//	case map[interface{}]interface{}:
		//		*r.Map = ref.(map[interface{}]interface{})
		//
		//	case interface{}:
		//		*r.Ptr = ref.(interface{})
		//
		//	case string:
		//		*r.String = ref.(string)
		//	case []string:
		//		*r.StringArray = ref.([]string)
		//
		//	case byte:
		//		*r.Byte = ref.(byte)
		//	case []byte:
		//		*r.ByteArray = ref.([]byte)
		//
		//	case struct{}:
		//		*r.Struct = ref.(struct{})
		//}

		ok = true
	}

	return ok
}


func (r *TypeResponse) GetType() reflect.Type {
	return r.ofType
}

func (r *TypeResponse) GetData() interface{} {
	return r.data
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

func (r *TypeResponse) IsNotOfType(t string) bool {
	return !r.IsOfType(t)
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
