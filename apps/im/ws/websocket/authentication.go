package websocket

import (
	"fmt"
	"net/http"
	"time"
)

type Authentication interface {
	Auth(w http.ResponseWriter, r *http.Request) bool
	UserId(r *http.Request) string
}

type authentication struct {
}

func (a *authentication) Auth(w http.ResponseWriter, r *http.Request) bool {
	return true
}

func (a *authentication) UserId(r *http.Request) string {
	query := r.URL.Query()
	if query != nil && query["userId"] != nil {
		return fmt.Sprintf("%v", query["userId"])
	}
	return fmt.Sprintf("%v", time.Now().UnixMilli())
}
