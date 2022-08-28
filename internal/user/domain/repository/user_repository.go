package repository

import (
	"github.com/9d77v/go-ddd-demo/internal/user/domain/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetDB() *gorm.DB
	GetUserByPhone(string) (*entity.User, error)
}
