package pangos

import (
	sugar "github.com/jopbrown/gtk-sugar"
)

type FontDescription struct {
	sugar.CandyWrapper
}

func NewFontDescription(candy sugar.Candy, id string) *FontDescription {
	obj := FontDescription{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}
