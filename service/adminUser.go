package service

import (
	"gin-app/model"
	"gin-app/pkg/util"
)

func FindAdmUserById(id int) (err error, admUser model.AdminUser) {
	err = model.GetDB().Where("`id` = ?", id).First(&admUser).Error
	return err, admUser
}

func CheckAuth(username, password string) bool {
	var admUser model.AdminUser

	password = util.MD5([]byte(password))

	model.GetDB().Select("id").Where("user_name = ? AND password = ?", username, password).First(&admUser)

	if admUser.ID > 0 {
		return true
	}

	return false
}
