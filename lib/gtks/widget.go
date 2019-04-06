package gtks

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/gdks"
	"github.com/jopbrown/gtk-sugar/lib/glibs"
	"github.com/jopbrown/gtk-sugar/lib/pangos"
)

type IWidget interface {
	sugar.CandyWrapper
	ShowAll()
	Show()
	Hide()
}

type Widget struct {
	glibs.Object
}

func NewWidget(candy sugar.Candy, id string) *Widget {
	w := Widget{}
	w.CandyWrapper = candy.NewWrapper(id)
	return &w
}

// FUNCTION_NAME = gtk_widget_show_all, NONE, NONE, 1, WIDGET
func (w *Widget) ShowAll() {
	w.Candy().Guify("gtk_widget_show_all", w)
}

// FUNCTION_NAME = gtk_widget_show, NONE, NONE, 1, WIDGET
func (w *Widget) Show() {
	w.Candy().Guify("gtk_widget_show", w)
}

// FUNCTION_NAME = gtk_widget_hide, NONE, NONE, 1, WIDGET
func (w *Widget) Hide() {
	w.Candy().Guify("gtk_widget_hide", w)
}

// FUNCTION_NAME = gtk_widget_realize, NONE, NONE, 1, WIDGET
func (w *Widget) Realize() {
	w.Candy().Guify("gtk_widget_realize", w)
}

// FUNCTION_NAME = gtk_widget_unrealize, NONE, NONE, 1, WIDGET
func (w *Widget) Unrealize() {
	w.Candy().Guify("gtk_widget_unrealize", w)
}

// FUNCTION_NAME = gtk_widget_destroy, NONE, NONE, 1, WIDGET
func (w *Widget) Destroy() {
	w.Candy().Guify("gtk_widget_destroy", w)
}

// FUNCTION_NAME = gtk_widget_grab_focus, NONE, NONE, 1, WIDGET
func (w *Widget) GrabFocus() {
	w.Candy().Guify("gtk_widget_grab_focus", w)
}

// FUNCTION_NAME = gtk_widget_set_size_request, NONE, NONE, 3, WIDGET, INT, INT
func (w *Widget) SetSizeRequest(width, height int) {
	w.Candy().Guify("gtk_widget_set_size_request", w, width, height)
}

type Requisition struct {
	Width, Height int
}

// FUNCTION_NAME = gtk_widget_size_request, NONE, NONE, 2, WIDGET, PTR_BASE64
func (w *Widget) SizeRequest() *Requisition {
	r := Requisition{}
	packer := sugar.NewBase64Packer(&r)
	format := packer.Format()
	w.Candy().ServerDataFormat(format)
	base64Data := w.Candy().Guify("gtk_widget_size_request", w, 0).String()
	fields := w.Candy().ServerUnpack(format, base64Data)
	fields.MustUnmarshal(packer)
	return &r
}

// FUNCTION_NAME = gtk_widget_set_sensitive, NONE, NONE, 2, WIDGET, BOOL
func (w *Widget) SetSensitive(sensitive bool) {
	w.Candy().Guify("gtk_widget_set_sensitive", w, sensitive)
}

// FUNCTION_NAME = gtk_widget_add_accelerator, NONE, NONE, 6, WIDGET, STRING, WIDGET, INT, INT, INT
func (w *Widget) AddAccelerator(accelSignal string, accelGroup *AccelGroup, accelKey uint, accelMods gdks.ModifierType, accelFlags AccelFlags) {
	w.Candy().Guify("gtk_widget_add_accelerator", w, accelSignal, accelGroup, accelKey, accelMods, accelFlags)
}

// FUNCTION_NAME = gtk_widget_get_parent, NONE, WIDGET, 1, WIDGET
func (w *Widget) GetParent() *Widget {
	id := w.Candy().Guify("gtk_widget_get_parent", w).String()
	return NewWidget(w.Candy(), id)
}

// FUNCTION_NAME = gtk_widget_get_toplevel, NONE, WIDGET, 1, WIDGET
func (w *Widget) GetToplevel() *Widget {
	id := w.Candy().Guify("gtk_widget_get_toplevel", w).String()
	return NewWidget(w.Candy(), id)
}

// FUNCTION_NAME = gtk_widget_set_name, NONE, NONE, 2, WIDGET, STRING
func (w *Widget) SetName(name string) {
	w.Candy().Guify("gtk_widget_set_name", w, name)
}

// FUNCTION_NAME = gtk_widget_get_size_request, NONE, NONE, 3, WIDGET, PTR_INT, PTR_INT
func (w *Widget) GetSizeRequest(name string) (width, height int) {
	fields := w.Candy().Guify("gtk_widget_get_size_request", w, 0, 0).Fields()
	width = fields[0].MustInt()
	height = fields[1].MustInt()
	return
}

// FUNCTION_NAME = gtk_widget_add_events, NONE, NONE, 2, WIDGET, INT
func (w *Widget) AddEvents(events int) {
	w.Candy().Guify("gtk_widget_add_events", w, events)
}

// FUNCTION_NAME = gtk_widget_queue_draw, NONE, NONE, 1, WIDGET
func (w *Widget) QueueDraw() {
	w.Candy().Guify("gtk_widget_queue_draw", w)
}

// FUNCTION_NAME = gtk_widget_get_parent_window, NONE, WIDGET, 1, WIDGET
func (w *Widget) GetParentWindow() *Window {
	id := w.Candy().Guify("gtk_widget_get_parent_window", w).String()
	return NewWindow(w.Candy(), id)
}

// FUNCTION_NAME = gtk_widget_create_pango_layout, NONE, WIDGET, 2, WIDGET, STRING
func (w *Widget) CreatePangoLayout() *pangos.Layout {
	id := w.Candy().Guify("gtk_widget_create_pango_layout", w).String()
	return pangos.NewLayout(w.Candy(), id)
}

// FUNCTION_NAME = gtk_widget_get_window, NONE, WIDGET, 1, WIDGET
func (w *Widget) GetWindow() *Window {
	id := w.Candy().Guify("gtk_widget_get_window", w).String()
	return NewWindow(w.Candy(), id)
}

// FUNCTION_NAME = gtk_widget_set_tooltip_text, NONE, NONE, 2, WIDGET, STRING
func (w *Widget) SetTooltipText(text string) {
	w.Candy().Guify("gtk_widget_set_tooltip_text", w, text)
}
