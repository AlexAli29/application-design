package core

import "time"

type ReservationStatus int

const (
	Closed ReservationStatus = iota
	Ongoing
	Upcoming
)

type Reservation struct {
	Id        int               `json:"id"`
	GuestId   int               `json:"guestId"`
	Status    ReservationStatus `json:"status"`
	StartDate time.Time         `json:"startDate"`
	EndDate   time.Time         `json:"endDate"`
}
