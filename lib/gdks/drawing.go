package gdks

import sugar "github.com/jopbrown/gtk-sugar"

type Color struct {
	sugar.CandyWrapper
}

func NewColor(candy sugar.Candy, id string) *Color {
	obj := Color{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}
