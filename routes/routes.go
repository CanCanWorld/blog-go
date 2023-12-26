package routes

import (
	"blog/api"
	"blog/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	var r = gin.Default()
	var store = cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("my_session", store))
	var v1 = r.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		var authed = v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("task", api.CreateTask)
		}
	}
	return r
}
