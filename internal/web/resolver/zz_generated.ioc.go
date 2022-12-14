//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package resolver

import (
	"github.com/9d77v/go-ddd-demo/internal/web/generated"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &resolver_{}
		},
	})
	resolverStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Resolver{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(resolverStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &standaloneResolver_{}
		},
	})
	standaloneResolverStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &StandaloneResolver{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(standaloneResolverStructDescriptor)
}

type resolver_ struct {
	Mutation_ func() generated.MutationResolver
	Query_    func() generated.QueryResolver
}

func (r *resolver_) Mutation() generated.MutationResolver {
	return r.Mutation_()
}

func (r *resolver_) Query() generated.QueryResolver {
	return r.Query_()
}

type standaloneResolver_ struct {
	Mutation_ func() generated.MutationResolver
	Query_    func() generated.QueryResolver
}

func (s *standaloneResolver_) Mutation() generated.MutationResolver {
	return s.Mutation_()
}

func (s *standaloneResolver_) Query() generated.QueryResolver {
	return s.Query_()
}

type ResolverIOCInterface interface {
	Mutation() generated.MutationResolver
	Query() generated.QueryResolver
}

type StandaloneResolverIOCInterface interface {
	Mutation() generated.MutationResolver
	Query() generated.QueryResolver
}

var _resolverSDID string

func GetResolverSingleton() (*Resolver, error) {
	if _resolverSDID == "" {
		_resolverSDID = util.GetSDIDByStructPtr(new(Resolver))
	}
	i, err := singleton.GetImpl(_resolverSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*Resolver)
	return impl, nil
}

func GetResolverIOCInterfaceSingleton() (ResolverIOCInterface, error) {
	if _resolverSDID == "" {
		_resolverSDID = util.GetSDIDByStructPtr(new(Resolver))
	}
	i, err := singleton.GetImplWithProxy(_resolverSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ResolverIOCInterface)
	return impl, nil
}

var _standaloneResolverSDID string

func GetStandaloneResolverSingleton() (*StandaloneResolver, error) {
	if _standaloneResolverSDID == "" {
		_standaloneResolverSDID = util.GetSDIDByStructPtr(new(StandaloneResolver))
	}
	i, err := singleton.GetImpl(_standaloneResolverSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*StandaloneResolver)
	return impl, nil
}

func GetStandaloneResolverIOCInterfaceSingleton() (StandaloneResolverIOCInterface, error) {
	if _standaloneResolverSDID == "" {
		_standaloneResolverSDID = util.GetSDIDByStructPtr(new(StandaloneResolver))
	}
	i, err := singleton.GetImplWithProxy(_standaloneResolverSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(StandaloneResolverIOCInterface)
	return impl, nil
}
