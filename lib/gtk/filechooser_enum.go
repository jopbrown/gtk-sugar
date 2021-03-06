package gtk

// ENUM_NAME = GTK_FILE_CHOOSER_ACTION_OPEN, 0
// ENUM_NAME = GTK_FILE_CHOOSER_ACTION_SAVE, 1
// ENUM_NAME = GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER, 2
// ENUM_NAME = GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER, 3
type FileChooserAction int

const (
	FILE_CHOOSER_ACTION_OPEN FileChooserAction = iota
	FILE_CHOOSER_ACTION_SAVE
	FILE_CHOOSER_ACTION_SELECT_FOLDER
	FILE_CHOOSER_ACTION_CREATE_FOLDER
)
