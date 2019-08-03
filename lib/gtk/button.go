package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Button struct {
	Bin
}

func NewButtonFromCandy(candy sugar.Candy, id string) *Button {
	v := Button{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_button_new, clicked, WIDGET, 0
func ButtonNew() *Button {
	id := Candy().Guify("gtk_button_new").String()
	return NewButtonFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_button_new_with_label, clicked, WIDGET, 1, STRING
func ButtonNewWithLabel(label string) *Button {
	id := Candy().Guify("gtk_button_new_with_label", label).String()
	return NewButtonFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_button_new_with_mnemonic, clicked, WIDGET, 1, STRING
func ButtonNewWithMnemonic(label string) *Button {
	id := Candy().Guify("gtk_button_new_with_mnemonic", label).String()
	return NewButtonFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_button_new_from_stock, clicked, WIDGET, 1, STRING
func ButtonNewWithStock(stockID string) *Button {
	id := Candy().Guify("gtk_button_new_from_stock", stockID).String()
	return NewButtonFromCandy(Candy(), id)
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

func NewToggleButtonFromCandy(candy sugar.Candy, id string) *ToggleButton {
	v := ToggleButton{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_toggle_button_new, clicked, WIDGET, 0
func ToggleButtonNew() *ToggleButton {
	id := Candy().Guify("gtk_toggle_button_new").String()
	return NewToggleButtonFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_toggle_button_new_with_label, clicked, WIDGET, 1, STRING
func ToggleButtonNewWithLabel(label string) *ToggleButton {
	id := Candy().Guify("gtk_toggle_button_new_with_label", label).String()
	return NewToggleButtonFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_toggle_button_new_with_mnemonic, clicked, WIDGET, 1, STRING
func ToggleButtonNewWithMnemonic(label string) *ToggleButton {
	id := Candy().Guify("gtk_toggle_button_new_with_mnemonic", label).String()
	return NewToggleButtonFromCandy(Candy(), id)
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

func NewCheckButtonFromCandy(candy sugar.Candy, id string) *CheckButton {
	v := CheckButton{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_check_button_new_with_label, clicked, WIDGET, 1, STRING
func CheckButtonNewWithLabel(label string) *CheckButton {
	id := Candy().Guify("gtk_check_button_new_with_label", label).String()
	return NewCheckButtonFromCandy(Candy(), id)
}

type RadioButton struct {
	CheckButton
}

func NewRadioButtonFromCandy(candy sugar.Candy, id string) *RadioButton {
	obj := RadioButton{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_radio_button_new, clicked, WIDGET, 1, WIDGET
func RadioButtonNew() *RadioButton {
	id := Candy().Guify("gtk_radio_button_new", nil).String()
	return NewRadioButtonFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_radio_button_new_with_label, clicked, WIDGET, 2, WIDGET, STRING
func RadioButtonNewWithLabel(label string) *RadioButton {
	id := Candy().Guify("gtk_radio_button_new_with_label", nil, label).String()
	return NewRadioButtonFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_radio_button_new_from_widget, clicked, WIDGET, 1, WIDGET
func RadioButtonNewFromWidget(radioGroupMember *RadioButton) *RadioButton {
	id := Candy().Guify("gtk_radio_button_new_from_widget", radioGroupMember).String()
	return NewRadioButtonFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_radio_button_new_with_label_from_widget, clicked, WIDGET, 2, WIDGET, STRING
func RadioButtonNewWithLabelFromWidget(radioGroupMember *RadioButton, label string) *RadioButton {
	id := Candy().Guify("gtk_radio_button_new_with_label_from_widget", radioGroupMember, label).String()
	return NewRadioButtonFromCandy(Candy(), id)
}

type SpinButton struct {
	Entry
}

func NewSpinButtonFromCandy(candy sugar.Candy, id string) *SpinButton {
	obj := SpinButton{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_spin_button_new, NONE, WIDGET, 3, WIDGET, DOUBLE, INT
func SpinButtonNew(adjustment Adjustment, climbRate float64, digits uint) *SpinButton {
	id := Candy().Guify("gtk_spin_button_new", adjustment, climbRate, digits).String()
	return NewSpinButtonFromCandy(Candy(), id)
}

// FUNCTION_NAME = gtk_spin_button_new_with_range, NONE, WIDGET, 3, DOUBLE, DOUBLE, DOUBLE
func SpinButtonNewWithRange(min, max, step float64) *SpinButton {
	id := Candy().Guify("gtk_spin_button_new_with_range", min, max, step).String()
	return NewSpinButtonFromCandy(Candy(), id)
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

func NewButtonBoxFromCandy(candy sugar.Candy, id string) *ButtonBox {
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

func NewHButtonBoxFromCandy(candy sugar.Candy, id string) *HButtonBox {
	v := HButtonBox{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_hbutton_box_new, NONE, WIDGET, 0
func HButtonBoxNew() *HButtonBox {
	id := Candy().Guify("gtk_hbutton_box_new").String()
	return NewHButtonBoxFromCandy(Candy(), id)
}

type VButtonBox struct {
	ButtonBox
}

func NewVButtonBoxFromCandy(candy sugar.Candy, id string) *VButtonBox {
	v := VButtonBox{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_vbutton_box_new, NONE, WIDGET, 0
func VButtonBoxNew() *VButtonBox {
	id := Candy().Guify("gtk_vbutton_box_new").String()
	return NewVButtonBoxFromCandy(Candy(), id)
}
