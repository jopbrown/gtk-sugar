package gtk

// FUNCTION_NAME = gtk_combo_box_new_text, changed, WIDGET, 0
func ComboBoxNewWithText() *ComboBox {
	id := Candy().Guify("gtk_combo_box_new_text").String()
	return NewComboBox(Candy(), id)
}

// FUNCTION_NAME = gtk_combo_box_append_text, NONE, NONE, 2, WIDGET, STRING
func (obj *ComboBox) AppendText(text string) {
	obj.Candy().Guify("gtk_combo_box_append_text", obj, text)
}

// FUNCTION_NAME = gtk_combo_box_prepend_text, NONE, NONE, 2, WIDGET, STRING
func (obj *ComboBox) PrependText(text string) {
	obj.Candy().Guify("gtk_combo_box_prepend_text", obj, text)
}

// FUNCTION_NAME = gtk_combo_box_insert_text, NONE, NONE, 3, WIDGET, INT, STRING
func (obj *ComboBox) InsertText(position int, text string) {
	obj.Candy().Guify("gtk_combo_box_insert_text", obj, position, text)
}

// FUNCTION_NAME = gtk_combo_box_remove_text, NONE, NONE, 2, WIDGET, INT
func (obj *ComboBox) RemoveText(position int) {
	obj.Candy().Guify("gtk_combo_box_remove_text", obj, position)
}

// FUNCTION_NAME = gtk_combo_box_get_active_text, NONE, STRING, 1, WIDGET
func (obj *ComboBox) GetActiveText() string {
	return obj.Candy().Guify("gtk_combo_box_get_active_text", obj).String()
}
