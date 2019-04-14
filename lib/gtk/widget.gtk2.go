package gtk

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/gdks"
	"github.com/jopbrown/gtk-sugar/lib/pangos"
)

// FUNCTION_NAME = gtk_widget_get_colormap, NONE, WIDGET, 1, WIDGET
func (w *Widget) GetColormap() *gdks.Colormap {
	id := w.Candy().Guify("gtk_widget_get_colormap", w).String()
	return gdks.NewColormapFromCandy(w.Candy(), id)
}

// FUNCTION_NAME = gtk_widget_modify_base, NONE, NONE, 3, WIDGET, INT, WIDGET
func (w *Widget) ModifyBase(state StateType, color *gdks.Color) {
	w.Candy().Guify("gtk_widget_modify_base", w, state, color)
}

// FUNCTION_NAME = gtk_widget_modify_text, NONE, NONE, 3, WIDGET, INT, WIDGET
func (w *Widget) ModifyText(state StateType, color *gdks.Color) {
	w.Candy().Guify("gtk_widget_modify_text", w, state, color)
}

// FUNCTION_NAME = gtk_widget_modify_bg, NONE, NONE, 3, WIDGET, INT, WIDGET
func (w *Widget) ModifyBg(state StateType, color *gdks.Color) {
	w.Candy().Guify("gtk_widget_modify_bg", w, state, color)
}

// FUNCTION_NAME = gtk_widget_modify_fg, NONE, NONE, 3, WIDGET, INT, WIDGET
func (w *Widget) ModifyFg(state StateType, color *gdks.Color) {
	w.Candy().Guify("gtk_widget_modify_fg", w, state, color)
}

// FUNCTION_NAME = gtk_widget_modify_font, NONE, NONE, 2, WIDGET, WIDGET
func (w *Widget) ModifyFont(fontDesc *pangos.FontDescription) {
	w.Candy().Guify("gtk_widget_modify_font", w, fontDesc)
}

type Allocation gdks.Rectangle

// FUNCTION_NAME = gtk_widget_get_allocation, NONE, NONE, 2, WIDGET, PTR_BASE64
func (w *Widget) GetAllocation() *Allocation {
	a := Allocation{}
	packer := sugar.NewBase64Packer(&a)
	format := packer.Format()
	w.Candy().ServerDataFormat(format)
	base64Data := w.Candy().Guify("gtk_widget_get_allocation", w, 0).String()
	fields := w.Candy().ServerUnpack(format, base64Data)
	fields.MustUnmarshal(packer)
	return &a
}
