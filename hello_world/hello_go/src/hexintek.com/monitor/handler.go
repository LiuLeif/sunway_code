// 2022-05-20 11:03
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetDevices(c *gin.Context) {
	devices := GetDeviceInfo(nil)
	c.JSON(http.StatusOK, gin.H{"devices": devices})
}

func handleInsertDevice(c *gin.Context) {
	var device DeviceInfo
	if err := c.ShouldBindJSON(&device); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	InsertDeviceInfo(device)
	c.JSON(http.StatusOK, gin.H{})
}
