package gtk

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/glibs"
)

type AccelGroup struct {
	glibs.Object
}

func NewAccelGroupFromCandy(candy sugar.Candy, id string) *AccelGroup {
	v := AccelGroup{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_accel_group_new, NONE, WIDGET, 0
func NewAccelGroup() *AccelGroup {
	id := Candy().Guify("gtk_accel_group_new").String()
	return NewAccelGroupFromCandy(Candy(), id)
}
