//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package impl

import (
	contextx "context"
	"github.com/9d77v/go-ddd-demo/internal/web/model"
	"github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &userServiceImpl_{}
		},
	})
	userServiceImplStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &UserServiceImpl{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(userServiceImplStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &userServiceStandaloneImpl_{}
		},
	})
	userServiceStandaloneImplStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &UserServiceStandaloneImpl{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(userServiceStandaloneImplStructDescriptor)
}

type userServiceImpl_ struct {
	UserInfo_ func(ctx contextx.Context, id string) (*model.User, error)
	UserPage_ func(ctx contextx.Context, page, size int) (*model.UserConnection, error)
}

func (u *userServiceImpl_) UserInfo(ctx contextx.Context, id string) (*model.User, error) {
	return u.UserInfo_(ctx, id)
}

func (u *userServiceImpl_) UserPage(ctx contextx.Context, page, size int) (*model.UserConnection, error) {
	return u.UserPage_(ctx, page, size)
}

type userServiceStandaloneImpl_ struct {
	UserInfo_ func(ctx contextx.Context, id string) (*model.User, error)
	UserPage_ func(ctx contextx.Context, page, size int) (*model.UserConnection, error)
}

func (u *userServiceStandaloneImpl_) UserInfo(ctx contextx.Context, id string) (*model.User, error) {
	return u.UserInfo_(ctx, id)
}

func (u *userServiceStandaloneImpl_) UserPage(ctx contextx.Context, page, size int) (*model.UserConnection, error) {
	return u.UserPage_(ctx, page, size)
}

type UserServiceImplIOCInterface interface {
	UserInfo(ctx contextx.Context, id string) (*model.User, error)
	UserPage(ctx contextx.Context, page, size int) (*model.UserConnection, error)
}

type UserServiceStandaloneImplIOCInterface interface {
	UserInfo(ctx contextx.Context, id string) (*model.User, error)
	UserPage(ctx contextx.Context, page, size int) (*model.UserConnection, error)
}

var _userServiceImplSDID string

func GetUserServiceImplSingleton() (*UserServiceImpl, error) {
	if _userServiceImplSDID == "" {
		_userServiceImplSDID = util.GetSDIDByStructPtr(new(UserServiceImpl))
	}
	i, err := singleton.GetImpl(_userServiceImplSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*UserServiceImpl)
	return impl, nil
}

func GetUserServiceImplIOCInterfaceSingleton() (UserServiceImplIOCInterface, error) {
	if _userServiceImplSDID == "" {
		_userServiceImplSDID = util.GetSDIDByStructPtr(new(UserServiceImpl))
	}
	i, err := singleton.GetImplWithProxy(_userServiceImplSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(UserServiceImplIOCInterface)
	return impl, nil
}

var _userServiceStandaloneImplSDID string

func GetUserServiceStandaloneImplSingleton() (*UserServiceStandaloneImpl, error) {
	if _userServiceStandaloneImplSDID == "" {
		_userServiceStandaloneImplSDID = util.GetSDIDByStructPtr(new(UserServiceStandaloneImpl))
	}
	i, err := singleton.GetImpl(_userServiceStandaloneImplSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*UserServiceStandaloneImpl)
	return impl, nil
}

func GetUserServiceStandaloneImplIOCInterfaceSingleton() (UserServiceStandaloneImplIOCInterface, error) {
	if _userServiceStandaloneImplSDID == "" {
		_userServiceStandaloneImplSDID = util.GetSDIDByStructPtr(new(UserServiceStandaloneImpl))
	}
	i, err := singleton.GetImplWithProxy(_userServiceStandaloneImplSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(UserServiceStandaloneImplIOCInterface)
	return impl, nil
}
