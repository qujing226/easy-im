package models

import (
	"github.com/peninsula12/easy-im/go-im/pkg/status"
	"time"
)

type Friend struct {
	ID        string    `gorm:"primaryKey;autoIncrement"`
	UserID    string    `gorm:"type:varchar(64);not null"`
	FriendUID string    `gorm:"type:varchar(64);not null"`
	Remark    string    `gorm:"type:varchar(255)"`
	AddSource *int8     `gorm:"type:tinyint"`
	CreatedAt time.Time `gorm:"type:timestamp"`
}

type FriendRequest struct {
	ID           string               `gorm:"primaryKey;autoIncrement"`
	UserID       string               `gorm:"type:varchar(64);not null"`
	ReqUID       string               `gorm:"type:varchar(64);not null"`
	ReqMsg       string               `gorm:"type:varchar(255)"`
	ReqTime      time.Time            `gorm:"type:timestamp;not null"`
	HandleResult status.HandlerResult `gorm:"type:tinyint"`
	HandleMsg    string               `gorm:"type:varchar(255)"`
	HandledAt    time.Time            `gorm:"type:timestamp"`
}

type Group struct {
	ID              string    `gorm:"type:varchar(24);primaryKey"`
	Name            string    `gorm:"type:varchar(255);not null"`
	Icon            string    `gorm:"type:varchar(255);not null;default:'https://c-ssl.duitang.com/uploads/item/201802/24/20180224083913_yhrX2.jpeg'"`
	Status          *int8     `gorm:"type:tinyint"`
	CreatorUID      string    `gorm:"type:varchar(64);not null"`
	GroupType       *int8     `gorm:"type:tinyint;not null"`
	IsVerify        bool      `gorm:"type:boolean;not null"`
	Notification    string    `gorm:"type:varchar(255)"`
	NotificationUID string    `gorm:"type:varchar(64)"`
	CreatedAt       time.Time `gorm:"type:timestamp"`
	UpdatedAt       time.Time `gorm:"type:timestamp"`
}

type GroupMember struct {
	ID          string                      `gorm:"primaryKey;autoIncrement"`
	GroupID     string                      `gorm:"type:varchar(64);not null"`
	UserID      string                      `gorm:"type:varchar(64);not null"`
	RoleLevel   status.GroupMemberRoleLevel `gorm:"type:tinyint;not null"`
	JoinTime    time.Time                   `gorm:"type:timestamp"`
	JoinSource  *int8                       `gorm:"type:tinyint;default:0"`
	InviterUID  string                      `gorm:"type:varchar(64)"`
	OperatorUID string                      `gorm:"type:varchar(64)"`
}

type GroupRequest struct {
	ID            string               `gorm:"primaryKey;autoIncrement"`
	ReqID         string               `gorm:"type:varchar(64);not null"`
	GroupID       string               `gorm:"type:varchar(64);not null"`
	ReqMsg        string               `gorm:"type:varchar(255)"`
	ReqTime       time.Time            `gorm:"type:timestamp"`
	JoinSource    *int8                `gorm:"type:tinyint"`
	InviterUserID string               `gorm:"type:varchar(64)"`
	HandleUserID  string               `gorm:"type:varchar(64)"`
	HandleTime    time.Time            `gorm:"type:timestamp"`
	HandleResult  status.HandlerResult `gorm:"type:tinyint"`
}
