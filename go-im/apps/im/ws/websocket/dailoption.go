package websocket

import "net/http"

type DailOption func(option *dailOption)

type dailOption struct {
	patten string
	header http.Header
}

func newDailOptions(opts ...DailOption) dailOption {
	o := dailOption{
		patten: "/ws",
		header: nil,
	}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func WithClientPatten(patten string) DailOption {
	return func(option *dailOption) {
		option.patten = patten
	}
}
func WithClientHead(head http.Header) DailOption {
	return func(option *dailOption) {
		option.header = head
	}
}
