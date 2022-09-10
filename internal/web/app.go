package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/9d77v/go-ddd-demo/internal/web/generated"
	"github.com/9d77v/go-ddd-demo/internal/web/resolver"
	"github.com/9d77v/go-ddd-demo/pkg/base"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:paramType=Param
// +ioc:autowire:constructFunc=Init
type App struct {
	base.BaseApp
	Resolver resolver.ResolverIOCInterface `singleton:""`
}

type Param struct {
	base.BaseParam
}

func (p *Param) Init(a *App) (*App, error) {
	a.BaseApp = base.NewBaseApp(&p.BaseParam)
	return a, nil
}

func (a *App) Run() {
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.ServerPort),
		Handler: a.getServerMux(),
	}
	go func() {
		errc <- srv.ListenAndServe()
		log.Printf("connect to http://localhost:%d/ for GraphQL playground", a.ServerPort)
	}()
	log.Printf("exiting (%v)", <-errc)
	srvCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := srv.Shutdown(srvCtx)
	if err != nil {
		log.Println("server shut down error:", err)
	}
	log.Println("exited")
}

func (a *App) getServerMux() *http.ServeMux {
	mux := http.NewServeMux()
	apiHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: a.Resolver,
			},
		),
	)
	mux.Handle("/docs", playground.Handler("GraphQL playground", "/api"))
	mux.Handle("/api", apiHandler)
	return mux
}
