package gtks

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/glibs"
)

type Tooltips struct {
	glibs.Object
}

func NewTooltips(candy sugar.Candy, id string) *Tooltips {
	obj := Tooltips{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_tooltips_new, NONE, WIDGET, 0
func (gtk *Gtk) NewTooltips() *Tooltips {
	id := gtk.Guify("gtk_tooltips_new").String()
	return NewTooltips(gtk, id)
}

// FUNCTION_NAME = gtk_tooltips_enable, NONE, NONE, 1, WIDGET
func (obj *Tooltips) Enable() {
	obj.Candy().Guify("gtk_tooltips_enable", obj)
}

// FUNCTION_NAME = gtk_tooltips_disable, NONE, NONE, 1, WIDGET
func (obj *Tooltips) Disable() {
	obj.Candy().Guify("gtk_tooltips_disable", obj)
}

// FUNCTION_NAME = gtk_tooltips_set_tip, NONE, NONE, 4, WIDGET, WIDGET, STRING, STRING
func (obj *Tooltips) SetTip(widget IWidget, tipText, tipPrivate string) {
	obj.Candy().Guify("gtk_tooltips_set_tip", obj, widget, tipText, tipPrivate)
}

// FUNCTION_NAME = gtk_tooltips_force_window, NONE, NONE, 1, WIDGET
func (obj *Tooltips) ForceWindow() {
	obj.Candy().Guify("gtk_tooltips_force_window", obj)
}
