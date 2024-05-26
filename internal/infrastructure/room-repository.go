package infrastructure

import (
	"applicationDesignTest/internal/core"
	"errors"
)

type RoomRepository struct {
	rooms []core.Room
}

func NewRoomRepository() *RoomRepository {
	rooms := []core.Room{
		{Id: 1, HotelId: 1, Number: "101", Reservations: []core.Reservation{}, Version: 1},
		{Id: 2, HotelId: 1, Number: "102", Reservations: []core.Reservation{}, Version: 1},
		{Id: 3, HotelId: 1, Number: "103", Reservations: []core.Reservation{}, Version: 1},
		{Id: 4, HotelId: 1, Number: "104", Reservations: []core.Reservation{}, Version: 1},
		{Id: 5, HotelId: 2, Number: "201", Reservations: []core.Reservation{}, Version: 1},
		{Id: 6, HotelId: 2, Number: "202", Reservations: []core.Reservation{}, Version: 1},
		{Id: 7, HotelId: 2, Number: "203", Reservations: []core.Reservation{}, Version: 1},
		{Id: 8, HotelId: 3, Number: "301", Reservations: []core.Reservation{}, Version: 1},
		{Id: 9, HotelId: 3, Number: "302", Reservations: []core.Reservation{}, Version: 1},
		{Id: 10, HotelId: 3, Number: "303", Reservations: []core.Reservation{}, Version: 1},
	}

	return &RoomRepository{rooms: rooms}
}

func (r *RoomRepository) GetRoomById(roomId int) (core.Room, error) {
	for _, room := range r.rooms {
		if room.Id == roomId {
			return room, nil
		}
	}
	return core.Room{}, errors.New("room not found")
}

func (r *RoomRepository) FindRoom(hotelId int, roomNumber string) (core.Room, error) {
	for _, room := range r.rooms {
		if room.HotelId == hotelId && room.Number == roomNumber {
			return room, nil
		}
	}
	return core.Room{}, errors.New("room not found")
}
func (r *RoomRepository) Save(room core.Room) error {

	for i, rm := range r.rooms {
		if rm.Id == room.Id {
			r.rooms[i] = room
			return nil
		}
	}

	return errors.New("room not found")
}
