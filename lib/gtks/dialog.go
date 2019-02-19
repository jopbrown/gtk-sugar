package gtks

import (
	"fmt"

	sugar "github.com/jopbrown/gtk-sugar"
)

type DialogButton struct {
	Name string
	ID   int
}

func WithDialogButton(name string, id int) DialogButton {
	return DialogButton{Name: name, ID: id}
}

type Dialog struct {
	Window
}

func NewDialog(candy sugar.Candy, id string) *Dialog {
	v := Dialog{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_dialog_run, NONE, INT, 1, WIDGET
func (w *Dialog) Run() int {
	return w.Candy().Guify("gtk_dialog_run", w).MustInt()
}

type MessageDialog struct {
	Dialog
}

func NewMessageDialog(candy sugar.Candy, id string) *MessageDialog {
	v := MessageDialog{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_message_dialog_new, response, WIDGET, 6, WIDGET, INT, INT, INT, STRING, STRING
func (gtk *Gtk) NewMessageDialog(parent *Window, flags DialogFlags, msgType MessageType, buttons ButtonsType, format string, v ...interface{}) *MessageDialog {
	msg := fmt.Sprintf(format, v...)
	id := gtk.Guify("gtk_message_dialog_new", parent, flags, msgType, buttons, msg, "''").String()
	return NewMessageDialog(gtk, id)
}

// FUNCTION_NAME = gtk_message_dialog_new_with_markup, response, WIDGET, 6, WIDGET, INT, INT, INT, STRING, STRING
func (gtk *Gtk) NewMessageDialogWithMarkup(parent *Window, flags DialogFlags, msgType MessageType, buttons ButtonsType, format string, v ...interface{}) *MessageDialog {
	msg := fmt.Sprintf(format, v...)
	id := gtk.Guify("gtk_message_dialog_new_with_markup", parent, flags, msgType, buttons, msg, "''").String()
	return NewMessageDialog(gtk, id)
}

type FontSelectionDialog struct {
	Dialog
}

func NewFontSelectionDialog(candy sugar.Candy, id string) *FontSelectionDialog {
	obj := FontSelectionDialog{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_font_selection_dialog_new, button-press-event, WIDGET, 1, STRING
func (gtk *Gtk) NewFontSelectionDialog(title string) *FontSelectionDialog {
	id := gtk.Guify("gtk_font_selection_dialog_new", title).String()
	return NewFontSelectionDialog(gtk, id)
}

// FUNCTION_NAME = gtk_font_selection_dialog_get_font_name, NONE, STRING, 1, WIDGET
func (obj *FontSelectionDialog) GetFontName() string {
	return obj.Candy().Guify("gtk_font_selection_dialog_get_font_name", obj).String()
}

// FUNCTION_NAME = gtk_font_selection_dialog_set_font_name, NONE, NONE, 2, WIDGET, STRING
func (obj *FontSelectionDialog) SetFontName(name string) {
	obj.Candy().Guify("gtk_font_selection_dialog_set_font_name", obj, name)
}
