// 2022-05-20 11:03
package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func error(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"msg": msg})
	log.Fatal(msg)
}

func showDashBoard(c *gin.Context) {
	if needLogin(c) {
		return
	}
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
	c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{"Devices": devices, "filter": filter, "username": GetLoginUser(c)})
}

func enroll(c *gin.Context) {
	if needLogin(c) {
		return
	}
	imei := strings.TrimSpace(c.PostForm("imei"))
	vendor := strings.TrimSpace(c.PostForm("vendor"))
	InsertDeviceInfo(DeviceInfo{imei, vendor})

	filter := map[string]interface{}{"imei": imei, "vendor": vendor}
	devices := GetDeviceInfo(filter)
	c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{"Devices": devices, "filter": filter})
}

func bulkEnroll(c *gin.Context) {
	if needLogin(c) {
		return
	}
	file, err := c.FormFile("upload_file")
	if err != nil {
		error(c, err.Error())
		return
	}

	if file.Size > 10240000 {
		error(c, "file size execeed 10MB")
		return
	}

	tmpfile, _ := ioutil.TempFile("/tmp", "gin")
	defer os.Remove(tmpfile.Name())
	c.SaveUploadedFile(file, tmpfile.Name())

	f, err := os.Open(tmpfile.Name())
	if err != nil {
		error(c, "internal error")
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		error(c, "bad csv format:"+err.Error())
		return
	}

	devices := []DeviceInfo{}

	for i, line := range data {
		if i == 0 {
			// skip header line
			continue
		}
		device := DeviceInfo{}
		for j, v := range line {
			if j == 0 {
				device.IMEI = v
			} else if j == 1 {
				device.Vendor = v
			} else {

			}
		}
		devices = append(devices, device)
	}

	for _, device := range devices {
		InsertDeviceInfo(device)
	}

	devices = GetDeviceInfo(nil)
	c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{"Devices": devices})
}

func showLogin(c *gin.Context) {
	Logout(c)
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if IsUserValid(username, password) {
		Login(c, username)
		showDashBoard(c)
	} else {
		showLogin(c)
	}
}

func needLogin(c *gin.Context) bool {
	if GetLoginUser(c) == "unknown" {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
		return true
	}
	return false
}
