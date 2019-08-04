package gtk

import "time"

// FUNCTION_NAME = gtk_main, NONE, NONE, 0
func Main() {
	Candy().RunLoop()
}

// FUNCTION_NAME = gtk_main_quit, NONE, NONE, 0
func MainQuit() {
	Candy().StopLoop()
}

// FUNCTION_NAME = gtk_main_iteration, NONE, BOOL, 0
func MainIteration() bool {
	return Candy().Guify("gtk_main_iteration").MustBool()
}

// FUNCTION_NAME = gtk_main_iteration_do, NONE, BOOL, 1, BOOL
func MainIterationDo(blocking bool) bool {
	return Candy().Guify("gtk_main_iteration_do", blocking).MustBool()
}

// FUNCTION_NAME = gtk_events_pending, NONE, BOOL, 0
func EventsPending() bool {
	return Candy().Guify("gtk_events_pending").MustBool()
}

// FUNCTION_NAME = gtk_get_current_event_time, NONE, INT, 0
func GetCurrentEventTime() time.Time {
	return time.Unix(Candy().Guify("gtk_get_current_event_time").MustInt64(), 0)
}
