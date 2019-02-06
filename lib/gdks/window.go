package gdks

type Window struct {
	Drawable
}

// FUNCTION_NAME = gdk_window_process_all_updates, NONE, NONE, 0
func (gdk *Gdk) WindowProcessAllUpdates() {
	gdk.Guify("gdk_window_process_all_updates")
}

// FUNCTION_NAME = gdk_window_freeze_updates, NONE, NONE, 1, WIDGET
func (w *Window) FreezeUpdates() {
	w.Candy().Guify("gdk_window_freeze_updates", w)
}

// FUNCTION_NAME = gdk_window_thaw_updates, NONE, NONE, 1, WIDGET
func (w *Window) ThawUpdates() {
	w.Candy().Guify("gdk_window_thaw_updates", w)
}

// FUNCTION_NAME = gdk_window_get_geometry, NONE, NONE, 6, WIDGET, PTR_INT, PTR_INT, PTR_INT, PTR_INT, PTR_INT
func (w *Window) GetGeometry() (x, y, width, height, depth int) {
	fields := w.Candy().Guify("gdk_window_get_geometry", w, x, y, width, height, depth).Fields()
	x = fields[0].MustInt()
	y = fields[1].MustInt()
	width = fields[2].MustInt()
	height = fields[3].MustInt()
	depth = fields[4].MustInt()
	return
}

// FUNCTION_NAME = gdk_window_set_cursor, NONE, NONE, 2, WIDGET, WIDGET
func (w *Window) SetCursor(*Cursor) {
	w.Candy().Guify("gdk_window_set_cursor", w)
}

// FUNCTION_NAME = gdk_window_get_pointer, NONE, WIDGET, 4, WIDGET, PTR_INT, PTR_INT, PTR_WIDGET
// FUNCTION_NAME = gdk_window_create_similar_surface, NONE, WIDGET, 4, WIDGET, INT, INT, INT
