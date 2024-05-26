package core

type RoomRepository interface {
	GetRoomById(roomId int) (Room, error)
	FindRoom(hotelId int, roomNumber string) (Room, error)
	Save(room Room) error
}
