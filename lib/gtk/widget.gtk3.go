package gtk

import "github.com/jopbrown/gtk-sugar/lib/gdk"

// FUNCTION_NAME = gtk_widget_override_background_color, NONE, NONE, 3, WIDGET, INT, WIDGET
func (w *Widget) OverrideBackgroundColor(state StateType, color *gdk.RGBA) {
	w.Candy().Guify("gtk_widget_override_background_color", w, state, color)
}

// FUNCTION_NAME = gtk_widget_override_color, NONE, NONE, 3, WIDGET, INT, WIDGET
func (w *Widget) OverrideColor(state StateType, color *gdk.RGBA) {
	w.Candy().Guify("gtk_widget_override_color", w, state, color)
}

// FUNCTION_NAME = gtk_widget_get_allocated_width, NONE, INT, 1, WIDGET
func (w *Widget) GetAllocatedWidth() int {
	return w.Candy().Guify("gtk_widget_get_allocated_width", w).MustInt()
}

// FUNCTION_NAME = gtk_widget_get_allocated_height, NONE, INT, 1, WIDGET
func (w *Widget) GetAllocatedHeight() int {
	return w.Candy().Guify("gtk_widget_get_allocated_height", w).MustInt()
}
