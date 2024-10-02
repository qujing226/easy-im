package test

import (
	"context"
	"easy-chat/apps/social/rpc/internal/logic"
	"easy-chat/apps/social/rpc/social"
	"testing"
)

func TestGroupPutInHandleLogic_GroupPutInHandle(t *testing.T) {
	type args struct {
		in *social.GroupPutInHandleReq
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "正常方式申请,进行拒绝",
			args: args{
				in: &social.GroupPutInHandleReq{
					GroupReqId:    "1841501295475691520",
					GroupId:       "1841495356521582592",
					HandleUid:     "1841486243565670400",
					HandleResult:  3,
					Username:      "",
					UserAvatarUrl: "",
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewGroupPutInHandleLogic(context.Background(), svcCtx)
			got, err := l.GroupPutInHandle(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupList() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				t.Logf("GroupList() got = %v, want %v", got, tt.want)
			}
		})

	}
}
