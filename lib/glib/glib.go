package glib

import sugar "github.com/jopbrown/gtk-sugar"

var candy sugar.Candy

func GiveCandy(c sugar.Candy) {
	candy = c
}

func Candy() sugar.Candy {
	return candy
}

// FUNCTION_NAME = g_free, NONE, NONE, 1, WIDGET
func Free(obj sugar.CandyWrapper) {
	obj.Candy().Guify("g_free", obj)
}

// FUNCTION_NAME = g_io_channel_read_line_string, NONE, INT, 4, WIDGET, STRING, NULL, NULL
// FUNCTION_NAME = g_io_channel_unix_new, NONE, WIDGET, 1, INT
// FUNCTION_NAME = g_locale_from_utf8, NONE, STRING, 5, STRING, LONG, NULL, NULL, NULL
// FUNCTION_NAME = g_locale_to_utf8, NONE, STRING, 5, STRING, LONG, NULL, NULL, NULL
// FUNCTION_NAME = g_malloc, NONE, POINTER, 1, INT
// FUNCTION_NAME = g_printf, NONE, NONE, 2, STRING, POINTER
// FUNCTION_NAME = g_snprintf, NONE, NONE, 4, POINTER, INT, STRING, DOUBLE
// FUNCTION_NAME = g_random_int_range, NONE, INT, 2, INT, INT
// FUNCTION_NAME = g_random_double, NONE, DOUBLE, 0
// FUNCTION_NAME = g_signal_stop_emission_by_name, NONE, NONE, 2, WIDGET, STRING
