package gtks

import sugar "github.com/jopbrown/gtk-sugar"

type Clipboard struct {
	sugar.CandyWrapper
}

func NewClipboard(candy sugar.Candy, id string) *Clipboard {
	v := Clipboard{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_clipboard_get, NONE, WIDGET, 1, INT
// FUNCTION_NAME = gtk_clipboard_set_text, NONE, NONE, 3, WIDGET, STRING, INT
// FUNCTION_NAME = gtk_clipboard_wait_for_text, NONE, STRING, 1, WIDGET
