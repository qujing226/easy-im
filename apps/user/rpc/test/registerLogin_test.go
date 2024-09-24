package logic

import (
	"context"
	"easy-chat/apps/user/rpc/internal/config"
	"easy-chat/apps/user/rpc/internal/logic"
	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"testing"
)

var configFile = flag.String("f", "../etc/dev/user.yaml", "the config file")

var svcCtx *svc.ServiceContext

func init() {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx = svc.NewServiceContext(c)
}

func TestRegisterLogic_Register(t *testing.T) {
	type args struct {
		in *user.RegisterReq
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "1", args: args{in: &user.RegisterReq{
				Phone:    "17309710356",
				Password: "yining2024",
				Nickname: "Qwyk",
				Avatar:   "",
				Sex:      0,
			}}, want: true, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewRegisterLogic(context.Background(), svcCtx)
			got, err := l.Register(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want {
				t.Log(tt.name, got)
			}
		})
	}
}

func TestLoginLogic_Login(t *testing.T) {
	type args struct {
		in *user.LoginReq
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "1", args: args{in: &user.LoginReq{
				Phone:    "17309710356",
				Password: "yining2024",
			}}, want: true, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewLoginLogic(context.Background(), svcCtx)
			got, err := l.Login(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				t.Log(tt.name, got)
			}
		})
	}
}
