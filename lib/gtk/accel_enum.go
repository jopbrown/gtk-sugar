package gtk

// ENUM_NAME = GTK_ACCEL_VISIBLE, 1
// ENUM_NAME = GTK_ACCEL_LOCKED, 2
// ENUM_NAME = GTK_ACCEL_MASK, 4
type AccelFlags int

const (
	ACCEL_VISIBLE AccelFlags = 1 << iota
	ACCEL_LOCKED
	ACCEL_MASK
)
