package gtks

import sugar "github.com/jopbrown/gtk-sugar"

type Adjustment struct {
	sugar.CandyWrapper
}

func NewAdjustment(candy sugar.Candy, id string) *Adjustment {
	v := Adjustment{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_adjustment_new, NONE, WIDGET, 6, DOUBLE, DOUBLE, DOUBLE, DOUBLE, DOUBLE, DOUBLE
// FUNCTION_NAME = gtk_adjustment_get_value, NONE, DOUBLE, 1, WIDGET
// FUNCTION_NAME = gtk_adjustment_set_value, NONE, NONE, 2, WIDGET, DOUBLE
// FUNCTION_NAME = gtk_adjustment_clamp_page, NONE, NONE, 3, WIDGET, DOUBLE, DOUBLE
