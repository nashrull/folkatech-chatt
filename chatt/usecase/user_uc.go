package usecase

import (
	entity "chattsimulator/chatt/entity"
	"chattsimulator/chatt/models"

	"github.com/gin-gonic/gin"
)

func RegisterUserUseCase(c *gin.Context, user entity.RegisterUser) (registResponse entity.RegisterUserResponse, err error) {
	// call model
	id, err := models.RegisterUser(c, user)
	registResponse.Id = id
	return registResponse, err
}

func FindUserByPhone(c *gin.Context, phone string) (registResponse entity.RegisterUserResponse, err error) {
	user, err := models.FindUserByPhone(c, phone)

	if user["id"] != nil {
		registResponse.Id = user["id"]
		registResponse.Nohp = user["hp"].(string)
	}

	return registResponse, err
}
