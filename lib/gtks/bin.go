package gtks

import sugar "github.com/jopbrown/gtk-sugar"

type Bin struct {
	Container
}

func NewBin(candy sugar.Candy, id string) *Bin {
	v := Bin{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_bin_get_child, NONE, WIDGET, 1, WIDGET
func (bin *Bin) GetChild() *Widget {
	id := bin.Candy().Guify("gtk_bin_get_child").String()
	return NewWidget(bin.Candy(), id)
}
