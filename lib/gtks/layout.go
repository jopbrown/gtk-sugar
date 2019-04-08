package gtks

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
func (gtk *Gtk) NewLayout(hadjustment, vadjustment *Adjustment) *Layout {
	id := gtk.Guify("gtk_layout_new", hadjustment, vadjustment).String()
	return NewLayout(gtk, id)
}
