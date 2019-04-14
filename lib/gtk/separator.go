package gtks

import sugar "github.com/jopbrown/gtk-sugar"

type Separator struct {
	Widget
}

type VSeparator struct {
	Separator
}

func NewVSeparatorFromCandy(candy sugar.Candy, id string) *VSeparator {
	obj := VSeparator{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_vseparator_new, NONE, WIDGET, 0
func NewVSeparator() *VSeparator {
	id := Candy().Guify("gtk_vseparator_new").String()
	return NewVSeparatorFromCandy(Candy(), id)
}

type HSeparator struct {
	Separator
}

func NewHSeparatorFromCandy(candy sugar.Candy, id string) *HSeparator {
	obj := HSeparator{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_hseparator_new, NONE, WIDGET, 0
func NewHSeparator() *HSeparator {
	id := Candy().Guify("gtk_hseparator_new").String()
	return NewHSeparatorFromCandy(Candy(), id)
}
