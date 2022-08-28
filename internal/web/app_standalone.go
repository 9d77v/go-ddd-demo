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
	"github.com/9d77v/go-pkg/env"
)

const defaultStandalonePort = 8080

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type StandaloneApp struct {
	Resolver resolver.StandaloneResolverIOCInterface `singleton:""`
}

func (a *StandaloneApp) Run() {
	port := env.Int("PORT")
	if port == 0 {
		port = defaultStandalonePort
	}
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: a.getServerMux(),
	}
	go func() {
		errc <- srv.ListenAndServe()
		log.Printf("connect to http://localhost:%d/ for GraphQL playground", port)
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

func (a *StandaloneApp) getServerMux() *http.ServeMux {
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
