package glib

import (
	sugar "github.com/jopbrown/gtk-sugar"
)

type Value struct {
	sugar.CandyWrapper
}

func NewValue(candy sugar.Candy, id string) *Value {
	obj := Value{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = g_value_init, NONE, NONE, 2, WIDGET, INT
func ValueInit(t Type) *Value {
	id := Candy().ServerOpaque()
	Candy().Guify("g_value_init", id, t)
	obj := NewValue(Candy(), id)
	return obj
}

// FUNCTION_NAME = g_value_unset, NONE, NONE, 1, WIDGET
func (obj *Value) Unset() {
	Candy().Guify("g_value_unset", obj)
}

// FUNCTION_NAME = G_VALUE_TYPE, NONE, INT, 1, WIDGET
func (obj *Value) Type() Type {
	return Type(obj.Candy().Guify("G_VALUE_TYPE", obj).MustUint())
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

// FUNCTION_NAME = g_value_get_long, NONE, LONG, 1, WIDGET
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
