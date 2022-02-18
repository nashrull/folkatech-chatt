package models

import (
	entity "chattsimulator/chatt/entity"
	"errors"
	"fmt"

	"github.com/bandros/framework"
	"github.com/gin-gonic/gin"
)

func RoomMapToStruct(data map[string]interface{}) (room entity.RoomsEntity) {
	if data["id"] != nil {
		room.ID = data["id"]
	}
	if data["name"] != nil {
		room.Name = data["name"].(string)
	}
	if data["created_by"] != nil {
		room.CreatedBy = data["created_by"].(string)
	}
	if data["date_created"] != nil {
		room.DateCreated = data["date_created"].(string)
	}
	return room
}

func RoomMemberMapToStruct(data map[string]interface{}) (roomMember entity.RoomMembers) {
	if data["id"] != nil {
		roomMember.ID = fmt.Sprint(data["id"])
	}
	if data["room_id"] != nil {
		roomMember.RoomID = fmt.Sprint(data["room_id"])
	}
	if data["date_joined"] != nil {
		roomMember.DateJoined = fmt.Sprint(data["date_joined"])
	}
	if data["user"] != nil {
		roomMember.UserId = fmt.Sprint(data["user"])
	}
	if data["is_admin"] != nil {
		roomMember.IsAdmin = fmt.Sprint(data["is_admin"])
	}

	if data["user_phone"] != nil {
		roomMember.UserPhone = fmt.Sprint(data["user_phone"])
	}
	return roomMember
}

func MessageMapToStruct(data map[string]interface{}) (message entity.MessageEntity) {
	if data["id"] != nil {
		message.Id = fmt.Sprint(data["id"])
	}
	if data["message"] != nil {
		message.Message = fmt.Sprint(data["message"])
	}
	if data["room_member"] != nil {
		message.RoomMember = fmt.Sprint(data["room_member"])
	}
	if data["date_created"] != nil {
		message.Status = fmt.Sprint(data["date_created"])
	}
	if data["user_id"] != nil {
		message.UserId = fmt.Sprint(data["user_id"])
	}
	return message
}

func CekRoomExistenceByTwoMember(c *gin.Context, sender, receiver string) (roomMember entity.RoomMembers, err error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("rm1.room_id").From("room_members rm1")
	db.Join("room_members rm2", "rm2.room_id = rm1.room_id", "")
	db.StartGroup("AND")
	db.Where("rm1.user", sender).Where("rm2.user", receiver)
	db.EndGroup()
	db.StartGroup("OR")
	db.Where("rm2.user", sender).Where("rm1.user", receiver)
	db.EndGroup()

	r, err := db.Row()
	if err != nil {
		return roomMember, err
	}
	if r["room_id"] != nil {
		roomMember.RoomID = fmt.Sprint(r["room_id"])
	}

	return roomMember, nil
}

func CreateRoom(c *gin.Context, data map[string]interface{}) (room entity.RoomsEntity, err error) {
	db := framework.Database{}
	defer db.Close()

	db.From("rooms")
	r, err := db.Insert(data)
	if err != nil {
		return room, err
	}
	room.ID = int(r.(int64))
	data["id"] = room.ID

	return RoomMapToStruct(data), nil
}

func RegisterToMemberToRoom(c *gin.Context, data []map[string]interface{}) (err error) {
	db := framework.Database{}
	defer db.Close()

	db.From("room_members")
	_, err = db.InsertBatch(data)
	return err
}

func GetMyRooms(c *gin.Context, userId string) (MyRooms []entity.RoomMembers, err error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("room_id").From("room_members")
	db.Where("user", userId)
	r, err := db.Result()
	for _, v := range r {
		MyRooms = append(MyRooms, RoomMemberMapToStruct(v))
	}
	return MyRooms, err
}

func GetRoomMembers(c *gin.Context, room_id string) (MyRooms []entity.RoomMembers, err error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("rm.*, u.hp user_phone").From("room_members rm")
	db.Join("user u", "u.id = rm.user", "")
	db.Where("room_id", room_id)
	r, err := db.Result()
	for _, v := range r {
		MyRooms = append(MyRooms, RoomMemberMapToStruct(v))
	}
	return MyRooms, err
}

func GetMemberIDBySenderIdAndRoomID(c *gin.Context, senderID, roomID string) (roomMember entity.RoomMembers, err error) {
	db := framework.Database{}
	defer db.Close()
	db.Select("*").From("room_members")
	db.Where("user", senderID).Where("room_id", roomID)
	r, err := db.Row()
	if err != nil {
		return roomMember, err
	}
	if r["id"] == nil {
		return roomMember, errors.New("anda bukan anggota dari group ini")
	}
	roomMember = RoomMemberMapToStruct(r)
	return roomMember, err
}

func StoreMessage(c *gin.Context, message map[string]interface{}) (messageresponse entity.MessageEntity, err error) {
	db := framework.Database{}
	defer db.Close()
	db.From("message")
	r, e := db.Insert(message)
	if e != nil {
		return messageresponse, err
	}

	messageresponse.Id = fmt.Sprint(r)
	return messageresponse, e
}

func GetMessageByRoom(c *gin.Context, roomid, user_id string) (messageresponse []entity.MessageEntity, err error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("m.id, m.message, m.date_created, m.room_member")
	db.From("message m")
	db.Join("room_members rm", "m.room_member = rm.id", "")
	// db.Join("room_members rm", "rm.user=u.id", "")
	db.Where("rm.room_id", roomid)
	db.Where("rm.user", user_id)

	r, e := db.Result()
	if e != nil {
		return messageresponse, e
	}
	for _, v := range r {
		v["user_id"] = user_id
		messageresponse = append(messageresponse, MessageMapToStruct(v))
	}
	return messageresponse, nil
}

func DeleteUserFromRoomMember(c *gin.Context, id string) error {
	db := framework.Database{}
	defer db.Close()

	db.From("room_members")
	db.Where("id", id)
	e := db.Delete()
	return e
}
