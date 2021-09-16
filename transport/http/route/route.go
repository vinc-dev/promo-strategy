package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vinc-dev/promo-strategy/transport/container"
	"github.com/vinc-dev/promo-strategy/transport/http/controller/promo"
)

// Compile compile the data http endpoint and middleware
func Compile(app *container.AppContainer) *mux.Router {
	r := mux.NewRouter()
	r.Use(
		mux.CORSMethodMiddleware(r),
	)
	r.HandleFunc("/checkout", promo.Checkout(app.Services.Promo)).Methods(http.MethodPost)
	http.Handle("/", r)
	return r
}
