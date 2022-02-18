package controller

import (
	entity "chattsimulator/chatt/entity"
	"chattsimulator/chatt/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	var CreateRoomForm entity.CreateRoom
	err := c.ShouldBindJSON(&CreateRoomForm)
	if err != nil {
		ErrorResponseJson(c, http.StatusBadRequest, "Kesalahan dalam pengiriman", err)
		return
	}
	result, err := usecase.CreateRoom(c, CreateRoomForm.Destination)
	if err != nil {
		ErrorResponseJson(c, http.StatusBadRequest, "Kesalahan dalam server", err)
		return
	}
	SuccessResponseJson(c, "Berhasil Membuat Room", result)
}
func CreateRoomGroup(c *gin.Context) {
	var CreateRoomForm entity.CreateRoomGroup
	err := c.ShouldBindJSON(&CreateRoomForm)
	if err != nil {
		ErrorResponseJson(c, http.StatusBadRequest, "Kesalahan dalam pengiriman", err)
		return
	}
	result, err := usecase.CreateRoomGroup(c, CreateRoomForm)
	if err != nil {
		ErrorResponseJson(c, http.StatusBadRequest, "Kesalahan dalam server", err)
		return
	}
	SuccessResponseJson(c, "Berhasil Membuat Room", result)
}

func MyRoom(c *gin.Context) {
	result, err := usecase.MyRoom(c)
	if err != nil {
		ErrorResponseJson(c, http.StatusBadRequest, "Kesalahan dalam server", err)
		return
	}
	SuccessResponseJson(c, "Sukses menampilkan data", result)
}

func MyRoomMessage(c *gin.Context) {
	room_id := c.Param("room_id")
	result, err := usecase.MessageInTheRoom(c, room_id)
	if err != nil {
		ErrorResponseJson(c, http.StatusBadRequest, "Kesalahan dalam server", err)
		return
	}
	SuccessResponseJson(c, "Sukses menampilkan data", result)
}

func KickFromRoom(c *gin.Context) {
	room_id := c.Param("room_id")
	usertokick := c.Param("nohp")
	err := usecase.KickUserFromRoom(c, room_id, usertokick)
	if err != nil {
		ErrorResponseJson(c, http.StatusBadRequest, "Kesalahan dalam server", err)
		return
	}
	SuccessResponseJson(c, "Sukses menampilkan data", nil)
}

func RoomMember(c *gin.Context) {
	room_id := c.Param("room_id")
	result, err := usecase.RoomMembers(c, room_id)
	if err != nil {
		ErrorResponseJson(c, http.StatusBadRequest, "Kesalahan dalam server", err)
		return
	}
	SuccessResponseJson(c, "Sukses menampilkan data", result)
}
