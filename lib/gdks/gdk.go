package gdks

import sugar "github.com/jopbrown/gtk-sugar"

type Gdk struct {
	sugar.Candy
}

func NewGdk(candy sugar.Candy) *Gdk {
	return &Gdk{Candy: candy}
}
