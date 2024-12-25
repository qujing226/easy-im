package models

import (
	"time"
)

type User struct {
	ID        string    `gorm:"type:varchar(24);primary_key"`
	Avatar    string    `gorm:"type:varchar(255);default:'https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png'"`
	Nickname  string    `gorm:"type:varchar(24);not null"`
	Phone     string    `gorm:"type:varchar(24);not null"`
	Email     string    `gorm:"type:varchar(24);"`
	Password  string    `gorm:"type:varchar(191);"`
	Status    *int8     `gorm:"type:tinyint(1);default:0"`
	Sex       *int8     `gorm:"type:tinyint(1);default:0"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
}
