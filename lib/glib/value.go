package glib

import (
	"fmt"

	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/libc"
)

type Value struct {
	sugar.CandyWrapper
}

func NewValue(candy sugar.Candy, id string) *Value {
	obj := Value{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

func GValue(govalue interface{}) *Value {
	switch v := govalue.(type) {
	case nil:
		return nil
	case bool:
		val := ValueInit(TYPE_BOOLEAN)
		val.SetBoolean(v)
		return val

	case byte:
		val := ValueInit(TYPE_CHAR)
		val.SetChar(v)
		return val

	case int64:
		val := ValueInit(TYPE_INT64)
		val.SetInt64(v)
		return val

	case int:
		val := ValueInit(TYPE_INT)
		val.SetInt(v)
		return val

	case uint64:
		val := ValueInit(TYPE_UINT64)
		val.SetUint64(v)
		return val

	case uint:
		val := ValueInit(TYPE_UINT)
		val.SetUint(v)
		return val

	case float32:
		val := ValueInit(TYPE_FLOAT)
		val.SetFloat(v)
		return val

	case float64:
		val := ValueInit(TYPE_DOUBLE)
		val.SetDouble(v)
		return val

	case string:
		val := ValueInit(TYPE_STRING)
		val.SetString(v)
		return val

	case IObject:
		val := ValueInit(TYPE_OBJECT)
		val.SetObject(v)

	case sugar.CandyWrapper:
		val := ValueInit(TYPE_POINTER)
		val.SetPointer(v)
	default:
		panic(fmt.Sprintf("Type %T not implemented", v))
	}

	return nil
}

// FUNCTION_NAME = g_value_init, NONE, WIDGET, 2, WIDGET, INT
func ValueInit(t Type) *Value {
	in := libc.Malloc(24) // 8bytes x 3
	libc.Memset(in, 0, 24)
	out := Candy().Guify("g_value_init", in, t).String()
	obj := NewValue(Candy(), out)
	return obj
}

// FUNCTION_NAME = g_value_unset, NONE, NONE, 1, WIDGET
func (obj *Value) Unset() {
	Candy().Guify("g_value_unset", obj)
}

func (obj *Value) Type() Type {
	return Type(obj.Candy().ServerUnpackFromPointer("%l", obj.ID())[0].MustUint())
}

// FUNCTION_NAME = g_value_get_boolean, NONE, BOOL, 1, WIDGET
func (obj *Value) GetBoolean() bool {
	return obj.Candy().Guify("g_value_get_boolean", obj).MustBool()
}

// FUNCTION_NAME = g_value_get_char, NONE, INT, 1, WIDGET
func (obj *Value) GetChar() byte {
	return byte(obj.Candy().Guify("g_value_get_char", obj).MustInt())
}

// FUNCTION_NAME = g_value_get_int, NONE, INT, 1, WIDGET
func (obj *Value) GetInt() int {
	return obj.Candy().Guify("g_value_get_int", obj).MustInt()
}

// FUNCTION_NAME = g_value_get_uint, NONE, INT, 1, WIDGET
func (obj *Value) GetUint() uint {
	return obj.Candy().Guify("g_value_get_uint", obj).MustUint()
}

// FUNCTION_NAME = g_value_get_int64, NONE, LONG, 1, WIDGET
func (obj *Value) GetInt64() int64 {
	return obj.Candy().Guify("g_value_get_int64", obj).MustInt64()
}

// FUNCTION_NAME = g_value_get_long, NONE, LONG, 1, WIDGET
func (obj *Value) GetLong() int64 {
	return obj.Candy().Guify("g_value_get_long", obj).MustInt64()
}

// FUNCTION_NAME = g_value_get_uint64, NONE, LONG, 1, WIDGET
func (obj *Value) GetUint64() uint64 {
	return obj.Candy().Guify("g_value_get_uint64", obj).MustUint64()
}

// FUNCTION_NAME = g_value_get_ulong, NONE, LONG, 1, WIDGET
func (obj *Value) GetUlong() uint64 {
	return obj.Candy().Guify("g_value_get_ulong", obj).MustUint64()
}

// FUNCTION_NAME = g_value_get_float, NONE, FLOAT, 1, WIDGET
func (obj *Value) GetFloat() float32 {
	return obj.Candy().Guify("g_value_get_float", obj).MustFloat32()
}

// FUNCTION_NAME = g_value_get_double, NONE, DOUBLE, 1, WIDGET
func (obj *Value) GetDouble() float64 {
	return obj.Candy().Guify("g_value_get_double", obj).MustFloat64()
}

// FUNCTION_NAME = g_value_get_string, NONE, STRING, 1, WIDGET
func (obj *Value) GetString() string {
	return obj.Candy().Guify("g_value_get_string", obj).String()
}

// FUNCTION_NAME = g_value_get_object, NONE, WIDGET, 1, WIDGET
func (obj *Value) GetObject() string {
	return obj.Candy().Guify("g_value_get_object", obj).String()
}

// FUNCTION_NAME = g_value_get_pointer, NONE, POINTER, 1, WIDGET
func (obj *Value) GetPointer() string {
	return obj.Candy().Guify("g_value_get_pointer", obj).String()
}

// FUNCTION_NAME = g_value_set_boolean, NONE, NONE, 2, WIDGET, BOOL
func (obj *Value) SetBoolean(v bool) {
	obj.Candy().Guify("g_value_set_boolean", obj, v)
}

// FUNCTION_NAME = g_value_set_char, NONE, NONE, 2, WIDGET, INT
func (obj *Value) SetChar(v byte) {
	obj.Candy().Guify("g_value_set_char", obj, int(v))
}

// FUNCTION_NAME = g_value_set_int, NONE, NONE, 2, WIDGET, INT
func (obj *Value) SetInt(v int) {
	obj.Candy().Guify("g_value_set_int", obj, v)
}

// FUNCTION_NAME = g_value_set_uint, NONE, NONE, 2, WIDGET, INT
func (obj *Value) SetUint(v uint) {
	obj.Candy().Guify("g_value_set_uint", obj, v)
}

// FUNCTION_NAME = g_value_set_int64, NONE, NONE, 2, WIDGET, LONG
func (obj *Value) SetInt64(v int64) {
	obj.Candy().Guify("g_value_set_int64", obj, v)
}

// FUNCTION_NAME = g_value_set_long, NONE, NONE, 2, WIDGET, LONG
func (obj *Value) SetLong(v int64) {
	obj.Candy().Guify("g_value_set_long", obj, v)
}

// FUNCTION_NAME = g_value_set_uint64, NONE, NONE, 2, WIDGET, LONG
func (obj *Value) SetUint64(v uint64) {
	obj.Candy().Guify("g_value_set_uint64", obj, v)
}

// FUNCTION_NAME = g_value_set_ulong, NONE, NONE, 2, WIDGET, LONG
func (obj *Value) SetUlong(v uint64) {
	obj.Candy().Guify("g_value_set_ulong", obj, v)
}

// FUNCTION_NAME = g_value_set_float, NONE, NONE, 2, WIDGET, FLOAT
func (obj *Value) SetFloat(v float32) {
	obj.Candy().Guify("g_value_set_float", obj, v)
}

// FUNCTION_NAME = g_value_set_double, NONE, NONE, 2, WIDGET, DOUBLE
func (obj *Value) SetDouble(v float64) {
	obj.Candy().Guify("g_value_set_double", obj, v)
}

// FUNCTION_NAME = g_value_set_string, NONE, NONE, 2, WIDGET, STRING
func (obj *Value) SetString(v string) {
	obj.Candy().Guify("g_value_set_string", obj, v)
}

// FUNCTION_NAME = g_value_set_object, NONE, NONE, 2, WIDGET, WIDGET
func (obj *Value) SetObject(v IObject) {
	obj.Candy().Guify("g_value_set_object", obj, v)
}

// FUNCTION_NAME = g_value_set_pointer, NONE, NONE, 2, WIDGET, POINTER
func (obj *Value) SetPointer(v sugar.CandyWrapper) {
	obj.Candy().Guify("g_value_set_pointer", obj, v)
}
