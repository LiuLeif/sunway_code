// 2022-05-20 10:20
package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"hexintek.com/license_server/model"
	"hexintek.com/license_server/session"
)

func reportError(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"msg": msg})
	log.Fatal(msg)
}

func needLogin(c *gin.Context) bool {
	if _, err := session.GetLoginAccount(c); err != nil {
		c.Redirect(http.StatusFound, "/login")
		return true
	}
	return false
}

func handleShowDashboard(c *gin.Context) {
	if needLogin(c) {
		return
	}
	imei := strings.TrimSpace(c.Query("imei"))
	version := strings.TrimSpace(c.Query("version"))
	filter := map[string]interface{}{}
	_, isExport := c.GetQuery("export")
	if len(imei) != 0 {
		filter["imei"] = imei
	}
	if len(version) != 0 {
		filter["version"] = version
	}
	devices := model.GetDeviceInfo(session.LimitToVendor(c, filter))

	if isExport {
		exportCSV(c, devices)
		return
	}
	account, _ := session.GetLoginAccount(c)
	c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{"Devices": devices, "filter": filter, "username": account.Username, "count": len(devices)})
}

func exportCSV(c *gin.Context, devices []model.DeviceInfo) {
	tmpfile, _ := ioutil.TempFile("/tmp", "gin")
	defer os.Remove(tmpfile.Name())

	file, err := os.Create(tmpfile.Name())
	if err != nil {
		reportError(c, "failed create tmpfile ")
		return
	}
	writer := csv.NewWriter(file)
	for _, device := range devices {
		writer.Write(device.ToCSV())
	}
	writer.Flush()

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "text")
	c.Header("Content-Disposition", "attachment; filename=devices.txt")
	c.Header("Content-Type", "application/text")
	c.File(tmpfile.Name())
}

func handleEnroll(c *gin.Context) {
	if needLogin(c) {
		return
	}
	imei := strings.TrimSpace(c.PostForm("imei"))
	version := strings.TrimSpace(c.PostForm("version"))

	account, _ := session.GetLoginAccount(c)
	model.InsertDeviceInfo(model.DeviceInfo{imei, version, account.Vendor})

	filter := map[string]interface{}{"imei": imei, "version": version}
	devices := model.GetDeviceInfo(session.LimitToVendor(c, filter))
	c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{"Devices": devices, "filter": filter})
}

func handleBulkEnroll(c *gin.Context) {
	if needLogin(c) {
		return
	}
	file, err := c.FormFile("upload_file")
	if err != nil {
		reportError(c, err.Error())
		return
	}

	if file.Size > 10240000 {
		reportError(c, "file size execeed 10MB")
		return
	}

	account, _ := session.GetLoginAccount(c)

	tmpfile, _ := ioutil.TempFile("/tmp", "gin")
	defer os.Remove(tmpfile.Name())
	c.SaveUploadedFile(file, tmpfile.Name())

	f, err := os.Open(tmpfile.Name())
	if err != nil {
		reportError(c, "internal error")
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		reportError(c, "bad csv format:"+err.Error())
		return
	}

	devices := []model.DeviceInfo{}

	for i, line := range data {
		if i == 0 {
			// skip header line
			continue
		}
		device := model.DeviceInfo{}
		device.Vendor = account.Vendor
		for j, v := range line {
			if j == 0 {
				device.IMEI = v
			} else if j == 1 {
				device.Version = v
			} else {

			}
		}
		devices = append(devices, device)
	}

	for _, device := range devices {
		model.InsertDeviceInfo(device)
	}

	devices = model.GetDeviceInfo(session.LimitToVendor(c, nil))
	c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{"Devices": devices})
}

func handleShowLogin(c *gin.Context) {
	session.Logout(c)
	retry, _ := c.Get("retry")
	c.HTML(http.StatusOK, "login.tmpl", gin.H{"retry": retry})
}

func handleLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if account, err := model.GetAccount(username, password); err == nil {
		session.Login(c, account)
		handleShowDashboard(c)
	} else {
		c.Set("retry", true)
		handleShowLogin(c)
	}
}

func main() {
	router := gin.Default()

	session.Init(router)
	model.Init()

	router.Static("/assets", "src/hexintek.com/license_server/assets")
	router.LoadHTMLGlob("src/hexintek.com/license_server/ui/*")

	// for web client
	router.GET("/login", handleShowLogin)
	router.POST("/login", handleLogin)

	router.GET("/dashboard/", handleShowDashboard)
	router.POST("/enroll/", handleEnroll)
	router.POST("/bulk_enroll/", handleBulkEnroll)

	router.Run()
}
