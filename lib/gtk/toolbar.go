package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Toolbar struct {
	Container
}

func NewToolbar(candy sugar.Candy, id string) *Toolbar {
	obj := Toolbar{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_toolbar_new, NONE, WIDGET, 0
func ToolbarNew() *Toolbar {
	id := Candy().Guify("gtk_toolbar_new").String()
	return NewToolbar(Candy(), id)
}

// FUNCTION_NAME = gtk_toolbar_insert, NONE, WIDGET, 3, WIDGET, WIDGET, INT
func (obj *Toolbar) Insert(item IToolItem, pos int) {
	obj.Candy().Guify("gtk_toolbar_insert", obj, item, pos)
}

// FUNCTION_NAME = gtk_toolbar_set_style, NONE, NONE, 2, WIDGET, INT
func (obj *Toolbar) SetStyle(style ToolbarStyle) {
	obj.Candy().Guify("gtk_toolbar_set_style", obj, style)
}

type IToolItem interface {
	IWidget
}

type ToolItem struct {
	Bin
}

func NewToolItem(candy sugar.Candy, id string) *ToolItem {
	obj := ToolItem{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

type ToolButton struct {
	ToolItem
}

func NewToolButton(candy sugar.Candy, id string) *ToolButton {
	obj := ToolButton{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_tool_button_new, clicked, WIDGET, 2, WIDGET, STRING
func ToolButtonNew(iconWidget IWidget, label string) *ToolButton {
	id := Candy().Guify("gtk_tool_button_new", iconWidget, label).String()
	return NewToolButton(Candy(), id)
}

type SeparatorToolItem struct {
	ToolItem
}

func NewSeparatorToolItem(candy sugar.Candy, id string) *SeparatorToolItem {
	obj := SeparatorToolItem{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_separator_tool_item_new, NONE, WIDGET, 0
func SeparatorToolItemNew() *SeparatorToolItem {
	id := Candy().Guify("gtk_separator_tool_item_new").String()
	return NewSeparatorToolItem(Candy(), id)
}
