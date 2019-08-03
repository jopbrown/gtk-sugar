package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Misc struct {
	Widget
}

func NewMisc(candy sugar.Candy, id string) *Misc {
	v := Misc{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_misc_set_alignment, NONE, NONE, 3, WIDGET, FLOAT, FLOAT
func (obj *Misc) SetAlignment(xalign, yalign float32) {
	obj.Candy().Guify("gtk_misc_set_alignment", obj, xalign, yalign)
}

// FUNCTION_NAME = gtk_misc_set_padding, NONE, NONE, 3, WIDGET, INT, INT
func (obj *Misc) SetPadding(xpad, ypad int) {
	obj.Candy().Guify("gtk_misc_set_padding", obj, xpad, ypad)
}

// FUNCTION_NAME = gtk_misc_get_alignment, NONE, NONE, 3, WIDGET, PTR_FLOAT, PTR_FLOAT
func (obj *Misc) GetAlignment() (xalign, yalign float32) {
	fields := obj.Candy().Guify("gtk_misc_get_alignment", obj, 0, 0).Fields()
	xalign = fields[0].MustFloat32()
	yalign = fields[1].MustFloat32()
	return
}

// FUNCTION_NAME = gtk_misc_get_padding, NONE, NONE, 3, WIDGET, PTR_INT, PTR_INT
func (obj *Misc) GetPadding() (xpad, ypad int) {
	fields := obj.Candy().Guify("gtk_misc_get_padding", obj, 0, 0).Fields()
	xpad = fields[0].MustInt()
	ypad = fields[1].MustInt()
	return
}
