package gdks

import sugar "github.com/jopbrown/gtk-sugar"

type IDrawable interface {
	sugar.CandyWrapper
}

type Drawable struct {
	sugar.CandyWrapper
}

type Cursor struct {
	sugar.CandyWrapper
}

type Color struct {
	sugar.CandyWrapper
}

type Colormap struct {
	sugar.CandyWrapper
}

type GC struct {
	sugar.CandyWrapper
}

// FUNCTION_NAME = gdk_gc_new, NONE, WIDGET, 1, WIDGET
func (gdk *Gdk) NewGC(drawable IDrawable) *GC {
	id := gdk.Guify("gdk_gc_new", drawable).String()
	w := GC{}
	w.CandyWrapper = gdk.NewWrapper(id)
	return &w
}

// FUNCTION_NAME = gdk_gc_set_rgb_fg_color, NONE, NONE, 2, WIDGET, WIDGET
func (gc *GC) SetRgbFgColor(color *Color) {
	gc.Candy().Guify("gdk_gc_set_rgb_fg_color", gc, color)
}

// FUNCTION_NAME = gdk_gc_set_rgb_bg_color, NONE, NONE, 2, WIDGET, WIDGET
func (gc *GC) SetRgbBgColor(color *Color) {
	gc.Candy().Guify("gdk_gc_set_rgb_bg_color", gc, color)
}

// FUNCTION_NAME = gdk_gc_set_foreground, NONE, NONE, 2, WIDGET, WIDGET
func (gc *GC) SetForeground(color *Color) {
	gc.Candy().Guify("gdk_gc_set_foreground", gc, color)
}

// FUNCTION_NAME = gdk_gc_set_background, NONE, NONE, 2, WIDGET, WIDGET
func (gc *GC) SetBackground(color *Color) {
	gc.Candy().Guify("gdk_gc_set_background", gc, color)
}

// FUNCTION_NAME = gdk_gc_set_colormap, NONE, NONE, 2, WIDGET, WIDGET
func (gc *GC) SetColormap(colormap *Colormap) {
	gc.Candy().Guify("gdk_gc_set_colormap", gc, colormap)
}

type Font struct {
	sugar.CandyWrapper
}

// FUNCTION_NAME = gdk_font_load, NONE, WIDGET, 1, STRING
func (gdk *Gdk) FontLoad(fontName string) *Font {
	id := gdk.Guify("gdk_font_load", fontName).String()
	w := Font{}
	w.CandyWrapper = gdk.NewWrapper(id)

	return &w
}

type Pixmap struct {
	Drawable
}

// FUNCTION_NAME = gdk_pixmap_new, NONE, WIDGET, 4, WIDGET, INT, INT, INT
func (gdk *Gdk) NewPixmap(drawable IDrawable, width, height, depth int) *Pixmap {
	id := gdk.Guify("gdk_pixmap_new", drawable, width, height, depth).String()
	w := Pixmap{}
	w.CandyWrapper = gdk.NewWrapper(id)
	return &w
}

// FUNCTION_NAME = gdk_pixmap_unref, NONE, NONE, 1, WIDGET
func (p *Pixmap) Unref() {
	p.Candy().Guify("gdk_pixmap_unref", p)
}

// FUNCTION_NAME = gdk_draw_rectangle, NONE, NONE, 7, WIDGET, WIDGET, BOOL, INT, INT, INT, INT
func (gdk *Gdk) DrawRectangle(drawable IDrawable, gc *GC, filled bool, x, y, width, height int) {
	gdk.Guify("gdk_draw_rectangle", drawable, gc, filled, x, y, width, height)
}

// FUNCTION_NAME = gdk_draw_arc, NONE, NONE, 9, WIDGET, WIDGET, BOOL, INT, INT, INT, INT, INT, INT
func (gdk *Gdk) DrawArc(drawable IDrawable, gc *GC, filled bool, x, y, width, height, angle1, angle2 int) {
	gdk.Guify("gdk_draw_arc", drawable, gc, filled, x, y, width, height, angle1, angle2)
}

// FUNCTION_NAME = gdk_draw_line, NONE, NONE, 6, WIDGET, WIDGET, INT, INT, INT, INT
func (gdk *Gdk) DrawLine(drawable IDrawable, gc *GC, x1, y1, x2, y2 int) {
	gdk.Guify("gdk_draw_line", drawable, gc, x1, y1, x2, y2)
}

// FUNCTION_NAME = gdk_draw_point, NONE, NONE, 4, WIDGET, WIDGET, INT, INT
func (gdk *Gdk) DrawPoint(drawable IDrawable, gc *GC, x, y int) {
	gdk.Guify("gdk_draw_point", drawable, gc, x, y)
}

// FUNCTION_NAME = gdk_draw_layout, NONE, NONE, 5, WIDGET, WIDGET, INT, INT, WIDGET
// FUNCTION_NAME = gdk_draw_layout_with_colors, NONE, NONE, 7, WIDGET, WIDGET, INT, INT, WIDGET, WIDGET, WIDGET
// FUNCTION_NAME = gdk_draw_drawable, NONE, NONE, 9, WIDGET, WIDGET, WIDGET, INT, INT, INT, INT, INT, INT
// FUNCTION_NAME = gdk_color_alloc, NONE, INT, 2, WIDGET, WIDGET
// FUNCTION_NAME = gdk_color_parse, NONE, NONE, 2, STRING, WIDGET
// FUNCTION_NAME = gdk_colormap_get_system, NONE, WIDGET, 0
// FUNCTION_NAME = gdk_colormap_alloc_color, NONE, BOOL, 4, WIDGET, WIDGET, BOOL, BOOL
// FUNCTION_NAME = gdk_get_default_root_window, NONE, WIDGET, 0
// FUNCTION_NAME = gdk_rgb_find_color, NONE, NONE, 2, WIDGET, WIDGET
// FUNCTION_NAME = gdk_drawable_set_colormap, NONE, NONE, 2, WIDGET, WIDGET
// FUNCTION_NAME = gdk_drawable_get_size, NONE, NONE, 3, WIDGET, WIDGET, WIDGET
// FUNCTION_NAME = gdk_drawable_get_depth, NONE, INT, 1, WIDGET
// FUNCTION_NAME = gdk_keymap_translate_keyboard_state, NONE, BOOL, 8, NULL, INT, INT, INT, WIDGET, NULL, NULL, NULL
// FUNCTION_NAME = gdk_screen_get_default, NONE, WIDGET, 0
// FUNCTION_NAME = gdk_screen_get_width, NONE, INT, 1, WIDGET
// FUNCTION_NAME = gdk_screen_get_height, NONE, INT, 1, WIDGET
// FUNCTION_NAME = gdk_screen_width, NONE, INT, 0
// FUNCTION_NAME = gdk_screen_height, NONE, INT, 0
// FUNCTION_NAME = gdk_flush, NONE, NONE, 0
// FUNCTION_NAME = gdk_init, NONE, NONE, 2, NULL, NULL
// FUNCTION_NAME = gdk_display_get_default, NONE, WIDGET, 0
// FUNCTION_NAME = gdk_display_get_pointer, NONE, NONE, 5, WIDGET, NULL, WIDGET, WIDGET, NULL
// FUNCTION_NAME = gdk_cursor_new, NONE, WIDGET, 1, INT
// FUNCTION_NAME = gdk_visual_get_best_depth, NONE, INT, 0
// FUNCTION_NAME = gdk_visual_get_system, NONE, WIDGET, 0
// FUNCTION_NAME = gdk_visual_get_screen, NONE, WIDGET, 1, WIDGET
// FUNCTION_NAME = gdk_threads_enter, NONE, NONE, 0
// FUNCTION_NAME = gdk_threads_leave, NONE, NONE, 0
// FUNCTION_NAME = gdk_cairo_create, NONE, WIDGET, 1, WIDGET
// FUNCTION_NAME = gdk_cairo_set_source_rgba, NONE, NONE, 2, WIDGET, WIDGET
// FUNCTION_NAME = gdk_rgba_parse, NONE, INT, 2, WIDGET, STRING
