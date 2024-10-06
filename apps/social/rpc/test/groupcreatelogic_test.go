package test

import (
	"context"
	"easy-chat/apps/social/rpc/internal/logic"
	"easy-chat/apps/social/rpc/social"
	"testing"
)

func TestGroupCreateLogic_GroupCreate(t *testing.T) {
	type args struct {
		in *social.GroupCreateReq
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
				in: &social.GroupCreateReq{
					Name:       "Go语言开发者(4)",
					Status:     0,
					CreatorUid: "1842843962742673408",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				in: &social.GroupCreateReq{
					Name:       "扫黄小分队",
					Status:     0,
					CreatorUid: "1842843971269693440",
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewGroupCreateLogic(context.Background(), svcCtx)
			got, err := l.GroupCreate(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				t.Logf("GroupCreate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
