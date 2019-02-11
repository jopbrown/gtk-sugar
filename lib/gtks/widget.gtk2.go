package gtks

import "github.com/jopbrown/gtk-sugar/lib/gdks"

// FUNCTION_NAME = gtk_widget_get_colormap, NONE, WIDGET, 1, WIDGET
func (widget *Widget) GetColormap() *gdks.Colormap {
	id := widget.Candy().Guify("gtk_widget_get_colormap", widget).String()
	return gdks.NewColormap(widget.Candy(), id)
}
