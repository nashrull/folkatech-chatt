package usecase

import (
	entity "chattsimulator/chatt/entity"
	"chattsimulator/chatt/models"
	"errors"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context, destination string) (room entity.RoomsEntity, err error) {
	sesion := sessions.Default(c)
	senderID := sesion.Get("user_id")
	// cek if receiver is not found
	receiver, err := FindUserByPhone(c, destination)
	if err != nil {
		return room, err
	}
	if receiver.Id == nil {
		return room, errors.New("receiver not found")
	}

	// cek if room name is existed or not
	roomMember, err := CekRoom(c, senderID.(string), receiver.Id.(string))
	if err != nil {
		return room, err
	}
	// fmt.Println(room)
	if roomMember.RoomID != nil {
		return room, errors.New("room pernah dibuat sebelumnya")
	}
	// create room
	// room name is sender_id + receiver_id
	name1 := senderID.(string) + "_" + receiver.Id.(string)
	var roomInpt entity.RoomsEntity
	roomInpt.CreatedBy = senderID.(string)
	roomInpt.Name = name1
	room, err = StoreRoom(c, roomInpt)
	if err != nil {
		return room, err
	}
	// REGISTER AS ROOM MEMBER
	var RegisterRoomMember []entity.RoomMembers
	RegisterRoomMember = append(RegisterRoomMember, entity.RoomMembers{
		UserId:  senderID,
		IsAdmin: "0",
		RoomID:  room.ID,
	})
	RegisterRoomMember = append(RegisterRoomMember, entity.RoomMembers{
		UserId:  receiver.Id,
		IsAdmin: "0",
		RoomID:  room.ID,
	})
	err = RegisterMembersToRoom(c, RegisterRoomMember)
	return room, err
}

func CreateRoomGroup(c *gin.Context, roomGroup entity.CreateRoomGroup) (room entity.RoomsEntity, err error) {
	sesion := sessions.Default(c)
	senderID := sesion.Get("user_id")

	var roomInpt entity.RoomsEntity
	roomInpt.CreatedBy = senderID.(string)
	roomInpt.Name = roomGroup.GroupName
	room, err = StoreRoom(c, roomInpt)
	if err != nil {
		return room, err
	}
	// REGISTER AS ROOM MEMBER
	var RegisterRoomMember []entity.RoomMembers
	RegisterRoomMember = append(RegisterRoomMember, entity.RoomMembers{
		IsAdmin: "1",
		UserId:  senderID,
		RoomID:  room.ID,
	})
	for _, v := range roomGroup.Destination {
		// get user id by phone
		userID, err := FindUserByPhone(c, v)
		if err != nil || userID.Id == nil {
			fmt.Println("This user is not found =>", v)
			fmt.Println("Or error =>", err.Error())
		}
		var tmp entity.RoomMembers
		tmp.IsAdmin = "0"
		tmp.UserId = userID.Id
		tmp.RoomID = room.ID
		RegisterRoomMember = append(RegisterRoomMember, tmp)
	}
	err = RegisterMembersToRoom(c, RegisterRoomMember)
	return room, err
}

func StoreRoom(c *gin.Context, roomInpt entity.RoomsEntity) (room entity.RoomsEntity, err error) {
	data := map[string]interface{}{}
	data["name"] = roomInpt.Name
	data["created_by"] = roomInpt.CreatedBy
	room, err = models.CreateRoom(c, data)
	return room, err
}

func CekRoom(c *gin.Context, sender, receiver string) (room entity.RoomMembers, err error) {
	room, err = models.CekRoomExistenceByTwoMember(c, sender, receiver)
	return room, err
}

func RegisterMembersToRoom(c *gin.Context, inpt []entity.RoomMembers) (err error) {
	var data []map[string]interface{}
	for _, v := range inpt {
		t := map[string]interface{}{}
		t["is_admin"] = v.IsAdmin
		t["room_id"] = v.RoomID
		t["user"] = v.UserId
		data = append(data, t)
	}
	err = models.RegisterToMemberToRoom(c, data)
	return err
}

func MyRoom(c *gin.Context) (MyRooms []entity.RoomMembers, err error) {
	sesion := sessions.Default(c)
	senderID := sesion.Get("user_id")
	MyRooms, err = models.GetMyRooms(c, senderID.(string))
	return MyRooms, err
}

func RoomMembers(c *gin.Context, room_id string) (MyRooms []entity.RoomMembers, err error) {
	MyRooms, err = models.GetRoomMembers(c, room_id)
	return MyRooms, err
}

func GetMemberIDBySenderIdAndRoomID(c *gin.Context, senderID, roomID string) (roomMember entity.RoomMembers, err error) {
	roomMember, err = models.GetMemberIDBySenderIdAndRoomID(c, senderID, roomID)
	return roomMember, err
}

func KickUserFromRoom(c *gin.Context, room_id, usertokick string) (err error) {
	sesion := sessions.Default(c)
	senderID := sesion.Get("user_id")

	// cek if sender is admin or not
	sender, err := GetMemberIDBySenderIdAndRoomID(c, senderID.(string), room_id)
	if err != nil {
		return err
	}
	fmt.Println(sender.IsAdmin)
	if sender.IsAdmin != "1" {
		return errors.New("anda bukan admin group ini")
	}

	// cek user to kick id
	usertokickid, err := FindUserByPhone(c, usertokick)
	if err != nil {
		return err
	}
	tokickuser, err := GetMemberIDBySenderIdAndRoomID(c, usertokickid.Id.(string), room_id)
	if err != nil {
		return err
	}
	return DeleteUserFromRoom(c, tokickuser.ID)
}

func DeleteUserFromRoom(c *gin.Context, idmember string) error {
	return models.DeleteUserFromRoomMember(c, idmember)
}
