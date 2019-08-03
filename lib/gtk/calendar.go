package gtk

import (
	"time"

	sugar "github.com/jopbrown/gtk-sugar"
)

type Calendar struct {
	Widget
}

func NewCalendarFromCandy(candy sugar.Candy, id string) *Calendar {
	obj := Calendar{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_calendar_new, NONE, WIDGET, 0
func CalendarNew() *Calendar {
	id := Candy().Guify("gtk_calendar_new").String()
	return NewCalendarFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_calendar_select_month, NONE, NONE, 3, WIDGET, INT, INT
func (obj *Calendar) SelectMonth(month, year uint) {
	obj.Candy().Guify("gtk_calendar_select_month", obj, month, year)
}

// FUNCTION_NAME = gtk_calendar_select_day, NONE, NONE, 2, WIDGET, INT
func (obj *Calendar) SelectDay(day uint) {
	obj.Candy().Guify("gtk_calendar_select_day", obj, day)
}

// FUNCTION_NAME = gtk_calendar_mark_day, NONE, NONE, 2, WIDGET, INT
func (obj *Calendar) MarkDay(day uint) {
	obj.Candy().Guify("gtk_calendar_mark_day", obj, day)
}

// FUNCTION_NAME = gtk_calendar_unmark_day, NONE, NONE, 2, WIDGET, INT
func (obj *Calendar) UnmarkDay(day uint) {
	obj.Candy().Guify("gtk_calendar_unmark_day", obj, day)
}

// FUNCTION_NAME = gtk_calendar_clear_marks, NONE, NONE, 1, WIDGET
func (obj *Calendar) ClearMarks() {
	obj.Candy().Guify("gtk_calendar_clear_marks", obj)
}

// FUNCTION_NAME = gtk_calendar_set_display_options, NONE, NONE, 2, WIDGET, INT
func (obj *Calendar) SetDisplayOptions(flags CalendarDisplayOptions) {
	obj.Candy().Guify("gtk_calendar_set_display_options", obj, flags)
}

// FUNCTION_NAME = gtk_calendar_get_display_options, NONE, INT, 1, WIDGET
func (obj *Calendar) GetDisplayOptions() CalendarDisplayOptions {
	flags := obj.Candy().Guify("gtk_calendar_get_display_options", obj).MustInt()
	return CalendarDisplayOptions(flags)
}

// FUNCTION_NAME = gtk_calendar_get_date, NONE, NONE, 4, WIDGET, PTR_INT, PTR_INT, PTR_INT
func (obj *Calendar) GetDate() time.Time {
	fields := obj.Candy().Guify("gtk_calendar_get_date", obj, 0, 0, 0).Fields()
	year := fields[0].MustInt()
	month := fields[1].MustInt()
	day := fields[2].MustInt()
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}
