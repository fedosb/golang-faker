package faker

import (
	"math/rand"
	"reflect"
)

func createBool(_ reflect.Type) bool {
	return rand.Intn(2) == 0
}

func createInt(_ reflect.Type) int {
	return rand.Int()
}

func createInt8(_ reflect.Type) int8 {
	return int8(rand.Int())
}

func createInt16(_ reflect.Type) int16 {
	return int16(rand.Int())
}

func createInt32(_ reflect.Type) int32 {
	return int32(rand.Int())
}

func createInt64(_ reflect.Type) int64 {
	return int64(rand.Int())
}

func createUint(_ reflect.Type) uint {
	return uint(rand.Int())
}

func createUint8(_ reflect.Type) uint8 {
	return uint8(rand.Int())
}

func createUint16(_ reflect.Type) uint16 {
	return uint16(rand.Int())
}

func createUint32(_ reflect.Type) uint32 {
	return uint32(rand.Int())
}

func createUint64(_ reflect.Type) uint64 {
	return uint64(rand.Int())
}

func createFloat32(_ reflect.Type) float32 {
	return rand.Float32()
}

func createFloat64(_ reflect.Type) float64 {
	return rand.Float64()
}

func createComplex64(_ reflect.Type) complex64 {
	return complex(rand.Float32(), rand.Float32())
}

func createComplex128(_ reflect.Type) complex128 {
	return complex(rand.Float64(), rand.Float64())
}
