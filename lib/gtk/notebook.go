package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Notebook struct {
	Container
}

func NewNotebook(candy sugar.Candy, id string) *Notebook {
	obj := Notebook{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_notebook_new, NONE, WIDGET, 0
func NotebookNew() *Notebook {
	id := Candy().Guify("gtk_notebook_new").String()
	return NewNotebook(Candy(), id)
}

// FUNCTION_NAME = gtk_notebook_set_tab_pos, NONE, NONE, 2, WIDGET, INT
func (obj *Notebook) SetTabPos(pos PositionType) {
	obj.Candy().Guify("gtk_notebook_set_tab_pos", obj, pos)
}

// FUNCTION_NAME = gtk_notebook_popup_enable, NONE, NONE, 1, WIDGET
func (obj *Notebook) PopupEnable() {
	obj.Candy().Guify("gtk_notebook_popup_enable", obj)
}

// FUNCTION_NAME = gtk_notebook_popup_disable, NONE, NONE, 1, WIDGET
func (obj *Notebook) PopupDisable() {
	obj.Candy().Guify("gtk_notebook_popup_disable", obj)
}

// FUNCTION_NAME = gtk_notebook_get_n_pages, NONE, INT, 1, WIDGET
func (obj *Notebook) GetNPages() int {
	return obj.Candy().Guify("gtk_notebook_get_n_pages", obj).MustInt()
}

// FUNCTION_NAME = gtk_notebook_get_nth_page, NONE, WIDGET, 2, WIDGET, INT
func (obj *Notebook) GetNthPage(pageNum int) *Widget {
	id := obj.Candy().Guify("gtk_notebook_get_nth_page", obj, pageNum).String()
	return NewWidget(obj.Candy(), id)
}

// FUNCTION_NAME = gtk_notebook_append_page, NONE, INT, 3, WIDGET, WIDGET, WIDGET
func (obj *Notebook) AppendPage(child, tabLabel IWidget) {
	obj.Candy().Guify("gtk_notebook_append_page", obj, child, tabLabel)
}

// FUNCTION_NAME = gtk_notebook_insert_page, NONE, INT, 4, WIDGET, WIDGET, WIDGET, INT
func (obj *Notebook) InsertPage(child, tabLabel IWidget, position int) {
	obj.Candy().Guify("gtk_notebook_insert_page", obj, child, tabLabel, position)
}

// FUNCTION_NAME = gtk_notebook_remove_page, NONE, NONE, 2, WIDGET, INT
func (obj *Notebook) RemovePage(pageNum int) {
	obj.Candy().Guify("gtk_notebook_remove_page", obj, pageNum)
}

// FUNCTION_NAME = gtk_notebook_get_current_page, NONE, INT, 1, WIDGET
func (obj *Notebook) GetCurrentPage() int {
	return obj.Candy().Guify("gtk_notebook_get_current_page", obj).MustInt()
}

// FUNCTION_NAME = gtk_notebook_get_menu_label, NONE, WIDGET, 2, WIDGET, WIDGET
func (obj *Notebook) GetMenuLabel() *Widget {
	id := obj.Candy().Guify("gtk_notebook_get_menu_label", obj).String()
	return NewWidget(obj.Candy(), id)
}

// FUNCTION_NAME = gtk_notebook_set_current_page, NONE, NONE, 2, WIDGET, INT
func (obj *Notebook) SetCurrentPage(pageNum int) {
	obj.Candy().Guify("gtk_notebook_set_current_page", obj, pageNum)
}

// FUNCTION_NAME = gtk_notebook_set_page, NONE, NONE, 2, WIDGET, INT
func (obj *Notebook) SetPage(pageNum int) {
	obj.Candy().Guify("gtk_notebook_set_page", obj, pageNum)
}

// FUNCTION_NAME = gtk_notebook_set_tab_label_text, NONE, NONE, 3, WIDGET, WIDGET, STRING
func (obj *Notebook) SetTabLabelText(child IWidget, tabText string) {
	obj.Candy().Guify("gtk_notebook_set_tab_label_text", obj, child, tabText)
}

// FUNCTION_NAME = gtk_notebook_set_scrollable, NONE, NONE, 2, WIDGET, BOOL
func (obj *Notebook) SetScrollable(scrollable bool) {
	obj.Candy().Guify("gtk_notebook_set_scrollable", obj, scrollable)
}

// FUNCTION_NAME = gtk_notebook_set_tab_reorderable, NONE, NONE, 3, WIDGET, WIDGET, BOOL
func (obj *Notebook) SetTabReorderable(child IWidget, reorderable bool) {
	obj.Candy().Guify("gtk_notebook_set_tab_reorderable", obj, child, reorderable)
}
