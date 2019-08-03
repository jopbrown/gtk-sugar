package gtk

// FUNCTION_NAME = gtk_button_new_from_icon_name, clicked, WIDGET, 2, STRING, INT
func ButtonNewFromIconName(name string, size IconSize) *Button {
	id := Candy().Guify("gtk_button_new_from_icon_name", name, size).String()
	return NewButton(Candy(), id)
}
