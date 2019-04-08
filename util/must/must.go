package must

import (
	"fmt"
)

func String(v string, err error) string {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return v
}

func Int(v int, err error) int {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return v
}

func Int32(v int32, err error) int32 {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return v
}

func Int64(v int64, err error) int64 {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return v
}

func Uint(v uint, err error) uint {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return v
}

func Uint32(v uint32, err error) uint32 {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return v
}

func Uint64(v uint64, err error) uint64 {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return v
}

func Float32(v float32, err error) float32 {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return v
}

func Float64(v float64, err error) float64 {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return v
}

func Bool(v bool, err error) bool {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return v
}

func StringSlice(v []string, err error) []string {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return v
}

func Interface(v interface{}, err error) interface{} {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return v
}

func InterfaceSlice(v []interface{}, err error) []interface{} {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return v
}

func Bytes(v []byte, err error) []byte {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return v
}
