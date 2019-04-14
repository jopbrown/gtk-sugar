package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Range struct {
	Widget
}

func NewRangeFromCandy(candy sugar.Candy, id string) *Range {
	obj := Range{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_range_get_adjustment, NONE, WIDGET, 1, WIDGET
func (obj *Range) GetAdjustment() *Adjustment {
	id := obj.Candy().Guify("gtk_range_get_adjustment", obj).String()
	return NewAdjustmentFromCandy(obj.Candy(), id)
}

// FUNCTION_NAME = gtk_range_get_value, NONE, DOUBLE, 1, WIDGET
func (obj *Range) GetValue() float64 {
	return obj.Candy().Guify("gtk_range_get_value", obj).MustFloat64()
}

// FUNCTION_NAME = gtk_range_set_value, NONE, NONE, 2, WIDGET, DOUBLE
func (obj *Range) SetValue(value float64) {
	obj.Candy().Guify("gtk_range_set_value", obj, value)
}
