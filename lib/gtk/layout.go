package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Layout struct {
	Container
}

func NewLayout(candy sugar.Candy, id string) *Layout {
	obj := Layout{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_layout_new, NONE, WIDGET, 2, WIDGET, WIDGET
func LayoutNew(hadjustment, vadjustment *Adjustment) *Layout {
	id := Candy().Guify("gtk_layout_new", hadjustment, vadjustment).String()
	return NewLayout(Candy(), id)
}
