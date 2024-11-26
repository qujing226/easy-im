package websocket

import "time"

const (
	defaultMaxConnectionIdle = time.Duration(60 * time.Second)
	defaultAckTimeout        = time.Duration(30 * time.Second)
	defaultConCurrency       = 30
)
