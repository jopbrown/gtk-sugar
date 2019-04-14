package pangos

import (
	sugar "github.com/jopbrown/gtk-sugar"
)

type FontDescription struct {
	sugar.CandyWrapper
}

func NewFontDescriptionFromCandy(candy sugar.Candy, id string) *FontDescription {
	obj := FontDescription{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = pango_font_description_from_string, NONE, WIDGET, 1, STRING
// FUNCTION_NAME = pango_font_description_free, NONE, NONE, 1, WIDGET
