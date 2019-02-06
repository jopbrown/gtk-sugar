package gtks

import sugar "github.com/jopbrown/gtk-sugar"

type IWidget interface {
	sugar.CandyWrapper
	ShowAll()
	Show()
	Hide()
}

type Widget struct {
	sugar.CandyWrapper
}

func newWidget(candy sugar.Candy, id string) *Widget {
	return &Widget{CandyWrapper: sugar.NewCandyWrapper(candy, id)}
}

// FUNCTION_NAME = gtk_widget_show_all, NONE, NONE, 1, WIDGET
func (widget *Widget) ShowAll() {
	widget.Candy().Guify("gtk_widget_show_all", widget)
}

// FUNCTION_NAME = gtk_widget_show, NONE, NONE, 1, WIDGET
func (widget *Widget) Show() {
	widget.Candy().Guify("gtk_widget_show", widget)
}

// FUNCTION_NAME = gtk_widget_hide, NONE, NONE, 1, WIDGET
func (widget *Widget) Hide() {
	widget.Candy().Guify("gtk_widget_hide", widget)
}

// FUNCTION_NAME = gtk_widget_realize, NONE, NONE, 1, WIDGET
// FUNCTION_NAME = gtk_widget_unrealize, NONE, NONE, 1, WIDGET
// FUNCTION_NAME = gtk_widget_destroy, NONE, NONE, 1, WIDGET
// FUNCTION_NAME = gtk_widget_grab_focus, NONE, NONE, 1, WIDGET
// FUNCTION_NAME = gtk_widget_set_size_request, NONE, NONE, 3, WIDGET, INT, INT
// FUNCTION_NAME = gtk_widget_size_request, NONE, NONE, 2, WIDGET, WIDGET
// FUNCTION_NAME = gtk_widget_set_usize, NONE, NONE, 3, WIDGET, INT, INT
// FUNCTION_NAME = gtk_widget_modify_base, NONE, NONE, 3, WIDGET, INT, WIDGET
// FUNCTION_NAME = gtk_widget_override_background_color, NONE, NONE, 3, WIDGET, INT, WIDGET
// FUNCTION_NAME = gtk_widget_modify_text, NONE, NONE, 3, WIDGET, INT, WIDGET
// FUNCTION_NAME = gtk_widget_override_color, NONE, NONE, 3, WIDGET, INT, WIDGET
// FUNCTION_NAME = gtk_widget_modify_bg, NONE, NONE, 3, WIDGET, INT, WIDGET
// FUNCTION_NAME = gtk_widget_modify_fg, NONE, NONE, 3, WIDGET, INT, WIDGET
// FUNCTION_NAME = gtk_widget_modify_font, NONE, NONE, 2, WIDGET, WIDGET
// FUNCTION_NAME = gtk_widget_set_sensitive, NONE, NONE, 2, WIDGET, BOOL
func (w *Widget) SetSensitive(sensitive bool) {
	w.Candy().Guify("gtk_widget_set_sensitive", w, sensitive)
}

// FUNCTION_NAME = gtk_widget_add_accelerator, NONE, NONE, 6, WIDGET, STRING, WIDGET, INT, INT, INT
// FUNCTION_NAME = gtk_widget_get_parent, NONE, WIDGET, 1, WIDGET
// FUNCTION_NAME = gtk_widget_get_toplevel, NONE, WIDGET, 1, WIDGET
// FUNCTION_NAME = gtk_widget_set_name, NONE, NONE, 2, WIDGET, STRING
// FUNCTION_NAME = gtk_widget_get_size_request, NONE, NONE, 3, WIDGET, WIDGET, WIDGET
// FUNCTION_NAME = gtk_widget_add_events, NONE, NONE, 2, WIDGET, INT
// FUNCTION_NAME = gtk_widget_get_allocated_width, NONE, INT, 1, WIDGET
// FUNCTION_NAME = gtk_widget_get_allocated_height, NONE, INT, 1, WIDGET

// FUNCTION_NAME = gtk_widget_queue_draw, NONE, NONE, 1, WIDGET
// FUNCTION_NAME = gtk_widget_get_colormap, NONE, WIDGET, 1, WIDGET
// FUNCTION_NAME = gtk_widget_get_parent_window, NONE, WIDGET, 1, WIDGET
// FUNCTION_NAME = gtk_widget_create_pango_layout, NONE, WIDGET, 2, WIDGET, STRING
// FUNCTION_NAME = gtk_widget_get_window, NONE, WIDGET, 1, WIDGET
// FUNCTION_NAME = gtk_widget_set_tooltip_text, NONE, VOID, 2, WIDGET, STRING
