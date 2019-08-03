package gtk

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/glib"
)

type TextMark struct {
	glib.Object
}

func NewTextMark(candy sugar.Candy, id string) *TextMark {
	v := TextMark{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

type TextBuffer struct {
	glib.Object
}

func NewTextBuffer(candy sugar.Candy, id string) *TextBuffer {
	v := TextBuffer{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_text_buffer_new, NONE, WIDGET, 1, NULL
func TextBufferNew() *TextBuffer {
	id := Candy().Guify("gtk_text_buffer_new").String()
	return NewTextBuffer(Candy(), id)
}

// FUNCTION_NAME = gtk_text_buffer_set_text, NONE, NONE, 3, WIDGET, STRING, INT
func (tb *TextBuffer) SetText(text string) {
	tb.Candy().Guify("gtk_text_buffer_set_text", tb, text, len(text))
}

// FUNCTION_NAME = gtk_text_buffer_get_insert, NONE, WIDGET, 1, WIDGET
func (tb *TextBuffer) GetInsert() *TextMark {
	id := tb.Candy().Guify("gtk_text_buffer_get_insert", tb).String()
	tm := TextMark{}
	tm.CandyWrapper = tb.Candy().NewWrapper(id)
	return &tm
}

// FUNCTION_NAME = gtk_text_buffer_get_start_iter, NONE, NONE, 2, WIDGET, WIDGET
func (tb *TextBuffer) GetStartIter() *TextIter {
	id := tb.Candy().ServerOpaque()
	tb.Candy().Guify("gtk_text_buffer_get_start_iter", tb, id)
	ti := TextIter{}
	ti.CandyWrapper = tb.Candy().NewWrapper(id)
	return &ti
}

// FUNCTION_NAME = gtk_text_buffer_get_end_iter, NONE, NONE, 2, WIDGET, WIDGET
func (tb *TextBuffer) GetEndIter() *TextIter {
	id := tb.Candy().ServerOpaque()
	tb.Candy().Guify("gtk_text_buffer_get_end_iter", tb, id)
	ti := TextIter{}
	ti.CandyWrapper = tb.Candy().NewWrapper(id)
	return &ti
}

// FUNCTION_NAME = gtk_text_buffer_get_bounds, NONE, NONE, 3, WIDGET, WIDGET, WIDGET
func (tb *TextBuffer) GetBounds() (start, end *TextIter) {
	sid := tb.Candy().ServerOpaque()
	eid := tb.Candy().ServerOpaque()
	tb.Candy().Guify("gtk_text_buffer_get_bounds", tb, sid, eid)
	start = NewTextIter(tb.Candy(), sid)
	end = NewTextIter(tb.Candy(), eid)
	return
}

// FUNCTION_NAME = gtk_text_buffer_get_selection_bounds, NONE, BOOL, 3, WIDGET, WIDGET, WIDGET
func (tb *TextBuffer) GetSelectionBounds() (start, end *TextIter) {
	sid := tb.Candy().ServerOpaque()
	eid := tb.Candy().ServerOpaque()
	tb.Candy().Guify("gtk_text_buffer_get_selection_bounds", tb, sid, eid)
	start = NewTextIter(tb.Candy(), sid)
	end = NewTextIter(tb.Candy(), eid)
	return
}

// FUNCTION_NAME = gtk_text_buffer_get_iter_at_offset, NONE, NONE, 3, WIDGET, WIDGET, INT
func (tb *TextBuffer) GetIterAtOffset(offset int) *TextIter {
	id := tb.Candy().ServerOpaque()
	tb.Candy().Guify("gtk_text_buffer_get_iter_at_offset", tb, id, offset)
	ti := TextIter{}
	ti.CandyWrapper = tb.Candy().NewWrapper(id)
	return &ti
}

// FUNCTION_NAME = gtk_text_buffer_get_text, NONE, STRING, 4, WIDGET, WIDGET, WIDGET, BOOL
func (tb *TextBuffer) GetText(start, end *TextIter, includeHiddenChars bool) string {
	return tb.Candy().Guify("gtk_text_buffer_get_text", tb, start, end, includeHiddenChars).MustUnquote()
}

// FUNCTION_NAME = gtk_text_buffer_create_tag, NONE, WIDGET, 4, WIDGET, STRING, STRING, VARARGS
func (tb *TextBuffer) CreateTag(tagName, firstPropertyName string, firstPropertyValue interface{}, namesAndValues ...interface{}) *TextTag {
	args := make(sugar.Varargs, 0, len(namesAndValues)+1)
	args = append(args, firstPropertyValue)
	args = append(args, namesAndValues...)
	id := tb.Candy().Guify("gtk_text_buffer_create_tag", tb, tagName, firstPropertyName, args, nil).String()
	w := TextTag{}
	w.CandyWrapper = tb.Candy().NewWrapper(id)
	return &w
}

// FUNCTION_NAME = gtk_text_buffer_insert, NONE, NONE, 4, WIDGET, WIDGET, STRING, INT
func (tb *TextBuffer) Insert(text string) {
	tb.Candy().Guify("gtk_text_buffer_insert", tb, text, len(text))
}

// FUNCTION_NAME = gtk_text_buffer_insert_at_cursor, NONE, NONE, 3, WIDGET, STRING, INT
func (tb *TextBuffer) InsertAtCursor(text string) {
	tb.Candy().Guify("gtk_text_buffer_insert_at_cursor", tb, text, len(text))
}

// FUNCTION_NAME = gtk_text_buffer_insert_with_tags_by_name, NONE, NONE, 5, WIDGET, WIDGET, STRING, INT, VARARGS
func (tb *TextBuffer) InsertWithTagsByName(iter *TextIter, text string, length int, firstTagName string, moreTagNames ...string) {
	args := make(sugar.Varargs, len(moreTagNames))
	for i, n := range moreTagNames {
		args[i] = n
	}
	tb.Candy().Guify("gtk_text_buffer_insert_with_tags_by_name", tb, iter, text, length, args, nil)
}

// FUNCTION_NAME = gtk_text_buffer_apply_tag_by_name, NONE, NONE, 4, WIDGET, STRING, WIDGET, WIDGET
func (tb *TextBuffer) ApplyTagByName(name string, start, end *TextIter) {
	tb.Candy().Guify("gtk_text_buffer_apply_tag_by_name", tb, name, start, end)
}

// FUNCTION_NAME = gtk_text_buffer_remove_tag_by_name, NONE, NONE, 4, WIDGET, STRING, WIDGET, WIDGET
func (tb *TextBuffer) RemoveTagByName(name string, start, end *TextIter) {
	tb.Candy().Guify("gtk_text_buffer_remove_tag_by_name", tb, name, start, end)
}

// FUNCTION_NAME = gtk_text_buffer_remove_all_tags, NONE, NONE, 3, WIDGET, WIDGET, WIDGET
func (tb *TextBuffer) RemoveAllTags(start, end *TextIter) {
	tb.Candy().Guify("gtk_text_buffer_remove_all_tags", tb, start, end)
}

// FUNCTION_NAME = gtk_text_buffer_get_tag_table, NONE, WIDGET, 1, WIDGET
func (tb *TextBuffer) GetTagTable() *TextTagTable {
	id := tb.Candy().Guify("gtk_text_buffer_get_tag_table", tb).String()
	w := TextTagTable{}
	w.CandyWrapper = tb.Candy().NewWrapper(id)
	return &w
}

func (tb *TextBuffer) SelectRange(ins, bound *TextIter) {
	tb.Candy().Guify("gtk_text_buffer_select_range", tb, ins, bound)
}

// FUNCTION_NAME = gtk_text_buffer_select_range, NONE, NONE, 3, WIDGET, WIDGET, WIDGET
func (tb *TextBuffer) GetSelectionBound() *TextMark {
	id := tb.Candy().Guify("gtk_text_buffer_get_selection_bound", tb).String()
	tm := TextMark{}
	tm.CandyWrapper = tb.Candy().NewWrapper(id)
	return &tm
}

// FUNCTION_NAME = gtk_text_buffer_get_line_count, NONE, INT, 1, WIDGET
func (tb *TextBuffer) GetLineCount() int {
	return tb.Candy().Guify("gtk_text_buffer_get_line_count", tb).MustInt()
}

// FUNCTION_NAME = gtk_text_buffer_create_mark, NONE, WIDGET, 4, WIDGET, STRING, WIDGET, BOOL
func (tb *TextBuffer) CreateMark(markName string, where *TextIter, leftGravity bool) *TextMark {
	id := tb.Candy().Guify("gtk_text_buffer_create_mark", tb, markName, where, leftGravity).String()
	tm := TextMark{}
	tm.CandyWrapper = tb.Candy().NewWrapper(id)
	return &tm
}

// FUNCTION_NAME = gtk_text_buffer_get_iter_at_mark, NONE, NONE, 3, WIDGET, WIDGET, WIDGET
func (tb *TextBuffer) GetIterAtMark(mark *TextMark) *TextIter {
	id := tb.Candy().ServerOpaque()
	tb.Candy().Guify("gtk_text_buffer_get_iter_at_mark", tb, id, mark)
	ti := TextIter{}
	ti.CandyWrapper = tb.Candy().NewWrapper(id)
	return &ti
}

// FUNCTION_NAME = gtk_text_buffer_get_iter_at_line, NONE, NONE, 3, WIDGET, WIDGET, INT
func (tb *TextBuffer) GetIterAtLine(lineNum int) *TextIter {
	id := tb.Candy().ServerOpaque()
	tb.Candy().Guify("gtk_text_buffer_get_iter_at_line", tb, id, lineNum)
	ti := TextIter{}
	ti.CandyWrapper = tb.Candy().NewWrapper(id)
	return &ti
}

// FUNCTION_NAME = gtk_text_buffer_delete, NONE, NONE, 3, WIDGET, WIDGET, WIDGET
func (tb *TextBuffer) Delete(start, end *TextIter) {
	tb.Candy().Guify("gtk_text_buffer_delete", tb, start, end)
}

// FUNCTION_NAME = gtk_text_buffer_delete_mark, NONE, NONE, 2, WIDGET, WIDGET
func (tb *TextBuffer) DeleteMark(mark *TextMark) {
	tb.Candy().Guify("gtk_text_buffer_delete_mark", tb, mark)
}

// FUNCTION_NAME = gtk_text_buffer_delete_mark_by_name, NONE, NONE, 2, WIDGET, STRING
func (tb *TextBuffer) DeleteMarkByName(name string) {
	tb.Candy().Guify("gtk_text_buffer_delete_mark_by_name", tb, name)
}

// FUNCTION_NAME = gtk_text_buffer_place_cursor, NONE, NONE, 2, WIDGET, WIDGET
func (tb *TextBuffer) PlaceCursor(where *TextIter) {
	tb.Candy().Guify("gtk_text_buffer_place_cursor ", tb, where)
}

// FUNCTION_NAME = gtk_text_buffer_copy_clipboard, NONE, NONE, 2, WIDGET, WIDGET
func (tb *TextBuffer) CopyClipboard(clipboard *Clipboard) {
	tb.Candy().Guify("gtk_text_buffer_copy_clipboard ", tb, clipboard)
}

// FUNCTION_NAME = gtk_text_buffer_cut_clipboard, NONE, NONE, 3, WIDGET, WIDGET, BOOL
func (tb *TextBuffer) CutClipboard(clipboard *Clipboard, defaultEditable bool) {
	tb.Candy().Guify("gtk_text_buffer_cut_clipboard ", tb, clipboard, defaultEditable)
}

// FUNCTION_NAME = gtk_text_buffer_paste_clipboard, NONE, NONE, 4, WIDGET, WIDGET, NULL, BOOL
func (tb *TextBuffer) PasteClipboard(clipboard *Clipboard, overrideLocation *TextIter, defaultEditable bool) {
	tb.Candy().Guify("gtk_text_buffer_paste_clipboard ", tb, clipboard, overrideLocation, defaultEditable)
}

type TextView struct {
	Container
}

func NewTextView(candy sugar.Candy, id string) *TextView {
	v := TextView{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_text_view_new, NONE, WIDGET, 0
func TextViewNew() *TextView {
	id := Candy().Guify("gtk_text_view_new").String()
	return NewTextView(Candy(), id)
}

// FUNCTION_NAME = gtk_text_view_new_with_buffer, NONE, WIDGET, 1, WIDGET
func TextViewNewWithBuffer(tb *TextBuffer) *TextView {
	id := Candy().Guify("gtk_text_view_new_with_buffer").String()
	return NewTextView(Candy(), id)
}

// FUNCTION_NAME = gtk_text_view_set_wrap_mode, NONE, NONE, 2, WIDGET, INT
func (w *TextView) SetWrapMode(wrapMode WrapMode) {
	w.Candy().Guify("gtk_text_view_set_wrap_mode", w, wrapMode)
}

// FUNCTION_NAME = gtk_text_view_set_editable, NONE, NONE, 2, WIDGET, BOOL
func (w *TextView) SetEditable(setting bool) {
	w.Candy().Guify("gtk_text_view_set_editable", w, setting)
}

// FUNCTION_NAME = gtk_text_view_set_border_window_size, NONE, NONE, 3, WIDGET, INT, INT
func (w *TextView) SetBorderWindowSize(t TextWindowType, size int) {
	w.Candy().Guify("gtk_text_view_set_border_window_size", w, t, size)
}

// FUNCTION_NAME = gtk_text_view_move_mark_onscreen, NONE, BOOL, 2, WIDGET, WIDGET
func (w *TextView) MoveMarkOnscreen(mark *TextMark) bool {
	return w.Candy().Guify("gtk_text_view_move_mark_onscreen", w, mark).MustBool()
}

// FUNCTION_NAME = gtk_text_view_scroll_to_mark, NONE, NONE, 6, WIDGET, WIDGET, DOUBLE, BOOL, DOUBLE, DOUBLE
func (w *TextView) ScrollToMark(mark *TextMark, withinMargin float64, useAlign bool, xAlign, yAlign float64) {
	w.Candy().Guify("gtk_text_view_scroll_to_mark", w, mark, withinMargin, useAlign, xAlign, yAlign)
}

// FUNCTION_NAME = gtk_text_view_scroll_mark_onscreen, NONE, NONE, 2, WIDGET, WIDGET
func (w *TextView) ScrollMarkOnscreen(mark *TextMark) {
	w.Candy().Guify("gtk_text_view_scroll_mark_onscreen", w, mark)
}

// FUNCTION_NAME = gtk_text_view_set_pixels_inside_wrap, NONE, NONE, 2, WIDGET, INT
func (w *TextView) SetPixelsInsideWrap(pixelsInsideWrap int) {
	w.Candy().Guify("gtk_text_view_set_pixels_inside_wrap", w, pixelsInsideWrap)
}

// FUNCTION_NAME = gtk_text_view_get_pixels_inside_wrap, NONE, INT, 1, WIDGET
func (w *TextView) GetPixelsInsideWrap() int {
	return w.Candy().Guify("gtk_text_view_get_pixels_inside_wrap", w).MustInt()
}

// FUNCTION_NAME = gtk_text_view_set_pixels_above_lines, NONE, NONE, 2, WIDGET, INT
func (w *TextView) SetPixelsAboveLines(pixelsAboveLines int) {
	w.Candy().Guify("gtk_text_view_set_pixels_above_lines", w, pixelsAboveLines)
}

// FUNCTION_NAME = gtk_text_view_get_pixels_above_lines, NONE, INT, 1, WIDGET
func (w *TextView) GetPixelsAboveLines() int {
	return w.Candy().Guify("gtk_text_view_get_pixels_above_lines", w).MustInt()
}

// FUNCTION_NAME = gtk_text_view_set_cursor_visible, NONE, NONE, 2, WIDGET, BOOL
func (w *TextView) SetCursorVisible(setting bool) {
	w.Candy().Guify("gtk_text_view_set_cursor_visible", w, setting)
}

// FUNCTION_NAME = gtk_text_view_window_to_buffer_coords, NONE, NONE, 6, WIDGET, INT, INT, INT, PTR_INT, PTR_INT
func (w *TextView) WindowToBufferCoords(winType TextWindowType, windowX, windowY int) (bufferX, bufferY int) {
	fields := w.Candy().Guify("gtk_text_view_window_to_buffer_coords", w, winType, windowX, windowY, bufferX, bufferY).Fields()
	bufferX = fields[0].MustInt()
	bufferY = fields[1].MustInt()
	return
}

// FUNCTION_NAME = gtk_text_view_get_buffer, NONE, WIDGET, 1, WIDGET
func (w *TextView) GetBuffer() *TextBuffer {
	id := w.Candy().Guify("gtk_text_view_get_buffer", w).String()
	tb := TextBuffer{}
	tb.CandyWrapper = w.Candy().NewWrapper(id)
	return &tb
}

type TextIter struct {
	sugar.CandyWrapper
}

func NewTextIter(candy sugar.Candy, id string) *TextIter {
	v := TextIter{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_text_iter_forward_search, NONE, BOOL, 6, WIDGET, STRING, INT, WIDGET, WIDGET, WIDGET
func (w *TextIter) ForwardSearch(str string, flags TextSearchFlags, start, end, limit *TextIter) bool {
	return w.Candy().Guify("gtk_text_iter_forward_search", w, str, flags, start, end, limit).MustBool()
}

// FUNCTION_NAME = gtk_text_iter_forward_visible_cursor_position, NONE, BOOL, 1, WIDGET
func (w *TextIter) ForwardVisibleCursorPosition() bool {
	return w.Candy().Guify("gtk_text_iter_forward_visible_cursor_position", w).MustBool()
}

// FUNCTION_NAME = gtk_text_iter_forward_to_line_end, NONE, BOOL, 1, WIDGET
func (w *TextIter) ForwardToLineEnd() bool {
	return w.Candy().Guify("gtk_text_iter_forward_to_line_end", w).MustBool()
}

// FUNCTION_NAME = gtk_text_iter_set_line, NONE, NONE, 2, WIDGET, INT
func (w *TextIter) SetLine(lineNumber int) {
	w.Candy().Guify("gtk_text_iter_set_line", w, lineNumber)
}

// FUNCTION_NAME = gtk_text_iter_set_line_offset, NONE, NONE, 2, WIDGET, INT
func (w *TextIter) SetLineOffset(charOnLine int) {
	w.Candy().Guify("gtk_text_iter_set_line_offset", w, charOnLine)
}

// FUNCTION_NAME = gtk_text_iter_set_line_index, NONE, NONE, 2, WIDGET, INT
func (w *TextIter) SetLineIndex(byteOnLine int) {
	w.Candy().Guify("gtk_text_iter_set_line_index", w, byteOnLine)
}

// FUNCTION_NAME = gtk_text_iter_get_line, NONE, INT, 1, WIDGET
func (w *TextIter) GetLine() int {
	return w.Candy().Guify("gtk_text_iter_get_line", w).MustInt()
}

// FUNCTION_NAME = gtk_text_iter_get_text, NONE, STRING, 2, WIDGET, WIDGET
func TextIterGetText(start, end *TextIter) string {
	return Candy().Guify("gtk_text_iter_get_text", start, end).String()
}

type TextTag struct {
	glib.Object
}

func NewTextTag(candy sugar.Candy, id string) *TextTag {
	v := TextTag{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

type TextTagTable struct {
	glib.Object
}

func NewTextTagTable(candy sugar.Candy, id string) *TextTagTable {
	v := TextTagTable{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_text_tag_table_add, NONE, BOOL, 2, WIDGET, WIDGET
func (w *TextTagTable) Add(tag *TextTag) bool {
	return w.Candy().Guify("gtk_text_tag_table_add", w, tag).MustBool()
}

// FUNCTION_NAME = gtk_text_tag_table_lookup, NONE, WIDGET, 2, WIDGET, STRING
func (w *TextTagTable) Lookup(name string) *TextTag {
	id := w.Candy().Guify("gtk_text_tag_table_lookup", w, name).String()
	tag := TextTag{}
	tag.CandyWrapper = w.Candy().NewWrapper(id)
	return &tag
}

// FUNCTION_NAME = gtk_text_tag_table_remove, NONE, NONE, 2, WIDGET, WIDGET
func (w *TextTagTable) Remove(tag *TextTag) {
	w.Candy().Guify("gtk_text_tag_table_remove", w, tag)
}
