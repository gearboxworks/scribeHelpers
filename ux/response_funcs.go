package ux

import (
	"reflect"
)


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

		switch response[0].Interface().(type) {
			case TypeResponse:
				ret = response[0].Interface().(TypeResponse)
			case *TypeResponse:
				ret = *(response[0].Interface().(*TypeResponse))
		}

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
