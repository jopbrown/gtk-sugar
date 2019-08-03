package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Frame struct {
	Bin
}

func NewFrame(candy sugar.Candy, id string) *Frame {
	obj := Frame{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_frame_new, NONE, WIDGET, 1, STRING
func FrameNew(label string) *Frame {
	id := Candy().Guify("gtk_frame_new", label).String()
	return NewFrame(Candy(), id)
}

// FUNCTION_NAME = gtk_frame_set_label, NONE, NONE, 2, WIDGET, STRING
func (obj *Frame) SetLabel(label string) {
	obj.Candy().Guify("gtk_frame_set_label", obj, label)
}

// FUNCTION_NAME = gtk_frame_set_label_align, NONE, NONE, 3, WIDGET, FLOAT, FLOAT
func (obj *Frame) SetLabelAlign(xalign, yalign float32) {
	obj.Candy().Guify("gtk_frame_set_label_align", obj, xalign, yalign)
}

// FUNCTION_NAME = gtk_frame_set_label_widget , NONE, NONE, 2, WIDGET, WIDGET
func (obj *Frame) SetLabelWidget(labelWidget IWidget) {
	obj.Candy().Guify("gtk_frame_set_label_widget", obj, labelWidget)
}

// FUNCTION_NAME = gtk_frame_set_shadow_type , NONE, NONE, 2, WIDGET, INT
func (obj *Frame) SetShadowType(t ShadowType) {
	obj.Candy().Guify("gtk_frame_set_shadow_type", obj, t)
}

// FUNCTION_NAME = gtk_frame_get_label, NONE, STRING, 1, WIDGET
func (obj *Frame) GetLabel() string {
	return obj.Candy().Guify("gtk_frame_get_label", obj).String()
}

// FUNCTION_NAME = gtk_frame_get_label_widget, NONE, WIDGET, 1, WIDGET
func (obj *Frame) GetLabelWidget() *Widget {
	id := obj.Candy().Guify("gtk_frame_get_label_widget", obj).String()
	return NewWidget(obj.Candy(), id)
}

type AspectFrame struct {
	Frame
}

func NewAspectFrame(candy sugar.Candy, id string) *AspectFrame {
	obj := AspectFrame{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_aspect_frame_new, NONE, WIDGET, 5, STRING, FLOAT, FLOAT, FLOAT, BOOL
func AspectFrameNew(label string, xalign, yalign, ratio float32, obeyChild bool) *AspectFrame {
	id := Candy().Guify("gtk_aspect_frame_new", label, xalign, yalign, ratio, obeyChild).String()
	return NewAspectFrame(Candy(), id)
}

// FUNCTION_NAME = gtk_aspect_frame_set, NONE, NONE, 5, WIDGET, FLOAT, FLOAT, FLOAT, BOOL
func (obj *AspectFrame) Set(xalign, yalign, ratio float32, obeyChild bool) {
	obj.Candy().Guify("gtk_aspect_frame_set", obj, xalign, yalign, ratio, obeyChild)
}
