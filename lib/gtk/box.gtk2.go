package gtk

// FUNCTION_NAME = gtk_box_pack_start_defaults, NONE, NONE, 2, WIDGET, WIDGET
func (box *Box) PackStartDefaults(child IWidget) {
	box.Candy().Guify("gtk_box_pack_start_defaults", box, child)
}

// FUNCTION_NAME = gtk_box_pack_end_defaults, NONE, NONE, 2, WIDGET, WIDGET
func (box *Box) PackEndDefaults(child IWidget) {
	box.Candy().Guify("gtk_box_pack_end_defaults", box, child)
}
