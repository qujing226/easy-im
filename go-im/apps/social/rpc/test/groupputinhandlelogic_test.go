package test

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/internal/logic"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
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
					GroupReqId:    "1843312468886032384",
					GroupId:       "1843311241150337024",
					HandleUid:     "1843306294396588032",
					HandleResult:  2,
					Username:      "",
					UserAvatarUrl: "",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "正常方式申请,同意",
			args: args{
				in: &social.GroupPutInHandleReq{
					GroupReqId:    "1843312469716504576",
					GroupId:       "1843311241150337024",
					HandleUid:     "1843306294396588032",
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
