package test

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/internal/logic"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
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
					FriendReqId:  "1843308560872640512",
					UserId:       "1843306302982328320",
					HandleResult: 2,
				},
			},
			want:    &social.FriendPutInHandleResp{},
			wantErr: false,
		},
		{
			name: "name2",
			args: args{
				in: &social.FriendPutInHandleReq{
					FriendReqId:  "1843308564307775488",
					UserId:       "1843306311148638208",
					HandleResult: 2,
				},
			},
			want:    &social.FriendPutInHandleResp{},
			wantErr: false,
		},
		{
			name: "name3",
			args: args{
				in: &social.FriendPutInHandleReq{
					FriendReqId:  "1843308567755493376",
					UserId:       "1843306319776321536",
					HandleResult: 2,
				},
			},
			want:    &social.FriendPutInHandleResp{},
			wantErr: false,
		},
		{
			name: "name3",
			args: args{
				in: &social.FriendPutInHandleReq{
					FriendReqId:  "1843308571182239744",
					UserId:       "1843306319776321536",
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
