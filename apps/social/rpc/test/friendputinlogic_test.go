package test

import (
	"context"
	"easy-chat/apps/social/rpc/internal/config"
	"easy-chat/apps/social/rpc/internal/logic"
	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"
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
			name: "test1",
			args: args{
				in: &social.FriendPutInReq{
					UserId:  "1842843962742673408",
					ReqUid:  "1842843971269693440",
					ReqMsg:  "hello",
					ReqTime: time.Now().Unix(),
				},
			}, want: true, wantErr: false,
		},
		{
			name: "test1",
			args: args{
				in: &social.FriendPutInReq{
					UserId:  "1842843962742673408",
					ReqUid:  "1842843978811052032",
					ReqMsg:  "hello",
					ReqTime: time.Now().Unix(),
				},
			}, want: true, wantErr: false,
		},
		{
			name: "test2",
			args: args{
				in: &social.FriendPutInReq{
					UserId:  "1842843962742673408",
					ReqUid:  "1842843985727459328",
					ReqMsg:  "hello~!",
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
