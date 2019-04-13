package sugar

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math"
	"runtime"
	"strings"
	"unsafe"

	"github.com/jopbrown/gtk-sugar/util/must"
	"github.com/pkg/errors"
)

var nativeEndian binary.ByteOrder

func init() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		nativeEndian = binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		nativeEndian = binary.BigEndian
	default:
		panic("unable to determine native endian")
	}
}

func (sugar *sugar) SugarUnpack(format, b64 string) RespFields {
	// gtk_server_unpack is not work correctly before 2.4.5
	// try to decode in client side
	bin := must.Bytes(base64.StdEncoding.DecodeString(b64))
	start := 0
	end := 0
	argTypes := strings.Split(format, "%")[1:]
	fields := make(RespFields, len(argTypes))
	var v interface{}
	for i, argType := range argTypes {
		switch argType {
		case "i":
			end = start + 4
			v = int32(nativeEndian.Uint32(bin[start:end]))
		case "s":
			end = start + 2
			v = int16(nativeEndian.Uint16(bin[start:end]))
		case "c":
			end = start + 1
			v = bin[start]
		case "l":
			if runtime.GOOS == "windows" {
				end = start + 2
				v = int32(nativeEndian.Uint32(bin[start:end]))
			} else {
				end = start + 4
				v = int64(nativeEndian.Uint64(bin[start:end]))
			}
		case "f":
			end = start + 2
			v = math.Float32frombits(nativeEndian.Uint32(bin[start:end]))
		case "d":
			end = start + 4
			v = math.Float64frombits(nativeEndian.Uint64(bin[start:end]))
		default:
			panic(errors.Errorf("unknow pack format %s in %s", argType, format))
		}

		fields[i] = Response(fmt.Sprintf("%v", v))
		start = end
	}

	return fields
}
