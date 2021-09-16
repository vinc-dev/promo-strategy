package container

import (
	"net/http"

	"github.com/vinc-dev/promo-strategy/transport/http/middleware"
)

// MiddlewareContainer handle all middleware used in project
type MiddlewareContainer struct {
	Cors func(next http.HandlerFunc) http.HandlerFunc
}

// CreateMiddlewareContainer construct all middlewares used in the app
func CreateMiddlewareContainer() *MiddlewareContainer {
	return &MiddlewareContainer{
		Cors: middleware.Cors(),
	}
}
