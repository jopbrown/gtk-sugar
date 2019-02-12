package gtks

import (
	"fmt"

	sugar "github.com/jopbrown/gtk-sugar"
)

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
