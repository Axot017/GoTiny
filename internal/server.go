package internal

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"

	"gotiny/internal/api"
	"gotiny/internal/api/handler"
	"gotiny/internal/core"
	"gotiny/internal/core/usecase"
	"gotiny/internal/data"
)

func StartServer() {
	fx.New(
		fx.Provide(
			newServer,
			fx.Annotate(NewConfig, fx.As(new(usecase.CreateShortLinkConfig))),
			fx.Annotate(
				newMux,
				fx.ParamTags(`group:"routes"`),
			),
		),
		fx.Provide(api.Providers()...),
		fx.Provide(data.Providers()...),
		fx.Provide(core.Providers()...),
		fx.Invoke(startServer),
	).Run()
}

func startServer(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", server.Addr)
			go server.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}

func newMux(handlers []api.RouteHandler, redirect_hander *handler.RedirectHandler) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Route("/api", func(r chi.Router) {
		for _, h := range handlers {
			r.Method(h.Method(), h.Path(), h)
		}
	})

	mux.Method(redirect_hander.Method(), redirect_hander.Path(), redirect_hander)

	return mux
}

func newServer(mux *chi.Mux) *http.Server {
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return &server
}
