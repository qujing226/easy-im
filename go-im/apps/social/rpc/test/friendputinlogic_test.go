package test

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/internal/config"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/internal/logic"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"testing"
	"time"
)

var configFile = flag.String("f", "../etc/dev/social.yaml", "the config file")

var svcCtx *svc.ServiceContext

func init() {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx = svc.NewServiceContext(c)
}

func TestFriendPutInLogic_FriendPutIn(t *testing.T) {
	type args struct {
		in *social.FriendPutInReq
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "admin and ming",
			args: args{
				in: &social.FriendPutInReq{
					UserId:  "1843306294396588032",
					ReqUid:  "1843306302982328320",
					ReqMsg:  "hello",
					ReqTime: time.Now().Unix(),
				},
			}, want: true, wantErr: false,
		},
		{
			name: "admin and hong",
			args: args{
				in: &social.FriendPutInReq{
					UserId:  "1843306294396588032",
					ReqUid:  "1843306311148638208",
					ReqMsg:  "hello",
					ReqTime: time.Now().Unix(),
				},
			}, want: true, wantErr: false,
		},
		{
			name: "admin and wang",
			args: args{
				in: &social.FriendPutInReq{
					UserId:  "1843306294396588032",
					ReqUid:  "1843306319776321536",
					ReqMsg:  "hello~!",
					ReqTime: time.Now().Unix(),
				},
			}, want: true, wantErr: false,
		},
		{
			name: "ming and wang, refused",
			args: args{
				in: &social.FriendPutInReq{
					UserId:  "1843306302982328320",
					ReqUid:  "1843306319776321536",
					ReqMsg:  "son of bitch!",
					ReqTime: time.Now().Unix(),
				},
			}, want: true, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewFriendPutInLogic(context.Background(), svcCtx)
			got, err := l.FriendPutIn(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendPutIn() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				t.Logf("FriendPutIn() got = %v", got)
			}
		})
	}
}
