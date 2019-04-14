package gdk

type Atom int

const (
	NONE                    Atom = 0
	SELECTION_PRIMARY            = 1
	SELECTION_SECONDARY          = 2
	SELECTION_CLIPBOARD          = 69
	TARGET_BITMAP                = 5
	TARGET_COLORMAP              = 7
	TARGET_DRAWABLE              = 17
	TARGET_PIXMAP                = 20
	TARGET_STRING                = 31
	SELECTION_TYPE_ATOM          = 4
	SELECTION_TYPE_BITMAP        = 5
	SELECTION_TYPE_COLORMAP      = 7
	SELECTION_TYPE_DRAWABLE      = 17
	SELECTION_TYPE_INTEGER       = 19
	SELECTION_TYPE_PIXMAP        = 20
	SELECTION_TYPE_WINDOW        = 33
	SELECTION_TYPE_STRING        = 31
)
