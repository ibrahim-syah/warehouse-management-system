package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"warehouse-management-system/middleware"
	databaseutils "warehouse-management-system/utils/database"
	"warehouse-management-system/utils/loggerutils"
	validatorutils "warehouse-management-system/utils/validator"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func StartApp() {
	db, err := databaseutils.ConnectDB()
	if err != nil {
		loggerutils.LoggerSingleton.Fatalf("unable to connect to the database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			loggerutils.LoggerSingleton.Fatal()
		}
	}()

	validatorutils.SetupValidator()

	appRepositories := SetupRepositories(db)
	appUsecases := SetupUsecases(appRepositories)
	appHandlers := SetupHandler(appUsecases)

	s := setupServer(appHandlers)
	startServer(s)
	shutdownServer(s)
}

func setupServer(handlers *appHandlers) *http.Server {
	env := viper.GetString("ENV")
	if env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))
	r.Use(middleware.LoggerMiddleware)
	// r.NoRoute(handler.NotFoundHandler)

	SetupRouter(r, handlers)

	s := &http.Server{
		Addr:        fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT")),
		Handler:     r,
		ReadTimeout: time.Duration(viper.GetInt("GIN_TIMEOUT_SECONDS")) * time.Second,
		// WriteTimeout: time.Duration(viper.GetInt("GIN_TIMEOUT_SECONDS")) * time.Second,
	}
	return s
}

func startServer(s *http.Server) {
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			loggerutils.LoggerSingleton.Fatalf("Server error: %v", err)
		}
	}()
}

func shutdownServer(s *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	loggerutils.LoggerSingleton.Infof("shutdown server...")

	timeoutSeconds := viper.GetInt("GIN_TIMEOUT_SECONDS")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		loggerutils.LoggerSingleton.Fatalf("Server shutdown error: %v", err)
	}

	<-ctx.Done()

	loggerutils.LoggerSingleton.Infof("server exited gracefully")
}
