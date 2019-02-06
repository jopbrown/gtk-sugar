package sugar

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/jopbrown/gtk-sugar/pkg/gfmt"
	"github.com/pkg/errors"
)

const (
	INVALID_RESPONSE = "-1"
)

// Guier provide low level api to communicate with gtk-server.
type Guier interface {
	// Gui send raw command to gtk-server and return raw response.
	Gui(command string) (string, error)
}

type guier struct {
	conn io.ReadWriter
	buf  *bufio.Reader
}

// NewGuier get a Guier from connection of gtk-server
func NewGuier(conn io.ReadWriter) Guier {
	var g Guier = &guier{conn: conn, buf: bufio.NewReader(conn)}
	return g
}

// Gui send raw command to gtk-server and return raw response.
func (guier *guier) Gui(command string) (string, error) {
	_, err := io.WriteString(guier.conn, command)
	if err != nil {
		return "", errors.Wrapf(err, "unable to send command:%v", command)
	}

	b, err := guier.buf.ReadBytes('\n')
	if err != nil && errors.Cause(err) != io.EOF {
		return "", errors.Wrapf(err, "unable to recived from command:%v", command)
	}

	b = bytes.TrimSpace(b)
	res := string(b)

	if res == INVALID_RESPONSE {
		return "", errors.Errorf("recived invalid response -1 from command:%v", command)
	}

	return res, nil
}

// Guifyer provide middle level api to communicate with gtk-server.
type Guifyer interface {
	// Guify format the cmd and args to raw command and send to gtk-server
	// return Response type for further parsing
	Guify(cmd string, args ...interface{}) Response
	// Guifyf format command as printf style and send to gtk-server
	// return Response type for further parsing
	Guifyf(format string, v ...interface{}) Response
}

type guifyer struct {
	guier Guier
}

// NewGuifyer get a Guifyer from connection of gtk-server
func NewGuifyer(conn io.ReadWriter) Guifyer {
	return &guifyer{guier: NewGuier(conn)}
}

func (guifyer *guifyer) guify(cmd string, args ...interface{}) (Response, error) {
	command := gfmt.FormatCommand(cmd, args...)
	res, err := guifyer.guier.Gui(command)
	return Response(res), err
}

func (guifyer *guifyer) Guify(cmd string, args ...interface{}) Response {
	res, err := guifyer.guify(cmd, args...)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return res
}

func (guifyer *guifyer) guifyf(format string, v ...interface{}) (Response, error) {
	command := fmt.Sprintf(format, v...)
	res, err := guifyer.guier.Gui(command)
	return Response(res), err
}

func (guifyer *guifyer) Guifyf(format string, v ...interface{}) Response {
	res, err := guifyer.guifyf(format, v...)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	return res
}

type Varargs []interface{}

func (args Varargs) String() string {
	var sb strings.Builder
	for _, arg := range args {
		sb.WriteByte(' ')

		switch v := arg.(type) {
		case nil:
			sb.WriteString(gfmt.STRING_NIL)
		case string:
			sb.WriteString("s:")
			sb.WriteString(v)
		case int, int8, int16, int32, uint, uint8, uint16, uint32:
			sb.WriteString("i:")
			sb.WriteString(fmt.Sprintf("%v", v))
		case int64, uint64:
			sb.WriteString("l:")
			sb.WriteString(fmt.Sprintf("%v", v))
		case float32:
			sb.WriteString("f:")
			sb.WriteString(fmt.Sprintf("%v", v))
		case float64:
			sb.WriteString("d:")
			sb.WriteString(fmt.Sprintf("%v", v))
		case bool:
			sb.WriteString("b:")
			if v {
				sb.WriteString(gfmt.STRING_BOOL_TRUE)
			} else {
				sb.WriteString(gfmt.STRING_BOOL_FALSE)
			}
		case CandyWrapper:
			sb.WriteString("w:")
			sb.WriteString(v.ID())
		default:
			sb.WriteString("p:")
			sb.WriteString(fmt.Sprintf("%v", v))
		}
	}

	return sb.String()
}

// Response of gtk-server
type Response string

func (resp Response) String() string {
	return string(resp)
}

func (resp Response) Unquote() (string, error) {
	return gfmt.Unquote(resp.String())
}

func (resp Response) MustUnquote() string {
	return gfmt.MustUnquote(resp.String())
}

func (resp Response) Bool() (bool, error) {
	return gfmt.Bool(resp.String())
}

func (resp Response) MustBool() bool {
	return gfmt.MustBool(resp.String())
}

func (resp Response) Int64() (int64, error) {
	return gfmt.Int64(resp.String())
}

func (resp Response) MustInt64() int64 {
	return gfmt.MustInt64(resp.String())
}

func (resp Response) Int32() (int32, error) {
	return gfmt.Int32(resp.String())
}

func (resp Response) MustInt32() int32 {
	return gfmt.MustInt32(resp.String())
}

func (resp Response) Int() (int, error) {
	return gfmt.Int(resp.String())
}

func (resp Response) MustInt() int {
	return gfmt.MustInt(resp.String())
}

func (resp Response) Float64() (float64, error) {
	return gfmt.Float64(resp.String())
}

func (resp Response) MustFloat64() float64 {
	return gfmt.MustFloat64(resp.String())
}

func (resp Response) Fields() []Response {
	fields := strings.Fields(resp.String())
	resps := make([]Response, len(fields))
	for i, field := range fields {
		resps[i] = Response(field)
	}

	return resps
}
