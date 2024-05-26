package application

import (
	"applicationDesignTest/internal/core"
)

type ReservationService struct {
	roomRepository core.RoomRepository
}

func NewReservationService(roomRepository core.RoomRepository) *ReservationService {

	return &ReservationService{roomRepository: roomRepository}
}

func (s *ReservationService) ReserveRoom(request ReservationRequest) error {
	room, err := s.roomRepository.GetRoomById(request.RoomId)

	if err != nil {
		return err
	}

	err = room.Reserve(request.GuestId, request.StartDate, request.EndDate)

	if err != nil {
		return err
	}

	err = s.roomRepository.Save(room)

	if err != nil {
		return err
	}

	return nil
}
