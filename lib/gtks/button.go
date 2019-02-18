package gtks

import sugar "github.com/jopbrown/gtk-sugar"

type Button struct {
	Bin
}

func NewButton(candy sugar.Candy, id string) *Button {
	v := Button{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_button_new, clicked, WIDGET, 0
func (gtk *Gtk) NewButton() *Button {
	id := gtk.Guify("gtk_button_new").String()
	return NewButton(gtk, id)
}

// FUNCTION_NAME = gtk_button_new_with_label, clicked, WIDGET, 1, STRING
func (gtk *Gtk) NewButtonWithLabel(label string) *Button {
	id := gtk.Guify("gtk_button_new_with_label", label).String()
	return NewButton(gtk, id)
}

// FUNCTION_NAME = gtk_button_new_with_mnemonic, clicked, WIDGET, 1, STRING
func (gtk *Gtk) NewButtonWithMnemonic(label string) *Button {
	id := gtk.Guify("gtk_button_new_with_mnemonic", label).String()
	return NewButton(gtk, id)
}

// FUNCTION_NAME = gtk_button_new_from_stock, clicked, WIDGET, 1, STRING
func (gtk *Gtk) NewButtonWithStock(stockID string) *Button {
	id := gtk.Guify("gtk_button_new_from_stock", stockID).String()
	return NewButton(gtk, id)
}

// FUNCTION_NAME = gtk_button_set_use_stock, NONE, NONE, 2, WIDGET, BOOL
func (btn *Button) SetUseStock(useStock bool) {
	btn.Candy().Guify("gtk_button_set_use_stock", btn, useStock)
}

// FUNCTION_NAME = gtk_button_set_label, NONE, NONE, 2, WIDGET, STRING
func (btn *Button) SetLabel(label string) {
	btn.Candy().Guify("gtk_button_set_label", btn, label)
}

// FUNCTION_NAME = gtk_button_set_relief, NONE, NONE, 2, WIDGET, INT
func (btn *Button) SetRelief(newstyle ReliefStyle) {
	btn.Candy().Guify("gtk_button_set_relief", btn, newstyle)
}

// FUNCTION_NAME = gtk_button_clicked, NONE, NONE, 1, WIDGET
func (btn *Button) Clicked() {
	btn.Candy().Guify("gtk_button_clicked", btn)
}

type ToggleButton struct {
	Button
}

func NewToggleButton(candy sugar.Candy, id string) *ToggleButton {
	v := ToggleButton{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_toggle_button_new, clicked, WIDGET, 0
func (gtk *Gtk) NewToggleButton() *ToggleButton {
	id := gtk.Guify("gtk_toggle_button_new").String()
	return NewToggleButton(gtk, id)
}

// FUNCTION_NAME = gtk_toggle_button_new_with_label, clicked, WIDGET, 1, STRING
func (gtk *Gtk) NewToggleButtonWithLabel(label string) *ToggleButton {
	id := gtk.Guify("gtk_toggle_button_new_with_label", label).String()
	return NewToggleButton(gtk, id)
}

// FUNCTION_NAME = gtk_toggle_button_new_with_mnemonic, clicked, WIDGET, 1, STRING
func (gtk *Gtk) NewToggleButtonWithMnemonic(label string) *ToggleButton {
	id := gtk.Guify("gtk_toggle_button_new_with_mnemonic", label).String()
	return NewToggleButton(gtk, id)
}

// FUNCTION_NAME = gtk_toggle_button_get_active, NONE, BOOL, 1, WIDGET
func (btn *ToggleButton) GetActive() bool {
	return btn.Candy().Guify("gtk_toggle_button_get_active", btn).MustBool()
}

// FUNCTION_NAME = gtk_toggle_button_set_active, NONE, NONE, 2, WIDGET, BOOL
func (btn *ToggleButton) SetActive(isActive bool) {
	btn.Candy().Guify("gtk_toggle_button_set_active", btn, isActive)
}

// FUNCTION_NAME = gtk_toggle_button_get_mode, NONE, BOOL, 1, WIDGET
func (btn *ToggleButton) GetMode() bool {
	return btn.Candy().Guify("gtk_toggle_button_get_mode", btn).MustBool()
}

// FUNCTION_NAME = gtk_toggle_button_set_mode, NONE, NONE, 2, WIDGET, BOOL
func (btn *ToggleButton) SetMode(drawIndicator bool) {
	btn.Candy().Guify("gtk_toggle_button_set_mode", btn, drawIndicator)
}

// FUNCTION_NAME = gtk_toggle_button_get_inconsistent, NONE, BOOL, 1, WIDGET
func (btn *ToggleButton) GetInconsistent() bool {
	return btn.Candy().Guify("gtk_toggle_button_get_inconsistent", btn).MustBool()
}

// FUNCTION_NAME = gtk_toggle_button_set_inconsistent, NONE, NONE, 2, WIDGET, BOOL
func (btn *ToggleButton) SetInconsistent(setting bool) {
	btn.Candy().Guify("gtk_toggle_button_set_inconsistent", btn, setting)
}

// FUNCTION_NAME = gtk_toggle_button_toggled, NONE, NONE, 1, WIDGET
func (btn *ToggleButton) Toggled() {
	btn.Candy().Guify("gtk_toggle_button_toggled", btn)
}

type CheckButton struct {
	ToggleButton
}

func NewCheckButton(candy sugar.Candy, id string) *CheckButton {
	v := CheckButton{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_check_button_new_with_label, clicked, WIDGET, 1, STRING
func (gtk *Gtk) NewCheckButtonWithLabel(label string) *CheckButton {
	id := gtk.Guify("gtk_check_button_new_with_label", label).String()
	return NewCheckButton(gtk, id)
}

type RadioButton struct {
	CheckButton
}

func NewRadioButton(candy sugar.Candy, id string) *RadioButton {
	obj := RadioButton{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_radio_button_new, clicked, WIDGET, 1, WIDGET
func (gtk *Gtk) NewRadioButton() *RadioButton {
	id := gtk.Guify("gtk_radio_button_new", nil).String()
	return NewRadioButton(gtk, id)
}

// FUNCTION_NAME = gtk_radio_button_new_with_label, clicked, WIDGET, 2, WIDGET, STRING
func (gtk *Gtk) NewRadioButtonWithLabel(label string) *RadioButton {
	id := gtk.Guify("gtk_radio_button_new_with_label", nil, label).String()
	return NewRadioButton(gtk, id)
}

// FUNCTION_NAME = gtk_radio_button_new_from_widget, clicked, WIDGET, 1, WIDGET
func (gtk *Gtk) NewRadioButtonFromWidget(radioGroupMember *RadioButton) *RadioButton {
	id := gtk.Guify("gtk_radio_button_new_from_widget", radioGroupMember).String()
	return NewRadioButton(gtk, id)
}

// FUNCTION_NAME = gtk_radio_button_new_with_label_from_widget, clicked, WIDGET, 2, WIDGET, STRING
func (gtk *Gtk) NewRadioButtonWithLabelFromWidget(radioGroupMember *RadioButton, label string) *RadioButton {
	id := gtk.Guify("gtk_radio_button_new_with_label_from_widget", radioGroupMember, label).String()
	return NewRadioButton(gtk, id)
}

type SpinButton struct {
	Entry
}

func NewSpinButton(candy sugar.Candy, id string) *SpinButton {
	obj := SpinButton{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_spin_button_new, NONE, WIDGET, 3, WIDGET, DOUBLE, INT
func (gtk *Gtk) NewSpinButton(adjustment Adjustment, climbRate float64, digits uint) *SpinButton {
	id := gtk.Guify("gtk_spin_button_new", adjustment, climbRate, digits).String()
	return NewSpinButton(gtk, id)
}

// FUNCTION_NAME = gtk_spin_button_new_with_range, NONE, WIDGET, 3, DOUBLE, DOUBLE, DOUBLE
func (gtk *Gtk) NewSpinButtonWithRange(min, max, step float64) *SpinButton {
	id := gtk.Guify("gtk_spin_button_new_with_range", min, max, step).String()
	return NewSpinButton(gtk, id)
}

// FUNCTION_NAME = gtk_spin_button_get_value_as_int, NONE, INT, 1, WIDGET
func (obj *SpinButton) GetValueAsInt() int {
	return obj.Candy().Guify("gtk_spin_button_get_value_as_int", obj).MustInt()
}

// FUNCTION_NAME = gtk_spin_button_get_value, NONE, FLOAT, 1, WIDGET
func (obj *SpinButton) GetValue() float64 {
	return obj.Candy().Guify("gtk_spin_button_get_value", obj).MustFloat64()
}

// FUNCTION_NAME = gtk_spin_button_set_value, NONE, NONE, 2, WIDGET, DOUBLE
func (obj *SpinButton) SetValue(value float64) {
	obj.Candy().Guify("gtk_spin_button_set_value", obj, value)
}

// FUNCTION_NAME = gtk_spin_button_set_wrap, NONE, NONE, 2, WIDGET, BOOL
func (obj *SpinButton) SetWrap(wrap bool) {
	obj.Candy().Guify("gtk_spin_button_set_wrap", obj, wrap)
}

type ButtonBox struct {
	Button
}

func NewButtonBox(candy sugar.Candy, id string) *ButtonBox {
	v := ButtonBox{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_button_box_set_layout, NONE, NONE, 2, WIDGET, INT
func (btnbox *ButtonBox) SetLayout(layout ButtonBoxStyle) {
	btnbox.Candy().Guify("gtk_button_box_set_layout", btnbox, layout)
}

type HButtonBox struct {
	ButtonBox
}

func NewHButtonBox(candy sugar.Candy, id string) *HButtonBox {
	v := HButtonBox{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_hbutton_box_new, NONE, WIDGET, 0
func (gtk *Gtk) NewHButtonBox() *HButtonBox {
	id := gtk.Guify("gtk_hbutton_box_new").String()
	return NewHButtonBox(gtk, id)
}

type VButtonBox struct {
	ButtonBox
}

func NewVButtonBox(candy sugar.Candy, id string) *VButtonBox {
	v := VButtonBox{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_vbutton_box_new, NONE, WIDGET, 0
func (gtk *Gtk) NewVButtonBox() *VButtonBox {
	id := gtk.Guify("gtk_vbutton_box_new").String()
	return NewVButtonBox(gtk, id)
}
