package job

import "time"

const (
	DefaultRetryJetLag  = time.Second
	DefaultRetryTimeout = time.Second * 2
	DefaultRetryNums    = 3
)
