package gtk

import (
	"errors"
	"fmt"

	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/glib"
)

type TreeIter struct {
	sugar.CandyWrapper
}

func NewTreeIter(candy sugar.Candy, id string) *TreeIter {
	obj := TreeIter{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

type TreePath struct {
	sugar.CandyWrapper
}

func NewTreePath(candy sugar.Candy, id string) *TreePath {
	obj := TreePath{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_tree_path_new_from_string, NONE, WIDGET, 1, STRING
func TreePathFromString(path string) *TreePath {
	id := Candy().Guify("gtk_tree_path_new_from_string", path).String()
	return NewTreePath(Candy(), id)
}

// FUNCTION_NAME = gtk_tree_path_free, NONE, NONE, 1, WIDGET
func (obj *TreePath) Free() {
	obj.Candy().Guify("gtk_tree_path_free", obj)
}

// FUNCTION_NAME = gtk_tree_path_prev, NONE, BOOL, 1, WIDGET
func (obj *TreePath) Prev() bool {
	return obj.Candy().Guify("gtk_tree_path_prev", obj).MustBool()
}

// FUNCTION_NAME = gtk_tree_path_next, NONE, BOOL, 1, WIDGET
func (obj *TreePath) Next() bool {
	return obj.Candy().Guify("gtk_tree_path_next", obj).MustBool()
}

type TreeModel struct {
	sugar.CandyWrapper
}

func NewTreeModel(candy sugar.Candy, id string) *TreeModel {
	obj := TreeModel{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_tree_model_get_value, NONE, NONE, 4, WIDGET, WIDGET, INT, WIDGET
func (obj *TreeModel) GetValue(iter *TreeIter, column int) *glib.Value {
	id := obj.Candy().ServerOpaque()
	obj.Candy().Guify("gtk_tree_model_get_value", obj, iter, column, id)
	v := glib.NewValue(obj.Candy(), id)
	return v
}

// FUNCTION_NAME = gtk_tree_model_get_iter, NONE, BOOL, 3, WIDGET, WIDGET, WIDGET
func (obj *TreeModel) GetIter(path *TreePath) (*TreeIter, error) {
	id := obj.Candy().ServerOpaque()
	iter := NewTreeIter(obj.Candy(), id)
	ok := obj.Candy().Guify("gtk_tree_model_get_iter", obj, iter, path).MustBool()
	if ok {
		glib.Free(iter)
		return nil, errors.New("gtk_tree_model_get_iter fail to set iter")
	}
	return iter, nil
}

// FUNCTION_NAME = gtk_tree_model_get_iter_first, NONE, NONE, 2, WIDGET, WIDGET
func (obj *TreeModel) GetIterFirst() (*TreeIter, error) {
	id := obj.Candy().ServerOpaque()
	iter := NewTreeIter(obj.Candy(), id)
	ok := obj.Candy().Guify("gtk_tree_model_get_iter_first", obj, iter).MustBool()
	if ok {
		glib.Free(iter)
		return nil, errors.New("gtk_tree_model_get_iter_first fail to set iter")
	}
	return iter, nil
}

// FUNCTION_NAME = gtk_tree_model_iter_next, NONE, BOOL, 2, WIDGET, WIDGET
func (obj *TreeModel) IterNext(iter *TreeIter) bool {
	return obj.Candy().Guify("gtk_tree_model_iter_next", obj, iter).MustBool()
}

// FUNCTION_NAME = gtk_tree_model_get_string_from_iter, NONE, STRING, 2, WIDGET, WIDGET
func (obj *TreeModel) GetStringFromIter(iter *TreeIter) string {
	return obj.Candy().Guify("gtk_tree_model_get_string_from_iter", obj, iter).String()
}

type TreeView struct {
	Container
}

func NewTreeView(candy sugar.Candy, id string) *TreeView {
	obj := TreeView{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_tree_view_new, NONE, WIDGET, 0
func TreeViewNew() *TreeView {
	id := Candy().Guify("gtk_tree_view_new").String()
	return NewTreeView(Candy(), id)
}

// FUNCTION_NAME = gtk_tree_view_new_with_model, NONE, WIDGET, 1, WIDGET
func TreeViewNewWithModel(model *TreeModel) *TreeView {
	id := Candy().Guify("gtk_tree_view_new_with_model", model).String()
	return NewTreeView(Candy(), id)
}

// FUNCTION_NAME = gtk_tree_view_set_model, NONE, NONE, 2, WIDGET, WIDGET
func (obj *TreeView) SetModel(model *TreeModel) {
	obj.Candy().Guify("gtk_tree_view_set_model", obj, model)
}

// FUNCTION_NAME = gtk_tree_view_append_column, NONE, NONE, 2, WIDGET, WIDGET
func (obj *TreeView) AppendColumn(col *TreeViewColumn) {
	obj.Candy().Guify("gtk_tree_view_append_column", obj, col)
}

// FUNCTION_NAME = gtk_tree_view_set_headers_visible, NONE, NONE, 2, WIDGET, BOOL
func (obj *TreeView) SetHeadersVisible(headersVisible bool) {
	obj.Candy().Guify("gtk_tree_view_set_headers_visible", obj, headersVisible)
}

// FUNCTION_NAME = gtk_tree_view_set_headers_clickable, NONE, NONE, 2, WIDGET, BOOL
func (obj *TreeView) SetHeadersClickable(clickable bool) {
	obj.Candy().Guify("gtk_tree_view_set_headers_clickable", obj, clickable)
}

// FUNCTION_NAME = gtk_tree_view_get_selection, NONE, WIDGET, 1, WIDGET
func (obj *TreeView) GetSelection() *TreeSelection {
	id := obj.Candy().Guify("gtk_tree_view_get_selection", obj).String()
	return NewTreeSelection(obj.Candy(), id)
}

// FUNCTION_NAME = gtk_tree_view_get_hadjustment, NONE, WIDGET, 1, WIDGET
func (obj *TreeView) GetHAdjustment() *Adjustment {
	id := obj.Candy().Guify("gtk_tree_view_get_hadjustment", obj).String()
	return NewAdjustment(obj.Candy(), id)
}

// FUNCTION_NAME = gtk_tree_view_get_vadjustment, NONE, WIDGET, 1, WIDGET
func (obj *TreeView) GetVAdjustment() *Adjustment {
	id := obj.Candy().Guify("gtk_tree_view_get_vadjustment", obj).String()
	return NewAdjustment(obj.Candy(), id)
}

// FUNCTION_NAME = gtk_tree_view_scroll_to_cell, NONE, NONE, 6, WIDGET, WIDGET, WIDGET, BOOL, FLOAT, FLOAT
func (obj *TreeView) ScrollToCell(path *TreePath, col *TreeViewColumn, useAlign bool, rowAlign, colAlign float32) {
	obj.Candy().Guify("gtk_tree_view_scroll_to_cell", obj, path, col, useAlign, rowAlign, colAlign)
}

// FUNCTION_NAME = gtk_tree_view_set_grid_lines, NONE, NONE, 2, WIDGET, INT
func (obj *TreeView) SetGridLines(gridLines TreeViewGridLines) {
	obj.Candy().Guify("gtk_tree_view_set_grid_lines", obj, gridLines)
}

type TreeViewColumn struct {
	glib.Object
}

func NewTreeViewColumn(candy sugar.Candy, id string) *TreeViewColumn {
	obj := TreeViewColumn{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_tree_view_column_new, clicked, WIDGET, 0
// FUNCTION_NAME = gtk_tree_view_column_new_with_attributes, clicked, WIDGET, 5, STRING, WIDGET, STRING, INT, NULL
// FUNCTION_NAME = gtk_tree_view_column_pack_start, NONE, NONE, 3, WIDGET, WIDGET, BOOL
// FUNCTION_NAME = gtk_tree_view_column_set_title, NONE, NONE, 2, WIDGET, STRING
// FUNCTION_NAME = gtk_tree_view_column_set_resizable, NONE, NONE, 2, WIDGET, BOOL
func (obj *TreeViewColumn) SetResizable(resizable bool) {
	obj.Candy().Guify("gtk_tree_view_column_set_resizable", obj, resizable)
}

// FUNCTION_NAME = gtk_tree_view_column_set_clickable, NONE, NONE, 2, WIDGET, BOOL
func (obj *TreeViewColumn) SetClickable(clickable bool) {
	obj.Candy().Guify("gtk_tree_view_column_set_clickable", obj, clickable)
}

type TreeSelection struct {
	glib.Object
}

func NewTreeSelection(candy sugar.Candy, id string) *TreeSelection {
	obj := TreeSelection{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_tree_selection_get_selected, NONE, BOOL, 3, WIDGET, NULL, WIDGET
func (obj *TreeSelection) GetSelected() (model *TreeModel, iter *TreeIter, ok bool) {
	mid := obj.Candy().ServerOpaque()
	iid := obj.Candy().ServerOpaque()
	model = NewTreeModel(obj.Candy(), mid)
	iter = NewTreeIter(obj.Candy(), iid)
	ok = obj.Candy().Guify("gtk_tree_selection_get_selected", obj, model, iter).MustBool()
	return
}

// FUNCTION_NAME = gtk_tree_selection_get_tree_view, NONE, WIDGET, 1, WIDGET
func (obj *TreeSelection) GetTreeView() *TreeView {
	id := obj.Candy().Guify("gtk_tree_selection_get_tree_view", obj).String()
	return NewTreeView(obj.Candy(), id)
}

// FUNCTION_NAME = gtk_tree_selection_select_iter, NONE, NONE, 2, WIDGET, WIDGET
func (obj *TreeSelection) SelectIter(iter *TreeIter) {
	obj.Candy().Guify("gtk_tree_selection_select_iter", obj, iter)
}

// FUNCTION_NAME = gtk_tree_selection_select_path, NONE, NONE, 2, WIDGET, WIDGET
func (obj *TreeSelection) SelectPath(path *TreePath) {
	obj.Candy().Guify("gtk_tree_selection_select_path", obj, path)
}

// FUNCTION_NAME = gtk_tree_selection_path_is_selected, NONE, BOOL, 2, WIDGET, WIDGET
func (obj *TreeSelection) PathIsSelected(path *TreePath) bool {
	return obj.Candy().Guify("gtk_tree_selection_path_is_selected", obj, path).MustBool()
}

// FUNCTION_NAME = gtk_tree_selection_set_mode, NONE, NONE, 2, WIDGET, INT
func (obj *TreeSelection) SetMode(mode SelectionMode) {
	obj.Candy().Guify("gtk_tree_selection_set_mode", obj, mode)
}

// FUNCTION_NAME = gtk_tree_sortable_set_sort_column_id, NONE, NONE, 3, WIDGET, INT, INT
func (obj *TreeSelection) SetSortColumnID(sortColumnID int, order SortType) {
	obj.Candy().Guify("gtk_tree_sortable_set_sort_column_id", obj, sortColumnID, order)
}

type TreeStore struct {
	glib.Object
}

func NewTreeStore(candy sugar.Candy, id string) *TreeStore {
	obj := TreeStore{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_tree_store_new, NONE, WIDGET, 2, INT, VARARGS
func TreeStoreNew(types ...glib.Type) *TreeStore {
	vargs := make(sugar.Varargs, 0, len(types))
	for _, t := range types {
		vargs = append(vargs, uint(t))
	}
	id := Candy().Guify("gtk_tree_store_new", len(types), vargs).String()
	return NewTreeStore(Candy(), id)
}

// FUNCTION_NAME = gtk_tree_store_append, NONE, NONE, 3, WIDGET, WIDGET, WIDGET
func (obj *TreeStore) Append(parent *TreeIter) *TreeIter {
	id := obj.Candy().ServerOpaque()
	obj.Candy().Guify("gtk_tree_store_append", obj, id, parent)
	iter := NewTreeIter(obj.Candy(), id)
	return iter
}

// FUNCTION_NAME = gtk_tree_store_set, NONE, NONE, 3, WIDGET, WIDGET, VARARGS
func (obj *TreeStore) Set(iter *TreeIter, cols []int, values []interface{}) {
	if len(cols) != len(values) {
		panic(fmt.Sprintf("cols(%d) and values(%d) length not match", len(cols), len(values)))
	}

	vargs := make(sugar.Varargs, 0, len(cols)+1)
	for i, col := range cols {
		vargs = append(vargs, col, values[i])
	}
	vargs = append(vargs, -1)
	obj.Candy().Guify("gtk_tree_store_set", obj, iter, vargs)
}

type CellRenderer struct {
	glib.Object
}

func NewCellRenderer(candy sugar.Candy, id string) *CellRenderer {
	obj := CellRenderer{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_cell_renderer_text_new, NONE, WIDGET, 0
func CellRendererNew() *CellRenderer {
	id := Candy().Guify("gtk_cell_renderer_text_new").String()
	return NewCellRenderer(Candy(), id)
}
