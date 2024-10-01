package test

import (
	"context"
	"easy-chat/apps/social/rpc/internal/logic"
	"easy-chat/apps/social/rpc/social"
	"testing"
)

func TestFriendListLogic_FriendList(t *testing.T) {
	type args struct {
		in *social.FriendListReq
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "name1",
			args: args{
				in: &social.FriendListReq{
					UserId: "1838501776039350272",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "name2",
			args: args{
				in: &social.FriendListReq{
					UserId: "1840324474826657792",
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewFriendListLogic(context.Background(), svcCtx)
			got, err := l.FriendList(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendList() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				l.Logger.Infof("got: %v", got)
			}
		})
	}
}
