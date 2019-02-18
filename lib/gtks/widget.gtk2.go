package gtks

import (
	"github.com/jopbrown/gtk-sugar/lib/gdks"
	"github.com/jopbrown/gtk-sugar/lib/pangos"
)

// FUNCTION_NAME = gtk_widget_get_colormap, NONE, WIDGET, 1, WIDGET
func (widget *Widget) GetColormap() *gdks.Colormap {
	id := widget.Candy().Guify("gtk_widget_get_colormap", widget).String()
	return gdks.NewColormap(widget.Candy(), id)
}

// FUNCTION_NAME = gtk_widget_modify_base, NONE, NONE, 3, WIDGET, INT, WIDGET
func (widget *Widget) ModifyBase(state StateType, color *gdks.Color) {
	widget.Candy().Guify("gtk_widget_modify_base", widget, state, color)
}

// FUNCTION_NAME = gtk_widget_modify_text, NONE, NONE, 3, WIDGET, INT, WIDGET
func (widget *Widget) ModifyText(state StateType, color *gdks.Color) {
	widget.Candy().Guify("gtk_widget_modify_text", widget, state, color)
}

// FUNCTION_NAME = gtk_widget_modify_bg, NONE, NONE, 3, WIDGET, INT, WIDGET
func (widget *Widget) ModifyBg(state StateType, color *gdks.Color) {
	widget.Candy().Guify("gtk_widget_modify_bg", widget, state, color)
}

// FUNCTION_NAME = gtk_widget_modify_fg, NONE, NONE, 3, WIDGET, INT, WIDGET
func (widget *Widget) ModifyFg(state StateType, color *gdks.Color) {
	widget.Candy().Guify("gtk_widget_modify_fg", widget, state, color)
}

// FUNCTION_NAME = gtk_widget_modify_font, NONE, NONE, 2, WIDGET, WIDGET
func (widget *Widget) ModifyFont(fontDesc *pangos.FontDescription) {
	widget.Candy().Guify("gtk_widget_modify_font", widget, fontDesc)
}
