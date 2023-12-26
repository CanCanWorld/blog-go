package api

import (
	"blog/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userService service.UserService
	var err = c.ShouldBind(&userService)
	if err == nil {
		var res = userService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

func UserLogin(c *gin.Context) {
	var userService service.UserService
	var err = c.ShouldBind(&userService)
	if err == nil {
		var res = userService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
