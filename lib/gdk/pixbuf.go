package gdk

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/glib"
)

type Pixbuf struct {
	glib.Object
}

func NewPixbufFromCandy(candy sugar.Candy, id string) *Pixbuf {
	obj := Pixbuf{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gdk_pixbuf_new, NONE, WIDGET, 5, INT, BOOL, INT, INT, INT
func PixbufNew(colorspace Colorspace, hasAlpha bool, bitsPerSample, width, height int) *Pixbuf {
	id := Candy().Guify("gdk_pixbuf_new", colorspace, hasAlpha, bitsPerSample, width, height).String()
	return NewPixbufFromCandy(Candy(), id)
}

// FUNCTION_NAME = gdk_pixbuf_new_from_file, NONE, WIDGET, 2, STRING, PTR_LONG
func PixbufNewFromFile(filename string) (*Pixbuf, error) {
	fields := Candy().Guify("gdk_pixbuf_new_from_file", filename, 0).Fields()
	if fields[0] == "0" {
		p := fields[1].String()
		err := glib.NewGErrorFromPointer(p)
		glib.FreePointer(p)
		return nil, err
	}

	return NewPixbufFromCandy(Candy(), fields[0].String()), nil
}

// FUNCTION_NAME = gdk_pixbuf_new_from_file_at_size, NONE, WIDGET, 4, STRING, INT, INT, PTR_LONG
func PixbufNewFromFileAtSize(filename string, width, height int) (*Pixbuf, error) {
	fields := Candy().Guify("gdk_pixbuf_new_from_file_at_size", filename, width, height, 0).Fields()
	if fields[0] == "0" {
		p := fields[1].String()
		err := glib.NewGErrorFromPointer(p)
		glib.FreePointer(p)
		return nil, err
	}

	return NewPixbufFromCandy(Candy(), fields[0].String()), nil
}

// FUNCTION_NAME = gdk_pixbuf_new_from_file_at_scale, NONE, WIDGET, 5, STRING, INT, INT, BOOL, PTR_LONG
func PixbufNewFromFileAtScale(filename string, width, height int, preserveAspectRatio bool) (*Pixbuf, error) {
	fields := Candy().Guify("gdk_pixbuf_new_from_file_at_scale", filename, width, height, preserveAspectRatio, 0).Fields()
	if fields[0] == "0" {
		p := fields[1].String()
		err := glib.NewGErrorFromPointer(p)
		glib.FreePointer(p)
		return nil, err
	}

	return NewPixbufFromCandy(Candy(), fields[0].String()), nil
}

// FUNCTION_NAME = gdk_pixbuf_get_type, NONE, INT, 0
func GetTypePixbuf() glib.Type {
	return glib.Type(Candy().Guify("gdk_pixbuf_get_type").MustInt())
}

// FUNCTION_NAME = gdk_pixbuf_scale_simple, NONE, WIDGET, 4, WIDGET, INT, INT, INT
func (obj *Pixbuf) ScaleSimple(destWidth, destHeight int, interpType InterpType) *Pixbuf {
	id := obj.Candy().Guify("gdk_pixbuf_scale_simple", obj, destWidth, destHeight, interpType).String()
	return NewPixbufFromCandy(obj.Candy(), id)
}

// FUNCTION_NAME = gdk_pixbuf_rotate_simple, NONE, WIDGET, 2, WIDGET, INT
func (obj *Pixbuf) RotateSimple(angle Rotation) *Pixbuf {
	id := obj.Candy().Guify("gdk_pixbuf_rotate_simple", obj, angle).String()
	return NewPixbufFromCandy(obj.Candy(), id)
}

// FUNCTION_NAME = gdk_pixbuf_get_height, NONE, INT, 1, WIDGET
func (obj *Pixbuf) GetHeight() int {
	return obj.Candy().Guify("gdk_pixbuf_get_height", obj).MustInt()
}

// FUNCTION_NAME = gdk_pixbuf_get_width, NONE, INT, 1, WIDGET
func (obj *Pixbuf) GetWidth() int {
	return obj.Candy().Guify("gdk_pixbuf_get_width", obj).MustInt()
}
