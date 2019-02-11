package gtks

type Bin struct {
	Container
}

// FUNCTION_NAME = gtk_bin_get_child, NONE, WIDGET, 1, WIDGET
func (bin *Bin) GetChild() *Widget {
	id := bin.Candy().Guify("gtk_bin_get_child").String()
	return NewWidget(bin.Candy(), id)
}
