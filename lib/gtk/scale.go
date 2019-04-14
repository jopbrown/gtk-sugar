package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Scale struct {
	Range
}

func NewScaleFromCandy(candy sugar.Candy, id string) *Scale {
	obj := Scale{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_scale_set_draw_value, NONE, NONE, 2, WIDGET, BOOL
func (obj *Scale) SetDrawValue(drawValue bool) {
	obj.Candy().Guify("gtk_scale_set_draw_value", obj, drawValue)
}

// FUNCTION_NAME = gtk_scale_set_value_pos, NONE, NONE, 2, WIDGET, INT
func (obj *Scale) SetValuePos(pos PositionType) {
	obj.Candy().Guify("gtk_scale_set_value_pos", obj, pos)
}

type HScale struct {
	Scale
}

func NewHScaleFromCandy(candy sugar.Candy, id string) *HScale {
	obj := HScale{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_hscale_new_with_range, value-changed, WIDGET, 3, DOUBLE, DOUBLE, DOUBLE
func NewHScaleWithRange(min, max, step float64) *HScale {
	id := Candy().Guify("gtk_hscale_new_with_range", min, max, step).String()
	return NewHScaleFromCandy(Candy(), id)
}

type VScale struct {
	Scale
}

func NewVScaleFromCandy(candy sugar.Candy, id string) *VScale {
	obj := VScale{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_vscale_new_with_range, value-changed, WIDGET, 3, DOUBLE, DOUBLE, DOUBLE
func NewVScaleWithRange(min, max, step float64) *VScale {
	id := Candy().Guify("gtk_hscale_new_with_range", min, max, step).String()
	return NewVScaleFromCandy(Candy(), id)
}
