package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Editable struct {
	Widget
}

func NewEditable(candy sugar.Candy, id string) *Editable {
	v := Editable{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
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

// FUNCTION_NAME = gtk_editable_copy_clipboard, NONE, NONE, 1, WIDGET
func (e *Editable) CopyClipboard() {
	e.Candy().Guify("gtk_editable_copy_clipboard", e)
}

// FUNCTION_NAME = gtk_editable_cut_clipboard, NONE, NONE, 1, WIDGET
func (e *Editable) CutClipboard() {
	e.Candy().Guify("gtk_editable_cut_clipboard", e)
}

// FUNCTION_NAME = gtk_editable_paste_clipboard, NONE, NONE, 1, WIDGET
func (e *Editable) PasteClipboard() {
	e.Candy().Guify("gtk_editable_paste_clipboard", e)
}
