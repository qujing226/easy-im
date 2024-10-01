package test

import (
	"context"
	"easy-chat/apps/social/rpc/internal/config"
	"easy-chat/apps/social/rpc/internal/logic"
	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"reflect"
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
		want    *social.FriendPutInResp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				in: &social.FriendPutInReq{
					UserId:  "1838501776039350272",
					ReqUid:  "1840324474826657792",
					ReqMsg:  "hello",
					ReqTime: time.Now().Unix(),
				},
			}, want: &social.FriendPutInResp{}, wantErr: false,
		},
		{
			name: "test1",
			args: args{
				in: &social.FriendPutInReq{
					UserId:  "1838501776039350272",
					ReqUid:  "1840324474826657792",
					ReqMsg:  "hello~!",
					ReqTime: time.Now().Unix(),
				},
			}, want: &social.FriendPutInResp{}, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewFriendPutInLogic(context.Background(), svcCtx)
			got, err := l.FriendPutIn(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendPutIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendPutIn() got = %v, want %v", got, tt.want)
			}
		})
	}
}
