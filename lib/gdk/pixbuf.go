package gdk

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/glibs"
	"github.com/jopbrown/gtk-sugar/util/fs"
	"github.com/pkg/errors"
)

type Pixbuf struct {
	glibs.Object
}

func NewPixbufFromCandy(candy sugar.Candy, id string) *Pixbuf {
	obj := Pixbuf{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gdk_pixbuf_new, NONE, WIDGET, 5, INT, BOOL, INT, INT, INT
func NewPixbuf(colorspace Colorspace, hasAlpha bool, bitsPerSample, width, height int) *Pixbuf {
	id := Candy().Guify("gdk_pixbuf_new", colorspace, hasAlpha, bitsPerSample, width, height).String()
	return NewPixbufFromCandy(Candy(), id)
}

// FUNCTION_NAME = gdk_pixbuf_new_from_file, NONE, WIDGET, 2, STRING, NULL
func NewPixbufFromFile(filename string) (*Pixbuf, error) {
	if err := fs.CheckExistsFile(filename); err != nil {
		return nil, errors.WithStack(err)
	}
	id := Candy().Guify("gdk_pixbuf_new_from_file", filename, 0).String()
	if id == "0" {
		return nil, errors.Errorf("unable to load pixbuf from %s", filename)
	}
	return NewPixbufFromCandy(Candy(), id), nil
}

// FUNCTION_NAME = gdk_pixbuf_new_from_file_at_size, NONE, WIDGET, 4, STRING, INT, INT, NULL
func NewPixbufFromFileAtSize(filename string, width, height int) (*Pixbuf, error) {
	if err := fs.CheckExistsFile(filename); err != nil {
		return nil, errors.WithStack(err)
	}
	id := Candy().Guify("gdk_pixbuf_new_from_file_at_size", filename, width, height, 0).String()
	if id == "0" {
		return nil, errors.Errorf("unable to load pixbuf from %s", filename)
	}
	return NewPixbufFromCandy(Candy(), id), nil
}

// FUNCTION_NAME = gdk_pixbuf_new_from_file_at_scale, NONE, WIDGET, 5, STRING, INT, INT, BOOL, NULL
func NewPixbufFromFileAtScale(filename string, width, height int, preserveAspectRatio bool) (*Pixbuf, error) {
	if err := fs.CheckExistsFile(filename); err != nil {
		return nil, errors.WithStack(err)
	}
	id := Candy().Guify("gdk_pixbuf_new_from_file_at_scale", filename, width, height, preserveAspectRatio, 0).String()
	if id == "0" {
		return nil, errors.Errorf("unable to load pixbuf from %s", filename)
	}
	return NewPixbufFromCandy(Candy(), id), nil
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
