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

func createArray(t reflect.Type, m map[reflect.Kind]func(reflect.Type) any) any {
	arrayLen := t.Len()
	array := reflect.New(t).Elem()
	for i := 0; i < arrayLen; i++ {
		array.Index(i).Set(reflect.ValueOf(createValue(t.Elem(), m)))
	}
	return array.Interface()
}

func createSlice(t reflect.Type, m map[reflect.Kind]func(reflect.Type) any) any {
	sliceType := t.Elem()
	sliceLen := rand.Intn(10) + 1
	slice := reflect.MakeSlice(t, sliceLen, sliceLen)
	for i := 0; i < sliceLen; i++ {
		slice.Index(i).Set(reflect.ValueOf(createValue(sliceType, m)))
	}
	return slice.Interface()
}

func createMap(t reflect.Type, m map[reflect.Kind]func(reflect.Type) any) any {
	mapType := t.Elem()
	mapLen := rand.Intn(10) + 1
	newMap := reflect.MakeMapWithSize(t, mapLen)
	for i := 0; i < mapLen; i++ {
		key := reflect.ValueOf(createValue(t.Key(), m))
		value := reflect.ValueOf(createValue(mapType, m))
		newMap.SetMapIndex(key, value)
	}
	return newMap.Interface()
}
