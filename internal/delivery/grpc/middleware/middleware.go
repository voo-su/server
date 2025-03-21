package middleware

type Middleware struct {
	Auth    *AuthMiddleware
	Logging *LoggingMiddleware
}
