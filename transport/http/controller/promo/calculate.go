package promo

import (
	"encoding/json"
	"net/http"

	"github.com/vinc-dev/promo-strategy/internal/domain/promo"
)

// Checkout...
func Checkout(service promo.ServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := r.Context()
		var body struct {
			ProductNames []string `json:"productNames"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); nil != err {
			response(w, struct {
				Error string `json:"error"`
			}{
				Error: err.Error(),
			}, http.StatusBadRequest)
			return
		}
		result, err := service.Checkout(ctx, body.ProductNames)
		if nil != err {
			response(w, struct {
				Error string `json:"error"`
			}{
				Error: err.Error(),
			}, http.StatusBadRequest)
			return
		}
		response(w, struct {
			Total float64 `json:"total"`
		}{
			Total: result,
		}, http.StatusOK)
	}
}

func response(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}
