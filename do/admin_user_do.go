package do

import (
	"free/datasource"
	"free/models"
)

type AdminUser interface {
	Find() (adminuser []models.Admin_user)
}

func NewAdminUser() AdminUser {
	return &adminUser{}
}

type adminUser struct{}

var db = datasource.GetDB()

func (u adminUser) Find() (adminuser []models.Admin_user) {
	db.Where("id = 1").First(&adminuser)
	// log.Print(user.Rows())
	return
}
