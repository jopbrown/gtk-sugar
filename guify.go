package sugar

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

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
	command := formatCommand(cmd, args...)
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
