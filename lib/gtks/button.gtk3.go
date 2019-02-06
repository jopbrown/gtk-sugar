package gtks

// FUNCTION_NAME = gtk_button_new_from_icon_name, clicked, WIDGET, 2, STRING, INT
func (gtk *Gtk) NewButtonFromIconName(name string, size IconSize) *Button {
	id := gtk.Guify("gtk_button_new_from_icon_name", name, size).String()
	btn := Button{}
	btn.CandyWrapper = gtk.NewWrapper(id)
	return &btn
}
