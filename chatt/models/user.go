package models

import (
	entity "chattsimulator/chatt/entity"

	"github.com/bandros/framework"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context, inpt entity.RegisterUser) (interface{}, error) {
	db := framework.Database{}
	defer db.Close()

	db.From("user")
	return db.Insert(map[string]interface{}{
		"hp": inpt.Nohp,
	})
}

func FindUserByPhone(c *gin.Context, phone string) (result map[string]interface{}, err error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("*").From("user")
	db.Where("hp", phone)
	r, e := db.Row()

	return r, e
}
