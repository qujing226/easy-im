package test

import (
	"context"
	"easy-chat/apps/social/rpc/internal/logic"
	"easy-chat/apps/social/rpc/social"
	"testing"
)

func TestGroupUsersLogic_GroupUsers(t *testing.T) {
	type args struct {
		in *social.GroupUsersReq
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
				in: &social.GroupUsersReq{
					GroupId: "1842848623063207936",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "name2",
			args: args{
				in: &social.GroupUsersReq{
					GroupId: "1842848625386852352",
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewGroupUsersLogic(context.Background(), svcCtx)
			got, err := l.GroupUsers(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupList() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				t.Logf("GroupList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
