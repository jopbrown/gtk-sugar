package gtks

import (
	sugar "github.com/jopbrown/gtk-sugar"
)

type Gtk struct {
	sugar.Candy
}

func NewGtk(candy sugar.Candy) *Gtk {
	return &Gtk{Candy: candy}
}

// FUNCTION_NAME = gtk_init, NONE, NONE, 2, NULL, NULL
func (gtk *Gtk) Init() {
	gtk.Guify("gtk_init", nil, nil)
}
