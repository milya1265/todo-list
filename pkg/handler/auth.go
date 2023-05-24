package handler

import (
	"github.com/gin-gonic/gin"
	todo_app "github.com/milya1265/todo-list.git"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo_app.User
	if err := c.BindJSON(&input); err != nil {
		//logrus.Errorf();
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Fatalf("failing sign up: ", err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
func (h *Handler) signIn(c *gin.Context) {

}
