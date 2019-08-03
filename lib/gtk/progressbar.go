package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type ProgressBar struct {
	Widget
}

func NewProgressBar(candy sugar.Candy, id string) *ProgressBar {
	obj := ProgressBar{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_progress_bar_new, NONE, WIDGET, 0
func ProgressBarNew() *ProgressBar {
	id := Candy().Guify("gtk_progress_bar_new").String()
	return NewProgressBar(Candy(), id)
}

// FUNCTION_NAME = gtk_progress_bar_set_text, NONE, NONE, 2, WIDGET, STRING
func (obj *ProgressBar) SetText(text string) {
	obj.Candy().Guify("gtk_progress_bar_set_text", obj, text)
}

// FUNCTION_NAME = gtk_progress_bar_set_fraction, NONE, NONE, 2, WIDGET, DOUBLE
func (obj *ProgressBar) SetFraction(fraction float64) {
	obj.Candy().Guify("gtk_progress_bar_set_fraction", obj, fraction)
}

// FUNCTION_NAME = gtk_progress_bar_pulse, NONE, NONE, 1, WIDGET
func (obj *ProgressBar) Pulse() {
	obj.Candy().Guify("gtk_progress_bar_pulse", obj)
}

// FUNCTION_NAME = gtk_progress_bar_set_pulse_step, NONE, NONE, 2, WIDGET, DOUBLE
func (obj *ProgressBar) SetPulseStep(fraction float64) {
	obj.Candy().Guify("gtk_progress_bar_set_pulse_step", obj, fraction)
}
