package websocket

import "time"

type Options func(opt *option)

type option struct {
	Authentication
	patten string

	maxConnectionIdle time.Duration
}

func newOption(opts ...Options) option {
	o := option{
		Authentication:    new(authentication),
		maxConnectionIdle: defaultMaxConnectionIdle,
		patten:            "/ws",
	}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func WithAuthentication(authentication Authentication) Options {
	return func(opt *option) {
		opt.Authentication = authentication
	}
}

func WithHandlerPattern(pattern string) Options {
	return func(opt *option) {
		opt.patten = pattern
	}
}

func WithServerMaxConnectionIdle(maxConnection time.Duration) Options {
	return func(opt *option) {
		if maxConnection > 0 {
			opt.maxConnectionIdle = maxConnection
		}
	}
}
