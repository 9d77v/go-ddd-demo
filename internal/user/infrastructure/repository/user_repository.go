package repository

import (
	"github.com/9d77v/go-ddd-demo/internal/user/domain/entity"
	"github.com/9d77v/go-ddd-demo/internal/user/domain/repository"
	"github.com/9d77v/go-ddd-demo/pkg/db/postgres"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type UserRepository struct {
	PgDB postgres.PgDBIOCInterface `normal:",pg_local"`
}

var _ repository.UserRepository = &UserRepository{}

func (r *UserRepository) GetUserByPhone(phone string) (*entity.User, error) {
	user := &entity.User{}
	err := r.GetDB().Where("phone=?", phone).First(user).Error
	return user, err
}

func (r *UserRepository) GetDB() *gorm.DB {
	return r.PgDB.GetDB()
}
