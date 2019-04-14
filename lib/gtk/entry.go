package gtks

import sugar "github.com/jopbrown/gtk-sugar"

type Entry struct {
	Editable
}

func NewEntryFromCandy(candy sugar.Candy, id string) *Entry {
	v := Entry{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_entry_new, activate, WIDGET, 0
func NewEntry() *Entry {
	id := Candy().Guify("gtk_entry_new").String()
	return NewEntryFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_entry_get_text, NONE, STRING, 1, WIDGET
func (entry *Entry) GetText() string {
	return entry.Candy().Guify("gtk_entry_get_text", entry).String()
}

// FUNCTION_NAME = gtk_entry_set_text, NONE, NONE, 2, WIDGET, STRING
func (entry *Entry) SetText(text string) {
	entry.Candy().Guify("gtk_entry_set_text", entry, text)
}

// FUNCTION_NAME = gtk_entry_set_visibility, NONE, NONE, 2, WIDGET, BOOL
func (entry *Entry) SetVisibility(visible bool) {
	entry.Candy().Guify("gtk_entry_set_visibility", entry, visible)
}
