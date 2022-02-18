package request

import "time"

type SentMessage struct {
	RoomID  string `json:"room_id"`
	Message string `json:"message"`
}

type SentMessageResponse struct {
	RoomID        string    `json:"room_id"`
	MemberID      string    `json:"member_id"`
	MessageID     string    `json:"message_id"`
	DestinationID string    `json:"destination_id"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	DateCreated   time.Time `json:"date_created"`
}

type MessageEntity struct {
	Id          string `json:"id"`
	Message     string `json:"message"`
	RoomMember  string `json:"room_member"`
	UserId      string `json:"user_id"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
}
