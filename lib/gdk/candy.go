package gdk

import sugar "github.com/jopbrown/gtk-sugar"

var candy sugar.Candy

func GiveCandy(c sugar.Candy) {
	candy = c
}

func Candy() sugar.Candy {
	return candy
}

func Invoke(callback func()) {
	Candy().Invoke(callback)
}

func Main() {
	Candy().Main()
}

func MainQuit() {
	Candy().MainQuit()
}

func init() {
	sugar.RegisterGlobalCandyGiver("gdk", GiveCandy)
}
