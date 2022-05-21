// 2022-05-20 10:20
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	InitSession(router)
	InitModel()

	router.Static("/assets", "src/hexintek.com/monitor/assets")
	router.LoadHTMLGlob("src/hexintek.com/monitor/ui/*")

	// for web client
	router.GET("/login", showLogin)
	router.POST("/login", login)

	router.GET("/dashboard/", showDashBoard)
	router.POST("/enroll/", enroll)
	router.POST("/bulk_enroll/", bulkEnroll)	

	router.Run()
}
