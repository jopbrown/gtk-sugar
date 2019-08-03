package gtk

import (
	sugar "github.com/jopbrown/gtk-sugar"
	"github.com/jopbrown/gtk-sugar/lib/glib"
)

type FileChooser struct {
	sugar.CandyWrapper
}

func NewFileChooser(candy sugar.Candy, id string) *FileChooser {
	obj := FileChooser{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_file_chooser_get_filename, NONE, STRING, 1, WIDGET
func (obj *FileChooser) GetFilename() string {
	return obj.Candy().Guify("gtk_file_chooser_get_filename", obj).String()
}

// FUNCTION_NAME = gtk_file_chooser_set_filename, NONE, BOOL, 2, WIDGET, STRING
func (obj *FileChooser) SetFilename(filename string) bool {
	return obj.Candy().Guify("gtk_file_chooser_set_filename", obj, filename).MustBool()
}

// FUNCTION_NAME = gtk_file_chooser_get_uri, NONE, STRING, 1, WIDGET
func (obj *FileChooser) GetURI() string {
	return obj.Candy().Guify("gtk_file_chooser_get_uri", obj).String()
}

// FUNCTION_NAME = gtk_file_chooser_set_uri, NONE, BOOL, 2, WIDGET, STRING
func (obj *FileChooser) SetURI(uri string) bool {
	return obj.Candy().Guify("gtk_file_chooser_set_uri", obj, uri).MustBool()
}

// FUNCTION_NAME = gtk_file_chooser_get_current_folder, NONE, STRING, 1, WIDGET
func (obj *FileChooser) GetCurrentFolder() string {
	return obj.Candy().Guify("gtk_file_chooser_get_current_folder", obj).String()
}

// FUNCTION_NAME = gtk_file_chooser_set_current_folder, NONE, BOOL, 2, WIDGET, STRING
func (obj *FileChooser) SetCurrentFolder(folder string) bool {
	return obj.Candy().Guify("gtk_file_chooser_set_current_folder", obj, folder).MustBool()
}

// FUNCTION_NAME = gtk_file_chooser_set_current_name, NONE, NONE, 2, WIDGET, STRING
func (obj *FileChooser) SetCurrentName(name string) {
	obj.Candy().Guify("gtk_file_chooser_set_current_name", obj, name)
}

// FUNCTION_NAME = gtk_file_chooser_add_filter, NONE, NONE, 2, WIDGET, WIDGET
func (obj *FileChooser) AddFilter(filter *FileFilter) {
	obj.Candy().Guify("gtk_file_chooser_add_filter", obj, filter)
}

// FUNCTION_NAME = gtk_file_chooser_set_filter, NONE, NONE, 2, WIDGET, WIDGET
func (obj *FileChooser) SetFilter(filter *FileFilter) {
	obj.Candy().Guify("gtk_file_chooser_set_filter", obj, filter)
}

// FUNCTION_NAME = gtk_file_chooser_set_select_multiple, NONE, NONE, 2, WIDGET, BOOL
func (obj *FileChooser) SetSelectMultiple(selectMultiple bool) {
	obj.Candy().Guify("gtk_file_chooser_set_select_multiple", obj, selectMultiple)
}

// FUNCTION_NAME = gtk_file_chooser_set_local_only, NONE, NONE, 2, WIDGET, BOOL
func (obj *FileChooser) SetLocalOnly(localOnly bool) {
	obj.Candy().Guify("gtk_file_chooser_set_local_only", obj, localOnly)
}

// FUNCTION_NAME = gtk_file_chooser_set_show_hidden, NONE, NONE, 2, WIDGET, BOOL
func (obj *FileChooser) SetShowHidden(showHidden bool) {
	obj.Candy().Guify("gtk_file_chooser_set_show_hidden", obj, showHidden)
}

// FUNCTION_NAME = gtk_file_chooser_set_do_overwrite_confirmation, NONE, NONE, 2, WIDGET, BOOL
func (obj *FileChooser) SetDoOverwriteConfirmation(doOverwriteConfirmation bool) {
	obj.Candy().Guify("gtk_file_chooser_set_do_overwrite_confirmation", obj, doOverwriteConfirmation)
}

// FUNCTION_NAME = gtk_file_chooser_set_create_folders, NONE, NONE, 2, WIDGET, BOOL
func (obj *FileChooser) SetCreateFolders(createFolders bool) {
	obj.Candy().Guify("gtk_file_chooser_set_create_folders", obj, createFolders)
}

type FileFilter struct {
	glib.Object
}

func NewFileFilter(candy sugar.Candy, id string) *FileFilter {
	obj := FileFilter{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_file_filter_new, NONE, WIDGET, 0
func FileFilterNew() *FileFilter {
	id := Candy().Guify("gtk_file_filter_new").String()
	return NewFileFilter(Candy(), id)
}

// FUNCTION_NAME = gtk_file_filter_set_name, NONE, NONE, 2, WIDGET, STRING
func (obj *FileFilter) SetName(name string) {
	obj.Candy().Guify("gtk_file_filter_set_name", obj, name)
}

// FUNCTION_NAME = gtk_file_filter_add_pattern, NONE, NONE, 2, WIDGET, STRING
func (obj *FileFilter) AddPattern(pattern string) {
	obj.Candy().Guify("gtk_file_filter_add_pattern", obj, pattern)
}

type FileChooserDialog struct {
	Dialog
	FileChooser
}

func NewFileChooserDialog(candy sugar.Candy, id string) *FileChooserDialog {
	obj := FileChooserDialog{}
	obj.CandyWrapper = candy.NewWrapper(id)
	obj.FileChooser = *NewFileChooser(candy, id)
	return &obj
}

// FUNCTION_NAME = gtk_file_chooser_dialog_new, NONE, WIDGET, 5, STRING, WIDGET, INT, STRING, VARARGS
func FileChooserDialogNew(title string, parent *Window, action FileChooserAction, firstBtn DialogButton, otherBtns ...DialogButton) *FileChooserDialog {
	args := make(sugar.Varargs, 0, len(otherBtns)*2+1)
	args = append(args, firstBtn.ID)
	for _, p := range otherBtns {
		args = append(args, p.Name)
		args = append(args, p.ID)
	}
	id := Candy().Guify("gtk_file_chooser_dialog_new", title, parent, action, firstBtn.Name, args, nil).String()
	return NewFileChooserDialog(Candy(), id)
}

type FileChooserWidget struct {
	Box
}

func NewFileChooserWidget(candy sugar.Candy, id string) *FileChooserWidget {
	obj := FileChooserWidget{}
	obj.CandyWrapper = candy.NewWrapper(id)
	return &obj
}

// FUNCTION_NAME = gtk_file_chooser_widget_new, NONE, WIDGET, 1, INT
func FileChooserWidgetNew(action FileChooserAction) *FileChooserWidget {
	id := Candy().Guify("gtk_file_chooser_widget_new", action).String()
	return NewFileChooserWidget(Candy(), id)
}
