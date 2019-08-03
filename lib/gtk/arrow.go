package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Arrow struct {
	Misc
}

func NewArrow(candy sugar.Candy, id string) *Arrow {
	obj := Arrow{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_arrow_new, NONE, WIDGET, 2, INT, INT
func ArrowNew(arrowType ArrowType, shadowType ShadowType) *Arrow {
	id := Candy().Guify("gtk_arrow_new", arrowType, shadowType).String()
	return NewArrow(Candy(), id)
}
