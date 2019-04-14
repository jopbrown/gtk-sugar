package gtks

import sugar "github.com/jopbrown/gtk-sugar"

type Misc struct {
	Widget
}

func NewMiscFromCandy(candy sugar.Candy, id string) *Misc {
	v := Misc{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_misc_set_alignment, NONE, NONE, 3, WIDGET, FLOAT, FLOAT
// FUNCTION_NAME = gtk_misc_set_padding, NONE, NONE, 3, WIDGET, INT, INT
// FUNCTION_NAME = gtk_misc_get_alignment, NONE, NONE, 3, WIDGET, PTR_FLOAT, PTR_FLOAT
// FUNCTION_NAME = gtk_misc_get_padding, NONE, NONE, 3, WIDGET, PTR_INT, PTR_INT
