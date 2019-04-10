package gtks

import sugar "github.com/jopbrown/gtk-sugar"

type Separator struct {
	Widget
}

type VSeparator struct {
	Separator
}

func NewVSeparator(candy sugar.Candy, id string) *VSeparator {
	obj := VSeparator{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_vseparator_new, NONE, WIDGET, 0
func (gtk *Gtk) NewVSeparator() *VSeparator {
	id := gtk.Guify("gtk_vseparator_new").String()
	return NewVSeparator(gtk, id)
}

type HSeparator struct {
	Separator
}

func NewHSeparator(candy sugar.Candy, id string) *HSeparator {
	obj := HSeparator{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_hseparator_new, NONE, WIDGET, 0
func (gtk *Gtk) NewHSeparator() *HSeparator {
	id := gtk.Guify("gtk_hseparator_new").String()
	return NewHSeparator(gtk, id)
}
