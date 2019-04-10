package gdks

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/glibs"
)

type RGBA struct {
	sugar.CandyWrapper
}

func NewRGBA(candy sugar.Candy, id string) *RGBA {
	obj := RGBA{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// NewRGBAFromSpec, need to Free() after used
func (gdk *Gdk) NewRGBAFromSpec(spec string) *RGBA {
	id := gdk.ServerOpaque()
	obj := NewRGBA(gdk, id)
	obj.Parse(spec)
	return obj
}

// FUNCTION_NAME = gdk_rgba_parse, NONE, BOOL, 2, WIDGET, STRING
func (obj *RGBA) Parse(spec string) bool {
	return obj.Candy().Guify("gdk_rgba_parse", obj, spec).MustBool()
}

func (obj *RGBA) Free() {
	glibs.Free(obj)
}
