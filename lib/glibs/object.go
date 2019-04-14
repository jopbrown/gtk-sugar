package glibs

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
