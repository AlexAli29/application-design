package core

import (
	"applicationDesignTest/pkg/helpers"
	"errors"
	"time"
)

type Room struct {
	Id           int           `json:"id"`
	HotelId      int           `json:"hotelId"`
	Number       string        `json:"number"`
	Reservations []Reservation `json:"reservations"`
	Version      int           `json:"version"`
}

func (r *Room) Reserve(GuestId int, startDate time.Time, endDate time.Time) error {

	now := time.Now()
	if startDate.Before(now) || endDate.Before(now) {
		return errors.New("invalid period")
	}

	for _, res := range r.Reservations {
		if helpers.CheckDateIntersection(startDate, endDate, res.StartDate, res.EndDate) {
			return errors.New("room already reserved for this period")
		}
	}

	newReservation := Reservation{
		Id:        len(r.Reservations) + 1,
		GuestId:   GuestId,
		Status:    Upcoming,
		StartDate: startDate,
		EndDate:   endDate,
	}

	r.Reservations = append(r.Reservations, newReservation)
	return nil
}
