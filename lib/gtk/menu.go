package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type MenuShell struct {
	Container
}

func NewMenuShell(candy sugar.Candy, id string) *MenuShell {
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

func NewMenu(candy sugar.Candy, id string) *Menu {
	obj := Menu{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_menu_new, NONE, WIDGET, 0
func MenuNew() *Menu {
	id := Candy().Guify("gtk_menu_new").String()
	return NewMenu(Candy(), id)
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

func NewMenuBar(candy sugar.Candy, id string) *MenuBar {
	obj := MenuBar{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_menu_bar_new, NONE, WIDGET, 0
func MenuBarNew() *MenuBar {
	id := Candy().Guify("gtk_menu_bar_new").String()
	return NewMenuBar(Candy(), id)
}

type MenuItem struct {
	Bin
}

func NewMenuItem(candy sugar.Candy, id string) *MenuItem {
	obj := MenuItem{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_menu_item_new, activate, WIDGET, 0
func MenuItemNew() *MenuItem {
	id := Candy().Guify("gtk_menu_item_new").String()
	return NewMenuItem(Candy(), id)
}

// FUNCTION_NAME = gtk_menu_item_new_with_label, activate, WIDGET, 1, STRING
func MenuItemNewWithLabel(label string) *MenuItem {
	id := Candy().Guify("gtk_menu_item_new_with_label", label).String()
	return NewMenuItem(Candy(), id)
}

// FUNCTION_NAME = gtk_menu_item_new_with_mnemonic, activate, WIDGET, 1, STRING
func MenuItemNewWithMnemonic(label string) *MenuItem {
	id := Candy().Guify("gtk_menu_item_new_with_mnemonic", label).String()
	return NewMenuItem(Candy(), id)
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

func NewTearoffMenuItem(candy sugar.Candy, id string) *TearoffMenuItem {
	obj := TearoffMenuItem{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_tearoff_menu_item_new, activate, WIDGET, 0
func TearoffMenuItemNew() *TearoffMenuItem {
	id := Candy().Guify("gtk_tearoff_menu_item_new").String()
	return NewTearoffMenuItem(Candy(), id)
}

type SeparatorMenuItem struct {
	MenuItem
}

func NewSeparatorMenuItem(candy sugar.Candy, id string) *SeparatorMenuItem {
	obj := SeparatorMenuItem{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_separator_menu_item_new, NONE, WIDGET, 0
func SeparatorMenuItemNew() *SeparatorMenuItem {
	id := Candy().Guify("gtk_separator_menu_item_new").String()
	return NewSeparatorMenuItem(Candy(), id)
}

type CheckMenuItem struct {
	MenuItem
}

func NewCheckMenuItem(candy sugar.Candy, id string) *CheckMenuItem {
	obj := CheckMenuItem{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_check_menu_item_new_with_label, activate, WIDGET, 1, STRING
func CheckMenuItemNewWithLabel(label string) *CheckMenuItem {
	id := Candy().Guify("gtk_check_menu_item_new_with_label", label).String()
	return NewCheckMenuItem(Candy(), id)
}

// FUNCTION_NAME = gtk_check_menu_item_new_with_mnemonic, activate, WIDGET, 1, STRING
func CheckMenuItemNewWithMnemonic(label string) *CheckMenuItem {
	id := Candy().Guify("gtk_check_menu_item_new_with_mnemonic", label).String()
	return NewCheckMenuItem(Candy(), id)
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

func NewImageMenuItem(candy sugar.Candy, id string) *ImageMenuItem {
	obj := ImageMenuItem{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_image_menu_item_new_from_stock, NONE, WIDGET, 2, STRING, WIDGET
func ImageMenuItemNewFromStock(stockID string, accelGroup *AccelGroup) *ImageMenuItem {
	id := Candy().Guify("gtk_image_menu_item_new_from_stock", stockID, accelGroup).String()
	return NewImageMenuItem(Candy(), id)
}
