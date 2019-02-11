package pangos

import sugar "github.com/jopbrown/gtk-sugar"

type Layout struct {
	sugar.CandyWrapper
}

func NewLayout(candy sugar.Candy, id string) *Layout {
	obj := Layout{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}
