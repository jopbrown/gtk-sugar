package gdks

import sugar "github.com/jopbrown/gtk-sugar"

type RGBA struct {
	sugar.CandyWrapper
}

func NewRGBA(candy sugar.Candy, id string) *RGBA {
	obj := RGBA{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}
