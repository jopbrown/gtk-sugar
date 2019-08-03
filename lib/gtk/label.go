package gtk

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/pango"
)

type Label struct {
	Misc
}

func NewLabel(candy sugar.Candy, id string) *Label {
	v := Label{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_label_new, NONE, WIDGET, 1, STRING
func LabelNew(str string) *Label {
	id := Candy().Guify("gtk_label_new", str).String()
	return NewLabel(Candy(), id)
}

// FUNCTION_NAME = gtk_label_set_text, NONE, NONE, 2, WIDGET, STRING
func (l *Label) SetText(str string) {
	l.Candy().Guify("gtk_label_set_text", l, str)
}

// FUNCTION_NAME = gtk_label_get_text, NONE, STRING, 1, WIDGET
func (l *Label) GetText() string {
	return l.Candy().Guify("gtk_label_get_text", l).String()
}

// FUNCTION_NAME = gtk_label_set_line_wrap, NONE, NONE, 2, WIDGET, BOOL
func (l *Label) SetLineWrap(wrap bool) {
	l.Candy().Guify("gtk_label_set_line_wrap", l, wrap)
}

// FUNCTION_NAME = gtk_label_set_selectable, NONE, NONE, 2, WIDGET, BOOL
func (l *Label) SetSelectable(setting bool) {
	l.Candy().Guify("gtk_label_set_selectable", l, setting)
}

// FUNCTION_NAME = gtk_label_set_use_markup, NONE, NONE, 2, WIDGET, BOOL
func (l *Label) SetUseMarkup(setting bool) {
	l.Candy().Guify("gtk_label_set_use_markup", l, setting)
}

// FUNCTION_NAME = gtk_label_set_justify, NONE, NONE, 2, WIDGET, INT
func (l *Label) SetJustify(jtype Justification) {
	l.Candy().Guify("gtk_label_set_justify", l, jtype)
}

// FUNCTION_NAME = gtk_label_get_width_chars, NONE, INT, 1, WIDGET
func (l *Label) GetWidthChars() int {
	return l.Candy().Guify("gtk_label_get_width_chars", l).MustInt()
}

// FUNCTION_NAME = gtk_label_get_max_width_chars, NONE, INT, 1, WIDGET
func (l *Label) GetMaxWidthChars() int {
	return l.Candy().Guify("gtk_label_get_max_width_chars", l).MustInt()
}

// FUNCTION_NAME = gtk_label_set_markup, NONE, NONE, 2, WIDGET, STRING
func (l *Label) SetMarkup(str string) {
	l.Candy().Guify("gtk_label_set_markup", l, str)
}

// FUNCTION_NAME = gtk_label_set_markup_with_mnemonic, NONE, NONE, 2, WIDGET, STRING
func (l *Label) SetMarkupWithMnemonic(str string) {
	l.Candy().Guify("gtk_label_set_markup_with_mnemonic", l, str)
}

// FUNCTION_NAME = gtk_label_set_ellipsize, NONE, NONE, 2, WIDGET, INT
func (l *Label) SetEllipsize(mode pango.EllipsizeMode) {
	l.Candy().Guify("gtk_label_set_ellipsize", l, mode)
}

// FUNCTION_NAME = gtk_label_set_max_width_chars, NONE, NONE, 2, WIDGET, INT
func (l *Label) SetMaxWidthChars(nChars int) {
	l.Candy().Guify("gtk_label_set_max_width_chars", l, nChars)
}
