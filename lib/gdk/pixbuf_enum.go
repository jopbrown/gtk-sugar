package gdk

// ENUM_NAME = GDK_COLORSPACE_RGB, 0
type Colorspace int

const (
	COLORSPACE_RGB Colorspace = iota
)

// ENUM_NAME = GDK_INTERP_NEAREST, 0
// ENUM_NAME = GDK_INTERP_TILES, 1
// ENUM_NAME = GDK_INTERP_BILINEAR, 2
// ENUM_NAME = GDK_INTERP_HYPER, 3
type InterpType int

const (
	INTERP_NEAREST InterpType = iota
	INTERP_TILES
	INTERP_BILINEAR
	INTERP_HYPER
)

// ENUM_NAME = GDK_PIXBUF_ROTATE_NONE, 0
// ENUM_NAME = GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE, 90
// ENUM_NAME = GDK_PIXBUF_ROTATE_UPSIDEDOWN, 180
// ENUM_NAME = GDK_PIXBUF_ROTATE_CLOCKWISE, 270
type Rotation int

const (
	ROTATE_NONE             Rotation = 0
	ROTATE_COUNTERCLOCKWISE Rotation = 90
	ROTATE_UPSIDEDOWN       Rotation = 180
	ROTATE_CLOCKWISE        Rotation = 270
)