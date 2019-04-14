package gtk

type CalendarDisplayOptions int

// ENUM_NAME = GTK_CALENDAR_SHOW_HEADING, 1
// ENUM_NAME = GTK_CALENDAR_SHOW_DAY_NAMES, 2
// ENUM_NAME = GTK_CALENDAR_NO_MONTH_CHANGE, 4
// ENUM_NAME = GTK_CALENDAR_SHOW_WEEK_NUMBERS, 8
// ENUM_NAME = GTK_CALENDAR_SHOW_DETAILS, 32

const (
	CALENDAR_SHOW_HEADING      CalendarDisplayOptions = 1 << 0
	CALENDAR_SHOW_DAY_NAMES                           = 1 << 1
	CALENDAR_NO_MONTH_CHANGE                          = 1 << 2
	CALENDAR_SHOW_WEEK_NUMBERS                        = 1 << 3
	CALENDAR_SHOW_DETAILS                             = 1 << 5
)
