package faker

import (
	"fmt"
	"reflect"
)

type FakerImpl[T any] struct{}

func NewFaker[T any]() *FakerImpl[T] {
	return &FakerImpl[T]{}
}

func (f *FakerImpl[T]) Create() *T {
	res := createValue(reflect.TypeOf(*new(T)))
	resTVal := res.(T)
	return &resTVal
}

func createValue(t reflect.Type) any {
	switch t.Kind() {

	case reflect.Bool:
		return createBool(t)

	case reflect.Int:
		return createInt(t)
	case reflect.Int8:
		return createInt8(t)
	case reflect.Int16:
		return createInt16(t)
	case reflect.Int32:
		return createInt32(t)
	case reflect.Int64:
		return createInt64(t)

	case reflect.Uint:
		return createUint(t)
	case reflect.Uint8:
		return createUint8(t)
	case reflect.Uint16:
		return createUint16(t)
	case reflect.Uint32:
		return createUint32(t)
	case reflect.Uint64:
		return createUint64(t)

	case reflect.Float32:
		return createFloat32(t)
	case reflect.Float64:
		return createFloat64(t)

	case reflect.Complex64:
		return createComplex64(t)
	case reflect.Complex128:
		return createComplex128(t)

	// TODO: create strings
	// TODO: create structs
	// TODO: create pointers
	// TODO: create arrays
	// TODO: create slices
	// TODO: create maps
	// TODO: create chans
	// TODO: create funcs

	default:
		panic(fmt.Sprintf("INVALID TYPE: %+v", t))
	}
}


