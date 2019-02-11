package gtks

type Editable struct {
	Widget
}

// FUNCTION_NAME = gtk_editable_delete_text, NONE, NONE, 3, WIDGET, INT, INT
func (e *Editable) DeleteText(startPos, endPos int) {
	e.Candy().Guify("gtk_editable_delete_text", e, startPos, endPos)
}

// FUNCTION_NAME = gtk_editable_get_chars, NONE, STRING, 3, WIDGET, INT, INT
func (e *Editable) GetChars(startPos, endPos int) string {
	return e.Candy().Guify("gtk_editable_get_chars", e, startPos, endPos).String()
}

// FUNCTION_NAME = gtk_editable_set_editable, NONE, NONE, 2, WIDGET, BOOL
func (e *Editable) SetEditable(isEditable bool) {
	e.Candy().Guify("gtk_editable_set_editable", e, isEditable)
}

// FUNCTION_NAME = gtk_editable_select_region, NONE, NONE, 3, WIDGET, INT, INT
func (e *Editable) SelectRegion(startPos, endPos int) {
	e.Candy().Guify("gtk_editable_select_region", e, startPos, endPos)
}

// FUNCTION_NAME = gtk_editable_set_position, NONE, NONE, 2, WIDGET, INT
func (e *Editable) SetPosition(position int) {
	e.Candy().Guify("gtk_editable_set_position", e, position)
}

// FUNCTION_NAME = gtk_editable_get_position, NONE, INT, 1, WIDGET
func (e *Editable) GetPosition() int {
	return e.Candy().Guify("gtk_editable_get_position", e).MustInt()
}

// FUNCTION_NAME = gtk_editable_get_selection_bounds, NONE, BOOL, 3, WIDGET, PTR_INT, PTR_INT
func (e *Editable) GetSelectionBounds() (isSelected bool, startPos, endPos int) {
	fields := e.Candy().Guify("gtk_editable_get_selection_bounds", e, 0, 0).Fields()
	isSelected = fields[0].MustBool()
	startPos = fields[1].MustInt()
	endPos = fields[2].MustInt()
	return
}

// FUNCTION_NAME = gtk_editable_insert_text, NONE, NONE, 4, WIDGET, STRING, INT, PTR_INT
func (e *Editable) InsertText(text string, position int) int {
	return e.Candy().Guify("gtk_editable_insert_text", e, text, len(text), position).MustInt()
}

type Entry struct {
	Editable
}

// FUNCTION_NAME = gtk_entry_new, activate, WIDGET, 0
func (gtk *Gtk) NewEntry() *Entry {
	id := gtk.Guify("gtk_entry_new").String()
	entry := Entry{}
	entry.CandyWrapper = gtk.NewWrapper(id)
	return &entry
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
