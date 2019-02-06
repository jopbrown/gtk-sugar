package gtks

type Button struct {
	Bin
}

// FUNCTION_NAME = gtk_button_new, clicked, WIDGET, 0
func (gtk *Gtk) NewButton() *Button {
	id := gtk.Guify("gtk_button_new").String()
	btn := Button{}
	btn.CandyWrapper = gtk.NewWrapper(id)
	return &btn
}

// FUNCTION_NAME = gtk_button_new_with_label, clicked, WIDGET, 1, STRING
func (gtk *Gtk) NewButtonWithLabel(label string) *Button {
	id := gtk.Guify("gtk_button_new_with_label", label).String()
	btn := Button{}
	btn.CandyWrapper = gtk.NewWrapper(id)
	return &btn
}

// FUNCTION_NAME = gtk_button_new_with_mnemonic, clicked, WIDGET, 1, STRING
func (gtk *Gtk) NewButtonWithMnemonic(label string) *Button {
	id := gtk.Guify("gtk_button_new_with_mnemonic", label).String()
	btn := Button{}
	btn.CandyWrapper = gtk.NewWrapper(id)
	return &btn
}

// FUNCTION_NAME = gtk_button_new_from_stock, clicked, WIDGET, 1, STRING
func (gtk *Gtk) NewButtonWithStock(stockID string) *Button {
	id := gtk.Guify("gtk_button_new_from_stock", stockID).String()
	btn := Button{}
	btn.CandyWrapper = gtk.NewWrapper(id)
	return &btn
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

// FUNCTION_NAME = gtk_toggle_button_new, clicked, WIDGET, 0
func (gtk *Gtk) NewToggleButton() *ToggleButton {
	id := gtk.Guify("gtk_toggle_button_new").String()
	btn := ToggleButton{}
	btn.CandyWrapper = gtk.NewWrapper(id)
	return &btn
}

// FUNCTION_NAME = gtk_toggle_button_new_with_label, clicked, WIDGET, 1, STRING
func (gtk *Gtk) NewToggleButtonWithLabel(label string) *ToggleButton {
	id := gtk.Guify("gtk_toggle_button_new_with_label", label).String()
	btn := ToggleButton{}
	btn.CandyWrapper = gtk.NewWrapper(id)
	return &btn
}

// FUNCTION_NAME = gtk_toggle_button_new_with_mnemonic, clicked, WIDGET, 1, STRING
func (gtk *Gtk) NewToggleButtonWithMnemonic(label string) *ToggleButton {
	id := gtk.Guify("gtk_toggle_button_new_with_mnemonic", label).String()
	btn := ToggleButton{}
	btn.CandyWrapper = gtk.NewWrapper(id)
	return &btn
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

// FUNCTION_NAME = gtk_check_button_new_with_label, clicked, WIDGET, 1, STRING
func (gtk *Gtk) NewCheckButtonWithLabel(label string) *CheckButton {
	id := gtk.Guify("gtk_check_button_new_with_label", label).String()
	btn := CheckButton{}
	btn.CandyWrapper = gtk.NewWrapper(id)
	return &btn
}

type ButtonBox struct {
	Button
}

// FUNCTION_NAME = gtk_button_box_set_layout, NONE, NONE, 2, WIDGET, INT
func (btnbox *ButtonBox) SetLayout(layout ButtonBoxStyle) {
	btnbox.Candy().Guify("gtk_button_box_set_layout", btnbox, layout)
}

type HButtonBox struct {
	ButtonBox
}

// FUNCTION_NAME = gtk_hbutton_box_new, NONE, WIDGET, 0
func (gtk *Gtk) NewHButtonBox() *HButtonBox {
	id := gtk.Guify("gtk_hbutton_box_new").String()
	btn := HButtonBox{}
	btn.CandyWrapper = gtk.NewWrapper(id)
	return &btn
}

type VButtonBox struct {
	ButtonBox
}

// FUNCTION_NAME = gtk_vbutton_box_new, NONE, WIDGET, 0
func (gtk *Gtk) NewVButtonBox() *VButtonBox {
	id := gtk.Guify("gtk_vbutton_box_new").String()
	btn := VButtonBox{}
	btn.CandyWrapper = gtk.NewWrapper(id)
	return &btn
}
