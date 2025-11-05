package service

import (
	"base_app/auth"
	"base_app/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/swaggo/files"
	httpSwagger "github.com/swaggo/http-swagger"
)

func MetricsHandler() http.Handler {

	storeMetricsGatherer := prometheus.NewRegistry()

	multiGatherer := prometheus.Gatherers{
		prometheus.DefaultGatherer,
		storeMetricsGatherer,
	}

	return promhttp.InstrumentMetricHandler(
		prometheus.DefaultRegisterer, promhttp.HandlerFor(multiGatherer, promhttp.HandlerOpts{}),
	)
}

func Router() *mux.Router {
	router := mux.NewRouter()

	// Handlers initialization
	apiHandler := ApiHandler()

	router.HandleFunc("/ping", PingHandler()).Methods("GET")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	groupV1 := router.PathPrefix("/api/v1").Subrouter()
	router.PathPrefix("/metrics").Handler(MetricsHandler()).Methods("GET")
	groupV1.Use(auth.Auth)
	groupV1.Use(PanicRecoveryMiddleware)

	groupV1.HandleFunc("/results", apiHandler.GetResults()).Methods("POST")

	return router
}

func PanicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				msg := fmt.Sprintf("panic occured in api: %s with trace: %v ", rec, string(debug.Stack()))
				logger.Log.Error(msg)

				decoder := json.NewDecoder(r.Body)
				decoder.DisallowUnknownFields()

				var searchRequest map[string]interface{}
				err := decoder.Decode(&searchRequest)

				logger.Log.Infof("Raw Request %+v, Error: %+v", searchRequest, err)

				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
