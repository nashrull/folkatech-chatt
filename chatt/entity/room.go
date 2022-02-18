package request

type CreateRoom struct {
	Destination string `json:"destination"`
}

type CreateRoomGroup struct {
	GroupName   string   `json:"group_name"`
	Destination []string `json:"destinations"`
}

type RoomsEntity struct {
	ID          interface{} `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	CreatedBy   string      `json:"created_by,omitempty"`
	DateCreated string      `json:"date_created,omitempty"`
}

type RegisterToRomMember struct {
	Users  []string    `json:"users"`
	RoomID interface{} `json:"room_id"`
}

type RoomMembers struct {
	ID         string      `json:"id,omitempty"`
	RoomID     interface{} `json:"room_id,omitempty"`
	IsAdmin    string      `json:"is_admin,omitempty"`
	UserId     interface{} `json:"user_id,omitempty"`
	UserPhone  string      `json:"user_phone,omitempty"`
	DateJoined string      `json:"date_joined,omitempty"`
}
