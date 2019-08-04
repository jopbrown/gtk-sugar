package gtk

// FUNCTION_NAME = gtk_rc_parse, NONE, NONE, 1, STRING
func RcParse(filename string) {
	Candy().Guify("gtk_rc_parse", filename)
}

// FUNCTION_NAME = gtk_rc_parse_string, NONE, NONE, 1, STRING
func RcParseString(rcString string) {
	Candy().Guify("gtk_rc_parse_string", rcString)
}

// FUNCTION_NAME = gtk_rc_reparse_all, NONE, BOOL, 0
func RcReparseAll() bool {
	return Candy().Guify("gtk_rc_reparse_all").MustBool()
}

// FUNCTION_NAME = gtk_rc_add_default_file, NONE, NONE, 1, STRING
func RcAddDefaultFile(filename string) {
	Candy().Guify("gtk_rc_add_default_file", filename)
}
