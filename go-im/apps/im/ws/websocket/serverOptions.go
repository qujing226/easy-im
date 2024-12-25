package websocket

import "time"

type ServerOption func(opt *serverOption)

type serverOption struct {
	Authentication
	ack        AckType
	ackTimeout time.Duration
	patten     string

	maxConnectionIdle time.Duration

	// 设置并发发送数据的量级
	concurrency int
}

func newOption(opts ...ServerOption) serverOption {
	o := serverOption{
		Authentication:    new(authentication),
		maxConnectionIdle: defaultMaxConnectionIdle,
		ackTimeout:        defaultAckTimeout,
		patten:            "/ws",
		concurrency:       defaultConCurrency,
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

func WithHandlerPatten(pattern string) ServerOption {
	return func(opt *serverOption) {
		opt.patten = pattern
	}
}

func WithServerAck(ack AckType) ServerOption {
	return func(opt *serverOption) {
		opt.ack = ack
	}
}

func WithServerMaxConnectionIdle(maxConnection time.Duration) ServerOption {
	return func(opt *serverOption) {
		if maxConnection > 0 {
			opt.maxConnectionIdle = maxConnection
		}
	}
}
