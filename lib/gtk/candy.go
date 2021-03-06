package gtk

import sugar "github.com/jopbrown/gtk-sugar"

var candy sugar.Candy

func GiveCandy(c sugar.Candy) {
	candy = c
}

func Candy() sugar.Candy {
	return candy
}

func Invoke(action func()) {
	Candy().Invoke(action)
}

func init() {
	sugar.RegisterCandyGiver("gtk", GiveCandy)
}
