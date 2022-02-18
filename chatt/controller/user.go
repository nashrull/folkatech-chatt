package controller

import (
	entity "chattsimulator/chatt/entity"
	"chattsimulator/chatt/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var RegisterUserForm entity.RegisterUser
	err := c.ShouldBindJSON(&RegisterUserForm)
	if err != nil {
		ErrorResponseJson(c, http.StatusBadRequest, "Kesalahan dalam pengiriman", err)
		return
	}

	// set usecase register user
	result, err := usecase.RegisterUserUseCase(c, RegisterUserForm)
	if err != nil {
		ErrorResponseJson(c, http.StatusBadRequest, "Kesalahan dalam penyimpanan data", err)
		return
	}

	SuccessResponseJson(c, "Berhasil register user", result)
}

func SentMessage(c *gin.Context) {
	var SentMessageForm entity.SentMessage
	err := c.ShouldBindJSON(&SentMessageForm)
	if err != nil {
		ErrorResponseJson(c, http.StatusBadRequest, "Kesalahan dalam pengiriman", err)
		return
	}
	result, err := usecase.SentMessageUsecase(c, SentMessageForm)
	if err != nil {
		ErrorResponseJson(c, http.StatusBadRequest, "Kesalahan dalam server", err)
		return
	}
	SuccessResponseJson(c, "Berhasil Mengirim Pesan", result)
}
