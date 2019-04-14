package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Fixed struct {
	Container
}

func NewFixedFromCandy(candy sugar.Candy, id string) *Fixed {
	obj := Fixed{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_fixed_new, NONE, WIDGET, 0
func NewFixed() *Fixed {
	id := Candy().Guify("gtk_fixed_new").String()
	return NewFixedFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_fixed_put, NONE, NONE, 4, WIDGET, WIDGET, INT, INT
func (obj *Fixed) Put(widget IWidget, x, y int) {
	obj.Candy().Guify("gtk_fixed_put", obj, widget, x, y)
}

// FUNCTION_NAME = gtk_fixed_move, NONE, NONE, 4, WIDGET, WIDGET, INT, INT
func (obj *Fixed) Move(widget IWidget, x, y int) {
	obj.Candy().Guify("gtk_fixed_move", obj, widget, x, y)
}
