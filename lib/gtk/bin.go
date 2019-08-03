package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Bin struct {
	Container
}

func NewBinFromCandy(candy sugar.Candy, id string) *Bin {
	v := Bin{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_bin_get_child, NONE, WIDGET, 1, WIDGET
func (bin *Bin) GetChild() *Widget {
	id := bin.Candy().Guify("gtk_bin_get_child").String()
	return NewWidgetFromCandy(bin.Candy(), id)
}

type EventBox struct {
	Bin
}

func NewEventBoxFromCandy(candy sugar.Candy, id string) *EventBox {
	obj := EventBox{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_event_box_new, NONE, WIDGET, 0
func EventBoxNew() *EventBox {
	id := Candy().Guify("gtk_event_box_new").String()
	return NewEventBoxFromCandy(Candy(), id)
}
