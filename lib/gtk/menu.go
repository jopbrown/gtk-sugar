package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type MenuShell struct {
	Container
}

func NewMenuShellFromCandy(candy sugar.Candy, id string) *MenuShell {
	obj := MenuShell{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_menu_shell_append, NONE, NONE, 2, WIDGET, WIDGET
func (obj *MenuShell) Append(child IWidget) {
	obj.Candy().Guify("gtk_menu_shell_append", obj, child)
}

type Menu struct {
	MenuShell
}

func NewMenuFromCandy(candy sugar.Candy, id string) *Menu {
	obj := Menu{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_menu_new, NONE, WIDGET, 0
func NewMenu() *Menu {
	id := Candy().Guify("gtk_menu_new").String()
	return NewMenuFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_menu_set_title, NONE, NONE, 2, WIDGET, STRING
func (obj *Menu) SetTitle(title string) {
	obj.Candy().Guify("gtk_menu_set_title", obj, title)
}

// FUNCTION_NAME = gtk_menu_popup, NONE, NONE, 7, WIDGET, NULL, NULL, NULL, NULL, INT, INT
func (obj *Menu) Popup(button uint, activateTime uint32) {
	obj.Candy().Guify("gtk_menu_popup", obj, nil, nil, nil, nil, button, activateTime)
}

type MenuBar struct {
	MenuShell
}

func NewMenuBarFromCandy(candy sugar.Candy, id string) *MenuBar {
	obj := MenuBar{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_menu_bar_new, NONE, WIDGET, 0
func NewMenuBar() *MenuBar {
	id := Candy().Guify("gtk_menu_bar_new").String()
	return NewMenuBarFromCandy(Candy(), id)
}

type MenuItem struct {
	Bin
}

func NewMenuItemFromCandy(candy sugar.Candy, id string) *MenuItem {
	obj := MenuItem{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_menu_item_new, activate, WIDGET, 0
func NewMenuItem() *MenuItem {
	id := Candy().Guify("gtk_menu_item_new").String()
	return NewMenuItemFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_menu_item_new_with_label, activate, WIDGET, 1, STRING
func NewMenuItemWithLabel(label string) *MenuItem {
	id := Candy().Guify("gtk_menu_item_new_with_label", label).String()
	return NewMenuItemFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_menu_item_new_with_mnemonic, activate, WIDGET, 1, STRING
func NewMenuItemWithMnemonic(label string) *MenuItem {
	id := Candy().Guify("gtk_menu_item_new_with_mnemonic", label).String()
	return NewMenuItemFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_menu_item_set_right_justified, NONE, NONE, 2, WIDGET, BOOL
func (obj *MenuItem) SetRightJustified(rightJustified bool) {
	obj.Candy().Guify("gtk_menu_item_set_right_justified", obj, rightJustified)
}

// FUNCTION_NAME = gtk_menu_item_set_submenu, NONE, NONE, 2, WIDGET, WIDGET
func (obj *MenuItem) SetSubmenu(subMenu IWidget) {
	obj.Candy().Guify("gtk_menu_item_set_submenu", obj, subMenu)
}

type TearoffMenuItem struct {
	MenuItem
}

func NewTearoffMenuItemFromCandy(candy sugar.Candy, id string) *TearoffMenuItem {
	obj := TearoffMenuItem{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_tearoff_menu_item_new, activate, WIDGET, 0
func NewTearoffMenuItem() *TearoffMenuItem {
	id := Candy().Guify("gtk_tearoff_menu_item_new").String()
	return NewTearoffMenuItemFromCandy(Candy(), id)
}

type SeparatorMenuItem struct {
	MenuItem
}

func NewSeparatorMenuItemFromCandy(candy sugar.Candy, id string) *SeparatorMenuItem {
	obj := SeparatorMenuItem{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_separator_menu_item_new, NONE, WIDGET, 0
func NewSeparatorMenuItem() *SeparatorMenuItem {
	id := Candy().Guify("gtk_separator_menu_item_new").String()
	return NewSeparatorMenuItemFromCandy(Candy(), id)
}

type CheckMenuItem struct {
	MenuItem
}

func NewCheckMenuItemFromCandy(candy sugar.Candy, id string) *CheckMenuItem {
	obj := CheckMenuItem{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_check_menu_item_new_with_label, activate, WIDGET, 1, STRING
func NewCheckMenuItemWithLabel(label string) *CheckMenuItem {
	id := Candy().Guify("gtk_check_menu_item_new_with_label", label).String()
	return NewCheckMenuItemFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_check_menu_item_new_with_mnemonic, activate, WIDGET, 1, STRING
func NewCheckMenuItemWithMnemonic(label string) *CheckMenuItem {
	id := Candy().Guify("gtk_check_menu_item_new_with_mnemonic", label).String()
	return NewCheckMenuItemFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_check_menu_item_get_active, NONE, BOOL, 1, WIDGET
func (obj *CheckMenuItem) GetActive() bool {
	return obj.Candy().Guify("gtk_check_menu_item_get_active", obj).MustBool()
}

// FUNCTION_NAME = gtk_check_menu_item_set_active, NONE, NONE, 2, WIDGET, BOOL
func (obj *CheckMenuItem) SetActive(isActive bool) {
	obj.Candy().Guify("gtk_check_menu_item_set_active", obj, isActive)
}

type ImageMenuItem struct {
	sugar.CandyWrapper
}

func NewImageMenuItemFromCandy(candy sugar.Candy, id string) *ImageMenuItem {
	obj := ImageMenuItem{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_image_menu_item_new_from_stock, NONE, WIDGET, 2, STRING, WIDGET
func NewImageMenuItemFromStock(stockID string, accelGroup *AccelGroup) *ImageMenuItem {
	id := Candy().Guify("gtk_image_menu_item_new_from_stock", stockID, accelGroup).String()
	return NewImageMenuItemFromCandy(Candy(), id)
}
