package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mskydream/ashyq/api/model"
)

func (h *Handler) signUp(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{IsSuccess: false, Message: "Введенные данные некорректны", Data: err})
		return
	}

	status, err := h.service.CreateUser(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Ошибка в сервере", Data: err})
		return
	}

	c.JSON(http.StatusOK, model.Response{IsSuccess: true, Message: "Успешно создано", Data: status})
}

func (h *Handler) signIn(c *gin.Context) {
	var input model.SignInInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Введенные данные некорректны", Data: err})
		return
	}

	token, err := h.service.GenerateToken(input.IIN, input.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{IsSuccess: false, Message: "Веденные данные не корректны.", Data: err})
		return
	}

	c.JSON(http.StatusOK, model.Response{IsSuccess: true, Message: "Вход выполнено", Data: token})
}
