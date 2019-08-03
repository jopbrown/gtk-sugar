package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type ScrolledWindow struct {
	Bin
}

func NewScrolledWindowFromCandy(candy sugar.Candy, id string) *ScrolledWindow {
	v := ScrolledWindow{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_scrolled_window_new, NONE, WIDGET, 2, WIDGET, WIDGET
func ScrolledWindowNew(h, v *Adjustment) *ScrolledWindow {
	id := Candy().Guify("gtk_scrolled_window_new", h, v).String()
	return NewScrolledWindowFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_scrolled_window_set_policy, NONE, NONE, 3, WIDGET, INT, INT
func (sw *ScrolledWindow) SetPolicy(hpolicy, vpolicy PolicyType) {
	sw.Candy().Guify("gtk_scrolled_window_set_policy", sw, hpolicy, vpolicy)
}

// FUNCTION_NAME = gtk_scrolled_window_set_shadow_type, NONE, NONE, 2, WIDGET, INT
func (sw *ScrolledWindow) SetShadowType(t ShadowType) {
	sw.Candy().Guify("gtk_scrolled_window_set_shadow_type", sw, t)
}

// FUNCTION_NAME = gtk_scrolled_window_add_with_viewport, NONE, NONE, 2, WIDGET, WIDGET
func (sw *ScrolledWindow) AddWithViewport(widget IWidget) {
	sw.Candy().Guify("gtk_scrolled_window_add_with_viewport", sw, widget)
}

// FUNCTION_NAME = gtk_scrolled_window_get_hadjustment, NONE, WIDGET, 1, WIDGET
func (sw *ScrolledWindow) GetHdjustment() *Adjustment {
	id := sw.Candy().Guify("gtk_scrolled_window_get_hadjustment", sw).String()
	a := Adjustment{}
	a.CandyWrapper = sw.Candy().NewWrapper(id)
	return &a
}

// FUNCTION_NAME = gtk_scrolled_window_get_vadjustment, NONE, WIDGET, 1, WIDGET
func (sw *ScrolledWindow) GetVdjustment() *Adjustment {
	id := sw.Candy().Guify("gtk_scrolled_window_get_vadjustment", sw).String()
	a := Adjustment{}
	a.CandyWrapper = sw.Candy().NewWrapper(id)
	return &a
}

// FUNCTION_NAME = gtk_scrolled_window_set_hadjustment, NONE, NONE, 2, WIDGET, WIDGET
func (sw *ScrolledWindow) SetHdjustment(a *Adjustment) {
	sw.Candy().Guify("gtk_scrolled_window_set_hadjustment", sw, a)
}

// FUNCTION_NAME = gtk_scrolled_window_set_vadjustment, NONE, NONE, 2, WIDGET, WIDGET
func (sw *ScrolledWindow) SetVdjustment(a *Adjustment) {
	sw.Candy().Guify("gtk_scrolled_window_set_vadjustment", sw, a)
}

// FUNCTION_NAME = gtk_scrolled_window_set_placement, NONE, NONE, 2, WIDGET, INT
func (sw *ScrolledWindow) SetPlacement(placement int) {
	sw.Candy().Guify("gtk_scrolled_window_set_placement", sw, placement)
}
