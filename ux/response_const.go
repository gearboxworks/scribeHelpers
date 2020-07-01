package ux


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

const (
	TypeInvalid			= "bool"

	TypeBool			= "bool"

	TypeInt				= "int"
	TypeInt8			= "int8"
	TypeInt16			= "int16"
	TypeInt32			= "int32"
	TypeInt64			= "int64"

	TypeUint			= "uint"
	TypeUint8			= "uint8"
	TypeUint16			= "uint16"
	TypeUint32			= "uint32"
	TypeUint64			= "uint64"

	TypeUintptr			= "uintptr"

	TypeFloat32			= "float32"
	TypeFloat64			= "float64"

	TypeComplex64		= "complex64"
	TypeComplex128		= "complex128"

	TypeInterfaceArray	= "[]interface {}"

	TypeFunc				= "func()"
	TypeFuncReturn			= "func() ux.TypeResponse"
	TypeFuncVariadic		= "func(...interface {})"
	TypeFuncVariadicReturn	= "func(...interface {}) ux.TypeResponse"

	TypeInterface		= "interface {}"

	TypeMap				= "map[interface {}]interface {}"

	TypePtr				= "interface {}"

	TypeSlice			= "[]interface {}"

	TypeString			= "string"
	TypeStringArray		= "[]string"

	TypeByte			= "byte"
	TypeByteArray		= "[]byte"

	TypeStruct			= "struct {}"
)
