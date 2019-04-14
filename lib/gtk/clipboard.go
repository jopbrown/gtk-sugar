package gtk

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/gdks"
	"github.com/jopbrown/gtk-sugar/lib/glibs"
)

type Clipboard struct {
	glibs.Object
}

func NewClipboardFromCandy(candy sugar.Candy, id string) *Clipboard {
	v := Clipboard{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_clipboard_get, NONE, WIDGET, 1, INT
func NewClipboard(selection gdks.Atom) *Clipboard {
	id := Candy().Guify("gtk_clipboard_get").String()
	return NewClipboardFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_clipboard_clear, NONE, NONE, 1, WIDGET
func (obj *Clipboard) Clear() {
	obj.Candy().Guify("gtk_clipboard_clear", obj)
}

// FUNCTION_NAME = gtk_clipboard_set_text, NONE, NONE, 3, WIDGET, STRING, INT
func (obj *Clipboard) SetText(text string) {
	obj.Candy().Guify("gtk_clipboard_set_text", obj, text, len(text))
}

// FUNCTION_NAME = gtk_clipboard_wait_for_text, NONE, STRING, 1, WIDGET
func (obj *Clipboard) WaitForText() string {
	return obj.Candy().Guify("gtk_clipboard_wait_for_text", obj).String()
}

// FUNCTION_NAME = gtk_clipboard_set_image, NONE, NONE, 2, WIDGET, WIDGET
func (obj *Clipboard) SetImage(pixbuf *gdks.Pixbuf) {
	obj.Candy().Guify("gtk_clipboard_set_image", obj, pixbuf)
}

// FUNCTION_NAME = gtk_clipboard_wait_for_image, NONE, WIDGET, 1, WIDGET
func (obj *Clipboard) WaitForImage() *gdks.Pixbuf {
	id := obj.Candy().Guify("gtk_clipboard_wait_for_image", obj).String()
	return gdks.NewPixbufFromCandy(obj.Candy(), id)
}
