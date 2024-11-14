package websocket

import "time"

type ServerOption func(opt *serverOption)

type serverOption struct {
	Authentication
	patten string

	maxConnectionIdle time.Duration
}

func newOption(opts ...ServerOption) serverOption {
	o := serverOption{
		Authentication:    new(authentication),
		maxConnectionIdle: defaultMaxConnectionIdle,
		patten:            "/ws",
	}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func WithAuthentication(authentication Authentication) ServerOption {
	return func(opt *serverOption) {
		opt.Authentication = authentication
	}
}

func WithHandlerPattern(pattern string) ServerOption {
	return func(opt *serverOption) {
		opt.patten = pattern
	}
}

func WithServerMaxConnectionIdle(maxConnection time.Duration) ServerOption {
	return func(opt *serverOption) {
		if maxConnection > 0 {
			opt.maxConnectionIdle = maxConnection
		}
	}
}
