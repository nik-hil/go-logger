package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type logData struct {
	Data       string `json:"data"`
	StatusCode int    `json:"statusCode"`
	Level      string `json:"level"`
}

func postLog(c *gin.Context) {
	var newLog logData

	if err := c.BindJSON(&newLog); err != nil {
		return
	}
	var statusCode int
	var data string

	statusCode = newLog.StatusCode
	data = newLog.Data

	switch level := strings.ToUpper(newLog.Level); level {
	case "INFO":

		InfoLogger.Println(newLog.Data)
	case "WARNING":

		WarningLogger.Println(newLog.Data)
	case "ERROR":
		ErrorLogger.Println(newLog.Data)
	default:
		statusCode = http.StatusBadRequest
		data = "Incorrect Level"
	}
	if len(http.StatusText(newLog.StatusCode)) == 0 {
		data = "Incorrect status code"
	}
	c.JSON(statusCode, data)
}
func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	router := gin.Default()
	router.POST("/log", postLog)
	router.GET("/ping", getPing)

	router.Run(":8080")
}
