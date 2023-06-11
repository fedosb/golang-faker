package faker

import (
	"fmt"
	"reflect"
)

type FakerImpl[T any] struct {
	createFuncMap map[reflect.Kind]func(reflect.Type) any
}

func NewFaker[T any]() *FakerImpl[T] {

	createFuncMap := map[reflect.Kind]func(reflect.Type) any{
		reflect.Bool:       createBool,
		reflect.Int:        createInt,
		reflect.Int8:       createInt8,
		reflect.Int16:      createInt16,
		reflect.Int32:      createInt32,
		reflect.Int64:      createInt64,
		reflect.Uint:       createUint,
		reflect.Uint8:      createUint8,
		reflect.Uint16:     createUint16,
		reflect.Uint32:     createUint32,
		reflect.Uint64:     createUint64,
		reflect.Float32:    createFloat32,
		reflect.Float64:    createFloat64,
		reflect.Complex64:  createComplex64,
		reflect.Complex128: createComplex128,
		reflect.String:     createString,
		reflect.Uintptr:    createUintptr,
		reflect.Func:       createFunc,
		// TODO: create structs
		// TODO: create chans
	}

	createFuncMap[reflect.Array] = func(t reflect.Type) any { return createArray(t, createFuncMap) }
	createFuncMap[reflect.Slice] = func(t reflect.Type) any { return createSlice(t, createFuncMap) }
	createFuncMap[reflect.Map] = func(t reflect.Type) any { return createMap(t, createFuncMap) }
	createFuncMap[reflect.Pointer] = func(t reflect.Type) any { return createPointer(t, createFuncMap) }

	return &FakerImpl[T]{
		createFuncMap: createFuncMap,
	}
}

func (f *FakerImpl[T]) Create() *T {
	res := createValue(reflect.TypeOf(*new(T)), f.createFuncMap)
	resTVal := res.(T)
	return &resTVal
}

func createValue(t reflect.Type, createFuncMap map[reflect.Kind]func(reflect.Type) any) any {
	createFunc, ok := createFuncMap[t.Kind()]
	if !ok {
		panic(fmt.Sprintf("INVALID TYPE: %+v", t))
	}
	return createFunc(t)
}
