package gtk

import (
	"fmt"

	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/glib"
)

type ListStore struct {
	glib.Object
}

func NewListStore(candy sugar.Candy, id string) *ListStore {
	obj := ListStore{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_list_store_new, NONE, WIDGET, 2, INT, VARARGS
func ListStoreNew(types ...glib.Type) *ListStore {
	vargs := make(sugar.Varargs, 0, len(types))
	for _, t := range types {
		vargs = append(vargs, uint(t))
	}
	id := Candy().Guify("gtk_list_store_new", len(types), vargs).String()
	return NewListStore(Candy(), id)
}

// FUNCTION_NAME = gtk_list_store_append, NONE, NONE, 2, WIDGET, WIDGET
func (obj *ListStore) Append() *TreeIter {
	id := obj.Candy().ServerOpaque()
	obj.Candy().Guify("gtk_list_store_append", obj, id)
	iter := NewTreeIter(obj.Candy(), id)
	return iter
}

// FUNCTION_NAME = gtk_list_store_set, NONE, NONE, 3, WIDGET, WIDGET, VARARGS
func (obj *ListStore) Set(iter *TreeIter, cols []int, values []interface{}) {
	if len(cols) != len(values) {
		panic(fmt.Sprintf("cols(%d) and values(%d) length not match", len(cols), len(values)))
	}

	vargs := make(sugar.Varargs, 0, len(cols)+1)
	for i, col := range cols {
		vargs = append(vargs, col, values[i])
	}
	vargs = append(vargs, -1)
	obj.Candy().Guify("gtk_list_store_set", obj, iter, vargs)
}

// FUNCTION_NAME = gtk_list_store_set_value, NONE, NONE, 4, WIDGET, WIDGET, INT, STRING
func (obj *ListStore) SetValue(iter *TreeIter, col int, value interface{}) {
	obj.Candy().Guify("gtk_list_store_set_value", obj, iter, col, value)
}

// FUNCTION_NAME = gtk_list_store_clear, NONE, NONE, 1, WIDGET
func (obj *ListStore) Clear() {
	obj.Candy().Guify("gtk_list_store_clear", obj)
}

// FUNCTION_NAME = gtk_list_store_remove, NONE, BOOL, 2, WIDGET, WIDGET
func (obj *ListStore) Remove(iter *TreeIter) bool {
	return obj.Candy().Guify("gtk_list_store_remove", obj, iter).MustBool()
}

// FUNCTION_NAME = gtk_list_store_move_before, NONE, NONE, 3, WIDGET, WIDGET, WIDGET
func (obj *ListStore) Before(iter, position *TreeIter) {
	obj.Candy().Guify("gtk_list_store_move_before", obj, iter, position)
}

// FUNCTION_NAME = gtk_list_store_move_after, NONE, NONE, 3, WIDGET, WIDGET, WIDGET
func (obj *ListStore) After(iter, position *TreeIter) {
	obj.Candy().Guify("gtk_list_store_move_after", obj, iter, position)
}

// FUNCTION_NAME = gtk_list_store_insert, NONE, NONE, 3, WIDGET, WIDGET, INT
func (obj *ListStore) Insert(iter *TreeIter, position int) {
	obj.Candy().Guify("gtk_list_store_insert", obj, iter, position)
}

// FUNCTION_NAME = gtk_list_store_insert_with_values, NONE, NONE, 4, WIDGET, WIDGET, INT, VARARGS
func (obj *ListStore) InsertWithValues(iter *TreeIter, position int, cols []int, values []interface{}) {
	if len(cols) != len(values) {
		panic(fmt.Sprintf("cols(%d) and values(%d) length not match", len(cols), len(values)))
	}

	vargs := make(sugar.Varargs, 0, len(cols)+1)
	for i, col := range cols {
		vargs = append(vargs, col, values[i])
	}
	vargs = append(vargs, -1)
	obj.Candy().Guify("gtk_list_store_insert_with_values", obj, iter, position, vargs)
}
