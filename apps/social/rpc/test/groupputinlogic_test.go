package test

import (
	"context"
	"easy-chat/apps/social/rpc/internal/logic"
	"easy-chat/apps/social/rpc/social"
	"testing"
	"time"
)

func TestGroupPutinLogic_GroupPutin(t *testing.T) {
	type args struct {
		in *social.GroupPutinReq
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "管理员邀请进群，期望能直接看到群成员",
			args: args{
				in: &social.GroupPutinReq{
					GroupId:    "1842848623063207936",
					ReqId:      "1842843971269693440",
					ReqMsg:     "我是高端go玩家",
					ReqTime:    time.Now().Unix(),
					JoinSource: 0,
					InviterUid: "1842843962742673408",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "正常方式申请",
			args: args{
				in: &social.GroupPutinReq{
					GroupId:    "1842848623063207936",
					ReqId:      "1842843978811052032",
					ReqMsg:     "我是低端go玩家QWQ",
					ReqTime:    time.Now().Unix(),
					JoinSource: 1,
					InviterUid: "",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "该群不需要验证便可加入",
			args: args{
				in: &social.GroupPutinReq{
					GroupId:    "1842848625386852352",
					ReqId:      "1842843985727459328",
					ReqMsg:     "扫黄大队长请求加入QWQ",
					ReqTime:    time.Now().Unix(),
					JoinSource: 2,
					InviterUid: "",
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewGroupPutinLogic(context.Background(), svcCtx)
			got, err := l.GroupPutin(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupList() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				t.Logf("GroupList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
