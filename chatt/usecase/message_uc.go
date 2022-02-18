package usecase

import (
	entity "chattsimulator/chatt/entity"
	"chattsimulator/chatt/models"
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SentMessageUsecase(c *gin.Context, SentMessageForm entity.SentMessage) (response entity.MessageEntity, err error) {
	s := sessions.Default(c)
	user_id := s.Get("user_id").(string)

	// cek if rooms is existed or not
	memberID, err := GetMemberIDBySenderIdAndRoomID(c, user_id, SentMessageForm.RoomID)
	if err != nil {
		return response, err
	}

	if memberID.ID == "" {
		return response, errors.New("anda bukan bagian dari room ini")
	}

	// sent message
	var messageInpt entity.MessageEntity
	messageInpt.RoomMember = memberID.ID
	messageInpt.Message = SentMessageForm.Message

	message, err := StoreMessage(c, messageInpt)
	if err != nil {
		return message, err
	}
	message.UserId = user_id
	message.Message = SentMessageForm.Message
	message.RoomMember = messageInpt.RoomMember
	return message, err
}

func StoreMessage(c *gin.Context, message entity.MessageEntity) (messageresponse entity.MessageEntity, err error) {
	data := map[string]interface{}{}
	data["message"] = message.Message
	data["room_member"] = message.RoomMember
	messageresponse, err = models.StoreMessage(c, data)
	return messageresponse, err
}

func MessageInTheRoom(c *gin.Context, roomid string) (response []entity.MessageEntity, err error) {
	s := sessions.Default(c)
	user_id := s.Get("user_id").(string)

	response, err = models.GetMessageByRoom(c, roomid, user_id)
	return response, err
}
