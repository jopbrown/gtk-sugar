package gtks

type Table struct {
	Container
}

// FUNCTION_NAME = gtk_table_new, NONE, WIDGET, 3, INT, INT, BOOL
func (gtk *Gtk) NewTable(rows, cols int, homogeneous bool) *Table {
	id := gtk.Guify("gtk_table_new", rows, cols, homogeneous).String()
	table := Table{}
	table.CandyWrapper = gtk.NewWrapper(id)
	return &table
}

// FUNCTION_NAME = gtk_table_attach_defaults, NONE, NONE, 6, WIDGET, WIDGET, INT, INT, INT, INT
func (table *Table) AttachDefaults(widget IWidget, left, right, top, bottom int) {
	table.Candy().Guify("gtk_table_attach_defaults", table, widget, left, right, top, bottom)
}
