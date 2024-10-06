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
			name: "正常方式申请,同意",
			args: args{
				in: &social.GroupPutInHandleReq{
					GroupReqId:    "1842849198509133824",
					GroupId:       "1842848623063207936",
					HandleUid:     "1842843962742673408",
					HandleResult:  2,
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
