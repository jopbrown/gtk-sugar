package glib

import sugar "github.com/jopbrown/gtk-sugar"

type Object struct {
	sugar.CandyWrapper
}

func NewObjectFromCandy(candy sugar.Candy, id string) *Object {
	obj := Object{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = g_object_get, NONE, NONE, 4, WIDGET, STRING, PTR_STRING, NULL
// FUNCTION_NAME = g_object_set, NONE, NONE, 4, WIDGET, STRING, INT, NULL

// FUNCTION_NAME = g_object_ref_sink, NONE, NONE, 1, WIDGET
func (obj *Object) RefSink() {
	obj.Candy().Guify("g_object_ref_sink", obj)
}

// FUNCTION_NAME = g_object_unref, NONE, NONE, 1, WIDGET
func (obj *Object) Unref() {
	obj.Candy().Guify("g_object_unref", obj)
}

// FUNCTION_NAME = g_signal_stop_emission_by_name, NONE, NONE, 2, WIDGET, STRING
func (obj *Object) StopEmission(name string) {
	obj.Candy().Guify("g_signal_stop_emission_by_name", obj, name)
}
