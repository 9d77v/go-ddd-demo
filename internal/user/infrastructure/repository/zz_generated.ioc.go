//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package repository

import (
	"github.com/9d77v/go-ddd-demo/internal/user/domain/entity"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
	"gorm.io/gorm"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &userRepository_{}
		},
	})
	userRepositoryStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &UserRepository{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(userRepositoryStructDescriptor)
}

type userRepository_ struct {
	GetUserByPhone_ func(phone string) (*entity.User, error)
	GetDB_          func() *gorm.DB
}

func (u *userRepository_) GetUserByPhone(phone string) (*entity.User, error) {
	return u.GetUserByPhone_(phone)
}

func (u *userRepository_) GetDB() *gorm.DB {
	return u.GetDB_()
}

type UserRepositoryIOCInterface interface {
	GetUserByPhone(phone string) (*entity.User, error)
	GetDB() *gorm.DB
}

var _userRepositorySDID string

func GetUserRepositorySingleton() (*UserRepository, error) {
	if _userRepositorySDID == "" {
		_userRepositorySDID = util.GetSDIDByStructPtr(new(UserRepository))
	}
	i, err := singleton.GetImpl(_userRepositorySDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*UserRepository)
	return impl, nil
}

func GetUserRepositoryIOCInterfaceSingleton() (UserRepositoryIOCInterface, error) {
	if _userRepositorySDID == "" {
		_userRepositorySDID = util.GetSDIDByStructPtr(new(UserRepository))
	}
	i, err := singleton.GetImplWithProxy(_userRepositorySDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(UserRepositoryIOCInterface)
	return impl, nil
}
