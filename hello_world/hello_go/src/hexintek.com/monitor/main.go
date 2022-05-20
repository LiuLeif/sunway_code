// 2022-05-20 10:20
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("src/hexintek.com/monitor/ui/*")
	router.Static("/assets", "src/hexintek.com/monitor/assets")

	router.GET("/devices/", handleGetDevices)
	router.POST("/devices/", handleInsertDevice)
	router.POST("/upload/", handleUpload)	
	router.Run()
}
