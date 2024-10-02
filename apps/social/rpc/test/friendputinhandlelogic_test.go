package test

import (
	"context"
	"easy-chat/apps/social/rpc/internal/logic"
	"easy-chat/apps/social/rpc/social"
	"reflect"
	"testing"
)

func TestFriendPutInHandleLogic_FriendPutInHandle(t *testing.T) {
	type args struct {
		in *social.FriendPutInHandleReq
	}
	tests := []struct {
		name    string
		args    args
		want    *social.FriendPutInHandleResp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "name1",
			args: args{
				in: &social.FriendPutInHandleReq{
					FriendReqId:  "1841490323889459200",
					UserId:       "1841486243565670400",
					HandleResult: 2,
				},
			},
			want:    &social.FriendPutInHandleResp{},
			wantErr: false,
		},
		{
			name: "name1",
			args: args{
				in: &social.FriendPutInHandleReq{
					FriendReqId:  "1841490325604929536",
					UserId:       "1841486293704380416",
					HandleResult: 3,
				},
			},
			want:    &social.FriendPutInHandleResp{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewFriendPutInHandleLogic(context.Background(), svcCtx)
			got, err := l.FriendPutInHandle(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendPutInHandle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendPutInHandle() got = %v, want %v", got, tt.want)
			}
		})
	}
}
