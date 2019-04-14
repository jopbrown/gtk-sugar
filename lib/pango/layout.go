package pango

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/glib"
)

type Layout struct {
	glib.Object
}

func NewLayoutFromCandy(candy sugar.Candy, id string) *Layout {
	obj := Layout{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = pango_layout_set_text, NONE, NONE, 3, WIDGET, STRING, INT
// FUNCTION_NAME = pango_layout_set_font_description, NONE, NONE, 2, WIDGET, WIDGET
// FUNCTION_NAME = pango_layout_get_size, NONE, NONE, 3, WIDGET, PTR_INT, PTR_INT
// FUNCTION_NAME = pango_layout_get_width, NONE, INT, 1, WIDGET
// FUNCTION_NAME = pango_layout_get_height, NONE, INT, 1, WIDGET
