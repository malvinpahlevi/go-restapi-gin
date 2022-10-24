package main

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type ResponseError struct {
	Status  int
	Message []string
	Error   []string
}

func SendResponse(c *gin.Context, response ResponseError) {
	if len(response.Message) > 0 {
		c.JSON(response.Status, map[string]interface{}{"message": strings.Join(response.Message, "; ")})
	} else if len(response.Error) > 0 {
		c.JSON(response.Status, map[string]interface{}{"error": strings.Join(response.Error, "; ")})
	}
}

// to use
// Unauthorized Access
//SendResponse(c, ResponseError{Status: http.StatusUnauthorized, Error: []string{"Username and password do not match"}})
// Wrong parameters
//SendResponse(c, ResponseError{Status: http.StatusBadRequest, Error: []string{"One or more params are wrong"}})
