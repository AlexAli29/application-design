package handler

import (
	"applicationDesignTest/internal/application"
	"applicationDesignTest/internal/transport"
	"applicationDesignTest/pkg/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ReservationHandler struct {
	reservationService *application.ReservationService
	logger             transport.Logger
}

func NewReservationHandler(reservationService *application.ReservationService, logger transport.Logger) *ReservationHandler {
	return &ReservationHandler{
		reservationService, logger,
	}
}

func (h *ReservationHandler) InitRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.reserve)

	return r
}

func (h *ReservationHandler) reserve(w http.ResponseWriter, r *http.Request) {
	reservation := application.ReservationRequest{}
	err := helpers.DecodeJSONRequest(r, &reservation)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.reservationService.ReserveRoom(reservation)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error(err.Error())
	}
	h.logger.Debug("Room reserved")
}
