package main

import (
	"applicationDesignTest/internal/application"
	"applicationDesignTest/internal/infrastructure"
	"applicationDesignTest/internal/infrastructure/logger"
	"applicationDesignTest/internal/transport/rest/handler"
	"context"
	"os"
	"os/signal"
	"syscall"

	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()

	logger := logger.NewLogger(logger.DEBUG)

	roomRepository := infrastructure.NewRoomRepository()

	reservationService := application.NewReservationService(roomRepository)

	reservationHandler := handler.NewReservationHandler(reservationService, logger)

	router.Mount("/reservation", reservationHandler.InitRouter())

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("failed to start server")
		}
	}()

	logger.Info("server started")

	<-done
	logger.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("failed to stop server:%s", err.Error())

		return
	}

	logger.Info("server stopped")
}
