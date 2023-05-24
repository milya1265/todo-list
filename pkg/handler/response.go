package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	Messege string `json:"messege"`
}

func newErrorResponse(c *gin.Context, statusCode int, messege string) {
	logrus.Error(messege)
	c.AbortWithStatusJSON(statusCode, error{messege})
}
