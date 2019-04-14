package gtk

import (
	"runtime"

	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/gdks"
)

type Image struct {
	Misc
}

func NewImageFromCandy(candy sugar.Candy, id string) *Image {
	obj := Image{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_image_new, NONE, WIDGET, 0
func NewImage() *Image {
	id := Candy().Guify("gtk_image_new").String()
	return NewImageFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_image_new_from_pixbuf, NONE, WIDGET, 1, WIDGET
func NewImageFromPixbuf(pixbuf *gdks.Pixbuf) *Image {
	id := Candy().Guify("gtk_image_new_from_pixbuf", pixbuf).String()
	return NewImageFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_image_new_from_file, NONE, WIDGET, 1, STRING
// FUNCTION_NAME = gtk_image_new_from_file_utf8, NONE, WIDGET, 1, STRING
func NewImageFromFile(filename string) *Image {
	gtkfunc := "gtk_image_new_from_file"
	if runtime.GOOS == "windows" {
		gtkfunc = "gtk_image_new_from_file_utf8"
	}
	id := Candy().Guify(gtkfunc, filename).String()
	return NewImageFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_image_new_from_stock, NONE, WIDGET, 2, STRING, INT
func NewImageFromStock(stockID string, size IconSize) *Image {
	id := Candy().Guify("gtk_image_new_from_stock", stockID, size).String()
	return NewImageFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_image_set_from_pixbuf, NONE, WIDGET, 2, WIDGET, WIDGET
func (obj *Image) SetFromPixbuf(pixbuf *gdks.Pixbuf) {
	obj.Candy().Guify("gtk_image_set_from_pixbuf", obj, pixbuf)
}

// FUNCTION_NAME = gtk_image_set_from_file, NONE, NONE, 2, WIDGET, STRING
func (obj *Image) SetFromFile(filename string) {
	obj.Candy().Guify("gtk_image_set_from_file", obj, filename)
}

// FUNCTION_NAME = gtk_image_get_pixbuf, NONE, WIDGET, 1, WIDGET
func (obj *Image) GetPixbuf() *gdks.Pixbuf {
	id := obj.Candy().Guify("gtk_image_get_pixbuf", obj).String()
	return gdks.NewPixbufFromCandy(obj.Candy(), id)
}
