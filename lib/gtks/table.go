package gtks

import sugar "github.com/jopbrown/gtk-sugar"

type Table struct {
	Container
}

func NewTable(candy sugar.Candy, id string) *Table {
	v := Table{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_table_new, NONE, WIDGET, 3, INT, INT, BOOL
func (gtk *Gtk) NewTable(rows, cols int, homogeneous bool) *Table {
	id := gtk.Guify("gtk_table_new", rows, cols, homogeneous).String()
	return NewTable(gtk, id)
}

// FUNCTION_NAME = gtk_table_attach_defaults, NONE, NONE, 6, WIDGET, WIDGET, INT, INT, INT, INT
func (table *Table) AttachDefaults(widget IWidget, left, right, top, bottom int) {
	table.Candy().Guify("gtk_table_attach_defaults", table, widget, left, right, top, bottom)
}
