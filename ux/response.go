package ux

import (
	"fmt"
	"reflect"
)


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

type TypeReflect struct {
	Valid         bool

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

	Func          		*func()
	FuncReturn    		*func() *TypeResponse
	FuncVariadic  		*func(args ...interface{})
	FuncVariadicReturn  *func(args ...interface{}) *TypeResponse

	Interface     interface{}

	Map           *map[interface{}]interface{}

	Ptr           *interface{}

	Slice         *[]interface{}

	TypeString    *string
	StringArray   *[]string

	Byte          *byte
	ByteArray     *[]byte

	Struct        *struct{}

	//UnsafePointer *UnsafePointer
	//Chan          *chan()
}


type TypeResponse struct {
	ofType reflect.Type
	data   interface{}
	TypeReflect
}


type ResponseGetter interface {
	GetResponse() interface{}
}


func NewResponse() *TypeResponse {
	return &TypeResponse{ ofType: nil, data: nil, TypeReflect: TypeReflect{} }
}


func (state *State) GetResponse() *TypeResponse {
	return &state.response
}


func (state *State) SetResponse(r interface{}) bool {
	return state.response.Set(r)
}


func (state *State) GetResponseType() *reflect.Type {
	return &state.response.ofType
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
//
//
//func (r *TypeResponse) Set(ref interface{}) bool {
//	var ok bool
//
//	for range onlyOnce {
//		v := reflect.ValueOf(ref)
//		PrintflnGreen("ref - v.Type().Name():%s\tv.Type().String():%s\tv.String():%s\tv.Kind():%s",
//			v.Type().Name(),
//			v.Type().String(),
//			v.String(),
//			v.Kind(),
//		)
//
//		if v.Kind() == reflect.Ptr {
//			r.TypeReflect = TypeReflect{}
//			s := v.Elem()
//			r.ofType = s.Type()
//			r.data = v.Interface()
//
//			// If we have a pointer, then call again with the value of that pointer.
//			if r.ofType.Kind().String() == "ptr" {
//				ok = r.Set(s.Addr().Elem().Interface())
//				break
//			}
//
//			if r.setter(ref) {
//				break
//			}
//		}
//
//		if v.Kind() != reflect.Ptr {
//			r.TypeReflect = TypeReflect{}
//			PrintflnError("SetResponse requires a pointer type, but is a '%s' kind of type '%s'", v.Kind().String(), v.Type().String())
//			//PrintflnError("Example: State.SetResponse(&xyzzy)")
//			//if !r.Set(&ref) {
//			//	panic("ABORTING")
//			//}
//			foo := v.Convert(v.Type())
//			PrintflnError("foo is a '%s' kind of type '%s'", foo.Kind().String(), foo.String())
//
//			//foo2 := v.Addr()
//			//PrintflnError("foo is a '%s' kind of type '%s'", foo2.Kind().String(), foo2.String())
//
//			//refPtr := &ref
//			//v = reflect.ValueOf(refPtr)
//			//PrintflnError("refPtr is a '%s' kind of type '%s'", v.Kind().String(), v.String())
//
//			r.ofType = reflect.TypeOf(ref)
//			r.data = ref
//			//switch r.ofType.String() {
//			//	case TypeBool:
//			//		r.Bool = r.data.(*bool)
//			//
//			//	case TypeInt:
//			//		r.Int = r.data.(*int)
//			//	case TypeInt8:
//			//		r.Int8 = r.data.(*int8)
//			//	case TypeInt16:
//			//		r.Int16 = r.data.(*int16)
//			//	case TypeInt32:
//			//		r.Int32 = r.data.(*int32)
//			//	case TypeInt64:
//			//		r.Int64 = r.data.(*int64)
//			//
//			//	case TypeUint:
//			//		r.Uint = r.data.(*uint)
//			//	//case TypeUint8:
//			//	//	r.Uint8 = r.data.(*uint8)
//			//	case TypeUint16:
//			//		r.Uint16 = r.data.(*uint16)
//			//	case TypeUint32:
//			//		r.Uint32 = r.data.(*uint32)
//			//	case TypeUint64:
//			//		r.Uint64 = r.data.(*uint64)
//			//
//			//	case TypeUintptr:
//			//		r.Uintptr = r.data.(*uintptr)
//			//
//			//	case TypeFloat32:
//			//		r.Float32 = r.data.(*float32)
//			//	case TypeFloat64:
//			//		r.Float64 = r.data.(*float64)
//			//
//			//	case TypeComplex64:
//			//		r.Complex64 = r.data.(*complex64)
//			//	case TypeComplex128:
//			//		r.Complex128 = r.data.(*complex128)
//			//
//			//	case TypeInterfaceArray:
//			//		r.InterfaceArray = r.data.(*[]interface{})
//			//
//			//	case TypeFunc:
//			//		r.Func = r.data.(*func())
//			//
//			//	case TypeMap:
//			//		r.Map = r.data.(*map[interface{}]interface{})
//			//
//			//	case TypeInterface:
//			//		r.Interface = r.data.(*interface{})
//			//
//			//	case TypeString:
//			//		r.TypeString = r.data.(*string)
//			//	case TypeStringArray:
//			//		t := r.data.([]string)
//			//		r.StringArray = &t
//			//
//			//	case TypeByte:
//			//		r.Byte = r.data.(*byte)
//			//	case TypeByteArray:
//			//		r.ByteArray = r.data.(*[]byte)
//			//
//			//	case TypeStruct:
//			//		r.Struct = r.data.(*struct{})
//			//}
//
//			if r.setter(ref) {
//				break
//			}
//
//		}
//
//		if v.Kind() == reflect.Func {
//			r.TypeReflect = TypeReflect{}
//			r.ofType = reflect.TypeOf(ref)
//			r.data = ref
//
//			//fmt.Printf(">>%s<<\t>>%s<<\t>>%s<<\n",
//			//	r.ofType.String(),
//			//	r.ofType.Name(),
//			//	r.ofType.Kind(),
//			//	)
//			ok = false
//			switch r.ofType.String() {
//			case TypeFunc:
//				f := ref.(func())
//				r.Func = &f
//				ok = true
//			case TypeFuncReturn:
//				f := ref.(func() *TypeResponse)
//				r.FuncReturn = &f
//				ok = true
//			case TypeFuncVariadic:
//				f := ref.(func(args ...interface{}))
//				r.FuncVariadic = &f
//				ok = true
//			case TypeFuncVariadicReturn:
//				f := ref.(func(args ...interface{}) *TypeResponse)
//				r.FuncVariadicReturn = &f
//				ok = true
//			}
//			break
//		}
//
//
//
//		//switch v.Kind() {
//		//	case reflect.Ptr:
//		//		// OK.
//		//	case reflect.Func:
//		//		f := ref.(func())
//		//		r.Func = &f
//		//		ok = true
//		//		break
//		//		//// fptr is a pointer to a function.
//		//		//// Obtain the function value itself (likely nil) as a reflect.Value
//		//		//// so that we can query its type and then set the value.
//		//		//v = reflect.ValueOf(ref).Elem()
//		//		//swap := func(in []reflect.Value) []reflect.Value {
//		//		//	return []reflect.Value{in[1], in[0]}
//		//		//}
//		//		//// Make a function of the right type.
//		//		//v = reflect.MakeFunc(v.Type(), swap)
//		//		//// OK.
//		//	default:
//		//		PrintflnError("SetResponse requires a pointer type, but is a '%s' kind of type '%s'", v.Kind().String(), v.String())
//		//		PrintflnError("Example: State.SetResponse(&xyzzy)")
//		//		panic("ABORTING")
//		//}
//		// @TODO MICKMAKE
//		//fmt.Printf("String: %s\t", r.ofType.String())
//		//fmt.Printf("Name: %s\t", r.ofType.Name())
//		//fmt.Printf("Kind: %s\n", r.ofType.Kind())
//		// @TODO MICKMAKE
//		//switch ref.(type) {
//		//	case bool:
//		//		*r.Bool = ref.(bool)
//		//
//		//	case int:
//		//		*r.Int = ref.(int)
//		//	case int8:
//		//		*r.Int8 = ref.(int8)
//		//	case int16:
//		//		*r.Int16 = ref.(int16)
//		//	case int32:
//		//		*r.Int32 = ref.(int32)
//		//	case int64:
//		//		*r.Int64 = ref.(int64)
//		//
//		//	case uint:
//		//		*r.Uint = ref.(uint)
//		//	//case uint8:
//		//	//	*r.Uint8 = ref.(uint8)
//		//	case uint16:
//		//		*r.Uint16 = ref.(uint16)
//		//	case uint32:
//		//		*r.Uint32 = ref.(uint32)
//		//	case uint64:
//		//		*r.Uint64 = ref.(uint64)
//		//
//		//	case uintptr:
//		//		*r.Uintptr = ref.(uintptr)
//		//
//		//	case float32:
//		//		*r.Float32 = ref.(float32)
//		//	case float64:
//		//		*r.Float64 = ref.(float64)
//		//
//		//	case complex64:
//		//		*r.Complex64 = ref.(complex64)
//		//	case complex128:
//		//		*r.Complex128 = ref.(complex128)
//		//
//		//	case []interface{}:
//		//		*r.Array = ref.([]interface{})
//		//
//		//	case func():
//		//		*r.Func = ref.(func())
//		//
//		//	case map[interface{}]interface{}:
//		//		*r.Map = ref.(map[interface{}]interface{})
//		//
//		//	case interface{}:
//		//		*r.Ptr = ref.(interface{})
//		//
//		//	case string:
//		//		*r.String = ref.(string)
//		//	case []string:
//		//		*r.StringArray = ref.([]string)
//		//
//		//	case byte:
//		//		*r.Byte = ref.(byte)
//		//	case []byte:
//		//		*r.ByteArray = ref.([]byte)
//		//
//		//	case struct{}:
//		//		*r.Struct = ref.(struct{})
//		//}
//
//		ok = true
//	}
//
//	return ok
//}


func (r *TypeResponse) Set(ref interface{}) bool {
	for range onlyOnce {
		v := reflect.ValueOf(ref)

		//PrintflnGreen("ref - v.Type().Name():%s\tv.Type().String():%s\tv.String():%s\tv.Kind():%s",
		//	v.Type().Name(),
		//	v.Type().String(),
		//	v.String(),
		//	v.Kind(),
		//)

		// We expect a pointer type to be given, but we can also handle a value type, (see below).
		if v.Kind() == reflect.Ptr {
			r.TypeReflect = TypeReflect{}
			s := v.Elem()
			r.ofType = s.Type()
			r.data = v.Interface()

			// If we have a pointer to a pointer, then call again with the value of that pointer.
			if r.ofType.Kind().String() == "ptr" {
				PrintflnCyan("POINTER TO POINTER")
				r.Valid = r.Set(s.Addr().Elem().Interface())
				break
			}

			r.Valid = r.setter(ref)
			break
		}

		// Functions are special types.
		// We support only four types - TypeFunc, TypeFuncReturn, TypeFuncVariadic, TypeFuncVariadicReturn
		if v.Kind() == reflect.Func {
			r.TypeReflect = TypeReflect{}
			r.ofType = reflect.TypeOf(ref)
			r.data = ref

			switch r.ofType.String() {
				case TypeFunc:
					f := ref.(func())
					r.Func = &f
					r.Valid = true
				case TypeFuncReturn:
					f := ref.(func() *TypeResponse)
					r.FuncReturn = &f
					r.Valid = true
				case TypeFuncVariadic:
					f := ref.(func(args ...interface{}))
					r.FuncVariadic = &f
					r.Valid = true
				case TypeFuncVariadicReturn:
					f := ref.(func(args ...interface{}) *TypeResponse)
					r.FuncVariadicReturn = &f
					r.Valid = true
				default:
					r.Valid = false
			}

			break
		}


		// We expect a pointer type to be given, but we can also handle a value type, (see below).
		r.TypeReflect = TypeReflect{}
		r.ofType = reflect.TypeOf(ref)
		r.data = ref

		//PrintflnError("SetResponse requires a pointer type, but is a '%s' kind of type '%s'", v.Kind().String(), v.Type().String())
		////PrintflnError("Example: State.SetResponse(&xyzzy)")
		////if !r.Set(&ref) {
		////	panic("ABORTING")
		////}
		//foo := v.Convert(v.Type())
		//PrintflnError("foo is a '%s' kind of type '%s'", foo.Kind().String(), foo.String())
		//
		////foo2 := v.Addr()
		////PrintflnError("foo is a '%s' kind of type '%s'", foo2.Kind().String(), foo2.String())
		//
		////refPtr := &ref
		////v = reflect.ValueOf(refPtr)
		////PrintflnError("refPtr is a '%s' kind of type '%s'", v.Kind().String(), v.String())
		//
		//r.ofType = reflect.TypeOf(ref)
		//r.data = ref

		r.Valid = r.setter(ref)
		break
	}

	return r.Valid
}


// I KNOW there's a better way to do this, but for now. It is what it is.
// ref can either be a pointer to a type value, or a type value.
// We can't blindly take the pointer of the type value as we need to know what it is before we do.
// Chicken and egg issue...
func (r *TypeResponse) setter(ref interface{}) bool {

	for range onlyOnce {
		v := reflect.ValueOf(ref)
		r.TypeReflect = TypeReflect{}
		switch v.Type().String() {
			// Two different ways of obtaining a pointer to a value.
			//p := v.Interface().([]string)
			//r.StringArray = &p
			//p = r.data.([]string)
			//r.StringArray = &p

			case TypeBool:
				p := r.data.(bool)
				r.Bool = &p
				r.Valid = true

			case TypeInt:
				p := r.data.(int)
				r.Int = &p
				r.Valid = true

			case TypeInt8:
				p := r.data.(int8)
				r.Int8 = &p
				r.Valid = true

			case TypeInt16:
				p := r.data.(int16)
				r.Int16 = &p
				r.Valid = true

			case TypeInt32:
				p := r.data.(int32)
				r.Int32 = &p
				r.Valid = true

			case TypeInt64:
				p := r.data.(int64)
				r.Int64 = &p
				r.Valid = true

			case TypeUint:
				p := r.data.(uint)
				r.Uint = &p
				r.Valid = true

			case TypeUint8:
				p := r.data.(uint8)
				r.Uint8 = &p
				r.Valid = true

			case TypeUint16:
				p := r.data.(uint16)
				r.Uint16 = &p
				r.Valid = true

			case TypeUint32:
				p := r.data.(uint32)
				r.Uint32 = &p
				r.Valid = true

			case TypeUint64:
				p := r.data.(uint64)
				r.Uint64 = &p
				r.Valid = true

			case TypeUintptr:
				p := r.data.(uintptr)
				r.Uintptr = &p
				r.Valid = true

			case TypeFloat32:
				p := r.data.(float32)
				r.Float32 = &p
				r.Valid = true

			case TypeFloat64:
				p := r.data.(float64)
				r.Float64 = &p
				r.Valid = true

			case TypeComplex64:
				p := r.data.(complex64)
				r.Complex64 = &p
				r.Valid = true

			case TypeComplex128:
				p := r.data.(complex128)
				r.Complex128 = &p
				r.Valid = true

			case TypeInterfaceArray:
				p := r.data.([]interface{})
				r.InterfaceArray = &p
				r.Valid = true

			case TypeFunc:
				p := r.data.(func())
				r.Func = &p
				r.Valid = true

			case TypeMap:
				p := r.data.(map[interface{}]interface{})
				r.Map = &p
				r.Valid = true

			case TypeInterface:
				r.Interface = r.data.(interface{})
				r.Valid = true

			case TypeString:
				p := r.data.(string)
				r.TypeString = &p
				r.Valid = true

			case TypeStringArray:
				p := r.data.([]string)
				r.StringArray = &p
				r.Valid = true

			case TypeByte:
				p := r.data.(byte)
				r.Byte = &p
				r.Valid = true

			case TypeByteArray:
				p := r.data.([]byte)
				r.ByteArray = &p
				r.Valid = true

			case TypeStruct:
				p := r.data.(struct{})
				r.Struct = &p
				r.Valid = true


			case "*" + TypeBool:
				r.Bool = r.data.(*bool)
				r.Valid = true

			case "*" + TypeInt:
				r.Int = r.data.(*int)
				r.Valid = true

			case "*" + TypeInt8:
				r.Int8 = r.data.(*int8)
				r.Valid = true

			case "*" + TypeInt16:
				r.Int16 = r.data.(*int16)
				r.Valid = true

			case "*" + TypeInt32:
				r.Int32 = r.data.(*int32)
				r.Valid = true

			case "*" + TypeInt64:
				r.Int64 = r.data.(*int64)
				r.Valid = true

			case "*" + TypeUint:
				r.Uint = r.data.(*uint)
				r.Valid = true

			case "*" + TypeUint8:
				r.Uint8 = r.data.(*uint8)
				r.Valid = true

			case "*" + TypeUint16:
				r.Uint16 = r.data.(*uint16)
				r.Valid = true

			case "*" + TypeUint32:
				r.Uint32 = r.data.(*uint32)
				r.Valid = true

			case "*" + TypeUint64:
				r.Uint64 = r.data.(*uint64)
				r.Valid = true

			case "*" + TypeUintptr:
				r.Uintptr = r.data.(*uintptr)
				r.Valid = true

			case "*" + TypeFloat32:
				r.Float32 = r.data.(*float32)
				r.Valid = true

			case "*" + TypeFloat64:
				r.Float64 = r.data.(*float64)
				r.Valid = true

			case "*" + TypeComplex64:
				r.Complex64 = r.data.(*complex64)
				r.Valid = true

			case "*" + TypeComplex128:
				r.Complex128 = r.data.(*complex128)
				r.Valid = true

			case "*" + TypeInterfaceArray:
				r.InterfaceArray = r.data.(*[]interface{})
				r.Valid = true

			case "*" + TypeFunc:
				r.Func = r.data.(*func())
				r.Valid = true

			case "*" + TypeMap:
				r.Map = r.data.(*map[interface{}]interface{})
				r.Valid = true

			case "*" + TypeInterface:
				r.Interface = r.data.(*interface{})
				r.Valid = true

			case "*" + TypeString:
				r.TypeString = r.data.(*string)
				r.Valid = true

			case "*" + TypeStringArray:
				r.StringArray = r.data.(*[]string)
				r.Valid = true

			case "*" + TypeByte:
				r.Byte = r.data.(*byte)
				r.Valid = true

			case "*" + TypeByteArray:
				r.ByteArray = r.data.(*[]byte)
				r.Valid = true

			case "*" + TypeStruct:
				r.Struct = r.data.(*struct{})
				r.Valid = true
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
	}

	return r.Valid
}


func (r TypeResponse) getter(getPtr bool) interface{} {
	var ret interface{}
	switch {
		case r.Bool != nil:
			if getPtr {
				ret = r.Bool
				break
			}
			ret = *r.Bool
		case r.Int != nil:
			if getPtr {
				ret = r.Int
				break
			}
			ret = *r.Int
		case r.Int8 != nil:
			if getPtr {
				ret = r.Int8
				break
			}
			ret = *r.Int8
		case r.Int16 != nil:
			if getPtr {
				ret = r.Int16
				break
			}
			ret = *r.Int16
		case r.Int32 != nil:
			if getPtr {
				ret = r.Int32
				break
			}
			ret = *r.Int32
		case r.Int64 != nil:
			if getPtr {
				ret = r.Int64
				break
			}
			ret = *r.Int64
		case r.Uint != nil:
			if getPtr {
				ret = r.Uint
				break
			}
			ret = *r.Uint
		case r.Uint8 != nil:
			if getPtr {
				ret = r.Uint8
				break
			}
			ret = *r.Uint8
		case r.Uint16 != nil:
			if getPtr {
				ret = r.Uint16
				break
			}
			ret = *r.Uint16
		case r.Uint32 != nil:
			if getPtr {
				ret = r.Uint32
				break
			}
			ret = *r.Uint32
		case r.Uint64 != nil:
			if getPtr {
				ret = r.Uint64
				break
			}
			ret = *r.Uint64
		case r.Uintptr != nil:
			if getPtr {
				ret = r.Uintptr
				break
			}
			ret = *r.Uintptr
		case r.Float32 != nil:
			if getPtr {
				ret = r.Float32
				break
			}
			ret = *r.Float32
		case r.Float64 != nil:
			if getPtr {
				ret = r.Float64
				break
			}
			ret = *r.Float64
		case r.Complex64 != nil:
			if getPtr {
				ret = r.Complex64
				break
			}
			ret = *r.Complex64
		case r.Complex128 != nil:
			if getPtr {
				ret = r.Complex128
				break
			}
			ret = *r.Complex128
		case r.InterfaceArray != nil:
			if getPtr {
				ret = r.InterfaceArray
				break
			}
			ret = *r.InterfaceArray
		case r.Func != nil:
			if getPtr {
				ret = r.Func
				break
			}
			ret = *r.Func
		case r.FuncReturn != nil:
			if getPtr {
				ret = r.FuncReturn
				break
			}
			ret = *r.FuncReturn
		case r.FuncVariadic != nil:
			if getPtr {
				ret = r.FuncVariadic
				break
			}
			ret = *r.FuncVariadic
		case r.FuncVariadicReturn != nil:
			if getPtr {
				ret = r.FuncVariadicReturn
				break
			}
			ret = *r.FuncVariadicReturn
		case r.Interface != nil:
			ret = r.Interface
		case r.Map != nil:
			if getPtr {
				ret = r.Map
				break
			}
			ret = *r.Map
		case r.Ptr != nil:
			if getPtr {
				ret = r.Ptr
				break
			}
			ret = *r.Ptr
		case r.Slice != nil:
			if getPtr {
				ret = r.Slice
				break
			}
			ret = *r.Slice
		case r.TypeString != nil:
			if getPtr {
				ret = r.TypeString
				break
			}
			ret = *r.TypeString
		case r.StringArray != nil:
			if getPtr {
				ret = r.StringArray
				break
			}
			ret = *r.StringArray
		case r.Byte != nil:
			if getPtr {
				ret = r.Byte
				break
			}
			ret = *r.Byte
		case r.ByteArray != nil:
			if getPtr {
				ret = r.ByteArray
				break
			}
			ret = *r.ByteArray
		case r.Struct != nil:
			if getPtr {
				ret = r.Struct
				break
			}
			ret = *r.Struct
	}
	return ret
}


func (r TypeResponse) Pointer() interface{} {
	return r.getter(true)
}


func (r TypeResponse) Value() interface{} {
	return r.getter(false)
}


func (r TypeResponse) String() string {
	var ret string
	switch {
		case r.Bool != nil:
			ret = fmt.Sprintf("%v", *r.Bool)
		case r.Int != nil:
			ret = fmt.Sprintf("%v", *r.Int)
		case r.Int8 != nil:
			ret = fmt.Sprintf("%v", *r.Int8)
		case r.Int16 != nil:
			ret = fmt.Sprintf("%v", *r.Int16)
		case r.Int32 != nil:
			ret = fmt.Sprintf("%v", *r.Int32)
		case r.Int64 != nil:
			ret = fmt.Sprintf("%v", *r.Int64)
		case r.Uint != nil:
			ret = fmt.Sprintf("%v", *r.Uint)
		case r.Uint8 != nil:
			ret = fmt.Sprintf("%v", *r.Uint8)
		case r.Uint16 != nil:
			ret = fmt.Sprintf("%v", *r.Uint16)
		case r.Uint32 != nil:
			ret = fmt.Sprintf("%v", *r.Uint32)
		case r.Uint64 != nil:
			ret = fmt.Sprintf("%v", *r.Uint64)
		case r.Uintptr != nil:
			ret = fmt.Sprintf("%v", *r.Uintptr)
		case r.Float32 != nil:
			ret = fmt.Sprintf("%v", *r.Float32)
		case r.Float64 != nil:
			ret = fmt.Sprintf("%v", *r.Float64)
		case r.Complex64 != nil:
			ret = fmt.Sprintf("%v", *r.Complex64)
		case r.Complex128 != nil:
			ret = fmt.Sprintf("%v", *r.Complex128)
		case r.InterfaceArray != nil:
			ret = fmt.Sprintf("%v", *r.InterfaceArray)
		case r.Func != nil:
			ret = fmt.Sprintf("%v", r.Func)
		case r.FuncReturn != nil:
			ret = fmt.Sprintf("%v", r.FuncReturn)
		case r.FuncVariadic != nil:
			ret = fmt.Sprintf("%v", r.FuncVariadic)
		case r.FuncVariadicReturn != nil:
			ret = fmt.Sprintf("%v", r.FuncVariadicReturn)
		case r.Interface != nil:
			ret = fmt.Sprintf("%v", r.Interface)
		case r.Map != nil:
			ret = fmt.Sprintf("%v", *r.Map)
		case r.Ptr != nil:
			ret = fmt.Sprintf("%v", *r.Ptr)
		case r.Slice != nil:
			ret = fmt.Sprintf("%v", *r.Slice)
		case r.TypeString != nil:
			ret = fmt.Sprintf("%v", *r.TypeString)
		case r.StringArray != nil:
			ret = fmt.Sprintf("%v", *r.StringArray)
		case r.Byte != nil:
			ret = fmt.Sprintf("%v", *r.Byte)
		case r.ByteArray != nil:
			ret = fmt.Sprintf("%v", *r.ByteArray)
		case r.Struct != nil:
			ret = fmt.Sprintf("%v", *r.Struct)
	}
	return ret
}


//func (r *TypeResponse) Set(ref interface{}) bool {
//	var ok bool
//
//	for range onlyOnce {
//		v := reflect.ValueOf(ref)
//
//		if v.Kind() == reflect.Func {
//			r.TypeReflect = TypeReflect{}
//			r.ofType = reflect.TypeOf(ref)
//			r.data = ref
//
//			//fmt.Printf(">>%s<<\t>>%s<<\t>>%s<<\n",
//			//	r.ofType.String(),
//			//	r.ofType.Name(),
//			//	r.ofType.Kind(),
//			//	)
//			ok = false
//			switch r.ofType.String() {
//			case TypeFunc:
//				f := ref.(func())
//				r.Func = &f
//				ok = true
//			case TypeFuncReturn:
//				f := ref.(func() *TypeResponse)
//				r.FuncReturn = &f
//				ok = true
//			case TypeFuncVariadic:
//				f := ref.(func(args ...interface{}))
//				r.FuncVariadic = &f
//				ok = true
//			case TypeFuncVariadicReturn:
//				f := ref.(func(args ...interface{}) *TypeResponse)
//				r.FuncVariadicReturn = &f
//				ok = true
//			}
//			break
//		}
//
//		if v.Kind() != reflect.Ptr {
//			PrintflnError("SetResponse requires a pointer type, but is a '%s' kind of type '%s'", v.Kind().String(), v.Type().String())
//			//PrintflnError("Example: State.SetResponse(&xyzzy)")
//			//if !r.Set(&ref) {
//			//	panic("ABORTING")
//			//}
//			foo := v.Convert(v.Type())
//			PrintflnError("foo is a '%s' kind of type '%s'", foo.Kind().String(), foo.String())
//
//			//foo2 := v.Addr()
//			//PrintflnError("foo is a '%s' kind of type '%s'", foo2.Kind().String(), foo2.String())
//
//			//refPtr := &ref
//			//v = reflect.ValueOf(refPtr)
//			//PrintflnError("refPtr is a '%s' kind of type '%s'", v.Kind().String(), v.String())
//
//			r.TypeReflect = TypeReflect{}
//			r.ofType = reflect.TypeOf(ref)
//			r.data = ref
//			switch r.ofType.String() {
//			case TypeBool:
//				r.Bool = r.data.(*bool)
//
//			case TypeInt:
//				r.Int = r.data.(*int)
//			case TypeInt8:
//				r.Int8 = r.data.(*int8)
//			case TypeInt16:
//				r.Int16 = r.data.(*int16)
//			case TypeInt32:
//				r.Int32 = r.data.(*int32)
//			case TypeInt64:
//				r.Int64 = r.data.(*int64)
//
//			case TypeUint:
//				r.Uint = r.data.(*uint)
//			//case TypeUint8:
//			//	r.Uint8 = r.data.(*uint8)
//			case TypeUint16:
//				r.Uint16 = r.data.(*uint16)
//			case TypeUint32:
//				r.Uint32 = r.data.(*uint32)
//			case TypeUint64:
//				r.Uint64 = r.data.(*uint64)
//
//			case TypeUintptr:
//				r.Uintptr = r.data.(*uintptr)
//
//			case TypeFloat32:
//				r.Float32 = r.data.(*float32)
//			case TypeFloat64:
//				r.Float64 = r.data.(*float64)
//
//			case TypeComplex64:
//				r.Complex64 = r.data.(*complex64)
//			case TypeComplex128:
//				r.Complex128 = r.data.(*complex128)
//
//			case TypeInterfaceArray:
//				r.InterfaceArray = r.data.(*[]interface{})
//
//			case TypeFunc:
//				r.Func = r.data.(*func())
//
//			case TypeMap:
//				r.Map = r.data.(*map[interface{}]interface{})
//
//			case TypeInterface:
//				r.Interface = r.data.(*interface{})
//
//			case TypeString:
//				r.TypeString = r.data.(*string)
//			case TypeStringArray:
//				t := r.data.([]string)
//				r.StringArray = &t
//
//			case TypeByte:
//				r.Byte = r.data.(*byte)
//			case TypeByteArray:
//				r.ByteArray = r.data.(*[]byte)
//
//			case TypeStruct:
//				r.Struct = r.data.(*struct{})
//			}
//
//		}
//
//		//switch v.Kind() {
//		//	case reflect.Ptr:
//		//		// OK.
//		//	case reflect.Func:
//		//		f := ref.(func())
//		//		r.Func = &f
//		//		ok = true
//		//		break
//		//		//// fptr is a pointer to a function.
//		//		//// Obtain the function value itself (likely nil) as a reflect.Value
//		//		//// so that we can query its type and then set the value.
//		//		//v = reflect.ValueOf(ref).Elem()
//		//		//swap := func(in []reflect.Value) []reflect.Value {
//		//		//	return []reflect.Value{in[1], in[0]}
//		//		//}
//		//		//// Make a function of the right type.
//		//		//v = reflect.MakeFunc(v.Type(), swap)
//		//		//// OK.
//		//	default:
//		//		PrintflnError("SetResponse requires a pointer type, but is a '%s' kind of type '%s'", v.Kind().String(), v.String())
//		//		PrintflnError("Example: State.SetResponse(&xyzzy)")
//		//		panic("ABORTING")
//		//}
//
//		s := v.Elem()
//		r.ofType = s.Type()
//		r.data = v.Interface()
//
//		// @TODO MICKMAKE
//		//fmt.Printf("String: %s\t", r.ofType.String())
//		//fmt.Printf("Name: %s\t", r.ofType.Name())
//		//fmt.Printf("Kind: %s\n", r.ofType.Kind())
//		// @TODO MICKMAKE
//
//		// If we have a pointer, then call again with the value of that pointer.
//		if r.ofType.Kind().String() == "ptr" {
//			ok = r.Set(s.Addr().Elem().Interface())
//			break
//		}
//
//		//PrintflnError("Checking type '%s'", r.ofType.String())
//		r.TypeReflect = TypeReflect{}
//		switch r.ofType.String() {
//		case TypeBool:
//			r.Bool = r.data.(*bool)
//
//		case TypeInt:
//			r.Int = r.data.(*int)
//		case TypeInt8:
//			r.Int8 = r.data.(*int8)
//		case TypeInt16:
//			r.Int16 = r.data.(*int16)
//		case TypeInt32:
//			r.Int32 = r.data.(*int32)
//		case TypeInt64:
//			r.Int64 = r.data.(*int64)
//
//		case TypeUint:
//			r.Uint = r.data.(*uint)
//		//case TypeUint8:
//		//	r.Uint8 = r.data.(*uint8)
//		case TypeUint16:
//			r.Uint16 = r.data.(*uint16)
//		case TypeUint32:
//			r.Uint32 = r.data.(*uint32)
//		case TypeUint64:
//			r.Uint64 = r.data.(*uint64)
//
//		case TypeUintptr:
//			r.Uintptr = r.data.(*uintptr)
//
//		case TypeFloat32:
//			r.Float32 = r.data.(*float32)
//		case TypeFloat64:
//			r.Float64 = r.data.(*float64)
//
//		case TypeComplex64:
//			r.Complex64 = r.data.(*complex64)
//		case TypeComplex128:
//			r.Complex128 = r.data.(*complex128)
//
//		case TypeInterfaceArray:
//			r.InterfaceArray = r.data.(*[]interface{})
//
//		case TypeFunc:
//			r.Func = r.data.(*func())
//
//		case TypeMap:
//			r.Map = r.data.(*map[interface{}]interface{})
//
//		case TypeInterface:
//			r.Interface = r.data.(*interface{})
//
//		case TypeString:
//			r.TypeString = r.data.(*string)
//		case TypeStringArray:
//			r.StringArray = r.data.(*[]string)
//
//		case TypeByte:
//			r.Byte = r.data.(*byte)
//		case TypeByteArray:
//			r.ByteArray = r.data.(*[]byte)
//
//		case TypeStruct:
//			r.Struct = r.data.(*struct{})
//		}
//
//		//switch ref.(type) {
//		//	case bool:
//		//		*r.Bool = ref.(bool)
//		//
//		//	case int:
//		//		*r.Int = ref.(int)
//		//	case int8:
//		//		*r.Int8 = ref.(int8)
//		//	case int16:
//		//		*r.Int16 = ref.(int16)
//		//	case int32:
//		//		*r.Int32 = ref.(int32)
//		//	case int64:
//		//		*r.Int64 = ref.(int64)
//		//
//		//	case uint:
//		//		*r.Uint = ref.(uint)
//		//	//case uint8:
//		//	//	*r.Uint8 = ref.(uint8)
//		//	case uint16:
//		//		*r.Uint16 = ref.(uint16)
//		//	case uint32:
//		//		*r.Uint32 = ref.(uint32)
//		//	case uint64:
//		//		*r.Uint64 = ref.(uint64)
//		//
//		//	case uintptr:
//		//		*r.Uintptr = ref.(uintptr)
//		//
//		//	case float32:
//		//		*r.Float32 = ref.(float32)
//		//	case float64:
//		//		*r.Float64 = ref.(float64)
//		//
//		//	case complex64:
//		//		*r.Complex64 = ref.(complex64)
//		//	case complex128:
//		//		*r.Complex128 = ref.(complex128)
//		//
//		//	case []interface{}:
//		//		*r.Array = ref.([]interface{})
//		//
//		//	case func():
//		//		*r.Func = ref.(func())
//		//
//		//	case map[interface{}]interface{}:
//		//		*r.Map = ref.(map[interface{}]interface{})
//		//
//		//	case interface{}:
//		//		*r.Ptr = ref.(interface{})
//		//
//		//	case string:
//		//		*r.String = ref.(string)
//		//	case []string:
//		//		*r.StringArray = ref.([]string)
//		//
//		//	case byte:
//		//		*r.Byte = ref.(byte)
//		//	case []byte:
//		//		*r.ByteArray = ref.([]byte)
//		//
//		//	case struct{}:
//		//		*r.Struct = ref.(struct{})
//		//}
//
//		ok = true
//	}
//
//	return ok
//}


func (r *TypeResponse) GetType() reflect.Type {
	return r.ofType
}

//func (r *TypeResponse) GetData() interface{} {
//	return r.data
//}


func (r *TypeResponse) IsOfType(t string) bool {
	var ok bool

	//fmt.Printf("%s - %s\n", r.ofType.String(), r.ofType.Name())
	for range onlyOnce {
		if r.ofType == nil {
			ok = false
			break
		}

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
