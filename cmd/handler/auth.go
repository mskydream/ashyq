package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mskydream/qr-code/cmd/model"
)

func (h *Handler) signUp(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{IsSuccess: false, Message: "Введенные данные некорректны", Data: err})
		return
	}
	// check
	fmt.Println("BornDate in handler", input.BornDate)
	id, err := h.service.CreateUser(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Ошибка в сервере", Data: err})
	} else {

		c.JSON(http.StatusOK, model.Response{IsSuccess: true, Message: "Успешно создано", Data: id})
	}
}
