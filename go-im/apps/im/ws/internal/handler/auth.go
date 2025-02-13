package handler

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/token"
	"net/http"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/internal/svc"
	"github.com/peninsula12/easy-im/go-im/pkg/ctxdata"
)

type JwtAuth struct {
	svc    *svc.ServiceContext
	parser *token.TokenParser
	logx.Logger
}

func NewJwtAuth(svc *svc.ServiceContext) *JwtAuth {
	return &JwtAuth{
		svc:    svc,
		parser: token.NewTokenParser(),
		Logger: logx.WithContext(context.Background()),
	}
}
func (j *JwtAuth) Auth(w http.ResponseWriter, r *http.Request) bool {
	if tok := r.Header.Get("sec-websocket-protocol"); tok != "" {
		r.Header.Set("Authorization", tok)
	}
	tok, err := j.parser.ParseToken(r, j.svc.Config.JwtAuth.AccessSecret, "")
	if err != nil {
		return false

	}
	if !tok.Valid {
		return false
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	*r = *r.WithContext(context.WithValue(r.Context(), ctxdata.IdentityKey, claims[ctxdata.IdentityKey]))
	return true
}

func (j *JwtAuth) UserId(r *http.Request) string {
	return ctxdata.GetUId(r.Context())
}
