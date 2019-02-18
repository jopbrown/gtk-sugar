package gtks

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/glibs"
)

type AccelGroup struct {
	glibs.Object
}

func NewAccelGroup(candy sugar.Candy, id string) *AccelGroup {
	v := AccelGroup{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_accel_group_new, NONE, WIDGET, 0
func (gtk *Gtk) NewAccelGroup() *AccelGroup {
	id := gtk.Guify("gtk_accel_group_new").String()
	return NewAccelGroup(gtk, id)
}
