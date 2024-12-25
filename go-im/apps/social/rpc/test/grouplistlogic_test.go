package test

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/internal/logic"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
	"testing"
)

func TestGroupListLogic_GroupList(t *testing.T) {
	type args struct {
		in *social.GroupListReq
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				in: &social.GroupListReq{
					UserId: "1841486189794693120",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				in: &social.GroupListReq{
					UserId: "1841486243565670400",
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewGroupListLogic(context.Background(), svcCtx)
			got, err := l.GroupList(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupList() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				t.Logf("GroupList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
