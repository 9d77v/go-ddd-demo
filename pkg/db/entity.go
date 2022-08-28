package db

import (
	"time"
)

// Entity 基础模型
type Entity struct {
	ID        string `gorm:"primarykey"`
	CreateBy  string
	CreatedAt *time.Time
	UpdateBy  string
	UpdatedAt *time.Time
	DeleteBy  string
	DeletedAt *time.Time `gorm:"index"`
}
