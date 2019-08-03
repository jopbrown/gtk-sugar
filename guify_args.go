package sugar

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	STRING_NIL               = "NULL"
	STRING_WHITE_SPACE_CHARS = " \t\n\r"
)

func formatCommand(cmd string, args ...interface{}) string {
	var sb strings.Builder
	sb.WriteString(cmd)
	writeArgs(&sb, args)
	return sb.String()
}

func formatArgs(args ...interface{}) string {
	var sb strings.Builder
	writeArgs(&sb, args)
	return sb.String()
}

func writeArgs(sb *strings.Builder, args []interface{}) {
	for _, arg := range args {
		sb.WriteByte(' ')

		switch v := arg.(type) {
		case nil:
			sb.WriteString(STRING_NIL)
		case string:
			if len(v) == 0 {
				sb.WriteString(STRING_NIL)
			} else {
				if strings.IndexAny(v, STRING_WHITE_SPACE_CHARS) > 0 {
					v = strconv.Quote(v)
				}
				sb.WriteString(v)
			}

		case bool:
			if v {
				sb.WriteString(STRING_BOOL_TRUE)
			} else {
				sb.WriteString(STRING_BOOL_FALSE)
			}
		default:
			sb.WriteString(fmt.Sprintf("%v", v))
		}
	}
}

type Args []interface{}

func (args Args) String() string {
	return formatArgs(args...)
}

type Varargs []interface{}

func (args Varargs) String() string {
	var sb strings.Builder
	for _, arg := range args {
		sb.WriteByte(' ')

		switch v := arg.(type) {
		case nil:
			sb.WriteString(STRING_NIL)
		case string:
			sb.WriteString("s:")
			sb.WriteString(v)
		case int, int8, int16, int32, uint, uint8, uint16, uint32:
			sb.WriteString("i:")
			sb.WriteString(fmt.Sprintf("%v", v))
		case int64, uint64:
			sb.WriteString("l:")
			sb.WriteString(fmt.Sprintf("%v", v))
		case float32:
			sb.WriteString("f:")
			sb.WriteString(fmt.Sprintf("%v", v))
		case float64:
			sb.WriteString("d:")
			sb.WriteString(fmt.Sprintf("%v", v))
		case bool:
			sb.WriteString("b:")
			if v {
				sb.WriteString(STRING_BOOL_TRUE)
			} else {
				sb.WriteString(STRING_BOOL_FALSE)
			}
		case CandyWrapper:
			sb.WriteString("w:")
			sb.WriteString(v.ID())
		default:
			sb.WriteString("p:")
			sb.WriteString(fmt.Sprintf("%v", v))
		}
	}

	return sb.String()
}

type Base64Packer reflect.Value

func NewBase64Packer(s interface{}) *Base64Packer {
	v := reflect.ValueOf(s).Elem()
	if v.Kind() != reflect.Struct {
		panic(errors.Errorf("invalid value to base64 packer: not struct but %v", v.Type()))
	}
	packer := Base64Packer(v)
	return &packer
}

func (packer *Base64Packer) Format() string {
	t := (*reflect.Value)(packer).Type()
	var sb strings.Builder
	n := t.NumField()
	for i := 0; i < n; i++ {
		switch t.Field(i).Type.Kind() {
		case reflect.Int8, reflect.Uint8:
			sb.WriteString("%c")
		case reflect.Int16, reflect.Uint16:
			sb.WriteString("%s")
		case reflect.Int32, reflect.Uint32, reflect.Int64, reflect.Uint64:
			sb.WriteString("%l")
		case reflect.Bool, reflect.Int, reflect.Uint:
			sb.WriteString("%i")
		case reflect.Float32:
			sb.WriteString("%f")
		case reflect.Float64:
			sb.WriteString("%d")
		case reflect.String: // char pointer
			sb.WriteString("%l")
		}
	}

	return sb.String()
}

func (packer *Base64Packer) Args() Args {
	v := (*reflect.Value)(packer)
	n := v.NumField()
	args := make(Args, 0, n)
	for i := 0; i < n; i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.String, reflect.Int64, reflect.Uint64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Float32, reflect.Float64, reflect.Bool:
			args = append(args, field.Interface())
		}
	}

	return args
}

func (packer *Base64Packer) Unmarshal(fields RespFields) error {
	v := (*reflect.Value)(packer)
	n := v.NumField()
	for i := 0; i < n; i++ {
		f := v.Field(i)
		if i >= len(fields) {
			return errors.Errorf("number of RespFields(%v) is less than struct", fields)
		}
		resp := fields[i]
		switch f.Kind() {
		case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			v, err := resp.Int64()
			if nil != err {
				return errors.WithStack(err)
			}
			f.SetInt(v)
		case reflect.Uint64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
			v, err := resp.Uint64()
			if nil != err {
				return errors.WithStack(err)
			}
			f.SetUint(v)
		case reflect.Float32, reflect.Float64:
			v, err := resp.Float64()
			if nil != err {
				return errors.WithStack(err)
			}
			f.SetFloat(v)
		case reflect.Bool:
			v, err := resp.Bool()
			if nil != err {
				return errors.WithStack(err)
			}
			f.SetBool(v)
		}
	}

	return nil
}
