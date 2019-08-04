package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Invisible struct {
	Widget
}

func NewInvisible(candy sugar.Candy, id string) *Invisible {
	obj := Invisible{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_invisible_new, NONE, WIDGET, 0
func InvisibleNew() *Invisible {
	id := Candy().Guify("gtk_invisible_new").String()
	return NewInvisible(Candy(), id)
}
