// 2022-05-20 10:20
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/devices/", handleGetDevices)
	router.PUT("/devices/", handleInsertDevice)
	router.Run()
}
