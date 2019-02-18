package gdks

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/glibs"
)

type Pixbuf struct {
	glibs.Object
}

func NewPixbuf(candy sugar.Candy, id string) *Pixbuf {
	obj := Pixbuf{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gdk_pixbuf_new_from_file, NONE, WIDGET, 2, STRING, NULL
// FUNCTION_NAME = gdk_pixbuf_new_from_file_at_size, NONE, WIDGET, 4, STRING, INT, INT, NULL
// FUNCTION_NAME = gdk_pixbuf_new_from_file_at_scale, NONE, WIDGET, 5, STRING, INT, INT, BOOL, NULL
// FUNCTION_NAME = gdk_pixbuf_rotate_simple, NONE, WIDGET, 2, WIDGET, INT
// FUNCTION_NAME = gdk_pixbuf_get_height, NONE, INT, 1, WIDGET
// FUNCTION_NAME = gdk_pixbuf_get_width, NONE, INT, 1, WIDGET
// FUNCTION_NAME = gdk_pixbuf_scale_simple, NONE, WIDGET, 4, WIDGET, INT, INT, INT
// FUNCTION_NAME = gdk_pixbuf_new, NONE, WIDGET, 5, INT, BOOL, INT, INT, INT
// FUNCTION_NAME = gdk_pixbuf_get_from_surface, NONE, WIDGET, 5, WIDGET, INT, INT, INT, INT
