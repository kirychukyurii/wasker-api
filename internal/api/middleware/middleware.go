package middleware

import "go.uber.org/fx"

// Module Middleware exported
var Module = fx.Options(
	fx.Provide(NewLoggerMiddleware),
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewAuthMiddleware),
	fx.Provide(New),
)

// Middleware interface
type Middleware interface {
	Setup()
}

// Middlewares contains multiple middleware
type Middlewares []Middleware

// New creates new middlewares
// Register the middleware that should be applied directly (globally)
func New(loggerMiddleware LoggerMiddleware, corsMiddleware CorsMiddleware, authMiddleware AuthMiddleware) Middlewares {
	return Middlewares{
		loggerMiddleware,
		corsMiddleware,
		authMiddleware,
	}
}

func (a Middlewares) Setup() {
	for _, middleware := range a {
		middleware.Setup()
	}
}
