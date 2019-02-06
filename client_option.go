package sugar

import "context"

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
