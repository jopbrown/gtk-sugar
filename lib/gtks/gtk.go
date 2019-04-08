package gtks

import (
	sugar "github.com/jopbrown/gtk-sugar"
)

type Gtk struct {
	sugar.Candy
}

func NewGtk(candy sugar.Candy) *Gtk {
	return &Gtk{Candy: candy}
}

// FUNCTION_NAME = gtk_init, NONE, NONE, 2, NULL, NULL
func (gtk *Gtk) Init() {
	gtk.Guify("gtk_init", nil, nil)
}

// FUNCTION_NAME = gtk_settings_get_default, NONE, WIDGET, 0
// FUNCTION_NAME = gtk_misc_set_alignment, NONE, NONE, 3, WIDGET, FLOAT, FLOAT
// FUNCTION_NAME = gtk_main, NONE, NONE, 0
// FUNCTION_NAME = gtk_main_iteration, NONE, BOOL, 0
// FUNCTION_NAME = gtk_main_iteration_do, NONE, BOOL, 1, BOOL
// FUNCTION_NAME = gtk_events_pending, NONE, BOOL, 0
// FUNCTION_NAME = gtk_exit, NONE, NONE, 1, INT
// FUNCTION_NAME = gtk_main_quit, NONE, NONE, 0

// FUNCTION_NAME = gtk_check_version, NONE, STRING, 3, INT, INT, INT
// FUNCTION_NAME = gtk_drag_source_set, NONE, NONE, 5, WIDGET, INT, WIDGET, INT, INT
// FUNCTION_NAME = gtk_drag_dest_set, NONE, NONE, 5, WIDGET, INT, WIDGET, INT, INT
// FUNCTION_NAME = gtk_drag_finish, NONE, NONE, 4, WIDGET, BOOL, BOOL, INT
// FUNCTION_NAME = gtk_drag_set_icon_stock, NONE, NONE, 4, WIDGET, STRING, INT, INT
// FUNCTION_NAME = gtk_get_current_event_time, NONE, INT, 0
// FUNCTION_NAME = gtk_signal_emit_by_name, NONE, NONE, 2, WIDGET, STRING
// FUNCTION_NAME = gtk_signal_connect_while_alive, NONE, NONE, 5, WIDGET, STRING, MACRO, POINTER, WIDGET
// FUNCTION_NAME = gtk_invisible_new, NONE, WIDGET, 0
// FUNCTION_NAME = gtk_target_list_new, NONE, WIDGET, 2, NULL, INT
// FUNCTION_NAME = gtk_target_list_add, NONE, NONE, 4, WIDGET, WIDGET, INT, INT
// FUNCTION_NAME = gtk_target_table_new_from_list, NONE, WIDGET, 2, WIDGET, WIDGET
// FUNCTION_NAME = gtk_object_set_data, NONE, NONE, 3, WIDGET, STRING, INT
// FUNCTION_NAME = gtk_object_get_data, NONE, INT, 2, WIDGET, STRING
// FUNCTION_NAME = gtk_object_ref, NONE, WIDGET, 1, WIDGET
