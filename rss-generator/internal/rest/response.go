package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespondMessage(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":      http.StatusText(code),
		"status_code": code,
		"message":     message,
	})
}

func RespondOK(c *gin.Context) {
	body := gin.H{}
	code := http.StatusOK
	body["status"] = http.StatusText(code)
	body["status_code"] = code

	c.JSON(http.StatusOK, body)
}

func RespondOKWithData(c *gin.Context, data interface{}) {
	var body gin.H
	if data == nil {
		body = gin.H{}
	} else {
		dataJson, err := json.Marshal(data)
		if err != nil {
			body = gin.H{"data": dataJson}
		} else {
			err = fmt.Errorf("JSON Marshal Error: %w", err)
			RespondMessage(c, http.StatusInternalServerError, err.Error())
		}
	}
	code := http.StatusOK
	body["status"] = http.StatusText(code)
	body["status_code"] = code

	c.JSON(http.StatusOK, data)
}
