package gtks

import sugar "github.com/jopbrown/gtk-sugar"

type ComboBox struct {
	Bin
}

func NewComboBox(candy sugar.Candy, id string) *ComboBox {
	obj := ComboBox{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_combo_box_get_active, NONE, INT, 1, WIDGET
func (obj *ComboBox) GetActive() int {
	return obj.Candy().Guify("gtk_combo_box_get_active", obj).MustInt()
}

// FUNCTION_NAME = gtk_combo_box_set_active, NONE, NONE, 2, WIDGET, INT
func (obj *ComboBox) SetActive(index int) {
	obj.Candy().Guify("gtk_combo_box_get_active", obj, index)
}

type ComboBoxText struct {
	ComboBox
}

func NewComboBoxText(candy sugar.Candy, id string) *ComboBoxText {
	obj := ComboBoxText{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_combo_box_text_new, changed, WIDGET, 0
func (gtk *Gtk) NewComboBoxText() *ComboBoxText {
	id := gtk.Guify("gtk_combo_box_text_new").String()
	return NewComboBoxText(gtk, id)
}

// FUNCTION_NAME = gtk_combo_box_text_append_text, NONE, NONE, 2, WIDGET, STRING
func (obj *ComboBoxText) AppendText(text string) {
	obj.Candy().Guify("gtk_combo_box_text_append_text", obj, text)
}

// FUNCTION_NAME = gtk_combo_box_text_prepend_text, NONE, NONE, 2, WIDGET, STRING
func (obj *ComboBoxText) PrependText(text string) {
	obj.Candy().Guify("gtk_combo_box_text_prepend_text", obj, text)
}

// FUNCTION_NAME = gtk_combo_box_text_insert_text, NONE, NONE, 3, WIDGET, INT, STRING
func (obj *ComboBoxText) InsertText(position int, text string) {
	obj.Candy().Guify("gtk_combo_box_text_insert_text", obj, position, text)
}

// FUNCTION_NAME = gtk_combo_box_text_remove, NONE, NONE, 2, WIDGET, INT
func (obj *ComboBoxText) Remove(position int) {
	obj.Candy().Guify("gtk_combo_box_text_remove", obj, position)
}

// FUNCTION_NAME = gtk_combo_box_text_get_active_text, NONE, STRING, 1, WIDGET
func (obj *ComboBoxText) GetActiveText() string {
	return obj.Candy().Guify("gtk_combo_box_text_get_active_text", obj).String()
}
