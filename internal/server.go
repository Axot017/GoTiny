package internal

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	openapi "github.com/go-openapi/runtime/middleware"
	"go.uber.org/fx"

	"gotiny/internal/api"
	"gotiny/internal/api/handler"
	"gotiny/internal/api/util"
	"gotiny/internal/core"
	"gotiny/internal/core/usecase"
	"gotiny/internal/data"
	"gotiny/internal/data/adapter"
)

func StartServer() {
	fx.New(
		fx.Provide(
			newServer,
			fx.Annotate(
				NewConfig,
				fx.As(new(usecase.CreateShortLinkConfig)),
				fx.As(new(core.LoggingConfig)),
				fx.As(new(adapter.DynamodbConfig)),
			),
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

func newMux(
	handlers []api.RouteHandler,
	redirect_hander *handler.RedirectHandler,
	logger *util.StructuredLogger,
) *chi.Mux {
	mux := chi.NewRouter()

	redocOpts := openapi.RedocOpts{
		SpecURL:  "swagger.yaml",
		BasePath: "/api",
		Path:     "/docs",
	}
	redoc := openapi.Redoc(redocOpts, nil)

	swatterUIOpts := openapi.SwaggerUIOpts{
		SpecURL:  "swagger.yaml",
		BasePath: "/api",
		Path:     "/swagger-ui",
	}
	swaggerUI := openapi.SwaggerUI(swatterUIOpts, nil)

	mux.Use(middleware.RealIP)
	mux.Use(middleware.RequestLogger(logger))
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Recoverer)

	mux.Route("/api", func(r chi.Router) {
		r.Method(http.MethodGet, "/docs", redoc)
		r.Method(http.MethodGet, "/swagger-ui", swaggerUI)
		r.Method(http.MethodGet, "/swagger.yaml", http.FileServer(http.Dir("./")))

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
