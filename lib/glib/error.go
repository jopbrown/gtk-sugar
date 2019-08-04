package glib

import sugar "github.com/jopbrown/gtk-sugar"

type Error struct {
	Domain  uint32
	Code    int
	Message string
}

func NewErrorFromPointer(id string) *Error {
	err := &Error{}
	packer := sugar.NewBase64Packer(err)
	format := packer.Format()
	fields := Candy().ServerUnpackFromPointer(format, id)
	fields.MustUnmarshal(packer)

	err.Message = Candy().ServerStringFromPointer(fields[2].String())
	return err
}

func (err *Error) Error() string {
	return err.Message
}
