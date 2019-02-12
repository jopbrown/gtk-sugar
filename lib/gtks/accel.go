package gtks

import sugar "github.com/jopbrown/gtk-sugar"

type AccelGroup struct {
	sugar.CandyWrapper
}

func NewAccelGroup(candy sugar.Candy, id string) *AccelGroup {
	v := AccelGroup{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}
