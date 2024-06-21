package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/caarlos0/env/v11"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

// Route is a http route.
type Route struct {
	Method      string
	Path        string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}

type HttpServer struct {
	Echo   *echo.Echo
	config *Config
	log    zerolog.Logger
}

func New() *HttpServer {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		panic(fmt.Sprintf("failed to parse config: %v", err))
	}

	var (
		e      = echo.New()
		logger = zerolog.New(zerolog.NewConsoleWriter())
	)

	// set request logger
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))

	// set CORS headers if Origin is present
	e.Use(middleware.CORS())

	e.HideBanner = true
	e.HidePort = true
	//e.Logger.SetLevel(log.OFF)

	return &HttpServer{
		Echo:   e,
		config: &cfg,
		log:    logger,
	}
}

func (s *HttpServer) AddRoute(route *Route) {
	s.Echo.Add(route.Method, route.Path, route.Handler, route.Middlewares...)
}

func (s *HttpServer) InitRoutes(rg RouteConfigurator) {
	for _, route := range rg.GetRoutes() {
		s.AddRoute(&route)
	}
}

func Start(lc fx.Lifecycle, srv *HttpServer) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				on := fmt.Sprintf("%s:%s", srv.config.Host, srv.config.Port)

				srv.log.Info().Msgf("starting server on %s", on)

				go func() {
					err := srv.Echo.Start(on)

					if errors.Is(err, http.ErrServerClosed) {
						srv.log.Info().Err(err).Msgf("server %s stopped", on)
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				err := srv.Echo.Shutdown(ctx)
				if err != nil {
					srv.log.Info().Err(err).Msg("couldn't stop server")
				}

				return nil
			},
		},
	)
}

func InitAndRun() fx.Option {
	return fx.Module("http-server",
		// Init http server
		fx.Provide(New),
		fx.Invoke(
			// Registration routes and handlers for http server
			InitHandlerGroups,
			// Run HTTP server
			Start,
		),
	)
}
