package sugar

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jopbrown/gtk-sugar/util/fs"
	"github.com/jopbrown/gtk-sugar/util/must"
)

type Options struct {
	BinPath     string
	CfgPath     string
	LogPath     string
	Args        []string
	Debug       bool
	EnableMutex bool
	Ctx         context.Context
}

type Option func(*Options)

func WithBinPath(path string) Option {
	return func(o *Options) {
		o.BinPath = path
	}
}

func WithCfgPath(path string) Option {
	return func(o *Options) {
		o.CfgPath = path
	}
}

func WithLogPath(path string) Option {
	return func(o *Options) {
		o.LogPath = path
	}
}

func WithArgs(args ...string) Option {
	return func(o *Options) {
		o.Args = append(o.Args, args...)
	}
}

func WithDebug(enable bool) Option {
	return func(o *Options) {
		o.Debug = enable
	}
}

func WithMutex(enable bool) Option {
	return func(o *Options) {
		o.EnableMutex = enable
	}
}

func WithContext(ctx context.Context) Option {
	return func(o *Options) {
		o.Ctx = ctx
	}
}

func LookupCfg() Option {
	return func(o *Options) {
		// look at current setting
		cfgPath := o.CfgPath
		if fs.ExistsFile(cfgPath) {
			return
		}

		defer func() { o.CfgPath = cfgPath }()

		// look at pwd
		cfgPath = DEFAULT_CFG_FILENAME
		if fs.ExistsFile(cfgPath) {
			return
		}

		// look at current executable dir
		cfgPath = filepath.Join(filepath.Dir(must.String(os.Executable())), DEFAULT_CFG_FILENAME)
		if fs.ExistsFile(cfgPath) {
			return
		}

		// look at gtk-server executable dir
		binPath, err := exec.LookPath(o.BinPath)
		if err == nil {
			cfgPath = filepath.Join(filepath.Dir(binPath), DEFAULT_CFG_FILENAME)
			if fs.ExistsFile(cfgPath) {
				return
			}
		}

		// look at /etc
		cfgPath = filepath.Join("/etc", DEFAULT_CFG_FILENAME)
		if fs.ExistsFile(cfgPath) {
			return
		}

		// look at /usr/local/etc
		cfgPath = filepath.Join("/usr/local/etc", DEFAULT_CFG_FILENAME)
		if fs.ExistsFile(cfgPath) {
			return
		}

		cfgPath = ""
	}
}
