package gtks

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/gdks"
)

type Window struct {
	Bin
}

func NewWindow(candy sugar.Candy, id string) *Window {
	win := Window{}
	win.CandyWrapper = candy.NewWrapper(id)
	return &win
}

// FUNCTION_NAME = gtk_window_new, delete-event, WIDGET, 1, INT
func (gtk *Gtk) NewWindow(t WindowType) *Window {
	id := gtk.Guify("gtk_window_new", t).String()
	return NewWindow(gtk, id)
}

// FUNCTION_NAME = gtk_window_set_title, NONE, NONE, 2, WIDGET, STRING
func (win *Window) SetTitle(title string) {
	win.Candy().Guify("gtk_window_set_title", win, title)
}

// FUNCTION_NAME = gtk_window_get_title, NONE, STRING, 1, WIDGET
func (win *Window) GetTitle() string {
	return win.Candy().Guify("gtk_window_get_title", win).String()
}

// FUNCTION_NAME = gtk_window_set_default_size, NONE, NONE, 3, WIDGET, INT, INT
func (win *Window) SetDefaultSize(width, height int) {
	win.Candy().Guify("gtk_window_set_default_size", win, width, height)
}

// FUNCTION_NAME = gtk_window_set_position, NONE, NONE, 2, WIDGET, INT
func (win *Window) SetPosition(pos WindowPosition) {
	win.Candy().Guify("gtk_window_set_position", win, pos)
}

// FUNCTION_NAME = gtk_window_set_resizable, NONE, NONE, 2, WIDGET, BOOL
func (win *Window) SetResizable(resizable bool) {
	win.Candy().Guify("gtk_window_set_resizable", win, resizable)
}

// FUNCTION_NAME = gtk_window_set_transient_for, NONE, NONE, 2, WIDGET, WIDGET
func (win *Window) SetTransientFor(parent Window) {
	win.Candy().Guify("gtk_window_set_transient_for", win, parent)
}

// FUNCTION_NAME = gtk_window_set_modal, NONE, NONE, 2, WIDGET, BOOL
func (win *Window) SetModal(modal bool) {
	win.Candy().Guify("gtk_window_set_modal", win, modal)
}

// FUNCTION_NAME = gtk_window_maximize, NONE, NONE, 1, WIDGET
func (win *Window) Maximize() {
	win.Candy().Guify("gtk_window_maximize", win)
}

// FUNCTION_NAME = gtk_window_set_icon_from_file, NONE, BOOL, 3, WIDGET, STRING, NULL
func (win *Window) SetIconFromFile(fileName string) bool {
	res := win.Candy().Guify("gtk_window_set_icon_from_file", win, fileName, nil)
	return res.MustBool()
}

// FUNCTION_NAME = gtk_window_set_keep_above, NONE, NONE, 2, WIDGET, BOOL
func (win *Window) SetKeepAbove(setting bool) {
	win.Candy().Guify("gtk_window_set_keep_above", win, setting)
}

// FUNCTION_NAME = gtk_window_set_keep_below, NONE, NONE, 2, WIDGET, BOOL
func (win *Window) SetKeepBelow(setting bool) {
	win.Candy().Guify("gtk_window_set_keep_below", win, setting)
}

// FUNCTION_NAME = gtk_window_set_policy, NONE, NONE, 4, WIDGET, INT, INT, INT
func (win *Window) SetPolicy(allowShrink, allowGrow, autoShrink int) {
	win.Candy().Guify("gtk_window_set_policy", win, allowShrink, allowGrow, autoShrink)
}

// FUNCTION_NAME = gtk_window_fullscreen, NONE, NONE, 1, WIDGET
func (win *Window) Fullscreen() {
	win.Candy().Guify("gtk_window_fullscreen", win)
}

// FUNCTION_NAME = gtk_window_unfullscreen, NONE, NONE, 1, WIDGET
func (win *Window) Unfullscreen() {
	win.Candy().Guify("gtk_window_unfullscreen", win)
}

// FUNCTION_NAME = gtk_window_set_icon_name, NONE, NONE, 2, WIDGET, STRING
func (win *Window) SetIconName(name string) {
	win.Candy().Guify("gtk_window_set_icon_name", win, name)
}

// FUNCTION_NAME = gtk_window_add_accel_group, NONE, NONE, 2, WIDGET, WIDGET
func (win *Window) AddAccelGroup(accelGroup *AccelGroup) {
	win.Candy().Guify("gtk_window_add_accel_group", win, accelGroup)
}

// FUNCTION_NAME = gtk_window_set_type_hint, NONE, NONE, 2, WIDGET, INT
func (win *Window) SetTypeHint(hint gdks.WindowTypeHint) {
	win.Candy().Guify("gtk_window_set_type_hint", win, hint)
}

// FUNCTION_NAME = gtk_window_reshow_with_initial_size, NONE, NONE, 1, WIDGET
func (win *Window) ReshowWithInitialSize() {
	win.Candy().Guify("gtk_window_reshow_with_initial_size", win)
}

// FUNCTION_NAME = gtk_window_set_accept_focus, NONE, NONE, 2, WIDGET, BOOL
func (win *Window) SetAcceptFocus(setting bool) {
	win.Candy().Guify("gtk_window_set_accept_focus", win, setting)
}

// FUNCTION_NAME = gtk_window_present, NONE, NONE, 1, WIDGET
func (win *Window) WindowPresent() {
	win.Candy().Guify("gtk_window_present", win)
}

// FUNCTION_NAME = gtk_window_is_active, NONE, BOOL, 1, WIDGET
func (win *Window) IsActive() bool {
	return win.Candy().Guify("gtk_window_is_active", win).MustBool()
}
