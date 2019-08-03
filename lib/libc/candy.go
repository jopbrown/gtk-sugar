package libc

import sugar "github.com/jopbrown/gtk-sugar"

var candy sugar.Candy

func GiveCandy(c sugar.Candy) {
	candy = c
}

func Candy() sugar.Candy {
	return candy
}

func init() {
	sugar.RegisterCandyGiver("libc", GiveCandy)
}
