package pango

// ENUM_NAME = ELLIPSIZE_NONE, 1
// ENUM_NAME = ELLIPSIZE_START, 2
// ENUM_NAME = ELLIPSIZE_MIDDLE, 3
// ENUM_NAME = ELLIPSIZE_END, 4
type EllipsizeMode int

const (
	ELLIPSIZE_NONE = iota
	ELLIPSIZE_START
	ELLIPSIZE_MIDDLE
	ELLIPSIZE_END
)
