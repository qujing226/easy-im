package rpcserver

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	zerr "github.com/zeromicro/x/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp any, err error) {
	resp, err = handler(ctx, req)
	if err == nil {
		return resp, err
	}

	causeErr := errors.Cause(err)
	var e *zerr.CodeMsg
	if errors.As(causeErr, &e) {
		err = status.Error(codes.Code(e.Code), e.Msg)
	}

	logx.WithContext(ctx).Errorf("『 RPC SRV ERR 』%v", err)
	return
}
