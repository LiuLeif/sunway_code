// 2022-05-20 19:25
package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitSession(router *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
}

func Login(c *gin.Context, username string) {
	session := sessions.Default(c)
	session.Set("username", username)
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, // 6 hours
	})
	session.Save()
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("username", "unknown")
	session.Save()
}

func GetLoginUser(c *gin.Context) string {
	session := sessions.Default(c)
	v := session.Get("username")
	if v == nil {
		return "unknown"
	}
	return v.(string)
}
