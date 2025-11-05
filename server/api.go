package server

import (
	"base_app/config"
	"base_app/constants/api_constants"
	"base_app/logger"
	"base_app/mongo"
	"base_app/redis"
	"base_app/service"
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/profile"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/urfave/negroni"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func StartAPI(c *cli.Context) error {
	router := service.Router()

	err := setupRedis()
	if err != nil {
		logger.Log.Fatalf("error setting up redis server: %+v", err)
		return err
	}
	err = setUpMongo()
	if err != nil {
		logger.Log.Fatalf("error setting up Mongo: %+v", err)
		return err
	}
	startMemoryProfiling()
	startServer(router)
	return nil
}

func startServer(router *mux.Router) {
	server := negroni.New(negroni.NewRecovery())
	handlerFunc := router.ServeHTTP

	server.Use(httpStatLoggerMiddleware())

	server.UseHandlerFunc(handlerFunc)
	portInfo := ":" + strconv.Itoa(int(config.Port()))

	logger.Log.Info("Starting http server on port ", portInfo)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	corsServer := c.Handler(server)

	apiServer := http.Server{
		Addr:    portInfo,
		Handler: corsServer,
	}

	// Create a context and cancel function for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up a signal handler to catch termination signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// Wait for a signal to shut down
		<-sigChan
		logger.Log.Info("Received termination signal. Shutting down...")

		// Cancel the context to signal the server to shut down
		cancel()
	}()

	// Start the HTTP server in a goroutine, why ? Because it will let the current server to still serve the incoming requests
	go func() {
		if err := apiServer.ListenAndServe(); err != http.ErrServerClosed {
			logger.Log.Errorf("ListenAndServe: %v", err)
		}
	}()

	// Wait for the context to be canceled, indicating a graceful shutdown
	<-ctx.Done()

	// Create a new context with a timeout for forcefully shutting down the server
	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	// Shut down the HTTP server gracefully
	if err := apiServer.Shutdown(shutdownCtx); err != nil {
		logger.Log.Errorf("Error shutting down the server: %v", err)
	}

	logger.Log.Info("Gracefully shutting down the http server")
}

func httpStatLoggerMiddleware() negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		startTime := time.Now()
		ctx := r.Context()

		// Will use the app requestId in complete HTTP request if present
		requestId := r.Header.Get("requestid")
		if requestId == "" {
			requestId = uuid.New().String()
		}
		ctx = context.WithValue(ctx, api_constants.RequestId, requestId)

		r = r.WithContext(ctx)

		logger.Log.WithFields(logrus.Fields{
			api_constants.RequestTime:   startTime.Format(time.RFC3339),
			api_constants.RequestUrl:    r.RequestURI,
			api_constants.RequestMethod: r.Method,
			api_constants.RemoteAddr:    r.RemoteAddr,
			api_constants.RequestId:     requestId,
		}).Debug("Incoming HTTP Request Logs")

		next(rw, r)

		// logs after request is completed
		responseTime := time.Now()
		deltaTime := responseTime.Sub(startTime).Seconds() * 1000

		if r.URL.Path != "/ping" {
			logger.Log.WithFields(logrus.Fields{
				api_constants.RequestId:     requestId,
				api_constants.RequestTime:   startTime.Format(time.RFC3339),
				api_constants.ResponseTime:  responseTime.Format(time.RFC3339),
				api_constants.DeltaTime:     deltaTime,
				api_constants.RequestUrl:    r.URL.Path,
				api_constants.RequestMethod: r.Method,
			}).Debug("Http Logs")
		}
	}
}

func startMemoryProfiling() {
	logContext := logger.BuildContext(api_constants.MemoryProfiling)
	defer profile.Start(profile.MemProfile).Stop()

	go func() {
		err := http.ListenAndServe(":6060", nil)
		if err != nil {
			logger.Log.WithFields(logContext).Errorf("Failed at http server: %+v", err)
		}
	}()

}

func setupRedis() error {
	_, err := redis.InitRedis()
	return err
}

func setUpMongo() error {
	_, err := mongo.InitMongo()
	return err
}
