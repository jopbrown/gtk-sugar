package sugar

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"

	"github.com/pkg/errors"
)

type Conn interface {
	Open() error
	io.ReadWriteCloser
}

type mutexConn struct {
	inner Conn
	mux   sync.Mutex
}

func wrapMutex(conn Conn) Conn {
	return &mutexConn{inner: conn}
}

func (conn *mutexConn) Open() error {
	conn.mux.Lock()
	defer conn.mux.Unlock()
	return conn.inner.Open()
}

func (conn *mutexConn) Write(p []byte) (n int, err error) {
	conn.mux.Lock()
	defer conn.mux.Unlock()
	return conn.inner.Write(p)
}

func (conn *mutexConn) Read(p []byte) (n int, err error) {
	conn.mux.Lock()
	defer conn.mux.Unlock()
	return conn.inner.Read(p)
}

func (conn *mutexConn) Close() error {
	conn.mux.Lock()
	defer conn.mux.Unlock()
	return conn.inner.Close()
}

type ConnAgent func(*client)

type stdinConn struct {
	w io.Writer
	r io.Reader
}

func (conn *stdinConn) Open() error {
	return nil
}

func (conn *stdinConn) Write(p []byte) (n int, err error) {
	return conn.w.Write(p)
}

func (conn *stdinConn) Read(p []byte) (n int, err error) {
	return conn.r.Read(p)
}

func (conn *stdinConn) Close() error {
	return nil
}

func ConnStdin() ConnAgent {
	return func(c *client) {
		w, err := c.cmd.StdinPipe()
		if err != nil {
			panic(fmt.Sprintf("FATAL:%+v", errors.Wrap(err, "invalid operation to create pipe for stdin")))
		}

		r, err := c.cmd.StdoutPipe()
		if err != nil {
			panic(fmt.Sprintf("FATAL:%+v", errors.Wrap(err, "invalid operation to create pipe for stdout")))
		}

		c.cmd.Args = append(c.cmd.Args, "-stdin")
		c.conn = &stdinConn{w: w, r: r}
	}
}

type fifoConn struct {
	inFileName  string
	outFileName string
	w           io.WriteCloser
	r           io.ReadCloser
}

func (conn *fifoConn) Open() error {
	w, err := os.OpenFile(conn.outFileName, os.O_WRONLY, 0666)
	if err != nil {
		return errors.Wrapf(err, "fifo: unable to open as output file for %v", conn.outFileName)
	}

	r, err := os.OpenFile(conn.inFileName, os.O_RDONLY, 0666)
	if err != nil {
		return errors.Wrapf(err, "fifo: unable to open as input file for %v", conn.inFileName)
	}

	conn.w = w
	conn.r = r

	return nil
}

func (conn *fifoConn) Write(p []byte) (n int, err error) {
	return conn.w.Write(p)
}

func (conn *fifoConn) Read(p []byte) (n int, err error) {
	return conn.r.Read(p)
}

func (conn *fifoConn) Close() (err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "unable to close fifo")
		}
	}()

	defer func() {
		errw := conn.w.Close()
		if errw != nil {
			err = errw
		}
	}()

	defer func() {
		errr := conn.r.Close()
		if errr != nil {
			err = errr
		}
	}()

	return
}

func ConnFifo(fileName string) ConnAgent {
	return func(c *client) {
		if runtime.GOOS == "windows" {
			c.cmd.Args = append(c.cmd.Args, "-fifo")
			c.conn = &fifoConn{inFileName: WINDOWS_PIPEIN, outFileName: WINDOWS_PIPEOUT}
		} else {
			c.cmd.Args = append(c.cmd.Args, "-fifo="+fileName)
			c.conn = &fifoConn{inFileName: fileName, outFileName: fileName}
		}
	}
}
