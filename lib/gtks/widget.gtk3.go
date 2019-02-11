package gtks

import "github.com/jopbrown/gtk-sugar/lib/gdks"

// FUNCTION_NAME = gtk_widget_override_background_color, NONE, NONE, 3, WIDGET, INT, WIDGET
func (widget *Widget) OverrideBackgroundColor(state StateType, color *gdks.RGBA) {
	widget.Candy().Guify("gtk_widget_override_background_color", widget, state, color)
}

// FUNCTION_NAME = gtk_widget_override_color, NONE, NONE, 3, WIDGET, INT, WIDGET
func (widget *Widget) OverrideColor(state StateType, color *gdks.RGBA) {
	widget.Candy().Guify("gtk_widget_override_color", widget, state, color)
}

// FUNCTION_NAME = gtk_widget_get_allocated_width, NONE, INT, 1, WIDGET
func (widget *Widget) GetAllocatedWidth() int {
	return widget.Candy().Guify("gtk_widget_get_allocated_width", widget).MustInt()
}

// FUNCTION_NAME = gtk_widget_get_allocated_height, NONE, INT, 1, WIDGET
func (widget *Widget) GetAllocatedHeight() int {
	return widget.Candy().Guify("gtk_widget_get_allocated_height", widget).MustInt()
}
