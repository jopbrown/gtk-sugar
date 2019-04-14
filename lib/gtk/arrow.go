package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Arrow struct {
	Misc
}

func NewArrowFromCandy(candy sugar.Candy, id string) *Arrow {
	obj := Arrow{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_arrow_new, NONE, WIDGET, 2, INT, INT
func NewArrow(arrowType ArrowType, shadowType ShadowType) *Arrow {
	id := Candy().Guify("gtk_arrow_new", arrowType, shadowType).String()
	return NewArrowFromCandy(Candy(), id)
}
