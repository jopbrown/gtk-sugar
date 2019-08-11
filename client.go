package sugar

import (
	"context"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/jopbrown/gtk-sugar/util/fs"
	"github.com/jopbrown/gtk-sugar/util/must"
	"github.com/pkg/errors"
)

const (
	_DEFAULT_BIN_PATH     = "gtk-server"
	_DEFAULT_CFG_FILENAME = "gtk-server.cfg"
	_WINDOWS_PIPEIN       = `\\.\pipe\in`
	_WINDOWS_PIPEOUT      = `\\.\pipe\out`
	_DEFAULT_START_DELAY  = 100 * time.Millisecond
)

type Client interface {
	Start() error
	StartAfter(delay time.Duration) error
	Stop() error
	Conn() io.ReadWriter
}

type client struct {
	cmd    *exec.Cmd
	cancel func()
	conn   Conn
	ctx    context.Context
}

func NewClient(connAgent ConnAgent, opts ...Option) Client {
	o := &Options{BinPath: _DEFAULT_BIN_PATH}
	for _, opt := range opts {
		opt(o)
	}

	return NewClientWithOptions(connAgent, o)
}

func NewClientWithOptions(connAgent ConnAgent, o *Options) Client {
	c := &client{}

	optCxt := o.Ctx
	if optCxt == nil {
		optCxt = context.Background()
	}

	ctx, cancel := context.WithCancel(optCxt)
	c.cmd = exec.CommandContext(ctx, o.BinPath)
	c.cancel = cancel

	applyCfgPath(o)
	if len(o.CfgPath) != 0 {
		c.cmd.Args = append(c.cmd.Args, "-cfg="+o.CfgPath)
	}

	if len(o.LogPath) != 0 {
		c.cmd.Args = append(c.cmd.Args, "-log="+o.LogPath)
	}

	if o.Debug {
		c.cmd.Args = append(c.cmd.Args, "-debug")
	}

	if len(o.Args) != 0 {
		c.cmd.Args = append(c.cmd.Args, o.Args...)
	}

	connAgent(c)

	if o.EnableMutex {
		c.conn = wrapMutex(c.conn)
	}

	c.ctx = o.Ctx

	return c
}

func (c *client) Start() error {
	return c.StartAfter(_DEFAULT_START_DELAY)
}

func (c *client) StartAfter(delay time.Duration) error {
	err := c.cmd.Start()
	if err != nil {
		return errors.Wrapf(err, "unable to start gtk-server with command: '%v %v'", c.cmd.Path, c.cmd.Args)
	}

	time.Sleep(delay)

	err = c.conn.Open()
	if err != nil {
		defer c.cancel()
		return errors.Wrap(err, "unable to open connect to gtk-server")
	}

	if c.ctx != nil {
		go func() {
			<-c.ctx.Done()
			c.conn.Close()
		}()
	}

	return nil
}

func (c *client) Stop() error {
	err := c.conn.Close()
	if err != nil {
		return errors.Wrap(err, "unable to close connect fot gtk-server")
	}

	c.cancel()

	return nil
}

func (c *client) Conn() io.ReadWriter {
	return c.conn
}

func applyCfgPath(o *Options) {
	if len(o.CfgPath) != 0 {
		return
	}

	cfgPath := o.CfgPath
	defer func() {
		o.CfgPath = cfgPath
	}()

	if fs.ExistsFile(cfgPath) {
		return
	}

	// look at pwd
	cfgPath = _DEFAULT_CFG_FILENAME
	if fs.ExistsFile(cfgPath) {
		return
	}

	// look at current executable dir
	cfgPath = filepath.Join(filepath.Dir(must.String(os.Executable())), _DEFAULT_CFG_FILENAME)
	if fs.ExistsFile(cfgPath) {
		return
	}

	// look at gtk-server executable dir
	binPath, err := exec.LookPath(o.BinPath)
	if err == nil {
		cfgPath = filepath.Join(filepath.Dir(binPath), _DEFAULT_CFG_FILENAME)
		if fs.ExistsFile(cfgPath) {
			return
		}
	}

	// look at /etc
	cfgPath = filepath.Join("/etc", _DEFAULT_CFG_FILENAME)
	if fs.ExistsFile(cfgPath) {
		return
	}

	// look at /usr/local/etc
	cfgPath = filepath.Join("/usr/local/etc", _DEFAULT_CFG_FILENAME)
	if fs.ExistsFile(cfgPath) {
		return
	}

	cfgPath = ""
}
