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

// FUNCTION_NAME = gtk_file_chooser_dialog_new, NONE, WIDGET, 8, STRING, WIDGET, INT, STRING, INT, STRING, INT, NULL
// FUNCTION_NAME = gtk_file_chooser_widget_new, NONE, WIDGET, 1, INT
// FUNCTION_NAME = gtk_file_chooser_set_local_only, NONE, NONE, 2, WIDGET, BOOL
// FUNCTION_NAME = gtk_file_chooser_get_filename, NONE, STRING, 1, WIDGET
// FUNCTION_NAME = gtk_file_chooser_get_uri, NONE, STRING, 1, WIDGET
// FUNCTION_NAME = gtk_file_chooser_get_current_folder, NONE, STRING, 1, WIDGET
// FUNCTION_NAME = gtk_file_chooser_set_filename, NONE, BOOL, 2, WIDGET, STRING
// FUNCTION_NAME = gtk_file_filter_new, NONE, WIDGET, 0
// FUNCTION_NAME = gtk_file_filter_add_pattern, NONE, NONE, 2, WIDGET, STRING
// FUNCTION_NAME = gtk_file_filter_set_name, NONE, NONE, 2, WIDGET, STRING
// FUNCTION_NAME = gtk_file_chooser_add_filter, NONE, NONE, 2, WIDGET, WIDGET
// FUNCTION_NAME = gtk_file_chooser_set_select_multiple, NONE, NONE, 2, WIDGET, BOOL

// FUNCTION_NAME = gtk_font_selection_dialog_new, button-press-event, WIDGET, 1, STRING
// FUNCTION_NAME = gtk_font_selection_dialog_get_font_name, NONE, STRING, 1, WIDGET
// FUNCTION_NAME = gtk_font_selection_new, NONE, WIDGET, 0
// FUNCTION_NAME = gtk_font_selection_get_font_name, NONE, STRING, 1, WIDGET
// FUNCTION_NAME = gtk_font_selection_set_font_name, NONE, BOOL, 2, WIDGET, STRING

// FUNCTION_NAME = gtk_color_selection_new, NONE, WIDGET, 0
// FUNCTION_NAME = gtk_color_selection_set_has_opacity_control, NONE, NONE, 2, WIDGET, BOOL
// FUNCTION_NAME = gtk_color_selection_set_current_color, NONE, NONE, 2, WIDGET, STRING
// FUNCTION_NAME = gtk_color_selection_get_current_color, NONE, NONE, 2, WIDGET, WIDGET
// FUNCTION_NAME = gtk_color_selection_set_color, NONE, NONE, 2, WIDGET, STRING
