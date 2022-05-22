// 2022-05-20 19:25
package session

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"hexintek.com/license_server/model"
)

func Init(router *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
}

func Login(c *gin.Context, account model.Account) {
	session := sessions.Default(c)
	session.Set("account", account)
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, // 6 hours
	})
	session.Save()
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("account")
	session.Save()
}

func GetLoginAccount(c *gin.Context) (model.Account, error) {
	session := sessions.Default(c)
	v := session.Get("account")
	if v == nil {
		return model.Account{}, errors.New("not logged in")
	}
	switch v.(type) {
	case model.Account:
		return (v.(model.Account)), nil
	case *model.Account:
		return *(v.(*model.Account)), nil
	}
	return model.Account{}, errors.New("not logged in")
}

func LimitToVendor(c *gin.Context, config map[string]interface{}) map[string]interface{} {
	account, _ := GetLoginAccount(c)
	if config == nil {
		config = map[string]interface{}{}
	}
	config["vendor"] = account.Vendor
	return config
}
