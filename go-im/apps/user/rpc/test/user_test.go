package logic

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/internal/config"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/internal/logic"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/user"
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
			name: "name1", args: args{in: &user.RegisterReq{
				Phone:    "11122223333",
				Password: "yujie2024",
				Nickname: "peninsula",
				Avatar:   "",
				Sex:      1,
			}}, want: true, wantErr: false,
		},
		{
			name: "name2", args: args{in: &user.RegisterReq{
				Phone:    "17309710356",
				Password: "yining2024",
				Nickname: "Qwyk",
				Avatar:   "",
				Sex:      1,
			}}, want: true, wantErr: false,
		},
		{
			name: "name3", args: args{in: &user.RegisterReq{
				Phone:    "17344995006",
				Password: "admin",
				Nickname: "admin",
				Avatar:   "",
				Sex:      1,
			}}, want: true, wantErr: false,
		},
		{
			name: "name4", args: args{in: &user.RegisterReq{
				Phone:    "11122223333",
				Password: "123456",
				Nickname: "xiaoming",
				Avatar:   "",
				Sex:      1,
			}}, want: true, wantErr: false,
		},
		{
			name: "name5", args: args{in: &user.RegisterReq{
				Phone:    "22233334444",
				Password: "123456",
				Nickname: "xiaohong",
				Avatar:   "",
				Sex:      0,
			}}, want: true, wantErr: false,
		},
		{
			name: "name6", args: args{in: &user.RegisterReq{
				Phone:    "33344445555",
				Password: "123456",
				Nickname: "xiaowang",
				Avatar:   "",
				Sex:      1,
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
			}
			if !tt.wantErr {
				t.Log(tt.name, got, err)
			}
		})
	}
}
func TestGetUserInfoLogic_GetUserInfo(t *testing.T) {
	type args struct {
		in *user.GetUserInfoReq
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "1", args: args{in: &user.GetUserInfoReq{
				User: "1838501776039350272",
			}}, want: true, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewGetUserInfoLogic(context.Background(), svcCtx)
			got, err := l.GetUserInfo(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				t.Logf("GetUserInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindUserLogic_FindUser(t *testing.T) {
	type args struct {
		in *user.FindUserReq
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "1", args: args{in: &user.FindUserReq{
				Name:  "Qwyk",
				Phone: "",
				Ids:   []string{},
			}}, want: true, wantErr: false,
		},
		{
			name: "1", args: args{in: &user.FindUserReq{
				Name:  "",
				Phone: "17309710356",
				Ids:   []string{},
			}}, want: true, wantErr: false,
		},
		{
			name: "1", args: args{in: &user.FindUserReq{
				Name:  "",
				Phone: "",
				Ids:   []string{"1838501776039350272", "1840324474826657792", "1840643788553326592"},
			}}, want: true, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewFindUserLogic(context.Background(), svcCtx)
			got, err := l.FindUser(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				t.Logf("FindUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
