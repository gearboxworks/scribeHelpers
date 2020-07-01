package ux

import (
	"fmt"
	"reflect"
)


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


func (r *TypeResponse) AsFunc() func() {
	if r.Func == nil {
		return nilFunc
	}
	return *r.Func
}
func nilFunc() {
	panic("This is a nil function")
}


func (r *TypeResponse) AsFuncReturn() func() *TypeResponse {
	if r.FuncReturn == nil {
		return nilFuncReturn
	}
	return *(r.FuncReturn)
}
func nilFuncReturn() *TypeResponse {
	panic("This is a nil function")
}


func (r *TypeResponse) AsFuncVariadic() func(args ...interface{}) {
	if r.FuncVariadic == nil {
		return nilFuncVariadic
	}
	return *r.FuncVariadic
}
func nilFuncVariadic(args ...interface{}) {
	panic("This is a nil function")
}


func (r *TypeResponse) AsFuncVariadicReturn() func(args ...interface{}) *TypeResponse {
	if r.FuncVariadicReturn == nil {
		return nilFuncVariadicReturn
	}
	return *(r.FuncVariadicReturn)
}
func nilFuncVariadicReturn(args ...interface{}) *TypeResponse {
	panic("This is a nil function")
}


func (r *TypeResponse) AsFuncCall(args ...interface{}) *TypeResponse {
	var ret TypeResponse

	for range onlyOnce {
		var fn reflect.Value
		var rargs []reflect.Value
		var ok bool

		//fmt.Printf("I AM '%s'\n", r.ofType.String())
		switch r.ofType.String() {
			case TypeFunc:
				if r.Func == nil {
					break
				}
				fn = reflect.ValueOf(*r.Func)
				ok = true

			case TypeFuncReturn:
				if r.FuncReturn == nil {
					break
				}
				fn = reflect.ValueOf(*r.FuncReturn)
				ok = true

			case TypeFuncVariadic:
				if r.FuncVariadic == nil {
					break
				}
				fn = reflect.ValueOf(*r.FuncVariadic)
				for _, a := range args {
					rargs = append(rargs, reflect.ValueOf(a))
				}
				ok = true

			case TypeFuncVariadicReturn:
				if r.FuncVariadicReturn == nil {
					break
				}
				fn = reflect.ValueOf(*r.FuncVariadicReturn)
				for _, a := range args {
					rargs = append(rargs, reflect.ValueOf(a))
				}
				ok = true
		}

		//fmt.Printf("I AM '%s'\n", fn.String())
		if !ok {
			fn = reflect.ValueOf(nilFunc)
		}
		var response []reflect.Value
		response = fn.Call(rargs)
		if len(response) == 0 {
			break
		}

		ret = response[0].Interface().(TypeResponse)

		//v := response[0]
		//fmt.Printf("I AM v.String() == '%s'\nv.Type().Name() == '%s'\nv.Type().String() == '%s'\nv.Type().Kind() == '%s'\nv.Kind() == '%s'\n",
		//	v.String(),
		//	v.Type().Name(),
		//	v.Type().String(),
		//	v.Type().Kind(),
		//	v.Kind(),
		//)
		//if v.Type().String() != "ux.TypeResponse" {
		//	break
		//}
		//
		//r := reflect.TypeOf(v)
		//fmt.Printf(">>%s<<\t>>%s<<\t>>%s<<\n",
		//	r.String(),
		//	r.Name(),
		//	r.Kind(),
		//)
		//
		//ret2 := v.Interface().(TypeResponse)
		//v = reflect.ValueOf(ret2)
		//fmt.Printf("I AM v.String() == '%s'\nv.Type().Name() == '%s'\nv.Type().String() == '%s'\nv.Type().Kind() == '%s'\nv.Kind() == '%s'\n",
		//	v.String(),
		//	v.Type().Name(),
		//	v.Type().String(),
		//	v.Type().Kind(),
		//	v.Kind(),
		//)
	}

	return &ret
}
