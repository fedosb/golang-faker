package faker

import (
	"math/rand"
	"reflect"
)

func createBool(_ reflect.Type) any {
	return rand.Intn(2) == 0
}

func createInt(_ reflect.Type) any {
	return rand.Int()
}

func createInt8(_ reflect.Type) any {
	return int8(rand.Int())
}

func createInt16(_ reflect.Type) any {
	return int16(rand.Int())
}

func createInt32(_ reflect.Type) any {
	return int32(rand.Int())
}

func createInt64(_ reflect.Type) any {
	return int64(rand.Int())
}

func createUint(_ reflect.Type) any {
	return uint(rand.Int())
}

func createUint8(_ reflect.Type) any {
	return uint8(rand.Int())
}

func createUint16(_ reflect.Type) any {
	return uint16(rand.Int())
}

func createUint32(_ reflect.Type) any {
	return uint32(rand.Int())
}

func createUint64(_ reflect.Type) any {
	return uint64(rand.Int())
}

func createFloat32(_ reflect.Type) any {
	return rand.Float32()
}

func createFloat64(_ reflect.Type) any {
	return rand.Float64()
}

func createComplex64(_ reflect.Type) any {
	return complex(rand.Float32(), rand.Float32())
}

func createComplex128(_ reflect.Type) any {
	return complex(rand.Float64(), rand.Float64())
}

func createString(_ reflect.Type) any {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, rand.Int()%len(charset))
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
