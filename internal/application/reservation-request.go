package application

import "time"

type ReservationRequest struct {
	GuestId   int       `json:"guestId"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	RoomId    int       `json:"roomId"`
}
