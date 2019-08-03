package gtk

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

type HBox struct {
	Box
}

func NewHBox(candy sugar.Candy, id string) *HBox {
	v := HBox{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_hbox_new, NONE, WIDGET, 2, BOOL, INT
func HBoxNew(homogeneous bool, spacing int) *HBox {
	id := Candy().Guify("gtk_hbox_new", homogeneous, spacing).String()
	return NewHBox(Candy(), id)
}

type VBox struct {
	Box
}

func NewVBox(candy sugar.Candy, id string) *VBox {
	v := VBox{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_vbox_new, NONE, WIDGET, 2, BOOL, INT
func VBoxNew(homogeneous bool, spacing int) *VBox {
	id := Candy().Guify("gtk_vbox_new", homogeneous, spacing).String()
	return NewVBox(Candy(), id)
}
