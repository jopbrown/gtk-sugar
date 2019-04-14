package gtk

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/glibs"
)

type Adjustment struct {
	glibs.Object
}

func NewAdjustmentFromCandy(candy sugar.Candy, id string) *Adjustment {
	v := Adjustment{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_adjustment_new, NONE, WIDGET, 6, DOUBLE, DOUBLE, DOUBLE, DOUBLE, DOUBLE, DOUBLE
func NewAdjustment(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) *Adjustment {
	id := Candy().Guify("gtk_adjustment_new", value, lower, upper, stepIncrement, pageIncrement, pageSize).String()
	return NewAdjustmentFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_adjustment_get_value, NONE, DOUBLE, 1, WIDGET
func (obj *Adjustment) GetValue() float64 {
	return obj.Candy().Guify("gtk_adjustment_get_value", obj).MustFloat64()
}

// FUNCTION_NAME = gtk_adjustment_set_value, NONE, NONE, 2, WIDGET, DOUBLE
func (obj *Adjustment) SetValue(value float64) {
	obj.Candy().Guify("gtk_adjustment_set_value", obj, value)
}

// FUNCTION_NAME = gtk_adjustment_clamp_page, NONE, NONE, 3, WIDGET, DOUBLE, DOUBLE
func (obj *Adjustment) ClampPage(lower, upper float64) {
	obj.Candy().Guify("gtk_adjustment_clamp_page", obj, lower, upper)
}
