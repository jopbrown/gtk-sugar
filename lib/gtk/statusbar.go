package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Statusbar struct {
	Box
}

func NewStatusbarFromCandy(candy sugar.Candy, id string) *Statusbar {
	obj := Statusbar{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_statusbar_new, NONE, WIDGET, 0
func StatusbarNew() *Statusbar {
	id := Candy().Guify("gtk_statusbar_new").String()
	return NewStatusbarFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_statusbar_get_context_id, NONE, INT, 2, WIDGET, STRING
func (obj *Statusbar) GetContextId(contextDescription string) uint {
	return obj.Candy().Guify("gtk_statusbar_get_context_id", obj, contextDescription).MustUint()
}

// FUNCTION_NAME = gtk_statusbar_push, NONE, INT, 3, WIDGET, INT, STRING
func (obj *Statusbar) Push(contextID uint, text string) uint {
	return obj.Candy().Guify("gtk_statusbar_push", obj, contextID, text).MustUint()
}

// FUNCTION_NAME = gtk_statusbar_pop, NONE, NONE, 2, WIDGET, INT
func (obj *Statusbar) Pop(contextID uint) {
	obj.Candy().Guify("gtk_statusbar_pop", obj, contextID)
}

// FUNCTION_NAME = gtk_statusbar_remove, NONE, NONE, 3, WIDGET, INT, INT
func (obj *Statusbar) Remove(contextID, messageID uint) {
	obj.Candy().Guify("gtk_statusbar_remove", obj, contextID, messageID)
}

// FUNCTION_NAME = gtk_statusbar_remove_all, NONE, NONE, 2, WIDGET, INT
func (obj *Statusbar) RemoveAll(contextID uint) {
	obj.Candy().Guify("gtk_statusbar_remove_all", obj, contextID)
}

// FUNCTION_NAME = gtk_statusbar_set_has_resize_grip, NONE, NONE, 2, WIDGET, BOOL
func (obj *Statusbar) SetHasResizeGrip(setting bool) {
	obj.Candy().Guify("gtk_statusbar_set_has_resize_grip", obj, setting)
}
