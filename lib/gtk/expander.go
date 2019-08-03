package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Expander struct {
	Bin
}

func NewExpander(candy sugar.Candy, id string) *Expander {
	obj := Expander{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_expander_new, activate, WIDGET, 1, STRING
func ExpanderNew(label string) *Expander {
	id := Candy().Guify("gtk_expander_new", label).String()
	return NewExpander(Candy(), id)
}

// FUNCTION_NAME = gtk_expander_new_with_mnemonic, NONE, WIDGET, 1, STRING
func ExpanderNewWithMnemonic(label string) *Expander {
	id := Candy().Guify("gtk_expander_new_with_mnemonic", label).String()
	return NewExpander(Candy(), id)
}

// FUNCTION_NAME = gtk_expander_set_expanded, NONE, NONE, 2, WIDGET, BOOL
func (obj *Expander) SetExpanded(expanded bool) {
	obj.Candy().Guify("gtk_expander_set_expanded", obj, expanded)
}

// FUNCTION_NAME = gtk_expander_get_expanded, NONE, BOOL, 1, WIDGET
func (obj *Expander) GetExpanded() bool {
	return obj.Candy().Guify("gtk_expander_get_expanded", obj).MustBool()
}

// FUNCTION_NAME = gtk_expander_set_spacing, NONE, NONE, 2, WIDGET, INT
func (obj *Expander) SetSpacing(spacing int) {
	obj.Candy().Guify("gtk_expander_set_spacing", obj, spacing)
}

// FUNCTION_NAME = gtk_expander_get_spacing, NONE, INT, 1, WIDGET
func (obj *Expander) GetSpacing() int {
	return obj.Candy().Guify("gtk_expander_get_spacing", obj).MustInt()
}
