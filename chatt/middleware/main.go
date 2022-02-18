package middleware

import (
	entity "chattsimulator/chatt/entity"
	"chattsimulator/chatt/usecase"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	// get user phone/token by header
	header := c.Request.Header.Get("token")
	// cek user by header
	FindUserByPhone, err := usecase.FindUserByPhone(c, header)
	if err != nil {
		c.Abort()
		return
	}

	if (FindUserByPhone == entity.RegisterUserResponse{}) {
		fmt.Println("User tidak terdaftar")
		c.Abort()
		return
	}

	sesion := sessions.Default(c)
	sesion.Set("user_id", FindUserByPhone.Id)
	sesion.Save()
}
