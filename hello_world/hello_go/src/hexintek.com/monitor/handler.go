// 2022-05-20 11:03
package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func handleGetDevices(c *gin.Context) {
	imei := strings.TrimSpace(c.Query("imei"))
	vendor := strings.TrimSpace(c.Query("vendor"))
	filter := map[string]interface{}{}
	if len(imei) != 0 {
		filter["imei"] = imei
	}
	if len(vendor) != 0 {
		filter["vendor"] = vendor
	}
	devices := GetDeviceInfo(filter)
	c.HTML(http.StatusOK, "device_list.tmpl", gin.H{"Devices": devices, "filter":filter})
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
