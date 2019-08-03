package sugar

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jopbrown/gtk-sugar/util/must"

	"github.com/pkg/errors"
)

const (
	_STRING_BOOL_TRUE  = "1"
	_STRING_BOOL_FALSE = "0"
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
	return must.String(resp.Unquote())
}

func (resp Response) Bool() (bool, error) {
	if resp == _STRING_BOOL_TRUE {
		return true, nil
	}

	if resp == _STRING_BOOL_FALSE {
		return false, nil
	}

	return false, errors.Errorf("invalid BOOL value: %v", resp)
}

func (resp Response) MustBool() bool {
	return must.Bool(resp.Bool())
}

func (resp Response) Int64() (int64, error) {
	return strconv.ParseInt(resp.String(), 0, 64)
}

func (resp Response) MustInt64() int64 {
	return must.Int64(resp.Int64())
}

func (resp Response) Int32() (int32, error) {
	i, err := strconv.ParseInt(resp.String(), 0, 32)
	return int32(i), err
}

func (resp Response) MustInt32() int32 {
	return must.Int32(resp.Int32())
}

func (resp Response) Int() (int, error) {
	i, err := strconv.ParseInt(resp.String(), 0, 64)
	return int(i), err
}

func (resp Response) MustInt() int {
	return must.Int(resp.Int())
}

func (resp Response) Uint64() (uint64, error) {
	return strconv.ParseUint(resp.String(), 0, 64)
}

func (resp Response) MustUint64() uint64 {
	return must.Uint64(resp.Uint64())
}

func (resp Response) Uint32() (uint32, error) {
	i, err := strconv.ParseUint(resp.String(), 0, 32)
	return uint32(i), err
}

func (resp Response) MustUint32() uint32 {
	return must.Uint32(resp.Uint32())
}

func (resp Response) Uint() (uint, error) {
	i, err := strconv.ParseUint(resp.String(), 0, 64)
	return uint(i), err
}

func (resp Response) MustUint() uint {
	return must.Uint(resp.Uint())
}

func (resp Response) Float32() (float32, error) {
	v, err := strconv.ParseFloat(resp.String(), 64)
	return float32(v), err
}

func (resp Response) MustFloat32() float32 {
	return must.Float32(resp.Float32())
}

func (resp Response) Float64() (float64, error) {
	v, err := strconv.ParseFloat(resp.String(), 64)
	return v, err
}

func (resp Response) MustFloat64() float64 {
	return must.Float64(resp.Float64())
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
