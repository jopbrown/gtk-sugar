// Package sugar is client side of gtk-server
package sugar

import (
	"io"
)

// Sugar provide a high level api to communicate with gtk-server.
type Sugar interface {
	Guifyer
	ServerConnect(widget, signal, description string) string
	ServerDisconnect(widget, description string)
	ServerExit()
	ServerEcho(msg string) string
	ServerDefine(define string)
	ServerRedefine(define string)
	ServerRequire(libName string) bool
	ServerPack(format string, values ...interface{}) string
	ServerPackStruct(s interface{}) string
	ServerUnpack(format, b64 string) RespFields
	ServerUnpackFromPointer(format, id string) RespFields
	ServerStringFromPointer(id string) string
	ServerDataFormat(format string)
	ServerCallback(t ServerCallbackType) string
	ServerCallbackValue(argIdx int, argType CallValueType) Response
	ServerOpaque() string
	ServerKey() int
	ServerKeyState() int
	ServerMouse() *Mouse
}

type sugar struct {
	Guifyer
}

// NewSugar get a Sugar from connection of gtk-server
func NewSugar(conn io.ReadWriter) Sugar {
	return &sugar{Guifyer: NewGuifyer(conn)}
}

type CallValueType int

const (
	CALLBACK_VALUE_TYPE_INT CallValueType = iota
	CALLBACK_VALUE_TYPE_STRING
)

func (t CallValueType) String() string {
	switch t {
	case CALLBACK_VALUE_TYPE_INT:
		return "INT"
	case CALLBACK_VALUE_TYPE_STRING:
		return "STRING"
	}

	return ""
}

func (sugar *sugar) ServerConnect(widget, signal, description string) string {
	return sugar.Guify("gtk_server_connect", widget, signal, description).String()
}

func (sugar *sugar) ServerDisconnect(widget, description string) {
	sugar.Guify("gtk_server_disconnect", widget, description)
}

func (sugar *sugar) ServerExit() {
	sugar.Guify("gtk_server_exit")
}

func (sugar *sugar) ServerEcho(msg string) string {
	res := sugar.Guify("gtk_server_echo", msg)
	return res.String()
}

func (sugar *sugar) ServerDefine(define string) {
	sugar.Guifyf("%s %s", "gtk_server_define", define)
}

func (sugar *sugar) ServerRedefine(define string) {
	sugar.Guifyf("%s %s", "gtk_server_redefine", define)
}

func (sugar *sugar) ServerRequire(libName string) bool {
	return sugar.Guify("gtk_server_require", libName).MustBool()
}

func (sugar *sugar) ServerPack(format string, values ...interface{}) string {
	return sugar.Guify("gtk_server_pack", format, Args(values)).String()
}

func (sugar *sugar) ServerPackStruct(s interface{}) string {
	packer := NewBase64Packer(s)
	return sugar.Guify("gtk_server_pack", packer.Format(), packer.Args()).String()
}

func (sugar *sugar) ServerUnpack(format, b64 string) RespFields {
	return sugar.Guify("gtk_server_unpack", format, b64).Fields()
}

func (sugar *sugar) ServerUnpackFromPointer(format, id string) RespFields {
	return sugar.Guify("gtk_server_unpack_from_pointer", format, id).Fields()
}

func (sugar *sugar) ServerStringFromPointer(id string) string {
	return sugar.Guify("gtk_server_string_from_pointer", id).String()
}

func (sugar *sugar) ServerDataFormat(format string) {
	sugar.Guify("gtk_server_data_format", format)
}

type ServerCallbackType int

const (
	SERVER_CALLBACK_NO_ITERATION = ServerCallbackType(iota)
	SERVER_CALLBACK_WAIT
	SERVER_CALLBACK_UPDATE
)

func (sugar *sugar) ServerCallback(t ServerCallbackType) string {
	resp := sugar.Guify("gtk_server_callback", t)
	return resp.String()
}

func (sugar *sugar) ServerCallbackValue(argIdx int, argType CallValueType) Response {
	return sugar.Guify("gtk_server_callback_value", argIdx, argType.String())
}

func (sugar *sugar) ServerOpaque() string {
	res := sugar.Guify("gtk_server_opaque")

	return res.String()
}

func (sugar *sugar) ServerKey() int {
	return sugar.Guify("gtk_server_key").MustInt()
}

func (sugar *sugar) ServerKeyState() int {
	return sugar.Guify("gtk_server_state").MustInt()
}

type ServerMouseStatus int

const (
	SERVER_MOUSE_STATUS_LEFT ServerMouseStatus = iota + 1
	SERVER_MOUSE_STATUS_MIDDLE
	SERVER_MOUSE_STATUS_RIGHT
)

type ServerMouseScroll int

const (
	SERVER_MOUSE_SCROLL_UP ServerMouseScroll = iota
	SERVER_MOUSE_SCROLL_DOWN
	SERVER_MOUSE_SCROLL_LEFT
	SERVER_MOUSE_SCROLL_RIGHT
)

type Mouse struct {
	X, Y   int
	Status ServerMouseStatus
	Scroll ServerMouseScroll
}

func (sugar *sugar) ServerMouse() *Mouse {
	x := sugar.Guify("gtk_server_mouse", 0).MustInt()
	y := sugar.Guify("gtk_server_mouse", 1).MustInt()
	status := sugar.Guify("gtk_server_mouse", 2).MustInt()
	scroll := sugar.Guify("gtk_server_mouse", 3).MustInt()

	return &Mouse{X: x, Y: y, Status: ServerMouseStatus(status), Scroll: ServerMouseScroll(scroll)}
}
