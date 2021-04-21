package v1

import (
	"gin-app/pkg/e"
	"gin-app/pkg/logging"
	"gin-app/pkg/util"
	"gin-app/service"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

var Auth auth

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func (*auth) GetAuth(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := service.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
