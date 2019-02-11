package sugar

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	STRING_BOOL_TRUE  = "1"
	STRING_BOOL_FALSE = "0"
)

// Response of gtk-server
type Response string

func (resp Response) String() string {
	return string(resp)
}

func (resp Response) Unquote() (string, error) {
	return strconv.Unquote(resp.String())
}

func (resp Response) MustUnquote() string {
	s, err := resp.Unquote()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return s
}

func (resp Response) Bool() (bool, error) {
	if resp == STRING_BOOL_TRUE {
		return true, nil
	}

	if resp == STRING_BOOL_FALSE {
		return false, nil
	}

	return false, errors.Errorf("invalid BOOL value: %v", resp)
}

func (resp Response) MustBool() bool {
	bo, err := resp.Bool()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return bo
}

func (resp Response) Int64() (int64, error) {
	return strconv.ParseInt(resp.String(), 0, 64)
}

func (resp Response) MustInt64() int64 {
	i, err := resp.Int64()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return i
}

func (resp Response) Int32() (int32, error) {
	i, err := strconv.ParseInt(resp.String(), 0, 32)
	return int32(i), err
}

func (resp Response) MustInt32() int32 {
	i, err := resp.Int32()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return i
}

func (resp Response) Int() (int, error) {
	i, err := strconv.ParseInt(resp.String(), 0, 64)
	return int(i), err
}

func (resp Response) MustInt() int {
	i, err := resp.Int()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return i
}

func (resp Response) Float32() (float32, error) {
	v, err := strconv.ParseFloat(resp.String(), 64)
	return float32(v), err
}

func (resp Response) MustFloat32() float32 {
	i, err := resp.Float32()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return i
}

func (resp Response) Float64() (float64, error) {
	v, err := strconv.ParseFloat(resp.String(), 64)
	return v, err
}

func (resp Response) MustFloat64() float64 {
	i, err := resp.Float64()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return i
}

type RespFields []Response

func (resp Response) Fields() RespFields {
	fields := strings.Fields(resp.String())
	resps := make([]Response, len(fields))
	for i, field := range fields {
		resps[i] = Response(field)
	}

	return RespFields(resps)
}

func (fields RespFields) Unmarshal(packer *Base64Packer) error {
	return packer.Unmarshal(fields)
}

func (fields RespFields) MustUnmarshal(packer *Base64Packer) {
	err := fields.Unmarshal(packer)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
}

func (fields RespFields) String() string {
	var sb strings.Builder
	for _, field := range fields[:len(fields)-1] {
		sb.WriteString(field.String())
		sb.WriteByte(' ')
	}

	sb.WriteString(fields[len(fields)-1].String())

	return sb.String()
}
