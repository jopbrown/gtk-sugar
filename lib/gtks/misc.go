package gtks

import sugar "github.com/jopbrown/gtk-sugar"

type Misc struct {
	Widget
}

func NewMisc(candy sugar.Candy, id string) *Misc {
	v := Misc{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}
