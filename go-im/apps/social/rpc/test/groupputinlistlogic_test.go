package test

import (
	"context"
	"easy-chat/apps/social/rpc/internal/logic"
	"easy-chat/apps/social/rpc/social"
	"testing"
)

func TestGroupPutinListLogic_GroupPutinList(t *testing.T) {
	type args struct {
		in *social.GroupPutinListReq
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
				in: &social.GroupPutinListReq{
					GroupId: "1841495356521582592",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "name2",
			args: args{
				in: &social.GroupPutinListReq{
					GroupId: "1841496420683616256",
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewGroupPutinListLogic(context.Background(), svcCtx)
			got, err := l.GroupPutinList(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupList() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				t.Logf("GroupList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
