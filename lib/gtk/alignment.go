package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Alignment struct {
	Bin
}

func NewAlignment(candy sugar.Candy, id string) *Alignment {
	obj := Alignment{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_alignment_new, NONE, WIDGET, 4, FLOAT, FLOAT, FLOAT, FLOAT
func AlignmentNew(xalign, yalign, xscale, yscale float32) *Alignment {
	id := Candy().Guify("gtk_alignment_new", xalign, yalign, xscale, yscale).String()
	return NewAlignment(Candy(), id)
}

// FUNCTION_NAME = gtk_alignment_set, NONE, NONE, 5, WIDGET, FLOAT, FLOAT, FLOAT, FLOAT
func (obj *FileChooser) Set(xalign, yalign, xscale, yscale float32) {
	obj.Candy().Guify("gtk_alignment_set", obj, xalign, yalign, xscale, yscale)
}

// FUNCTION_NAME = gtk_alignment_set_padding, NONE, NONE, 5, WIDGET, INT, INT, INT, INT
func (obj *FileChooser) SetPadding(top, bottom, left, right uint) {
	obj.Candy().Guify("gtk_alignment_set_padding", obj, top, bottom, left, right)
}

// FUNCTION_NAME = gtk_alignment_get_padding, NONE, NONE, 5, WIDGET, PTR_INT, PTR_INT, PTR_INT, PTR_INT
func (obj *FileChooser) GetPadding() (top, bottom, left, right uint) {
	fields := obj.Candy().Guify("gtk_alignment_get_padding", obj, 0, 0, 0, 0).Fields()
	top = fields[0].MustUint()
	bottom = fields[1].MustUint()
	left = fields[2].MustUint()
	right = fields[3].MustUint()
	return
}
