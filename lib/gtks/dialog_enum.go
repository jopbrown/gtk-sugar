package gtks

// ENUM_NAME = GTK_DIALOG_MODAL, 1
// ENUM_NAME = GTK_DIALOG_DESTROY_WITH_PARENT, 2
// ENUM_NAME = GTK_DIALOG_NO_SEPARATOR, 4
type DialogFlags int

const (
	DIALOG_MODAL DialogFlags = 1 << iota
	DIALOG_DESTROY_WITH_PARENT
	DIALOG_NO_SEPARATOR
)

// ENUM_NAME = GTK_RESPONSE_NONE, -1
// ENUM_NAME = GTK_RESPONSE_REJECT, -2
// ENUM_NAME = GTK_RESPONSE_ACCEPT, -3
// ENUM_NAME = GTK_RESPONSE_DELETE_EVENT, -4
// ENUM_NAME = GTK_RESPONSE_OK, -5
// ENUM_NAME = GTK_RESPONSE_CANCEL, -6
// ENUM_NAME = GTK_RESPONSE_CLOSE, -7
// ENUM_NAME = GTK_RESPONSE_YES, -8
// ENUM_NAME = GTK_RESPONSE_NO, -9
// ENUM_NAME = GTK_RESPONSE_APPLY, -10
// ENUM_NAME = GTK_RESPONSE_HELP, -11
type ResponseType int

const (
	RESPONSE_NONE         ResponseType = -1
	RESPONSE_REJECT                    = -2
	RESPONSE_ACCEPT                    = -3
	RESPONSE_DELETE_EVENT              = -4
	RESPONSE_OK                        = -5
	RESPONSE_CANCEL                    = -6
	RESPONSE_CLOSE                     = -7
	RESPONSE_YES                       = -8
	RESPONSE_NO                        = -9
	RESPONSE_APPLY                     = -10
	RESPONSE_HELP                      = -11
)

// ENUM_NAME = GTK_MESSAGE_INFO, 0
// ENUM_NAME = GTK_MESSAGE_WARNING, 1
// ENUM_NAME = GTK_MESSAGE_QUESTION, 2
// ENUM_NAME = GTK_MESSAGE_ERROR, 3
// ENUM_NAME = GTK_MESSAGE_OTHER, 4
type MessageType int

const (
	MESSAGE_INFO MessageType = iota
	MESSAGE_WARNING
	MESSAGE_QUESTION
	MESSAGE_ERROR
	MESSAGE_OTHER
)

// ENUM_NAME = GTK_BUTTONS_NONE, 0
// ENUM_NAME = GTK_BUTTONS_OK, 1
// ENUM_NAME = GTK_BUTTONS_CLOSE, 2
// ENUM_NAME = GTK_BUTTONS_CANCEL, 3
// ENUM_NAME = GTK_BUTTONS_YES_NO, 4
// ENUM_NAME = GTK_BUTTONS_OK_CANCEL, 5
type ButtonsType int

const (
	BUTTONS_NONE ButtonsType = iota
	BUTTONS_OK
	BUTTONS_CLOSE
	BUTTONS_CANCEL
	BUTTONS_YES_NO
	BUTTONS_OK_CANCEL
)
