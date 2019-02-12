package gtks

import sugar "github.com/jopbrown/gtk-sugar"

type Box struct {
	Container
}

func NewBox(candy sugar.Candy, id string) *Box {
	v := Box{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_box_pack_start, NONE, NONE, 5, WIDGET, WIDGET, BOOL, BOOL, INT
func (box *Box) PackStart(child IWidget, expand, fill bool, padding int) {
	box.Candy().Guify("gtk_box_pack_start", box, child, expand, fill, padding)
}

// FUNCTION_NAME = gtk_box_pack_end, NONE, NONE, 5, WIDGET, WIDGET, BOOL, BOOL, INT
func (box *Box) PackEnd(child IWidget, expand, fill bool, padding int) {
	box.Candy().Guify("gtk_box_pack_end", box, child, expand, fill, padding)
}

// FUNCTION_NAME = gtk_box_set_spacing, NONE, NONE, 2, WIDGET, INT
func (box *Box) SetSpacing(spacing int) {
	box.Candy().Guify("gtk_box_set_spacing", box, spacing)
}

// FUNCTION_NAME = gtk_box_get_spacing, NONE, INT, 1, WIDGET
func (box *Box) GetSpacing() int {
	return box.Candy().Guify("gtk_box_get_spacing ", box).MustInt()
}

type Hbox struct {
	Box
}

func NewHbox(candy sugar.Candy, id string) *Hbox {
	v := Hbox{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_hbox_new, NONE, WIDGET, 2, BOOL, INT
func (gtk *Gtk) NewHbox(homogeneous bool, spacing int) *Hbox {
	id := gtk.Guify("gtk_hbox_new", homogeneous, spacing).String()
	return NewHbox(gtk, id)
}

type Vbox struct {
	Box
}

func NewVbox(candy sugar.Candy, id string) *Vbox {
	v := Vbox{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_vbox_new, NONE, WIDGET, 2, BOOL, INT
func (gtk *Gtk) NewVbox(homogeneous bool, spacing int) *Vbox {
	id := gtk.Guify("gtk_vbox_new", homogeneous, spacing).String()
	return NewVbox(gtk, id)
}
