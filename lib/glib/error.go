package glib

import sugar "github.com/jopbrown/gtk-sugar"

type GError struct {
	Domain  uint32
	Code    int
	Message string
}

func NewGErrorFromPointer(id string) *GError {
	err := &GError{}
	packer := sugar.NewBase64Packer(err)
	format := packer.Format()
	fields := Candy().ServerUnpackFromPointer(format, id)
	fields.MustUnmarshal(packer)

	err.Message = Candy().ServerStringFromPointer(fields[2].String())
	return err
}

func (err *GError) Error() string {
	return err.Message
}
