package gtk

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

// FUNCTION_NAME = gtk_dialog_new, WIDGET, INT, 0
func DialogNew() *Dialog {
	id := Candy().Guify("gtk_dialog_new").String()
	return NewDialog(Candy(), id)
}

// FUNCTION_NAME = gtk_dialog_run, NONE, INT, 1, WIDGET
func (obj *Dialog) Run() int {
	return obj.Candy().Guify("gtk_dialog_run", obj).MustInt()
}

// FUNCTION_NAME = gtk_dialog_add_button, NONE, WIDGET, 3, WIDGET, STRING, INT
func (obj *Dialog) AddButton(buttonText string, responseID int) *Button {
	id := obj.Candy().Guify("gtk_dialog_add_button", obj, buttonText, responseID).String()
	return NewButton(obj.Candy(), id)
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
func MessageDialogNew(parent *Window, flags DialogFlags, msgType MessageType, buttons ButtonsType, format string, v ...interface{}) *MessageDialog {
	msg := fmt.Sprintf(format, v...)
	id := Candy().Guify("gtk_message_dialog_new", parent, flags, msgType, buttons, msg, "''").String()
	return NewMessageDialog(Candy(), id)
}

// FUNCTION_NAME = gtk_message_dialog_new_with_markup, response, WIDGET, 6, WIDGET, INT, INT, INT, STRING, STRING
func MessageDialogNewWithMarkup(parent *Window, flags DialogFlags, msgType MessageType, buttons ButtonsType, format string, v ...interface{}) *MessageDialog {
	msg := fmt.Sprintf(format, v...)
	id := Candy().Guify("gtk_message_dialog_new_with_markup", parent, flags, msgType, buttons, msg, "''").String()
	return NewMessageDialog(Candy(), id)
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
func FontSelectionDialogNew(title string) *FontSelectionDialog {
	id := Candy().Guify("gtk_font_selection_dialog_new", title).String()
	return NewFontSelectionDialog(Candy(), id)
}

// FUNCTION_NAME = gtk_font_selection_dialog_get_font_name, NONE, STRING, 1, WIDGET
func (obj *FontSelectionDialog) GetFontName() string {
	return obj.Candy().Guify("gtk_font_selection_dialog_get_font_name", obj).String()
}

// FUNCTION_NAME = gtk_font_selection_dialog_set_font_name, NONE, NONE, 2, WIDGET, STRING
func (obj *FontSelectionDialog) SetFontName(name string) {
	obj.Candy().Guify("gtk_font_selection_dialog_set_font_name", obj, name)
}
