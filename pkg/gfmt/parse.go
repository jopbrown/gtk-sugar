package gfmt

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

const (
	STRING_BOOL_TRUE  = "1"
	STRING_BOOL_FALSE = "0"
)

func Bool(s string) (bool, error) {
	if s == STRING_BOOL_TRUE {
		return true, nil
	}

	if s == STRING_BOOL_FALSE {
		return false, nil
	}

	return false, errors.Errorf("invalid BOOL value: %v", s)
}

func MustBool(s string) bool {
	bo, err := Bool(s)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return bo
}

func Unquote(s string) (string, error) {
	return strconv.Unquote(s)
}

func MustUnquote(s string) string {
	s, err := Unquote(s)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return s
}

func Int64(s string) (int64, error) {
	return strconv.ParseInt(s, 0, 64)
}

func MustInt64(s string) int64 {
	i, err := Int64(s)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return i
}

func Int32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 0, 32)
	return int32(i), err
}

func MustInt32(s string) int32 {
	i, err := Int32(s)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return i
}

func Int(s string) (int, error) {
	i, err := strconv.ParseInt(s, 0, 64)
	return int(i), err
}

func MustInt(s string) int {
	i, err := Int(s)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return i
}

func Float32(s string) (float32, error) {
	v, err := strconv.ParseFloat(s, 32)
	return float32(v), err
}

func MustFloat32(s string) float32 {
	i, err := Float32(s)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return i
}

func Float64(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	return v, err
}

func MustFloat64(s string) float64 {
	i, err := Float64(s)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return i
}
