package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Text struct {
	Editable
}

func NewTextFromCandy(candy sugar.Candy, id string) *Text {
	v := Text{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_text_new, NONE, WIDGET, 2, NULL, NULL
func TextNew() *Text {
	id := Candy().Guify("gtk_text_new", nil, nil).String()
	return NewTextFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_text_set_editable, NONE, NONE, 2, WIDGET, BOOL
func (w *Text) SetEditable(editable bool) {
	w.Candy().Guify("gtk_text_set_editable", w, editable)
}

// FUNCTION_NAME = gtk_text_insert, NONE, NONE, 6, WIDGET, NULL, NULL, NULL, STRING, INT
func (w *Text) Insert(text string) {
	w.Candy().Guify("gtk_text_insert", w, nil, nil, nil, text, len(text))
}

// FUNCTION_NAME = gtk_text_set_adjustments, NONE, NONE, 3, WIDGET, WIDGET, WIDGET
func (w *Text) SetAdjustments(h, v *Adjustment) {
	w.Candy().Guify("gtk_text_set_adjustments", w, h, v)
}

// FUNCTION_NAME = gtk_text_get_length, NONE, INT, 1, WIDGET
func (w *Text) GetLength() int {
	return w.Candy().Guify("gtk_text_get_length", w).MustInt()
}

// FUNCTION_NAME = gtk_text_set_word_wrap, NONE, NONE, 2, WIDGET, BOOL
func (w *Text) SetWordWrap(isWrap bool) {
	w.Candy().Guify("gtk_text_set_word_wrap", w, isWrap)
}

// FUNCTION_NAME = gtk_text_backward_delete, NONE, BOOL, 2, WIDGET, INT
func (w *Text) BackwardDelete(n int) {
	w.Candy().Guify("gtk_text_backward_delete", w, n)
}

// FUNCTION_NAME = gtk_text_forward_delete, NONE, BOOL, 2, WIDGET, INT
func (w *Text) ForwardDelete(n int) {
	w.Candy().Guify("gtk_text_forward_delete", w, n)
}

// FUNCTION_NAME = gtk_text_set_point, NONE, NONE, 2, WIDGET, INT
func (w *Text) SetPoint(index int) {
	w.Candy().Guify("gtk_text_set_point", w, index)
}
