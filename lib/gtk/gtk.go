package gtk

// FUNCTION_NAME = gtk_init, NONE, NONE, 2, NULL, NULL
func Init() {
	Candy().Guify("gtk_init", nil, nil)
}

// FUNCTION_NAME = gtk_exit, NONE, NONE, 1, INT
func Exit(errorCode int) {
	Candy().Guify("gtk_exit", errorCode)
}

// FUNCTION_NAME = gtk_check_version, NONE, STRING, 3, INT, INT, INT
func CheckVersion(major, minor, micro int) string {
	return Candy().Guify("gtk_check_version", major, minor, micro).String()
}
